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
	internalversion "github.com/fire833/morfic/pkg/client/clientset/internalversion/typed/certificates/internalversion"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeCertificates struct {
	*testing.Fake
}

func (c *FakeCertificates) Certificates(namespace string) internalversion.CertificateInterface {
	return &FakeCertificates{c, namespace}
}

func (c *FakeCertificates) CertificateLists(namespace string) internalversion.CertificateListInterface {
	return &FakeCertificateLists{c, namespace}
}

func (c *FakeCertificates) CertificateSigners(namespace string) internalversion.CertificateSignerInterface {
	return &FakeCertificateSigners{c, namespace}
}

func (c *FakeCertificates) CertificateSignerLists(namespace string) internalversion.CertificateSignerListInterface {
	return &FakeCertificateSignerLists{c, namespace}
}

func (c *FakeCertificates) CertificateSigningRequests(namespace string) internalversion.CertificateSigningRequestInterface {
	return &FakeCertificateSigningRequests{c, namespace}
}

func (c *FakeCertificates) CertificateSigningRequestLists(namespace string) internalversion.CertificateSigningRequestListInterface {
	return &FakeCertificateSigningRequestLists{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeCertificates) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
