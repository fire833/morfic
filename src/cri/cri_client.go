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

package cri

import (
	"context"
	"errors"
	"time"

	"github.com/fire833/vroute/src/config"

	grpc "google.golang.org/grpc"
	criapi "k8s.io/cri-api/pkg/apis/runtime/v1"
)

var RuntimeClient criapi.RuntimeServiceClient
var ImageClient criapi.ImageServiceClient

var (
	runtimeSockets []string = []string{"unix:///var/run/crio/crio.sock", "unix:///run/containerd/containerd.sock", "unix:///var/run/docker.sock"}
)

func DialCRIEndpoint() error {

	if config.CPRF.CRISocket != "" {
		ctx, close := context.WithTimeout(context.Background(), time.Duration(time.Second*5))
		defer close()
		c, err := grpc.DialContext(ctx, config.CPRF.CRISocket)
		if err == nil {
			createClients(c)
			return nil
		}
	} else {
		for _, sock := range runtimeSockets {
			ctx, close := context.WithTimeout(context.Background(), time.Duration(time.Second*5))
			defer close()
			c, err := grpc.DialContext(ctx, sock)
			if err == nil {
				createClients(c)
				return nil
			} else {
				continue
			}
		}
	}
	return errors.New("error with creating runtime connection")
}

func createClients(conn *grpc.ClientConn) {
	RuntimeClient = criapi.NewRuntimeServiceClient(conn)
	ImageClient = criapi.NewImageServiceClient(conn)
}
