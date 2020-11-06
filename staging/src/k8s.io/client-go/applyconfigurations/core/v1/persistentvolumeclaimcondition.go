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

// PersistentVolumeClaimConditionApplyConfiguration represents an declarative configuration of the PersistentVolumeClaimCondition type for use
// with apply.
type PersistentVolumeClaimConditionApplyConfiguration struct {
	Type               *v1.PersistentVolumeClaimConditionType `json:"type,omitempty"`
	Status             *v1.ConditionStatus                    `json:"status,omitempty"`
	LastProbeTime      *metav1.Time                           `json:"lastProbeTime,omitempty"`
	LastTransitionTime *metav1.Time                           `json:"lastTransitionTime,omitempty"`
	Reason             *string                                `json:"reason,omitempty"`
	Message            *string                                `json:"message,omitempty"`
}

// PersistentVolumeClaimConditionApplyConfiguration constructs an declarative configuration of the PersistentVolumeClaimCondition type for use with
// apply.
func PersistentVolumeClaimCondition() *PersistentVolumeClaimConditionApplyConfiguration {
	return &PersistentVolumeClaimConditionApplyConfiguration{}
}

// SetType sets the Type field in the declarative configuration to the given value.
func (b *PersistentVolumeClaimConditionApplyConfiguration) SetType(value v1.PersistentVolumeClaimConditionType) *PersistentVolumeClaimConditionApplyConfiguration {
	b.Type = &value
	return b
}

// RemoveType removes the Type field from the declarative configuration.
func (b *PersistentVolumeClaimConditionApplyConfiguration) RemoveType() *PersistentVolumeClaimConditionApplyConfiguration {
	b.Type = nil
	return b
}

// GetType gets the Type field from the declarative configuration.
func (b *PersistentVolumeClaimConditionApplyConfiguration) GetType() (value v1.PersistentVolumeClaimConditionType, ok bool) {
	if v := b.Type; v != nil {
		return *v, true
	}
	return value, false
}

// SetStatus sets the Status field in the declarative configuration to the given value.
func (b *PersistentVolumeClaimConditionApplyConfiguration) SetStatus(value v1.ConditionStatus) *PersistentVolumeClaimConditionApplyConfiguration {
	b.Status = &value
	return b
}

// RemoveStatus removes the Status field from the declarative configuration.
func (b *PersistentVolumeClaimConditionApplyConfiguration) RemoveStatus() *PersistentVolumeClaimConditionApplyConfiguration {
	b.Status = nil
	return b
}

// GetStatus gets the Status field from the declarative configuration.
func (b *PersistentVolumeClaimConditionApplyConfiguration) GetStatus() (value v1.ConditionStatus, ok bool) {
	if v := b.Status; v != nil {
		return *v, true
	}
	return value, false
}

// SetLastProbeTime sets the LastProbeTime field in the declarative configuration to the given value.
func (b *PersistentVolumeClaimConditionApplyConfiguration) SetLastProbeTime(value metav1.Time) *PersistentVolumeClaimConditionApplyConfiguration {
	b.LastProbeTime = &value
	return b
}

// RemoveLastProbeTime removes the LastProbeTime field from the declarative configuration.
func (b *PersistentVolumeClaimConditionApplyConfiguration) RemoveLastProbeTime() *PersistentVolumeClaimConditionApplyConfiguration {
	b.LastProbeTime = nil
	return b
}

// GetLastProbeTime gets the LastProbeTime field from the declarative configuration.
func (b *PersistentVolumeClaimConditionApplyConfiguration) GetLastProbeTime() (value metav1.Time, ok bool) {
	if v := b.LastProbeTime; v != nil {
		return *v, true
	}
	return value, false
}

// SetLastTransitionTime sets the LastTransitionTime field in the declarative configuration to the given value.
func (b *PersistentVolumeClaimConditionApplyConfiguration) SetLastTransitionTime(value metav1.Time) *PersistentVolumeClaimConditionApplyConfiguration {
	b.LastTransitionTime = &value
	return b
}

// RemoveLastTransitionTime removes the LastTransitionTime field from the declarative configuration.
func (b *PersistentVolumeClaimConditionApplyConfiguration) RemoveLastTransitionTime() *PersistentVolumeClaimConditionApplyConfiguration {
	b.LastTransitionTime = nil
	return b
}

// GetLastTransitionTime gets the LastTransitionTime field from the declarative configuration.
func (b *PersistentVolumeClaimConditionApplyConfiguration) GetLastTransitionTime() (value metav1.Time, ok bool) {
	if v := b.LastTransitionTime; v != nil {
		return *v, true
	}
	return value, false
}

// SetReason sets the Reason field in the declarative configuration to the given value.
func (b *PersistentVolumeClaimConditionApplyConfiguration) SetReason(value string) *PersistentVolumeClaimConditionApplyConfiguration {
	b.Reason = &value
	return b
}

// RemoveReason removes the Reason field from the declarative configuration.
func (b *PersistentVolumeClaimConditionApplyConfiguration) RemoveReason() *PersistentVolumeClaimConditionApplyConfiguration {
	b.Reason = nil
	return b
}

// GetReason gets the Reason field from the declarative configuration.
func (b *PersistentVolumeClaimConditionApplyConfiguration) GetReason() (value string, ok bool) {
	if v := b.Reason; v != nil {
		return *v, true
	}
	return value, false
}

// SetMessage sets the Message field in the declarative configuration to the given value.
func (b *PersistentVolumeClaimConditionApplyConfiguration) SetMessage(value string) *PersistentVolumeClaimConditionApplyConfiguration {
	b.Message = &value
	return b
}

// RemoveMessage removes the Message field from the declarative configuration.
func (b *PersistentVolumeClaimConditionApplyConfiguration) RemoveMessage() *PersistentVolumeClaimConditionApplyConfiguration {
	b.Message = nil
	return b
}

// GetMessage gets the Message field from the declarative configuration.
func (b *PersistentVolumeClaimConditionApplyConfiguration) GetMessage() (value string, ok bool) {
	if v := b.Message; v != nil {
		return *v, true
	}
	return value, false
}

// PersistentVolumeClaimConditionList represents a listAlias of PersistentVolumeClaimConditionApplyConfiguration.
type PersistentVolumeClaimConditionList []*PersistentVolumeClaimConditionApplyConfiguration

// PersistentVolumeClaimConditionList represents a map of PersistentVolumeClaimConditionApplyConfiguration.
type PersistentVolumeClaimConditionMap map[string]PersistentVolumeClaimConditionApplyConfiguration
