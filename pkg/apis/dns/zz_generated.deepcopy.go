//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
*	Copyright (C) 2023 Kendall Tauser
*
*	This program is free software; you can redistribute it and/or modify
*	it under the terms of the GNU General Public License as published by
*	the Free Software Foundation; either version 2 of the License, or
*	(at your option) any later version.
*
*	This program is distributed in the hope that it will be useful,
*	but WITHOUT ANY WARRANTY; without even the implied warranty of
*	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
*	GNU General Public License for more details.
*
*	You should have received a copy of the GNU General Public License along
*	with this program; if not, write to the Free Software Foundation, Inc.,
*	51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA.
 */

// Code generated by deepcopy-gen. DO NOT EDIT.

package dns

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudflareProviderSpec) DeepCopyInto(out *CloudflareProviderSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudflareProviderSpec.
func (in *CloudflareProviderSpec) DeepCopy() *CloudflareProviderSpec {
	if in == nil {
		return nil
	}
	out := new(CloudflareProviderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSProvider) DeepCopyInto(out *DNSProvider) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSProvider.
func (in *DNSProvider) DeepCopy() *DNSProvider {
	if in == nil {
		return nil
	}
	out := new(DNSProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DNSProvider) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSProviderList) DeepCopyInto(out *DNSProviderList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DNSProvider, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSProviderList.
func (in *DNSProviderList) DeepCopy() *DNSProviderList {
	if in == nil {
		return nil
	}
	out := new(DNSProviderList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DNSProviderList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSProviderSpec) DeepCopyInto(out *DNSProviderSpec) {
	*out = *in
	out.Cloudflare = in.Cloudflare
	out.Route53 = in.Route53
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSProviderSpec.
func (in *DNSProviderSpec) DeepCopy() *DNSProviderSpec {
	if in == nil {
		return nil
	}
	out := new(DNSProviderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSProviderStatus) DeepCopyInto(out *DNSProviderStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSProviderStatus.
func (in *DNSProviderStatus) DeepCopy() *DNSProviderStatus {
	if in == nil {
		return nil
	}
	out := new(DNSProviderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSRecord) DeepCopyInto(out *DNSRecord) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSRecord.
func (in *DNSRecord) DeepCopy() *DNSRecord {
	if in == nil {
		return nil
	}
	out := new(DNSRecord)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DNSRecord) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSRecordList) DeepCopyInto(out *DNSRecordList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DNSRecord, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSRecordList.
func (in *DNSRecordList) DeepCopy() *DNSRecordList {
	if in == nil {
		return nil
	}
	out := new(DNSRecordList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DNSRecordList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSRecordSpec) DeepCopyInto(out *DNSRecordSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSRecordSpec.
func (in *DNSRecordSpec) DeepCopy() *DNSRecordSpec {
	if in == nil {
		return nil
	}
	out := new(DNSRecordSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSRecordStatus) DeepCopyInto(out *DNSRecordStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSRecordStatus.
func (in *DNSRecordStatus) DeepCopy() *DNSRecordStatus {
	if in == nil {
		return nil
	}
	out := new(DNSRecordStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Route53ProviderSpec) DeepCopyInto(out *Route53ProviderSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Route53ProviderSpec.
func (in *Route53ProviderSpec) DeepCopy() *Route53ProviderSpec {
	if in == nil {
		return nil
	}
	out := new(Route53ProviderSpec)
	in.DeepCopyInto(out)
	return out
}
