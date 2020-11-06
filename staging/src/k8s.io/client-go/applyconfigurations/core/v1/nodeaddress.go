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

// NodeAddressApplyConfiguration represents an declarative configuration of the NodeAddress type for use
// with apply.
type NodeAddressApplyConfiguration struct {
	Type    *v1.NodeAddressType `json:"type,omitempty"`
	Address *string             `json:"address,omitempty"`
}

// NodeAddressApplyConfiguration constructs an declarative configuration of the NodeAddress type for use with
// apply.
func NodeAddress() *NodeAddressApplyConfiguration {
	return &NodeAddressApplyConfiguration{}
}

// SetType sets the Type field in the declarative configuration to the given value.
func (b *NodeAddressApplyConfiguration) SetType(value v1.NodeAddressType) *NodeAddressApplyConfiguration {
	b.Type = &value
	return b
}

// RemoveType removes the Type field from the declarative configuration.
func (b *NodeAddressApplyConfiguration) RemoveType() *NodeAddressApplyConfiguration {
	b.Type = nil
	return b
}

// GetType gets the Type field from the declarative configuration.
func (b *NodeAddressApplyConfiguration) GetType() (value v1.NodeAddressType, ok bool) {
	if v := b.Type; v != nil {
		return *v, true
	}
	return value, false
}

// SetAddress sets the Address field in the declarative configuration to the given value.
func (b *NodeAddressApplyConfiguration) SetAddress(value string) *NodeAddressApplyConfiguration {
	b.Address = &value
	return b
}

// RemoveAddress removes the Address field from the declarative configuration.
func (b *NodeAddressApplyConfiguration) RemoveAddress() *NodeAddressApplyConfiguration {
	b.Address = nil
	return b
}

// GetAddress gets the Address field from the declarative configuration.
func (b *NodeAddressApplyConfiguration) GetAddress() (value string, ok bool) {
	if v := b.Address; v != nil {
		return *v, true
	}
	return value, false
}

// NodeAddressList represents a listAlias of NodeAddressApplyConfiguration.
type NodeAddressList []*NodeAddressApplyConfiguration

// NodeAddressList represents a map of NodeAddressApplyConfiguration.
type NodeAddressMap map[string]NodeAddressApplyConfiguration
