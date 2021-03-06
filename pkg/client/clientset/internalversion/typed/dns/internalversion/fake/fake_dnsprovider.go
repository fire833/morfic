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

// FakeDNSProviders implements DNSProviderInterface
type FakeDNSProviders struct {
	Fake *FakeDns
	ns   string
}

var dnsprovidersResource = schema.GroupVersionResource{Group: "dns.vroute.io", Version: "", Resource: "dnsproviders"}

var dnsprovidersKind = schema.GroupVersionKind{Group: "dns.vroute.io", Version: "", Kind: "DNSProvider"}

// Get takes name of the dNSProvider, and returns the corresponding dNSProvider object, and an error if there is any.
func (c *FakeDNSProviders) Get(ctx context.Context, name string, options v1.GetOptions) (result *dns.DNSProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(dnsprovidersResource, c.ns, name), &dns.DNSProvider{})

	if obj == nil {
		return nil, err
	}
	return obj.(*dns.DNSProvider), err
}

// List takes label and field selectors, and returns the list of DNSProviders that match those selectors.
func (c *FakeDNSProviders) List(ctx context.Context, opts v1.ListOptions) (result *dns.DNSProviderList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(dnsprovidersResource, dnsprovidersKind, c.ns, opts), &dns.DNSProviderList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &dns.DNSProviderList{ListMeta: obj.(*dns.DNSProviderList).ListMeta}
	for _, item := range obj.(*dns.DNSProviderList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dNSProviders.
func (c *FakeDNSProviders) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(dnsprovidersResource, c.ns, opts))

}

// Create takes the representation of a dNSProvider and creates it.  Returns the server's representation of the dNSProvider, and an error, if there is any.
func (c *FakeDNSProviders) Create(ctx context.Context, dNSProvider *dns.DNSProvider, opts v1.CreateOptions) (result *dns.DNSProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(dnsprovidersResource, c.ns, dNSProvider), &dns.DNSProvider{})

	if obj == nil {
		return nil, err
	}
	return obj.(*dns.DNSProvider), err
}

// Update takes the representation of a dNSProvider and updates it. Returns the server's representation of the dNSProvider, and an error, if there is any.
func (c *FakeDNSProviders) Update(ctx context.Context, dNSProvider *dns.DNSProvider, opts v1.UpdateOptions) (result *dns.DNSProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(dnsprovidersResource, c.ns, dNSProvider), &dns.DNSProvider{})

	if obj == nil {
		return nil, err
	}
	return obj.(*dns.DNSProvider), err
}

// Delete takes name of the dNSProvider and deletes it. Returns an error if one occurs.
func (c *FakeDNSProviders) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(dnsprovidersResource, c.ns, name, opts), &dns.DNSProvider{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDNSProviders) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(dnsprovidersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &dns.DNSProviderList{})
	return err
}

// Patch applies the patch and returns the patched dNSProvider.
func (c *FakeDNSProviders) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *dns.DNSProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(dnsprovidersResource, c.ns, name, pt, data, subresources...), &dns.DNSProvider{})

	if obj == nil {
		return nil, err
	}
	return obj.(*dns.DNSProvider), err
}
