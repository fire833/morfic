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
	"github.com/jsimonetti/rtnetlink"
	"github.com/vishvananda/netlink"

	api "github.com/fire833/vroute/pkg/apis/ipcapi/v1alpha1"
)

func ConvertAPIAddrtoNetlinkAddrNew(addr *api.LinkAddress) *netlink.Addr {
	return nil
}

func ConvertAPIAddrtoNetlinkAddr(addr *api.IPCIDR) *rtnetlink.AddressMessage {
	return nil
}

func ConvertNetlinkAddrNewToAPIAddr(addr *netlink.Addr) *api.LinkAddress {
	return nil
}

func ConvertNetlinkAddrToAPIAddr(addr *rtnetlink.AddressMessage) *api.LinkAddress {
	return nil
}
