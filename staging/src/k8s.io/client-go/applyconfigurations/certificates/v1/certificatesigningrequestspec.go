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
	v1 "k8s.io/api/certificates/v1"
)

// CertificateSigningRequestSpecApplyConfiguration represents an declarative configuration of the CertificateSigningRequestSpec type for use
// with apply.
type CertificateSigningRequestSpecApplyConfiguration struct {
	Request    *[]byte                   `json:"request,omitempty"`
	SignerName *string                   `json:"signerName,omitempty"`
	Usages     *[]v1.KeyUsage            `json:"usages,omitempty"`
	Username   *string                   `json:"username,omitempty"`
	UID        *string                   `json:"uid,omitempty"`
	Groups     *[]string                 `json:"groups,omitempty"`
	Extra      *map[string]v1.ExtraValue `json:"extra,omitempty"`
}

// CertificateSigningRequestSpecApplyConfiguration constructs an declarative configuration of the CertificateSigningRequestSpec type for use with
// apply.
func CertificateSigningRequestSpec() *CertificateSigningRequestSpecApplyConfiguration {
	return &CertificateSigningRequestSpecApplyConfiguration{}
}

// SetRequest sets the Request field in the declarative configuration to the given value.
func (b *CertificateSigningRequestSpecApplyConfiguration) SetRequest(value []byte) *CertificateSigningRequestSpecApplyConfiguration {
	b.Request = &value
	return b
}

// RemoveRequest removes the Request field from the declarative configuration.
func (b *CertificateSigningRequestSpecApplyConfiguration) RemoveRequest() *CertificateSigningRequestSpecApplyConfiguration {
	b.Request = nil
	return b
}

// GetRequest gets the Request field from the declarative configuration.
func (b *CertificateSigningRequestSpecApplyConfiguration) GetRequest() (value []byte, ok bool) {
	if v := b.Request; v != nil {
		return *v, true
	}
	return value, false
}

// SetSignerName sets the SignerName field in the declarative configuration to the given value.
func (b *CertificateSigningRequestSpecApplyConfiguration) SetSignerName(value string) *CertificateSigningRequestSpecApplyConfiguration {
	b.SignerName = &value
	return b
}

// RemoveSignerName removes the SignerName field from the declarative configuration.
func (b *CertificateSigningRequestSpecApplyConfiguration) RemoveSignerName() *CertificateSigningRequestSpecApplyConfiguration {
	b.SignerName = nil
	return b
}

// GetSignerName gets the SignerName field from the declarative configuration.
func (b *CertificateSigningRequestSpecApplyConfiguration) GetSignerName() (value string, ok bool) {
	if v := b.SignerName; v != nil {
		return *v, true
	}
	return value, false
}

// SetUsages sets the Usages field in the declarative configuration to the given value.
func (b *CertificateSigningRequestSpecApplyConfiguration) SetUsages(value []v1.KeyUsage) *CertificateSigningRequestSpecApplyConfiguration {
	b.Usages = &value
	return b
}

// RemoveUsages removes the Usages field from the declarative configuration.
func (b *CertificateSigningRequestSpecApplyConfiguration) RemoveUsages() *CertificateSigningRequestSpecApplyConfiguration {
	b.Usages = nil
	return b
}

// GetUsages gets the Usages field from the declarative configuration.
func (b *CertificateSigningRequestSpecApplyConfiguration) GetUsages() (value []v1.KeyUsage, ok bool) {
	if v := b.Usages; v != nil {
		return *v, true
	}
	return value, false
}

// SetUsername sets the Username field in the declarative configuration to the given value.
func (b *CertificateSigningRequestSpecApplyConfiguration) SetUsername(value string) *CertificateSigningRequestSpecApplyConfiguration {
	b.Username = &value
	return b
}

// RemoveUsername removes the Username field from the declarative configuration.
func (b *CertificateSigningRequestSpecApplyConfiguration) RemoveUsername() *CertificateSigningRequestSpecApplyConfiguration {
	b.Username = nil
	return b
}

// GetUsername gets the Username field from the declarative configuration.
func (b *CertificateSigningRequestSpecApplyConfiguration) GetUsername() (value string, ok bool) {
	if v := b.Username; v != nil {
		return *v, true
	}
	return value, false
}

// SetUID sets the UID field in the declarative configuration to the given value.
func (b *CertificateSigningRequestSpecApplyConfiguration) SetUID(value string) *CertificateSigningRequestSpecApplyConfiguration {
	b.UID = &value
	return b
}

// RemoveUID removes the UID field from the declarative configuration.
func (b *CertificateSigningRequestSpecApplyConfiguration) RemoveUID() *CertificateSigningRequestSpecApplyConfiguration {
	b.UID = nil
	return b
}

// GetUID gets the UID field from the declarative configuration.
func (b *CertificateSigningRequestSpecApplyConfiguration) GetUID() (value string, ok bool) {
	if v := b.UID; v != nil {
		return *v, true
	}
	return value, false
}

// SetGroups sets the Groups field in the declarative configuration to the given value.
func (b *CertificateSigningRequestSpecApplyConfiguration) SetGroups(value []string) *CertificateSigningRequestSpecApplyConfiguration {
	b.Groups = &value
	return b
}

// RemoveGroups removes the Groups field from the declarative configuration.
func (b *CertificateSigningRequestSpecApplyConfiguration) RemoveGroups() *CertificateSigningRequestSpecApplyConfiguration {
	b.Groups = nil
	return b
}

// GetGroups gets the Groups field from the declarative configuration.
func (b *CertificateSigningRequestSpecApplyConfiguration) GetGroups() (value []string, ok bool) {
	if v := b.Groups; v != nil {
		return *v, true
	}
	return value, false
}

// SetExtra sets the Extra field in the declarative configuration to the given value.
func (b *CertificateSigningRequestSpecApplyConfiguration) SetExtra(value map[string]v1.ExtraValue) *CertificateSigningRequestSpecApplyConfiguration {
	b.Extra = &value
	return b
}

// RemoveExtra removes the Extra field from the declarative configuration.
func (b *CertificateSigningRequestSpecApplyConfiguration) RemoveExtra() *CertificateSigningRequestSpecApplyConfiguration {
	b.Extra = nil
	return b
}

// GetExtra gets the Extra field from the declarative configuration.
func (b *CertificateSigningRequestSpecApplyConfiguration) GetExtra() (value map[string]v1.ExtraValue, ok bool) {
	if v := b.Extra; v != nil {
		return *v, true
	}
	return value, false
}

// CertificateSigningRequestSpecList represents a listAlias of CertificateSigningRequestSpecApplyConfiguration.
type CertificateSigningRequestSpecList []*CertificateSigningRequestSpecApplyConfiguration

// CertificateSigningRequestSpecList represents a map of CertificateSigningRequestSpecApplyConfiguration.
type CertificateSigningRequestSpecMap map[string]CertificateSigningRequestSpecApplyConfiguration
