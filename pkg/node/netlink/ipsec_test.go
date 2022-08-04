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

func TestGetIPSecTunnel(t *testing.T) {
	type args struct {
		ctx context.Context
		req *api.GetIPSecTunnelRequest
	}
	tests := []struct {
		name     string
		args     args
		wantResp *api.GetIPSecTunnelResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := GetIPSecTunnel(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetIPSecTunnel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("GetIPSecTunnel() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestGetAllIPSecTunnels(t *testing.T) {
	type args struct {
		ctx context.Context
		req *api.GetAllIPSecTunnelsRequest
	}
	tests := []struct {
		name     string
		args     args
		wantResp *api.GetAllIPSecTunnelsResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := GetAllIPSecTunnels(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllIPSecTunnels() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("GetAllIPSecTunnels() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestDeleteIPSecTunnel(t *testing.T) {
	type args struct {
		ctx context.Context
		req *api.DeleteIPSecTunnelRequest
	}
	tests := []struct {
		name     string
		args     args
		wantResp *api.DeleteIPSecTunnelResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := DeleteIPSecTunnel(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteIPSecTunnel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("DeleteIPSecTunnel() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestCreateIPSecTunnel(t *testing.T) {
	type args struct {
		ctx context.Context
		req *api.CreateIPSecTunnelRequest
	}
	tests := []struct {
		name     string
		args     args
		wantResp *api.CreateIPSecTunnelResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := CreateIPSecTunnel(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateIPSecTunnel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("CreateIPSecTunnel() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestUpdateIPSecTunnel(t *testing.T) {
	type args struct {
		ctx context.Context
		req *api.UpdateIPSecTunnelRequest
	}
	tests := []struct {
		name     string
		args     args
		wantResp *api.UpdateIPSecTunnelResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := UpdateIPSecTunnel(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateIPSecTunnel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("UpdateIPSecTunnel() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
