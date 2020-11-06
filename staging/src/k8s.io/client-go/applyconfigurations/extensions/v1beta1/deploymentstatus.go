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

// DeploymentStatusApplyConfiguration represents an declarative configuration of the DeploymentStatus type for use
// with apply.
type DeploymentStatusApplyConfiguration struct {
	ObservedGeneration  *int64                   `json:"observedGeneration,omitempty"`
	Replicas            *int32                   `json:"replicas,omitempty"`
	UpdatedReplicas     *int32                   `json:"updatedReplicas,omitempty"`
	ReadyReplicas       *int32                   `json:"readyReplicas,omitempty"`
	AvailableReplicas   *int32                   `json:"availableReplicas,omitempty"`
	UnavailableReplicas *int32                   `json:"unavailableReplicas,omitempty"`
	Conditions          *DeploymentConditionList `json:"conditions,omitempty"`
	CollisionCount      *int32                   `json:"collisionCount,omitempty"`
}

// DeploymentStatusApplyConfiguration constructs an declarative configuration of the DeploymentStatus type for use with
// apply.
func DeploymentStatus() *DeploymentStatusApplyConfiguration {
	return &DeploymentStatusApplyConfiguration{}
}

// SetObservedGeneration sets the ObservedGeneration field in the declarative configuration to the given value.
func (b *DeploymentStatusApplyConfiguration) SetObservedGeneration(value int64) *DeploymentStatusApplyConfiguration {
	b.ObservedGeneration = &value
	return b
}

// RemoveObservedGeneration removes the ObservedGeneration field from the declarative configuration.
func (b *DeploymentStatusApplyConfiguration) RemoveObservedGeneration() *DeploymentStatusApplyConfiguration {
	b.ObservedGeneration = nil
	return b
}

// GetObservedGeneration gets the ObservedGeneration field from the declarative configuration.
func (b *DeploymentStatusApplyConfiguration) GetObservedGeneration() (value int64, ok bool) {
	if v := b.ObservedGeneration; v != nil {
		return *v, true
	}
	return value, false
}

// SetReplicas sets the Replicas field in the declarative configuration to the given value.
func (b *DeploymentStatusApplyConfiguration) SetReplicas(value int32) *DeploymentStatusApplyConfiguration {
	b.Replicas = &value
	return b
}

// RemoveReplicas removes the Replicas field from the declarative configuration.
func (b *DeploymentStatusApplyConfiguration) RemoveReplicas() *DeploymentStatusApplyConfiguration {
	b.Replicas = nil
	return b
}

// GetReplicas gets the Replicas field from the declarative configuration.
func (b *DeploymentStatusApplyConfiguration) GetReplicas() (value int32, ok bool) {
	if v := b.Replicas; v != nil {
		return *v, true
	}
	return value, false
}

// SetUpdatedReplicas sets the UpdatedReplicas field in the declarative configuration to the given value.
func (b *DeploymentStatusApplyConfiguration) SetUpdatedReplicas(value int32) *DeploymentStatusApplyConfiguration {
	b.UpdatedReplicas = &value
	return b
}

// RemoveUpdatedReplicas removes the UpdatedReplicas field from the declarative configuration.
func (b *DeploymentStatusApplyConfiguration) RemoveUpdatedReplicas() *DeploymentStatusApplyConfiguration {
	b.UpdatedReplicas = nil
	return b
}

// GetUpdatedReplicas gets the UpdatedReplicas field from the declarative configuration.
func (b *DeploymentStatusApplyConfiguration) GetUpdatedReplicas() (value int32, ok bool) {
	if v := b.UpdatedReplicas; v != nil {
		return *v, true
	}
	return value, false
}

// SetReadyReplicas sets the ReadyReplicas field in the declarative configuration to the given value.
func (b *DeploymentStatusApplyConfiguration) SetReadyReplicas(value int32) *DeploymentStatusApplyConfiguration {
	b.ReadyReplicas = &value
	return b
}

