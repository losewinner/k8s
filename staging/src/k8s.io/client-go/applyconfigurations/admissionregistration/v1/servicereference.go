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

// ServiceReferenceApplyConfiguration represents an declarative configuration of the ServiceReference type for use
// with apply.
type ServiceReferenceApplyConfiguration struct {
	Namespace *string `json:"namespace,omitempty"`
	Name      *string `json:"name,omitempty"`
	Path      *string `json:"path,omitempty"`
	Port      *int32  `json:"port,omitempty"`
}

// ServiceReferenceApplyConfiguration constructs an declarative configuration of the ServiceReference type for use with
// apply.
func ServiceReference() *ServiceReferenceApplyConfiguration {
	return &ServiceReferenceApplyConfiguration{}
}

// SetNamespace sets the Namespace field in the declarative configuration to the given value.
func (b *ServiceReferenceApplyConfiguration) SetNamespace(value string) *ServiceReferenceApplyConfiguration {
	b.Namespace = &value
	return b
}

// RemoveNamespace removes the Namespace field from the declarative configuration.
func (b *ServiceReferenceApplyConfiguration) RemoveNamespace() *ServiceReferenceApplyConfiguration {
	b.Namespace = nil
	return b
}

// GetNamespace gets the Namespace field from the declarative configuration.
func (b *ServiceReferenceApplyConfiguration) GetNamespace() (value string, ok bool) {
	if v := b.Namespace; v != nil {
		return *v, true
	}
	return value, false
}

// SetName sets the Name field in the declarative configuration to the given value.
func (b *ServiceReferenceApplyConfiguration) SetName(value string) *ServiceReferenceApplyConfiguration {
	b.Name = &value
	return b
}

// RemoveName removes the Name field from the declarative configuration.
func (b *ServiceReferenceApplyConfiguration) RemoveName() *ServiceReferenceApplyConfiguration {
	b.Name = nil
	return b
}

// GetName gets the Name field from the declarative configuration.
func (b *ServiceReferenceApplyConfiguration) GetName() (value string, ok bool) {
	if v := b.Name; v != nil {
		return *v, true
	}
	return value, false
}

// SetPath sets the Path field in the declarative configuration to the given value.
func (b *ServiceReferenceApplyConfiguration) SetPath(value string) *ServiceReferenceApplyConfiguration {
	b.Path = &value
	return b
}

// RemovePath removes the Path field from the declarative configuration.
func (b *ServiceReferenceApplyConfiguration) RemovePath() *ServiceReferenceApplyConfiguration {
	b.Path = nil
	return b
}

// GetPath gets the Path field from the declarative configuration.
func (b *ServiceReferenceApplyConfiguration) GetPath() (value string, ok bool) {
	if v := b.Path; v != nil {
		return *v, true
	}
	return value, false
}

// SetPort sets the Port field in the declarative configuration to the given value.
func (b *ServiceReferenceApplyConfiguration) SetPort(value int32) *ServiceReferenceApplyConfiguration {
	b.Port = &value
	return b
}

// RemovePort removes the Port field from the declarative configuration.
func (b *ServiceReferenceApplyConfiguration) RemovePort() *ServiceReferenceApplyConfiguration {
	b.Port = nil
	return b
}

// GetPort gets the Port field from the declarative configuration.
func (b *ServiceReferenceApplyConfiguration) GetPort() (value int32, ok bool) {
	if v := b.Port; v != nil {
		return *v, true
	}
	return value, false
}

// ServiceReferenceList represents a listAlias of ServiceReferenceApplyConfiguration.
type ServiceReferenceList []*ServiceReferenceApplyConfiguration

// ServiceReferenceList represents a map of ServiceReferenceApplyConfiguration.
type ServiceReferenceMap map[string]ServiceReferenceApplyConfiguration
