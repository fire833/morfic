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

func (s *NetlinkNodeServer) GetAddress(ctx context.Context, req *api.GetLinkRequest) (resp *api.GetLinkResponse, e error) {
	return nil, nil
}

func (s *NetlinkNodeServer) GetAllAddresses(ctx context.Context, req *api.GetLinkRequest) (resp *api.GetLinkResponse, e error) {
	return nil, nil
}

func (s *NetlinkNodeServer) CreateAddress(ctx context.Context, req *api.GetLinkRequest) (resp *api.GetLinkResponse, e error) {
	return nil, nil
}

func (s *NetlinkNodeServer) DeleteAddress(ctx context.Context, req *api.GetLinkRequest) (resp *api.GetLinkResponse, e error) {
	return nil, nil
}

func (s *NetlinkNodeServer) UpdateAddress(ctx context.Context, req *api.GetLinkRequest) (resp *api.GetLinkResponse, e error) {
	return nil, nil
}
