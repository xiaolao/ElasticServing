// +build !ignore_autogenerated

/*


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

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
	"knative.dev/pkg/apis/duck/v1beta1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointSpec) DeepCopyInto(out *EndpointSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointSpec.
func (in *EndpointSpec) DeepCopy() *EndpointSpec {
	if in == nil {
		return nil
	}
	out := new(EndpointSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PaddleService) DeepCopyInto(out *PaddleService) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PaddleService.
func (in *PaddleService) DeepCopy() *PaddleService {
	if in == nil {
		return nil
	}
	out := new(PaddleService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PaddleService) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PaddleServiceList) DeepCopyInto(out *PaddleServiceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PaddleService, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PaddleServiceList.
func (in *PaddleServiceList) DeepCopy() *PaddleServiceList {
	if in == nil {
		return nil
	}
	out := new(PaddleServiceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PaddleServiceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PaddleServiceSpec) DeepCopyInto(out *PaddleServiceSpec) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
	if in.Default != nil {
		in, out := &in.Default, &out.Default
		*out = new(EndpointSpec)
		**out = **in
	}
	if in.Canary != nil {
		in, out := &in.Canary, &out.Canary
		*out = new(EndpointSpec)
		**out = **in
	}
	if in.CanaryTrafficPercent != nil {
		in, out := &in.CanaryTrafficPercent, &out.CanaryTrafficPercent
		*out = new(int)
		**out = **in
	}
	in.Service.DeepCopyInto(&out.Service)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PaddleServiceSpec.
func (in *PaddleServiceSpec) DeepCopy() *PaddleServiceSpec {
	if in == nil {
		return nil
	}
	out := new(PaddleServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PaddleServiceStatus) DeepCopyInto(out *PaddleServiceStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	if in.Default != nil {
		in, out := &in.Default, &out.Default
		*out = new(StatusConfigurationSpec)
		**out = **in
	}
	if in.Canary != nil {
		in, out := &in.Canary, &out.Canary
		*out = new(StatusConfigurationSpec)
		**out = **in
	}
	if in.Address != nil {
		in, out := &in.Address, &out.Address
		*out = new(v1beta1.Addressable)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PaddleServiceStatus.
func (in *PaddleServiceStatus) DeepCopy() *PaddleServiceStatus {
	if in == nil {
		return nil
	}
	out := new(PaddleServiceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceSpec) DeepCopyInto(out *ServiceSpec) {
	*out = *in
	if in.MinScale != nil {
		in, out := &in.MinScale, &out.MinScale
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceSpec.
func (in *ServiceSpec) DeepCopy() *ServiceSpec {
	if in == nil {
		return nil
	}
	out := new(ServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatusConfigurationSpec) DeepCopyInto(out *StatusConfigurationSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatusConfigurationSpec.
func (in *StatusConfigurationSpec) DeepCopy() *StatusConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(StatusConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}
