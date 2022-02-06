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

package main

import (
	"os"
	"syscall"

	"github.com/fasthttp/router"

	api "github.com/fire833/vroute/pkg/apis/v1alpha1"
)

// Unprivileged API listener main function.
func api_main() {

	if os.Getuid() != unprivilegedUID || os.Getgid() != unprivilegedGID {
		if e := dropPriv(); e != nil {
			// TODO send calls to the other processes to kill themselves, and kill the whole control plane.
		}
	}

	router := router.New()

	api.RegisterAuthRoutes(router)
	api.RegisterInterfaceRoutes(router)
	api.RegisterNFRoutes(router)
	api.RegisterRouteRoutes(router)
	api.RegisterServiceRoutes(router)
	api.RegisterWireguardRoutes(router)

}

func dropPriv() error {
	if e := syscall.Setuid(unprivilegedUID); e != nil {
		return e
	}

	if e := syscall.Setgid(unprivilegedGID); e != nil {
		return e
	}
	return nil
}
