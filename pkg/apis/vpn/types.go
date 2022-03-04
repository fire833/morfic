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

package vpn

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen=true

// VPNTunnelList represents a list of VPN tunnels.
type VPNTunnelList struct {
	metav1.TypeMeta `json:",inline" yaml:",inline"`

	// Standard object metadata.
	// Utilizes the Kubernetes metadata object spec for now.
	metav1.ObjectMeta `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	// Items represents an array of VPN tunnels
	Items []VPNTunnel `json:"items" yaml:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen=true

// VPNTunnel represents a tunnel to a remote endpoint with a vpn configuration.
type VPNTunnel struct {
	metav1.TypeMeta `json:",inline" yaml:",inline"`

	// Standard object metadata.
	// Utilizes the Kubernetes metadata object spec for now.
	metav1.ObjectMeta `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	Spec VPNTunnelSpec `json:"spec" yaml:"spec"`

	Status VPNTunnelStatus `json:"status" yaml:"status"`
}

type VPNTunnelSpec struct {
	//
	//
	// +optional
	Wireguard WireguardTunnelConfig `json:"wireguard" yaml:"wireguard"`

	//
	//
	// +optional
	OpenVPN OpenVPNTunnelConfig `json:"openvpn" yaml:"openvpn"`

	//
	//
	// +optional
	L2TP L2TPTunnelConfig `json:"l2tp" yaml:"l2tp"`
}

type VPNTunnelStatus struct {
}

type WireguardTunnelConfig struct {
}

type OpenVPNTunnelConfig struct {
}

type L2TPTunnelConfig struct {
}
