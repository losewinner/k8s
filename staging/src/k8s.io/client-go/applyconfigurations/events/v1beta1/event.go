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
	corev1 "k8s.io/client-go/applyconfigurations/core/v1"
	v1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

// EventApplyConfiguration represents an declarative configuration of the Event type for use
// with apply.
type EventApplyConfiguration struct {
	v1.TypeMetaApplyConfiguration `json:",inline"`
	ObjectMeta                    *v1.ObjectMetaApplyConfiguration          `json:"metadata,omitempty"`
	EventTime                     *metav1.MicroTime                         `json:"eventTime,omitempty"`
	Series                        *EventSeriesApplyConfiguration            `json:"series,omitempty"`
	ReportingController           *string                                   `json:"reportingController,omitempty"`
	ReportingInstance             *string                                   `json:"reportingInstance,omitempty"`
	Action                        *string                                   `json:"action,omitempty"`
	Reason                        *string                                   `json:"reason,omitempty"`
	Regarding                     *corev1.ObjectReferenceApplyConfiguration `json:"regarding,omitempty"`
	Related                       *corev1.ObjectReferenceApplyConfiguration `json:"related,omitempty"`
	Note                          *string                                   `json:"note,omitempty"`
	Type                          *string                                   `json:"type,omitempty"`
	DeprecatedSource              *corev1.EventSourceApplyConfiguration     `json:"deprecatedSource,omitempty"`
	DeprecatedFirstTimestamp      *metav1.Time                              `json:"deprecatedFirstTimestamp,omitempty"`
	DeprecatedLastTimestamp       *metav1.Time                              `json:"deprecatedLastTimestamp,omitempty"`
	DeprecatedCount               *int32                                    `json:"deprecatedCount,omitempty"`
}

// EventApplyConfiguration constructs an declarative configuration of the Event type for use with
// apply.
func Event() *EventApplyConfiguration {
	return &EventApplyConfiguration{}
}

// SetTypeMeta sets the TypeMeta field in the declarative configuration to the given value.
func (b *EventApplyConfiguration) SetTypeMeta(value *v1.TypeMetaApplyConfiguration) *EventApplyConfiguration {
	if value != nil {
		b.TypeMetaApplyConfiguration = *value
	}
	return b
}

// GetTypeMeta gets the TypeMeta field from the declarative configuration.
func (b *EventApplyConfiguration) GetTypeMeta() (value *v1.TypeMetaApplyConfiguration, ok bool) {
	return &b.TypeMetaApplyConfiguration, true
}

// SetObjectMeta sets the ObjectMeta field in the declarative configuration to the given value.
func (b *EventApplyConfiguration) SetObjectMeta(value *v1.ObjectMetaApplyConfiguration) *EventApplyConfiguration {
	b.ObjectMeta = value
	return b
}

// RemoveObjectMeta removes the ObjectMeta field from the declarative configuration.
func (b *EventApplyConfiguration) RemoveObjectMeta() *EventApplyConfiguration {
	b.ObjectMeta = nil
	return b
}

// GetObjectMeta gets the ObjectMeta field from the declarative configuration.
func (b *EventApplyConfiguration) GetObjectMeta() (value *v1.ObjectMetaApplyConfiguration, ok bool) {
	return b.ObjectMeta, b.ObjectMeta != nil
}

// SetEventTime sets the EventTime field in the declarative configuration to the given value.
func (b *EventApplyConfiguration) SetEventTime(value metav1.MicroTime) *EventApplyConfiguration {
	b.EventTime = &value
	return b
}

// RemoveEventTime removes the EventTime field from the declarative configuration.
func (b *EventApplyConfiguration) RemoveEventTime() *EventApplyConfiguration {
	b.EventTime = nil
	return b
}

// GetEventTime gets the EventTime field from the declarative configuration.
func (b *EventApplyConfiguration) GetEventTime() (value metav1.MicroTime, ok bool) {
	if v := b.EventTime; v != nil {
		return *v, true
	}
	return value, false
}

// SetSeries sets the Series field in the declarative configuration to the given value.
func (b *EventApplyConfiguration) SetSeries(value *EventSeriesApplyConfiguration) *EventApplyConfiguration {
	b.Series = value
	return b
}

// RemoveSeries removes the Series field from the declarative configuration.
func (b *EventApplyConfiguration) RemoveSeries() *EventApplyConfiguration {
	b.Series = nil
	return b
}

