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

	neighbors "github.com/fire833/morfic/pkg/apis/neighbors"
	scheme "github.com/fire833/morfic/pkg/client/clientset/internalversion/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// NeighborListsGetter has a method to return a NeighborListInterface.
// A group's client should implement this interface.
type NeighborListsGetter interface {
	NeighborLists(namespace string) NeighborListInterface
}

// NeighborListInterface has methods to work with NeighborList resources.
type NeighborListInterface interface {
	Create(ctx context.Context, neighborList *neighbors.NeighborList, opts v1.CreateOptions) (*neighbors.NeighborList, error)
	Update(ctx context.Context, neighborList *neighbors.NeighborList, opts v1.UpdateOptions) (*neighbors.NeighborList, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*neighbors.NeighborList, error)
	List(ctx context.Context, opts v1.ListOptions) (*neighbors.NeighborListList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *neighbors.NeighborList, err error)
	NeighborListExpansion
}

// neighborLists implements NeighborListInterface
type neighborLists struct {
	client rest.Interface
	ns     string
}

// newNeighborLists returns a NeighborLists
func newNeighborLists(c *NeighborsClient, namespace string) *neighborLists {
	return &neighborLists{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the neighborList, and returns the corresponding neighborList object, and an error if there is any.
func (c *neighborLists) Get(ctx context.Context, name string, options v1.GetOptions) (result *neighbors.NeighborList, err error) {
	result = &neighbors.NeighborList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("neighborlists").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of NeighborLists that match those selectors.
func (c *neighborLists) List(ctx context.Context, opts v1.ListOptions) (result *neighbors.NeighborListList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &neighbors.NeighborListList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("neighborlists").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested neighborLists.
func (c *neighborLists) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("neighborlists").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a neighborList and creates it.  Returns the server's representation of the neighborList, and an error, if there is any.
func (c *neighborLists) Create(ctx context.Context, neighborList *neighbors.NeighborList, opts v1.CreateOptions) (result *neighbors.NeighborList, err error) {
	result = &neighbors.NeighborList{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("neighborlists").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(neighborList).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a neighborList and updates it. Returns the server's representation of the neighborList, and an error, if there is any.
func (c *neighborLists) Update(ctx context.Context, neighborList *neighbors.NeighborList, opts v1.UpdateOptions) (result *neighbors.NeighborList, err error) {
	result = &neighbors.NeighborList{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("neighborlists").
		Name(neighborList.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(neighborList).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the neighborList and deletes it. Returns an error if one occurs.
func (c *neighborLists) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("neighborlists").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *neighborLists) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("neighborlists").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched neighborList.
func (c *neighborLists) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *neighbors.NeighborList, err error) {
	result = &neighbors.NeighborList{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("neighborlists").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
