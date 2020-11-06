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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PodConditionApplyConfiguration represents an declarative configuration of the PodCondition type for use
// with apply.
type PodConditionApplyConfiguration struct {
	Type               *v1.PodConditionType `json:"type,omitempty"`
	Status             *v1.ConditionStatus  `json:"status,omitempty"`
	LastProbeTime      *metav1.Time         `json:"lastProbeTime,omitempty"`
	LastTransitionTime *metav1.Time         `json:"lastTransitionTime,omitempty"`
	Reason             *string              `json:"reason,omitempty"`
	Message            *string              `json:"message,omitempty"`
}

// PodConditionApplyConfiguration constructs an declarative configuration of the PodCondition type for use with
// apply.
func PodCondition() *PodConditionApplyConfiguration {
	return &PodConditionApplyConfiguration{}
}

// SetType sets the Type field in the declarative configuration to the given value.
func (b *PodConditionApplyConfiguration) SetType(value v1.PodConditionType) *PodConditionApplyConfiguration {
	b.Type = &value
	return b
}

// RemoveType removes the Type field from the declarative configuration.
func (b *PodConditionApplyConfiguration) RemoveType() *PodConditionApplyConfiguration {
	b.Type = nil
	return b
}

// GetType gets the Type field from the declarative configuration.
func (b *PodConditionApplyConfiguration) GetType() (value v1.PodConditionType, ok bool) {
	if v := b.Type; v != nil {
		return *v, true
	}
	return value, false
}

// SetStatus sets the Status field in the declarative configuration to the given value.
func (b *PodConditionApplyConfiguration) SetStatus(value v1.ConditionStatus) *PodConditionApplyConfiguration {
	b.Status = &value
	return b
}

// RemoveStatus removes the Status field from the declarative configuration.
func (b *PodConditionApplyConfiguration) RemoveStatus() *PodConditionApplyConfiguration {
	b.Status = nil
	return b
}

// GetStatus gets the Status field from the declarative configuration.
func (b *PodConditionApplyConfiguration) GetStatus() (value v1.ConditionStatus, ok bool) {
	if v := b.Status; v != nil {
		return *v, true
	}
	return value, false
}

// SetLastProbeTime sets the LastProbeTime field in the declarative configuration to the given value.
func (b *PodConditionApplyConfiguration) SetLastProbeTime(value metav1.Time) *PodConditionApplyConfiguration {
	b.LastProbeTime = &value
	return b
}

// RemoveLastProbeTime removes the LastProbeTime field from the declarative configuration.
func (b *PodConditionApplyConfiguration) RemoveLastProbeTime() *PodConditionApplyConfiguration {
	b.LastProbeTime = nil
	return b
}

// GetLastProbeTime gets the LastProbeTime field from the declarative configuration.
func (b *PodConditionApplyConfiguration) GetLastProbeTime() (value metav1.Time, ok bool) {
	if v := b.LastProbeTime; v != nil {
		return *v, true
	}
	return value, false
}

// SetLastTransitionTime sets the LastTransitionTime field in the declarative configuration to the given value.
func (b *PodConditionApplyConfiguration) SetLastTransitionTime(value metav1.Time) *PodConditionApplyConfiguration {
	b.LastTransitionTime = &value
	return b
}

// RemoveLastTransitionTime removes the LastTransitionTime field from the declarative configuration.
func (b *PodConditionApplyConfiguration) RemoveLastTransitionTime() *PodConditionApplyConfiguration {
	b.LastTransitionTime = nil
	return b
}

// GetLastTransitionTime gets the LastTransitionTime field from the declarative configuration.
func (b *PodConditionApplyConfiguration) GetLastTransitionTime() (value metav1.Time, ok bool) {
	if v := b.LastTransitionTime; v != nil {
		return *v, true
	}
	return value, false
}

// SetReason sets the Reason field in the declarative configuration to the given value.
func (b *PodConditionApplyConfiguration) SetReason(value string) *PodConditionApplyConfiguration {
	b.Reason = &value
	return b
}

// RemoveReason removes the Reason field from the declarative configuration.
func (b *PodConditionApplyConfiguration) RemoveReason() *PodConditionApplyConfiguration {
	b.Reason = nil
	return b
}

// GetReason gets the Reason field from the declarative configuration.
func (b *PodConditionApplyConfiguration) GetReason() (value string, ok bool) {
	if v := b.Reason; v != nil {
		return *v, true
	}
	return value, false
}

// SetMessage sets the Message field in the declarative configuration to the given value.
func (b *PodConditionApplyConfiguration) SetMessage(value string) *PodConditionApplyConfiguration {
	b.Message = &value
	return b
}

// RemoveMessage removes the Message field from the declarative configuration.
func (b *PodConditionApplyConfiguration) RemoveMessage() *PodConditionApplyConfiguration {
	b.Message = nil
	return b
}

// GetMessage gets the Message field from the declarative configuration.
func (b *PodConditionApplyConfiguration) GetMessage() (value string, ok bool) {
	if v := b.Message; v != nil {
		return *v, true
	}
	return value, false
}

// PodConditionList represents a listAlias of PodConditionApplyConfiguration.
type PodConditionList []*PodConditionApplyConfiguration

// PodConditionList represents a map of PodConditionApplyConfiguration.
type PodConditionMap map[string]PodConditionApplyConfiguration
