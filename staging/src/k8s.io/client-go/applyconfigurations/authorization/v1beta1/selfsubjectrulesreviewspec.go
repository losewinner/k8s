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

// SelfSubjectRulesReviewSpecApplyConfiguration represents an declarative configuration of the SelfSubjectRulesReviewSpec type for use
// with apply.
type SelfSubjectRulesReviewSpecApplyConfiguration struct {
	Namespace *string `json:"namespace,omitempty"`
}

// SelfSubjectRulesReviewSpecApplyConfiguration constructs an declarative configuration of the SelfSubjectRulesReviewSpec type for use with
// apply.
func SelfSubjectRulesReviewSpec() *SelfSubjectRulesReviewSpecApplyConfiguration {
	return &SelfSubjectRulesReviewSpecApplyConfiguration{}
}

// SetNamespace sets the Namespace field in the declarative configuration to the given value.
func (b *SelfSubjectRulesReviewSpecApplyConfiguration) SetNamespace(value string) *SelfSubjectRulesReviewSpecApplyConfiguration {
	b.Namespace = &value
	return b
}

// RemoveNamespace removes the Namespace field from the declarative configuration.
func (b *SelfSubjectRulesReviewSpecApplyConfiguration) RemoveNamespace() *SelfSubjectRulesReviewSpecApplyConfiguration {
	b.Namespace = nil
	return b
}

// GetNamespace gets the Namespace field from the declarative configuration.
func (b *SelfSubjectRulesReviewSpecApplyConfiguration) GetNamespace() (value string, ok bool) {
	if v := b.Namespace; v != nil {
		return *v, true
	}
	return value, false
}

// SelfSubjectRulesReviewSpecList represents a listAlias of SelfSubjectRulesReviewSpecApplyConfiguration.
type SelfSubjectRulesReviewSpecList []*SelfSubjectRulesReviewSpecApplyConfiguration

// SelfSubjectRulesReviewSpecList represents a map of SelfSubjectRulesReviewSpecApplyConfiguration.
type SelfSubjectRulesReviewSpecMap map[string]SelfSubjectRulesReviewSpecApplyConfiguration
