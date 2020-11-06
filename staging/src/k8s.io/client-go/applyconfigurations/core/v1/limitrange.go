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
	v1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

// LimitRangeApplyConfiguration represents an declarative configuration of the LimitRange type for use
// with apply.
type LimitRangeApplyConfiguration struct {
	v1.TypeMetaApplyConfiguration `json:",inline"`
	ObjectMeta                    *v1.ObjectMetaApplyConfiguration  `json:"metadata,omitempty"`
	Spec                          *LimitRangeSpecApplyConfiguration `json:"spec,omitempty"`
}

// LimitRangeApplyConfiguration constructs an declarative configuration of the LimitRange type for use with
// apply.
func LimitRange() *LimitRangeApplyConfiguration {
	return &LimitRangeApplyConfiguration{}
}

// SetTypeMeta sets the TypeMeta field in the declarative configuration to the given value.
func (b *LimitRangeApplyConfiguration) SetTypeMeta(value *v1.TypeMetaApplyConfiguration) *LimitRangeApplyConfiguration {
	if value != nil {
		b.TypeMetaApplyConfiguration = *value
	}
	return b
}

// GetTypeMeta gets the TypeMeta field from the declarative configuration.
func (b *LimitRangeApplyConfiguration) GetTypeMeta() (value *v1.TypeMetaApplyConfiguration, ok bool) {
	return &b.TypeMetaApplyConfiguration, true
}

// SetObjectMeta sets the ObjectMeta field in the declarative configuration to the given value.
func (b *LimitRangeApplyConfiguration) SetObjectMeta(value *v1.ObjectMetaApplyConfiguration) *LimitRangeApplyConfiguration {
	b.ObjectMeta = value
	return b
}

// RemoveObjectMeta removes the ObjectMeta field from the declarative configuration.
func (b *LimitRangeApplyConfiguration) RemoveObjectMeta() *LimitRangeApplyConfiguration {
	b.ObjectMeta = nil
	return b
}

// GetObjectMeta gets the ObjectMeta field from the declarative configuration.
func (b *LimitRangeApplyConfiguration) GetObjectMeta() (value *v1.ObjectMetaApplyConfiguration, ok bool) {
	return b.ObjectMeta, b.ObjectMeta != nil
}

// SetSpec sets the Spec field in the declarative configuration to the given value.
func (b *LimitRangeApplyConfiguration) SetSpec(value *LimitRangeSpecApplyConfiguration) *LimitRangeApplyConfiguration {
	b.Spec = value
	return b
}

// RemoveSpec removes the Spec field from the declarative configuration.
func (b *LimitRangeApplyConfiguration) RemoveSpec() *LimitRangeApplyConfiguration {
	b.Spec = nil
	return b
}

// GetSpec gets the Spec field from the declarative configuration.
func (b *LimitRangeApplyConfiguration) GetSpec() (value *LimitRangeSpecApplyConfiguration, ok bool) {
	return b.Spec, b.Spec != nil
}

// LimitRangeList represents a listAlias of LimitRangeApplyConfiguration.
type LimitRangeList []*LimitRangeApplyConfiguration

// LimitRangeList represents a map of LimitRangeApplyConfiguration.
type LimitRangeMap map[string]LimitRangeApplyConfiguration
