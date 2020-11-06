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

// MutatingWebhookConfigurationApplyConfiguration represents an declarative configuration of the MutatingWebhookConfiguration type for use
// with apply.
type MutatingWebhookConfigurationApplyConfiguration struct {
	v1.TypeMetaApplyConfiguration `json:",inline"`
	ObjectMeta                    *v1.ObjectMetaApplyConfiguration `json:"metadata,omitempty"`
	Webhooks                      *MutatingWebhookList             `json:"webhooks,omitempty"`
}

// MutatingWebhookConfigurationApplyConfiguration constructs an declarative configuration of the MutatingWebhookConfiguration type for use with
// apply.
func MutatingWebhookConfiguration() *MutatingWebhookConfigurationApplyConfiguration {
	return &MutatingWebhookConfigurationApplyConfiguration{}
}

// SetTypeMeta sets the TypeMeta field in the declarative configuration to the given value.
func (b *MutatingWebhookConfigurationApplyConfiguration) SetTypeMeta(value *v1.TypeMetaApplyConfiguration) *MutatingWebhookConfigurationApplyConfiguration {
	if value != nil {
		b.TypeMetaApplyConfiguration = *value
	}
	return b
}

// GetTypeMeta gets the TypeMeta field from the declarative configuration.
func (b *MutatingWebhookConfigurationApplyConfiguration) GetTypeMeta() (value *v1.TypeMetaApplyConfiguration, ok bool) {
	return &b.TypeMetaApplyConfiguration, true
}

// SetObjectMeta sets the ObjectMeta field in the declarative configuration to the given value.
func (b *MutatingWebhookConfigurationApplyConfiguration) SetObjectMeta(value *v1.ObjectMetaApplyConfiguration) *MutatingWebhookConfigurationApplyConfiguration {
	b.ObjectMeta = value
	return b
}

// RemoveObjectMeta removes the ObjectMeta field from the declarative configuration.
func (b *MutatingWebhookConfigurationApplyConfiguration) RemoveObjectMeta() *MutatingWebhookConfigurationApplyConfiguration {
	b.ObjectMeta = nil
	return b
}

// GetObjectMeta gets the ObjectMeta field from the declarative configuration.
func (b *MutatingWebhookConfigurationApplyConfiguration) GetObjectMeta() (value *v1.ObjectMetaApplyConfiguration, ok bool) {
	return b.ObjectMeta, b.ObjectMeta != nil
}

// SetWebhooks sets the Webhooks field in the declarative configuration to the given value.
func (b *MutatingWebhookConfigurationApplyConfiguration) SetWebhooks(value MutatingWebhookList) *MutatingWebhookConfigurationApplyConfiguration {
	b.Webhooks = &value
	return b
}

// RemoveWebhooks removes the Webhooks field from the declarative configuration.
func (b *MutatingWebhookConfigurationApplyConfiguration) RemoveWebhooks() *MutatingWebhookConfigurationApplyConfiguration {
	b.Webhooks = nil
	return b
}

// GetWebhooks gets the Webhooks field from the declarative configuration.
func (b *MutatingWebhookConfigurationApplyConfiguration) GetWebhooks() (value MutatingWebhookList, ok bool) {
	if v := b.Webhooks; v != nil {
		return *v, true
	}
	return value, false
}

// MutatingWebhookConfigurationList represents a listAlias of MutatingWebhookConfigurationApplyConfiguration.
type MutatingWebhookConfigurationList []*MutatingWebhookConfigurationApplyConfiguration

// MutatingWebhookConfigurationList represents a map of MutatingWebhookConfigurationApplyConfiguration.
type MutatingWebhookConfigurationMap map[string]MutatingWebhookConfigurationApplyConfiguration
