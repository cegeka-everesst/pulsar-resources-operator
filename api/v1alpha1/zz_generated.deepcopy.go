//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright (c) 2022 StreamNative, Inc.. All Rights Reserved.

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PulsarConnection) DeepCopyInto(out *PulsarConnection) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PulsarConnection.
func (in *PulsarConnection) DeepCopy() *PulsarConnection {
	if in == nil {
		return nil
	}
	out := new(PulsarConnection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PulsarConnection) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PulsarConnectionList) DeepCopyInto(out *PulsarConnectionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PulsarConnection, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PulsarConnectionList.
func (in *PulsarConnectionList) DeepCopy() *PulsarConnectionList {
	if in == nil {
		return nil
	}
	out := new(PulsarConnectionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PulsarConnectionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PulsarConnectionSpec) DeepCopyInto(out *PulsarConnectionSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PulsarConnectionSpec.
func (in *PulsarConnectionSpec) DeepCopy() *PulsarConnectionSpec {
	if in == nil {
		return nil
	}
	out := new(PulsarConnectionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PulsarConnectionStatus) DeepCopyInto(out *PulsarConnectionStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PulsarConnectionStatus.
func (in *PulsarConnectionStatus) DeepCopy() *PulsarConnectionStatus {
	if in == nil {
		return nil
	}
	out := new(PulsarConnectionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PulsarNamespace) DeepCopyInto(out *PulsarNamespace) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PulsarNamespace.
func (in *PulsarNamespace) DeepCopy() *PulsarNamespace {
	if in == nil {
		return nil
	}
	out := new(PulsarNamespace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PulsarNamespace) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PulsarNamespaceList) DeepCopyInto(out *PulsarNamespaceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PulsarNamespace, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PulsarNamespaceList.
func (in *PulsarNamespaceList) DeepCopy() *PulsarNamespaceList {
	if in == nil {
		return nil
	}
	out := new(PulsarNamespaceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PulsarNamespaceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PulsarNamespaceSpec) DeepCopyInto(out *PulsarNamespaceSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PulsarNamespaceSpec.
func (in *PulsarNamespaceSpec) DeepCopy() *PulsarNamespaceSpec {
	if in == nil {
		return nil
	}
	out := new(PulsarNamespaceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PulsarNamespaceStatus) DeepCopyInto(out *PulsarNamespaceStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PulsarNamespaceStatus.
func (in *PulsarNamespaceStatus) DeepCopy() *PulsarNamespaceStatus {
	if in == nil {
		return nil
	}
	out := new(PulsarNamespaceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PulsarPermission) DeepCopyInto(out *PulsarPermission) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PulsarPermission.
func (in *PulsarPermission) DeepCopy() *PulsarPermission {
	if in == nil {
		return nil
	}
	out := new(PulsarPermission)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PulsarPermission) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PulsarPermissionList) DeepCopyInto(out *PulsarPermissionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PulsarPermission, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PulsarPermissionList.
func (in *PulsarPermissionList) DeepCopy() *PulsarPermissionList {
	if in == nil {
		return nil
	}
	out := new(PulsarPermissionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PulsarPermissionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PulsarPermissionSpec) DeepCopyInto(out *PulsarPermissionSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PulsarPermissionSpec.
func (in *PulsarPermissionSpec) DeepCopy() *PulsarPermissionSpec {
	if in == nil {
		return nil
	}
	out := new(PulsarPermissionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PulsarPermissionStatus) DeepCopyInto(out *PulsarPermissionStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PulsarPermissionStatus.
func (in *PulsarPermissionStatus) DeepCopy() *PulsarPermissionStatus {
	if in == nil {
		return nil
	}
	out := new(PulsarPermissionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PulsarTenant) DeepCopyInto(out *PulsarTenant) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PulsarTenant.
func (in *PulsarTenant) DeepCopy() *PulsarTenant {
	if in == nil {
		return nil
	}
	out := new(PulsarTenant)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PulsarTenant) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PulsarTenantList) DeepCopyInto(out *PulsarTenantList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PulsarTenant, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PulsarTenantList.
func (in *PulsarTenantList) DeepCopy() *PulsarTenantList {
	if in == nil {
		return nil
	}
	out := new(PulsarTenantList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PulsarTenantList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PulsarTenantSpec) DeepCopyInto(out *PulsarTenantSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PulsarTenantSpec.
func (in *PulsarTenantSpec) DeepCopy() *PulsarTenantSpec {
	if in == nil {
		return nil
	}
	out := new(PulsarTenantSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PulsarTenantStatus) DeepCopyInto(out *PulsarTenantStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PulsarTenantStatus.
func (in *PulsarTenantStatus) DeepCopy() *PulsarTenantStatus {
	if in == nil {
		return nil
	}
	out := new(PulsarTenantStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PulsarTopic) DeepCopyInto(out *PulsarTopic) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PulsarTopic.
func (in *PulsarTopic) DeepCopy() *PulsarTopic {
	if in == nil {
		return nil
	}
	out := new(PulsarTopic)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PulsarTopic) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PulsarTopicList) DeepCopyInto(out *PulsarTopicList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PulsarTopic, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PulsarTopicList.
func (in *PulsarTopicList) DeepCopy() *PulsarTopicList {
	if in == nil {
		return nil
	}
	out := new(PulsarTopicList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PulsarTopicList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PulsarTopicSpec) DeepCopyInto(out *PulsarTopicSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PulsarTopicSpec.
func (in *PulsarTopicSpec) DeepCopy() *PulsarTopicSpec {
	if in == nil {
		return nil
	}
	out := new(PulsarTopicSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PulsarTopicStatus) DeepCopyInto(out *PulsarTopicStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PulsarTopicStatus.
func (in *PulsarTopicStatus) DeepCopy() *PulsarTopicStatus {
	if in == nil {
		return nil
	}
	out := new(PulsarTopicStatus)
	in.DeepCopyInto(out)
	return out
}
