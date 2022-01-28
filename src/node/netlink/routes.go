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

	api "github.com/fire833/vroute/src/api/ipcapi/v1alpha1"
)

func (nl *NetlinkNodeServer) CreateStaticRoute(ctx context.Context, req *api.CreateStaticRouteRequest) (resp *api.CreateStaticRouteResponse, e error) {
	return nil, nil
}

func (nl *NetlinkNodeServer) DeleteStaticRoute(ctx context.Context, req *api.DeleteStaticRouteRequest) (resp *api.DeleteStaticRouteResponse, e error) {
	return nil, nil
}

func (nl *NetlinkNodeServer) UpdateStaticRoute(ctx context.Context, req *api.UpdateStaticRouteRequest) (resp *api.UpdateStaticRouteResponse, e error) {
	return nil, nil
}

func (nl *NetlinkNodeServer) GetRoute(ctx context.Context, req *api.GetRouteRequest) (resp *api.GetRouteResponse, e error) {
	return nil, nil
}

func (nl *NetlinkNodeServer) GetAllRoutes(ctx context.Context, req *api.GetAllRoutesRequest) (resp *api.GetAllRoutesResponse, e error) {
	return nil, nil
}
