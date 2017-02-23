/*
Copyright 2016 The Kubernetes Authors.

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

package validation

import (
	"net"
	"net/url"
	"os"
	"strings"

	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
	"k8s.io/kubernetes/cmd/kubeadm/app/constants"
	tokenutil "k8s.io/kubernetes/cmd/kubeadm/app/util/token"
	authzmodes "k8s.io/kubernetes/pkg/kubeapiserver/authorizer/modes"
	"k8s.io/kubernetes/pkg/registry/core/service/ipallocator"
)

// TODO: Break out the cloudprovider functionality out of core and only support the new flow
// described in https://github.com/kubernetes/community/pull/128
var cloudproviders = []string{
	"aws",
	"azure",
	"cloudstack",
	"gce",
	"mesos",
	"openstack",
	"ovirt",
	"photon",
	"rackspace",
	"vsphere",
}

func ValidateMasterConfiguration(c *kubeadm.MasterConfiguration) field.ErrorList {
	allErrs := field.ErrorList{}
	allErrs = append(allErrs, ValidateServiceSubnet(c.Networking.ServiceSubnet, field.NewPath("service subnet"))...)
	allErrs = append(allErrs, ValidateCloudProvider(c.CloudProvider, field.NewPath("cloudprovider"))...)
	allErrs = append(allErrs, ValidateAuthorizationMode(c.AuthorizationMode, field.NewPath("authorization-mode"))...)
	return allErrs
}

func ValidateNodeConfiguration(c *kubeadm.NodeConfiguration) field.ErrorList {
	allErrs := field.ErrorList{}
	allErrs = append(allErrs, ValidateDiscovery(c, field.NewPath("discovery"))...)
	return allErrs
}

func ValidateAuthorizationMode(authzMode string, fldPath *field.Path) field.ErrorList {
	if !authzmodes.IsValidAuthorizationMode(authzMode) {
		return field.ErrorList{field.Invalid(fldPath, nil, "invalid authorization mode")}
	}
	return field.ErrorList{}
}

func ValidateDiscovery(c *kubeadm.NodeConfiguration, fldPath *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}
	allErrs = append(allErrs, ValidateTLSBootstrapToken(c, fldPath)...)
	var count int
	if len(c.DiscoveryToken) != 0 {
		allErrs = append(allErrs, ValidateDiscoveryToken(c, fldPath)...)
		count++
	}
	if len(c.DiscoveryFile) != 0 {
		allErrs = append(allErrs, ValidateDiscoveryFile(c, fldPath)...)
		count++
	}
	if len(c.DiscoveryURL) != 0 {
		allErrs = append(allErrs, ValidateDiscoveryURL(c, fldPath)...)
		count++
	}
	if count > 1 {
		allErrs = append(allErrs, field.Invalid(fldPath, nil, "exactly one discovery strategy can be provided"))
	}
	return allErrs
}

func ValidateTLSBootstrapToken(cfg *kubeadm.NodeConfiguration, fldPath *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}
	if len(cfg.Token) != 0 && (len(cfg.TLSBootstrapToken) != 0 || len(cfg.DiscoveryToken) != 0) {
		allErrs = append(allErrs, field.Invalid(fldPath, nil, "--token is mutually exclusive with --bootstrap-token and --discovery-token"))
	}
	if len(cfg.Token) == 0 && (len(cfg.TLSBootstrapToken) == 0 || len(cfg.DiscoveryToken) == 0) {
		allErrs = append(
			allErrs,
			field.Invalid(fldPath, nil, "If --token is not specified, both --bootstrap-token and --discovery-token have to be provided"),
		)
	}
	if len(cfg.Token) != 0 {
		if len(cfg.DiscoveryURL) != 0 {
			cfg.TLSBootstrapToken = cfg.Token
		} else {
			cfg.DiscoveryURL = cfg.Token
		}
		if len(cfg.DiscoveryToken) != 0 {
			cfg.TLSBootstrapToken = cfg.Token
		} else {
			cfg.DiscoveryToken = cfg.Token
		}
	}
	if len(cfg.Masters) == 0 {
		allErrs = append(allErrs, field.Invalid(fldPath, nil, "master address(es) not specified"))
	}
	return allErrs
}

func ValidateDiscoveryFile(c *kubeadm.NodeConfiguration, fldPath *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}
	if _, err := os.Stat(c.DiscoveryFile); os.IsNotExist(err) {
		allErrs = append(allErrs, field.Invalid(fldPath, nil, "file does not exist"))
	}

	return allErrs
}

func ValidateDiscoveryURL(c *kubeadm.NodeConfiguration, fldPath *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}
	u, err := url.Parse(c.DiscoveryURL)
	if err != nil {
		allErrs = append(allErrs, field.Invalid(fldPath, nil, "invalide URL"))
	}
	if u.Scheme != "https" {
		allErrs = append(allErrs, field.Invalid(fldPath, nil, "must be https"))
	}
	return allErrs
}

func ValidateDiscoveryToken(c *kubeadm.NodeConfiguration, fldPath *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}
	if !strings.Contains(c.DiscoveryToken, "@") {
		c.DiscoveryToken = c.DiscoveryToken + "@"
	}

	id, secret, err := tokenutil.ParseToken(c.DiscoveryToken)
	if err != nil {
		allErrs = append(allErrs, field.Invalid(fldPath, nil, err.Error()))
	}

	if len(id) == 0 || len(secret) == 0 {
		allErrs = append(allErrs, field.Invalid(fldPath, nil, "token must be specific as <ID>:<Secret>"))
	}
	return allErrs
}

func ValidateServiceSubnet(subnet string, fldPath *field.Path) field.ErrorList {
	_, svcSubnet, err := net.ParseCIDR(subnet)
	if err != nil {
		return field.ErrorList{field.Invalid(fldPath, nil, "couldn't parse the service subnet")}
	}
	numAddresses := ipallocator.RangeSize(svcSubnet)
	if numAddresses < constants.MinimumAddressesInServiceSubnet {
		return field.ErrorList{field.Invalid(fldPath, nil, "service subnet is too small")}
	}
	return field.ErrorList{}
}

func ValidateCloudProvider(provider string, fldPath *field.Path) field.ErrorList {
	if len(provider) == 0 {
		return field.ErrorList{}
	}
	for _, supported := range cloudproviders {
		if provider == supported {
			return field.ErrorList{}
		}
	}
	return field.ErrorList{field.Invalid(fldPath, nil, "cloudprovider not supported")}
}
