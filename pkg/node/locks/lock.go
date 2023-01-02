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

package locks

import "sync"

// Define a set of mutexes to be called for each network object that will
// be mutated via the node agent.
var (
	RouteLock    *routeLock    = new(routeLock)
	NeighborLock *neighborLock = new(neighborLock)
	AddrLock     *addrLock     = new(addrLock)
	LinkLock     *linkLock     = new(linkLock)
	TableLock    *tableLock    = new(tableLock)
)

// RouteLock serves as a wrapper around RWMutex as a means for consistent
// changes to host state via the node server. Should be called for reading/writing
// routes over netlink.
type routeLock struct {
	sync.RWMutex
}

// NeighborLock serves as a wrapper around RWMutex as a means for consistent
// changes to host state via the node server. Should be called for reading/writing
// neighbor states over netlink.
type neighborLock struct {
	sync.RWMutex
}

// AddrLock serves as a wrapper around RWMutex as a means for consistent
// changes to host state via the node server. Should be called for reading/writing
// addresses over netlink.
type addrLock struct {
	sync.RWMutex
}

// LinkLock serves as a wrapper around RWMutex as a means for consistent
// changes to host state via the node server. Should be called for reading/writing
// links over netlink. Should also be called whenever addresses are updated as well.
type linkLock struct {
	sync.RWMutex
}

// TableLock serves as a wrapper around RWMutex as a means for consistent
// changes to host state via the node server. Should be called for reading/writing
// netfilter tables over netlink. Should also be called whenever addresses are updated as well.
type tableLock struct {
	sync.RWMutex
}
