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

func TestNetlinkNodeServer_GetNeighbor(t *testing.T) {
	type fields struct {
		UnimplementedNodeControllerServiceServer         v1alpha1.UnimplementedNodeControllerServiceServer
		UnimplementedNodeFirewallControllerServiceServer v1alpha1.UnimplementedNodeFirewallControllerServiceServer
	}
	type args struct {
		ctx context.Context
		req *api.GetNeighborRequest
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp *api.GetNeighborResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &NetlinkNodeServer{
				UnimplementedNodeControllerServiceServer:         tt.fields.UnimplementedNodeControllerServiceServer,
				UnimplementedNodeFirewallControllerServiceServer: tt.fields.UnimplementedNodeFirewallControllerServiceServer,
			}
			gotResp, err := s.GetNeighbor(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("NetlinkNodeServer.GetNeighbor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("NetlinkNodeServer.GetNeighbor() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestNetlinkNodeServer_GetAllNeighbors(t *testing.T) {
	type fields struct {
		UnimplementedNodeControllerServiceServer         v1alpha1.UnimplementedNodeControllerServiceServer
		UnimplementedNodeFirewallControllerServiceServer v1alpha1.UnimplementedNodeFirewallControllerServiceServer
	}
	type args struct {
		ctx context.Context
		req *api.GetAllNeighborsRequest
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp *api.GetAllNeighborsResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &NetlinkNodeServer{
				UnimplementedNodeControllerServiceServer:         tt.fields.UnimplementedNodeControllerServiceServer,
				UnimplementedNodeFirewallControllerServiceServer: tt.fields.UnimplementedNodeFirewallControllerServiceServer,
			}
			gotResp, err := s.GetAllNeighbors(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("NetlinkNodeServer.GetAllNeighbors() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("NetlinkNodeServer.GetAllNeighbors() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestNetlinkNodeServer_CreateNeighbor(t *testing.T) {
	type fields struct {
		UnimplementedNodeControllerServiceServer         v1alpha1.UnimplementedNodeControllerServiceServer
		UnimplementedNodeFirewallControllerServiceServer v1alpha1.UnimplementedNodeFirewallControllerServiceServer
	}
	type args struct {
		ctx context.Context
		req *api.CreateNeighborRequest
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp *api.CreateNeighborResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &NetlinkNodeServer{
				UnimplementedNodeControllerServiceServer:         tt.fields.UnimplementedNodeControllerServiceServer,
				UnimplementedNodeFirewallControllerServiceServer: tt.fields.UnimplementedNodeFirewallControllerServiceServer,
			}
			gotResp, err := s.CreateNeighbor(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("NetlinkNodeServer.CreateNeighbor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("NetlinkNodeServer.CreateNeighbor() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestNetlinkNodeServer_DeleteNeighbor(t *testing.T) {
	type fields struct {
		UnimplementedNodeControllerServiceServer         v1alpha1.UnimplementedNodeControllerServiceServer
		UnimplementedNodeFirewallControllerServiceServer v1alpha1.UnimplementedNodeFirewallControllerServiceServer
	}
	type args struct {
		ctx context.Context
		req *api.DeleteNeighborRequest
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp *api.DeleteNeighborRequest
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &NetlinkNodeServer{
				UnimplementedNodeControllerServiceServer:         tt.fields.UnimplementedNodeControllerServiceServer,
				UnimplementedNodeFirewallControllerServiceServer: tt.fields.UnimplementedNodeFirewallControllerServiceServer,
			}
			gotResp, err := s.DeleteNeighbor(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("NetlinkNodeServer.DeleteNeighbor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("NetlinkNodeServer.DeleteNeighbor() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestNetlinkNodeServer_UpdateNeighbor(t *testing.T) {
	type fields struct {
		UnimplementedNodeControllerServiceServer         v1alpha1.UnimplementedNodeControllerServiceServer
		UnimplementedNodeFirewallControllerServiceServer v1alpha1.UnimplementedNodeFirewallControllerServiceServer
	}
	type args struct {
		ctx context.Context
		req *api.UpdateNeighborRequest
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp *api.UpdateNeighborResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &NetlinkNodeServer{
				UnimplementedNodeControllerServiceServer:         tt.fields.UnimplementedNodeControllerServiceServer,
				UnimplementedNodeFirewallControllerServiceServer: tt.fields.UnimplementedNodeFirewallControllerServiceServer,
			}
			gotResp, err := s.UpdateNeighbor(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("NetlinkNodeServer.UpdateNeighbor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("NetlinkNodeServer.UpdateNeighbor() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
