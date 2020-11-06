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
	v1 "k8s.io/api/core/v1"
)

// NamespaceStatusApplyConfiguration represents an declarative configuration of the NamespaceStatus type for use
// with apply.
type NamespaceStatusApplyConfiguration struct {
	Phase      *v1.NamespacePhase      `json:"phase,omitempty"`
	Conditions *NamespaceConditionList `json:"conditions,omitempty"`
}

// NamespaceStatusApplyConfiguration constructs an declarative configuration of the NamespaceStatus type for use with
// apply.
func NamespaceStatus() *NamespaceStatusApplyConfiguration {
	return &NamespaceStatusApplyConfiguration{}
}

// SetPhase sets the Phase field in the declarative configuration to the given value.
func (b *NamespaceStatusApplyConfiguration) SetPhase(value v1.NamespacePhase) *NamespaceStatusApplyConfiguration {
	b.Phase = &value
	return b
}

// RemovePhase removes the Phase field from the declarative configuration.
func (b *NamespaceStatusApplyConfiguration) RemovePhase() *NamespaceStatusApplyConfiguration {
	b.Phase = nil
	return b
}

// GetPhase gets the Phase field from the declarative configuration.
func (b *NamespaceStatusApplyConfiguration) GetPhase() (value v1.NamespacePhase, ok bool) {
	if v := b.Phase; v != nil {
		return *v, true
	}
	return value, false
}

// SetConditions sets the Conditions field in the declarative configuration to the given value.
func (b *NamespaceStatusApplyConfiguration) SetConditions(value NamespaceConditionList) *NamespaceStatusApplyConfiguration {
	b.Conditions = &value
	return b
}

// RemoveConditions removes the Conditions field from the declarative configuration.
func (b *NamespaceStatusApplyConfiguration) RemoveConditions() *NamespaceStatusApplyConfiguration {
	b.Conditions = nil
	return b
}

// GetConditions gets the Conditions field from the declarative configuration.
func (b *NamespaceStatusApplyConfiguration) GetConditions() (value NamespaceConditionList, ok bool) {
	if v := b.Conditions; v != nil {
		return *v, true
	}
	return value, false
}

// NamespaceStatusList represents a listAlias of NamespaceStatusApplyConfiguration.
type NamespaceStatusList []*NamespaceStatusApplyConfiguration

// NamespaceStatusList represents a map of NamespaceStatusApplyConfiguration.
type NamespaceStatusMap map[string]NamespaceStatusApplyConfiguration
