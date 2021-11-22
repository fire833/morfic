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

package wg

import "golang.zx2c4.com/wireguard/wgctrl"

var WgClient *wgctrl.Client

func init() {
	client, err := wgctrl.New()
	if err != nil {
		panic(err) // Solve this panic in the future, but probably results in exit of control plane.
	}

	WgClient = client
	return
}
