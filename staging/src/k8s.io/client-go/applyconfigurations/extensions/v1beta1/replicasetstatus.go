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

// ReplicaSetStatusApplyConfiguration represents an declarative configuration of the ReplicaSetStatus type for use
// with apply.
type ReplicaSetStatusApplyConfiguration struct {
	Replicas             *int32                   `json:"replicas,omitempty"`
	FullyLabeledReplicas *int32                   `json:"fullyLabeledReplicas,omitempty"`
	ReadyReplicas        *int32                   `json:"readyReplicas,omitempty"`
	AvailableReplicas    *int32                   `json:"availableReplicas,omitempty"`
	ObservedGeneration   *int64                   `json:"observedGeneration,omitempty"`
	Conditions           *ReplicaSetConditionList `json:"conditions,omitempty"`
}

// ReplicaSetStatusApplyConfiguration constructs an declarative configuration of the ReplicaSetStatus type for use with
// apply.
func ReplicaSetStatus() *ReplicaSetStatusApplyConfiguration {
	return &ReplicaSetStatusApplyConfiguration{}
}

// SetReplicas sets the Replicas field in the declarative configuration to the given value.
func (b *ReplicaSetStatusApplyConfiguration) SetReplicas(value int32) *ReplicaSetStatusApplyConfiguration {
	b.Replicas = &value
	return b
}

// RemoveReplicas removes the Replicas field from the declarative configuration.
func (b *ReplicaSetStatusApplyConfiguration) RemoveReplicas() *ReplicaSetStatusApplyConfiguration {
	b.Replicas = nil
	return b
}

// GetReplicas gets the Replicas field from the declarative configuration.
func (b *ReplicaSetStatusApplyConfiguration) GetReplicas() (value int32, ok bool) {
	if v := b.Replicas; v != nil {
		return *v, true
	}
	return value, false
}

// SetFullyLabeledReplicas sets the FullyLabeledReplicas field in the declarative configuration to the given value.
func (b *ReplicaSetStatusApplyConfiguration) SetFullyLabeledReplicas(value int32) *ReplicaSetStatusApplyConfiguration {
	b.FullyLabeledReplicas = &value
	return b
}

// RemoveFullyLabeledReplicas removes the FullyLabeledReplicas field from the declarative configuration.
func (b *ReplicaSetStatusApplyConfiguration) RemoveFullyLabeledReplicas() *ReplicaSetStatusApplyConfiguration {
	b.FullyLabeledReplicas = nil
	return b
}

// GetFullyLabeledReplicas gets the FullyLabeledReplicas field from the declarative configuration.
func (b *ReplicaSetStatusApplyConfiguration) GetFullyLabeledReplicas() (value int32, ok bool) {
	if v := b.FullyLabeledReplicas; v != nil {
		return *v, true
	}
	return value, false
}

// SetReadyReplicas sets the ReadyReplicas field in the declarative configuration to the given value.
func (b *ReplicaSetStatusApplyConfiguration) SetReadyReplicas(value int32) *ReplicaSetStatusApplyConfiguration {
	b.ReadyReplicas = &value
	return b
}

// RemoveReadyReplicas removes the ReadyReplicas field from the declarative configuration.
func (b *ReplicaSetStatusApplyConfiguration) RemoveReadyReplicas() *ReplicaSetStatusApplyConfiguration {
	b.ReadyReplicas = nil
	return b
}

// GetReadyReplicas gets the ReadyReplicas field from the declarative configuration.
func (b *ReplicaSetStatusApplyConfiguration) GetReadyReplicas() (value int32, ok bool) {
	if v := b.ReadyReplicas; v != nil {
		return *v, true
	}
	return value, false
}

// SetAvailableReplicas sets the AvailableReplicas field in the declarative configuration to the given value.
func (b *ReplicaSetStatusApplyConfiguration) SetAvailableReplicas(value int32) *ReplicaSetStatusApplyConfiguration {
	b.AvailableReplicas = &value
	return b
}

// RemoveAvailableReplicas removes the AvailableReplicas field from the declarative configuration.
func (b *ReplicaSetStatusApplyConfiguration) RemoveAvailableReplicas() *ReplicaSetStatusApplyConfiguration {
	b.AvailableReplicas = nil
	return b
}

// GetAvailableReplicas gets the AvailableReplicas field from the declarative configuration.
func (b *ReplicaSetStatusApplyConfiguration) GetAvailableReplicas() (value int32, ok bool) {
	if v := b.AvailableReplicas; v != nil {
		return *v, true
	}
	return value, false
}

// SetObservedGeneration sets the ObservedGeneration field in the declarative configuration to the given value.
func (b *ReplicaSetStatusApplyConfiguration) SetObservedGeneration(value int64) *ReplicaSetStatusApplyConfiguration {
	b.ObservedGeneration = &value
	return b
}

// RemoveObservedGeneration removes the ObservedGeneration field from the declarative configuration.
func (b *ReplicaSetStatusApplyConfiguration) RemoveObservedGeneration() *ReplicaSetStatusApplyConfiguration {
	b.ObservedGeneration = nil
	return b
}

// GetObservedGeneration gets the ObservedGeneration field from the declarative configuration.
func (b *ReplicaSetStatusApplyConfiguration) GetObservedGeneration() (value int64, ok bool) {
	if v := b.ObservedGeneration; v != nil {
		return *v, true
	}
	return value, false
}

// SetConditions sets the Conditions field in the declarative configuration to the given value.
func (b *ReplicaSetStatusApplyConfiguration) SetConditions(value ReplicaSetConditionList) *ReplicaSetStatusApplyConfiguration {
	b.Conditions = &value
	return b
}

// RemoveConditions removes the Conditions field from the declarative configuration.
func (b *ReplicaSetStatusApplyConfiguration) RemoveConditions() *ReplicaSetStatusApplyConfiguration {
	b.Conditions = nil
	return b
}

// GetConditions gets the Conditions field from the declarative configuration.
func (b *ReplicaSetStatusApplyConfiguration) GetConditions() (value ReplicaSetConditionList, ok bool) {
	if v := b.Conditions; v != nil {
		return *v, true
	}
	return value, false
}

// ReplicaSetStatusList represents a listAlias of ReplicaSetStatusApplyConfiguration.
type ReplicaSetStatusList []*ReplicaSetStatusApplyConfiguration

// ReplicaSetStatusList represents a map of ReplicaSetStatusApplyConfiguration.
type ReplicaSetStatusMap map[string]ReplicaSetStatusApplyConfiguration
