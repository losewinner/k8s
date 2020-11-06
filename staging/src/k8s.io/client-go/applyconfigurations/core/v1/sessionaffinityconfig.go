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

// SessionAffinityConfigApplyConfiguration represents an declarative configuration of the SessionAffinityConfig type for use
// with apply.
type SessionAffinityConfigApplyConfiguration struct {
	ClientIP *ClientIPConfigApplyConfiguration `json:"clientIP,omitempty"`
}

// SessionAffinityConfigApplyConfiguration constructs an declarative configuration of the SessionAffinityConfig type for use with
// apply.
func SessionAffinityConfig() *SessionAffinityConfigApplyConfiguration {
	return &SessionAffinityConfigApplyConfiguration{}
}

// SetClientIP sets the ClientIP field in the declarative configuration to the given value.
func (b *SessionAffinityConfigApplyConfiguration) SetClientIP(value *ClientIPConfigApplyConfiguration) *SessionAffinityConfigApplyConfiguration {
	b.ClientIP = value
	return b
}

// RemoveClientIP removes the ClientIP field from the declarative configuration.
func (b *SessionAffinityConfigApplyConfiguration) RemoveClientIP() *SessionAffinityConfigApplyConfiguration {
	b.ClientIP = nil
	return b
}

// GetClientIP gets the ClientIP field from the declarative configuration.
func (b *SessionAffinityConfigApplyConfiguration) GetClientIP() (value *ClientIPConfigApplyConfiguration, ok bool) {
	return b.ClientIP, b.ClientIP != nil
}

// SessionAffinityConfigList represents a listAlias of SessionAffinityConfigApplyConfiguration.
type SessionAffinityConfigList []*SessionAffinityConfigApplyConfiguration

// SessionAffinityConfigList represents a map of SessionAffinityConfigApplyConfiguration.
type SessionAffinityConfigMap map[string]SessionAffinityConfigApplyConfiguration
