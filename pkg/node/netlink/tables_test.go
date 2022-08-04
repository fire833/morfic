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

func TestNetlinkNodeServer_GetTable(t *testing.T) {
	type fields struct {
		UnimplementedNodeControllerServiceServer         v1alpha1.UnimplementedNodeControllerServiceServer
		UnimplementedNodeFirewallControllerServiceServer v1alpha1.UnimplementedNodeFirewallControllerServiceServer
	}
	type args struct {
		ctx context.Context
		req *api.GetTableRequest
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp *api.GetTableResponse
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
			gotResp, err := s.GetTable(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("NetlinkNodeServer.GetTable() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("NetlinkNodeServer.GetTable() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestNetlinkNodeServer_GetAllTables(t *testing.T) {
	type fields struct {
		UnimplementedNodeControllerServiceServer         v1alpha1.UnimplementedNodeControllerServiceServer
		UnimplementedNodeFirewallControllerServiceServer v1alpha1.UnimplementedNodeFirewallControllerServiceServer
	}
	type args struct {
		ctx context.Context
		req *api.GetAllTablesRequest
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp *api.GetAllTablesResponse
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
			gotResp, err := s.GetAllTables(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("NetlinkNodeServer.GetAllTables() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("NetlinkNodeServer.GetAllTables() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestNetlinkNodeServer_DeleteTable(t *testing.T) {
	type fields struct {
		UnimplementedNodeControllerServiceServer         v1alpha1.UnimplementedNodeControllerServiceServer
		UnimplementedNodeFirewallControllerServiceServer v1alpha1.UnimplementedNodeFirewallControllerServiceServer
	}
	type args struct {
		ctx context.Context
		req *api.DeleteTableRequest
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp *api.DeleteTableResponse
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
			gotResp, err := s.DeleteTable(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("NetlinkNodeServer.DeleteTable() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("NetlinkNodeServer.DeleteTable() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestNetlinkNodeServer_CreateTable(t *testing.T) {
	type fields struct {
		UnimplementedNodeControllerServiceServer         v1alpha1.UnimplementedNodeControllerServiceServer
		UnimplementedNodeFirewallControllerServiceServer v1alpha1.UnimplementedNodeFirewallControllerServiceServer
	}
	type args struct {
		ctx context.Context
		req *api.CreateTableRequest
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp *api.CreateTableResponse
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
			gotResp, err := s.CreateTable(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("NetlinkNodeServer.CreateTable() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("NetlinkNodeServer.CreateTable() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestNetlinkNodeServer_UpdateTable(t *testing.T) {
	type fields struct {
		UnimplementedNodeControllerServiceServer         v1alpha1.UnimplementedNodeControllerServiceServer
		UnimplementedNodeFirewallControllerServiceServer v1alpha1.UnimplementedNodeFirewallControllerServiceServer
	}
	type args struct {
		ctx context.Context
		req *api.UpdateTableRequest
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp *api.UpdateTableResponse
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
			gotResp, err := s.UpdateTable(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("NetlinkNodeServer.UpdateTable() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("NetlinkNodeServer.UpdateTable() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
