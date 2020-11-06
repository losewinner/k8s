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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/client-go/applyconfigurations/core/v1"
)

// CronJobStatusApplyConfiguration represents an declarative configuration of the CronJobStatus type for use
// with apply.
type CronJobStatusApplyConfiguration struct {
	Active           *v1.ObjectReferenceList `json:"active,omitempty"`
	LastScheduleTime *metav1.Time            `json:"lastScheduleTime,omitempty"`
}

// CronJobStatusApplyConfiguration constructs an declarative configuration of the CronJobStatus type for use with
// apply.
func CronJobStatus() *CronJobStatusApplyConfiguration {
	return &CronJobStatusApplyConfiguration{}
}

// SetActive sets the Active field in the declarative configuration to the given value.
func (b *CronJobStatusApplyConfiguration) SetActive(value v1.ObjectReferenceList) *CronJobStatusApplyConfiguration {
	b.Active = &value
	return b
}

// RemoveActive removes the Active field from the declarative configuration.
func (b *CronJobStatusApplyConfiguration) RemoveActive() *CronJobStatusApplyConfiguration {
	b.Active = nil
	return b
}

// GetActive gets the Active field from the declarative configuration.
func (b *CronJobStatusApplyConfiguration) GetActive() (value v1.ObjectReferenceList, ok bool) {
	if v := b.Active; v != nil {
		return *v, true
	}
	return value, false
}

// SetLastScheduleTime sets the LastScheduleTime field in the declarative configuration to the given value.
func (b *CronJobStatusApplyConfiguration) SetLastScheduleTime(value metav1.Time) *CronJobStatusApplyConfiguration {
	b.LastScheduleTime = &value
	return b
}

// RemoveLastScheduleTime removes the LastScheduleTime field from the declarative configuration.
func (b *CronJobStatusApplyConfiguration) RemoveLastScheduleTime() *CronJobStatusApplyConfiguration {
	b.LastScheduleTime = nil
	return b
}

// GetLastScheduleTime gets the LastScheduleTime field from the declarative configuration.
func (b *CronJobStatusApplyConfiguration) GetLastScheduleTime() (value metav1.Time, ok bool) {
	if v := b.LastScheduleTime; v != nil {
		return *v, true
	}
	return value, false
}

// CronJobStatusList represents a listAlias of CronJobStatusApplyConfiguration.
type CronJobStatusList []*CronJobStatusApplyConfiguration

// CronJobStatusList represents a map of CronJobStatusApplyConfiguration.
type CronJobStatusMap map[string]CronJobStatusApplyConfiguration
