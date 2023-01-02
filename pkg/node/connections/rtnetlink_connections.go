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

package connections

import (
	"errors"
	"sync"

	"github.com/jsimonetti/rtnetlink"
	"github.com/mdlayher/netlink"
)

var (
	NetlinkPool *NetLinkConnectionPool
)

func init() {
	NetlinkPool = NewConnPool(5) // Set to 5 for now, should be configurable upstream eventually.
}

type ConnInterface interface {
}

// Internal management interface to cache/queue netlink connections to the kernel asynchronously,
// for better performance on the API side. The interface asynchronously builds up a few queued
// connections at any one time, and then
type NetLinkConnectionPool struct {
	nummu               *sync.Mutex // Mutex for number of connections that should be kept at one time.
	maintainConnections int         // Number of connections to try and keep active at any one time
	connmu              *sync.Mutex // Mutex for writers of the connection array.
	// List of queued connection objects that can be passed to callers.
	// Uses pointers for theoretically quicker queueing/dequeueing.
	activeConns *List
	configmu    *sync.Mutex     // Mutex for writers of the config parameter.
	config      *netlink.Config // Condfiguration to be used for queues.
}

func NewConnPool(n int) *NetLinkConnectionPool {
	pool := &NetLinkConnectionPool{
		connmu:      &sync.Mutex{},
		nummu:       &sync.Mutex{},
		configmu:    &sync.Mutex{},
		config:      &netlink.Config{}, //Start default with an empty configuration, have editor callers for it.
		activeConns: NewList(),
	}

	e := pool.setDefaultConnections(n)
	if e != nil {
		pool.setDefaultConnections(5) // Failover and just set the active connection queue to 5 by default if there's an error.
	}

	return pool
}

// Manager thread will spin up connections asynchronously.
func (conn *NetLinkConnectionPool) StartManager() {

}

// Top/front of queue: highest index (where dequeing is done)
// Bottom/back of queue: index 0 / lowest index (where enqueing is done)

func (conn *NetLinkConnectionPool) enqueue(in *rtnetlink.Conn) {
	conn.connmu.Lock()
	conn.activeConns.PushBack(in)
	conn.connmu.Unlock()
}

// Called after it is known that there is at least one available connection at top of queue.
func (conn *NetLinkConnectionPool) dequeue() (out *rtnetlink.Conn) {
	conn.connmu.Lock()
	elem := conn.activeConns.Remove(conn.activeConns.Front()) //remove the front element from queue.
	conn.connmu.Unlock()

	return elem
}

func (conn *NetLinkConnectionPool) setDefaultConnections(n int) error {
	conn.nummu.Lock()

	if n <= 0 {
		conn.nummu.Unlock()
		return errors.New("incorrect value for default connection queue size")
	}

	conn.maintainConnections = n
	conn.nummu.Unlock()
	return nil
}

// Create a new connection asynchronously.
func (conn *NetLinkConnectionPool) createConn() {
	if conn.activeConns.Len() < conn.maintainConnections {
		c, _ := rtnetlink.Dial(conn.config)
		conn.enqueue(c)
	}
}

func (conn *NetLinkConnectionPool) createConnImmediate() (*rtnetlink.Conn, error) {
	c, e := rtnetlink.Dial(conn.config)
	if e != nil {
		return c, e
	}

	return c, nil
}

func (conn *NetLinkConnectionPool) UpdateQueueSize(n int) error {
	return conn.setDefaultConnections(n)
}

// Callers should use method to acquire a new connection fd to use for their operations.
func (conn *NetLinkConnectionPool) GetConn() (*rtnetlink.Conn, error) {
	if conn.activeConns.Len() < 1 {
		// If no connections left, block and synchronously create a new connection and return.
		return conn.createConnImmediate()
	} else {
		c := conn.dequeue()
		conn.createConn() // Replace the pulled connection.
		return c, nil     // Return the top of the queue.
	}
}

// Callers should "return" file descriptor to pool after being completing tasks with connection.
// Connections can be recycled internally or thrown out if the number of connections is bigger
// than the number of static connections that are requested for this pool.
func (conn *NetLinkConnectionPool) ReturnConn(in *rtnetlink.Conn) {
	if conn.activeConns.Len() < conn.maintainConnections {
		conn.enqueue(in)
	} else {
		in.Close() // Close the connection after return.
	}
}
