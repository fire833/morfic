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
	"fmt"
	"net"
	"strconv"

	api "github.com/fire833/morfic/pkg/ipcapi/v1alpha1"
)

// wrapper for validating the integrity of a link object.
// returns the status code for the link at this stage.
func ValidateLink(in *api.Link) error {

	if e := ValidateInterfaceName(in.GetName()); e != nil {
		return e
	}

	for _, addr := range in.Address {
		if e := ValidateLinkAddress(addr); e != nil {
			return e
		} else {
			continue
		}
	}

	if e := ValidateMACAddress(in.Mac.Address); e != nil {
		return e
	}

	if e := ValidateMACAddress(in.Mac.BroadcastAddress); e != nil {
		return e
	}

	return nil
}

// Validates an interface name for a link. Returns error if invalid.
func ValidateInterfaceName(in string) error {
	if len([]byte(in)) >= 16 {
		return NewError("invalid link name length (must be less than or equal to 16 bytes)", api.ReturnStatusCodes_INVALID_FIELD_ERROR)
	}
	return nil
}

// Validates incoming filter elements for filtering when getting a link from the kernel.
// Returns error if invalid.
func ValidateLinkFilterElems(in map[string]string) error {

	for key, value := range in {
		switch key {
		case api.LinkElements_name[int32(api.LinkElements_LINK_NAME)]:
			{
				if e := ValidateInterfaceName(value); e != nil {
					return e
				} else {
					continue
				}
			}
		case api.LinkElements_name[int32(api.LinkElements_LINK_ADDRESS)]:
			{
				if ip := net.ParseIP(value); ip == nil {
					return invalidIPStringError
				} else {
					continue
				}
			}
		case api.LinkElements_name[int32(api.LinkElements_LINK_MACADDRESS)]:
			{
				if _, e := net.ParseMAC(value); e != nil {
					return invalidMacError
				} else {
					continue
				}
			}
		case api.LinkElements_name[int32(api.LinkElements_LINK_TYPE)]:
			{
				if _, e := strconv.ParseInt(value, 10, 32); e != nil {
					return e
				} else {
					continue
				}
			}
		case api.LinkElements_name[int32(api.LinkElements_LINK_MTU)],
			api.LinkElements_name[int32(api.LinkElements_LINK_INDEX)]:
			{
				if _, e := strconv.ParseUint(value, 10, 32); e != nil {
					return e
				} else {
					continue
				}
			}
		case api.LinkElements_name[int32(api.LinkElements_LINK_UP)],
			api.LinkElements_name[int32(api.LinkElements_LINK_ARPENABLED)],
			api.LinkElements_name[int32(api.LinkElements_LINK_DYNAMIC)],
			api.LinkElements_name[int32(api.LinkElements_LINK_MULTICAST)]:
			{
				if _, e := strconv.ParseBool(value); e != nil {
					return e
				} else {
					continue
				}
			}
		default:
			{
				// There is a bad kv/pair.
				return NewError(fmt.Sprintf("invalid link address KV filter pair: %s:%s", key, value), api.ReturnStatusCodes_INVALID_FIELD_ERROR)
			}
		}
	}

	return nil
}
