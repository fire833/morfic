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

	core "k8s.io/api/core/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen=true
// +genclient:nonNamespaced
// +genclient
// VPNTunnelList represents a list of VPN tunnels.
type VPNTunnelList struct {
	metav1.TypeMeta `json:",inline" yaml:",inline"`

	// +optional
	metav1.ListMeta

	// Items represents an array of VPN tunnels
	Items []VPNTunnel `json:"items" yaml:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen=true
// +genclient:nonNamespaced
// +genclient
// VPNTunnel represents a tunnel to a remote endpoint with a vpn configuration.
type VPNTunnel struct {
	metav1.TypeMeta `json:",inline" yaml:",inline"`

	// Standard object metadata.
	// Utilizes the Kubernetes metadata object spec for now.
	metav1.ObjectMeta `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	Spec VPNTunnelSpec `json:"spec" yaml:"spec"`

	Status VPNTunnelStatus `json:"status" yaml:"status"`
}

// VPNTunnelSpec represents the desired specification for a VPN tunnel to a remote host.
// Only one tunnel configuration should be specified per spec object. Otherwise a conflict will occur.
type VPNTunnelSpec struct {

	// Wireguard is optional configuration for a wireguard tunnel that will be represented by this spec.
	//
	// +optional
	Wireguard WireguardTunnelConfig `json:"wireguard" yaml:"wireguard"`

	// OpenVPN is optional configuration for an OpenVPN tunnel that will be represented by this spec.
	// NOTE: Requires the openvpn service to be installed on the host and running.
	//
	// +optional
	OpenVPN OpenVPNTunnelConfig `json:"openvpn" yaml:"openvpn"`

	// L2TP is optional configuration for an L2TP tunnel that will be represented by this spec.
	//
	// +optional
	L2TP L2TPTunnelConfig `json:"l2tp" yaml:"l2tp"`

	// IPSEC is optional configuration for an IPSec tunnel that will be represented by this spec.
	//
	// +optional
	IPSEC IPSecTunnelConfig `json:"ipsec" yaml:"ipsec"`
}

// VPNTunnelStatus represents the current status for a VPN tunnel to a remote host.
type VPNTunnelStatus struct {

	// IsUp denotes whether the tunnel has been established with the remote peer.
	IsUp bool `json:"up" yaml:"up"`

	// ReceiveBytes represents the number of bytes that have been received by this tunnel from a peer.
	ReceiveBytes int64 `json:"receiveBytes" yaml:"receiveBytes"`

	// SendBytes represents the number of bytes that have been sent through this tunnel to a peer.
	SendBytes int64 `json:"sendBytes" yaml:"sendBytes"`
}

// WireguardTunnelConfig is configuration for establishing a wireguard tunnel with a remote peer
// using the kernel's built-in wireguard kernel module.
type WireguardTunnelConfig struct {

	// DeviceName represents the name of the network link that will be created for wireguard.
	DeviceName string `json:"deviceName" yaml:"deviceName"`

	// ListenPort specifies the port that the device should listen on for incoming Wireguard connections.
	// The port will always be a UDP listen port that is initialized.
	ListenPort uint16 `json:"listenPort" yaml:"listenPort"`

	// KeyRef represents the local secret object that contains the public/private key
	// to be used by this tunnel for establishing secure connections with the remote peer.
	//
	// The controller will look for two k/v pairs within said secret:
	// - wg.privateKey : This should have as the value that represents the private key that should
	// be used for said link.
	// - wg.publicKey : This should have another string as the value that represents the corresponding
	// public key of the afformentioned private key.
	KeyRef *core.LocalObjectReference `json:"keyRef" yaml:"keyRef"`

	// Peers represents the array of peers that the wireguard will looks for
	// and accept incoming connections from/to.
	Peers []WireguardPeer `json:"peers" yaml:"peers"`
}

// OpenVPNTunnelConfig is configuration for establishing an OpenVPN tunnel connection with a remote
// host. Requires that the OpneVPN service be configured and enabled on the host machine.
type OpenVPNTunnelConfig struct {
}

// L2TPTunnelConfig is configuration for establishing a L2TP tunnel connection with a remote peer.
// utilizes the kernel's built-in L2TP functionality.
type L2TPTunnelConfig struct {

	// DeviceName represents the name of the network link that will be created for l2tp.
	DeviceName string `json:"deviceName" yaml:"deviceName"`
}

// IPSecTunnelConfig is configuration for establishing an IPSec tunnel with a remote host using the kernel's
// built-in IPSec functionality.
type IPSecTunnelConfig struct {

	// DeviceName represents the name of the network link that will be created for ipsec.
	DeviceName string `json:"deviceName" yaml:"deviceName"`
}

// WireguardPeer represents a peer that a wireguard tunnel device will attempt to peer with.
type WireguardPeer struct {
}
