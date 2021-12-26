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

func (s *IPNodeServer) GetNeighbor(ctx context.Context, req *v1alpha1.GetNeighborRequest) (resp *v1alpha1.GetNeighborResponse, e error) {
	return nil, nil
}

func (s *IPNodeServer) GetAllNeighbors(ctx context.Context, req *v1alpha1.GetAllNeighborsRequest) (resp *v1alpha1.GetAllNeighborsResponse, e error) {
	return nil, nil
}

func (s *IPNodeServer) CreateNeighbor(ctx context.Context, req *v1alpha1.CreateNeighborRequest) (resp *v1alpha1.CreateNeighborResponse, e error) {
	return nil, nil
}

func (s *IPNodeServer) DeleteNeighbor(ctx context.Context, req *v1alpha1.DeleteNeighborRequest) (resp *v1alpha1.DeleteNeighborRequest, e error) {
	return nil, nil
}

func (s *IPNodeServer) UpdateNeighbor(ctx context.Context, req *v1alpha1.UpdateNeighborRequest) (resp *v1alpha1.UpdateNeighborResponse, e error) {
	return nil, nil
}
