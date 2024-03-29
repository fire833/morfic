/*
*	Copyright (C) 2023 Kendall Tauser
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

// Code generated by lister-gen. DO NOT EDIT.

package dns

import (
	dns "github.com/fire833/morfic/pkg/apis/dns"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DNSProviderLister helps list DNSProviders.
// All objects returned here must be treated as read-only.
type DNSProviderLister interface {
	// List lists all DNSProviders in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*dns.DNSProvider, err error)
	// Get retrieves the DNSProvider from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*dns.DNSProvider, error)
	DNSProviderListerExpansion
}

// dNSProviderLister implements the DNSProviderLister interface.
type dNSProviderLister struct {
	indexer cache.Indexer
}

// NewDNSProviderLister returns a new DNSProviderLister.
func NewDNSProviderLister(indexer cache.Indexer) DNSProviderLister {
	return &dNSProviderLister{indexer: indexer}
}

// List lists all DNSProviders in the indexer.
func (s *dNSProviderLister) List(selector labels.Selector) (ret []*dns.DNSProvider, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*dns.DNSProvider))
	})
	return ret, err
}

// Get retrieves the DNSProvider from the index for a given name.
func (s *dNSProviderLister) Get(name string) (*dns.DNSProvider, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(dns.Resource("dnsprovider"), name)
	}
	return obj.(*dns.DNSProvider), nil
}
