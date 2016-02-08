/*
Copyright 2014 The Kubernetes Authors All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package app does all of the work necessary to create a Kubernetes
// APIServer by binding together the API, master and APIServer infrastructure.
// It can be configured and called directly or via the hyperkube framework.
package app

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"

	systemd "github.com/coreos/go-systemd/daemon"
	"github.com/golang/glog"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"k8s.io/kubernetes/cmd/kube-apiserver/app/options"
	"k8s.io/kubernetes/pkg/admission"
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/latest"
	"k8s.io/kubernetes/pkg/api/meta"
	"k8s.io/kubernetes/pkg/api/unversioned"
	apiutil "k8s.io/kubernetes/pkg/api/util"
	"k8s.io/kubernetes/pkg/apis/extensions"
	"k8s.io/kubernetes/pkg/apiserver"
	"k8s.io/kubernetes/pkg/apiserver/authenticator"
	"k8s.io/kubernetes/pkg/capabilities"
	client "k8s.io/kubernetes/pkg/client/unversioned"
	"k8s.io/kubernetes/pkg/cloudprovider"
	serviceaccountcontroller "k8s.io/kubernetes/pkg/controller/serviceaccount"
	"k8s.io/kubernetes/pkg/genericapiserver"
	kubeletclient "k8s.io/kubernetes/pkg/kubelet/client"
	"k8s.io/kubernetes/pkg/master"
	"k8s.io/kubernetes/pkg/serviceaccount"
	"k8s.io/kubernetes/pkg/storage"
	etcdstorage "k8s.io/kubernetes/pkg/storage/etcd"
	"k8s.io/kubernetes/pkg/util"
)

// NewAPIServerCommand creates a *cobra.Command object with default parameters
func NewAPIServerCommand() *cobra.Command {
	s := options.NewAPIServer()
	s.AddFlags(pflag.CommandLine)
	cmd := &cobra.Command{
		Use: "kube-apiserver",
		Long: `The Kubernetes API server validates and configures data
for the api objects which include pods, services, replicationcontrollers, and
others. The API Server services REST operations and provides the frontend to the
cluster's shared state through which all other components interact.`,
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	return cmd
}

// TODO: Longer term we should read this from some config store, rather than a flag.
func verifyClusterIPFlags(s *options.APIServer) {
	if s.ServiceClusterIPRange.IP == nil {
		glog.Fatal("No --service-cluster-ip-range specified")
	}
	var ones, bits = s.ServiceClusterIPRange.Mask.Size()
	if bits-ones > 20 {
		glog.Fatal("Specified --service-cluster-ip-range is too large")
	}
}

type newEtcdFunc func([]string, meta.VersionInterfacesFunc, string, string) (storage.Interface, error)

func newEtcd(etcdServerList []string, interfacesFunc meta.VersionInterfacesFunc, storageGroupVersionString, pathPrefix string) (etcdStorage storage.Interface, err error) {
	if storageGroupVersionString == "" {
		return etcdStorage, fmt.Errorf("storageVersion is required to create a etcd storage")
	}
	storageVersion, err := unversioned.ParseGroupVersion(storageGroupVersionString)
	if err != nil {
		return nil, err
	}

	var storageConfig etcdstorage.EtcdConfig
	storageConfig.ServerList = etcdServerList
	storageConfig.Prefix = pathPrefix
	versionedInterface, err := interfacesFunc(storageVersion)
	if err != nil {
		return nil, err
	}
	storageConfig.Codec = versionedInterface.Codec
	return storageConfig.NewStorage()
}

// convert to a map between group and groupVersions.
func generateStorageVersionMap(legacyVersion string, storageVersions string) map[string]string {
	storageVersionMap := map[string]string{}
	if legacyVersion != "" {
		storageVersionMap[""] = legacyVersion
	}
	if storageVersions != "" {
		groupVersions := strings.Split(storageVersions, ",")
		for _, gv := range groupVersions {
			storageVersionMap[apiutil.GetGroup(gv)] = gv
		}
	}
	return storageVersionMap
}

// parse the value of --etcd-servers-overrides and update given storageDestinations.
func updateEtcdOverrides(overrides []string, storageVersions map[string]string, prefix string, storageDestinations *genericapiserver.StorageDestinations, newEtcdFn newEtcdFunc) {
	if len(overrides) == 0 {
		return
	}
	for _, override := range overrides {
		tokens := strings.Split(override, "#")
		if len(tokens) != 2 {
			glog.Errorf("invalid value of etcd server overrides: %s", override)
			continue
		}

		apiresource := strings.Split(tokens[0], "/")
		if len(apiresource) != 2 {
			glog.Errorf("invalid resource definition: %s", tokens[0])
		}
		group := apiresource[0]
		resource := apiresource[1]

		apigroup, err := latest.Group(group)
		if err != nil {
			glog.Errorf("invalid api group %s: %v", group, err)
			continue
		}
		if _, found := storageVersions[apigroup.GroupVersion.Group]; !found {
			glog.Errorf("Couldn't find the storage version for group %s", apigroup.GroupVersion.Group)
			continue
		}

		servers := strings.Split(tokens[1], ";")
		etcdOverrideStorage, err := newEtcdFn(servers, apigroup.InterfacesFor, storageVersions[apigroup.GroupVersion.Group], prefix)
		if err != nil {
			glog.Fatalf("Invalid storage version or misconfigured etcd for %s: %v", tokens[0], err)
		}

		storageDestinations.AddStorageOverride(group, resource, etcdOverrideStorage)
	}
}

// Run runs the specified APIServer.  This should never exit.
func Run(s *options.APIServer) error {
	verifyClusterIPFlags(s)

	// If advertise-address is not specified, use bind-address. If bind-address
	// is not usable (unset, 0.0.0.0, or loopback), we will use the host's default
	// interface as valid public addr for master (see: util#ValidPublicAddrForMaster)
	if s.AdvertiseAddress == nil || s.AdvertiseAddress.IsUnspecified() {
		hostIP, err := util.ValidPublicAddrForMaster(s.BindAddress)
		if err != nil {
			glog.Fatalf("Unable to find suitable network address.error='%v' . "+
				"Try to set the AdvertiseAddress directly or provide a valid BindAddress to fix this.", err)
		}
		s.AdvertiseAddress = hostIP
	}
	glog.Infof("Will report %v as public IP address.", s.AdvertiseAddress)

	if len(s.EtcdServerList) == 0 {
		glog.Fatalf("--etcd-servers must be specified")
	}

	if s.KubernetesServiceNodePort > 0 && !s.ServiceNodePortRange.Contains(s.KubernetesServiceNodePort) {
		glog.Fatalf("Kubernetes service port range %v doesn't contain %v", s.ServiceNodePortRange, (s.KubernetesServiceNodePort))
	}

	capabilities.Initialize(capabilities.Capabilities{
		AllowPrivileged: s.AllowPrivileged,
		// TODO(vmarmol): Implement support for HostNetworkSources.
		PrivilegedSources: capabilities.PrivilegedSources{
			HostNetworkSources: []string{},
			HostPIDSources:     []string{},
			HostIPCSources:     []string{},
		},
		PerConnectionBandwidthLimitBytesPerSec: s.MaxConnectionBytesPerSec,
	})

	cloud, err := cloudprovider.InitCloudProvider(s.CloudProvider, s.CloudConfigFile)
	if err != nil {
		glog.Fatalf("Cloud provider could not be initialized: %v", err)
	}

	// Setup tunneler if needed
	var tunneler master.Tunneler
	var proxyDialerFn apiserver.ProxyDialerFunc
	if len(s.SSHUser) > 0 {
		// Get ssh key distribution func, if supported
		var installSSH master.InstallSSHKey
		if cloud != nil {
			if instances, supported := cloud.Instances(); supported {
				installSSH = instances.AddSSHKeyToAllInstances
			}
		}

		// Set up the tunneler
		tunneler = master.NewSSHTunneler(s.SSHUser, s.SSHKeyfile, installSSH)

		// Use the tunneler's dialer to connect to the kubelet
		s.KubeletConfig.Dial = tunneler.Dial
		// Use the tunneler's dialer when proxying to pods, services, and nodes
		proxyDialerFn = tunneler.Dial
	}

	// Proxying to pods and services is IP-based... don't expect to be able to verify the hostname
	proxyTLSClientConfig := &tls.Config{InsecureSkipVerify: true}

	kubeletClient, err := kubeletclient.NewStaticKubeletClient(&s.KubeletConfig)
	if err != nil {
		glog.Fatalf("Failure to start kubelet client: %v", err)
	}

	apiGroupVersionOverrides, err := parseRuntimeConfig(s)
	if err != nil {
		glog.Fatalf("error in parsing runtime-config: %s", err)
	}

	clientConfig := &client.Config{
		Host: net.JoinHostPort(s.InsecureBindAddress.String(), strconv.Itoa(s.InsecurePort)),
	}
	if len(s.DeprecatedStorageVersion) != 0 {
		gv, err := unversioned.ParseGroupVersion(s.DeprecatedStorageVersion)
		if err != nil {
			glog.Fatalf("error in parsing group version: %s", err)
		}
		clientConfig.GroupVersion = &gv
	}

	client, err := client.New(clientConfig)
	if err != nil {
		glog.Fatalf("Invalid server address: %v", err)
	}

	legacyV1Group, err := latest.Group(api.GroupName)
	if err != nil {
		return err
	}

	storageDestinations := genericapiserver.NewStorageDestinations()

	storageVersions := generateStorageVersionMap(s.DeprecatedStorageVersion, s.StorageVersions)
	if _, found := storageVersions[legacyV1Group.GroupVersion.Group]; !found {
		glog.Fatalf("Couldn't find the storage version for group: %q in storageVersions: %v", legacyV1Group.GroupVersion.Group, storageVersions)
	}
	etcdStorage, err := newEtcd(s.EtcdServerList, legacyV1Group.InterfacesFor, storageVersions[legacyV1Group.GroupVersion.Group], s.EtcdPathPrefix)
	if err != nil {
		glog.Fatalf("Invalid storage version or misconfigured etcd: %v", err)
	}
	storageDestinations.AddAPIGroup("", etcdStorage)

	if !apiGroupVersionOverrides["extensions/v1beta1"].Disable {
		expGroup, err := latest.Group(extensions.GroupName)
		if err != nil {
			glog.Fatalf("Extensions API is enabled in runtime config, but not enabled in the environment variable KUBE_API_VERSIONS. Error: %v", err)
		}
		if _, found := storageVersions[expGroup.GroupVersion.Group]; !found {
			glog.Fatalf("Couldn't find the storage version for group: %q in storageVersions: %v", expGroup.GroupVersion.Group, storageVersions)
		}
		expEtcdStorage, err := newEtcd(s.EtcdServerList, expGroup.InterfacesFor, storageVersions[expGroup.GroupVersion.Group], s.EtcdPathPrefix)
		if err != nil {
			glog.Fatalf("Invalid extensions storage version or misconfigured etcd: %v", err)
		}
		storageDestinations.AddAPIGroup(extensions.GroupName, expEtcdStorage)
	}

	updateEtcdOverrides(s.EtcdServersOverrides, storageVersions, s.EtcdPathPrefix, &storageDestinations, newEtcd)

	n := s.ServiceClusterIPRange

	// Default to the private server key for service account token signing
	if s.ServiceAccountKeyFile == "" && s.TLSPrivateKeyFile != "" {
		if authenticator.IsValidServiceAccountKeyFile(s.TLSPrivateKeyFile) {
			s.ServiceAccountKeyFile = s.TLSPrivateKeyFile
		} else {
			glog.Warning("No RSA key provided, service account token authentication disabled")
		}
	}

	var serviceAccountGetter serviceaccount.ServiceAccountTokenGetter
	if s.ServiceAccountLookup {
		// If we need to look up service accounts and tokens,
		// go directly to etcd to avoid recursive auth insanity
		serviceAccountGetter = serviceaccountcontroller.NewGetterFromStorageInterface(etcdStorage)
	}

	authenticator, err := authenticator.New(authenticator.AuthenticatorConfig{
		BasicAuthFile:             s.BasicAuthFile,
		ClientCAFile:              s.ClientCAFile,
		TokenAuthFile:             s.TokenAuthFile,
		OIDCIssuerURL:             s.OIDCIssuerURL,
		OIDCClientID:              s.OIDCClientID,
		OIDCCAFile:                s.OIDCCAFile,
		OIDCUsernameClaim:         s.OIDCUsernameClaim,
		ServiceAccountKeyFile:     s.ServiceAccountKeyFile,
		ServiceAccountLookup:      s.ServiceAccountLookup,
		ServiceAccountTokenGetter: serviceAccountGetter,
		KeystoneURL:               s.KeystoneURL,
		PamPassword:               s.PamPassword,
	})

	if err != nil {
		glog.Fatalf("Invalid Authentication Config: %v", err)
	}

	authorizationModeNames := strings.Split(s.AuthorizationMode, ",")
	authorizer, err := apiserver.NewAuthorizerFromAuthorizationConfig(authorizationModeNames, s.AuthorizationPolicyFile)
	if err != nil {
		glog.Fatalf("Invalid Authorization Config: %v", err)
	}

	admissionControlPluginNames := strings.Split(s.AdmissionControl, ",")
	admissionController := admission.NewFromPlugins(client, admissionControlPluginNames, s.AdmissionControlConfigFile)

	if len(s.ExternalHost) == 0 {
		// TODO: extend for other providers
		if s.CloudProvider == "gce" {
			instances, supported := cloud.Instances()
			if !supported {
				glog.Fatalf("GCE cloud provider has no instances.  this shouldn't happen. exiting.")
			}
			name, err := os.Hostname()
			if err != nil {
				glog.Fatalf("Failed to get hostname: %v", err)
			}
			addrs, err := instances.NodeAddresses(name)
			if err != nil {
				glog.Warningf("Unable to obtain external host address from cloud provider: %v", err)
			} else {
				for _, addr := range addrs {
					if addr.Type == api.NodeExternalIP {
						s.ExternalHost = addr.Address
					}
				}
			}
		}
	}

	config := &master.Config{
		Config: &genericapiserver.Config{
			StorageDestinations:       storageDestinations,
			StorageVersions:           storageVersions,
			ServiceClusterIPRange:     &n,
			EnableLogsSupport:         s.EnableLogsSupport,
			EnableUISupport:           true,
			EnableSwaggerSupport:      true,
			EnableProfiling:           s.EnableProfiling,
			EnableWatchCache:          s.EnableWatchCache,
			EnableIndex:               true,
			APIPrefix:                 s.APIPrefix,
			APIGroupPrefix:            s.APIGroupPrefix,
			CorsAllowedOriginList:     s.CorsAllowedOriginList,
			ReadWritePort:             s.SecurePort,
			PublicAddress:             s.AdvertiseAddress,
			Authenticator:             authenticator,
			SupportsBasicAuth:         len(s.BasicAuthFile) > 0 || s.PamPassword,
			Authorizer:                authorizer,
			AdmissionControl:          admissionController,
			APIGroupVersionOverrides:  apiGroupVersionOverrides,
			MasterServiceNamespace:    s.MasterServiceNamespace,
			MasterCount:               s.MasterCount,
			ExternalHost:              s.ExternalHost,
			MinRequestTimeout:         s.MinRequestTimeout,
			ProxyDialer:               proxyDialerFn,
			ProxyTLSClientConfig:      proxyTLSClientConfig,
			ServiceNodePortRange:      s.ServiceNodePortRange,
			KubernetesServiceNodePort: s.KubernetesServiceNodePort,
		},
		EnableCoreControllers: true,
		EventTTL:              s.EventTTL,
		KubeletClient:         kubeletClient,

		Tunneler: tunneler,
	}
	m := master.New(config)

	// We serve on 2 ports.  See docs/accessing_the_api.md
	secureLocation := ""
	if s.SecurePort != 0 {
		secureLocation = net.JoinHostPort(s.BindAddress.String(), strconv.Itoa(s.SecurePort))
	}
	insecureLocation := net.JoinHostPort(s.InsecureBindAddress.String(), strconv.Itoa(s.InsecurePort))

	// See the flag commentary to understand our assumptions when opening the read-only and read-write ports.

	var sem chan bool
	if s.MaxRequestsInFlight > 0 {
		sem = make(chan bool, s.MaxRequestsInFlight)
	}

	longRunningRE := regexp.MustCompile(s.LongRunningRequestRE)
	longRunningTimeout := func(req *http.Request) (<-chan time.Time, string) {
		// TODO unify this with apiserver.MaxInFlightLimit
		if longRunningRE.MatchString(req.URL.Path) || req.URL.Query().Get("watch") == "true" {
			return nil, ""
		}
		return time.After(time.Minute), ""
	}

	if secureLocation != "" {
		handler := apiserver.TimeoutHandler(m.Handler, longRunningTimeout)
		secureServer := &http.Server{
			Addr:           secureLocation,
			Handler:        apiserver.MaxInFlightLimit(sem, longRunningRE, apiserver.RecoverPanics(handler)),
			MaxHeaderBytes: 1 << 20,
			TLSConfig: &tls.Config{
				// Change default from SSLv3 to TLSv1.0 (because of POODLE vulnerability)
				MinVersion: tls.VersionTLS10,
			},
		}

		if len(s.ClientCAFile) > 0 {
			clientCAs, err := util.CertPoolFromFile(s.ClientCAFile)
			if err != nil {
				glog.Fatalf("Unable to load client CA file: %v", err)
			}
			// Populate PeerCertificates in requests, but don't reject connections without certificates
			// This allows certificates to be validated by authenticators, while still allowing other auth types
			secureServer.TLSConfig.ClientAuth = tls.RequestClientCert
			// Specify allowed CAs for client certificates
			secureServer.TLSConfig.ClientCAs = clientCAs
		}

		glog.Infof("Serving securely on %s", secureLocation)
		if s.TLSCertFile == "" && s.TLSPrivateKeyFile == "" {
			s.TLSCertFile = path.Join(s.CertDirectory, "apiserver.crt")
			s.TLSPrivateKeyFile = path.Join(s.CertDirectory, "apiserver.key")
			// TODO (cjcullen): Is PublicAddress the right address to sign a cert with?
			alternateIPs := []net.IP{config.ServiceReadWriteIP}
			alternateDNS := []string{"kubernetes.default.svc", "kubernetes.default", "kubernetes"}
			// It would be nice to set a fqdn subject alt name, but only the kubelets know, the apiserver is clueless
			// alternateDNS = append(alternateDNS, "kubernetes.default.svc.CLUSTER.DNS.NAME")
			if err := util.GenerateSelfSignedCert(config.PublicAddress.String(), s.TLSCertFile, s.TLSPrivateKeyFile, alternateIPs, alternateDNS); err != nil {
				glog.Errorf("Unable to generate self signed cert: %v", err)
			} else {
				glog.Infof("Using self-signed cert (%s, %s)", s.TLSCertFile, s.TLSPrivateKeyFile)
			}
		}

		go func() {
			defer util.HandleCrash()
			for {
				// err == systemd.SdNotifyNoSocket when not running on a systemd system
				if err := systemd.SdNotify("READY=1\n"); err != nil && err != systemd.SdNotifyNoSocket {
					glog.Errorf("Unable to send systemd daemon successful start message: %v\n", err)
				}
				if err := secureServer.ListenAndServeTLS(s.TLSCertFile, s.TLSPrivateKeyFile); err != nil {
					glog.Errorf("Unable to listen for secure (%v); will try again.", err)
				}
				time.Sleep(15 * time.Second)
			}
		}()
	}
	handler := apiserver.TimeoutHandler(m.InsecureHandler, longRunningTimeout)
	http := &http.Server{
		Addr:           insecureLocation,
		Handler:        apiserver.RecoverPanics(handler),
		MaxHeaderBytes: 1 << 20,
	}
	if secureLocation == "" {
		// err == systemd.SdNotifyNoSocket when not running on a systemd system
		if err := systemd.SdNotify("READY=1\n"); err != nil && err != systemd.SdNotifyNoSocket {
			glog.Errorf("Unable to send systemd daemon successful start message: %v\n", err)
		}
	}
	glog.Infof("Serving insecurely on %s", insecureLocation)
	glog.Fatal(http.ListenAndServe())
	return nil
}

func getRuntimeConfigValue(s *options.APIServer, apiKey string, defaultValue bool) bool {
	flagValue, ok := s.RuntimeConfig[apiKey]
	if ok {
		if flagValue == "" {
			return true
		}
		boolValue, err := strconv.ParseBool(flagValue)
		if err != nil {
			glog.Fatalf("Invalid value of %s: %s, err: %v", apiKey, flagValue, err)
		}
		return boolValue
	}
	return defaultValue
}

// Parses the given runtime-config and formats it into map[string]ApiGroupVersionOverride
func parseRuntimeConfig(s *options.APIServer) (map[string]genericapiserver.APIGroupVersionOverride, error) {
	// "api/all=false" allows users to selectively enable specific api versions.
	disableAllAPIs := false
	allAPIFlagValue, ok := s.RuntimeConfig["api/all"]
	if ok && allAPIFlagValue == "false" {
		disableAllAPIs = true
	}

	// "api/legacy=false" allows users to disable legacy api versions.
	disableLegacyAPIs := false
	legacyAPIFlagValue, ok := s.RuntimeConfig["api/legacy"]
	if ok && legacyAPIFlagValue == "false" {
		disableLegacyAPIs = true
	}
	_ = disableLegacyAPIs // hush the compiler while we don't have legacy APIs to disable.

	// "api/v1={true|false} allows users to enable/disable v1 API.
	// This takes preference over api/all and api/legacy, if specified.
	disableV1 := disableAllAPIs
	v1GroupVersion := "api/v1"
	disableV1 = !getRuntimeConfigValue(s, v1GroupVersion, !disableV1)
	apiGroupVersionOverrides := map[string]genericapiserver.APIGroupVersionOverride{}
	if disableV1 {
		apiGroupVersionOverrides[v1GroupVersion] = genericapiserver.APIGroupVersionOverride{
			Disable: true,
		}
	}

	// "extensions/v1beta1={true|false} allows users to enable/disable the extensions API.
	// This takes preference over api/all, if specified.
	disableExtensions := disableAllAPIs
	extensionsGroupVersion := "extensions/v1beta1"
	// TODO: Make this a loop over all group/versions when there are more of them.
	disableExtensions = !getRuntimeConfigValue(s, extensionsGroupVersion, !disableExtensions)
	if disableExtensions {
		apiGroupVersionOverrides[extensionsGroupVersion] = genericapiserver.APIGroupVersionOverride{
			Disable: true,
		}
	}

	for key := range s.RuntimeConfig {
		if strings.HasPrefix(key, v1GroupVersion+"/") {
			return nil, fmt.Errorf("api/v1 resources cannot be enabled/disabled individually")
		} else if strings.HasPrefix(key, extensionsGroupVersion+"/") {
			resource := strings.TrimPrefix(key, extensionsGroupVersion+"/")

			apiGroupVersionOverride := apiGroupVersionOverrides[extensionsGroupVersion]
			if apiGroupVersionOverride.ResourceOverrides == nil {
				apiGroupVersionOverride.ResourceOverrides = map[string]bool{}
			}
			apiGroupVersionOverride.ResourceOverrides[resource] = getRuntimeConfigValue(s, key, false)
			apiGroupVersionOverrides[extensionsGroupVersion] = apiGroupVersionOverride
		}
	}
	return apiGroupVersionOverrides, nil
}
