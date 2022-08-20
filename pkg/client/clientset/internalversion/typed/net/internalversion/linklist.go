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

	net "github.com/fire833/morfic/pkg/apis/net"
	scheme "github.com/fire833/morfic/pkg/client/clientset/internalversion/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// LinkListsGetter has a method to return a LinkListInterface.
// A group's client should implement this interface.
type LinkListsGetter interface {
	LinkLists(namespace string) LinkListInterface
}

// LinkListInterface has methods to work with LinkList resources.
type LinkListInterface interface {
	Create(ctx context.Context, linkList *net.LinkList, opts v1.CreateOptions) (*net.LinkList, error)
	Update(ctx context.Context, linkList *net.LinkList, opts v1.UpdateOptions) (*net.LinkList, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*net.LinkList, error)
	List(ctx context.Context, opts v1.ListOptions) (*net.LinkListList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *net.LinkList, err error)
	LinkListExpansion
}

// linkLists implements LinkListInterface
type linkLists struct {
	client rest.Interface
	ns     string
}

// newLinkLists returns a LinkLists
func newLinkLists(c *NetClient, namespace string) *linkLists {
	return &linkLists{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the linkList, and returns the corresponding linkList object, and an error if there is any.
func (c *linkLists) Get(ctx context.Context, name string, options v1.GetOptions) (result *net.LinkList, err error) {
	result = &net.LinkList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("linklists").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of LinkLists that match those selectors.
func (c *linkLists) List(ctx context.Context, opts v1.ListOptions) (result *net.LinkListList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &net.LinkListList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("linklists").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested linkLists.
func (c *linkLists) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("linklists").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a linkList and creates it.  Returns the server's representation of the linkList, and an error, if there is any.
func (c *linkLists) Create(ctx context.Context, linkList *net.LinkList, opts v1.CreateOptions) (result *net.LinkList, err error) {
	result = &net.LinkList{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("linklists").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(linkList).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a linkList and updates it. Returns the server's representation of the linkList, and an error, if there is any.
func (c *linkLists) Update(ctx context.Context, linkList *net.LinkList, opts v1.UpdateOptions) (result *net.LinkList, err error) {
	result = &net.LinkList{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("linklists").
		Name(linkList.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(linkList).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the linkList and deletes it. Returns an error if one occurs.
func (c *linkLists) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("linklists").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *linkLists) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("linklists").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched linkList.
func (c *linkLists) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *net.LinkList, err error) {
	result = &net.LinkList{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("linklists").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}