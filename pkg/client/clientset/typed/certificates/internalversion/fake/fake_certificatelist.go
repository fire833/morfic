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

	certificates "github.com/fire833/morfic/pkg/apis/certificates"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCertificateLists implements CertificateListInterface
type FakeCertificateLists struct {
	Fake *FakeCertificates
	ns   string
}

var certificatelistsResource = schema.GroupVersionResource{Group: "certificates.morfic.io", Version: "", Resource: "certificatelists"}

var certificatelistsKind = schema.GroupVersionKind{Group: "certificates.morfic.io", Version: "", Kind: "CertificateList"}

// Get takes name of the certificateList, and returns the corresponding certificateList object, and an error if there is any.
func (c *FakeCertificateLists) Get(ctx context.Context, name string, options v1.GetOptions) (result *certificates.CertificateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(certificatelistsResource, c.ns, name), &certificates.CertificateList{})

	if obj == nil {
		return nil, err
	}
	return obj.(*certificates.CertificateList), err
}

// List takes label and field selectors, and returns the list of CertificateLists that match those selectors.
func (c *FakeCertificateLists) List(ctx context.Context, opts v1.ListOptions) (result *certificates.CertificateListList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(certificatelistsResource, certificatelistsKind, c.ns, opts), &certificates.CertificateListList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &certificates.CertificateListList{ListMeta: obj.(*certificates.CertificateListList).ListMeta}
	for _, item := range obj.(*certificates.CertificateListList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested certificateLists.
func (c *FakeCertificateLists) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(certificatelistsResource, c.ns, opts))

}

// Create takes the representation of a certificateList and creates it.  Returns the server's representation of the certificateList, and an error, if there is any.
func (c *FakeCertificateLists) Create(ctx context.Context, certificateList *certificates.CertificateList, opts v1.CreateOptions) (result *certificates.CertificateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(certificatelistsResource, c.ns, certificateList), &certificates.CertificateList{})

	if obj == nil {
		return nil, err
	}
	return obj.(*certificates.CertificateList), err
}

// Update takes the representation of a certificateList and updates it. Returns the server's representation of the certificateList, and an error, if there is any.
func (c *FakeCertificateLists) Update(ctx context.Context, certificateList *certificates.CertificateList, opts v1.UpdateOptions) (result *certificates.CertificateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(certificatelistsResource, c.ns, certificateList), &certificates.CertificateList{})

	if obj == nil {
		return nil, err
	}
	return obj.(*certificates.CertificateList), err
}

// Delete takes name of the certificateList and deletes it. Returns an error if one occurs.
func (c *FakeCertificateLists) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(certificatelistsResource, c.ns, name, opts), &certificates.CertificateList{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCertificateLists) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(certificatelistsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &certificates.CertificateListList{})
	return err
}

// Patch applies the patch and returns the patched certificateList.
func (c *FakeCertificateLists) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *certificates.CertificateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(certificatelistsResource, c.ns, name, pt, data, subresources...), &certificates.CertificateList{})

	if obj == nil {
		return nil, err
	}
	return obj.(*certificates.CertificateList), err
}
