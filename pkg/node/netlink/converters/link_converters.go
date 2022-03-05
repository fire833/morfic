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

package converters

import (
	"net"

	api "github.com/fire833/vroute/pkg/apis/ipcapi/v1alpha1"
	"github.com/jsimonetti/rtnetlink"
	"golang.org/x/sys/unix"
)

// Converts an API Link to a netlink message object to be sent over the wire.
//
// Should be validated before conversion.
func ConvertAPILinkToNetlinkLink(link *api.Link) *rtnetlink.LinkMessage {

	var flags uint32 = 0

	if link.Attributes.Up {
		flags = SetIsInterfaceUp(flags, true)
	}

	if link.Attributes.Dynamic {
		flags = SetIsDynamic(flags, true)
	}

	if link.Attributes.Broadcast {
		flags = SetIsBroadcast(flags, true)
	}

	if link.Attributes.Multicast {
		flags = SetIsMulticast(flags, true)
	}

	if link.Attributes.Loopback {
		flags = SetIsLoopback(flags, true)
	}

	if link.Attributes.NoArp {
		flags = SetIsArpEnabled(flags, true)
	}

	if link.Attributes.Promiscuous {
		flags = SetIsPromiscuous(flags, true)
	}

	if link.Attributes.P2P {
		flags = SetIsP2P(flags, true)
	}

	hwaddr, _ := net.ParseMAC(string(link.Mac.Address.Address))
	brdaddr, _ := net.ParseMAC(string(link.Mac.BroadcastAddress.Address))

	return &rtnetlink.LinkMessage{
		Flags:  flags,
		Type:   uint16(link.Type), // Need to look into this
		Family: 0,                 // Need to look into this
		Change: 0xFFFFFFFF,        // Reference: https://www.man7.org/linux/man-pages/man7/rtnetlink.7.html
		Attributes: &rtnetlink.LinkAttributes{
			Address:   hwaddr,
			Broadcast: brdaddr,
			Name:      link.Name,
			MTU:       link.Mtu,
			Type:      uint32(link.Type), // Need to look into this
		},
	}
}

// Converts a netlink message to an API link message.
func ConvertNetlinkLinkAPILink(link *rtnetlink.LinkMessage) *api.Link {

	addrbytes := []byte(link.Attributes.Address)
	brdbytes := []byte(link.Attributes.Broadcast)

	return &api.Link{
		Name:  link.Attributes.Name,
		Mtu:   link.Attributes.MTU,
		Type:  api.LinkTypes(link.Type),
		Index: link.Index,
		Mac: &api.LinkMACAddress{
			Address: &api.MACAddress{
				Address: addrbytes,
			},
			BroadcastAddress: &api.MACAddress{
				Address: brdbytes,
			},
		},
		Attributes: &api.LinkAttributes{
			Up:          GetIsInterfaceUp(link.Flags),
			NoArp:       GetIsArpEnabled(link.Flags),
			Dynamic:     GetIsDynamic(link.Flags),
			Multicast:   GetIsMulticast(link.Flags),
			Broadcast:   GetIsBroadcast(link.Flags),
			Loopback:    GetIsLoopback(link.Flags),
			Promiscuous: GetIsPromiscuous(link.Flags),
			P2P:         GetIsP2P(link.Flags),
		},
	}
}

// Extract information from device flags

func GetIsInterfaceUp(flags uint32) bool {
	return flags&unix.IFF_UP == uint32(unix.IFF_UP)
}

func GetIsBroadcast(flags uint32) bool {
	return flags&unix.IFF_BROADCAST == uint32(unix.IFF_BROADCAST)
}

func GetIsMulticast(flags uint32) bool {
	return flags&unix.IFF_MULTICAST == uint32(unix.IFF_MULTICAST)
}

func GetIsLoopback(flags uint32) bool {
	return flags&unix.IFF_LOOPBACK == uint32(unix.IFF_LOOPBACK)
}

func GetIsDynamic(flags uint32) bool {
	return flags&unix.IFF_DYNAMIC == uint32(unix.IFF_DYNAMIC)
}

func GetIsArpEnabled(flags uint32) bool {
	return flags&unix.IFF_NOARP == uint32(unix.IFF_NOARP)
}

func GetIsPromiscuous(flags uint32) bool {
	return flags&unix.IFF_PROMISC == uint32(unix.IFF_PROMISC)
}

func GetIsP2P(flags uint32) bool {
	return flags&unix.IFF_POINTOPOINT == uint32(unix.IFF_POINTOPOINT)
}

// Set device flag information

func SetIsInterfaceUp(flags uint32, set bool) uint32 {
	if set {
		return flags | uint32(unix.IFF_UP)
	} else {
		return flags // Assumes that the bit has already been set to 0 by initialization.
	}
}

func SetIsBroadcast(flags uint32, set bool) uint32 {
	if set {
		return flags | uint32(unix.IFF_BROADCAST)
	} else {
		return flags // Assumes that the bit has already been set to 0 by initialization.
	}
}

func SetIsMulticast(flags uint32, set bool) uint32 {
	if set {
		return flags | uint32(unix.IFF_MULTICAST)
	} else {
		return flags // Assumes that the bit has already been set to 0 by initialization.
	}
}

func SetIsLoopback(flags uint32, set bool) uint32 {
	if set {
		return flags | uint32(unix.IFF_LOOPBACK)
	} else {
		return flags // Assumes that the bit has already been set to 0 by initialization.
	}
}

func SetIsDynamic(flags uint32, set bool) uint32 {
	if set {
		return flags | uint32(unix.IFF_DYNAMIC)
	} else {
		return flags // Assumes that the bit has already been set to 0 by initialization.
	}
}

func SetIsArpEnabled(flags uint32, set bool) uint32 {
	if set {
		return flags | uint32(unix.IFF_NOARP)
	} else {
		return flags // Assumes that the bit has already been set to 0 by initialization.
	}
}

func SetIsPromiscuous(flags uint32, set bool) uint32 {
	if set {
		return flags | uint32(unix.IFF_PROMISC)
	} else {
		return flags // Assumes that the bit has already been set to 0 by initialization.
	}
}

func SetIsP2P(flags uint32, set bool) uint32 {
	if set {
		return flags | uint32(unix.IFF_POINTOPOINT)
	} else {
		return flags // Assumes that the bit has already been set to 0 by initialization.
	}
}
