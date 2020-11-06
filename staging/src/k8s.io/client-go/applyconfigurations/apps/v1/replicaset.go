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

// ReplicaSetApplyConfiguration represents an declarative configuration of the ReplicaSet type for use
// with apply.
type ReplicaSetApplyConfiguration struct {
	v1.TypeMetaApplyConfiguration `json:",inline"`
	ObjectMeta                    *v1.ObjectMetaApplyConfiguration    `json:"metadata,omitempty"`
	Spec                          *ReplicaSetSpecApplyConfiguration   `json:"spec,omitempty"`
	Status                        *ReplicaSetStatusApplyConfiguration `json:"status,omitempty"`
}

// ReplicaSetApplyConfiguration constructs an declarative configuration of the ReplicaSet type for use with
// apply.
func ReplicaSet() *ReplicaSetApplyConfiguration {
	return &ReplicaSetApplyConfiguration{}
}

// SetTypeMeta sets the TypeMeta field in the declarative configuration to the given value.
func (b *ReplicaSetApplyConfiguration) SetTypeMeta(value *v1.TypeMetaApplyConfiguration) *ReplicaSetApplyConfiguration {
	if value != nil {
		b.TypeMetaApplyConfiguration = *value
	}
	return b
}

// GetTypeMeta gets the TypeMeta field from the declarative configuration.
func (b *ReplicaSetApplyConfiguration) GetTypeMeta() (value *v1.TypeMetaApplyConfiguration, ok bool) {
	return &b.TypeMetaApplyConfiguration, true
}

// SetObjectMeta sets the ObjectMeta field in the declarative configuration to the given value.
func (b *ReplicaSetApplyConfiguration) SetObjectMeta(value *v1.ObjectMetaApplyConfiguration) *ReplicaSetApplyConfiguration {
	b.ObjectMeta = value
	return b
}

// RemoveObjectMeta removes the ObjectMeta field from the declarative configuration.
func (b *ReplicaSetApplyConfiguration) RemoveObjectMeta() *ReplicaSetApplyConfiguration {
	b.ObjectMeta = nil
	return b
}

// GetObjectMeta gets the ObjectMeta field from the declarative configuration.
func (b *ReplicaSetApplyConfiguration) GetObjectMeta() (value *v1.ObjectMetaApplyConfiguration, ok bool) {
	return b.ObjectMeta, b.ObjectMeta != nil
}

// SetSpec sets the Spec field in the declarative configuration to the given value.
func (b *ReplicaSetApplyConfiguration) SetSpec(value *ReplicaSetSpecApplyConfiguration) *ReplicaSetApplyConfiguration {
	b.Spec = value
	return b
}

// RemoveSpec removes the Spec field from the declarative configuration.
func (b *ReplicaSetApplyConfiguration) RemoveSpec() *ReplicaSetApplyConfiguration {
	b.Spec = nil
	return b
}

// GetSpec gets the Spec field from the declarative configuration.
func (b *ReplicaSetApplyConfiguration) GetSpec() (value *ReplicaSetSpecApplyConfiguration, ok bool) {
	return b.Spec, b.Spec != nil
}

// SetStatus sets the Status field in the declarative configuration to the given value.
func (b *ReplicaSetApplyConfiguration) SetStatus(value *ReplicaSetStatusApplyConfiguration) *ReplicaSetApplyConfiguration {
	b.Status = value
	return b
}

// RemoveStatus removes the Status field from the declarative configuration.
func (b *ReplicaSetApplyConfiguration) RemoveStatus() *ReplicaSetApplyConfiguration {
	b.Status = nil
	return b
}

// GetStatus gets the Status field from the declarative configuration.
func (b *ReplicaSetApplyConfiguration) GetStatus() (value *ReplicaSetStatusApplyConfiguration, ok bool) {
	return b.Status, b.Status != nil
}

// ReplicaSetList represents a listAlias of ReplicaSetApplyConfiguration.
type ReplicaSetList []*ReplicaSetApplyConfiguration

// ReplicaSetList represents a map of ReplicaSetApplyConfiguration.
type ReplicaSetMap map[string]ReplicaSetApplyConfiguration
