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

package ip_routes

import (
	"context"

	"github.com/fire833/vroute/src/api/ipcapi/v1alpha1"
)

func (s *CMDNodeServer) CreateLink(ctx context.Context, req *v1alpha1.CreateLinkRequest) (resp *v1alpha1.CreateLinkResponse, e error) {
	return nil, nil
}

func (s *CMDNodeServer) UpdateLink(ctx context.Context, req *v1alpha1.UpdateLinkRequest) (resp *v1alpha1.UpdateLinkResponse, e error) {
	return nil, nil
}

func (s *CMDNodeServer) DeleteLink(ctx context.Context, req *v1alpha1.DeleteLinkRequest) (resp *v1alpha1.DeleteLinkResponse, e error) {
	return nil, nil
}

func (s *CMDNodeServer) GetLink(ctx context.Context, req *v1alpha1.GetLinkRequest) (resp *v1alpha1.GetLinkResponse, e error) {
	return nil, nil
}

func (s *CMDNodeServer) GetAllLinks(ctx context.Context, req *v1alpha1.GetAllLinksRequest) (resp *v1alpha1.GetAllLinksResponse, e error) {
	return nil, nil
}
