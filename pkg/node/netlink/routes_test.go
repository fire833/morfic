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

package netlink

import (
	"context"
	"reflect"
	"testing"

	"github.com/fire833/morfic/pkg/ipcapi/v1alpha1"
	api "github.com/fire833/morfic/pkg/ipcapi/v1alpha1"
)

func TestNetlinkNodeServer_CreateStaticRoute(t *testing.T) {
	type fields struct {
		UnimplementedNodeControllerServiceServer         v1alpha1.UnimplementedNodeControllerServiceServer
		UnimplementedNodeFirewallControllerServiceServer v1alpha1.UnimplementedNodeFirewallControllerServiceServer
	}
	type args struct {
		ctx context.Context
		req *api.CreateStaticRouteRequest
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp *api.CreateStaticRouteResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nl := &NetlinkNodeServer{
				UnimplementedNodeControllerServiceServer:         tt.fields.UnimplementedNodeControllerServiceServer,
				UnimplementedNodeFirewallControllerServiceServer: tt.fields.UnimplementedNodeFirewallControllerServiceServer,
			}
			gotResp, err := nl.CreateStaticRoute(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("NetlinkNodeServer.CreateStaticRoute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("NetlinkNodeServer.CreateStaticRoute() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestNetlinkNodeServer_DeleteStaticRoute(t *testing.T) {
	type fields struct {
		UnimplementedNodeControllerServiceServer         v1alpha1.UnimplementedNodeControllerServiceServer
		UnimplementedNodeFirewallControllerServiceServer v1alpha1.UnimplementedNodeFirewallControllerServiceServer
	}
	type args struct {
		ctx context.Context
		req *api.DeleteStaticRouteRequest
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp *api.DeleteStaticRouteResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nl := &NetlinkNodeServer{
				UnimplementedNodeControllerServiceServer:         tt.fields.UnimplementedNodeControllerServiceServer,
				UnimplementedNodeFirewallControllerServiceServer: tt.fields.UnimplementedNodeFirewallControllerServiceServer,
			}
			gotResp, err := nl.DeleteStaticRoute(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("NetlinkNodeServer.DeleteStaticRoute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("NetlinkNodeServer.DeleteStaticRoute() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestNetlinkNodeServer_UpdateStaticRoute(t *testing.T) {
	type fields struct {
		UnimplementedNodeControllerServiceServer         v1alpha1.UnimplementedNodeControllerServiceServer
		UnimplementedNodeFirewallControllerServiceServer v1alpha1.UnimplementedNodeFirewallControllerServiceServer
	}
	type args struct {
		ctx context.Context
		req *api.UpdateStaticRouteRequest
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp *api.UpdateStaticRouteResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nl := &NetlinkNodeServer{
				UnimplementedNodeControllerServiceServer:         tt.fields.UnimplementedNodeControllerServiceServer,
				UnimplementedNodeFirewallControllerServiceServer: tt.fields.UnimplementedNodeFirewallControllerServiceServer,
			}
			gotResp, err := nl.UpdateStaticRoute(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("NetlinkNodeServer.UpdateStaticRoute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("NetlinkNodeServer.UpdateStaticRoute() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestNetlinkNodeServer_GetRoute(t *testing.T) {
	type fields struct {
		UnimplementedNodeControllerServiceServer         v1alpha1.UnimplementedNodeControllerServiceServer
		UnimplementedNodeFirewallControllerServiceServer v1alpha1.UnimplementedNodeFirewallControllerServiceServer
	}
	type args struct {
		ctx context.Context
		req *api.GetRouteRequest
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp *api.GetRouteResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nl := &NetlinkNodeServer{
				UnimplementedNodeControllerServiceServer:         tt.fields.UnimplementedNodeControllerServiceServer,
				UnimplementedNodeFirewallControllerServiceServer: tt.fields.UnimplementedNodeFirewallControllerServiceServer,
			}
			gotResp, err := nl.GetRoute(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("NetlinkNodeServer.GetRoute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("NetlinkNodeServer.GetRoute() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestNetlinkNodeServer_GetAllRoutes(t *testing.T) {
	type fields struct {
		UnimplementedNodeControllerServiceServer         v1alpha1.UnimplementedNodeControllerServiceServer
		UnimplementedNodeFirewallControllerServiceServer v1alpha1.UnimplementedNodeFirewallControllerServiceServer
	}
	type args struct {
		ctx context.Context
		req *api.GetAllRoutesRequest
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp *api.GetAllRoutesResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nl := &NetlinkNodeServer{
				UnimplementedNodeControllerServiceServer:         tt.fields.UnimplementedNodeControllerServiceServer,
				UnimplementedNodeFirewallControllerServiceServer: tt.fields.UnimplementedNodeFirewallControllerServiceServer,
			}
			gotResp, err := nl.GetAllRoutes(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("NetlinkNodeServer.GetAllRoutes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("NetlinkNodeServer.GetAllRoutes() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
