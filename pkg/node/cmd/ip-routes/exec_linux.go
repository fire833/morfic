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

package ip_routes

import "os/exec"

// Wrapper function to find the ip command to manage routing/interface configuration on host.
func execIP(args ...string) (cmd *exec.Cmd, e error) {
	path, e := exec.LookPath("ip")
	return exec.Command(path, args...), e
}

func execRoute(args ...string) (cmd *exec.Cmd, e error) {
	a1 := make([]string, len(args)+1)
	a1 = append(a1, "route")
	a1 = append(a1, args...)

	return execIP(a1...)
}

func execLink(args ...string) (cmd *exec.Cmd, e error) {
	a1 := make([]string, len(args)+1)
	a1 = append(a1, "link")
	a1 = append(a1, args...)

	return execIP(a1...)
}

func execAddr(args ...string) (cmd *exec.Cmd, e error) {
	a1 := make([]string, len(args)+1)
	a1 = append(a1, "address")
	a1 = append(a1, args...)

	return execIP(a1...)
}

func execNeigh(args ...string) (cmd *exec.Cmd, e error) {
	a1 := make([]string, len(args)+1)
	a1 = append(a1, "neigh")
	a1 = append(a1, args...)

	return execIP(a1...)
}

func execTunnel(args ...string) (cmd *exec.Cmd, e error) {
	a1 := make([]string, len(args)+1)
	a1 = append(a1, "tunnel")
	a1 = append(a1, args...)

	return execIP(a1...)
}
