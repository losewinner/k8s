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

package v1beta1

import (
	v1beta1 "k8s.io/api/policy/v1beta1"
)

// SupplementalGroupsStrategyOptionsApplyConfiguration represents an declarative configuration of the SupplementalGroupsStrategyOptions type for use
// with apply.
type SupplementalGroupsStrategyOptionsApplyConfiguration struct {
	Rule   *v1beta1.SupplementalGroupsStrategyType `json:"rule,omitempty"`
	Ranges *IDRangeList                            `json:"ranges,omitempty"`
}

// SupplementalGroupsStrategyOptionsApplyConfiguration constructs an declarative configuration of the SupplementalGroupsStrategyOptions type for use with
// apply.
func SupplementalGroupsStrategyOptions() *SupplementalGroupsStrategyOptionsApplyConfiguration {
	return &SupplementalGroupsStrategyOptionsApplyConfiguration{}
}

// SetRule sets the Rule field in the declarative configuration to the given value.
func (b *SupplementalGroupsStrategyOptionsApplyConfiguration) SetRule(value v1beta1.SupplementalGroupsStrategyType) *SupplementalGroupsStrategyOptionsApplyConfiguration {
	b.Rule = &value
	return b
}

// RemoveRule removes the Rule field from the declarative configuration.
func (b *SupplementalGroupsStrategyOptionsApplyConfiguration) RemoveRule() *SupplementalGroupsStrategyOptionsApplyConfiguration {
	b.Rule = nil
	return b
}

// GetRule gets the Rule field from the declarative configuration.
func (b *SupplementalGroupsStrategyOptionsApplyConfiguration) GetRule() (value v1beta1.SupplementalGroupsStrategyType, ok bool) {
	if v := b.Rule; v != nil {
		return *v, true
	}
	return value, false
}

// SetRanges sets the Ranges field in the declarative configuration to the given value.
func (b *SupplementalGroupsStrategyOptionsApplyConfiguration) SetRanges(value IDRangeList) *SupplementalGroupsStrategyOptionsApplyConfiguration {
	b.Ranges = &value
	return b
}

// RemoveRanges removes the Ranges field from the declarative configuration.
func (b *SupplementalGroupsStrategyOptionsApplyConfiguration) RemoveRanges() *SupplementalGroupsStrategyOptionsApplyConfiguration {
	b.Ranges = nil
	return b
}

// GetRanges gets the Ranges field from the declarative configuration.
func (b *SupplementalGroupsStrategyOptionsApplyConfiguration) GetRanges() (value IDRangeList, ok bool) {
	if v := b.Ranges; v != nil {
		return *v, true
	}
	return value, false
}

// SupplementalGroupsStrategyOptionsList represents a listAlias of SupplementalGroupsStrategyOptionsApplyConfiguration.
type SupplementalGroupsStrategyOptionsList []*SupplementalGroupsStrategyOptionsApplyConfiguration

// SupplementalGroupsStrategyOptionsList represents a map of SupplementalGroupsStrategyOptionsApplyConfiguration.
type SupplementalGroupsStrategyOptionsMap map[string]SupplementalGroupsStrategyOptionsApplyConfiguration
