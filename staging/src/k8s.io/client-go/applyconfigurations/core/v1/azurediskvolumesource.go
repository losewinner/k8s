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
	v1 "k8s.io/api/core/v1"
)

// AzureDiskVolumeSourceApplyConfiguration represents an declarative configuration of the AzureDiskVolumeSource type for use
// with apply.
type AzureDiskVolumeSourceApplyConfiguration struct {
	DiskName    *string                      `json:"diskName,omitempty"`
	DataDiskURI *string                      `json:"diskURI,omitempty"`
	CachingMode *v1.AzureDataDiskCachingMode `json:"cachingMode,omitempty"`
	FSType      *string                      `json:"fsType,omitempty"`
	ReadOnly    *bool                        `json:"readOnly,omitempty"`
	Kind        *v1.AzureDataDiskKind        `json:"kind,omitempty"`
}

// AzureDiskVolumeSourceApplyConfiguration constructs an declarative configuration of the AzureDiskVolumeSource type for use with
// apply.
func AzureDiskVolumeSource() *AzureDiskVolumeSourceApplyConfiguration {
	return &AzureDiskVolumeSourceApplyConfiguration{}
}

// SetDiskName sets the DiskName field in the declarative configuration to the given value.
func (b *AzureDiskVolumeSourceApplyConfiguration) SetDiskName(value string) *AzureDiskVolumeSourceApplyConfiguration {
	b.DiskName = &value
	return b
}

// RemoveDiskName removes the DiskName field from the declarative configuration.
func (b *AzureDiskVolumeSourceApplyConfiguration) RemoveDiskName() *AzureDiskVolumeSourceApplyConfiguration {
	b.DiskName = nil
	return b
}

// GetDiskName gets the DiskName field from the declarative configuration.
func (b *AzureDiskVolumeSourceApplyConfiguration) GetDiskName() (value string, ok bool) {
	if v := b.DiskName; v != nil {
		return *v, true
	}
	return value, false
}

// SetDataDiskURI sets the DataDiskURI field in the declarative configuration to the given value.
func (b *AzureDiskVolumeSourceApplyConfiguration) SetDataDiskURI(value string) *AzureDiskVolumeSourceApplyConfiguration {
	b.DataDiskURI = &value
	return b
}

// RemoveDataDiskURI removes the DataDiskURI field from the declarative configuration.
func (b *AzureDiskVolumeSourceApplyConfiguration) RemoveDataDiskURI() *AzureDiskVolumeSourceApplyConfiguration {
	b.DataDiskURI = nil
	return b
}

// GetDataDiskURI gets the DataDiskURI field from the declarative configuration.
func (b *AzureDiskVolumeSourceApplyConfiguration) GetDataDiskURI() (value string, ok bool) {
	if v := b.DataDiskURI; v != nil {
		return *v, true
	}
	return value, false
}

// SetCachingMode sets the CachingMode field in the declarative configuration to the given value.
func (b *AzureDiskVolumeSourceApplyConfiguration) SetCachingMode(value v1.AzureDataDiskCachingMode) *AzureDiskVolumeSourceApplyConfiguration {
	b.CachingMode = &value
	return b
}

// RemoveCachingMode removes the CachingMode field from the declarative configuration.
func (b *AzureDiskVolumeSourceApplyConfiguration) RemoveCachingMode() *AzureDiskVolumeSourceApplyConfiguration {
	b.CachingMode = nil
	return b
}

// GetCachingMode gets the CachingMode field from the declarative configuration.
func (b *AzureDiskVolumeSourceApplyConfiguration) GetCachingMode() (value v1.AzureDataDiskCachingMode, ok bool) {
	if v := b.CachingMode; v != nil {
		return *v, true
	}
	return value, false
}

// SetFSType sets the FSType field in the declarative configuration to the given value.
func (b *AzureDiskVolumeSourceApplyConfiguration) SetFSType(value string) *AzureDiskVolumeSourceApplyConfiguration {
	b.FSType = &value
	return b
}

// RemoveFSType removes the FSType field from the declarative configuration.
func (b *AzureDiskVolumeSourceApplyConfiguration) RemoveFSType() *AzureDiskVolumeSourceApplyConfiguration {
	b.FSType = nil
	return b
}

// GetFSType gets the FSType field from the declarative configuration.
func (b *AzureDiskVolumeSourceApplyConfiguration) GetFSType() (value string, ok bool) {
	if v := b.FSType; v != nil {
		return *v, true
	}
	return value, false
}

// SetReadOnly sets the ReadOnly field in the declarative configuration to the given value.
func (b *AzureDiskVolumeSourceApplyConfiguration) SetReadOnly(value bool) *AzureDiskVolumeSourceApplyConfiguration {
	b.ReadOnly = &value
	return b
}

// RemoveReadOnly removes the ReadOnly field from the declarative configuration.
func (b *AzureDiskVolumeSourceApplyConfiguration) RemoveReadOnly() *AzureDiskVolumeSourceApplyConfiguration {
	b.ReadOnly = nil
	return b
}

// GetReadOnly gets the ReadOnly field from the declarative configuration.
func (b *AzureDiskVolumeSourceApplyConfiguration) GetReadOnly() (value bool, ok bool) {
	if v := b.ReadOnly; v != nil {
		return *v, true
	}
	return value, false
}

// SetKind sets the Kind field in the declarative configuration to the given value.
func (b *AzureDiskVolumeSourceApplyConfiguration) SetKind(value v1.AzureDataDiskKind) *AzureDiskVolumeSourceApplyConfiguration {
	b.Kind = &value
	return b
}

// RemoveKind removes the Kind field from the declarative configuration.
func (b *AzureDiskVolumeSourceApplyConfiguration) RemoveKind() *AzureDiskVolumeSourceApplyConfiguration {
	b.Kind = nil
	return b
}

// GetKind gets the Kind field from the declarative configuration.
func (b *AzureDiskVolumeSourceApplyConfiguration) GetKind() (value v1.AzureDataDiskKind, ok bool) {
	if v := b.Kind; v != nil {
		return *v, true
	}
	return value, false
}

// AzureDiskVolumeSourceList represents a listAlias of AzureDiskVolumeSourceApplyConfiguration.
type AzureDiskVolumeSourceList []*AzureDiskVolumeSourceApplyConfiguration

// AzureDiskVolumeSourceList represents a map of AzureDiskVolumeSourceApplyConfiguration.
type AzureDiskVolumeSourceMap map[string]AzureDiskVolumeSourceApplyConfiguration
