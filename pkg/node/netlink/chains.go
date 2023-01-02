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
)

func (s *NetlinkNodeServer) GetChain(ctx context.Context, req *api.GetChainRequest) (resp *api.GetChainResponse, err error) {
	return nil, nil
}

func (s *NetlinkNodeServer) GetAllChains(ctx context.Context, req *api.GetAllChainsRequest) (resp *api.GetAllChainsResponse, err error) {
	return nil, nil
}

func (s *NetlinkNodeServer) DeleteChain(ctx context.Context, req *api.DeleteChainRequest) (resp *api.DeleteChainResponse, err error) {
	return nil, nil
}

func (s *NetlinkNodeServer) CreateChain(ctx context.Context, req *api.CreateChainRequest) (resp *api.CreateChainResponse, err error) {
	return nil, nil
}

func (s *NetlinkNodeServer) UpdateChain(ctx context.Context, req *api.UpdateChainRequest) (resp *api.UpdateChainResponse, err error) {
	return nil, nil
}
