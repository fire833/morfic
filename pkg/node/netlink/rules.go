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

	api "github.com/fire833/vroute/pkg/apis/ipcapi/v1alpha1"
)

func (s *NetlinkNodeServer) GetRule(ctx context.Context, req *api.GetRuleRequest) (resp *api.GetRuleResponse, e error) {
	return nil, nil
}

func (s *NetlinkNodeServer) GetAllRules(ctx context.Context, req *api.GetAllRulesRequest) (resp *api.GetAllRulesResponse, e error) {
	return nil, nil
}

func (s *NetlinkNodeServer) DeleteRule(ctx context.Context, req *api.DeleteRuleRequest) (resp *api.DeleteRuleResponse, e error) {
	return nil, nil
}

func (s *NetlinkNodeServer) CreateRule(ctx context.Context, req *api.CreateRuleRequest) (resp *api.CreateRuleResponse, e error) {
	return nil, nil
}

func (s *NetlinkNodeServer) UpdateRule(ctx context.Context, req *api.UpdateRuleRequest) (resp *api.UpdateRuleResponse, e error) {
	return nil, nil
}
