/*
*	Copyright (C) 2023 Kendall Tauser
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

package comparator

import (
	api "github.com/fire833/morfic/pkg/ipcapi/v1alpha1"
)

// Returns true if the link addresses are equivalent. Otherwise returns false.
func CompareLinkAddress(addr1 *api.LinkAddress, addr2 *api.LinkAddress) bool {

	if addr1.Assignment != addr2.Assignment {
		return false
	}

	if !CompareCIDR(addr1.Address, addr2.Address) {
		return false
	}

	if !CompareIPAddress(addr1.Gateway, addr2.Gateway) {
		return false
	}

	if !CompareIPAddress(addr1.DnsServer, addr2.DnsServer) {
		return false
	}

	return true
}

// Returns true if the IP addresses are equivalent. Otherwise returns false.
func CompareIPAddress(addr1 *api.IPAddress, addr2 *api.IPAddress) bool {

	if addr1.Type != addr2.Type {
		return false
	}

	if len(addr1.Address) != len(addr2.Address) {
		return false
	}

	for n, b := range addr1.Address {
		if b != addr2.Address[n] {
			return false
		}
	}

	return true
}

// Returns true if the link MAC addresses are equivalent. Otherwise returns false.
func CompareLinkMACAddress(addr1 *api.LinkMACAddress, addr2 *api.LinkMACAddress) bool {

	if addr1.Assignment != addr2.Assignment {
		return false
	}

	if !CompareMACAddress(addr1.Address, addr2.Address) {
		return false
	}

	if !CompareMACAddress(addr1.BroadcastAddress, addr2.BroadcastAddress) {
		return false
	}

	return true
}

// Returns true if the MAC addresses are equivalent. Otherwise returns false.
func CompareMACAddress(addr1 *api.MACAddress, addr2 *api.MACAddress) bool {

	if len(addr1.Address) != len(addr2.Address) {
		return false
	}

	for n, b := range addr1.Address {
		if b != addr2.Address[n] {
			return false
		}
	}

	return true
}

// Returns true if the CIDRs are equivalent. Otherwise returns false.
func CompareCIDR(cidr1 *api.IPCIDR, cidr2 *api.IPCIDR) bool {

	if !CompareIPAddress(cidr1.Address, cidr2.Address) {
		return false
	}

	if !CompareSubnetMask(cidr1.Mask, cidr2.Mask) {
		return false
	}

	return true
}

// Returns true if the subnet masks are equivalent. Otherwise returns false.
func CompareSubnetMask(mask1 *api.IPMask, mask2 *api.IPMask) bool {

	if mask1.Mask != mask2.Mask {
		return false
	}

	if mask1.Type != mask2.Type {
		return false
	}

	return true
}
