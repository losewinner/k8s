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

// IngressServiceBackendApplyConfiguration represents an declarative configuration of the IngressServiceBackend type for use
// with apply.
type IngressServiceBackendApplyConfiguration struct {
	Name *string                               `json:"name,omitempty"`
	Port *ServiceBackendPortApplyConfiguration `json:"port,omitempty"`
}

// IngressServiceBackendApplyConfiguration constructs an declarative configuration of the IngressServiceBackend type for use with
// apply.
func IngressServiceBackend() *IngressServiceBackendApplyConfiguration {
	return &IngressServiceBackendApplyConfiguration{}
}

// SetName sets the Name field in the declarative configuration to the given value.
func (b *IngressServiceBackendApplyConfiguration) SetName(value string) *IngressServiceBackendApplyConfiguration {
	b.Name = &value
	return b
}

// RemoveName removes the Name field from the declarative configuration.
func (b *IngressServiceBackendApplyConfiguration) RemoveName() *IngressServiceBackendApplyConfiguration {
	b.Name = nil
	return b
}

// GetName gets the Name field from the declarative configuration.
func (b *IngressServiceBackendApplyConfiguration) GetName() (value string, ok bool) {
	if v := b.Name; v != nil {
		return *v, true
	}
	return value, false
}

// SetPort sets the Port field in the declarative configuration to the given value.
func (b *IngressServiceBackendApplyConfiguration) SetPort(value *ServiceBackendPortApplyConfiguration) *IngressServiceBackendApplyConfiguration {
	b.Port = value
	return b
}

// RemovePort removes the Port field from the declarative configuration.
func (b *IngressServiceBackendApplyConfiguration) RemovePort() *IngressServiceBackendApplyConfiguration {
	b.Port = nil
	return b
}

// GetPort gets the Port field from the declarative configuration.
func (b *IngressServiceBackendApplyConfiguration) GetPort() (value *ServiceBackendPortApplyConfiguration, ok bool) {
	return b.Port, b.Port != nil
}

// IngressServiceBackendList represents a listAlias of IngressServiceBackendApplyConfiguration.
type IngressServiceBackendList []*IngressServiceBackendApplyConfiguration

// IngressServiceBackendList represents a map of IngressServiceBackendApplyConfiguration.
type IngressServiceBackendMap map[string]IngressServiceBackendApplyConfiguration
