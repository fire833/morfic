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

func TestNetlinkNodeServer_GetChain(t *testing.T) {
	type fields struct {
		UnimplementedNodeControllerServiceServer         v1alpha1.UnimplementedNodeControllerServiceServer
		UnimplementedNodeFirewallControllerServiceServer v1alpha1.UnimplementedNodeFirewallControllerServiceServer
	}
	type args struct {
		ctx context.Context
		req *api.GetChainRequest
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp *api.GetChainResponse
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
			gotResp, err := s.GetChain(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("NetlinkNodeServer.GetChain() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("NetlinkNodeServer.GetChain() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestNetlinkNodeServer_GetAllChains(t *testing.T) {
	type fields struct {
		UnimplementedNodeControllerServiceServer         v1alpha1.UnimplementedNodeControllerServiceServer
		UnimplementedNodeFirewallControllerServiceServer v1alpha1.UnimplementedNodeFirewallControllerServiceServer
	}
	type args struct {
		ctx context.Context
		req *api.GetAllChainsRequest
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp *api.GetAllChainsResponse
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
			gotResp, err := s.GetAllChains(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("NetlinkNodeServer.GetAllChains() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("NetlinkNodeServer.GetAllChains() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestNetlinkNodeServer_DeleteChain(t *testing.T) {
	type fields struct {
		UnimplementedNodeControllerServiceServer         v1alpha1.UnimplementedNodeControllerServiceServer
		UnimplementedNodeFirewallControllerServiceServer v1alpha1.UnimplementedNodeFirewallControllerServiceServer
	}
	type args struct {
		ctx context.Context
		req *api.DeleteChainRequest
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp *api.DeleteChainResponse
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
			gotResp, err := s.DeleteChain(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("NetlinkNodeServer.DeleteChain() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("NetlinkNodeServer.DeleteChain() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestNetlinkNodeServer_CreateChain(t *testing.T) {
	type fields struct {
		UnimplementedNodeControllerServiceServer         v1alpha1.UnimplementedNodeControllerServiceServer
		UnimplementedNodeFirewallControllerServiceServer v1alpha1.UnimplementedNodeFirewallControllerServiceServer
	}
	type args struct {
		ctx context.Context
		req *api.CreateChainRequest
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp *api.CreateChainResponse
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
			gotResp, err := s.CreateChain(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("NetlinkNodeServer.CreateChain() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("NetlinkNodeServer.CreateChain() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestNetlinkNodeServer_UpdateChain(t *testing.T) {
	type fields struct {
		UnimplementedNodeControllerServiceServer         v1alpha1.UnimplementedNodeControllerServiceServer
		UnimplementedNodeFirewallControllerServiceServer v1alpha1.UnimplementedNodeFirewallControllerServiceServer
	}
	type args struct {
		ctx context.Context
		req *api.UpdateChainRequest
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp *api.UpdateChainResponse
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
			gotResp, err := s.UpdateChain(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("NetlinkNodeServer.UpdateChain() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("NetlinkNodeServer.UpdateChain() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
