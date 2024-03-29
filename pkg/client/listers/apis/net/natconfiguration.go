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

// NATConfigurationLister helps list NATConfigurations.
// All objects returned here must be treated as read-only.
type NATConfigurationLister interface {
	// List lists all NATConfigurations in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*net.NATConfiguration, err error)
	// Get retrieves the NATConfiguration from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*net.NATConfiguration, error)
	NATConfigurationListerExpansion
}

// nATConfigurationLister implements the NATConfigurationLister interface.
type nATConfigurationLister struct {
	indexer cache.Indexer
}

// NewNATConfigurationLister returns a new NATConfigurationLister.
func NewNATConfigurationLister(indexer cache.Indexer) NATConfigurationLister {
	return &nATConfigurationLister{indexer: indexer}
}

// List lists all NATConfigurations in the indexer.
func (s *nATConfigurationLister) List(selector labels.Selector) (ret []*net.NATConfiguration, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*net.NATConfiguration))
	})
	return ret, err
}

// Get retrieves the NATConfiguration from the index for a given name.
func (s *nATConfigurationLister) Get(name string) (*net.NATConfiguration, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(net.Resource("natconfiguration"), name)
	}
	return obj.(*net.NATConfiguration), nil
}
