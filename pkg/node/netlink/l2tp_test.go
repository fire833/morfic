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

func TestGetL2TPTunnel(t *testing.T) {
	type args struct {
		ctx context.Context
		req *api.GetL2TPTunnelRequest
	}
	tests := []struct {
		name     string
		args     args
		wantResp *api.GetL2TPTunnelResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := GetL2TPTunnel(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetL2TPTunnel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("GetL2TPTunnel() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestGetAllL2TPTunnels(t *testing.T) {
	type args struct {
		ctx context.Context
		req *api.GetAllL2TPTunnelsRequest
	}
	tests := []struct {
		name     string
		args     args
		wantResp *api.GetAllL2TPTunnelsResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := GetAllL2TPTunnels(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllL2TPTunnels() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("GetAllL2TPTunnels() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestDeleteL2TPTunnel(t *testing.T) {
	type args struct {
		ctx context.Context
		req *api.DeleteL2TPTunnelRequest
	}
	tests := []struct {
		name     string
		args     args
		wantResp *api.DeleteL2TPTunnelResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := DeleteL2TPTunnel(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteL2TPTunnel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("DeleteL2TPTunnel() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestCreateL2TPTunnel(t *testing.T) {
	type args struct {
		ctx context.Context
		req *api.CreateL2TPTunnelRequest
	}
	tests := []struct {
		name     string
		args     args
		wantResp *api.CreateL2TPTunnelResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := CreateL2TPTunnel(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateL2TPTunnel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("CreateL2TPTunnel() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestUpdateL2TPTunnel(t *testing.T) {
	type args struct {
		ctx context.Context
		req *api.UpdateL2TPTunnelRequest
	}
	tests := []struct {
		name     string
		args     args
		wantResp *api.UpdateL2TPTunnelResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := UpdateL2TPTunnel(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateL2TPTunnel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("UpdateL2TPTunnel() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
