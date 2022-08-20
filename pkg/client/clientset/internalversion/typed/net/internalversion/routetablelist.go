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

// RouteTableListsGetter has a method to return a RouteTableListInterface.
// A group's client should implement this interface.
type RouteTableListsGetter interface {
	RouteTableLists(namespace string) RouteTableListInterface
}

// RouteTableListInterface has methods to work with RouteTableList resources.
type RouteTableListInterface interface {
	Create(ctx context.Context, routeTableList *net.RouteTableList, opts v1.CreateOptions) (*net.RouteTableList, error)
	Update(ctx context.Context, routeTableList *net.RouteTableList, opts v1.UpdateOptions) (*net.RouteTableList, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*net.RouteTableList, error)
	List(ctx context.Context, opts v1.ListOptions) (*net.RouteTableListList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *net.RouteTableList, err error)
	RouteTableListExpansion
}

// routeTableLists implements RouteTableListInterface
type routeTableLists struct {
	client rest.Interface
	ns     string
}

// newRouteTableLists returns a RouteTableLists
func newRouteTableLists(c *NetClient, namespace string) *routeTableLists {
	return &routeTableLists{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the routeTableList, and returns the corresponding routeTableList object, and an error if there is any.
func (c *routeTableLists) Get(ctx context.Context, name string, options v1.GetOptions) (result *net.RouteTableList, err error) {
	result = &net.RouteTableList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("routetablelists").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of RouteTableLists that match those selectors.
func (c *routeTableLists) List(ctx context.Context, opts v1.ListOptions) (result *net.RouteTableListList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &net.RouteTableListList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("routetablelists").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested routeTableLists.
func (c *routeTableLists) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("routetablelists").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a routeTableList and creates it.  Returns the server's representation of the routeTableList, and an error, if there is any.
func (c *routeTableLists) Create(ctx context.Context, routeTableList *net.RouteTableList, opts v1.CreateOptions) (result *net.RouteTableList, err error) {
	result = &net.RouteTableList{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("routetablelists").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(routeTableList).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a routeTableList and updates it. Returns the server's representation of the routeTableList, and an error, if there is any.
func (c *routeTableLists) Update(ctx context.Context, routeTableList *net.RouteTableList, opts v1.UpdateOptions) (result *net.RouteTableList, err error) {
	result = &net.RouteTableList{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("routetablelists").
		Name(routeTableList.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(routeTableList).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the routeTableList and deletes it. Returns an error if one occurs.
func (c *routeTableLists) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("routetablelists").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *routeTableLists) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("routetablelists").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched routeTableList.
func (c *routeTableLists) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *net.RouteTableList, err error) {
	result = &net.RouteTableList{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("routetablelists").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}