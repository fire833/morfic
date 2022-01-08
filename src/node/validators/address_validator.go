/*
*	Copyright (C) 2021  Kendall Tauser
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

package validators

import (
	"net"

	api "github.com/fire833/vroute/src/api/ipcapi/v1alpha1"
)

var (
	badLinkIPError           ValidatorStatus = NewError("invalid formatted link ip address", api.ReturnStatusCodes_INVALID_FIELD_ERROR)
	typeAndActualIPTypeError ValidatorStatus = NewError("provided ip address did not match the specified type", api.ReturnStatusCodes_INVALID_FIELD_ERROR)
	badSubnetMaskError       ValidatorStatus = NewError("incorrect subnet mask value for specified IP type", api.ReturnStatusCodes_INVALID_FIELD_ERROR)
)

// Validtes the link address object, returns nil if the
// link is fine and ready for sending to kernel.
func ValidateLinkAddress(in *api.LinkAddress) error {
	var l int

	if in.Type == api.IPType_IPV4 {
		l = net.IPv4len
	} else {
		l = net.IPv6len
	}

	// Validate the actual link address.
	if ip := net.ParseIP(in.Address); ip == nil {
		return badLinkIPError
	} else {
		if len(ip) != l {
			return typeAndActualIPTypeError
		}
	}

	// Validate the link DNS server address.
	ip := net.ParseIP(in.DnsServer)
	if ip == nil {
		return badLinkIPError
	} else {
		if len(ip) != l {
			return typeAndActualIPTypeError
		}
	}

	// Validate the gateway address.
	if ip := net.ParseIP(in.DnsServer); ip == nil {
		return badLinkIPError
	} else {
		if len(ip) != l {
			return typeAndActualIPTypeError
		}
	}

	if in.SubnetMask > uint32(l*8) { // If the mask is bigger than the address itself, fail it.
		return badSubnetMaskError
	}

	return nil
}
