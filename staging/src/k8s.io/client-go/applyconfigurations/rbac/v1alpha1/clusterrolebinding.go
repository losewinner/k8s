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

import (
	v1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

// ClusterRoleBindingApplyConfiguration represents an declarative configuration of the ClusterRoleBinding type for use
// with apply.
type ClusterRoleBindingApplyConfiguration struct {
	v1.TypeMetaApplyConfiguration `json:",inline"`
	ObjectMeta                    *v1.ObjectMetaApplyConfiguration `json:"metadata,omitempty"`
	Subjects                      *SubjectList                     `json:"subjects,omitempty"`
	RoleRef                       *RoleRefApplyConfiguration       `json:"roleRef,omitempty"`
}

// ClusterRoleBindingApplyConfiguration constructs an declarative configuration of the ClusterRoleBinding type for use with
// apply.
func ClusterRoleBinding() *ClusterRoleBindingApplyConfiguration {
	return &ClusterRoleBindingApplyConfiguration{}
}

// SetTypeMeta sets the TypeMeta field in the declarative configuration to the given value.
func (b *ClusterRoleBindingApplyConfiguration) SetTypeMeta(value *v1.TypeMetaApplyConfiguration) *ClusterRoleBindingApplyConfiguration {
	if value != nil {
		b.TypeMetaApplyConfiguration = *value
	}
	return b
}

// GetTypeMeta gets the TypeMeta field from the declarative configuration.
func (b *ClusterRoleBindingApplyConfiguration) GetTypeMeta() (value *v1.TypeMetaApplyConfiguration, ok bool) {
	return &b.TypeMetaApplyConfiguration, true
}

// SetObjectMeta sets the ObjectMeta field in the declarative configuration to the given value.
func (b *ClusterRoleBindingApplyConfiguration) SetObjectMeta(value *v1.ObjectMetaApplyConfiguration) *ClusterRoleBindingApplyConfiguration {
	b.ObjectMeta = value
	return b
}

// RemoveObjectMeta removes the ObjectMeta field from the declarative configuration.
func (b *ClusterRoleBindingApplyConfiguration) RemoveObjectMeta() *ClusterRoleBindingApplyConfiguration {
	b.ObjectMeta = nil
	return b
}

// GetObjectMeta gets the ObjectMeta field from the declarative configuration.
func (b *ClusterRoleBindingApplyConfiguration) GetObjectMeta() (value *v1.ObjectMetaApplyConfiguration, ok bool) {
	return b.ObjectMeta, b.ObjectMeta != nil
}

// SetSubjects sets the Subjects field in the declarative configuration to the given value.
func (b *ClusterRoleBindingApplyConfiguration) SetSubjects(value SubjectList) *ClusterRoleBindingApplyConfiguration {
	b.Subjects = &value
	return b
}

// RemoveSubjects removes the Subjects field from the declarative configuration.
func (b *ClusterRoleBindingApplyConfiguration) RemoveSubjects() *ClusterRoleBindingApplyConfiguration {
	b.Subjects = nil
	return b
}

// GetSubjects gets the Subjects field from the declarative configuration.
func (b *ClusterRoleBindingApplyConfiguration) GetSubjects() (value SubjectList, ok bool) {
	if v := b.Subjects; v != nil {
		return *v, true
	}
	return value, false
}

// SetRoleRef sets the RoleRef field in the declarative configuration to the given value.
func (b *ClusterRoleBindingApplyConfiguration) SetRoleRef(value *RoleRefApplyConfiguration) *ClusterRoleBindingApplyConfiguration {
	b.RoleRef = value
	return b
}

// RemoveRoleRef removes the RoleRef field from the declarative configuration.
func (b *ClusterRoleBindingApplyConfiguration) RemoveRoleRef() *ClusterRoleBindingApplyConfiguration {
	b.RoleRef = nil
	return b
}

// GetRoleRef gets the RoleRef field from the declarative configuration.
func (b *ClusterRoleBindingApplyConfiguration) GetRoleRef() (value *RoleRefApplyConfiguration, ok bool) {
	return b.RoleRef, b.RoleRef != nil
}

// ClusterRoleBindingList represents a listAlias of ClusterRoleBindingApplyConfiguration.
type ClusterRoleBindingList []*ClusterRoleBindingApplyConfiguration

// ClusterRoleBindingList represents a map of ClusterRoleBindingApplyConfiguration.
type ClusterRoleBindingMap map[string]ClusterRoleBindingApplyConfiguration
