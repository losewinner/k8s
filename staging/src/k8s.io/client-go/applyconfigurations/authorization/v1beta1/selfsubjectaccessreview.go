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

import (
	v1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

// SelfSubjectAccessReviewApplyConfiguration represents an declarative configuration of the SelfSubjectAccessReview type for use
// with apply.
type SelfSubjectAccessReviewApplyConfiguration struct {
	v1.TypeMetaApplyConfiguration `json:",inline"`
	ObjectMeta                    *v1.ObjectMetaApplyConfiguration               `json:"metadata,omitempty"`
	Spec                          *SelfSubjectAccessReviewSpecApplyConfiguration `json:"spec,omitempty"`
	Status                        *SubjectAccessReviewStatusApplyConfiguration   `json:"status,omitempty"`
}

// SelfSubjectAccessReviewApplyConfiguration constructs an declarative configuration of the SelfSubjectAccessReview type for use with
// apply.
func SelfSubjectAccessReview() *SelfSubjectAccessReviewApplyConfiguration {
	return &SelfSubjectAccessReviewApplyConfiguration{}
}

// SetTypeMeta sets the TypeMeta field in the declarative configuration to the given value.
func (b *SelfSubjectAccessReviewApplyConfiguration) SetTypeMeta(value *v1.TypeMetaApplyConfiguration) *SelfSubjectAccessReviewApplyConfiguration {
	if value != nil {
		b.TypeMetaApplyConfiguration = *value
	}
	return b
}

// GetTypeMeta gets the TypeMeta field from the declarative configuration.
func (b *SelfSubjectAccessReviewApplyConfiguration) GetTypeMeta() (value *v1.TypeMetaApplyConfiguration, ok bool) {
	return &b.TypeMetaApplyConfiguration, true
}

// SetObjectMeta sets the ObjectMeta field in the declarative configuration to the given value.
func (b *SelfSubjectAccessReviewApplyConfiguration) SetObjectMeta(value *v1.ObjectMetaApplyConfiguration) *SelfSubjectAccessReviewApplyConfiguration {
	b.ObjectMeta = value
	return b
}

// RemoveObjectMeta removes the ObjectMeta field from the declarative configuration.
func (b *SelfSubjectAccessReviewApplyConfiguration) RemoveObjectMeta() *SelfSubjectAccessReviewApplyConfiguration {
	b.ObjectMeta = nil
	return b
}

// GetObjectMeta gets the ObjectMeta field from the declarative configuration.
func (b *SelfSubjectAccessReviewApplyConfiguration) GetObjectMeta() (value *v1.ObjectMetaApplyConfiguration, ok bool) {
	return b.ObjectMeta, b.ObjectMeta != nil
}

// SetSpec sets the Spec field in the declarative configuration to the given value.
func (b *SelfSubjectAccessReviewApplyConfiguration) SetSpec(value *SelfSubjectAccessReviewSpecApplyConfiguration) *SelfSubjectAccessReviewApplyConfiguration {
	b.Spec = value
	return b
}

// RemoveSpec removes the Spec field from the declarative configuration.
func (b *SelfSubjectAccessReviewApplyConfiguration) RemoveSpec() *SelfSubjectAccessReviewApplyConfiguration {
	b.Spec = nil
	return b
}

// GetSpec gets the Spec field from the declarative configuration.
func (b *SelfSubjectAccessReviewApplyConfiguration) GetSpec() (value *SelfSubjectAccessReviewSpecApplyConfiguration, ok bool) {
	return b.Spec, b.Spec != nil
}

// SetStatus sets the Status field in the declarative configuration to the given value.
func (b *SelfSubjectAccessReviewApplyConfiguration) SetStatus(value *SubjectAccessReviewStatusApplyConfiguration) *SelfSubjectAccessReviewApplyConfiguration {
	b.Status = value
	return b
}

// RemoveStatus removes the Status field from the declarative configuration.
func (b *SelfSubjectAccessReviewApplyConfiguration) RemoveStatus() *SelfSubjectAccessReviewApplyConfiguration {
	b.Status = nil
	return b
}

// GetStatus gets the Status field from the declarative configuration.
func (b *SelfSubjectAccessReviewApplyConfiguration) GetStatus() (value *SubjectAccessReviewStatusApplyConfiguration, ok bool) {
	return b.Status, b.Status != nil
}

// SelfSubjectAccessReviewList represents a listAlias of SelfSubjectAccessReviewApplyConfiguration.
type SelfSubjectAccessReviewList []*SelfSubjectAccessReviewApplyConfiguration

// SelfSubjectAccessReviewList represents a map of SelfSubjectAccessReviewApplyConfiguration.
type SelfSubjectAccessReviewMap map[string]SelfSubjectAccessReviewApplyConfiguration
