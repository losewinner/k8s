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

package framework

import (
	"fmt"
	"strconv"
	"strings"

	"k8s.io/kubernetes/pkg/api/v1"
	clientset "k8s.io/kubernetes/pkg/client/clientset_generated/release_1_5"
	"k8s.io/kubernetes/pkg/cloudprovider"
	gcecloud "k8s.io/kubernetes/pkg/cloudprovider/providers/gce"
	"k8s.io/kubernetes/pkg/util/sets"

	. "github.com/onsi/gomega"
	compute "google.golang.org/api/compute/v1"
)

// This should match the formatting of makeFirewallName() in pkg/cloudprovider/providers/gce/gce.go
func MakeFirewallNameForLBSvc(name string) string {
	return fmt.Sprintf("k8s-fw-%s", name)
}

// ConstructFwForLBSvc returns the expected GCE firewall rule for a loadbalancer type service
func ConstructFwForLBSvc(svc *v1.Service, nodesTags []string) (*compute.Firewall, error) {
	if svc.Spec.Type != v1.ServiceTypeLoadBalancer {
		return nil, fmt.Errorf("can not construct firewall rule for non-loadbalancer type service")
	}
	fw := compute.Firewall{}
	fw.Name = MakeFirewallNameForLBSvc(cloudprovider.GetLoadBalancerName(svc))
	fw.TargetTags = nodesTags
	if svc.Spec.LoadBalancerSourceRanges == nil {
		fw.SourceRanges = []string{"0.0.0.0/0"}
	} else {
		fw.SourceRanges = svc.Spec.LoadBalancerSourceRanges
	}
	for _, sp := range svc.Spec.Ports {
		fw.Allowed = append(fw.Allowed, &compute.FirewallAllowed{
			IPProtocol: strings.ToLower(string(sp.Protocol)),
			Ports:      []string{strconv.Itoa(int(sp.Port))},
		})
	}
	return &fw, nil
}

func GetNodeTags(c clientset.Interface, cloudConfig CloudConfig) *compute.Tags {
	nodes := GetReadySchedulableNodesOrDie(c)
	Expect(len(nodes.Items) > 0).Should(BeTrue())
	nodeTags, err := GetInstanceTags(cloudConfig, nodes.Items[0].Name)
	Expect(err).ToNot(HaveOccurred())
	return nodeTags
}

func GetInstanceTags(cloudConfig CloudConfig, instanceName string) (*compute.Tags, error) {
	gceCloud := cloudConfig.Provider.(*gcecloud.GCECloud)
	res, err := gceCloud.GetComputeService().Instances.Get(cloudConfig.ProjectID, cloudConfig.Zone, instanceName).Do()
	if err != nil {
		return nil, err
	}
	return res.Tags, nil
}

func SetInstanceTags(cloudConfig CloudConfig, instanceName string, tags []string) ([]string, error) {
	gceCloud := cloudConfig.Provider.(*gcecloud.GCECloud)
	// Re-get instance everytime because we need the latest fingerprint for updating metadata
	resTags, err := GetInstanceTags(cloudConfig, instanceName)
	if err != nil {
		return nil, fmt.Errorf("failed to get instance tags: %v", err)
	}

	_, err = gceCloud.GetComputeService().Instances.SetTags(cloudConfig.ProjectID, cloudConfig.Zone, instanceName, &compute.Tags{Fingerprint: resTags.Fingerprint, Items: tags}).Do()
	if err != nil {
		return nil, fmt.Errorf("failed to set instance tags: %v", err)
	}
	Logf("Sent request to set tags %v on instance: %v", tags, instanceName)
	return resTags.Items, nil
}

// From cluster/gce/config-test.sh, master name is set up using below format:
// MASTER_NAME="${INSTANCE_PREFIX}-master"
func GetInstancePrefix(masterName string) (string, error) {
	if !strings.HasSuffix(masterName, "-master") {
		return "", fmt.Errorf("unexpected master name format: %v", masterName)
	}
	return masterName[:len(masterName)-7], nil
}

// From cluster/gce/config-test.sh, cluster ip range is set up using below command:
// CLUSTER_IP_RANGE="${CLUSTER_IP_RANGE:-10.180.0.0/14}"
// Warning: this need to be consistent with the CLUSTER_IP_RANGE in startup scripts,
// which is hardcoded currently
func GetClusterIpRange() string {
	return "10.180.0.0/14"
}

