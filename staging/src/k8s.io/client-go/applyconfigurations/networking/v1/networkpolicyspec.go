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

import (
	apinetworkingv1 "k8s.io/api/networking/v1"
	v1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

// NetworkPolicySpecApplyConfiguration represents an declarative configuration of the NetworkPolicySpec type for use
// with apply.
type NetworkPolicySpecApplyConfiguration struct {
	PodSelector *v1.LabelSelectorApplyConfiguration `json:"podSelector,omitempty"`
	Ingress     *NetworkPolicyIngressRuleList       `json:"ingress,omitempty"`
	Egress      *NetworkPolicyEgressRuleList        `json:"egress,omitempty"`
	PolicyTypes *[]apinetworkingv1.PolicyType       `json:"policyTypes,omitempty"`
}

// NetworkPolicySpecApplyConfiguration constructs an declarative configuration of the NetworkPolicySpec type for use with
// apply.
func NetworkPolicySpec() *NetworkPolicySpecApplyConfiguration {
	return &NetworkPolicySpecApplyConfiguration{}
}

// SetPodSelector sets the PodSelector field in the declarative configuration to the given value.
func (b *NetworkPolicySpecApplyConfiguration) SetPodSelector(value *v1.LabelSelectorApplyConfiguration) *NetworkPolicySpecApplyConfiguration {
	b.PodSelector = value
	return b
}

// RemovePodSelector removes the PodSelector field from the declarative configuration.
func (b *NetworkPolicySpecApplyConfiguration) RemovePodSelector() *NetworkPolicySpecApplyConfiguration {
	b.PodSelector = nil
	return b
}

// GetPodSelector gets the PodSelector field from the declarative configuration.
func (b *NetworkPolicySpecApplyConfiguration) GetPodSelector() (value *v1.LabelSelectorApplyConfiguration, ok bool) {
	return b.PodSelector, b.PodSelector != nil
}

// SetIngress sets the Ingress field in the declarative configuration to the given value.
func (b *NetworkPolicySpecApplyConfiguration) SetIngress(value NetworkPolicyIngressRuleList) *NetworkPolicySpecApplyConfiguration {
	b.Ingress = &value
	return b
}

// RemoveIngress removes the Ingress field from the declarative configuration.
func (b *NetworkPolicySpecApplyConfiguration) RemoveIngress() *NetworkPolicySpecApplyConfiguration {
	b.Ingress = nil
	return b
}

// GetIngress gets the Ingress field from the declarative configuration.
func (b *NetworkPolicySpecApplyConfiguration) GetIngress() (value NetworkPolicyIngressRuleList, ok bool) {
	if v := b.Ingress; v != nil {
		return *v, true
	}
	return value, false
}

// SetEgress sets the Egress field in the declarative configuration to the given value.
func (b *NetworkPolicySpecApplyConfiguration) SetEgress(value NetworkPolicyEgressRuleList) *NetworkPolicySpecApplyConfiguration {
	b.Egress = &value
	return b
}

// RemoveEgress removes the Egress field from the declarative configuration.
func (b *NetworkPolicySpecApplyConfiguration) RemoveEgress() *NetworkPolicySpecApplyConfiguration {
	b.Egress = nil
	return b
}

// GetEgress gets the Egress field from the declarative configuration.
func (b *NetworkPolicySpecApplyConfiguration) GetEgress() (value NetworkPolicyEgressRuleList, ok bool) {
	if v := b.Egress; v != nil {
		return *v, true
	}
	return value, false
}

// SetPolicyTypes sets the PolicyTypes field in the declarative configuration to the given value.
func (b *NetworkPolicySpecApplyConfiguration) SetPolicyTypes(value []apinetworkingv1.PolicyType) *NetworkPolicySpecApplyConfiguration {
	b.PolicyTypes = &value
	return b
}

// RemovePolicyTypes removes the PolicyTypes field from the declarative configuration.
func (b *NetworkPolicySpecApplyConfiguration) RemovePolicyTypes() *NetworkPolicySpecApplyConfiguration {
	b.PolicyTypes = nil
	return b
}

// GetPolicyTypes gets the PolicyTypes field from the declarative configuration.
func (b *NetworkPolicySpecApplyConfiguration) GetPolicyTypes() (value []apinetworkingv1.PolicyType, ok bool) {
	if v := b.PolicyTypes; v != nil {
		return *v, true
	}
	return value, false
}

// NetworkPolicySpecList represents a listAlias of NetworkPolicySpecApplyConfiguration.
type NetworkPolicySpecList []*NetworkPolicySpecApplyConfiguration

// NetworkPolicySpecList represents a map of NetworkPolicySpecApplyConfiguration.
type NetworkPolicySpecMap map[string]NetworkPolicySpecApplyConfiguration
