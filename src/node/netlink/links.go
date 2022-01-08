/*
*	Copyright (C) 2021  Kendall Tauser
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
	"net"

	api "github.com/fire833/vroute/src/api/ipcapi/v1alpha1"
	"github.com/fire833/vroute/src/node/validators"
	"github.com/jsimonetti/rtnetlink"
)

func (s *NetlinkNodeServer) CreateLink(ctx context.Context, req *api.CreateLinkRequest) (resp *api.CreateLinkResponse, e error) {
	linkcodes := make(map[int]error, len(req.Link)) // Map each of the link indexes to their respective return codes/error message

	// Validate the request before grabbing a file descriptor to write
	for n, link := range req.Link {
		linkcodes[n] = validators.ValidateLink(link) // Validate each link before trying to update.
	}

	conn, e1 := netlinkPool.GetConn()
	defer netlinkPool.ReturnConn(conn) // Try and return the fd after returning.
	if e1 != nil {                     // If there's an error, then there is something seriously wrong at this point with
		// getting connections, so just return the response.

		resp := &api.CreateLinkResponse{
			StatusCode:        api.ReturnStatusCodes_INTERNAL_ERROR,
			UnsuccessfulLinks: req.Link, // All the links failed
			SuccessLinks:      nil,
		}

		return resp, nil // Return the response.
	}

	for n, link := range req.Link {
		if linkcodes[n] != nil {
			continue
		}

		conn.Link.New(&rtnetlink.LinkMessage{
			Attributes: &rtnetlink.LinkAttributes{
				MTU:  link.Mtu,
				Name: link.Name,
			},
			Family: 0, // Need to look into this
			Index:  link.Index,
		})

	}

	return nil, nil
}

func (s *NetlinkNodeServer) UpdateLink(ctx context.Context, req *api.UpdateLinkRequest) (resp *api.UpdateLinkResponse, e error) {
	return nil, nil
}

func (s *NetlinkNodeServer) DeleteLink(ctx context.Context, req *api.DeleteLinkRequest) (resp *api.DeleteLinkResponse, e error) {
	for _, iface := range req.Name {
		if i, e := net.InterfaceByName(iface); e != nil {

		}
	}

	return nil, nil
}

func (s *NetlinkNodeServer) GetLink(ctx context.Context, req *api.GetLinkRequest) (resp *api.GetLinkResponse, e error) {
	return nil, nil
}

func (s *NetlinkNodeServer) GetAllLinks(ctx context.Context, req *api.GetAllLinksRequest) (resp *api.GetAllLinksResponse, e error) {
	return nil, nil
}
