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

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1beta1

import (
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	certificates "k8s.io/kubernetes/pkg/apis/certificates"
	unsafe "unsafe"
)

func init() {
	SchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1beta1_CertificateSigningRequest_To_certificates_CertificateSigningRequest,
		Convert_certificates_CertificateSigningRequest_To_v1beta1_CertificateSigningRequest,
		Convert_v1beta1_CertificateSigningRequestCondition_To_certificates_CertificateSigningRequestCondition,
		Convert_certificates_CertificateSigningRequestCondition_To_v1beta1_CertificateSigningRequestCondition,
		Convert_v1beta1_CertificateSigningRequestList_To_certificates_CertificateSigningRequestList,
		Convert_certificates_CertificateSigningRequestList_To_v1beta1_CertificateSigningRequestList,
		Convert_v1beta1_CertificateSigningRequestSpec_To_certificates_CertificateSigningRequestSpec,
		Convert_certificates_CertificateSigningRequestSpec_To_v1beta1_CertificateSigningRequestSpec,
		Convert_v1beta1_CertificateSigningRequestStatus_To_certificates_CertificateSigningRequestStatus,
		Convert_certificates_CertificateSigningRequestStatus_To_v1beta1_CertificateSigningRequestStatus,
	)
}

func autoConvert_v1beta1_CertificateSigningRequest_To_certificates_CertificateSigningRequest(in *CertificateSigningRequest, out *certificates.CertificateSigningRequest, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta1_CertificateSigningRequestSpec_To_certificates_CertificateSigningRequestSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_CertificateSigningRequestStatus_To_certificates_CertificateSigningRequestStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_CertificateSigningRequest_To_certificates_CertificateSigningRequest is an autogenerated conversion function.
func Convert_v1beta1_CertificateSigningRequest_To_certificates_CertificateSigningRequest(in *CertificateSigningRequest, out *certificates.CertificateSigningRequest, s conversion.Scope) error {
	return autoConvert_v1beta1_CertificateSigningRequest_To_certificates_CertificateSigningRequest(in, out, s)
}

func autoConvert_certificates_CertificateSigningRequest_To_v1beta1_CertificateSigningRequest(in *certificates.CertificateSigningRequest, out *CertificateSigningRequest, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_certificates_CertificateSigningRequestSpec_To_v1beta1_CertificateSigningRequestSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_certificates_CertificateSigningRequestStatus_To_v1beta1_CertificateSigningRequestStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_certificates_CertificateSigningRequest_To_v1beta1_CertificateSigningRequest is an autogenerated conversion function.
func Convert_certificates_CertificateSigningRequest_To_v1beta1_CertificateSigningRequest(in *certificates.CertificateSigningRequest, out *CertificateSigningRequest, s conversion.Scope) error {
	return autoConvert_certificates_CertificateSigningRequest_To_v1beta1_CertificateSigningRequest(in, out, s)
}

func autoConvert_v1beta1_CertificateSigningRequestCondition_To_certificates_CertificateSigningRequestCondition(in *CertificateSigningRequestCondition, out *certificates.CertificateSigningRequestCondition, s conversion.Scope) error {
	out.Type = certificates.RequestConditionType(in.Type)
	out.Reason = in.Reason
	out.Message = in.Message
	out.LastUpdateTime = in.LastUpdateTime
	return nil
}

// Convert_v1beta1_CertificateSigningRequestCondition_To_certificates_CertificateSigningRequestCondition is an autogenerated conversion function.
func Convert_v1beta1_CertificateSigningRequestCondition_To_certificates_CertificateSigningRequestCondition(in *CertificateSigningRequestCondition, out *certificates.CertificateSigningRequestCondition, s conversion.Scope) error {
	return autoConvert_v1beta1_CertificateSigningRequestCondition_To_certificates_CertificateSigningRequestCondition(in, out, s)
}

func autoConvert_certificates_CertificateSigningRequestCondition_To_v1beta1_CertificateSigningRequestCondition(in *certificates.CertificateSigningRequestCondition, out *CertificateSigningRequestCondition, s conversion.Scope) error {
	out.Type = RequestConditionType(in.Type)
	out.Reason = in.Reason
	out.Message = in.Message
	out.LastUpdateTime = in.LastUpdateTime
	return nil
}

// Convert_certificates_CertificateSigningRequestCondition_To_v1beta1_CertificateSigningRequestCondition is an autogenerated conversion function.
func Convert_certificates_CertificateSigningRequestCondition_To_v1beta1_CertificateSigningRequestCondition(in *certificates.CertificateSigningRequestCondition, out *CertificateSigningRequestCondition, s conversion.Scope) error {
	return autoConvert_certificates_CertificateSigningRequestCondition_To_v1beta1_CertificateSigningRequestCondition(in, out, s)
}

func autoConvert_v1beta1_CertificateSigningRequestList_To_certificates_CertificateSigningRequestList(in *CertificateSigningRequestList, out *certificates.CertificateSigningRequestList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]certificates.CertificateSigningRequest)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1beta1_CertificateSigningRequestList_To_certificates_CertificateSigningRequestList is an autogenerated conversion function.
func Convert_v1beta1_CertificateSigningRequestList_To_certificates_CertificateSigningRequestList(in *CertificateSigningRequestList, out *certificates.CertificateSigningRequestList, s conversion.Scope) error {
	return autoConvert_v1beta1_CertificateSigningRequestList_To_certificates_CertificateSigningRequestList(in, out, s)
}

func autoConvert_certificates_CertificateSigningRequestList_To_v1beta1_CertificateSigningRequestList(in *certificates.CertificateSigningRequestList, out *CertificateSigningRequestList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]CertificateSigningRequest)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_certificates_CertificateSigningRequestList_To_v1beta1_CertificateSigningRequestList is an autogenerated conversion function.
func Convert_certificates_CertificateSigningRequestList_To_v1beta1_CertificateSigningRequestList(in *certificates.CertificateSigningRequestList, out *CertificateSigningRequestList, s conversion.Scope) error {
	return autoConvert_certificates_CertificateSigningRequestList_To_v1beta1_CertificateSigningRequestList(in, out, s)
}

