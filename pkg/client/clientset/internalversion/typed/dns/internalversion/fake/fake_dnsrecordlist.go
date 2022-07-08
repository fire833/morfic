/*
*	Copyright (C) 2022  Kendall Tauser
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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	dns "github.com/fire833/morfic/pkg/apis/dns"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDNSRecordLists implements DNSRecordListInterface
type FakeDNSRecordLists struct {
	Fake *FakeDns
	ns   string
}

var dnsrecordlistsResource = schema.GroupVersionResource{Group: "dns.vroute.io", Version: "", Resource: "dnsrecordlists"}

var dnsrecordlistsKind = schema.GroupVersionKind{Group: "dns.vroute.io", Version: "", Kind: "DNSRecordList"}

// Get takes name of the dNSRecordList, and returns the corresponding dNSRecordList object, and an error if there is any.
func (c *FakeDNSRecordLists) Get(ctx context.Context, name string, options v1.GetOptions) (result *dns.DNSRecordList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(dnsrecordlistsResource, c.ns, name), &dns.DNSRecordList{})

	if obj == nil {
		return nil, err
	}
	return obj.(*dns.DNSRecordList), err
}

// List takes label and field selectors, and returns the list of DNSRecordLists that match those selectors.
func (c *FakeDNSRecordLists) List(ctx context.Context, opts v1.ListOptions) (result *dns.DNSRecordListList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(dnsrecordlistsResource, dnsrecordlistsKind, c.ns, opts), &dns.DNSRecordListList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &dns.DNSRecordListList{ListMeta: obj.(*dns.DNSRecordListList).ListMeta}
	for _, item := range obj.(*dns.DNSRecordListList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dNSRecordLists.
func (c *FakeDNSRecordLists) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(dnsrecordlistsResource, c.ns, opts))

}

// Create takes the representation of a dNSRecordList and creates it.  Returns the server's representation of the dNSRecordList, and an error, if there is any.
func (c *FakeDNSRecordLists) Create(ctx context.Context, dNSRecordList *dns.DNSRecordList, opts v1.CreateOptions) (result *dns.DNSRecordList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(dnsrecordlistsResource, c.ns, dNSRecordList), &dns.DNSRecordList{})

	if obj == nil {
		return nil, err
	}
	return obj.(*dns.DNSRecordList), err
}

// Update takes the representation of a dNSRecordList and updates it. Returns the server's representation of the dNSRecordList, and an error, if there is any.
func (c *FakeDNSRecordLists) Update(ctx context.Context, dNSRecordList *dns.DNSRecordList, opts v1.UpdateOptions) (result *dns.DNSRecordList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(dnsrecordlistsResource, c.ns, dNSRecordList), &dns.DNSRecordList{})

	if obj == nil {
		return nil, err
	}
	return obj.(*dns.DNSRecordList), err
}

// Delete takes name of the dNSRecordList and deletes it. Returns an error if one occurs.
func (c *FakeDNSRecordLists) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(dnsrecordlistsResource, c.ns, name, opts), &dns.DNSRecordList{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDNSRecordLists) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(dnsrecordlistsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &dns.DNSRecordListList{})
	return err
}

// Patch applies the patch and returns the patched dNSRecordList.
func (c *FakeDNSRecordLists) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *dns.DNSRecordList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(dnsrecordlistsResource, c.ns, name, pt, data, subresources...), &dns.DNSRecordList{})

	if obj == nil {
		return nil, err
	}
	return obj.(*dns.DNSRecordList), err
}