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
	invalidMacError = NewError("invalid mac address format", api.ReturnStatusCodes_INVALID_FIELD_ERROR)
)

func ValidateNeighbor(in *api.Neighbor) error {
	var l int

	if in.NeighborAddrType == api.IPType_IPV4 {
		l = net.IPv4len
	} else {
		l = net.IPv6len
	}

	// Validate the actual link address.
	if ip := net.ParseIP(in.IpAddress); ip == nil {
		return badLinkIPError
	} else {
		if len(ip) != l {
			return typeAndActualIPTypeError
		}
	}

	if _, e := net.ParseMAC(in.MacAddress); e != nil {
		return invalidMacError
	}

	return nil
}
