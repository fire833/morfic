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

package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// Unprivileged API listener main function.
func apiMain() {

	fmt.Printf("starting api process, dropping process privilege")

	if os.Getuid() != unprivilegedUID || os.Getgid() != unprivilegedGID {
		if e := dropPriv(); e != nil {
			// TODO send calls to the other processes to kill themselves, and kill the whole control plane.
		}
	}

	fmt.Printf("starting api process, initializing signal handler")

	sig := make(chan os.Signal, 5)
	signal.Notify(sig)

	go apisignalHandler(sig)

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

func apisignalHandler(sig chan os.Signal) error {
	for {
		signal := <-sig

		switch signal {
		case syscall.SIGHUP: // Reload configuration, gracefully stop and restart the server.
			{
				continue
			}
		case syscall.SIGKILL | syscall.SIGINT: // Hard stop the control plane.
			{
				break
			}
		case syscall.SIGTERM: //Gracefully stop the control plane
			{
				break
			}
		default:
			{
				continue
			}
		}
	}
	// return nil
}
