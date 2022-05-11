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
	"errors"
	"net"

	api "github.com/fire833/morfic/pkg/apis/ipcapi/v1alpha1"
	"github.com/fire833/morfic/pkg/node/converters"
	"github.com/fire833/morfic/pkg/node/locks"
	"github.com/fire833/morfic/pkg/node/validators"
	"github.com/jsimonetti/rtnetlink"

	c "github.com/fire833/morfic/pkg/node/connections"
)

var (
	unableToAcquireLinkError error = errors.New("unable to acquire socket file descriptor for connection")
	linkNotFoundError        error = errors.New("unable to find matching link on the host")
	duplicateLinkError       error = errors.New("link with same name already exists on the host")
)

func (s *NetlinkNodeServer) CreateLink(ctx context.Context, req *api.CreateLinkRequest) (resp *api.CreateLinkResponse, err error) {
	ll := locks.LinkLock

	// Validate the incoming link from the request.
	if e := validators.ValidateLink(req.Link); e != nil {
		return &api.CreateLinkResponse{
			StatusCode: api.ReturnStatusCodes_INVALID_FIELD_ERROR,
			Error:      e.Error(),
		}, nil
	}

	ll.RLock()

	// Check if a duplicate link already exists.
	if _, e := net.InterfaceByName(req.Link.Name); e == nil {
		ll.RUnlock()

		return &api.CreateLinkResponse{
			StatusCode: api.ReturnStatusCodes_DUPLICATE_ELEMENT_ERROR,
			Error:      duplicateLinkError.Error(),
		}, nil
	}

	ll.RUnlock()

	// Acquire connection to netlink.
	conn, e1 := c.NetlinkPool.GetConn()
	defer c.NetlinkPool.ReturnConn(conn) // Try and return the fd after returning.
	if e1 != nil {                       // If there's an error, then there is something seriously wrong at this point with

		resp := &api.CreateLinkResponse{
			StatusCode: api.ReturnStatusCodes_INTERNAL_ERROR,
			Error:      e1.Error(),
		}

		return resp, nil
	}

	ll.Lock()

	// Create the new link through netlink connection.
	if e := conn.Link.New(converters.ConvertAPILinkToNetlinkLink(req.Link)); e != nil {
		ll.Unlock()

		return &api.CreateLinkResponse{
			StatusCode: api.ReturnStatusCodes_INTERNAL_ERROR,
			Error:      e.Error(),
		}, nil
	}

	ll.Unlock()

	// Return that the link was created correctly.
	return &api.CreateLinkResponse{
		StatusCode: api.ReturnStatusCodes_OK,
		Error:      "",
	}, nil

}

func (s *NetlinkNodeServer) UpdateLink(ctx context.Context, req *api.UpdateLinkRequest) (resp *api.UpdateLinkResponse, err error) {
	ll := locks.LinkLock

	// Validate the incoming link from the request.
	if e := validators.ValidateLink(req.Link); e != nil {
		return &api.UpdateLinkResponse{
			StatusCode: api.ReturnStatusCodes_INVALID_FIELD_ERROR,
			Error:      e.Error(),
		}, nil
	}

	ll.RLock()

	// Check if the link doesn't already exist.
	if _, e := net.InterfaceByName(req.Link.Name); e == nil {
		ll.RUnlock()

		return &api.UpdateLinkResponse{
			StatusCode: api.ReturnStatusCodes_NON_EXISTENT_ELEMENT,
			Error:      duplicateLinkError.Error(),
		}, nil
	}

	ll.RUnlock()

	// Acquire connection to netlink.
	conn, e1 := c.NetlinkPool.GetConn()
	defer c.NetlinkPool.ReturnConn(conn) // Try and return the fd after returning.

	if e1 != nil {

		resp := &api.UpdateLinkResponse{
			StatusCode: api.ReturnStatusCodes_INTERNAL_ERROR,
			Error:      e1.Error(),
		}

		return resp, nil

	}

	ll.Lock()

	if e := conn.Link.Set(converters.ConvertAPILinkToNetlinkLink(req.Link)); e != nil {
		ll.Unlock()

		return &api.UpdateLinkResponse{
			StatusCode: api.ReturnStatusCodes_INTERNAL_ERROR,
			Error:      e.Error(),
		}, nil
	}

	ll.Unlock()

	// Return that the link was updated correctly.
	return &api.UpdateLinkResponse{
		StatusCode: api.ReturnStatusCodes_OK,
		Error:      "",
	}, nil

}

func (s *NetlinkNodeServer) DeleteLink(ctx context.Context, req *api.DeleteLinkRequest) (resp *api.DeleteLinkResponse, err error) {
	ll := locks.LinkLock

	// Validate the incoming link name to make sure it's valid.
	if e := validators.ValidateInterfaceName(req.Name); e == nil {
		return &api.DeleteLinkResponse{
			StatusCode:  api.ReturnStatusCodes_INVALID_FIELD_ERROR,
			Error:       e.Error(),
			DeletedLink: nil, // Return nothing, since nothing was deleted.
		}, nil
	}

	var i *net.Interface // Interface with the index of the name.

	ll.RLock()

	// Get the index for this address name so it can be deleted.
	if iface, e := net.InterfaceByName(req.Name); e != nil {
		ll.RUnlock()

		return &api.DeleteLinkResponse{
			StatusCode:  api.ReturnStatusCodes_NON_EXISTENT_ELEMENT,
			Error:       "",
			DeletedLink: nil, // Return nothing, since nothing was deleted.
		}, nil
	} else {
		ll.RUnlock()

		i = iface
	}

	// Acquire connection to netlink.
	conn, e1 := c.NetlinkPool.GetConn()
	defer c.NetlinkPool.ReturnConn(conn) // Try and return the fd after returning.

	if e1 != nil {

		resp := &api.DeleteLinkResponse{
			StatusCode:  api.ReturnStatusCodes_INTERNAL_ERROR,
			Error:       e1.Error(),
			DeletedLink: nil, // Return nothing, since nothing was deleted.
		}

		return resp, nil

	}

	var msg *rtnetlink.LinkMessage

	ll.RLock()

	if m, e := conn.Link.Get(uint32(i.Index)); e != nil {
		ll.RUnlock()

		return &api.DeleteLinkResponse{
			StatusCode:  api.ReturnStatusCodes_INTERNAL_ERROR,
			Error:       e.Error(),
			DeletedLink: nil, // Return nothing, since nothing was deleted.
		}, nil
	} else {
		ll.RUnlock()

		msg = &m
	}

	ll.Lock()

	if e := conn.Link.Delete(uint32(i.Index)); e != nil {
		ll.Unlock()

		return &api.DeleteLinkResponse{
			StatusCode:  api.ReturnStatusCodes_INTERNAL_ERROR,
			Error:       e.Error(),
			DeletedLink: nil, // Return nothing, since nothing was deleted.
		}, nil
	}

	ll.Unlock()

	return &api.DeleteLinkResponse{
		StatusCode:  api.ReturnStatusCodes_OK,
		Error:       "",
		DeletedLink: converters.ConvertNetlinkLinkAPILink(msg),
	}, nil

}

func (s *NetlinkNodeServer) GetLink(ctx context.Context, req *api.GetLinkRequest) (resp *api.GetLinkResponse, err error) {
	return nil, nil
}

func (s *NetlinkNodeServer) GetAllLinks(ctx context.Context, req *api.GetAllLinksRequest) (resp *api.GetAllLinksResponse, err error) {
	return nil, nil
}
