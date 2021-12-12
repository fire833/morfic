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

package config

import (
	"sync"

	"github.com/fire833/vroute/src/services"
)

type CpConfig struct {
	// mutex for locking the updater for the configuration.
	m sync.Mutex

	// Wireguard tunnels managed by the control plane.
	WgTuns []WireguardTun `json:"wireguard_tunnels"`

	// Interfaces that are controlled/managed by the control plane.
	Ifaces []Interface `json:"interfaces"`

	// Containerized services that are run on the host and controlled by the control plane.
	Services []services.ServiceDescriptor `json:"services"`
}
