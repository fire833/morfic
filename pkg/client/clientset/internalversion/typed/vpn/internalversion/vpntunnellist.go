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

	vpn "github.com/fire833/morfic/pkg/apis/vpn"
	scheme "github.com/fire833/morfic/pkg/client/clientset/internalversion/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// VPNTunnelListsGetter has a method to return a VPNTunnelListInterface.
// A group's client should implement this interface.
type VPNTunnelListsGetter interface {
	VPNTunnelLists(namespace string) VPNTunnelListInterface
}

// VPNTunnelListInterface has methods to work with VPNTunnelList resources.
type VPNTunnelListInterface interface {
	Create(ctx context.Context, vPNTunnelList *vpn.VPNTunnelList, opts v1.CreateOptions) (*vpn.VPNTunnelList, error)
	Update(ctx context.Context, vPNTunnelList *vpn.VPNTunnelList, opts v1.UpdateOptions) (*vpn.VPNTunnelList, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*vpn.VPNTunnelList, error)
	List(ctx context.Context, opts v1.ListOptions) (*vpn.VPNTunnelListList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *vpn.VPNTunnelList, err error)
	VPNTunnelListExpansion
}

// vPNTunnelLists implements VPNTunnelListInterface
type vPNTunnelLists struct {
	client rest.Interface
	ns     string
}

// newVPNTunnelLists returns a VPNTunnelLists
func newVPNTunnelLists(c *VpnClient, namespace string) *vPNTunnelLists {
	return &vPNTunnelLists{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the vPNTunnelList, and returns the corresponding vPNTunnelList object, and an error if there is any.
func (c *vPNTunnelLists) Get(ctx context.Context, name string, options v1.GetOptions) (result *vpn.VPNTunnelList, err error) {
	result = &vpn.VPNTunnelList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("vpntunnellists").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of VPNTunnelLists that match those selectors.
func (c *vPNTunnelLists) List(ctx context.Context, opts v1.ListOptions) (result *vpn.VPNTunnelListList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &vpn.VPNTunnelListList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("vpntunnellists").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested vPNTunnelLists.
func (c *vPNTunnelLists) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("vpntunnellists").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a vPNTunnelList and creates it.  Returns the server's representation of the vPNTunnelList, and an error, if there is any.
func (c *vPNTunnelLists) Create(ctx context.Context, vPNTunnelList *vpn.VPNTunnelList, opts v1.CreateOptions) (result *vpn.VPNTunnelList, err error) {
	result = &vpn.VPNTunnelList{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("vpntunnellists").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(vPNTunnelList).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a vPNTunnelList and updates it. Returns the server's representation of the vPNTunnelList, and an error, if there is any.
func (c *vPNTunnelLists) Update(ctx context.Context, vPNTunnelList *vpn.VPNTunnelList, opts v1.UpdateOptions) (result *vpn.VPNTunnelList, err error) {
	result = &vpn.VPNTunnelList{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("vpntunnellists").
		Name(vPNTunnelList.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(vPNTunnelList).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the vPNTunnelList and deletes it. Returns an error if one occurs.
func (c *vPNTunnelLists) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("vpntunnellists").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *vPNTunnelLists) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("vpntunnellists").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched vPNTunnelList.
func (c *vPNTunnelLists) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *vpn.VPNTunnelList, err error) {
	result = &vpn.VPNTunnelList{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("vpntunnellists").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
