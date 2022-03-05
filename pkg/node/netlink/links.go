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

	api "github.com/fire833/vroute/pkg/apis/ipcapi/v1alpha1"
	"github.com/fire833/vroute/pkg/node/validators"
	"github.com/jsimonetti/rtnetlink"

	c "github.com/fire833/vroute/pkg/node/netlink/connections"
)

var (
	unableToAcquireLinkError error = errors.New("unable to acquire socket file descriptor for connection")
	linkNotFoundError        error = errors.New("unable to find matching link on the host")
	duplicateLinkError       error = errors.New("link with same name already exists on the host")
)

func (s *NetlinkNodeServer) CreateLink(ctx context.Context, req *api.CreateLinkRequest) (resp *api.CreateLinkResponse, e error) {

	// Validate the incoming link from the request.
	if e := validators.ValidateLink(req.Link); e != nil {
		return &api.CreateLinkResponse{
			StatusCode: api.ReturnStatusCodes_INVALID_FIELD_ERROR,
			Error:      e.Error(),
		}, nil
	}

	// Check if a duplicate link already exists.
	if _, e := net.InterfaceByName(req.Link.Name); e == nil {
		return &api.CreateLinkResponse{
			StatusCode: api.ReturnStatusCodes_DUPLICATE_ELEMENT_ERROR,
			Error:      duplicateLinkError.Error(),
		}, nil
	}

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

	// Get the mac addresses.
	hwaddr, _ := net.ParseMAC(string(req.Link.Mac.Address.Address))
	bdaddr, _ := net.ParseMAC(string(req.Link.Mac.BroadcastAddress.Address))

	// Create the new link through netlink connection.
	if e := conn.Link.New(&rtnetlink.LinkMessage{
		Attributes: &rtnetlink.LinkAttributes{
			MTU:       req.Link.Mtu,
			Name:      req.Link.Name,
			Address:   hwaddr,
			Broadcast: bdaddr,
		},
		Family: 0, // TODO: Need to look into this
		Index:  req.Link.Index,
	}); e != nil {

		resp := &api.CreateLinkResponse{
			StatusCode: api.ReturnStatusCodes_INTERNAL_ERROR,
			Error:      e.Error(),
		}

		return resp, nil
	}

	// Return that the link was created correctly.
	return &api.CreateLinkResponse{
		StatusCode: api.ReturnStatusCodes_OK,
		Error:      "",
	}, nil
}

func (s *NetlinkNodeServer) UpdateLink(ctx context.Context, req *api.UpdateLinkRequest) (resp *api.UpdateLinkResponse, e error) {

	// Validate the incoming link from the request.
	if e := validators.ValidateLink(req.Link); e != nil {
		return &api.UpdateLinkResponse{
			StatusCode: api.ReturnStatusCodes_INVALID_FIELD_ERROR,
			Error:      e.Error(),
		}, nil
	}

	// Check if the link doesn't already exist.
	if _, e := net.InterfaceByName(req.Link.Name); e == nil {
		return &api.UpdateLinkResponse{
			StatusCode: api.ReturnStatusCodes_NON_EXISTENT_ELEMENT,
			Error:      duplicateLinkError.Error(),
		}, nil
	}

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

	// TODO need to update the link still

	// Return that the link was updated correctly.
	return &api.UpdateLinkResponse{
		StatusCode: api.ReturnStatusCodes_OK,
		Error:      "",
	}, nil
}

func (s *NetlinkNodeServer) DeleteLink(ctx context.Context, req *api.DeleteLinkRequest) (resp *api.DeleteLinkResponse, e error) {

	// Validate the incoming link name to make sure it's valid.
	if e := validators.ValidateInterfaceName(req.Name); e == nil {
		return &api.DeleteLinkResponse{
			StatusCode: api.ReturnStatusCodes_NON_EXISTENT_ELEMENT,
			Error:      e.Error(),
		}, nil
	}

	var i *net.Interface // Interface with the index of the name.

	// Get the index for this address name so it can be deleted.
	if iface, e := net.InterfaceByName(req.Name); e != nil {
		return &api.DeleteLinkResponse{
			StatusCode: api.ReturnStatusCodes_NON_EXISTENT_ELEMENT,
			Error:      "",
		}, nil
	} else {
		i = iface
	}

	// Acquire connection to netlink.
	conn, e1 := c.NetlinkPool.GetConn()
	defer c.NetlinkPool.ReturnConn(conn) // Try and return the fd after returning.

	if e1 != nil {

		resp := &api.DeleteLinkResponse{
			StatusCode: api.ReturnStatusCodes_INTERNAL_ERROR,
			Error:      e1.Error(),
		}

		return resp, nil

	}

	var msg *rtnetlink.LinkMessage

	if m, e := conn.Link.Get(uint32(i.Index)); e != nil {
		return &api.DeleteLinkResponse{
			StatusCode:  api.ReturnStatusCodes_INTERNAL_ERROR,
			Error:       e.Error(),
			DeletedLink: nil,
		}, nil
	} else {
		msg = &m
	}

	if e := conn.Link.Delete(uint32(i.Index)); e != nil {

		return &api.DeleteLinkResponse{
			StatusCode:  api.ReturnStatusCodes_INTERNAL_ERROR,
			Error:       e.Error(),
			DeletedLink: nil,
		}, nil

	}

	return &api.DeleteLinkResponse{
		StatusCode: api.ReturnStatusCodes_OK,
		Error:      "",
		DeletedLink: &api.Link{
			Name:    msg.Attributes.Name,
			Mtu:     uint32(i.MTU),
			Address: nil, // Just keep nil for now I guess
			Index:   uint32(i.Index),
			Attributes: &api.LinkAttributes{
				Up:         false, // link is obviouslyno longer up.
				ArpEnabled: false,
				Multicast:  false,
				Dynamic:    false,
			},
		},
	}, nil

}

func (s *NetlinkNodeServer) GetLink(ctx context.Context, req *api.GetLinkRequest) (resp *api.GetLinkResponse, e error) {
	return nil, nil
}

func (s *NetlinkNodeServer) GetAllLinks(ctx context.Context, req *api.GetAllLinksRequest) (resp *api.GetAllLinksResponse, e error) {
	return nil, nil
}
