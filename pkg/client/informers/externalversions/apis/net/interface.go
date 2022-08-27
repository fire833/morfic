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

// Code generated by informer-gen. DO NOT EDIT.

package net

import (
	internalinterfaces "pkg/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Addresses returns a AddressInformer.
	Addresses() AddressInformer
	// AddressLists returns a AddressListInformer.
	AddressLists() AddressListInformer
	// Links returns a LinkInformer.
	Links() LinkInformer
	// LinkLists returns a LinkListInformer.
	LinkLists() LinkListInformer
	// Neighbors returns a NeighborInformer.
	Neighbors() NeighborInformer
	// NeighborLists returns a NeighborListInformer.
	NeighborLists() NeighborListInformer
	// Routes returns a RouteInformer.
	Routes() RouteInformer
	// RouteTables returns a RouteTableInformer.
	RouteTables() RouteTableInformer
	// RouteTableLists returns a RouteTableListInformer.
	RouteTableLists() RouteTableListInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// Addresses returns a AddressInformer.
func (v *version) Addresses() AddressInformer {
	return &addressInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// AddressLists returns a AddressListInformer.
func (v *version) AddressLists() AddressListInformer {
	return &addressListInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Links returns a LinkInformer.
func (v *version) Links() LinkInformer {
	return &linkInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// LinkLists returns a LinkListInformer.
func (v *version) LinkLists() LinkListInformer {
	return &linkListInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Neighbors returns a NeighborInformer.
func (v *version) Neighbors() NeighborInformer {
	return &neighborInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// NeighborLists returns a NeighborListInformer.
func (v *version) NeighborLists() NeighborListInformer {
	return &neighborListInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Routes returns a RouteInformer.
func (v *version) Routes() RouteInformer {
	return &routeInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// RouteTables returns a RouteTableInformer.
func (v *version) RouteTables() RouteTableInformer {
	return &routeTableInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// RouteTableLists returns a RouteTableListInformer.
func (v *version) RouteTableLists() RouteTableListInformer {
	return &routeTableListInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
