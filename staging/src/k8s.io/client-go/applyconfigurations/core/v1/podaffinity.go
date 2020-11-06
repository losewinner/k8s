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

// PodAffinityApplyConfiguration represents an declarative configuration of the PodAffinity type for use
// with apply.
type PodAffinityApplyConfiguration struct {
	RequiredDuringSchedulingIgnoredDuringExecution  *PodAffinityTermList         `json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
	PreferredDuringSchedulingIgnoredDuringExecution *WeightedPodAffinityTermList `json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`
}

// PodAffinityApplyConfiguration constructs an declarative configuration of the PodAffinity type for use with
// apply.
func PodAffinity() *PodAffinityApplyConfiguration {
	return &PodAffinityApplyConfiguration{}
}

// SetRequiredDuringSchedulingIgnoredDuringExecution sets the RequiredDuringSchedulingIgnoredDuringExecution field in the declarative configuration to the given value.
func (b *PodAffinityApplyConfiguration) SetRequiredDuringSchedulingIgnoredDuringExecution(value PodAffinityTermList) *PodAffinityApplyConfiguration {
	b.RequiredDuringSchedulingIgnoredDuringExecution = &value
	return b
}

// RemoveRequiredDuringSchedulingIgnoredDuringExecution removes the RequiredDuringSchedulingIgnoredDuringExecution field from the declarative configuration.
func (b *PodAffinityApplyConfiguration) RemoveRequiredDuringSchedulingIgnoredDuringExecution() *PodAffinityApplyConfiguration {
	b.RequiredDuringSchedulingIgnoredDuringExecution = nil
	return b
}

// GetRequiredDuringSchedulingIgnoredDuringExecution gets the RequiredDuringSchedulingIgnoredDuringExecution field from the declarative configuration.
func (b *PodAffinityApplyConfiguration) GetRequiredDuringSchedulingIgnoredDuringExecution() (value PodAffinityTermList, ok bool) {
	if v := b.RequiredDuringSchedulingIgnoredDuringExecution; v != nil {
		return *v, true
	}
	return value, false
}

// SetPreferredDuringSchedulingIgnoredDuringExecution sets the PreferredDuringSchedulingIgnoredDuringExecution field in the declarative configuration to the given value.
func (b *PodAffinityApplyConfiguration) SetPreferredDuringSchedulingIgnoredDuringExecution(value WeightedPodAffinityTermList) *PodAffinityApplyConfiguration {
	b.PreferredDuringSchedulingIgnoredDuringExecution = &value
	return b
}

// RemovePreferredDuringSchedulingIgnoredDuringExecution removes the PreferredDuringSchedulingIgnoredDuringExecution field from the declarative configuration.
func (b *PodAffinityApplyConfiguration) RemovePreferredDuringSchedulingIgnoredDuringExecution() *PodAffinityApplyConfiguration {
	b.PreferredDuringSchedulingIgnoredDuringExecution = nil
	return b
}

// GetPreferredDuringSchedulingIgnoredDuringExecution gets the PreferredDuringSchedulingIgnoredDuringExecution field from the declarative configuration.
func (b *PodAffinityApplyConfiguration) GetPreferredDuringSchedulingIgnoredDuringExecution() (value WeightedPodAffinityTermList, ok bool) {
	if v := b.PreferredDuringSchedulingIgnoredDuringExecution; v != nil {
		return *v, true
	}
	return value, false
}

// PodAffinityList represents a listAlias of PodAffinityApplyConfiguration.
type PodAffinityList []*PodAffinityApplyConfiguration

// PodAffinityList represents a map of PodAffinityApplyConfiguration.
type PodAffinityMap map[string]PodAffinityApplyConfiguration
