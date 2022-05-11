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

package validators

import (
	"net"

	api "github.com/fire833/morfic/pkg/apis/ipcapi/v1alpha1"
)

var (
	badLinkIPError           ValidatorStatus = NewError("invalid formatted link ip address", api.ReturnStatusCodes_INVALID_FIELD_ERROR)
	typeAndActualIPTypeError ValidatorStatus = NewError("provided IP address did not match the specified type", api.ReturnStatusCodes_INVALID_FIELD_ERROR)
	badSubnetMaskError       ValidatorStatus = NewError("incorrect subnet mask value for specified IP type", api.ReturnStatusCodes_INVALID_FIELD_ERROR)
	invalidMacError          ValidatorStatus = NewError("invalid mac address format", api.ReturnStatusCodes_INVALID_FIELD_ERROR)
	invalidIPStringError     ValidatorStatus = NewError("invalid address string provided", api.ReturnStatusCodes_INVALID_FIELD_ERROR)
)

const (
	MACAddressLength = 6
)

// Validtes the link address object, returns nil if the
// link is fine and ready for sending to kernel.
func ValidateLinkAddress(in *api.LinkAddress) error {

	// Validate the actual link address.
	if e := ValidateIPCIDR(in.GetAddress()); e != nil {
		return e
	}

	// Validate the gateway address.
	if e := ValidateIPAddress(in.GetGateway()); e != nil {
		return e
	}

	// Validate the link DNS server address.
	if e := ValidateIPAddress(in.GetDnsServer()); e != nil {
		return e
	}

	return nil
}

func ValidateIPCIDR(in *api.IPCIDR) error {
	e := ValidateIPAddress(in.GetAddress())
	if e != nil {
		return e
	}

	e1 := ValidateSubnetMask(in.GetMask())
	if e1 != nil {
		return e1
	}

	return nil
}

func ValidateIPAddress(in *api.IPAddress) error {
	var l int

	if in.Type == api.IPType_IPV4 {
		l = net.IPv4len
	} else {
		l = net.IPv6len
	}

	if len(in.GetAddress()) != l {
		return typeAndActualIPTypeError
	}

	return nil
}

// Validates both the IP address and makes sure it matches with the defined type.
func ValidateIPAddressAndExternalType(in *api.IPAddress, t api.IPType) error {
	if e := ValidateIPAddress(in); e != nil {
		return e
	}

	var l int

	if t == api.IPType_IPV4 {
		l = net.IPv4len
	} else {
		l = net.IPv6len
	}

	if len(in.GetAddress()) != l {
		return typeAndActualIPTypeError
	}

	return nil
}

func ValidateMACAddress(in *api.MACAddress) error {
	if len(in.GetAddress()) != MACAddressLength {
		return invalidMacError
	} else {
		return nil
	}
}

func ValidateSubnetMask(in *api.IPMask) error {
	var l int

	if in.Type == api.IPType_IPV4 {
		l = net.IPv4len
	} else {
		l = net.IPv6len
	}

	if in.Mask > uint32(l*8) {
		return badSubnetMaskError
	}

	return nil
}