// GetSeries gets the Series field from the declarative configuration.
func (b *EventApplyConfiguration) GetSeries() (value *EventSeriesApplyConfiguration, ok bool) {
	return b.Series, b.Series != nil
}

// SetReportingController sets the ReportingController field in the declarative configuration to the given value.
func (b *EventApplyConfiguration) SetReportingController(value string) *EventApplyConfiguration {
	b.ReportingController = &value
	return b
}

// RemoveReportingController removes the ReportingController field from the declarative configuration.
func (b *EventApplyConfiguration) RemoveReportingController() *EventApplyConfiguration {
	b.ReportingController = nil
	return b
}

// GetReportingController gets the ReportingController field from the declarative configuration.
func (b *EventApplyConfiguration) GetReportingController() (value string, ok bool) {
	if v := b.ReportingController; v != nil {
		return *v, true
	}
	return value, false
}

// SetReportingInstance sets the ReportingInstance field in the declarative configuration to the given value.
func (b *EventApplyConfiguration) SetReportingInstance(value string) *EventApplyConfiguration {
	b.ReportingInstance = &value
	return b
}

// RemoveReportingInstance removes the ReportingInstance field from the declarative configuration.
func (b *EventApplyConfiguration) RemoveReportingInstance() *EventApplyConfiguration {
	b.ReportingInstance = nil
	return b
}

// GetReportingInstance gets the ReportingInstance field from the declarative configuration.
func (b *EventApplyConfiguration) GetReportingInstance() (value string, ok bool) {
	if v := b.ReportingInstance; v != nil {
		return *v, true
	}
	return value, false
}

// SetAction sets the Action field in the declarative configuration to the given value.
func (b *EventApplyConfiguration) SetAction(value string) *EventApplyConfiguration {
	b.Action = &value
	return b
}

// RemoveAction removes the Action field from the declarative configuration.
func (b *EventApplyConfiguration) RemoveAction() *EventApplyConfiguration {
	b.Action = nil
	return b
}

// GetAction gets the Action field from the declarative configuration.
func (b *EventApplyConfiguration) GetAction() (value string, ok bool) {
	if v := b.Action; v != nil {
		return *v, true
	}
	return value, false
}

// SetReason sets the Reason field in the declarative configuration to the given value.
func (b *EventApplyConfiguration) SetReason(value string) *EventApplyConfiguration {
	b.Reason = &value
	return b
}

// RemoveReason removes the Reason field from the declarative configuration.
func (b *EventApplyConfiguration) RemoveReason() *EventApplyConfiguration {
	b.Reason = nil
	return b
}

// GetReason gets the Reason field from the declarative configuration.
func (b *EventApplyConfiguration) GetReason() (value string, ok bool) {
	if v := b.Reason; v != nil {
		return *v, true
	}
	return value, false
}

// SetRegarding sets the Regarding field in the declarative configuration to the given value.
func (b *EventApplyConfiguration) SetRegarding(value *corev1.ObjectReferenceApplyConfiguration) *EventApplyConfiguration {
	b.Regarding = value
	return b
}

// RemoveRegarding removes the Regarding field from the declarative configuration.
func (b *EventApplyConfiguration) RemoveRegarding() *EventApplyConfiguration {
	b.Regarding = nil
	return b
}

// GetRegarding gets the Regarding field from the declarative configuration.
func (b *EventApplyConfiguration) GetRegarding() (value *corev1.ObjectReferenceApplyConfiguration, ok bool) {
	return b.Regarding, b.Regarding != nil
}

// SetRelated sets the Related field in the declarative configuration to the given value.
func (b *EventApplyConfiguration) SetRelated(value *corev1.ObjectReferenceApplyConfiguration) *EventApplyConfiguration {
	b.Related = value
	return b
}

// RemoveRelated removes the Related field from the declarative configuration.
func (b *EventApplyConfiguration) RemoveRelated() *EventApplyConfiguration {
	b.Related = nil
	return b
}

// GetRelated gets the Related field from the declarative configuration.
func (b *EventApplyConfiguration) GetRelated() (value *corev1.ObjectReferenceApplyConfiguration, ok bool) {
	return b.Related, b.Related != nil
}

// SetNote sets the Note field in the declarative configuration to the given value.
func (b *EventApplyConfiguration) SetNote(value string) *EventApplyConfiguration {
	b.Note = &value
	return b
}

// RemoveNote removes the Note field from the declarative configuration.
func (b *EventApplyConfiguration) RemoveNote() *EventApplyConfiguration {
	b.Note = nil
	return b
}

