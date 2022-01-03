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

package netlink

import (
	"syscall"

	"github.com/mdlayher/netlink"
)

var NetFilterLink *netlink.Conn
var RoutingLink *netlink.Conn

func DialSockets() {
	nf, err := netlink.Dial(syscall.NETLINK_NETFILTER, &netlink.Config{})
	if err != nil {
		panic(err) // Just panic for now, will need a proper error handling layer for startup errors.
	}

	NetFilterLink = nf

	rt, err1 := netlink.Dial(syscall.NETLINK_ROUTE, &netlink.Config{})
	if err1 != nil {
		panic(err1)
	}

	RoutingLink = rt
	return
}
