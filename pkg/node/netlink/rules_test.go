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

	"github.com/fire833/morfic/pkg/apis/ipcapi/v1alpha1"
	api "github.com/fire833/morfic/pkg/apis/ipcapi/v1alpha1"
)

func TestNetlinkNodeServer_GetRule(t *testing.T) {
	type fields struct {
		UnimplementedNodeControllerServiceServer         v1alpha1.UnimplementedNodeControllerServiceServer
		UnimplementedNodeFirewallControllerServiceServer v1alpha1.UnimplementedNodeFirewallControllerServiceServer
	}
	type args struct {
		ctx context.Context
		req *api.GetRuleRequest
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp *api.GetRuleResponse
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
			gotResp, err := s.GetRule(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("NetlinkNodeServer.GetRule() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("NetlinkNodeServer.GetRule() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestNetlinkNodeServer_GetAllRules(t *testing.T) {
	type fields struct {
		UnimplementedNodeControllerServiceServer         v1alpha1.UnimplementedNodeControllerServiceServer
		UnimplementedNodeFirewallControllerServiceServer v1alpha1.UnimplementedNodeFirewallControllerServiceServer
	}
	type args struct {
		ctx context.Context
		req *api.GetAllRulesRequest
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp *api.GetAllRulesResponse
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
			gotResp, err := s.GetAllRules(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("NetlinkNodeServer.GetAllRules() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("NetlinkNodeServer.GetAllRules() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestNetlinkNodeServer_DeleteRule(t *testing.T) {
	type fields struct {
		UnimplementedNodeControllerServiceServer         v1alpha1.UnimplementedNodeControllerServiceServer
		UnimplementedNodeFirewallControllerServiceServer v1alpha1.UnimplementedNodeFirewallControllerServiceServer
	}
	type args struct {
		ctx context.Context
		req *api.DeleteRuleRequest
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp *api.DeleteRuleResponse
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
			gotResp, err := s.DeleteRule(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("NetlinkNodeServer.DeleteRule() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("NetlinkNodeServer.DeleteRule() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestNetlinkNodeServer_CreateRule(t *testing.T) {
	type fields struct {
		UnimplementedNodeControllerServiceServer         v1alpha1.UnimplementedNodeControllerServiceServer
		UnimplementedNodeFirewallControllerServiceServer v1alpha1.UnimplementedNodeFirewallControllerServiceServer
	}
	type args struct {
		ctx context.Context
		req *api.CreateRuleRequest
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp *api.CreateRuleResponse
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
			gotResp, err := s.CreateRule(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("NetlinkNodeServer.CreateRule() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("NetlinkNodeServer.CreateRule() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestNetlinkNodeServer_UpdateRule(t *testing.T) {
	type fields struct {
		UnimplementedNodeControllerServiceServer         v1alpha1.UnimplementedNodeControllerServiceServer
		UnimplementedNodeFirewallControllerServiceServer v1alpha1.UnimplementedNodeFirewallControllerServiceServer
	}
	type args struct {
		ctx context.Context
		req *api.UpdateRuleRequest
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp *api.UpdateRuleResponse
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
			gotResp, err := s.UpdateRule(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("NetlinkNodeServer.UpdateRule() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("NetlinkNodeServer.UpdateRule() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
