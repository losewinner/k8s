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

// PodApplyConfiguration represents an declarative configuration of the Pod type for use
// with apply.
type PodApplyConfiguration struct {
	v1.TypeMetaApplyConfiguration `json:",inline"`
	ObjectMeta                    *v1.ObjectMetaApplyConfiguration `json:"metadata,omitempty"`
	Spec                          *PodSpecApplyConfiguration       `json:"spec,omitempty"`
	Status                        *PodStatusApplyConfiguration     `json:"status,omitempty"`
}

// PodApplyConfiguration constructs an declarative configuration of the Pod type for use with
// apply.
func Pod() *PodApplyConfiguration {
	return &PodApplyConfiguration{}
}

// SetTypeMeta sets the TypeMeta field in the declarative configuration to the given value.
func (b *PodApplyConfiguration) SetTypeMeta(value *v1.TypeMetaApplyConfiguration) *PodApplyConfiguration {
	if value != nil {
		b.TypeMetaApplyConfiguration = *value
	}
	return b
}

// GetTypeMeta gets the TypeMeta field from the declarative configuration.
func (b *PodApplyConfiguration) GetTypeMeta() (value *v1.TypeMetaApplyConfiguration, ok bool) {
	return &b.TypeMetaApplyConfiguration, true
}

// SetObjectMeta sets the ObjectMeta field in the declarative configuration to the given value.
func (b *PodApplyConfiguration) SetObjectMeta(value *v1.ObjectMetaApplyConfiguration) *PodApplyConfiguration {
	b.ObjectMeta = value
	return b
}

// RemoveObjectMeta removes the ObjectMeta field from the declarative configuration.
func (b *PodApplyConfiguration) RemoveObjectMeta() *PodApplyConfiguration {
	b.ObjectMeta = nil
	return b
}

// GetObjectMeta gets the ObjectMeta field from the declarative configuration.
func (b *PodApplyConfiguration) GetObjectMeta() (value *v1.ObjectMetaApplyConfiguration, ok bool) {
	return b.ObjectMeta, b.ObjectMeta != nil
}

// SetSpec sets the Spec field in the declarative configuration to the given value.
func (b *PodApplyConfiguration) SetSpec(value *PodSpecApplyConfiguration) *PodApplyConfiguration {
	b.Spec = value
	return b
}

// RemoveSpec removes the Spec field from the declarative configuration.
func (b *PodApplyConfiguration) RemoveSpec() *PodApplyConfiguration {
	b.Spec = nil
	return b
}

// GetSpec gets the Spec field from the declarative configuration.
func (b *PodApplyConfiguration) GetSpec() (value *PodSpecApplyConfiguration, ok bool) {
	return b.Spec, b.Spec != nil
}

// SetStatus sets the Status field in the declarative configuration to the given value.
func (b *PodApplyConfiguration) SetStatus(value *PodStatusApplyConfiguration) *PodApplyConfiguration {
	b.Status = value
	return b
}

// RemoveStatus removes the Status field from the declarative configuration.
func (b *PodApplyConfiguration) RemoveStatus() *PodApplyConfiguration {
	b.Status = nil
	return b
}

// GetStatus gets the Status field from the declarative configuration.
func (b *PodApplyConfiguration) GetStatus() (value *PodStatusApplyConfiguration, ok bool) {
	return b.Status, b.Status != nil
}

// PodList represents a listAlias of PodApplyConfiguration.
type PodList []*PodApplyConfiguration

// PodList represents a map of PodApplyConfiguration.
type PodMap map[string]PodApplyConfiguration
