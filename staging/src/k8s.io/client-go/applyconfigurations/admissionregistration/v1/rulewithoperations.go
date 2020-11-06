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
	v1 "k8s.io/api/admissionregistration/v1"
)

// RuleWithOperationsApplyConfiguration represents an declarative configuration of the RuleWithOperations type for use
// with apply.
type RuleWithOperationsApplyConfiguration struct {
	Operations             *[]v1.OperationType `json:"operations,omitempty"`
	RuleApplyConfiguration `json:",inline"`
}

// RuleWithOperationsApplyConfiguration constructs an declarative configuration of the RuleWithOperations type for use with
// apply.
func RuleWithOperations() *RuleWithOperationsApplyConfiguration {
	return &RuleWithOperationsApplyConfiguration{}
}

// SetOperations sets the Operations field in the declarative configuration to the given value.
func (b *RuleWithOperationsApplyConfiguration) SetOperations(value []v1.OperationType) *RuleWithOperationsApplyConfiguration {
	b.Operations = &value
	return b
}

// RemoveOperations removes the Operations field from the declarative configuration.
func (b *RuleWithOperationsApplyConfiguration) RemoveOperations() *RuleWithOperationsApplyConfiguration {
	b.Operations = nil
	return b
}

// GetOperations gets the Operations field from the declarative configuration.
func (b *RuleWithOperationsApplyConfiguration) GetOperations() (value []v1.OperationType, ok bool) {
	if v := b.Operations; v != nil {
		return *v, true
	}
	return value, false
}

// SetRule sets the Rule field in the declarative configuration to the given value.
func (b *RuleWithOperationsApplyConfiguration) SetRule(value *RuleApplyConfiguration) *RuleWithOperationsApplyConfiguration {
	if value != nil {
		b.RuleApplyConfiguration = *value
	}
	return b
}

// GetRule gets the Rule field from the declarative configuration.
func (b *RuleWithOperationsApplyConfiguration) GetRule() (value *RuleApplyConfiguration, ok bool) {
	return &b.RuleApplyConfiguration, true
}

// RuleWithOperationsList represents a listAlias of RuleWithOperationsApplyConfiguration.
type RuleWithOperationsList []*RuleWithOperationsApplyConfiguration

// RuleWithOperationsList represents a map of RuleWithOperationsApplyConfiguration.
type RuleWithOperationsMap map[string]RuleWithOperationsApplyConfiguration
