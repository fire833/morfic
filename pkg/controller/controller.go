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

package controller

// List of registered controller loops that are installed and begun at runtime.
var RegisteredControllers []ControllerInterface

// StopChan specifies a standard interface for killing worker threads within controllers.
type StopChan chan uint8

// ControllerInterface represents a controller that manages the state for
// components within the control plane. Controllers watch for updates to
// the state from the API, and work to both set the current state to correlate
// with the desired state. Controllers are not responsible for interfacing with
// the persistence controller to persist the desired state of the controlplane,
// with the notable exception of the storage controller, which directly
// interfaces with and manages backends.
//
// Otherwise, controllers are supposed to create a number of worker threads
// (the number to be determined by the user/main control loop), which then listen
// on channels for incoming events from the API, and then act accordingly with an
// incoming event. They interface with methods such as the node grpc server/
// cri grpc server in order to create the desired state on the host.
type ControllerInterface interface {
	// Begin workers starts up a number of threads for each controller in order
	// to start listening for incoming requests.
	BeginWorkers(number int)

	// Return the number of workers that should be default run on startup.
	RunWorkers() int

	// Stops all workers gracefully by sending them a signal to return and exit.
	// Send a exit code as well with this signal exit in order to invoke specific
	// exit behavior.
	GracefulStop(uint8)
}
