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

// StorageOSVolumeSourceApplyConfiguration represents an declarative configuration of the StorageOSVolumeSource type for use
// with apply.
type StorageOSVolumeSourceApplyConfiguration struct {
	VolumeName      *string                                 `json:"volumeName,omitempty"`
	VolumeNamespace *string                                 `json:"volumeNamespace,omitempty"`
	FSType          *string                                 `json:"fsType,omitempty"`
	ReadOnly        *bool                                   `json:"readOnly,omitempty"`
	SecretRef       *LocalObjectReferenceApplyConfiguration `json:"secretRef,omitempty"`
}

// StorageOSVolumeSourceApplyConfiguration constructs an declarative configuration of the StorageOSVolumeSource type for use with
// apply.
func StorageOSVolumeSource() *StorageOSVolumeSourceApplyConfiguration {
	return &StorageOSVolumeSourceApplyConfiguration{}
}

// SetVolumeName sets the VolumeName field in the declarative configuration to the given value.
func (b *StorageOSVolumeSourceApplyConfiguration) SetVolumeName(value string) *StorageOSVolumeSourceApplyConfiguration {
	b.VolumeName = &value
	return b
}

// RemoveVolumeName removes the VolumeName field from the declarative configuration.
func (b *StorageOSVolumeSourceApplyConfiguration) RemoveVolumeName() *StorageOSVolumeSourceApplyConfiguration {
	b.VolumeName = nil
	return b
}

// GetVolumeName gets the VolumeName field from the declarative configuration.
func (b *StorageOSVolumeSourceApplyConfiguration) GetVolumeName() (value string, ok bool) {
	if v := b.VolumeName; v != nil {
		return *v, true
	}
	return value, false
}

// SetVolumeNamespace sets the VolumeNamespace field in the declarative configuration to the given value.
func (b *StorageOSVolumeSourceApplyConfiguration) SetVolumeNamespace(value string) *StorageOSVolumeSourceApplyConfiguration {
	b.VolumeNamespace = &value
	return b
}

// RemoveVolumeNamespace removes the VolumeNamespace field from the declarative configuration.
func (b *StorageOSVolumeSourceApplyConfiguration) RemoveVolumeNamespace() *StorageOSVolumeSourceApplyConfiguration {
	b.VolumeNamespace = nil
	return b
}

// GetVolumeNamespace gets the VolumeNamespace field from the declarative configuration.
func (b *StorageOSVolumeSourceApplyConfiguration) GetVolumeNamespace() (value string, ok bool) {
	if v := b.VolumeNamespace; v != nil {
		return *v, true
	}
	return value, false
}

// SetFSType sets the FSType field in the declarative configuration to the given value.
func (b *StorageOSVolumeSourceApplyConfiguration) SetFSType(value string) *StorageOSVolumeSourceApplyConfiguration {
	b.FSType = &value
	return b
}

// RemoveFSType removes the FSType field from the declarative configuration.
func (b *StorageOSVolumeSourceApplyConfiguration) RemoveFSType() *StorageOSVolumeSourceApplyConfiguration {
	b.FSType = nil
	return b
}

// GetFSType gets the FSType field from the declarative configuration.
func (b *StorageOSVolumeSourceApplyConfiguration) GetFSType() (value string, ok bool) {
	if v := b.FSType; v != nil {
		return *v, true
	}
	return value, false
}

// SetReadOnly sets the ReadOnly field in the declarative configuration to the given value.
func (b *StorageOSVolumeSourceApplyConfiguration) SetReadOnly(value bool) *StorageOSVolumeSourceApplyConfiguration {
	b.ReadOnly = &value
	return b
}

// RemoveReadOnly removes the ReadOnly field from the declarative configuration.
func (b *StorageOSVolumeSourceApplyConfiguration) RemoveReadOnly() *StorageOSVolumeSourceApplyConfiguration {
	b.ReadOnly = nil
	return b
}

// GetReadOnly gets the ReadOnly field from the declarative configuration.
func (b *StorageOSVolumeSourceApplyConfiguration) GetReadOnly() (value bool, ok bool) {
	if v := b.ReadOnly; v != nil {
		return *v, true
	}
	return value, false
}

// SetSecretRef sets the SecretRef field in the declarative configuration to the given value.
func (b *StorageOSVolumeSourceApplyConfiguration) SetSecretRef(value *LocalObjectReferenceApplyConfiguration) *StorageOSVolumeSourceApplyConfiguration {
	b.SecretRef = value
	return b
}

// RemoveSecretRef removes the SecretRef field from the declarative configuration.
func (b *StorageOSVolumeSourceApplyConfiguration) RemoveSecretRef() *StorageOSVolumeSourceApplyConfiguration {
	b.SecretRef = nil
	return b
}

// GetSecretRef gets the SecretRef field from the declarative configuration.
func (b *StorageOSVolumeSourceApplyConfiguration) GetSecretRef() (value *LocalObjectReferenceApplyConfiguration, ok bool) {
	return b.SecretRef, b.SecretRef != nil
}

// StorageOSVolumeSourceList represents a listAlias of StorageOSVolumeSourceApplyConfiguration.
type StorageOSVolumeSourceList []*StorageOSVolumeSourceApplyConfiguration

// StorageOSVolumeSourceList represents a map of StorageOSVolumeSourceApplyConfiguration.
type StorageOSVolumeSourceMap map[string]StorageOSVolumeSourceApplyConfiguration