// GetNote gets the Note field from the declarative configuration.
func (b *EventApplyConfiguration) GetNote() (value string, ok bool) {
	if v := b.Note; v != nil {
		return *v, true
	}
	return value, false
}

// SetType sets the Type field in the declarative configuration to the given value.
func (b *EventApplyConfiguration) SetType(value string) *EventApplyConfiguration {
	b.Type = &value
	return b
}

// RemoveType removes the Type field from the declarative configuration.
func (b *EventApplyConfiguration) RemoveType() *EventApplyConfiguration {
	b.Type = nil
	return b
}

// GetType gets the Type field from the declarative configuration.
func (b *EventApplyConfiguration) GetType() (value string, ok bool) {
	if v := b.Type; v != nil {
		return *v, true
	}
	return value, false
}

// SetDeprecatedSource sets the DeprecatedSource field in the declarative configuration to the given value.
func (b *EventApplyConfiguration) SetDeprecatedSource(value *corev1.EventSourceApplyConfiguration) *EventApplyConfiguration {
	b.DeprecatedSource = value
	return b
}

// RemoveDeprecatedSource removes the DeprecatedSource field from the declarative configuration.
func (b *EventApplyConfiguration) RemoveDeprecatedSource() *EventApplyConfiguration {
	b.DeprecatedSource = nil
	return b
}

// GetDeprecatedSource gets the DeprecatedSource field from the declarative configuration.
func (b *EventApplyConfiguration) GetDeprecatedSource() (value *corev1.EventSourceApplyConfiguration, ok bool) {
	return b.DeprecatedSource, b.DeprecatedSource != nil
}

// SetDeprecatedFirstTimestamp sets the DeprecatedFirstTimestamp field in the declarative configuration to the given value.
func (b *EventApplyConfiguration) SetDeprecatedFirstTimestamp(value metav1.Time) *EventApplyConfiguration {
	b.DeprecatedFirstTimestamp = &value
	return b
}

// RemoveDeprecatedFirstTimestamp removes the DeprecatedFirstTimestamp field from the declarative configuration.
func (b *EventApplyConfiguration) RemoveDeprecatedFirstTimestamp() *EventApplyConfiguration {
	b.DeprecatedFirstTimestamp = nil
	return b
}

// GetDeprecatedFirstTimestamp gets the DeprecatedFirstTimestamp field from the declarative configuration.
func (b *EventApplyConfiguration) GetDeprecatedFirstTimestamp() (value metav1.Time, ok bool) {
	if v := b.DeprecatedFirstTimestamp; v != nil {
		return *v, true
	}
	return value, false
}

// SetDeprecatedLastTimestamp sets the DeprecatedLastTimestamp field in the declarative configuration to the given value.
func (b *EventApplyConfiguration) SetDeprecatedLastTimestamp(value metav1.Time) *EventApplyConfiguration {
	b.DeprecatedLastTimestamp = &value
	return b
}

// RemoveDeprecatedLastTimestamp removes the DeprecatedLastTimestamp field from the declarative configuration.
func (b *EventApplyConfiguration) RemoveDeprecatedLastTimestamp() *EventApplyConfiguration {
	b.DeprecatedLastTimestamp = nil
	return b
}

// GetDeprecatedLastTimestamp gets the DeprecatedLastTimestamp field from the declarative configuration.
func (b *EventApplyConfiguration) GetDeprecatedLastTimestamp() (value metav1.Time, ok bool) {
	if v := b.DeprecatedLastTimestamp; v != nil {
		return *v, true
	}
	return value, false
}

// SetDeprecatedCount sets the DeprecatedCount field in the declarative configuration to the given value.
func (b *EventApplyConfiguration) SetDeprecatedCount(value int32) *EventApplyConfiguration {
	b.DeprecatedCount = &value
	return b
}

// RemoveDeprecatedCount removes the DeprecatedCount field from the declarative configuration.
func (b *EventApplyConfiguration) RemoveDeprecatedCount() *EventApplyConfiguration {
	b.DeprecatedCount = nil
	return b
}

// GetDeprecatedCount gets the DeprecatedCount field from the declarative configuration.
func (b *EventApplyConfiguration) GetDeprecatedCount() (value int32, ok bool) {
	if v := b.DeprecatedCount; v != nil {
		return *v, true
	}
	return value, false
}

// EventList represents a listAlias of EventApplyConfiguration.
type EventList []*EventApplyConfiguration

// EventList represents a map of EventApplyConfiguration.
type EventMap map[string]EventApplyConfiguration
