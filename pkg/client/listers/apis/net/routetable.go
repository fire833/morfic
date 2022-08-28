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

package net

import (
	net "github.com/fire833/morfic/pkg/apis/net"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// RouteTableLister helps list RouteTables.
// All objects returned here must be treated as read-only.
type RouteTableLister interface {
	// List lists all RouteTables in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*net.RouteTable, err error)
	// Get retrieves the RouteTable from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*net.RouteTable, error)
	RouteTableListerExpansion
}

// routeTableLister implements the RouteTableLister interface.
type routeTableLister struct {
	indexer cache.Indexer
}

// NewRouteTableLister returns a new RouteTableLister.
func NewRouteTableLister(indexer cache.Indexer) RouteTableLister {
	return &routeTableLister{indexer: indexer}
}

// List lists all RouteTables in the indexer.
func (s *routeTableLister) List(selector labels.Selector) (ret []*net.RouteTable, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*net.RouteTable))
	})
	return ret, err
}

// Get retrieves the RouteTable from the index for a given name.
func (s *routeTableLister) Get(name string) (*net.RouteTable, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(net.Resource("routetable"), name)
	}
	return obj.(*net.RouteTable), nil
}
