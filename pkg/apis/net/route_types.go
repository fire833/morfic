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

package net

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen=true
// +genclient
// +genclient:nonNamespaced
// RouteTableList specifies a list of RouteTables.
type RouteTableList struct {
	metav1.TypeMeta `json:",inline" yaml:",inline"`

	// +optional
	metav1.ListMeta

	// Items specifies the array of RouteTables.
	Items []RouteTable `json:"items" yaml:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen=true
// +genclient
// +genclient:nonNamespaced
// RouteTable specifies a routing table located in the kernel.
type RouteTable struct {
	metav1.TypeMeta `json:",inline" yaml:",inline"`

	// Standard object metadata.
	// Utilizes the Kubernetes metadata object spec for now.
	metav1.ObjectMeta `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	// Spec is the desired spec of this route table.
	Spec RouteTableSpec `json:"spec" yaml:"spec"`

	// Status is the current state of this route table on the host.
	//
	// Should not be filled out by the user, will be filled/managed
	// by the server. Can be read by user at runtime.
	Status RouteTableStatus `json:"status" yaml:"status"`
}

// RouteTableSpec specifies the desired state of the route table.
type RouteTableSpec struct {
	// Name specifies the name of this routing table.
	Name string `json:"name" yaml:"name"`

	// Index specifies the index of this route table.
	Index uint8 `json:"index" yaml:"index"`
}

// RouteTableStatus specifies the current status of the routing table.
type RouteTableStatus struct {
	// FoundOnBoot specifies if this route table was found at bootup.
	FoundOnBoot bool `json:"foundOnBoot" yaml:"foundOnBoot"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen=true
// +genclient
// +genclient:nonNamespaced
// Route specifies a route on the host.
type Route struct {
	metav1.TypeMeta `json:",inline" yaml:",inline"`

	// Standard object metadata.
	// Utilizes the Kubernetes metadata object spec for now.
	metav1.ObjectMeta `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	// Spec is the desired spec of this route.
	Spec RouteSpec `json:"spec" yaml:"spec"`

	// Status is the current state of this route on the host.
	//
	// Should not be filled out by the user, will be filled/managed
	// by the server. Can be read by user at runtime.
	Status RouteStatus `json:"status" yaml:"status"`
}

type RouteSpec struct {
	Type string `json:"type" yaml:"type"`
}

type RouteStatus struct {
}
