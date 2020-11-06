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

package v1beta2

import (
	v1beta2 "k8s.io/api/apps/v1beta2"
	corev1 "k8s.io/client-go/applyconfigurations/core/v1"
	v1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

// StatefulSetSpecApplyConfiguration represents an declarative configuration of the StatefulSetSpec type for use
// with apply.
type StatefulSetSpecApplyConfiguration struct {
	Replicas             *int32                                       `json:"replicas,omitempty"`
	Selector             *v1.LabelSelectorApplyConfiguration          `json:"selector,omitempty"`
	Template             *corev1.PodTemplateSpecApplyConfiguration    `json:"template,omitempty"`
	VolumeClaimTemplates *corev1.PersistentVolumeClaimList            `json:"volumeClaimTemplates,omitempty"`
	ServiceName          *string                                      `json:"serviceName,omitempty"`
	PodManagementPolicy  *v1beta2.PodManagementPolicyType             `json:"podManagementPolicy,omitempty"`
	UpdateStrategy       *StatefulSetUpdateStrategyApplyConfiguration `json:"updateStrategy,omitempty"`
	RevisionHistoryLimit *int32                                       `json:"revisionHistoryLimit,omitempty"`
}

// StatefulSetSpecApplyConfiguration constructs an declarative configuration of the StatefulSetSpec type for use with
// apply.
func StatefulSetSpec() *StatefulSetSpecApplyConfiguration {
	return &StatefulSetSpecApplyConfiguration{}
}

// SetReplicas sets the Replicas field in the declarative configuration to the given value.
func (b *StatefulSetSpecApplyConfiguration) SetReplicas(value int32) *StatefulSetSpecApplyConfiguration {
	b.Replicas = &value
	return b
}

// RemoveReplicas removes the Replicas field from the declarative configuration.
func (b *StatefulSetSpecApplyConfiguration) RemoveReplicas() *StatefulSetSpecApplyConfiguration {
	b.Replicas = nil
	return b
}

// GetReplicas gets the Replicas field from the declarative configuration.
func (b *StatefulSetSpecApplyConfiguration) GetReplicas() (value int32, ok bool) {
	if v := b.Replicas; v != nil {
		return *v, true
	}
	return value, false
}

// SetSelector sets the Selector field in the declarative configuration to the given value.
func (b *StatefulSetSpecApplyConfiguration) SetSelector(value *v1.LabelSelectorApplyConfiguration) *StatefulSetSpecApplyConfiguration {
	b.Selector = value
	return b
}

// RemoveSelector removes the Selector field from the declarative configuration.
func (b *StatefulSetSpecApplyConfiguration) RemoveSelector() *StatefulSetSpecApplyConfiguration {
	b.Selector = nil
	return b
}

// GetSelector gets the Selector field from the declarative configuration.
func (b *StatefulSetSpecApplyConfiguration) GetSelector() (value *v1.LabelSelectorApplyConfiguration, ok bool) {
	return b.Selector, b.Selector != nil
}

// SetTemplate sets the Template field in the declarative configuration to the given value.
func (b *StatefulSetSpecApplyConfiguration) SetTemplate(value *corev1.PodTemplateSpecApplyConfiguration) *StatefulSetSpecApplyConfiguration {
	b.Template = value
	return b
}

// RemoveTemplate removes the Template field from the declarative configuration.
func (b *StatefulSetSpecApplyConfiguration) RemoveTemplate() *StatefulSetSpecApplyConfiguration {
	b.Template = nil
	return b
}

// GetTemplate gets the Template field from the declarative configuration.
func (b *StatefulSetSpecApplyConfiguration) GetTemplate() (value *corev1.PodTemplateSpecApplyConfiguration, ok bool) {
	return b.Template, b.Template != nil
}

// SetVolumeClaimTemplates sets the VolumeClaimTemplates field in the declarative configuration to the given value.
func (b *StatefulSetSpecApplyConfiguration) SetVolumeClaimTemplates(value corev1.PersistentVolumeClaimList) *StatefulSetSpecApplyConfiguration {
	b.VolumeClaimTemplates = &value
	return b
}

// RemoveVolumeClaimTemplates removes the VolumeClaimTemplates field from the declarative configuration.
func (b *StatefulSetSpecApplyConfiguration) RemoveVolumeClaimTemplates() *StatefulSetSpecApplyConfiguration {
	b.VolumeClaimTemplates = nil
	return b
}

// GetVolumeClaimTemplates gets the VolumeClaimTemplates field from the declarative configuration.
func (b *StatefulSetSpecApplyConfiguration) GetVolumeClaimTemplates() (value corev1.PersistentVolumeClaimList, ok bool) {
	if v := b.VolumeClaimTemplates; v != nil {
		return *v, true
	}
	return value, false
}

// SetServiceName sets the ServiceName field in the declarative configuration to the given value.
func (b *StatefulSetSpecApplyConfiguration) SetServiceName(value string) *StatefulSetSpecApplyConfiguration {
	b.ServiceName = &value
	return b
}

// RemoveServiceName removes the ServiceName field from the declarative configuration.
func (b *StatefulSetSpecApplyConfiguration) RemoveServiceName() *StatefulSetSpecApplyConfiguration {
	b.ServiceName = nil
	return b
}

// GetServiceName gets the ServiceName field from the declarative configuration.
func (b *StatefulSetSpecApplyConfiguration) GetServiceName() (value string, ok bool) {
	if v := b.ServiceName; v != nil {
		return *v, true
	}
	return value, false
}

// SetPodManagementPolicy sets the PodManagementPolicy field in the declarative configuration to the given value.
func (b *StatefulSetSpecApplyConfiguration) SetPodManagementPolicy(value v1beta2.PodManagementPolicyType) *StatefulSetSpecApplyConfiguration {
	b.PodManagementPolicy = &value
	return b
}

// RemovePodManagementPolicy removes the PodManagementPolicy field from the declarative configuration.
func (b *StatefulSetSpecApplyConfiguration) RemovePodManagementPolicy() *StatefulSetSpecApplyConfiguration {
	b.PodManagementPolicy = nil
	return b
}

// GetPodManagementPolicy gets the PodManagementPolicy field from the declarative configuration.
func (b *StatefulSetSpecApplyConfiguration) GetPodManagementPolicy() (value v1beta2.PodManagementPolicyType, ok bool) {
	if v := b.PodManagementPolicy; v != nil {
		return *v, true
	}
	return value, false
}

// SetUpdateStrategy sets the UpdateStrategy field in the declarative configuration to the given value.
func (b *StatefulSetSpecApplyConfiguration) SetUpdateStrategy(value *StatefulSetUpdateStrategyApplyConfiguration) *StatefulSetSpecApplyConfiguration {
	b.UpdateStrategy = value
	return b
}

// RemoveUpdateStrategy removes the UpdateStrategy field from the declarative configuration.
func (b *StatefulSetSpecApplyConfiguration) RemoveUpdateStrategy() *StatefulSetSpecApplyConfiguration {
	b.UpdateStrategy = nil
	return b
}

// GetUpdateStrategy gets the UpdateStrategy field from the declarative configuration.
func (b *StatefulSetSpecApplyConfiguration) GetUpdateStrategy() (value *StatefulSetUpdateStrategyApplyConfiguration, ok bool) {
	return b.UpdateStrategy, b.UpdateStrategy != nil
}

// SetRevisionHistoryLimit sets the RevisionHistoryLimit field in the declarative configuration to the given value.
func (b *StatefulSetSpecApplyConfiguration) SetRevisionHistoryLimit(value int32) *StatefulSetSpecApplyConfiguration {
	b.RevisionHistoryLimit = &value
	return b
}

// RemoveRevisionHistoryLimit removes the RevisionHistoryLimit field from the declarative configuration.
func (b *StatefulSetSpecApplyConfiguration) RemoveRevisionHistoryLimit() *StatefulSetSpecApplyConfiguration {
	b.RevisionHistoryLimit = nil
	return b
}

// GetRevisionHistoryLimit gets the RevisionHistoryLimit field from the declarative configuration.
func (b *StatefulSetSpecApplyConfiguration) GetRevisionHistoryLimit() (value int32, ok bool) {
	if v := b.RevisionHistoryLimit; v != nil {
		return *v, true
	}
	return value, false
}

// StatefulSetSpecList represents a listAlias of StatefulSetSpecApplyConfiguration.
type StatefulSetSpecList []*StatefulSetSpecApplyConfiguration

// StatefulSetSpecList represents a map of StatefulSetSpecApplyConfiguration.
type StatefulSetSpecMap map[string]StatefulSetSpecApplyConfiguration