// From cluster/gce/util.sh, all firewall rules should be consistent with the ones created by startup scripts
func GetE2eFirewalls(masterName, masterTag, instancePrefix, nodeTag, network string) []*compute.Firewall {
	fws := []*compute.Firewall{}
	fws = append(fws, &compute.Firewall{
		Name:         network + "-default-internal-master",
		SourceRanges: []string{"10.0.0.0/8"},
		TargetTags:   []string{masterTag},
		Allowed: []*compute.FirewallAllowed{
			{
				IPProtocol: "tcp",
				Ports:      []string{"1-2379"},
			},
			{
				IPProtocol: "tcp",
				Ports:      []string{"2382-65535"},
			},
			{
				IPProtocol: "udp",
				Ports:      []string{"1-65535"},
			},
			{
				IPProtocol: "icmp",
			},
		},
	})
	fws = append(fws, &compute.Firewall{
		Name:         network + "-default-internal-node",
		SourceRanges: []string{"10.0.0.0/8"},
		TargetTags:   []string{nodeTag},
		Allowed: []*compute.FirewallAllowed{
			{
				IPProtocol: "tcp",
				Ports:      []string{"1-65535"},
			},
			{
				IPProtocol: "udp",
				Ports:      []string{"1-65535"},
			},
			{
				IPProtocol: "icmp",
			},
		},
	})
	fws = append(fws, &compute.Firewall{
		Name:         network + "-default-ssh",
		SourceRanges: []string{"0.0.0.0/0"},
		Allowed: []*compute.FirewallAllowed{
			{
				IPProtocol: "tcp",
				Ports:      []string{"22"},
			},
		},
	})
	fws = append(fws, &compute.Firewall{
		Name:       masterName + "-etcd",
		SourceTags: []string{masterTag},
		TargetTags: []string{masterTag},
		Allowed: []*compute.FirewallAllowed{
			{
				IPProtocol: "tcp",
				Ports:      []string{"2380"},
			},
			{
				IPProtocol: "tcp",
				Ports:      []string{"2381"},
			},
		},
	})
	fws = append(fws, &compute.Firewall{
		Name:         masterName + "-https",
		SourceRanges: []string{"0.0.0.0/0"},
		TargetTags:   []string{masterTag},
		Allowed: []*compute.FirewallAllowed{
			{
				IPProtocol: "tcp",
				Ports:      []string{"443"},
			},
		},
	})
	fws = append(fws, &compute.Firewall{
		Name:         nodeTag + "-all",
		SourceRanges: []string{GetClusterIpRange()},
		TargetTags:   []string{nodeTag},
		Allowed: []*compute.FirewallAllowed{
			{
				IPProtocol: "tcp",
			},
			{
				IPProtocol: "udp",
			},
			{
				IPProtocol: "icmp",
			},
			{
				IPProtocol: "esp",
			},
			{
				IPProtocol: "ah",
			},
			{
				IPProtocol: "sctp",
			},
		},
	})
	fws = append(fws, &compute.Firewall{
		Name:         nodeTag + "-" + instancePrefix + "-http-alt",
		SourceRanges: []string{"0.0.0.0/0"},
		TargetTags:   []string{nodeTag},
		Allowed: []*compute.FirewallAllowed{
			{
				IPProtocol: "tcp",
				Ports:      []string{"80"},
			},
			{
				IPProtocol: "tcp",
				Ports:      []string{"8080"},
			},
		},
	})
	fws = append(fws, &compute.Firewall{
		Name:         nodeTag + "-" + instancePrefix + "-nodeports",
		SourceRanges: []string{"0.0.0.0/0"},
		TargetTags:   []string{nodeTag},
		Allowed: []*compute.FirewallAllowed{
			{
				IPProtocol: "tcp",
				Ports:      []string{"30000-32767"},
			},
			{
				IPProtocol: "udp",
				Ports:      []string{"30000-32767"},
			},
		},
	})
	return fws
}

func PackProtocolsPortsFromFw(alloweds []*compute.FirewallAllowed) []string {
	protocolPorts := []string{}
	for _, allowed := range alloweds {
		for _, port := range allowed.Ports {
			protocolPorts = append(protocolPorts, strings.ToLower(allowed.IPProtocol+"/"+port))
		}
	}
	return protocolPorts
}

// SameStringArray verifies whether two string arrays have the same strings, return error if not.
// Order does not matter.
// When `include` is set to true, verifies whether result includes all elements from expected.
func SameStringArray(result, expected []string, include bool) error {
	res := sets.NewString(result...)
	exp := sets.NewString(expected...)
	if !include {
		diff := res.Difference(exp)
		if len(diff) != 0 {
			return fmt.Errorf("found differences: %v", diff)
		}
	} else {
		if !res.IsSuperset(exp) {
			return fmt.Errorf("some elements are missing: expected %v, got %v", expected, result)
		}
	}
	return nil
}

func VerifyFirewallRule(res, exp *compute.Firewall, network string, portsSubset bool) error {
	if res.Name != exp.Name {
		return fmt.Errorf("incorrect name: %v, expected %v", res.Name, exp.Name)
	}
	// Sample Network value: https://www.googleapis.com/compute/v1/projects/{project-id}/global/networks/e2e
	if !strings.HasSuffix(res.Network, "/"+network) {
		return fmt.Errorf("incorrect network: %v, expected ends with: %v", res.Network, "/"+network)
	}
	if err := SameStringArray(PackProtocolsPortsFromFw(res.Allowed), PackProtocolsPortsFromFw(exp.Allowed), portsSubset); err != nil {
		return fmt.Errorf("incorrect allowed protocols ports: %v", err)
	}
	if err := SameStringArray(res.SourceRanges, exp.SourceRanges, false); err != nil {
		return fmt.Errorf("incorrect source ranges %v, expected %v: %v", res.SourceRanges, exp.SourceRanges, err)
	}
	if err := SameStringArray(res.SourceTags, exp.SourceTags, false); err != nil {
		return fmt.Errorf("incorrect source tags %v, expected %v: %v", res.SourceTags, exp.SourceTags, err)
	}
	if err := SameStringArray(res.TargetTags, exp.TargetTags, false); err != nil {
		return fmt.Errorf("incorrect target tags %v, expected %v: %v", res.TargetTags, exp.TargetTags, err)
	}
	return nil
}
