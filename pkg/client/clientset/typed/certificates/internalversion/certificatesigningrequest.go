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

	certificates "github.com/fire833/morfic/pkg/apis/certificates"
	scheme "github.com/fire833/morfic/pkg/client/clientset/internalversion/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CertificateSigningRequestsGetter has a method to return a CertificateSigningRequestInterface.
// A group's client should implement this interface.
type CertificateSigningRequestsGetter interface {
	CertificateSigningRequests(namespace string) CertificateSigningRequestInterface
}

// CertificateSigningRequestInterface has methods to work with CertificateSigningRequest resources.
type CertificateSigningRequestInterface interface {
	Create(ctx context.Context, certificateSigningRequest *certificates.CertificateSigningRequest, opts v1.CreateOptions) (*certificates.CertificateSigningRequest, error)
	Update(ctx context.Context, certificateSigningRequest *certificates.CertificateSigningRequest, opts v1.UpdateOptions) (*certificates.CertificateSigningRequest, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*certificates.CertificateSigningRequest, error)
	List(ctx context.Context, opts v1.ListOptions) (*certificates.CertificateSigningRequestList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *certificates.CertificateSigningRequest, err error)
	CertificateSigningRequestExpansion
}

// certificateSigningRequests implements CertificateSigningRequestInterface
type certificateSigningRequests struct {
	client rest.Interface
	ns     string
}

// newCertificateSigningRequests returns a CertificateSigningRequests
func newCertificateSigningRequests(c *CertificatesClient, namespace string) *certificateSigningRequests {
	return &certificateSigningRequests{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the certificateSigningRequest, and returns the corresponding certificateSigningRequest object, and an error if there is any.
func (c *certificateSigningRequests) Get(ctx context.Context, name string, options v1.GetOptions) (result *certificates.CertificateSigningRequest, err error) {
	result = &certificates.CertificateSigningRequest{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("certificatesigningrequests").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CertificateSigningRequests that match those selectors.
func (c *certificateSigningRequests) List(ctx context.Context, opts v1.ListOptions) (result *certificates.CertificateSigningRequestList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &certificates.CertificateSigningRequestList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("certificatesigningrequests").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested certificateSigningRequests.
func (c *certificateSigningRequests) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("certificatesigningrequests").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a certificateSigningRequest and creates it.  Returns the server's representation of the certificateSigningRequest, and an error, if there is any.
func (c *certificateSigningRequests) Create(ctx context.Context, certificateSigningRequest *certificates.CertificateSigningRequest, opts v1.CreateOptions) (result *certificates.CertificateSigningRequest, err error) {
	result = &certificates.CertificateSigningRequest{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("certificatesigningrequests").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(certificateSigningRequest).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a certificateSigningRequest and updates it. Returns the server's representation of the certificateSigningRequest, and an error, if there is any.
func (c *certificateSigningRequests) Update(ctx context.Context, certificateSigningRequest *certificates.CertificateSigningRequest, opts v1.UpdateOptions) (result *certificates.CertificateSigningRequest, err error) {
	result = &certificates.CertificateSigningRequest{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("certificatesigningrequests").
		Name(certificateSigningRequest.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(certificateSigningRequest).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the certificateSigningRequest and deletes it. Returns an error if one occurs.
func (c *certificateSigningRequests) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("certificatesigningrequests").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *certificateSigningRequests) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("certificatesigningrequests").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched certificateSigningRequest.
func (c *certificateSigningRequests) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *certificates.CertificateSigningRequest, err error) {
	result = &certificates.CertificateSigningRequest{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("certificatesigningrequests").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