// RemoveReadyReplicas removes the ReadyReplicas field from the declarative configuration.
func (b *DeploymentStatusApplyConfiguration) RemoveReadyReplicas() *DeploymentStatusApplyConfiguration {
	b.ReadyReplicas = nil
	return b
}

// GetReadyReplicas gets the ReadyReplicas field from the declarative configuration.
func (b *DeploymentStatusApplyConfiguration) GetReadyReplicas() (value int32, ok bool) {
	if v := b.ReadyReplicas; v != nil {
		return *v, true
	}
	return value, false
}

// SetAvailableReplicas sets the AvailableReplicas field in the declarative configuration to the given value.
func (b *DeploymentStatusApplyConfiguration) SetAvailableReplicas(value int32) *DeploymentStatusApplyConfiguration {
	b.AvailableReplicas = &value
	return b
}

// RemoveAvailableReplicas removes the AvailableReplicas field from the declarative configuration.
func (b *DeploymentStatusApplyConfiguration) RemoveAvailableReplicas() *DeploymentStatusApplyConfiguration {
	b.AvailableReplicas = nil
	return b
}

// GetAvailableReplicas gets the AvailableReplicas field from the declarative configuration.
func (b *DeploymentStatusApplyConfiguration) GetAvailableReplicas() (value int32, ok bool) {
	if v := b.AvailableReplicas; v != nil {
		return *v, true
	}
	return value, false
}

// SetUnavailableReplicas sets the UnavailableReplicas field in the declarative configuration to the given value.
func (b *DeploymentStatusApplyConfiguration) SetUnavailableReplicas(value int32) *DeploymentStatusApplyConfiguration {
	b.UnavailableReplicas = &value
	return b
}

// RemoveUnavailableReplicas removes the UnavailableReplicas field from the declarative configuration.
func (b *DeploymentStatusApplyConfiguration) RemoveUnavailableReplicas() *DeploymentStatusApplyConfiguration {
	b.UnavailableReplicas = nil
	return b
}

// GetUnavailableReplicas gets the UnavailableReplicas field from the declarative configuration.
func (b *DeploymentStatusApplyConfiguration) GetUnavailableReplicas() (value int32, ok bool) {
	if v := b.UnavailableReplicas; v != nil {
		return *v, true
	}
	return value, false
}

// SetConditions sets the Conditions field in the declarative configuration to the given value.
func (b *DeploymentStatusApplyConfiguration) SetConditions(value DeploymentConditionList) *DeploymentStatusApplyConfiguration {
	b.Conditions = &value
	return b
}

// RemoveConditions removes the Conditions field from the declarative configuration.
func (b *DeploymentStatusApplyConfiguration) RemoveConditions() *DeploymentStatusApplyConfiguration {
	b.Conditions = nil
	return b
}

// GetConditions gets the Conditions field from the declarative configuration.
func (b *DeploymentStatusApplyConfiguration) GetConditions() (value DeploymentConditionList, ok bool) {
	if v := b.Conditions; v != nil {
		return *v, true
	}
	return value, false
}

// SetCollisionCount sets the CollisionCount field in the declarative configuration to the given value.
func (b *DeploymentStatusApplyConfiguration) SetCollisionCount(value int32) *DeploymentStatusApplyConfiguration {
	b.CollisionCount = &value
	return b
}

// RemoveCollisionCount removes the CollisionCount field from the declarative configuration.
func (b *DeploymentStatusApplyConfiguration) RemoveCollisionCount() *DeploymentStatusApplyConfiguration {
	b.CollisionCount = nil
	return b
}

// GetCollisionCount gets the CollisionCount field from the declarative configuration.
func (b *DeploymentStatusApplyConfiguration) GetCollisionCount() (value int32, ok bool) {
	if v := b.CollisionCount; v != nil {
		return *v, true
	}
	return value, false
}

// DeploymentStatusList represents a listAlias of DeploymentStatusApplyConfiguration.
type DeploymentStatusList []*DeploymentStatusApplyConfiguration

// DeploymentStatusList represents a map of DeploymentStatusApplyConfiguration.
type DeploymentStatusMap map[string]DeploymentStatusApplyConfiguration
