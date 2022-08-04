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

	api "github.com/fire833/morfic/pkg/ipcapi/v1alpha1"
)

func TestGetWireguardTunnel(t *testing.T) {
	type args struct {
		ctx context.Context
		req *api.GetWireguardTunnelRequest
	}
	tests := []struct {
		name     string
		args     args
		wantResp *api.GetWireguardTunnelResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := GetWireguardTunnel(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetWireguardTunnel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("GetWireguardTunnel() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestGetAllWireguardTunnels(t *testing.T) {
	type args struct {
		ctx context.Context
		req *api.GetAllWireguardTunnelsRequest
	}
	tests := []struct {
		name     string
		args     args
		wantResp *api.GetAllWireguardTunnelsResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := GetAllWireguardTunnels(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllWireguardTunnels() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("GetAllWireguardTunnels() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestDeleteWireguardTunnel(t *testing.T) {
	type args struct {
		ctx context.Context
		req *api.DeleteWireguardTunnelRequest
	}
	tests := []struct {
		name     string
		args     args
		wantResp *api.DeleteWireguardTunnelResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := DeleteWireguardTunnel(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteWireguardTunnel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("DeleteWireguardTunnel() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestCreateWireguardTunnel(t *testing.T) {
	type args struct {
		ctx context.Context
		req *api.CreateWireguardTunnelRequest
	}
	tests := []struct {
		name     string
		args     args
		wantResp *api.CreateWireguardTunnelResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := CreateWireguardTunnel(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateWireguardTunnel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("CreateWireguardTunnel() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestUpdateWireguardTunnel(t *testing.T) {
	type args struct {
		ctx context.Context
		req *api.UpdateWireguardTunnelRequest
	}
	tests := []struct {
		name     string
		args     args
		wantResp *api.UpdateWireguardTunnelResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := UpdateWireguardTunnel(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateWireguardTunnel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("UpdateWireguardTunnel() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
