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

package v1beta1

import (
	intstr "k8s.io/apimachinery/pkg/util/intstr"
	v1 "k8s.io/client-go/applyconfigurations/core/v1"
)

// IngressBackendApplyConfiguration represents an declarative configuration of the IngressBackend type for use
// with apply.
type IngressBackendApplyConfiguration struct {
	ServiceName *string                                         `json:"serviceName,omitempty"`
	ServicePort *intstr.IntOrString                             `json:"servicePort,omitempty"`
	Resource    *v1.TypedLocalObjectReferenceApplyConfiguration `json:"resource,omitempty"`
}

// IngressBackendApplyConfiguration constructs an declarative configuration of the IngressBackend type for use with
// apply.
func IngressBackend() *IngressBackendApplyConfiguration {
	return &IngressBackendApplyConfiguration{}
}

// SetServiceName sets the ServiceName field in the declarative configuration to the given value.
func (b *IngressBackendApplyConfiguration) SetServiceName(value string) *IngressBackendApplyConfiguration {
	b.ServiceName = &value
	return b
}

// RemoveServiceName removes the ServiceName field from the declarative configuration.
func (b *IngressBackendApplyConfiguration) RemoveServiceName() *IngressBackendApplyConfiguration {
	b.ServiceName = nil
	return b
}

// GetServiceName gets the ServiceName field from the declarative configuration.
func (b *IngressBackendApplyConfiguration) GetServiceName() (value string, ok bool) {
	if v := b.ServiceName; v != nil {
		return *v, true
	}
	return value, false
}

// SetServicePort sets the ServicePort field in the declarative configuration to the given value.
func (b *IngressBackendApplyConfiguration) SetServicePort(value intstr.IntOrString) *IngressBackendApplyConfiguration {
	b.ServicePort = &value
	return b
}

// RemoveServicePort removes the ServicePort field from the declarative configuration.
func (b *IngressBackendApplyConfiguration) RemoveServicePort() *IngressBackendApplyConfiguration {
	b.ServicePort = nil
	return b
}

// GetServicePort gets the ServicePort field from the declarative configuration.
func (b *IngressBackendApplyConfiguration) GetServicePort() (value intstr.IntOrString, ok bool) {
	if v := b.ServicePort; v != nil {
		return *v, true
	}
	return value, false
}

// SetResource sets the Resource field in the declarative configuration to the given value.
func (b *IngressBackendApplyConfiguration) SetResource(value *v1.TypedLocalObjectReferenceApplyConfiguration) *IngressBackendApplyConfiguration {
	b.Resource = value
	return b
}

// RemoveResource removes the Resource field from the declarative configuration.
func (b *IngressBackendApplyConfiguration) RemoveResource() *IngressBackendApplyConfiguration {
	b.Resource = nil
	return b
}

// GetResource gets the Resource field from the declarative configuration.
func (b *IngressBackendApplyConfiguration) GetResource() (value *v1.TypedLocalObjectReferenceApplyConfiguration, ok bool) {
	return b.Resource, b.Resource != nil
}

// IngressBackendList represents a listAlias of IngressBackendApplyConfiguration.
type IngressBackendList []*IngressBackendApplyConfiguration

// IngressBackendList represents a map of IngressBackendApplyConfiguration.
type IngressBackendMap map[string]IngressBackendApplyConfiguration
