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

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	"context"

	v1beta1 "k8s.io/api/authorization/v1beta1"
	rest "k8s.io/client-go/rest"
)

// LocalSubjectAccessReviewsGetter has a method to return a LocalSubjectAccessReviewInterface.
// A group's client should implement this interface.
type LocalSubjectAccessReviewsGetter interface {
	LocalSubjectAccessReviews(namespace string) LocalSubjectAccessReviewInterface
}

// LocalSubjectAccessReviewInterface has methods to work with LocalSubjectAccessReview resources.
type LocalSubjectAccessReviewInterface interface {
	Create(*v1beta1.LocalSubjectAccessReview) (*v1beta1.LocalSubjectAccessReview, error)
	LocalSubjectAccessReviewExpansion
}

// localSubjectAccessReviews implements LocalSubjectAccessReviewInterface
type localSubjectAccessReviews struct {
	client rest.Interface
	ns     string
}

// newLocalSubjectAccessReviews returns a LocalSubjectAccessReviews
func newLocalSubjectAccessReviews(c *AuthorizationV1beta1Client, namespace string) *localSubjectAccessReviews {
	return &localSubjectAccessReviews{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Create takes the representation of a localSubjectAccessReview and creates it.  Returns the server's representation of the localSubjectAccessReview, and an error, if there is any.
func (c *localSubjectAccessReviews) Create(localSubjectAccessReview *v1beta1.LocalSubjectAccessReview) (result *v1beta1.LocalSubjectAccessReview, err error) {
	result = &v1beta1.LocalSubjectAccessReview{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("localsubjectaccessreviews").
		Body(localSubjectAccessReview).
		Do(context.TODO()).
		Into(result)
	return
}
