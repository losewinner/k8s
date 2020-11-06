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

// ContainerImageApplyConfiguration represents an declarative configuration of the ContainerImage type for use
// with apply.
type ContainerImageApplyConfiguration struct {
	Names     *[]string `json:"names,omitempty"`
	SizeBytes *int64    `json:"sizeBytes,omitempty"`
}

// ContainerImageApplyConfiguration constructs an declarative configuration of the ContainerImage type for use with
// apply.
func ContainerImage() *ContainerImageApplyConfiguration {
	return &ContainerImageApplyConfiguration{}
}

// SetNames sets the Names field in the declarative configuration to the given value.
func (b *ContainerImageApplyConfiguration) SetNames(value []string) *ContainerImageApplyConfiguration {
	b.Names = &value
	return b
}

// RemoveNames removes the Names field from the declarative configuration.
func (b *ContainerImageApplyConfiguration) RemoveNames() *ContainerImageApplyConfiguration {
	b.Names = nil
	return b
}

// GetNames gets the Names field from the declarative configuration.
func (b *ContainerImageApplyConfiguration) GetNames() (value []string, ok bool) {
	if v := b.Names; v != nil {
		return *v, true
	}
	return value, false
}

// SetSizeBytes sets the SizeBytes field in the declarative configuration to the given value.
func (b *ContainerImageApplyConfiguration) SetSizeBytes(value int64) *ContainerImageApplyConfiguration {
	b.SizeBytes = &value
	return b
}

// RemoveSizeBytes removes the SizeBytes field from the declarative configuration.
func (b *ContainerImageApplyConfiguration) RemoveSizeBytes() *ContainerImageApplyConfiguration {
	b.SizeBytes = nil
	return b
}

// GetSizeBytes gets the SizeBytes field from the declarative configuration.
func (b *ContainerImageApplyConfiguration) GetSizeBytes() (value int64, ok bool) {
	if v := b.SizeBytes; v != nil {
		return *v, true
	}
	return value, false
}

// ContainerImageList represents a listAlias of ContainerImageApplyConfiguration.
type ContainerImageList []*ContainerImageApplyConfiguration

// ContainerImageList represents a map of ContainerImageApplyConfiguration.
type ContainerImageMap map[string]ContainerImageApplyConfiguration
