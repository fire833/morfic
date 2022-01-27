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

package node

import (
	"fmt"
	"net"
	"os"

	"github.com/fire833/vroute/src/api/ipcapi/v1alpha1"
	"github.com/fire833/vroute/src/config"
	"github.com/fire833/vroute/src/node/netlink"
	"google.golang.org/grpc"
)

var (
	NodeControllerServer *grpc.Server
)

func BeginNodeServer() error {
	// create the unix bind listener.
	l, e := net.Listen("unix", config.CPRF.NodeControllerSocket)
	if e != nil {
		fmt.Printf("Unable to start node server: %s", e.Error())
		os.Exit(1)
	}

	s := grpc.NewServer(grpc.EmptyServerOption{})

	// Register the services
	s.RegisterService(&v1alpha1.NodeControllerService_ServiceDesc, netlink.NetlinkNodeServer{})

	NodeControllerServer = s
	// This will basically be the last main function for the process.
	return s.Serve(l)
}
