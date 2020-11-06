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
	resource "k8s.io/apimachinery/pkg/api/resource"
	v1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

// CSIStorageCapacityApplyConfiguration represents an declarative configuration of the CSIStorageCapacity type for use
// with apply.
type CSIStorageCapacityApplyConfiguration struct {
	v1.TypeMetaApplyConfiguration `json:",inline"`
	ObjectMeta                    *v1.ObjectMetaApplyConfiguration    `json:"metadata,omitempty"`
	NodeTopology                  *v1.LabelSelectorApplyConfiguration `json:"nodeTopology,omitempty"`
	StorageClassName              *string                             `json:"storageClassName,omitempty"`
	Capacity                      *resource.Quantity                  `json:"capacity,omitempty"`
}

// CSIStorageCapacityApplyConfiguration constructs an declarative configuration of the CSIStorageCapacity type for use with
// apply.
func CSIStorageCapacity() *CSIStorageCapacityApplyConfiguration {
	return &CSIStorageCapacityApplyConfiguration{}
}

// SetTypeMeta sets the TypeMeta field in the declarative configuration to the given value.
func (b *CSIStorageCapacityApplyConfiguration) SetTypeMeta(value *v1.TypeMetaApplyConfiguration) *CSIStorageCapacityApplyConfiguration {
	if value != nil {
		b.TypeMetaApplyConfiguration = *value
	}
	return b
}

// GetTypeMeta gets the TypeMeta field from the declarative configuration.
func (b *CSIStorageCapacityApplyConfiguration) GetTypeMeta() (value *v1.TypeMetaApplyConfiguration, ok bool) {
	return &b.TypeMetaApplyConfiguration, true
}

// SetObjectMeta sets the ObjectMeta field in the declarative configuration to the given value.
func (b *CSIStorageCapacityApplyConfiguration) SetObjectMeta(value *v1.ObjectMetaApplyConfiguration) *CSIStorageCapacityApplyConfiguration {
	b.ObjectMeta = value
	return b
}

// RemoveObjectMeta removes the ObjectMeta field from the declarative configuration.
func (b *CSIStorageCapacityApplyConfiguration) RemoveObjectMeta() *CSIStorageCapacityApplyConfiguration {
	b.ObjectMeta = nil
	return b
}

// GetObjectMeta gets the ObjectMeta field from the declarative configuration.
func (b *CSIStorageCapacityApplyConfiguration) GetObjectMeta() (value *v1.ObjectMetaApplyConfiguration, ok bool) {
	return b.ObjectMeta, b.ObjectMeta != nil
}

// SetNodeTopology sets the NodeTopology field in the declarative configuration to the given value.
func (b *CSIStorageCapacityApplyConfiguration) SetNodeTopology(value *v1.LabelSelectorApplyConfiguration) *CSIStorageCapacityApplyConfiguration {
	b.NodeTopology = value
	return b
}

// RemoveNodeTopology removes the NodeTopology field from the declarative configuration.
func (b *CSIStorageCapacityApplyConfiguration) RemoveNodeTopology() *CSIStorageCapacityApplyConfiguration {
	b.NodeTopology = nil
	return b
}

// GetNodeTopology gets the NodeTopology field from the declarative configuration.
func (b *CSIStorageCapacityApplyConfiguration) GetNodeTopology() (value *v1.LabelSelectorApplyConfiguration, ok bool) {
	return b.NodeTopology, b.NodeTopology != nil
}

// SetStorageClassName sets the StorageClassName field in the declarative configuration to the given value.
func (b *CSIStorageCapacityApplyConfiguration) SetStorageClassName(value string) *CSIStorageCapacityApplyConfiguration {
	b.StorageClassName = &value
	return b
}

// RemoveStorageClassName removes the StorageClassName field from the declarative configuration.
func (b *CSIStorageCapacityApplyConfiguration) RemoveStorageClassName() *CSIStorageCapacityApplyConfiguration {
	b.StorageClassName = nil
	return b
}

// GetStorageClassName gets the StorageClassName field from the declarative configuration.
func (b *CSIStorageCapacityApplyConfiguration) GetStorageClassName() (value string, ok bool) {
	if v := b.StorageClassName; v != nil {
		return *v, true
	}
	return value, false
}

// SetCapacity sets the Capacity field in the declarative configuration to the given value.
func (b *CSIStorageCapacityApplyConfiguration) SetCapacity(value resource.Quantity) *CSIStorageCapacityApplyConfiguration {
	b.Capacity = &value
	return b
}

// RemoveCapacity removes the Capacity field from the declarative configuration.
func (b *CSIStorageCapacityApplyConfiguration) RemoveCapacity() *CSIStorageCapacityApplyConfiguration {
	b.Capacity = nil
	return b
}

// GetCapacity gets the Capacity field from the declarative configuration.
func (b *CSIStorageCapacityApplyConfiguration) GetCapacity() (value resource.Quantity, ok bool) {
	if v := b.Capacity; v != nil {
		return *v, true
	}
	return value, false
}

// CSIStorageCapacityList represents a listAlias of CSIStorageCapacityApplyConfiguration.
type CSIStorageCapacityList []*CSIStorageCapacityApplyConfiguration

// CSIStorageCapacityList represents a map of CSIStorageCapacityApplyConfiguration.
type CSIStorageCapacityMap map[string]CSIStorageCapacityApplyConfiguration
