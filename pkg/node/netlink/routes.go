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
	nl "github.com/vishvananda/netlink"
)

func (s *NetlinkNodeServer) CreateStaticRoute(ctx context.Context, req *api.CreateStaticRouteRequest) (resp *api.CreateStaticRouteResponse, err error) {

	// Validate the incoming route request.
	if e := validators.ValidateRoute(req.GetRoute()); e != nil {
		return &api.CreateStaticRouteResponse{
			StatusCode: api.ReturnStatusCodes_INVALID_FIELD_ERROR,
			Error:      e.Error(),
		}, nil
	}

	// Add the route after converting it to avalid form for netlink.
	if e := nl.RouteAdd(converters.ConvertAPIRouteToNetlinkRouteNew(req.GetRoute())); e != nil {
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

func (s *NetlinkNodeServer) DeleteStaticRoute(ctx context.Context, req *api.DeleteStaticRouteRequest) (resp *api.DeleteStaticRouteResponse, err error) {

	// Validate the incoming route request.
	if e := validators.ValidateRoute(req.GetRoute()); e != nil {
		return &api.DeleteStaticRouteResponse{
			StatusCode: api.ReturnStatusCodes_INVALID_FIELD_ERROR,
			Error:      e.Error(),
			Route:      nil, // Return nothing, since nothing was deleted.
		}, nil
	}

	var delRoute nl.Route

	// Check that the route exists, otherwise abort and return invalid element error.
	if r, e := nl.RouteGet(nil); r == nil || e != nil { // Need to update the destination with a value
		return &api.DeleteStaticRouteResponse{
			StatusCode: api.ReturnStatusCodes_NON_EXISTENT_ELEMENT,
			Error:      e.Error(),
			Route:      nil, // Return nothing, since nothing was deleted.
		}, nil
	} else {
		delRoute = r[0] // TODO need to figure out which route we are looking to delete from the returned options.
	}

	// Now try to delete the route.
	if e := nl.RouteDel(converters.ConvertAPIRouteToNetlinkRouteNew(req.GetRoute())); e != nil {
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

func (s *NetlinkNodeServer) UpdateStaticRoute(ctx context.Context, req *api.UpdateStaticRouteRequest) (resp *api.UpdateStaticRouteResponse, err error) {

	// Validate the incoming route request.
	if e := validators.ValidateRoute(req.GetRoute()); e != nil {
		return &api.UpdateStaticRouteResponse{
			StatusCode: api.ReturnStatusCodes_INVALID_FIELD_ERROR,
			Error:      e.Error(),
		}, nil
	}

	// Check that the route exists, otherwise abort and return invalid element error.
	if r, e := nl.RouteGet(nil); r == nil || e != nil { // Need to update the destination with a value
		return &api.UpdateStaticRouteResponse{
			StatusCode: api.ReturnStatusCodes_NON_EXISTENT_ELEMENT,
			Error:      e.Error(),
		}, nil
	}

	// Run a replace on the route.
	if e := nl.RouteReplace(converters.ConvertAPIRouteToNetlinkRouteNew(req.GetRoute())); e != nil {
		return &api.UpdateStaticRouteResponse{
			StatusCode: api.ReturnStatusCodes_INTERNAL_ERROR,
			Error:      e.Error(),
		}, nil
	}

	return &api.UpdateStaticRouteResponse{
		StatusCode: api.ReturnStatusCodes_OK,
		Error:      "",
	}, nil

}

func (s *NetlinkNodeServer) GetRoute(ctx context.Context, req *api.GetRouteRequest) (resp *api.GetRouteResponse, err error) {
	return nil, nil
}

func (s *NetlinkNodeServer) GetAllRoutes(ctx context.Context, req *api.GetAllRoutesRequest) (resp *api.GetAllRoutesResponse, err error) {

	if r, e := nl.RouteList(nil, 0); e != nil { // TODO: Need to check that family value to see what should be put
		return &api.GetAllRoutesResponse{
			StatusCode: api.ReturnStatusCodes_NON_EXISTENT_ELEMENT,
			Error:      e.Error(),
			Routes:     nil, // Return nil, since nothing was returned
		}, nil
	} else {
		var retRoutes []*api.Route

		// Convert all the returned routes to API route object for return.
		for _, route := range r {
			retRoutes = append(retRoutes, converters.ConvertNetlinkRouteNewToAPIRoute(&route))
		}

		return &api.GetAllRoutesResponse{
			StatusCode: api.ReturnStatusCodes_OK,
			Error:      "",
			Routes:     retRoutes,
		}, nil
	}

}
