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
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/fire833/vroute/pkg/controller"
)

// Unprivileged controllerloop subprocess main function.
func controllerMain() {

	fmt.Printf("starting api process, dropping process privilege")

	if os.Getuid() != unprivilegedUID || os.Getgid() != unprivilegedGID {
		if e := dropPriv(); e != nil {
			// TODO send calls to the other processes to kill themselves, and kill the whole control plane.
		}
	}

	fmt.Printf("starting node control process, initializing signal handler")

	sig := make(chan os.Signal, 5)
	signal.Notify(sig)

	for _, controller := range controller.RegisteredControllers {
		go controller.BeginWorkers(controller.RunWorkers())
	}

	// Wait for signals for this process.
	controllerSignalHandler(sig)

}

func controllerSignalHandler(sig chan os.Signal) error {
	for {
		signal := <-sig

		switch signal {
		case syscall.SIGHUP: // Reload configuration
			{

			}
		case syscall.SIGKILL | syscall.SIGINT: // Hard stop process.
			{
				break
			}
		case syscall.SIGTERM: // Gracefully kill the process and all control loops
			{
				break
			}
		default:
			{
				continue // Basically just ignore any other signal.
			}
		}
	}
}
