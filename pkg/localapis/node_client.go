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

package localapis

import (
	"context"
	"log"
	"time"

	ipcapi "github.com/fire833/morfic/pkg/apis/ipcapi/v1alpha1"
	"google.golang.org/grpc"
)

var (
	/*
		Client endpoint to be used by callers for node control operations.
		This client is used for handling sysctls, links, neighbors, routes,
		IP addresses for host links, etc.
		This is a gRPC endpoint, which callers can use to invoke remote
		operations using the specified methods in the client.

		This client SHOULD NOT be referenced by callers AFTER a termination
		signal has been sent to the control plane AND after the REST API has
		been instructed to gracefully shutdown. For callers within API handlers,
		connections will not be closed until the REST API has gracefully shut down,
		so you are free to call at any point within an API handler.

		This client SHOULD NOT be referenced by callers BEFORE the client has
		been established by the bootstrapping process. For callers within API
		handlers, the API will not begin to listen for requests until after client
		establishment, so you are free to call at any point within an API handler.
	*/
	NodeClient ipcapi.NodeControllerServiceClient
	/*
		Client endpoint to be used by callers for node control operations.
		This client is used for handling netfilter/nftables rules on the host
		and managing the firewall state of the host.
		This is a gRPC endpoint, which callers can use to invoke remote
		operations using the specified methods in the client.

		This client SHOULD NOT be referenced by callers AFTER a termination
		signal has been sent to the control plane AND after the REST API has
		been instructed to gracefully shutdown. For callers within API handlers,
		connections will not be closed until the REST API has gracefully shut down,
		so you are free to call at any point within an API handler.

		This client SHOULD NOT be referenced by callers BEFORE the client has
		been established by the bootstrapping process. For callers within API
		handlers, the API will not begin to listen for requests until after client
		establishment, so you are free to call at any point within an API handler.
	*/
	NodeFirewallClient ipcapi.NodeFirewallControllerServiceClient

	/*
	 */
	NodeVPNClient ipcapi.NodeVPNControllerServiceClient

	// The local node connection endpoint that is maintained by the gRPC library.
	nodeConn *grpc.ClientConn
)

const (
	nodeSocket string = "unix:///var/run/morfic.sock"
)

// Used internally for bootstrapping connections to the local node controller
// instance that was forked from the control plane not long ago in the bootup process.
func dialNodeEndpoints() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*2))
	defer cancel()
	c, err := grpc.DialContext(ctx, nodeSocket)
	if err != nil {
		return err
	} else {
		nodeConn = c

		createNodeClients(c)
		log.Printf("established node API gRPC connection with %s", nodeSocket)
		return nil
	}

}

func createNodeClients(conn *grpc.ClientConn) {
	NodeClient = ipcapi.NewNodeControllerServiceClient(conn)
	NodeFirewallClient = ipcapi.NewNodeFirewallControllerServiceClient(conn)
	NodeVPNClient = ipcapi.NewNodeVPNControllerServiceClient(conn)
}
