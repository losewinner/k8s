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

// VolumeAttachmentStatusApplyConfiguration represents an declarative configuration of the VolumeAttachmentStatus type for use
// with apply.
type VolumeAttachmentStatusApplyConfiguration struct {
	Attached           *bool                          `json:"attached,omitempty"`
	AttachmentMetadata *map[string]string             `json:"attachmentMetadata,omitempty"`
	AttachError        *VolumeErrorApplyConfiguration `json:"attachError,omitempty"`
	DetachError        *VolumeErrorApplyConfiguration `json:"detachError,omitempty"`
}

// VolumeAttachmentStatusApplyConfiguration constructs an declarative configuration of the VolumeAttachmentStatus type for use with
// apply.
func VolumeAttachmentStatus() *VolumeAttachmentStatusApplyConfiguration {
	return &VolumeAttachmentStatusApplyConfiguration{}
}

// SetAttached sets the Attached field in the declarative configuration to the given value.
func (b *VolumeAttachmentStatusApplyConfiguration) SetAttached(value bool) *VolumeAttachmentStatusApplyConfiguration {
	b.Attached = &value
	return b
}

// RemoveAttached removes the Attached field from the declarative configuration.
func (b *VolumeAttachmentStatusApplyConfiguration) RemoveAttached() *VolumeAttachmentStatusApplyConfiguration {
	b.Attached = nil
	return b
}

// GetAttached gets the Attached field from the declarative configuration.
func (b *VolumeAttachmentStatusApplyConfiguration) GetAttached() (value bool, ok bool) {
	if v := b.Attached; v != nil {
		return *v, true
	}
	return value, false
}

// SetAttachmentMetadata sets the AttachmentMetadata field in the declarative configuration to the given value.
func (b *VolumeAttachmentStatusApplyConfiguration) SetAttachmentMetadata(value map[string]string) *VolumeAttachmentStatusApplyConfiguration {
	b.AttachmentMetadata = &value
	return b
}

// RemoveAttachmentMetadata removes the AttachmentMetadata field from the declarative configuration.
func (b *VolumeAttachmentStatusApplyConfiguration) RemoveAttachmentMetadata() *VolumeAttachmentStatusApplyConfiguration {
	b.AttachmentMetadata = nil
	return b
}

// GetAttachmentMetadata gets the AttachmentMetadata field from the declarative configuration.
func (b *VolumeAttachmentStatusApplyConfiguration) GetAttachmentMetadata() (value map[string]string, ok bool) {
	if v := b.AttachmentMetadata; v != nil {
		return *v, true
	}
	return value, false
}

// SetAttachError sets the AttachError field in the declarative configuration to the given value.
func (b *VolumeAttachmentStatusApplyConfiguration) SetAttachError(value *VolumeErrorApplyConfiguration) *VolumeAttachmentStatusApplyConfiguration {
	b.AttachError = value
	return b
}

// RemoveAttachError removes the AttachError field from the declarative configuration.
func (b *VolumeAttachmentStatusApplyConfiguration) RemoveAttachError() *VolumeAttachmentStatusApplyConfiguration {
	b.AttachError = nil
	return b
}

// GetAttachError gets the AttachError field from the declarative configuration.
func (b *VolumeAttachmentStatusApplyConfiguration) GetAttachError() (value *VolumeErrorApplyConfiguration, ok bool) {
	return b.AttachError, b.AttachError != nil
}

// SetDetachError sets the DetachError field in the declarative configuration to the given value.
func (b *VolumeAttachmentStatusApplyConfiguration) SetDetachError(value *VolumeErrorApplyConfiguration) *VolumeAttachmentStatusApplyConfiguration {
	b.DetachError = value
	return b
}

// RemoveDetachError removes the DetachError field from the declarative configuration.
func (b *VolumeAttachmentStatusApplyConfiguration) RemoveDetachError() *VolumeAttachmentStatusApplyConfiguration {
	b.DetachError = nil
	return b
}

// GetDetachError gets the DetachError field from the declarative configuration.
func (b *VolumeAttachmentStatusApplyConfiguration) GetDetachError() (value *VolumeErrorApplyConfiguration, ok bool) {
	return b.DetachError, b.DetachError != nil
}

// VolumeAttachmentStatusList represents a listAlias of VolumeAttachmentStatusApplyConfiguration.
type VolumeAttachmentStatusList []*VolumeAttachmentStatusApplyConfiguration

// VolumeAttachmentStatusList represents a map of VolumeAttachmentStatusApplyConfiguration.
type VolumeAttachmentStatusMap map[string]VolumeAttachmentStatusApplyConfiguration
