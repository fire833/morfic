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
	internalversion "github.com/fire833/morfic/pkg/client/clientset/internalversion/typed/dns/internalversion"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeDns struct {
	*testing.Fake
}

func (c *FakeDns) DNSProviders(namespace string) internalversion.DNSProviderInterface {
	return &FakeDNSProviders{c, namespace}
}

func (c *FakeDns) DNSProviderLists(namespace string) internalversion.DNSProviderListInterface {
	return &FakeDNSProviderLists{c, namespace}
}

func (c *FakeDns) DNSRecords(namespace string) internalversion.DNSRecordInterface {
	return &FakeDNSRecords{c, namespace}
}

func (c *FakeDns) DNSRecordLists(namespace string) internalversion.DNSRecordListInterface {
	return &FakeDNSRecordLists{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeDns) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
