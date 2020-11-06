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

package v1alpha1

import (
	v1alpha1 "k8s.io/api/flowcontrol/v1alpha1"
)

// FlowDistinguisherMethodApplyConfiguration represents an declarative configuration of the FlowDistinguisherMethod type for use
// with apply.
type FlowDistinguisherMethodApplyConfiguration struct {
	Type *v1alpha1.FlowDistinguisherMethodType `json:"type,omitempty"`
}

// FlowDistinguisherMethodApplyConfiguration constructs an declarative configuration of the FlowDistinguisherMethod type for use with
// apply.
func FlowDistinguisherMethod() *FlowDistinguisherMethodApplyConfiguration {
	return &FlowDistinguisherMethodApplyConfiguration{}
}

// SetType sets the Type field in the declarative configuration to the given value.
func (b *FlowDistinguisherMethodApplyConfiguration) SetType(value v1alpha1.FlowDistinguisherMethodType) *FlowDistinguisherMethodApplyConfiguration {
	b.Type = &value
	return b
}

// RemoveType removes the Type field from the declarative configuration.
func (b *FlowDistinguisherMethodApplyConfiguration) RemoveType() *FlowDistinguisherMethodApplyConfiguration {
	b.Type = nil
	return b
}

// GetType gets the Type field from the declarative configuration.
func (b *FlowDistinguisherMethodApplyConfiguration) GetType() (value v1alpha1.FlowDistinguisherMethodType, ok bool) {
	if v := b.Type; v != nil {
		return *v, true
	}
	return value, false
}

// FlowDistinguisherMethodList represents a listAlias of FlowDistinguisherMethodApplyConfiguration.
type FlowDistinguisherMethodList []*FlowDistinguisherMethodApplyConfiguration

// FlowDistinguisherMethodList represents a map of FlowDistinguisherMethodApplyConfiguration.
type FlowDistinguisherMethodMap map[string]FlowDistinguisherMethodApplyConfiguration
