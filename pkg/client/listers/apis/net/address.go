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

package net

import (
	net "github.com/fire833/morfic/pkg/apis/net"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AddressLister helps list Addresses.
// All objects returned here must be treated as read-only.
type AddressLister interface {
	// List lists all Addresses in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*net.Address, err error)
	// Addresses returns an object that can list and get Addresses.
	Addresses(namespace string) AddressNamespaceLister
	AddressListerExpansion
}

// addressLister implements the AddressLister interface.
type addressLister struct {
	indexer cache.Indexer
}

// NewAddressLister returns a new AddressLister.
func NewAddressLister(indexer cache.Indexer) AddressLister {
	return &addressLister{indexer: indexer}
}

// List lists all Addresses in the indexer.
func (s *addressLister) List(selector labels.Selector) (ret []*net.Address, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*net.Address))
	})
	return ret, err
}

// Addresses returns an object that can list and get Addresses.
func (s *addressLister) Addresses(namespace string) AddressNamespaceLister {
	return addressNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AddressNamespaceLister helps list and get Addresses.
// All objects returned here must be treated as read-only.
type AddressNamespaceLister interface {
	// List lists all Addresses in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*net.Address, err error)
	// Get retrieves the Address from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*net.Address, error)
	AddressNamespaceListerExpansion
}

// addressNamespaceLister implements the AddressNamespaceLister
// interface.
type addressNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Addresses in the indexer for a given namespace.
func (s addressNamespaceLister) List(selector labels.Selector) (ret []*net.Address, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*net.Address))
	})
	return ret, err
}

// Get retrieves the Address from the indexer for a given namespace and name.
func (s addressNamespaceLister) Get(name string) (*net.Address, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(net.Resource("address"), name)
	}
	return obj.(*net.Address), nil
}
