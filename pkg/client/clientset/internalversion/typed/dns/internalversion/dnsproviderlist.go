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

package internalversion

import (
	"context"
	"time"

	dns "github.com/fire833/morfic/pkg/apis/dns"
	scheme "github.com/fire833/morfic/pkg/client/clientset/internalversion/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DNSProviderListsGetter has a method to return a DNSProviderListInterface.
// A group's client should implement this interface.
type DNSProviderListsGetter interface {
	DNSProviderLists(namespace string) DNSProviderListInterface
}

// DNSProviderListInterface has methods to work with DNSProviderList resources.
type DNSProviderListInterface interface {
	Create(ctx context.Context, dNSProviderList *dns.DNSProviderList, opts v1.CreateOptions) (*dns.DNSProviderList, error)
	Update(ctx context.Context, dNSProviderList *dns.DNSProviderList, opts v1.UpdateOptions) (*dns.DNSProviderList, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*dns.DNSProviderList, error)
	List(ctx context.Context, opts v1.ListOptions) (*dns.DNSProviderListList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *dns.DNSProviderList, err error)
	DNSProviderListExpansion
}

// dNSProviderLists implements DNSProviderListInterface
type dNSProviderLists struct {
	client rest.Interface
	ns     string
}

// newDNSProviderLists returns a DNSProviderLists
func newDNSProviderLists(c *DnsClient, namespace string) *dNSProviderLists {
	return &dNSProviderLists{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the dNSProviderList, and returns the corresponding dNSProviderList object, and an error if there is any.
func (c *dNSProviderLists) Get(ctx context.Context, name string, options v1.GetOptions) (result *dns.DNSProviderList, err error) {
	result = &dns.DNSProviderList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dnsproviderlists").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DNSProviderLists that match those selectors.
func (c *dNSProviderLists) List(ctx context.Context, opts v1.ListOptions) (result *dns.DNSProviderListList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &dns.DNSProviderListList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dnsproviderlists").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested dNSProviderLists.
func (c *dNSProviderLists) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("dnsproviderlists").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a dNSProviderList and creates it.  Returns the server's representation of the dNSProviderList, and an error, if there is any.
func (c *dNSProviderLists) Create(ctx context.Context, dNSProviderList *dns.DNSProviderList, opts v1.CreateOptions) (result *dns.DNSProviderList, err error) {
	result = &dns.DNSProviderList{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("dnsproviderlists").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(dNSProviderList).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a dNSProviderList and updates it. Returns the server's representation of the dNSProviderList, and an error, if there is any.
func (c *dNSProviderLists) Update(ctx context.Context, dNSProviderList *dns.DNSProviderList, opts v1.UpdateOptions) (result *dns.DNSProviderList, err error) {
	result = &dns.DNSProviderList{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("dnsproviderlists").
		Name(dNSProviderList.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(dNSProviderList).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the dNSProviderList and deletes it. Returns an error if one occurs.
func (c *dNSProviderLists) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dnsproviderlists").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *dNSProviderLists) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dnsproviderlists").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched dNSProviderList.
func (c *dNSProviderLists) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *dns.DNSProviderList, err error) {
	result = &dns.DNSProviderList{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("dnsproviderlists").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
