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

package localapis

import (
	"log"
)

// Dials the local UNIX endpoints for both the node controller as well as the container
// runtime interface shim endpoints and stands up the client that will be called by the
// control plane API.
func DialLocalAPIs() error {
	if e := dialCRIEndpoints(); e != nil {
		log.Printf("unable to dial CRI APIs: %s", e.Error())
		return e
	}
	if e := dialNodeEndpoints(); e != nil {
		log.Printf("unable to dial Node APIs: %s", e.Error())
		return e
	}

	return nil
}

// Tears down local API connections gracefully on a termination of the control plane.
// Should be torn down AFTER API listener has hopefully gracefully shut down.
func CloseLocalAPIS() error {
	runtimeConn.Close()
	nodeConn.Close()

	return nil
}
