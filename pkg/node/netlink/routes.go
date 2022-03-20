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
	"github.com/fire833/vroute/pkg/node/converters"
	"github.com/fire833/vroute/pkg/node/validators"
	netl "github.com/vishvananda/netlink"
)

func (nl *NetlinkNodeServer) CreateStaticRoute(ctx context.Context, req *api.CreateStaticRouteRequest) (resp *api.CreateStaticRouteResponse, err error) {

	// Validate the incoming route request.
	if e := validators.ValidateRoute(req.GetRoute()); e != nil {
		return &api.CreateStaticRouteResponse{
			StatusCode: api.ReturnStatusCodes_INVALID_FIELD_ERROR,
			Error:      e.Error(),
		}, nil
	}

	// Add the route after converting it to avalid form for netlink.
	if e := netl.RouteAdd(converters.ConvertAPIRouteToNetlinkRouteNew(req.GetRoute())); e != nil {
		return &api.CreateStaticRouteResponse{
			StatusCode: api.ReturnStatusCodes_INTERNAL_ERROR,
			Error:      e.Error(),
		}, nil
	}

	return &api.CreateStaticRouteResponse{
		StatusCode: api.ReturnStatusCodes_OK,
		Error:      "",
	}, nil

}

func (nl *NetlinkNodeServer) DeleteStaticRoute(ctx context.Context, req *api.DeleteStaticRouteRequest) (resp *api.DeleteStaticRouteResponse, err error) {

	// Validate the incoming route request.
	if e := validators.ValidateRoute(req.GetRoute()); e != nil {
		return &api.DeleteStaticRouteResponse{
			StatusCode: api.ReturnStatusCodes_INVALID_FIELD_ERROR,
			Error:      e.Error(),
			Route:      nil, // Return nothing, since nothing was deleted.
		}, nil
	}

	var delRoute netl.Route

	// Check that the route exists, otherwise abort and return invalid element error.
	if r, e := netl.RouteGet(nil); r == nil || e != nil { // Need to update the destination with a value
		return &api.DeleteStaticRouteResponse{
			StatusCode: api.ReturnStatusCodes_NON_EXISTENT_ELEMENT,
			Error:      e.Error(),
			Route:      nil, // Return nothing, since nothing was deleted.
		}, nil
	} else {
		delRoute = r[0] // TODO need to figure out which route we are looking to delete from the returned options.
	}

	// Now try to delete the route.
	if e := netl.RouteDel(converters.ConvertAPIRouteToNetlinkRouteNew(req.GetRoute())); e != nil {
		return &api.DeleteStaticRouteResponse{
			StatusCode: api.ReturnStatusCodes_INTERNAL_ERROR,
			Error:      e.Error(),
			Route:      nil, // Return nothing, since nothing was deleted.
		}, nil
	}

	return &api.DeleteStaticRouteResponse{
		StatusCode: api.ReturnStatusCodes_OK,
		Error:      "",
		Route:      converters.ConvertNetlinkRouteNewToAPIRoute(&delRoute),
	}, nil

}

func (nl *NetlinkNodeServer) UpdateStaticRoute(ctx context.Context, req *api.UpdateStaticRouteRequest) (resp *api.UpdateStaticRouteResponse, err error) {

	// Validate the incoming route request.
	if e := validators.ValidateRoute(req.GetRoute()); e != nil {
		return &api.UpdateStaticRouteResponse{
			StatusCode: api.ReturnStatusCodes_INVALID_FIELD_ERROR,
			Error:      e.Error(),
		}, nil
	}

	return nil, nil
}

func (nl *NetlinkNodeServer) GetRoute(ctx context.Context, req *api.GetRouteRequest) (resp *api.GetRouteResponse, err error) {
	return nil, nil
}

func (nl *NetlinkNodeServer) GetAllRoutes(ctx context.Context, req *api.GetAllRoutesRequest) (resp *api.GetAllRoutesResponse, err error) {
	return nil, nil
}
