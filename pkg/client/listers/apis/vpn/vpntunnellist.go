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

package vpn

import (
	vpn "github.com/fire833/morfic/pkg/apis/vpn"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// VPNTunnelListLister helps list VPNTunnelLists.
// All objects returned here must be treated as read-only.
type VPNTunnelListLister interface {
	// List lists all VPNTunnelLists in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*vpn.VPNTunnelList, err error)
	// Get retrieves the VPNTunnelList from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*vpn.VPNTunnelList, error)
	VPNTunnelListListerExpansion
}

// vPNTunnelListLister implements the VPNTunnelListLister interface.
type vPNTunnelListLister struct {
	indexer cache.Indexer
}

// NewVPNTunnelListLister returns a new VPNTunnelListLister.
func NewVPNTunnelListLister(indexer cache.Indexer) VPNTunnelListLister {
	return &vPNTunnelListLister{indexer: indexer}
}

// List lists all VPNTunnelLists in the indexer.
func (s *vPNTunnelListLister) List(selector labels.Selector) (ret []*vpn.VPNTunnelList, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*vpn.VPNTunnelList))
	})
	return ret, err
}

// Get retrieves the VPNTunnelList from the index for a given name.
func (s *vPNTunnelListLister) Get(name string) (*vpn.VPNTunnelList, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(vpn.Resource("vpntunnellist"), name)
	}
	return obj.(*vpn.VPNTunnelList), nil
}
