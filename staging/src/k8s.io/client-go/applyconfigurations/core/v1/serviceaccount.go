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
	v1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

// ServiceAccountApplyConfiguration represents an declarative configuration of the ServiceAccount type for use
// with apply.
type ServiceAccountApplyConfiguration struct {
	v1.TypeMetaApplyConfiguration `json:",inline"`
	ObjectMeta                    *v1.ObjectMetaApplyConfiguration `json:"metadata,omitempty"`
	Secrets                       *ObjectReferenceList             `json:"secrets,omitempty"`
	ImagePullSecrets              *LocalObjectReferenceList        `json:"imagePullSecrets,omitempty"`
	AutomountServiceAccountToken  *bool                            `json:"automountServiceAccountToken,omitempty"`
}

// ServiceAccountApplyConfiguration constructs an declarative configuration of the ServiceAccount type for use with
// apply.
func ServiceAccount() *ServiceAccountApplyConfiguration {
	return &ServiceAccountApplyConfiguration{}
}

// SetTypeMeta sets the TypeMeta field in the declarative configuration to the given value.
func (b *ServiceAccountApplyConfiguration) SetTypeMeta(value *v1.TypeMetaApplyConfiguration) *ServiceAccountApplyConfiguration {
	if value != nil {
		b.TypeMetaApplyConfiguration = *value
	}
	return b
}

// GetTypeMeta gets the TypeMeta field from the declarative configuration.
func (b *ServiceAccountApplyConfiguration) GetTypeMeta() (value *v1.TypeMetaApplyConfiguration, ok bool) {
	return &b.TypeMetaApplyConfiguration, true
}

// SetObjectMeta sets the ObjectMeta field in the declarative configuration to the given value.
func (b *ServiceAccountApplyConfiguration) SetObjectMeta(value *v1.ObjectMetaApplyConfiguration) *ServiceAccountApplyConfiguration {
	b.ObjectMeta = value
	return b
}

// RemoveObjectMeta removes the ObjectMeta field from the declarative configuration.
func (b *ServiceAccountApplyConfiguration) RemoveObjectMeta() *ServiceAccountApplyConfiguration {
	b.ObjectMeta = nil
	return b
}

// GetObjectMeta gets the ObjectMeta field from the declarative configuration.
func (b *ServiceAccountApplyConfiguration) GetObjectMeta() (value *v1.ObjectMetaApplyConfiguration, ok bool) {
	return b.ObjectMeta, b.ObjectMeta != nil
}

// SetSecrets sets the Secrets field in the declarative configuration to the given value.
func (b *ServiceAccountApplyConfiguration) SetSecrets(value ObjectReferenceList) *ServiceAccountApplyConfiguration {
	b.Secrets = &value
	return b
}

// RemoveSecrets removes the Secrets field from the declarative configuration.
func (b *ServiceAccountApplyConfiguration) RemoveSecrets() *ServiceAccountApplyConfiguration {
	b.Secrets = nil
	return b
}

// GetSecrets gets the Secrets field from the declarative configuration.
func (b *ServiceAccountApplyConfiguration) GetSecrets() (value ObjectReferenceList, ok bool) {
	if v := b.Secrets; v != nil {
		return *v, true
	}
	return value, false
}

// SetImagePullSecrets sets the ImagePullSecrets field in the declarative configuration to the given value.
func (b *ServiceAccountApplyConfiguration) SetImagePullSecrets(value LocalObjectReferenceList) *ServiceAccountApplyConfiguration {
	b.ImagePullSecrets = &value
	return b
}

// RemoveImagePullSecrets removes the ImagePullSecrets field from the declarative configuration.
func (b *ServiceAccountApplyConfiguration) RemoveImagePullSecrets() *ServiceAccountApplyConfiguration {
	b.ImagePullSecrets = nil
	return b
}

// GetImagePullSecrets gets the ImagePullSecrets field from the declarative configuration.
func (b *ServiceAccountApplyConfiguration) GetImagePullSecrets() (value LocalObjectReferenceList, ok bool) {
	if v := b.ImagePullSecrets; v != nil {
		return *v, true
	}
	return value, false
}

// SetAutomountServiceAccountToken sets the AutomountServiceAccountToken field in the declarative configuration to the given value.
func (b *ServiceAccountApplyConfiguration) SetAutomountServiceAccountToken(value bool) *ServiceAccountApplyConfiguration {
	b.AutomountServiceAccountToken = &value
	return b
}

// RemoveAutomountServiceAccountToken removes the AutomountServiceAccountToken field from the declarative configuration.
func (b *ServiceAccountApplyConfiguration) RemoveAutomountServiceAccountToken() *ServiceAccountApplyConfiguration {
	b.AutomountServiceAccountToken = nil
	return b
}

// GetAutomountServiceAccountToken gets the AutomountServiceAccountToken field from the declarative configuration.
func (b *ServiceAccountApplyConfiguration) GetAutomountServiceAccountToken() (value bool, ok bool) {
	if v := b.AutomountServiceAccountToken; v != nil {
		return *v, true
	}
	return value, false
}

// ServiceAccountList represents a listAlias of ServiceAccountApplyConfiguration.
type ServiceAccountList []*ServiceAccountApplyConfiguration

// ServiceAccountList represents a map of ServiceAccountApplyConfiguration.
type ServiceAccountMap map[string]ServiceAccountApplyConfiguration
