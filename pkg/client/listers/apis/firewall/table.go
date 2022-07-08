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

package firewall

import (
	firewall "github.com/fire833/morfic/pkg/apis/firewall"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// TableLister helps list Tables.
// All objects returned here must be treated as read-only.
type TableLister interface {
	// List lists all Tables in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*firewall.Table, err error)
	// Tables returns an object that can list and get Tables.
	Tables(namespace string) TableNamespaceLister
	TableListerExpansion
}

// tableLister implements the TableLister interface.
type tableLister struct {
	indexer cache.Indexer
}

// NewTableLister returns a new TableLister.
func NewTableLister(indexer cache.Indexer) TableLister {
	return &tableLister{indexer: indexer}
}

// List lists all Tables in the indexer.
func (s *tableLister) List(selector labels.Selector) (ret []*firewall.Table, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*firewall.Table))
	})
	return ret, err
}

// Tables returns an object that can list and get Tables.
func (s *tableLister) Tables(namespace string) TableNamespaceLister {
	return tableNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TableNamespaceLister helps list and get Tables.
// All objects returned here must be treated as read-only.
type TableNamespaceLister interface {
	// List lists all Tables in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*firewall.Table, err error)
	// Get retrieves the Table from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*firewall.Table, error)
	TableNamespaceListerExpansion
}

// tableNamespaceLister implements the TableNamespaceLister
// interface.
type tableNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Tables in the indexer for a given namespace.
func (s tableNamespaceLister) List(selector labels.Selector) (ret []*firewall.Table, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*firewall.Table))
	})
	return ret, err
}

// Get retrieves the Table from the indexer for a given namespace and name.
func (s tableNamespaceLister) Get(name string) (*firewall.Table, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(firewall.Resource("table"), name)
	}
	return obj.(*firewall.Table), nil
}