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

// ServiceApplyConfiguration represents an declarative configuration of the Service type for use
// with apply.
type ServiceApplyConfiguration struct {
	v1.TypeMetaApplyConfiguration `json:",inline"`
	ObjectMeta                    *v1.ObjectMetaApplyConfiguration `json:"metadata,omitempty"`
	Spec                          *ServiceSpecApplyConfiguration   `json:"spec,omitempty"`
	Status                        *ServiceStatusApplyConfiguration `json:"status,omitempty"`
}

// ServiceApplyConfiguration constructs an declarative configuration of the Service type for use with
// apply.
func Service() *ServiceApplyConfiguration {
	return &ServiceApplyConfiguration{}
}

// SetTypeMeta sets the TypeMeta field in the declarative configuration to the given value.
func (b *ServiceApplyConfiguration) SetTypeMeta(value *v1.TypeMetaApplyConfiguration) *ServiceApplyConfiguration {
	if value != nil {
		b.TypeMetaApplyConfiguration = *value
	}
	return b
}

// GetTypeMeta gets the TypeMeta field from the declarative configuration.
func (b *ServiceApplyConfiguration) GetTypeMeta() (value *v1.TypeMetaApplyConfiguration, ok bool) {
	return &b.TypeMetaApplyConfiguration, true
}

// SetObjectMeta sets the ObjectMeta field in the declarative configuration to the given value.
func (b *ServiceApplyConfiguration) SetObjectMeta(value *v1.ObjectMetaApplyConfiguration) *ServiceApplyConfiguration {
	b.ObjectMeta = value
	return b
}

// RemoveObjectMeta removes the ObjectMeta field from the declarative configuration.
func (b *ServiceApplyConfiguration) RemoveObjectMeta() *ServiceApplyConfiguration {
	b.ObjectMeta = nil
	return b
}

// GetObjectMeta gets the ObjectMeta field from the declarative configuration.
func (b *ServiceApplyConfiguration) GetObjectMeta() (value *v1.ObjectMetaApplyConfiguration, ok bool) {
	return b.ObjectMeta, b.ObjectMeta != nil
}

// SetSpec sets the Spec field in the declarative configuration to the given value.
func (b *ServiceApplyConfiguration) SetSpec(value *ServiceSpecApplyConfiguration) *ServiceApplyConfiguration {
	b.Spec = value
	return b
}

// RemoveSpec removes the Spec field from the declarative configuration.
func (b *ServiceApplyConfiguration) RemoveSpec() *ServiceApplyConfiguration {
	b.Spec = nil
	return b
}

// GetSpec gets the Spec field from the declarative configuration.
func (b *ServiceApplyConfiguration) GetSpec() (value *ServiceSpecApplyConfiguration, ok bool) {
	return b.Spec, b.Spec != nil
}

// SetStatus sets the Status field in the declarative configuration to the given value.
func (b *ServiceApplyConfiguration) SetStatus(value *ServiceStatusApplyConfiguration) *ServiceApplyConfiguration {
	b.Status = value
	return b
}

// RemoveStatus removes the Status field from the declarative configuration.
func (b *ServiceApplyConfiguration) RemoveStatus() *ServiceApplyConfiguration {
	b.Status = nil
	return b
}

// GetStatus gets the Status field from the declarative configuration.
func (b *ServiceApplyConfiguration) GetStatus() (value *ServiceStatusApplyConfiguration, ok bool) {
	return b.Status, b.Status != nil
}

// ServiceList represents a listAlias of ServiceApplyConfiguration.
type ServiceList []*ServiceApplyConfiguration

// ServiceList represents a map of ServiceApplyConfiguration.
type ServiceMap map[string]ServiceApplyConfiguration
