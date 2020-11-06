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
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EventSeriesApplyConfiguration represents an declarative configuration of the EventSeries type for use
// with apply.
type EventSeriesApplyConfiguration struct {
	Count            *int32        `json:"count,omitempty"`
	LastObservedTime *v1.MicroTime `json:"lastObservedTime,omitempty"`
}

// EventSeriesApplyConfiguration constructs an declarative configuration of the EventSeries type for use with
// apply.
func EventSeries() *EventSeriesApplyConfiguration {
	return &EventSeriesApplyConfiguration{}
}

// SetCount sets the Count field in the declarative configuration to the given value.
func (b *EventSeriesApplyConfiguration) SetCount(value int32) *EventSeriesApplyConfiguration {
	b.Count = &value
	return b
}

// RemoveCount removes the Count field from the declarative configuration.
func (b *EventSeriesApplyConfiguration) RemoveCount() *EventSeriesApplyConfiguration {
	b.Count = nil
	return b
}

// GetCount gets the Count field from the declarative configuration.
func (b *EventSeriesApplyConfiguration) GetCount() (value int32, ok bool) {
	if v := b.Count; v != nil {
		return *v, true
	}
	return value, false
}

// SetLastObservedTime sets the LastObservedTime field in the declarative configuration to the given value.
func (b *EventSeriesApplyConfiguration) SetLastObservedTime(value v1.MicroTime) *EventSeriesApplyConfiguration {
	b.LastObservedTime = &value
	return b
}

// RemoveLastObservedTime removes the LastObservedTime field from the declarative configuration.
func (b *EventSeriesApplyConfiguration) RemoveLastObservedTime() *EventSeriesApplyConfiguration {
	b.LastObservedTime = nil
	return b
}

// GetLastObservedTime gets the LastObservedTime field from the declarative configuration.
func (b *EventSeriesApplyConfiguration) GetLastObservedTime() (value v1.MicroTime, ok bool) {
	if v := b.LastObservedTime; v != nil {
		return *v, true
	}
	return value, false
}

// EventSeriesList represents a listAlias of EventSeriesApplyConfiguration.
type EventSeriesList []*EventSeriesApplyConfiguration

// EventSeriesList represents a map of EventSeriesApplyConfiguration.
type EventSeriesMap map[string]EventSeriesApplyConfiguration
