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

	routes "github.com/fire833/morfic/pkg/apis/routes"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRouteTables implements RouteTableInterface
type FakeRouteTables struct {
	Fake *FakeRoutes
	ns   string
}

var routetablesResource = schema.GroupVersionResource{Group: "routes.morfic.io", Version: "", Resource: "routetables"}

var routetablesKind = schema.GroupVersionKind{Group: "routes.morfic.io", Version: "", Kind: "RouteTable"}

// Get takes name of the routeTable, and returns the corresponding routeTable object, and an error if there is any.
func (c *FakeRouteTables) Get(ctx context.Context, name string, options v1.GetOptions) (result *routes.RouteTable, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(routetablesResource, c.ns, name), &routes.RouteTable{})

	if obj == nil {
		return nil, err
	}
	return obj.(*routes.RouteTable), err
}

// List takes label and field selectors, and returns the list of RouteTables that match those selectors.
func (c *FakeRouteTables) List(ctx context.Context, opts v1.ListOptions) (result *routes.RouteTableList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(routetablesResource, routetablesKind, c.ns, opts), &routes.RouteTableList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &routes.RouteTableList{ListMeta: obj.(*routes.RouteTableList).ListMeta}
	for _, item := range obj.(*routes.RouteTableList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested routeTables.
func (c *FakeRouteTables) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(routetablesResource, c.ns, opts))

}

// Create takes the representation of a routeTable and creates it.  Returns the server's representation of the routeTable, and an error, if there is any.
func (c *FakeRouteTables) Create(ctx context.Context, routeTable *routes.RouteTable, opts v1.CreateOptions) (result *routes.RouteTable, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(routetablesResource, c.ns, routeTable), &routes.RouteTable{})

	if obj == nil {
		return nil, err
	}
	return obj.(*routes.RouteTable), err
}

// Update takes the representation of a routeTable and updates it. Returns the server's representation of the routeTable, and an error, if there is any.
func (c *FakeRouteTables) Update(ctx context.Context, routeTable *routes.RouteTable, opts v1.UpdateOptions) (result *routes.RouteTable, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(routetablesResource, c.ns, routeTable), &routes.RouteTable{})

	if obj == nil {
		return nil, err
	}
	return obj.(*routes.RouteTable), err
}

// Delete takes name of the routeTable and deletes it. Returns an error if one occurs.
func (c *FakeRouteTables) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(routetablesResource, c.ns, name, opts), &routes.RouteTable{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRouteTables) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(routetablesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &routes.RouteTableList{})
	return err
}

// Patch applies the patch and returns the patched routeTable.
func (c *FakeRouteTables) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *routes.RouteTable, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(routetablesResource, c.ns, name, pt, data, subresources...), &routes.RouteTable{})

	if obj == nil {
		return nil, err
	}
	return obj.(*routes.RouteTable), err
}
