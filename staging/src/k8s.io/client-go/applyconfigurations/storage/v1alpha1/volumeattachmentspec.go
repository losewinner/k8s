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

// VolumeAttachmentSpecApplyConfiguration represents an declarative configuration of the VolumeAttachmentSpec type for use
// with apply.
type VolumeAttachmentSpecApplyConfiguration struct {
	Attacher *string                                   `json:"attacher,omitempty"`
	Source   *VolumeAttachmentSourceApplyConfiguration `json:"source,omitempty"`
	NodeName *string                                   `json:"nodeName,omitempty"`
}

// VolumeAttachmentSpecApplyConfiguration constructs an declarative configuration of the VolumeAttachmentSpec type for use with
// apply.
func VolumeAttachmentSpec() *VolumeAttachmentSpecApplyConfiguration {
	return &VolumeAttachmentSpecApplyConfiguration{}
}

// SetAttacher sets the Attacher field in the declarative configuration to the given value.
func (b *VolumeAttachmentSpecApplyConfiguration) SetAttacher(value string) *VolumeAttachmentSpecApplyConfiguration {
	b.Attacher = &value
	return b
}

// RemoveAttacher removes the Attacher field from the declarative configuration.
func (b *VolumeAttachmentSpecApplyConfiguration) RemoveAttacher() *VolumeAttachmentSpecApplyConfiguration {
	b.Attacher = nil
	return b
}

// GetAttacher gets the Attacher field from the declarative configuration.
func (b *VolumeAttachmentSpecApplyConfiguration) GetAttacher() (value string, ok bool) {
	if v := b.Attacher; v != nil {
		return *v, true
	}
	return value, false
}

// SetSource sets the Source field in the declarative configuration to the given value.
func (b *VolumeAttachmentSpecApplyConfiguration) SetSource(value *VolumeAttachmentSourceApplyConfiguration) *VolumeAttachmentSpecApplyConfiguration {
	b.Source = value
	return b
}

// RemoveSource removes the Source field from the declarative configuration.
func (b *VolumeAttachmentSpecApplyConfiguration) RemoveSource() *VolumeAttachmentSpecApplyConfiguration {
	b.Source = nil
	return b
}

// GetSource gets the Source field from the declarative configuration.
func (b *VolumeAttachmentSpecApplyConfiguration) GetSource() (value *VolumeAttachmentSourceApplyConfiguration, ok bool) {
	return b.Source, b.Source != nil
}

// SetNodeName sets the NodeName field in the declarative configuration to the given value.
func (b *VolumeAttachmentSpecApplyConfiguration) SetNodeName(value string) *VolumeAttachmentSpecApplyConfiguration {
	b.NodeName = &value
	return b
}

// RemoveNodeName removes the NodeName field from the declarative configuration.
func (b *VolumeAttachmentSpecApplyConfiguration) RemoveNodeName() *VolumeAttachmentSpecApplyConfiguration {
	b.NodeName = nil
	return b
}

// GetNodeName gets the NodeName field from the declarative configuration.
func (b *VolumeAttachmentSpecApplyConfiguration) GetNodeName() (value string, ok bool) {
	if v := b.NodeName; v != nil {
		return *v, true
	}
	return value, false
}

// VolumeAttachmentSpecList represents a listAlias of VolumeAttachmentSpecApplyConfiguration.
type VolumeAttachmentSpecList []*VolumeAttachmentSpecApplyConfiguration

// VolumeAttachmentSpecList represents a map of VolumeAttachmentSpecApplyConfiguration.
type VolumeAttachmentSpecMap map[string]VolumeAttachmentSpecApplyConfiguration
