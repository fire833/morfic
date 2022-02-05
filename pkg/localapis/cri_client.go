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
	"errors"
	"log"
	"time"

	"github.com/fire833/vroute/pkg/config"

	grpc "google.golang.org/grpc"
	criapi "k8s.io/cri-api/pkg/apis/runtime/v1"
)

var (
	/*
		Client endpoint to be used by callers for node control operations.
		This client is used for managing container/pod sandbox state for
		containerized services that are being hosted on the host and managed by vroute.
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
	RuntimeClient criapi.RuntimeServiceClient
	/*
		Client endpoint to be used by callers for node control operations.
		This client is used for managing images that are stored/used by the local
		CRI implementation, such as pulling images, deleting unused ones, and checking
		disk usage for images stored on the host.
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
	ImageClient criapi.ImageServiceClient

	// The local CRI connection endpoint that is maintained by the gRPC library.
	runtimeConn *grpc.ClientConn

	runtimeSockets []string = []string{"unix:///var/run/crio/crio.sock", "unix:///run/containerd/containerd.sock", "unix:///var/run/docker.sock"}
)

// Used internally for bootstrapping connections to the local CRI instance.
func dialCRIEndpoints() error {

	if config.CPRF.CRISocket != "" {
		ctx, close := context.WithTimeout(context.Background(), time.Duration(time.Second*2))
		defer close()
		c, err := grpc.DialContext(ctx, config.CPRF.CRISocket)
		if err == nil {
			runtimeConn = c

			createRuntimeClients(c)
			log.Printf("established container service runtime API gRPC connection with %s", config.CPRF.CRISocket)
			return nil
		}
	} else {
		for _, sock := range runtimeSockets {
			ctx, close := context.WithTimeout(context.Background(), time.Duration(time.Second*2))
			defer close()
			c, err := grpc.DialContext(ctx, sock)
			if err == nil {
				runtimeConn = c

				createRuntimeClients(c)
				log.Printf("established container service runtime API gRPC connection with %s", sock)
				return nil
			} else {
				continue
			}
		}
	}
	return errors.New("error with creating runtime connection")
}

func createRuntimeClients(conn *grpc.ClientConn) {
	RuntimeClient = criapi.NewRuntimeServiceClient(conn)
	ImageClient = criapi.NewImageServiceClient(conn)
}
