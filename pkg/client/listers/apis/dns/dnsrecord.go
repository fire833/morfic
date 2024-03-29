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

// DNSRecordLister helps list DNSRecords.
// All objects returned here must be treated as read-only.
type DNSRecordLister interface {
	// List lists all DNSRecords in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*dns.DNSRecord, err error)
	// Get retrieves the DNSRecord from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*dns.DNSRecord, error)
	DNSRecordListerExpansion
}

// dNSRecordLister implements the DNSRecordLister interface.
type dNSRecordLister struct {
	indexer cache.Indexer
}

// NewDNSRecordLister returns a new DNSRecordLister.
func NewDNSRecordLister(indexer cache.Indexer) DNSRecordLister {
	return &dNSRecordLister{indexer: indexer}
}

// List lists all DNSRecords in the indexer.
func (s *dNSRecordLister) List(selector labels.Selector) (ret []*dns.DNSRecord, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*dns.DNSRecord))
	})
	return ret, err
}

// Get retrieves the DNSRecord from the index for a given name.
func (s *dNSRecordLister) Get(name string) (*dns.DNSRecord, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(dns.Resource("dnsrecord"), name)
	}
	return obj.(*dns.DNSRecord), nil
}
