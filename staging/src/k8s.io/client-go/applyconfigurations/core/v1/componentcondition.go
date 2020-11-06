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

// ComponentConditionApplyConfiguration represents an declarative configuration of the ComponentCondition type for use
// with apply.
type ComponentConditionApplyConfiguration struct {
	Type    *v1.ComponentConditionType `json:"type,omitempty"`
	Status  *v1.ConditionStatus        `json:"status,omitempty"`
	Message *string                    `json:"message,omitempty"`
	Error   *string                    `json:"error,omitempty"`
}

// ComponentConditionApplyConfiguration constructs an declarative configuration of the ComponentCondition type for use with
// apply.
func ComponentCondition() *ComponentConditionApplyConfiguration {
	return &ComponentConditionApplyConfiguration{}
}

// SetType sets the Type field in the declarative configuration to the given value.
func (b *ComponentConditionApplyConfiguration) SetType(value v1.ComponentConditionType) *ComponentConditionApplyConfiguration {
	b.Type = &value
	return b
}

// RemoveType removes the Type field from the declarative configuration.
func (b *ComponentConditionApplyConfiguration) RemoveType() *ComponentConditionApplyConfiguration {
	b.Type = nil
	return b
}

// GetType gets the Type field from the declarative configuration.
func (b *ComponentConditionApplyConfiguration) GetType() (value v1.ComponentConditionType, ok bool) {
	if v := b.Type; v != nil {
		return *v, true
	}
	return value, false
}

// SetStatus sets the Status field in the declarative configuration to the given value.
func (b *ComponentConditionApplyConfiguration) SetStatus(value v1.ConditionStatus) *ComponentConditionApplyConfiguration {
	b.Status = &value
	return b
}

// RemoveStatus removes the Status field from the declarative configuration.
func (b *ComponentConditionApplyConfiguration) RemoveStatus() *ComponentConditionApplyConfiguration {
	b.Status = nil
	return b
}

// GetStatus gets the Status field from the declarative configuration.
func (b *ComponentConditionApplyConfiguration) GetStatus() (value v1.ConditionStatus, ok bool) {
	if v := b.Status; v != nil {
		return *v, true
	}
	return value, false
}

// SetMessage sets the Message field in the declarative configuration to the given value.
func (b *ComponentConditionApplyConfiguration) SetMessage(value string) *ComponentConditionApplyConfiguration {
	b.Message = &value
	return b
}

// RemoveMessage removes the Message field from the declarative configuration.
func (b *ComponentConditionApplyConfiguration) RemoveMessage() *ComponentConditionApplyConfiguration {
	b.Message = nil
	return b
}

// GetMessage gets the Message field from the declarative configuration.
func (b *ComponentConditionApplyConfiguration) GetMessage() (value string, ok bool) {
	if v := b.Message; v != nil {
		return *v, true
	}
	return value, false
}

// SetError sets the Error field in the declarative configuration to the given value.
func (b *ComponentConditionApplyConfiguration) SetError(value string) *ComponentConditionApplyConfiguration {
	b.Error = &value
	return b
}

// RemoveError removes the Error field from the declarative configuration.
func (b *ComponentConditionApplyConfiguration) RemoveError() *ComponentConditionApplyConfiguration {
	b.Error = nil
	return b
}

// GetError gets the Error field from the declarative configuration.
func (b *ComponentConditionApplyConfiguration) GetError() (value string, ok bool) {
	if v := b.Error; v != nil {
		return *v, true
	}
	return value, false
}

// ComponentConditionList represents a listAlias of ComponentConditionApplyConfiguration.
type ComponentConditionList []*ComponentConditionApplyConfiguration

// ComponentConditionList represents a map of ComponentConditionApplyConfiguration.
type ComponentConditionMap map[string]ComponentConditionApplyConfiguration
