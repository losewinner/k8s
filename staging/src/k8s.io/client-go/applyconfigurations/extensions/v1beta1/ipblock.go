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

// IPBlockApplyConfiguration represents an declarative configuration of the IPBlock type for use
// with apply.
type IPBlockApplyConfiguration struct {
	CIDR   *string   `json:"cidr,omitempty"`
	Except *[]string `json:"except,omitempty"`
}

// IPBlockApplyConfiguration constructs an declarative configuration of the IPBlock type for use with
// apply.
func IPBlock() *IPBlockApplyConfiguration {
	return &IPBlockApplyConfiguration{}
}

// SetCIDR sets the CIDR field in the declarative configuration to the given value.
func (b *IPBlockApplyConfiguration) SetCIDR(value string) *IPBlockApplyConfiguration {
	b.CIDR = &value
	return b
}

// RemoveCIDR removes the CIDR field from the declarative configuration.
func (b *IPBlockApplyConfiguration) RemoveCIDR() *IPBlockApplyConfiguration {
	b.CIDR = nil
	return b
}

// GetCIDR gets the CIDR field from the declarative configuration.
func (b *IPBlockApplyConfiguration) GetCIDR() (value string, ok bool) {
	if v := b.CIDR; v != nil {
		return *v, true
	}
	return value, false
}

// SetExcept sets the Except field in the declarative configuration to the given value.
func (b *IPBlockApplyConfiguration) SetExcept(value []string) *IPBlockApplyConfiguration {
	b.Except = &value
	return b
}

// RemoveExcept removes the Except field from the declarative configuration.
func (b *IPBlockApplyConfiguration) RemoveExcept() *IPBlockApplyConfiguration {
	b.Except = nil
	return b
}

// GetExcept gets the Except field from the declarative configuration.
func (b *IPBlockApplyConfiguration) GetExcept() (value []string, ok bool) {
	if v := b.Except; v != nil {
		return *v, true
	}
	return value, false
}

// IPBlockList represents a listAlias of IPBlockApplyConfiguration.
type IPBlockList []*IPBlockApplyConfiguration

// IPBlockList represents a map of IPBlockApplyConfiguration.
type IPBlockMap map[string]IPBlockApplyConfiguration
