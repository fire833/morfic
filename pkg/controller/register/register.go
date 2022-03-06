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

package register

import (
	"github.com/fire833/vroute/pkg/controller"
	"github.com/fire833/vroute/pkg/controller/interfaces"
)

// Load in all the registered controller loops to be started
func RegisterControllers() {

	controller.RegisteredControllers = make([]controller.ControllerInterface, 1) // Set to number of controller loops

	iface := &interfaces.InterfaceController{}
	controller.RegisteredControllers = append(controller.RegisteredControllers, iface)

}
