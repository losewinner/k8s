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
	v1beta1 "k8s.io/api/admissionregistration/v1beta1"
)

// RuleApplyConfiguration represents an declarative configuration of the Rule type for use
// with apply.
type RuleApplyConfiguration struct {
	APIGroups   *[]string          `json:"apiGroups,omitempty"`
	APIVersions *[]string          `json:"apiVersions,omitempty"`
	Resources   *[]string          `json:"resources,omitempty"`
	Scope       *v1beta1.ScopeType `json:"scope,omitempty"`
}

// RuleApplyConfiguration constructs an declarative configuration of the Rule type for use with
// apply.
func Rule() *RuleApplyConfiguration {
	return &RuleApplyConfiguration{}
}

// SetAPIGroups sets the APIGroups field in the declarative configuration to the given value.
func (b *RuleApplyConfiguration) SetAPIGroups(value []string) *RuleApplyConfiguration {
	b.APIGroups = &value
	return b
}

// RemoveAPIGroups removes the APIGroups field from the declarative configuration.
func (b *RuleApplyConfiguration) RemoveAPIGroups() *RuleApplyConfiguration {
	b.APIGroups = nil
	return b
}

// GetAPIGroups gets the APIGroups field from the declarative configuration.
func (b *RuleApplyConfiguration) GetAPIGroups() (value []string, ok bool) {
	if v := b.APIGroups; v != nil {
		return *v, true
	}
	return value, false
}

// SetAPIVersions sets the APIVersions field in the declarative configuration to the given value.
func (b *RuleApplyConfiguration) SetAPIVersions(value []string) *RuleApplyConfiguration {
	b.APIVersions = &value
	return b
}

// RemoveAPIVersions removes the APIVersions field from the declarative configuration.
func (b *RuleApplyConfiguration) RemoveAPIVersions() *RuleApplyConfiguration {
	b.APIVersions = nil
	return b
}

// GetAPIVersions gets the APIVersions field from the declarative configuration.
func (b *RuleApplyConfiguration) GetAPIVersions() (value []string, ok bool) {
	if v := b.APIVersions; v != nil {
		return *v, true
	}
	return value, false
}

// SetResources sets the Resources field in the declarative configuration to the given value.
func (b *RuleApplyConfiguration) SetResources(value []string) *RuleApplyConfiguration {
	b.Resources = &value
	return b
}

// RemoveResources removes the Resources field from the declarative configuration.
func (b *RuleApplyConfiguration) RemoveResources() *RuleApplyConfiguration {
	b.Resources = nil
	return b
}

// GetResources gets the Resources field from the declarative configuration.
func (b *RuleApplyConfiguration) GetResources() (value []string, ok bool) {
	if v := b.Resources; v != nil {
		return *v, true
	}
	return value, false
}

// SetScope sets the Scope field in the declarative configuration to the given value.
func (b *RuleApplyConfiguration) SetScope(value v1beta1.ScopeType) *RuleApplyConfiguration {
	b.Scope = &value
	return b
}

// RemoveScope removes the Scope field from the declarative configuration.
func (b *RuleApplyConfiguration) RemoveScope() *RuleApplyConfiguration {
	b.Scope = nil
	return b
}

// GetScope gets the Scope field from the declarative configuration.
func (b *RuleApplyConfiguration) GetScope() (value v1beta1.ScopeType, ok bool) {
	if v := b.Scope; v != nil {
		return *v, true
	}
	return value, false
}

// RuleList represents a listAlias of RuleApplyConfiguration.
type RuleList []*RuleApplyConfiguration

// RuleList represents a map of RuleApplyConfiguration.
type RuleMap map[string]RuleApplyConfiguration
