// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TokenReview) DeepCopyInto(out *TokenReview) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new TokenReview.
func (x *TokenReview) DeepCopy() *TokenReview {
	if x == nil {
		return nil
	}
	out := new(TokenReview)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (x *TokenReview) DeepCopyObject() runtime.Object {
	if c := x.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TokenReviewSpec) DeepCopyInto(out *TokenReviewSpec) {
	*out = *in
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new TokenReviewSpec.
func (x *TokenReviewSpec) DeepCopy() *TokenReviewSpec {
	if x == nil {
		return nil
	}
	out := new(TokenReviewSpec)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TokenReviewStatus) DeepCopyInto(out *TokenReviewStatus) {
	*out = *in
	in.User.DeepCopyInto(&out.User)
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new TokenReviewStatus.
func (x *TokenReviewStatus) DeepCopy() *TokenReviewStatus {
	if x == nil {
		return nil
	}
	out := new(TokenReviewStatus)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserInfo) DeepCopyInto(out *UserInfo) {
	*out = *in
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Extra != nil {
		in, out := &in.Extra, &out.Extra
		*out = make(map[string]ExtraValue, len(*in))
		for key, val := range *in {
			(*out)[key] = make(ExtraValue, len(val))
			copy((*out)[key], val)
		}
	}
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new UserInfo.
func (x *UserInfo) DeepCopy() *UserInfo {
	if x == nil {
		return nil
	}
	out := new(UserInfo)
	x.DeepCopyInto(out)
	return out
}
