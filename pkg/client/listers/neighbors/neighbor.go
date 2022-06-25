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

package neighbors

import (
	neighbors "github.com/fire833/morfic/pkg/apis/neighbors"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// NeighborLister helps list Neighbors.
// All objects returned here must be treated as read-only.
type NeighborLister interface {
	// List lists all Neighbors in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*neighbors.Neighbor, err error)
	// Neighbors returns an object that can list and get Neighbors.
	Neighbors(namespace string) NeighborNamespaceLister
	NeighborListerExpansion
}

// neighborLister implements the NeighborLister interface.
type neighborLister struct {
	indexer cache.Indexer
}

// NewNeighborLister returns a new NeighborLister.
func NewNeighborLister(indexer cache.Indexer) NeighborLister {
	return &neighborLister{indexer: indexer}
}

// List lists all Neighbors in the indexer.
func (s *neighborLister) List(selector labels.Selector) (ret []*neighbors.Neighbor, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*neighbors.Neighbor))
	})
	return ret, err
}

// Neighbors returns an object that can list and get Neighbors.
func (s *neighborLister) Neighbors(namespace string) NeighborNamespaceLister {
	return neighborNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// NeighborNamespaceLister helps list and get Neighbors.
// All objects returned here must be treated as read-only.
type NeighborNamespaceLister interface {
	// List lists all Neighbors in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*neighbors.Neighbor, err error)
	// Get retrieves the Neighbor from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*neighbors.Neighbor, error)
	NeighborNamespaceListerExpansion
}

// neighborNamespaceLister implements the NeighborNamespaceLister
// interface.
type neighborNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Neighbors in the indexer for a given namespace.
func (s neighborNamespaceLister) List(selector labels.Selector) (ret []*neighbors.Neighbor, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*neighbors.Neighbor))
	})
	return ret, err
}

// Get retrieves the Neighbor from the indexer for a given namespace and name.
func (s neighborNamespaceLister) Get(name string) (*neighbors.Neighbor, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(neighbors.Resource("neighbor"), name)
	}
	return obj.(*neighbors.Neighbor), nil
}
