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

package routes

import (
	routes "github.com/fire833/morfic/pkg/apis/routes"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// RouteTableListLister helps list RouteTableLists.
// All objects returned here must be treated as read-only.
type RouteTableListLister interface {
	// List lists all RouteTableLists in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*routes.RouteTableList, err error)
	// RouteTableLists returns an object that can list and get RouteTableLists.
	RouteTableLists(namespace string) RouteTableListNamespaceLister
	RouteTableListListerExpansion
}

// routeTableListLister implements the RouteTableListLister interface.
type routeTableListLister struct {
	indexer cache.Indexer
}

// NewRouteTableListLister returns a new RouteTableListLister.
func NewRouteTableListLister(indexer cache.Indexer) RouteTableListLister {
	return &routeTableListLister{indexer: indexer}
}

// List lists all RouteTableLists in the indexer.
func (s *routeTableListLister) List(selector labels.Selector) (ret []*routes.RouteTableList, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*routes.RouteTableList))
	})
	return ret, err
}

// RouteTableLists returns an object that can list and get RouteTableLists.
func (s *routeTableListLister) RouteTableLists(namespace string) RouteTableListNamespaceLister {
	return routeTableListNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// RouteTableListNamespaceLister helps list and get RouteTableLists.
// All objects returned here must be treated as read-only.
type RouteTableListNamespaceLister interface {
	// List lists all RouteTableLists in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*routes.RouteTableList, err error)
	// Get retrieves the RouteTableList from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*routes.RouteTableList, error)
	RouteTableListNamespaceListerExpansion
}

// routeTableListNamespaceLister implements the RouteTableListNamespaceLister
// interface.
type routeTableListNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all RouteTableLists in the indexer for a given namespace.
func (s routeTableListNamespaceLister) List(selector labels.Selector) (ret []*routes.RouteTableList, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*routes.RouteTableList))
	})
	return ret, err
}

// Get retrieves the RouteTableList from the indexer for a given namespace and name.
func (s routeTableListNamespaceLister) Get(name string) (*routes.RouteTableList, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(routes.Resource("routetablelist"), name)
	}
	return obj.(*routes.RouteTableList), nil
}