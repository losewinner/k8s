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

// AWSElasticBlockStoreVolumeSourceApplyConfiguration represents an declarative configuration of the AWSElasticBlockStoreVolumeSource type for use
// with apply.
type AWSElasticBlockStoreVolumeSourceApplyConfiguration struct {
	VolumeID  *string `json:"volumeID,omitempty"`
	FSType    *string `json:"fsType,omitempty"`
	Partition *int32  `json:"partition,omitempty"`
	ReadOnly  *bool   `json:"readOnly,omitempty"`
}

// AWSElasticBlockStoreVolumeSourceApplyConfiguration constructs an declarative configuration of the AWSElasticBlockStoreVolumeSource type for use with
// apply.
func AWSElasticBlockStoreVolumeSource() *AWSElasticBlockStoreVolumeSourceApplyConfiguration {
	return &AWSElasticBlockStoreVolumeSourceApplyConfiguration{}
}

// SetVolumeID sets the VolumeID field in the declarative configuration to the given value.
func (b *AWSElasticBlockStoreVolumeSourceApplyConfiguration) SetVolumeID(value string) *AWSElasticBlockStoreVolumeSourceApplyConfiguration {
	b.VolumeID = &value
	return b
}

// RemoveVolumeID removes the VolumeID field from the declarative configuration.
func (b *AWSElasticBlockStoreVolumeSourceApplyConfiguration) RemoveVolumeID() *AWSElasticBlockStoreVolumeSourceApplyConfiguration {
	b.VolumeID = nil
	return b
}

// GetVolumeID gets the VolumeID field from the declarative configuration.
func (b *AWSElasticBlockStoreVolumeSourceApplyConfiguration) GetVolumeID() (value string, ok bool) {
	if v := b.VolumeID; v != nil {
		return *v, true
	}
	return value, false
}

// SetFSType sets the FSType field in the declarative configuration to the given value.
func (b *AWSElasticBlockStoreVolumeSourceApplyConfiguration) SetFSType(value string) *AWSElasticBlockStoreVolumeSourceApplyConfiguration {
	b.FSType = &value
	return b
}

// RemoveFSType removes the FSType field from the declarative configuration.
func (b *AWSElasticBlockStoreVolumeSourceApplyConfiguration) RemoveFSType() *AWSElasticBlockStoreVolumeSourceApplyConfiguration {
	b.FSType = nil
	return b
}

// GetFSType gets the FSType field from the declarative configuration.
func (b *AWSElasticBlockStoreVolumeSourceApplyConfiguration) GetFSType() (value string, ok bool) {
	if v := b.FSType; v != nil {
		return *v, true
	}
	return value, false
}

// SetPartition sets the Partition field in the declarative configuration to the given value.
func (b *AWSElasticBlockStoreVolumeSourceApplyConfiguration) SetPartition(value int32) *AWSElasticBlockStoreVolumeSourceApplyConfiguration {
	b.Partition = &value
	return b
}

// RemovePartition removes the Partition field from the declarative configuration.
func (b *AWSElasticBlockStoreVolumeSourceApplyConfiguration) RemovePartition() *AWSElasticBlockStoreVolumeSourceApplyConfiguration {
	b.Partition = nil
	return b
}

// GetPartition gets the Partition field from the declarative configuration.
func (b *AWSElasticBlockStoreVolumeSourceApplyConfiguration) GetPartition() (value int32, ok bool) {
	if v := b.Partition; v != nil {
		return *v, true
	}
	return value, false
}

// SetReadOnly sets the ReadOnly field in the declarative configuration to the given value.
func (b *AWSElasticBlockStoreVolumeSourceApplyConfiguration) SetReadOnly(value bool) *AWSElasticBlockStoreVolumeSourceApplyConfiguration {
	b.ReadOnly = &value
	return b
}

// RemoveReadOnly removes the ReadOnly field from the declarative configuration.
func (b *AWSElasticBlockStoreVolumeSourceApplyConfiguration) RemoveReadOnly() *AWSElasticBlockStoreVolumeSourceApplyConfiguration {
	b.ReadOnly = nil
	return b
}

// GetReadOnly gets the ReadOnly field from the declarative configuration.
func (b *AWSElasticBlockStoreVolumeSourceApplyConfiguration) GetReadOnly() (value bool, ok bool) {
	if v := b.ReadOnly; v != nil {
		return *v, true
	}
	return value, false
}

// AWSElasticBlockStoreVolumeSourceList represents a listAlias of AWSElasticBlockStoreVolumeSourceApplyConfiguration.
type AWSElasticBlockStoreVolumeSourceList []*AWSElasticBlockStoreVolumeSourceApplyConfiguration

// AWSElasticBlockStoreVolumeSourceList represents a map of AWSElasticBlockStoreVolumeSourceApplyConfiguration.
type AWSElasticBlockStoreVolumeSourceMap map[string]AWSElasticBlockStoreVolumeSourceApplyConfiguration
