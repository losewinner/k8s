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

package v2beta1

import (
	v1 "k8s.io/api/core/v1"
	resource "k8s.io/apimachinery/pkg/api/resource"
)

// ContainerResourceMetricStatusApplyConfiguration represents an declarative configuration of the ContainerResourceMetricStatus type for use
// with apply.
type ContainerResourceMetricStatusApplyConfiguration struct {
	Name                      *v1.ResourceName   `json:"name,omitempty"`
	CurrentAverageUtilization *int32             `json:"currentAverageUtilization,omitempty"`
	CurrentAverageValue       *resource.Quantity `json:"currentAverageValue,omitempty"`
	Container                 *string            `json:"container,omitempty"`
}

// ContainerResourceMetricStatusApplyConfiguration constructs an declarative configuration of the ContainerResourceMetricStatus type for use with
// apply.
func ContainerResourceMetricStatus() *ContainerResourceMetricStatusApplyConfiguration {
	return &ContainerResourceMetricStatusApplyConfiguration{}
}

// SetName sets the Name field in the declarative configuration to the given value.
func (b *ContainerResourceMetricStatusApplyConfiguration) SetName(value v1.ResourceName) *ContainerResourceMetricStatusApplyConfiguration {
	b.Name = &value
	return b
}

// RemoveName removes the Name field from the declarative configuration.
func (b *ContainerResourceMetricStatusApplyConfiguration) RemoveName() *ContainerResourceMetricStatusApplyConfiguration {
	b.Name = nil
	return b
}

// GetName gets the Name field from the declarative configuration.
func (b *ContainerResourceMetricStatusApplyConfiguration) GetName() (value v1.ResourceName, ok bool) {
	if v := b.Name; v != nil {
		return *v, true
	}
	return value, false
}

// SetCurrentAverageUtilization sets the CurrentAverageUtilization field in the declarative configuration to the given value.
func (b *ContainerResourceMetricStatusApplyConfiguration) SetCurrentAverageUtilization(value int32) *ContainerResourceMetricStatusApplyConfiguration {
	b.CurrentAverageUtilization = &value
	return b
}

// RemoveCurrentAverageUtilization removes the CurrentAverageUtilization field from the declarative configuration.
func (b *ContainerResourceMetricStatusApplyConfiguration) RemoveCurrentAverageUtilization() *ContainerResourceMetricStatusApplyConfiguration {
	b.CurrentAverageUtilization = nil
	return b
}

// GetCurrentAverageUtilization gets the CurrentAverageUtilization field from the declarative configuration.
func (b *ContainerResourceMetricStatusApplyConfiguration) GetCurrentAverageUtilization() (value int32, ok bool) {
	if v := b.CurrentAverageUtilization; v != nil {
		return *v, true
	}
	return value, false
}

// SetCurrentAverageValue sets the CurrentAverageValue field in the declarative configuration to the given value.
func (b *ContainerResourceMetricStatusApplyConfiguration) SetCurrentAverageValue(value resource.Quantity) *ContainerResourceMetricStatusApplyConfiguration {
	b.CurrentAverageValue = &value
	return b
}

// RemoveCurrentAverageValue removes the CurrentAverageValue field from the declarative configuration.
func (b *ContainerResourceMetricStatusApplyConfiguration) RemoveCurrentAverageValue() *ContainerResourceMetricStatusApplyConfiguration {
	b.CurrentAverageValue = nil
	return b
}

// GetCurrentAverageValue gets the CurrentAverageValue field from the declarative configuration.
func (b *ContainerResourceMetricStatusApplyConfiguration) GetCurrentAverageValue() (value resource.Quantity, ok bool) {
	if v := b.CurrentAverageValue; v != nil {
		return *v, true
	}
	return value, false
}

// SetContainer sets the Container field in the declarative configuration to the given value.
func (b *ContainerResourceMetricStatusApplyConfiguration) SetContainer(value string) *ContainerResourceMetricStatusApplyConfiguration {
	b.Container = &value
	return b
}

// RemoveContainer removes the Container field from the declarative configuration.
func (b *ContainerResourceMetricStatusApplyConfiguration) RemoveContainer() *ContainerResourceMetricStatusApplyConfiguration {
	b.Container = nil
	return b
}

// GetContainer gets the Container field from the declarative configuration.
func (b *ContainerResourceMetricStatusApplyConfiguration) GetContainer() (value string, ok bool) {
	if v := b.Container; v != nil {
		return *v, true
	}
	return value, false
}

// ContainerResourceMetricStatusList represents a listAlias of ContainerResourceMetricStatusApplyConfiguration.
type ContainerResourceMetricStatusList []*ContainerResourceMetricStatusApplyConfiguration

// ContainerResourceMetricStatusList represents a map of ContainerResourceMetricStatusApplyConfiguration.
type ContainerResourceMetricStatusMap map[string]ContainerResourceMetricStatusApplyConfiguration
