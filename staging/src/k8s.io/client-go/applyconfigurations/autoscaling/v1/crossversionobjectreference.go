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

// CrossVersionObjectReferenceApplyConfiguration represents an declarative configuration of the CrossVersionObjectReference type for use
// with apply.
type CrossVersionObjectReferenceApplyConfiguration struct {
	Kind       *string `json:"kind,omitempty"`
	Name       *string `json:"name,omitempty"`
	APIVersion *string `json:"apiVersion,omitempty"`
}

// CrossVersionObjectReferenceApplyConfiguration constructs an declarative configuration of the CrossVersionObjectReference type for use with
// apply.
func CrossVersionObjectReference() *CrossVersionObjectReferenceApplyConfiguration {
	return &CrossVersionObjectReferenceApplyConfiguration{}
}

// SetKind sets the Kind field in the declarative configuration to the given value.
func (b *CrossVersionObjectReferenceApplyConfiguration) SetKind(value string) *CrossVersionObjectReferenceApplyConfiguration {
	b.Kind = &value
	return b
}

// RemoveKind removes the Kind field from the declarative configuration.
func (b *CrossVersionObjectReferenceApplyConfiguration) RemoveKind() *CrossVersionObjectReferenceApplyConfiguration {
	b.Kind = nil
	return b
}

// GetKind gets the Kind field from the declarative configuration.
func (b *CrossVersionObjectReferenceApplyConfiguration) GetKind() (value string, ok bool) {
	if v := b.Kind; v != nil {
		return *v, true
	}
	return value, false
}

// SetName sets the Name field in the declarative configuration to the given value.
func (b *CrossVersionObjectReferenceApplyConfiguration) SetName(value string) *CrossVersionObjectReferenceApplyConfiguration {
	b.Name = &value
	return b
}

// RemoveName removes the Name field from the declarative configuration.
func (b *CrossVersionObjectReferenceApplyConfiguration) RemoveName() *CrossVersionObjectReferenceApplyConfiguration {
	b.Name = nil
	return b
}

// GetName gets the Name field from the declarative configuration.
func (b *CrossVersionObjectReferenceApplyConfiguration) GetName() (value string, ok bool) {
	if v := b.Name; v != nil {
		return *v, true
	}
	return value, false
}

// SetAPIVersion sets the APIVersion field in the declarative configuration to the given value.
func (b *CrossVersionObjectReferenceApplyConfiguration) SetAPIVersion(value string) *CrossVersionObjectReferenceApplyConfiguration {
	b.APIVersion = &value
	return b
}

// RemoveAPIVersion removes the APIVersion field from the declarative configuration.
func (b *CrossVersionObjectReferenceApplyConfiguration) RemoveAPIVersion() *CrossVersionObjectReferenceApplyConfiguration {
	b.APIVersion = nil
	return b
}

// GetAPIVersion gets the APIVersion field from the declarative configuration.
func (b *CrossVersionObjectReferenceApplyConfiguration) GetAPIVersion() (value string, ok bool) {
	if v := b.APIVersion; v != nil {
		return *v, true
	}
	return value, false
}

// CrossVersionObjectReferenceList represents a listAlias of CrossVersionObjectReferenceApplyConfiguration.
type CrossVersionObjectReferenceList []*CrossVersionObjectReferenceApplyConfiguration

// CrossVersionObjectReferenceList represents a map of CrossVersionObjectReferenceApplyConfiguration.
type CrossVersionObjectReferenceMap map[string]CrossVersionObjectReferenceApplyConfiguration
