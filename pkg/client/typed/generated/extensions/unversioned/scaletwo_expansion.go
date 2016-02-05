/*
Copyright 2016 The Kubernetes Authors All rights reserved.

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

package unversioned

import (
	"k8s.io/kubernetes/pkg/api/meta"
	"k8s.io/kubernetes/pkg/api/unversioned"
	"k8s.io/kubernetes/pkg/apis/extensions"
)

// The ScaleTwoExpansion interface allows manually adding extra methods to the ScaleInterface.
type ScaleTwoExpansion interface {
	Get(kind string, name string) (*extensions.ScaleTwo, error)
	Update(kind string, scale *extensions.ScaleTwo) (*extensions.ScaleTwo, error)
}

// Get takes the reference to scale subresource and returns the subresource or error, if one occurs.
func (c *scaleTwos) Get(kind string, name string) (result *extensions.ScaleTwo, err error) {
	result = &extensions.ScaleTwo{}

	// TODO this method needs to take a proper unambiguous kind
	fullyQualifiedKind := unversioned.GroupVersionKind{Kind: kind}
	resource, _ := meta.KindToResource(fullyQualifiedKind, false)

	err = c.client.Get().
		Namespace(c.ns).
		Resource(resource.Resource).
		Name(name).
		SubResource("scale").
		Do().
		Into(result)
	return
}

func (c *scaleTwos) Update(kind string, scale *extensions.ScaleTwo) (result *extensions.ScaleTwo, err error) {
	result = &extensions.ScaleTwo{}

	// TODO this method needs to take a proper unambiguous kind
	fullyQualifiedKind := unversioned.GroupVersionKind{Kind: kind}
	resource, _ := meta.KindToResource(fullyQualifiedKind, false)

	err = c.client.Put().
		Namespace(scale.Namespace).
		Resource(resource.Resource).
		Name(scale.Name).
		SubResource("scale").
		Body(scale).
		Do().
		Into(result)
	return
}
