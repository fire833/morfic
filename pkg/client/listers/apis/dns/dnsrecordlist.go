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

// Code generated by lister-gen. DO NOT EDIT.

package dns

import (
	dns "github.com/fire833/morfic/pkg/apis/dns"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DNSRecordListLister helps list DNSRecordLists.
// All objects returned here must be treated as read-only.
type DNSRecordListLister interface {
	// List lists all DNSRecordLists in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*dns.DNSRecordList, err error)
	// Get retrieves the DNSRecordList from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*dns.DNSRecordList, error)
	DNSRecordListListerExpansion
}

// dNSRecordListLister implements the DNSRecordListLister interface.
type dNSRecordListLister struct {
	indexer cache.Indexer
}

// NewDNSRecordListLister returns a new DNSRecordListLister.
func NewDNSRecordListLister(indexer cache.Indexer) DNSRecordListLister {
	return &dNSRecordListLister{indexer: indexer}
}

// List lists all DNSRecordLists in the indexer.
func (s *dNSRecordListLister) List(selector labels.Selector) (ret []*dns.DNSRecordList, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*dns.DNSRecordList))
	})
	return ret, err
}

// Get retrieves the DNSRecordList from the index for a given name.
func (s *dNSRecordListLister) Get(name string) (*dns.DNSRecordList, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(dns.Resource("dnsrecordlist"), name)
	}
	return obj.(*dns.DNSRecordList), nil
}
