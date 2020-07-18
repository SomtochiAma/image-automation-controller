// +build !ignore_autogenerated

/*
Copyright 2020 Michael Bridgen <mikeb@squaremobius.net>

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageUpdateAutomation) DeepCopyInto(out *ImageUpdateAutomation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageUpdateAutomation.
func (in *ImageUpdateAutomation) DeepCopy() *ImageUpdateAutomation {
	if in == nil {
		return nil
	}
	out := new(ImageUpdateAutomation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImageUpdateAutomation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageUpdateAutomationList) DeepCopyInto(out *ImageUpdateAutomationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ImageUpdateAutomation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageUpdateAutomationList.
func (in *ImageUpdateAutomationList) DeepCopy() *ImageUpdateAutomationList {
	if in == nil {
		return nil
	}
	out := new(ImageUpdateAutomationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImageUpdateAutomationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageUpdateAutomationSpec) DeepCopyInto(out *ImageUpdateAutomationSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageUpdateAutomationSpec.
func (in *ImageUpdateAutomationSpec) DeepCopy() *ImageUpdateAutomationSpec {
	if in == nil {
		return nil
	}
	out := new(ImageUpdateAutomationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageUpdateAutomationStatus) DeepCopyInto(out *ImageUpdateAutomationStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageUpdateAutomationStatus.
func (in *ImageUpdateAutomationStatus) DeepCopy() *ImageUpdateAutomationStatus {
	if in == nil {
		return nil
	}
	out := new(ImageUpdateAutomationStatus)
	in.DeepCopyInto(out)
	return out
}
