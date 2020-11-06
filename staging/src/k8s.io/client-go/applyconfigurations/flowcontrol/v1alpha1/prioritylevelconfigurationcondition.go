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
	v1alpha1 "k8s.io/api/flowcontrol/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PriorityLevelConfigurationConditionApplyConfiguration represents an declarative configuration of the PriorityLevelConfigurationCondition type for use
// with apply.
type PriorityLevelConfigurationConditionApplyConfiguration struct {
	Type               *v1alpha1.PriorityLevelConfigurationConditionType `json:"type,omitempty"`
	Status             *v1alpha1.ConditionStatus                         `json:"status,omitempty"`
	LastTransitionTime *v1.Time                                          `json:"lastTransitionTime,omitempty"`
	Reason             *string                                           `json:"reason,omitempty"`
	Message            *string                                           `json:"message,omitempty"`
}

// PriorityLevelConfigurationConditionApplyConfiguration constructs an declarative configuration of the PriorityLevelConfigurationCondition type for use with
// apply.
func PriorityLevelConfigurationCondition() *PriorityLevelConfigurationConditionApplyConfiguration {
	return &PriorityLevelConfigurationConditionApplyConfiguration{}
}

// SetType sets the Type field in the declarative configuration to the given value.
func (b *PriorityLevelConfigurationConditionApplyConfiguration) SetType(value v1alpha1.PriorityLevelConfigurationConditionType) *PriorityLevelConfigurationConditionApplyConfiguration {
	b.Type = &value
	return b
}

// RemoveType removes the Type field from the declarative configuration.
func (b *PriorityLevelConfigurationConditionApplyConfiguration) RemoveType() *PriorityLevelConfigurationConditionApplyConfiguration {
	b.Type = nil
	return b
}

// GetType gets the Type field from the declarative configuration.
func (b *PriorityLevelConfigurationConditionApplyConfiguration) GetType() (value v1alpha1.PriorityLevelConfigurationConditionType, ok bool) {
	if v := b.Type; v != nil {
		return *v, true
	}
	return value, false
}

// SetStatus sets the Status field in the declarative configuration to the given value.
func (b *PriorityLevelConfigurationConditionApplyConfiguration) SetStatus(value v1alpha1.ConditionStatus) *PriorityLevelConfigurationConditionApplyConfiguration {
	b.Status = &value
	return b
}

// RemoveStatus removes the Status field from the declarative configuration.
func (b *PriorityLevelConfigurationConditionApplyConfiguration) RemoveStatus() *PriorityLevelConfigurationConditionApplyConfiguration {
	b.Status = nil
	return b
}

// GetStatus gets the Status field from the declarative configuration.
func (b *PriorityLevelConfigurationConditionApplyConfiguration) GetStatus() (value v1alpha1.ConditionStatus, ok bool) {
	if v := b.Status; v != nil {
		return *v, true
	}
	return value, false
}

// SetLastTransitionTime sets the LastTransitionTime field in the declarative configuration to the given value.
func (b *PriorityLevelConfigurationConditionApplyConfiguration) SetLastTransitionTime(value v1.Time) *PriorityLevelConfigurationConditionApplyConfiguration {
	b.LastTransitionTime = &value
	return b
}

// RemoveLastTransitionTime removes the LastTransitionTime field from the declarative configuration.
func (b *PriorityLevelConfigurationConditionApplyConfiguration) RemoveLastTransitionTime() *PriorityLevelConfigurationConditionApplyConfiguration {
	b.LastTransitionTime = nil
	return b
}

// GetLastTransitionTime gets the LastTransitionTime field from the declarative configuration.
func (b *PriorityLevelConfigurationConditionApplyConfiguration) GetLastTransitionTime() (value v1.Time, ok bool) {
	if v := b.LastTransitionTime; v != nil {
		return *v, true
	}
	return value, false
}

// SetReason sets the Reason field in the declarative configuration to the given value.
func (b *PriorityLevelConfigurationConditionApplyConfiguration) SetReason(value string) *PriorityLevelConfigurationConditionApplyConfiguration {
	b.Reason = &value
	return b
}

// RemoveReason removes the Reason field from the declarative configuration.
func (b *PriorityLevelConfigurationConditionApplyConfiguration) RemoveReason() *PriorityLevelConfigurationConditionApplyConfiguration {
	b.Reason = nil
	return b
}

// GetReason gets the Reason field from the declarative configuration.
func (b *PriorityLevelConfigurationConditionApplyConfiguration) GetReason() (value string, ok bool) {
	if v := b.Reason; v != nil {
		return *v, true
	}
	return value, false
}

// SetMessage sets the Message field in the declarative configuration to the given value.
func (b *PriorityLevelConfigurationConditionApplyConfiguration) SetMessage(value string) *PriorityLevelConfigurationConditionApplyConfiguration {
	b.Message = &value
	return b
}

// RemoveMessage removes the Message field from the declarative configuration.
func (b *PriorityLevelConfigurationConditionApplyConfiguration) RemoveMessage() *PriorityLevelConfigurationConditionApplyConfiguration {
	b.Message = nil
	return b
}

// GetMessage gets the Message field from the declarative configuration.
func (b *PriorityLevelConfigurationConditionApplyConfiguration) GetMessage() (value string, ok bool) {
	if v := b.Message; v != nil {
		return *v, true
	}
	return value, false
}

// PriorityLevelConfigurationConditionList represents a listAlias of PriorityLevelConfigurationConditionApplyConfiguration.
type PriorityLevelConfigurationConditionList []*PriorityLevelConfigurationConditionApplyConfiguration

// PriorityLevelConfigurationConditionList represents a map of PriorityLevelConfigurationConditionApplyConfiguration.
type PriorityLevelConfigurationConditionMap map[string]PriorityLevelConfigurationConditionApplyConfiguration
