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
// +genclient:noStatus

// AddressList represents a list of Address objects.
type AddressList struct {
	metav1.TypeMeta `json:",inline" yaml:",inline"`

	// Standard object metadata.
	// Utilizes the Kubernetes metadata object spec for now.
	metav1.ObjectMeta `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	// Items specifies the array of addresses.
	Items []Address `json:"items" yaml:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen=true
// +genclient
// +genclient:noStatus

// Address represents an address that is to be assigned to a link on a host.
type Address struct {
	metav1.TypeMeta `json:",inline" yaml:",inline"`

	// Standard object metadata.
	// Utilizes the Kubernetes metadata object spec for now.
	metav1.ObjectMeta `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	// Spec is the desired spec of this address object.
	Spec AddressSpec `json:"spec" yaml:"spec"`

	// Status is the current state of this address on the host.
	//
	// Should not be filled out by the user, will be filled/managed
	// by the server. Can be read by user at runtime.
	Status AddressStatus `json:"status" yaml:"status"`
}

// AddressSpec represents the desired spec for an address.
type AddressSpec struct {

	// Namespace specifies within which net namespace should vroute look for the link
	// to assign this address to. If this value is set to null or is undefined,
	// the control plane defaults to using the main host net namespace and looks
	// for the link to bind to in that namespace.
	Namespace string `json:"namespace" yaml:"namespace"`

	// Linkname specifies the host link that this address should be bound/assigned to.
	// If the link does not exist, then the status for this address will have
	// bound set to false.
	Linkname string `json:"linkName" yaml:"linkName"`

	// IP Specifies the IPSpec of the address for this object.
	IP IPSpec `json:"ip" yaml:"ip"`
}

// AddressStatus specifies the status of an address.
type AddressStatus struct {

	// Bound notes if this address has been bound to a link on the host.
	Bound bool `json:"isBound" yaml:"isBound"`
}

// IPSpec represents an IP address that is either IPv4 or IPv6.
type IPSpec struct {

	// Type specifies whether this address is IPv4 or IPv6. Can be set to either "IPv4"
	// or to "IPv6".
	Type string `json:"type" yaml:"type"`

	// Address specifies the actual IP address for the IPSpec. Should be in string
	// format and parseable by net.ParseIP().
	Address string `json:"address" yaml:"address"`
}
