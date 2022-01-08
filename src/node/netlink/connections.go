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

package netlink

import (
	"errors"
	"sync"

	"github.com/jsimonetti/rtnetlink"
	"github.com/mdlayher/netlink"
)

var (
	netlinkPool *NetLinkConnectionPool
)

func init() {
	netlinkPool = NewConnPool(5) // Set to 5 for now, should be configurable upstream eventually.
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
	activeConns []*rtnetlink.Conn
	configmu    *sync.Mutex     // Mutex for writers of the config parameter.
	config      *netlink.Config // Condfiguration to be used for queues.
}

func NewConnPool(n int) *NetLinkConnectionPool {
	pool := &NetLinkConnectionPool{
		connmu:   &sync.Mutex{},
		nummu:    &sync.Mutex{},
		configmu: &sync.Mutex{},
		config:   &netlink.Config{}, //Start default with an empty configuration, have editor callers for it.
	}

	e := pool.setDefaultConnections(n)
	if e != nil {
		pool.setDefaultConnections(5) // Failover and just set the active connection queue to 5 by default if there's an error.
	}

	return pool
}

// Top of queue: highest index (where dequeing is done)
// Bottom of queue: index 0 / lowest index (where enqueing is done)

// Called after it is known there is an empty queue spot available at bottom of queue.
func (conn *NetLinkConnectionPool) enqueue(in *rtnetlink.Conn) {
	conn.connmu.Lock()
	for n, _ := range conn.activeConns {
		conn.activeConns[n+1] = conn.activeConns[n] // pull each connection up one.
	}

	conn.activeConns[0] = in
	conn.connmu.Unlock()
}

// Called after it is known that there is at least one available connection at top of queue.
func (conn *NetLinkConnectionPool) dequeue() (out *rtnetlink.Conn) {
	conn.connmu.Lock()

	last := len(conn.activeConns) - 1

	outptr := conn.activeConns[last] // Pull the last from queue.
	conn.activeConns[last] = nil     // Nil the last in the queue
	conn.connmu.Unlock()

	return outptr
}

func (conn *NetLinkConnectionPool) setDefaultConnections(n int) error {
	conn.nummu.Lock()

	if n <= 0 {
		return errors.New("incorrect value for default connection queue size")
	}

	conn.maintainConnections = n
	conn.nummu.Unlock()
	return nil
}

// Create a new connection asynchronously.
func (conn *NetLinkConnectionPool) createConn() {

	c, err := rtnetlink.Dial(conn.config)
	if err != nil {

	}

}

func (conn *NetLinkConnectionPool) createConnImmediate() (*rtnetlink.Conn, error) {
	return rtnetlink.Dial(conn.config)
}

// Internally tries to recycle the connection to the back of the queue.
func (conn *NetLinkConnectionPool) recycleConn(in *rtnetlink.Conn) {

}

func (conn *NetLinkConnectionPool) UpdateQueueSize(n int) error {
	return conn.setDefaultConnections(n)
}

// Callers should use method to acquire a new connection fd to use for their operations.
func (conn *NetLinkConnectionPool) GetConn() (*rtnetlink.Conn, error) {
	if len(conn.activeConns) < 1 {
		// If no connections left, block and synchronously create a new connection and return.
		return conn.createConnImmediate()
	}

}

// Callers should "return" file descriptor to pool after being completing tasks with connection.
// Connections can be recycled internally or thrown out if the number of connections is bigger
// than the number of static connections that are requested for this pool.
func (conn *NetLinkConnectionPool) ReturnConn(in *rtnetlink.Conn) {

}
