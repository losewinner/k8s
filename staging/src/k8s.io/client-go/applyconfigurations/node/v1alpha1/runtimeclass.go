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
	v1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

// RuntimeClassApplyConfiguration represents an declarative configuration of the RuntimeClass type for use
// with apply.
type RuntimeClassApplyConfiguration struct {
	v1.TypeMetaApplyConfiguration `json:",inline"`
	ObjectMeta                    *v1.ObjectMetaApplyConfiguration    `json:"metadata,omitempty"`
	Spec                          *RuntimeClassSpecApplyConfiguration `json:"spec,omitempty"`
}

// RuntimeClassApplyConfiguration constructs an declarative configuration of the RuntimeClass type for use with
// apply.
func RuntimeClass() *RuntimeClassApplyConfiguration {
	return &RuntimeClassApplyConfiguration{}
}

// SetTypeMeta sets the TypeMeta field in the declarative configuration to the given value.
func (b *RuntimeClassApplyConfiguration) SetTypeMeta(value *v1.TypeMetaApplyConfiguration) *RuntimeClassApplyConfiguration {
	if value != nil {
		b.TypeMetaApplyConfiguration = *value
	}
	return b
}

// GetTypeMeta gets the TypeMeta field from the declarative configuration.
func (b *RuntimeClassApplyConfiguration) GetTypeMeta() (value *v1.TypeMetaApplyConfiguration, ok bool) {
	return &b.TypeMetaApplyConfiguration, true
}

// SetObjectMeta sets the ObjectMeta field in the declarative configuration to the given value.
func (b *RuntimeClassApplyConfiguration) SetObjectMeta(value *v1.ObjectMetaApplyConfiguration) *RuntimeClassApplyConfiguration {
	b.ObjectMeta = value
	return b
}

// RemoveObjectMeta removes the ObjectMeta field from the declarative configuration.
func (b *RuntimeClassApplyConfiguration) RemoveObjectMeta() *RuntimeClassApplyConfiguration {
	b.ObjectMeta = nil
	return b
}

// GetObjectMeta gets the ObjectMeta field from the declarative configuration.
func (b *RuntimeClassApplyConfiguration) GetObjectMeta() (value *v1.ObjectMetaApplyConfiguration, ok bool) {
	return b.ObjectMeta, b.ObjectMeta != nil
}

// SetSpec sets the Spec field in the declarative configuration to the given value.
func (b *RuntimeClassApplyConfiguration) SetSpec(value *RuntimeClassSpecApplyConfiguration) *RuntimeClassApplyConfiguration {
	b.Spec = value
	return b
}

// RemoveSpec removes the Spec field from the declarative configuration.
func (b *RuntimeClassApplyConfiguration) RemoveSpec() *RuntimeClassApplyConfiguration {
	b.Spec = nil
	return b
}

// GetSpec gets the Spec field from the declarative configuration.
func (b *RuntimeClassApplyConfiguration) GetSpec() (value *RuntimeClassSpecApplyConfiguration, ok bool) {
	return b.Spec, b.Spec != nil
}

// RuntimeClassList represents a listAlias of RuntimeClassApplyConfiguration.
type RuntimeClassList []*RuntimeClassApplyConfiguration

// RuntimeClassList represents a map of RuntimeClassApplyConfiguration.
type RuntimeClassMap map[string]RuntimeClassApplyConfiguration
