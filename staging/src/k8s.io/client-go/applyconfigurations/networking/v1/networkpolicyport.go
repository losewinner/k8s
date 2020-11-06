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
	intstr "k8s.io/apimachinery/pkg/util/intstr"
)

// NetworkPolicyPortApplyConfiguration represents an declarative configuration of the NetworkPolicyPort type for use
// with apply.
type NetworkPolicyPortApplyConfiguration struct {
	Protocol *v1.Protocol        `json:"protocol,omitempty"`
	Port     *intstr.IntOrString `json:"port,omitempty"`
}

// NetworkPolicyPortApplyConfiguration constructs an declarative configuration of the NetworkPolicyPort type for use with
// apply.
func NetworkPolicyPort() *NetworkPolicyPortApplyConfiguration {
	return &NetworkPolicyPortApplyConfiguration{}
}

// SetProtocol sets the Protocol field in the declarative configuration to the given value.
func (b *NetworkPolicyPortApplyConfiguration) SetProtocol(value v1.Protocol) *NetworkPolicyPortApplyConfiguration {
	b.Protocol = &value
	return b
}

// RemoveProtocol removes the Protocol field from the declarative configuration.
func (b *NetworkPolicyPortApplyConfiguration) RemoveProtocol() *NetworkPolicyPortApplyConfiguration {
	b.Protocol = nil
	return b
}

// GetProtocol gets the Protocol field from the declarative configuration.
func (b *NetworkPolicyPortApplyConfiguration) GetProtocol() (value v1.Protocol, ok bool) {
	if v := b.Protocol; v != nil {
		return *v, true
	}
	return value, false
}

// SetPort sets the Port field in the declarative configuration to the given value.
func (b *NetworkPolicyPortApplyConfiguration) SetPort(value intstr.IntOrString) *NetworkPolicyPortApplyConfiguration {
	b.Port = &value
	return b
}

// RemovePort removes the Port field from the declarative configuration.
func (b *NetworkPolicyPortApplyConfiguration) RemovePort() *NetworkPolicyPortApplyConfiguration {
	b.Port = nil
	return b
}

// GetPort gets the Port field from the declarative configuration.
func (b *NetworkPolicyPortApplyConfiguration) GetPort() (value intstr.IntOrString, ok bool) {
	if v := b.Port; v != nil {
		return *v, true
	}
	return value, false
}

// NetworkPolicyPortList represents a listAlias of NetworkPolicyPortApplyConfiguration.
type NetworkPolicyPortList []*NetworkPolicyPortApplyConfiguration

// NetworkPolicyPortList represents a map of NetworkPolicyPortApplyConfiguration.
type NetworkPolicyPortMap map[string]NetworkPolicyPortApplyConfiguration
