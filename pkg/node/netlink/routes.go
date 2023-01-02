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
	"github.com/fire833/morfic/pkg/node/converters"
	"github.com/fire833/morfic/pkg/node/locks"
	"github.com/fire833/morfic/pkg/node/validators"
	nl "github.com/vishvananda/netlink"
)

func (s *NetlinkNodeServer) CreateStaticRoute(ctx context.Context, req *api.CreateStaticRouteRequest) (resp *api.CreateStaticRouteResponse, err error) {
	rl := locks.RouteLock

	// Validate the incoming route request.
	if e := validators.ValidateRoute(req.GetRoute()); e != nil {
		return &api.CreateStaticRouteResponse{
			StatusCode: api.ReturnStatusCodes_INVALID_FIELD_ERROR,
			Error:      e.Error(),
		}, nil
	}

	rl.Lock()

	// Add the route after converting it to avalid form for netlink.
	if e := nl.RouteAdd(converters.ConvertAPIRouteToNetlinkRouteNew(req.GetRoute())); e != nil {
		rl.Unlock()

		return &api.CreateStaticRouteResponse{
			StatusCode: api.ReturnStatusCodes_INTERNAL_ERROR,
			Error:      e.Error(),
		}, nil
	}

	rl.Unlock()

	return &api.CreateStaticRouteResponse{
		StatusCode: api.ReturnStatusCodes_OK,
		Error:      "",
	}, nil

}

func (s *NetlinkNodeServer) DeleteStaticRoute(ctx context.Context, req *api.DeleteStaticRouteRequest) (resp *api.DeleteStaticRouteResponse, err error) {
	rl := locks.RouteLock

	// Validate the incoming route request.
	if e := validators.ValidateRoute(req.GetRoute()); e != nil {
		return &api.DeleteStaticRouteResponse{
			StatusCode: api.ReturnStatusCodes_INVALID_FIELD_ERROR,
			Error:      e.Error(),
			Route:      nil, // Return nothing, since nothing was deleted.
		}, nil
	}

	var delRoute nl.Route

	rl.RLock()

	// Check that the route exists, otherwise abort and return invalid element error.
	if r, e := nl.RouteGet(nil); r == nil || e != nil { // Need to update the destination with a value
		rl.RUnlock()

		return &api.DeleteStaticRouteResponse{
			StatusCode: api.ReturnStatusCodes_NON_EXISTENT_ELEMENT,
			Error:      e.Error(),
			Route:      nil, // Return nothing, since nothing was deleted.
		}, nil
	} else {
		rl.RUnlock()

		delRoute = r[0] // TODO need to figure out which route we are looking to delete from the returned options.
	}

	rl.Lock()

	// Now try to delete the route.
	if e := nl.RouteDel(converters.ConvertAPIRouteToNetlinkRouteNew(req.GetRoute())); e != nil {
		rl.Unlock()

		return &api.DeleteStaticRouteResponse{
			StatusCode: api.ReturnStatusCodes_INTERNAL_ERROR,
			Error:      e.Error(),
			Route:      nil, // Return nothing, since nothing was deleted.
		}, nil
	}

	rl.Unlock()

	return &api.DeleteStaticRouteResponse{
		StatusCode: api.ReturnStatusCodes_OK,
		Error:      "",
		Route:      converters.ConvertNetlinkRouteNewToAPIRoute(&delRoute),
	}, nil

}

func (s *NetlinkNodeServer) UpdateStaticRoute(ctx context.Context, req *api.UpdateStaticRouteRequest) (resp *api.UpdateStaticRouteResponse, err error) {
	rl := locks.RouteLock

	// Validate the incoming route request.
	if e := validators.ValidateRoute(req.GetRoute()); e != nil {
		return &api.UpdateStaticRouteResponse{
			StatusCode: api.ReturnStatusCodes_INVALID_FIELD_ERROR,
			Error:      e.Error(),
		}, nil
	}

	rl.RLock()

	// Check that the route exists, otherwise abort and return invalid element error.
	if r, e := nl.RouteGet(nil); r == nil || e != nil { // Need to update the destination with a value
		rl.RUnlock()

		return &api.UpdateStaticRouteResponse{
			StatusCode: api.ReturnStatusCodes_NON_EXISTENT_ELEMENT,
			Error:      e.Error(),
		}, nil
	}

	rl.RUnlock()
	rl.Lock()

	// Run a replace on the route.
	if e := nl.RouteReplace(converters.ConvertAPIRouteToNetlinkRouteNew(req.GetRoute())); e != nil {
		rl.Unlock()

		return &api.UpdateStaticRouteResponse{
			StatusCode: api.ReturnStatusCodes_INTERNAL_ERROR,
			Error:      e.Error(),
		}, nil
	}

	rl.Unlock()

	return &api.UpdateStaticRouteResponse{
		StatusCode: api.ReturnStatusCodes_OK,
		Error:      "",
	}, nil

}

func (s *NetlinkNodeServer) GetRoute(ctx context.Context, req *api.GetRouteRequest) (resp *api.GetRouteResponse, err error) {
	return nil, nil
}

func (s *NetlinkNodeServer) GetAllRoutes(ctx context.Context, req *api.GetAllRoutesRequest) (resp *api.GetAllRoutesResponse, err error) {
	rl := locks.RouteLock

	rl.RLock()

	if r, e := nl.RouteList(nil, 0); e != nil { // TODO: Need to check that family value to see what should be put
		rl.RUnlock()

		return &api.GetAllRoutesResponse{
			StatusCode: api.ReturnStatusCodes_NON_EXISTENT_ELEMENT,
			Error:      e.Error(),
			Routes:     nil, // Return nil, since nothing was returned
		}, nil
	} else {
		rl.RUnlock()

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
