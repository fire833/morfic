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

// RouteLister helps list Routes.
// All objects returned here must be treated as read-only.
type RouteLister interface {
	// List lists all Routes in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*net.Route, err error)
	// Get retrieves the Route from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*net.Route, error)
	RouteListerExpansion
}

// routeLister implements the RouteLister interface.
type routeLister struct {
	indexer cache.Indexer
}

// NewRouteLister returns a new RouteLister.
func NewRouteLister(indexer cache.Indexer) RouteLister {
	return &routeLister{indexer: indexer}
}

// List lists all Routes in the indexer.
func (s *routeLister) List(selector labels.Selector) (ret []*net.Route, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*net.Route))
	})
	return ret, err
}

// Get retrieves the Route from the index for a given name.
func (s *routeLister) Get(name string) (*net.Route, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(net.Resource("route"), name)
	}
	return obj.(*net.Route), nil
}