func autoConvert_v1beta1_CertificateSigningRequestSpec_To_certificates_CertificateSigningRequestSpec(in *CertificateSigningRequestSpec, out *certificates.CertificateSigningRequestSpec, s conversion.Scope) error {
	out.Request = *(*[]byte)(unsafe.Pointer(&in.Request))
	out.Usages = *(*[]certificates.KeyUsage)(unsafe.Pointer(&in.Usages))
	out.Username = in.Username
	out.UID = in.UID
	out.Groups = *(*[]string)(unsafe.Pointer(&in.Groups))
	out.Extra = *(*map[string]certificates.ExtraValue)(unsafe.Pointer(&in.Extra))
	return nil
}

// Convert_v1beta1_CertificateSigningRequestSpec_To_certificates_CertificateSigningRequestSpec is an autogenerated conversion function.
func Convert_v1beta1_CertificateSigningRequestSpec_To_certificates_CertificateSigningRequestSpec(in *CertificateSigningRequestSpec, out *certificates.CertificateSigningRequestSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_CertificateSigningRequestSpec_To_certificates_CertificateSigningRequestSpec(in, out, s)
}

func autoConvert_certificates_CertificateSigningRequestSpec_To_v1beta1_CertificateSigningRequestSpec(in *certificates.CertificateSigningRequestSpec, out *CertificateSigningRequestSpec, s conversion.Scope) error {
	out.Request = *(*[]byte)(unsafe.Pointer(&in.Request))
	out.Usages = *(*[]KeyUsage)(unsafe.Pointer(&in.Usages))
	out.Username = in.Username
	out.UID = in.UID
	out.Groups = *(*[]string)(unsafe.Pointer(&in.Groups))
	out.Extra = *(*map[string]ExtraValue)(unsafe.Pointer(&in.Extra))
	return nil
}

// Convert_certificates_CertificateSigningRequestSpec_To_v1beta1_CertificateSigningRequestSpec is an autogenerated conversion function.
func Convert_certificates_CertificateSigningRequestSpec_To_v1beta1_CertificateSigningRequestSpec(in *certificates.CertificateSigningRequestSpec, out *CertificateSigningRequestSpec, s conversion.Scope) error {
	return autoConvert_certificates_CertificateSigningRequestSpec_To_v1beta1_CertificateSigningRequestSpec(in, out, s)
}

func autoConvert_v1beta1_CertificateSigningRequestStatus_To_certificates_CertificateSigningRequestStatus(in *CertificateSigningRequestStatus, out *certificates.CertificateSigningRequestStatus, s conversion.Scope) error {
	out.Conditions = *(*[]certificates.CertificateSigningRequestCondition)(unsafe.Pointer(&in.Conditions))
	out.Certificate = *(*[]byte)(unsafe.Pointer(&in.Certificate))
	return nil
}

// Convert_v1beta1_CertificateSigningRequestStatus_To_certificates_CertificateSigningRequestStatus is an autogenerated conversion function.
func Convert_v1beta1_CertificateSigningRequestStatus_To_certificates_CertificateSigningRequestStatus(in *CertificateSigningRequestStatus, out *certificates.CertificateSigningRequestStatus, s conversion.Scope) error {
	return autoConvert_v1beta1_CertificateSigningRequestStatus_To_certificates_CertificateSigningRequestStatus(in, out, s)
}

func autoConvert_certificates_CertificateSigningRequestStatus_To_v1beta1_CertificateSigningRequestStatus(in *certificates.CertificateSigningRequestStatus, out *CertificateSigningRequestStatus, s conversion.Scope) error {
	out.Conditions = *(*[]CertificateSigningRequestCondition)(unsafe.Pointer(&in.Conditions))
	out.Certificate = *(*[]byte)(unsafe.Pointer(&in.Certificate))
	return nil
}

// Convert_certificates_CertificateSigningRequestStatus_To_v1beta1_CertificateSigningRequestStatus is an autogenerated conversion function.
func Convert_certificates_CertificateSigningRequestStatus_To_v1beta1_CertificateSigningRequestStatus(in *certificates.CertificateSigningRequestStatus, out *CertificateSigningRequestStatus, s conversion.Scope) error {
	return autoConvert_certificates_CertificateSigningRequestStatus_To_v1beta1_CertificateSigningRequestStatus(in, out, s)
}
