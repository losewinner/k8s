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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// TaintApplyConfiguration represents an declarative configuration of the Taint type for use
// with apply.
type TaintApplyConfiguration struct {
	Key       *string         `json:"key,omitempty"`
	Value     *string         `json:"value,omitempty"`
	Effect    *v1.TaintEffect `json:"effect,omitempty"`
	TimeAdded *metav1.Time    `json:"timeAdded,omitempty"`
}

// TaintApplyConfiguration constructs an declarative configuration of the Taint type for use with
// apply.
func Taint() *TaintApplyConfiguration {
	return &TaintApplyConfiguration{}
}

// SetKey sets the Key field in the declarative configuration to the given value.
func (b *TaintApplyConfiguration) SetKey(value string) *TaintApplyConfiguration {
	b.Key = &value
	return b
}

// RemoveKey removes the Key field from the declarative configuration.
func (b *TaintApplyConfiguration) RemoveKey() *TaintApplyConfiguration {
	b.Key = nil
	return b
}

// GetKey gets the Key field from the declarative configuration.
func (b *TaintApplyConfiguration) GetKey() (value string, ok bool) {
	if v := b.Key; v != nil {
		return *v, true
	}
	return value, false
}

// SetValue sets the Value field in the declarative configuration to the given value.
func (b *TaintApplyConfiguration) SetValue(value string) *TaintApplyConfiguration {
	b.Value = &value
	return b
}

// RemoveValue removes the Value field from the declarative configuration.
func (b *TaintApplyConfiguration) RemoveValue() *TaintApplyConfiguration {
	b.Value = nil
	return b
}

// GetValue gets the Value field from the declarative configuration.
func (b *TaintApplyConfiguration) GetValue() (value string, ok bool) {
	if v := b.Value; v != nil {
		return *v, true
	}
	return value, false
}

// SetEffect sets the Effect field in the declarative configuration to the given value.
func (b *TaintApplyConfiguration) SetEffect(value v1.TaintEffect) *TaintApplyConfiguration {
	b.Effect = &value
	return b
}

// RemoveEffect removes the Effect field from the declarative configuration.
func (b *TaintApplyConfiguration) RemoveEffect() *TaintApplyConfiguration {
	b.Effect = nil
	return b
}

// GetEffect gets the Effect field from the declarative configuration.
func (b *TaintApplyConfiguration) GetEffect() (value v1.TaintEffect, ok bool) {
	if v := b.Effect; v != nil {
		return *v, true
	}
	return value, false
}

// SetTimeAdded sets the TimeAdded field in the declarative configuration to the given value.
func (b *TaintApplyConfiguration) SetTimeAdded(value metav1.Time) *TaintApplyConfiguration {
	b.TimeAdded = &value
	return b
}

// RemoveTimeAdded removes the TimeAdded field from the declarative configuration.
func (b *TaintApplyConfiguration) RemoveTimeAdded() *TaintApplyConfiguration {
	b.TimeAdded = nil
	return b
}

// GetTimeAdded gets the TimeAdded field from the declarative configuration.
func (b *TaintApplyConfiguration) GetTimeAdded() (value metav1.Time, ok bool) {
	if v := b.TimeAdded; v != nil {
		return *v, true
	}
	return value, false
}

// TaintList represents a listAlias of TaintApplyConfiguration.
type TaintList []*TaintApplyConfiguration

// TaintList represents a map of TaintApplyConfiguration.
type TaintMap map[string]TaintApplyConfiguration
