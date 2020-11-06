/*
Copyright The Kubernetes Authors.

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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// LoadBalancerStatusApplyConfiguration represents an declarative configuration of the LoadBalancerStatus type for use
// with apply.
type LoadBalancerStatusApplyConfiguration struct {
	Ingress *LoadBalancerIngressList `json:"ingress,omitempty"`
}

// LoadBalancerStatusApplyConfiguration constructs an declarative configuration of the LoadBalancerStatus type for use with
// apply.
func LoadBalancerStatus() *LoadBalancerStatusApplyConfiguration {
	return &LoadBalancerStatusApplyConfiguration{}
}

// SetIngress sets the Ingress field in the declarative configuration to the given value.
func (b *LoadBalancerStatusApplyConfiguration) SetIngress(value LoadBalancerIngressList) *LoadBalancerStatusApplyConfiguration {
	b.Ingress = &value
	return b
}

// RemoveIngress removes the Ingress field from the declarative configuration.
func (b *LoadBalancerStatusApplyConfiguration) RemoveIngress() *LoadBalancerStatusApplyConfiguration {
	b.Ingress = nil
	return b
}

// GetIngress gets the Ingress field from the declarative configuration.
func (b *LoadBalancerStatusApplyConfiguration) GetIngress() (value LoadBalancerIngressList, ok bool) {
	if v := b.Ingress; v != nil {
		return *v, true
	}
	return value, false
}

// LoadBalancerStatusList represents a listAlias of LoadBalancerStatusApplyConfiguration.
type LoadBalancerStatusList []*LoadBalancerStatusApplyConfiguration

// LoadBalancerStatusList represents a map of LoadBalancerStatusApplyConfiguration.
type LoadBalancerStatusMap map[string]LoadBalancerStatusApplyConfiguration
