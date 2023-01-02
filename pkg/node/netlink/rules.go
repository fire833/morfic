/*
*	Copyright (C) 2023 Kendall Tauser
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

	api "github.com/fire833/morfic/pkg/ipcapi/v1alpha1"
	"github.com/fire833/morfic/pkg/node/locks"

	"github.com/vishvananda/netlink"
)

func (s *NetlinkNodeServer) GetRule(ctx context.Context, req *api.GetRuleRequest) (resp *api.GetRuleResponse, err error) {
	return nil, nil
}

func (s *NetlinkNodeServer) GetAllRules(ctx context.Context, req *api.GetAllRulesRequest) (resp *api.GetAllRulesResponse, err error) {
	tl := locks.TableLock

	tl.RLock()

	if _, e := netlink.RuleList(0); e != nil {
		tl.RUnlock()
		return &api.GetAllRulesResponse{}, nil
	} else {
		tl.RUnlock()

		return &api.GetAllRulesResponse{}, nil
	}
}

func (s *NetlinkNodeServer) DeleteRule(ctx context.Context, req *api.DeleteRuleRequest) (resp *api.DeleteRuleResponse, err error) {
	return nil, nil
}

func (s *NetlinkNodeServer) CreateRule(ctx context.Context, req *api.CreateRuleRequest) (resp *api.CreateRuleResponse, err error) {
	return nil, nil
}

func (s *NetlinkNodeServer) UpdateRule(ctx context.Context, req *api.UpdateRuleRequest) (resp *api.UpdateRuleResponse, err error) {
	return nil, nil
}
