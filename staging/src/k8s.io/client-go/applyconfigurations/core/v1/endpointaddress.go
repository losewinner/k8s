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

// EndpointAddressApplyConfiguration represents an declarative configuration of the EndpointAddress type for use
// with apply.
type EndpointAddressApplyConfiguration struct {
	IP        *string                            `json:"ip,omitempty"`
	Hostname  *string                            `json:"hostname,omitempty"`
	NodeName  *string                            `json:"nodeName,omitempty"`
	TargetRef *ObjectReferenceApplyConfiguration `json:"targetRef,omitempty"`
}

// EndpointAddressApplyConfiguration constructs an declarative configuration of the EndpointAddress type for use with
// apply.
func EndpointAddress() *EndpointAddressApplyConfiguration {
	return &EndpointAddressApplyConfiguration{}
}

// SetIP sets the IP field in the declarative configuration to the given value.
func (b *EndpointAddressApplyConfiguration) SetIP(value string) *EndpointAddressApplyConfiguration {
	b.IP = &value
	return b
}

// RemoveIP removes the IP field from the declarative configuration.
func (b *EndpointAddressApplyConfiguration) RemoveIP() *EndpointAddressApplyConfiguration {
	b.IP = nil
	return b
}

// GetIP gets the IP field from the declarative configuration.
func (b *EndpointAddressApplyConfiguration) GetIP() (value string, ok bool) {
	if v := b.IP; v != nil {
		return *v, true
	}
	return value, false
}

// SetHostname sets the Hostname field in the declarative configuration to the given value.
func (b *EndpointAddressApplyConfiguration) SetHostname(value string) *EndpointAddressApplyConfiguration {
	b.Hostname = &value
	return b
}

// RemoveHostname removes the Hostname field from the declarative configuration.
func (b *EndpointAddressApplyConfiguration) RemoveHostname() *EndpointAddressApplyConfiguration {
	b.Hostname = nil
	return b
}

// GetHostname gets the Hostname field from the declarative configuration.
func (b *EndpointAddressApplyConfiguration) GetHostname() (value string, ok bool) {
	if v := b.Hostname; v != nil {
		return *v, true
	}
	return value, false
}

// SetNodeName sets the NodeName field in the declarative configuration to the given value.
func (b *EndpointAddressApplyConfiguration) SetNodeName(value string) *EndpointAddressApplyConfiguration {
	b.NodeName = &value
	return b
}

// RemoveNodeName removes the NodeName field from the declarative configuration.
func (b *EndpointAddressApplyConfiguration) RemoveNodeName() *EndpointAddressApplyConfiguration {
	b.NodeName = nil
	return b
}

// GetNodeName gets the NodeName field from the declarative configuration.
func (b *EndpointAddressApplyConfiguration) GetNodeName() (value string, ok bool) {
	if v := b.NodeName; v != nil {
		return *v, true
	}
	return value, false
}

// SetTargetRef sets the TargetRef field in the declarative configuration to the given value.
func (b *EndpointAddressApplyConfiguration) SetTargetRef(value *ObjectReferenceApplyConfiguration) *EndpointAddressApplyConfiguration {
	b.TargetRef = value
	return b
}

// RemoveTargetRef removes the TargetRef field from the declarative configuration.
func (b *EndpointAddressApplyConfiguration) RemoveTargetRef() *EndpointAddressApplyConfiguration {
	b.TargetRef = nil
	return b
}

// GetTargetRef gets the TargetRef field from the declarative configuration.
func (b *EndpointAddressApplyConfiguration) GetTargetRef() (value *ObjectReferenceApplyConfiguration, ok bool) {
	return b.TargetRef, b.TargetRef != nil
}

// EndpointAddressList represents a listAlias of EndpointAddressApplyConfiguration.
type EndpointAddressList []*EndpointAddressApplyConfiguration

// EndpointAddressList represents a map of EndpointAddressApplyConfiguration.
type EndpointAddressMap map[string]EndpointAddressApplyConfiguration
