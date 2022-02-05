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

package ip_routes

import (
	"context"

	"github.com/fire833/vroute/pkg/apis/ipcapi/v1alpha1"
)

func (s *CMDNodeServer) CreateStaticRoute(ctx context.Context, req *v1alpha1.CreateStaticRouteRequest) (resp *v1alpha1.CreateStaticRouteResponse, e error) {
	defer ctx.Done()

	return nil, nil
}

func (s *CMDNodeServer) DeleteStaticRoute(ctx context.Context, req *v1alpha1.DeleteStaticRouteRequest) (resp *v1alpha1.DeleteStaticRouteResponse, e error) {
	defer ctx.Done()

	for _, route := range req.Route {
		route.Descriptor()
	}

	return nil, nil
}

func (s *CMDNodeServer) UpdateStaticRoute(ctx context.Context, req *v1alpha1.UpdateStaticRouteRequest) (resp *v1alpha1.UpdateStaticRouteResponse, e error) {
	defer ctx.Done()

	return nil, nil
}

func (s *CMDNodeServer) GetRoute(ctx context.Context, req *v1alpha1.GetRouteRequest) (resp *v1alpha1.GetRouteResponse, e error) {
	defer ctx.Done()

	return nil, nil
}

func (s *CMDNodeServer) GetAllRoutes(ctx context.Context, req *v1alpha1.GetAllRoutesRequest) (resp *v1alpha1.GetAllRoutesResponse, e error) {
	defer ctx.Done()

	return nil, nil
}
