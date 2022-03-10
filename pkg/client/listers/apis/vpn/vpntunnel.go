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

package vpn

import (
	vpn "github.com/fire833/vroute/pkg/apis/vpn"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// VPNTunnelLister helps list VPNTunnels.
// All objects returned here must be treated as read-only.
type VPNTunnelLister interface {
	// List lists all VPNTunnels in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*vpn.VPNTunnel, err error)
	// VPNTunnels returns an object that can list and get VPNTunnels.
	VPNTunnels(namespace string) VPNTunnelNamespaceLister
	VPNTunnelListerExpansion
}

// vPNTunnelLister implements the VPNTunnelLister interface.
type vPNTunnelLister struct {
	indexer cache.Indexer
}

// NewVPNTunnelLister returns a new VPNTunnelLister.
func NewVPNTunnelLister(indexer cache.Indexer) VPNTunnelLister {
	return &vPNTunnelLister{indexer: indexer}
}

// List lists all VPNTunnels in the indexer.
func (s *vPNTunnelLister) List(selector labels.Selector) (ret []*vpn.VPNTunnel, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*vpn.VPNTunnel))
	})
	return ret, err
}

// VPNTunnels returns an object that can list and get VPNTunnels.
func (s *vPNTunnelLister) VPNTunnels(namespace string) VPNTunnelNamespaceLister {
	return vPNTunnelNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// VPNTunnelNamespaceLister helps list and get VPNTunnels.
// All objects returned here must be treated as read-only.
type VPNTunnelNamespaceLister interface {
	// List lists all VPNTunnels in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*vpn.VPNTunnel, err error)
	// Get retrieves the VPNTunnel from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*vpn.VPNTunnel, error)
	VPNTunnelNamespaceListerExpansion
}

// vPNTunnelNamespaceLister implements the VPNTunnelNamespaceLister
// interface.
type vPNTunnelNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all VPNTunnels in the indexer for a given namespace.
func (s vPNTunnelNamespaceLister) List(selector labels.Selector) (ret []*vpn.VPNTunnel, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*vpn.VPNTunnel))
	})
	return ret, err
}

// Get retrieves the VPNTunnel from the indexer for a given namespace and name.
func (s vPNTunnelNamespaceLister) Get(name string) (*vpn.VPNTunnel, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(vpn.Resource("vpntunnel"), name)
	}
	return obj.(*vpn.VPNTunnel), nil
}
