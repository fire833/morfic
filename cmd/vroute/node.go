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
	"os/signal"
	"syscall"

	"github.com/fire833/vroute/src/node"
)

// Priviledged node operator process main function.
func node_main() {

	sig := make(chan os.Signal)
	signal.Notify(sig)

	// Begin handler on thread to free up main for gRPC listener.
	go signalHandler(sig)

	// Start the node gRPC listener.
	if e := node.BeginNodeServer(); e != nil {
		node.NodeControllerServer.GracefulStop() // Kill the server gracefully and exit.
	}

	os.Exit(1) // Kill subprocess after shutdown of server.

}

func signalHandler(sig chan os.Signal) error {
	for {
		signal := <-sig

		switch signal {
		case syscall.SIGHUP: // Reload configuration, gracefully stop and restart the server.
			{

			}
		case syscall.SIGKILL | syscall.SIGINT: // Hard stop process.
			{
				node.NodeControllerServer.Stop()
				break
			}
		case syscall.SIGTERM:
			{
				node.NodeControllerServer.GracefulStop()
				break
			}
		default:
			{
				continue // Basically just ignore any other signal.
			}
		}
	}
	return nil
}
