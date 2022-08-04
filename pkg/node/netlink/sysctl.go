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
	"fmt"
	"strconv"

	api "github.com/fire833/morfic/pkg/ipcapi/v1alpha1"

	sysctl "github.com/lorenzosaino/go-sysctl"
)

func (s *NetlinkNodeServer) GetSysctl(ctx context.Context, req *api.GetSysctlRequest) (resp *api.GetSysctlResponse, err error) {

	// Set up a new client for each request for now.
	client, e := sysctl.NewClient(sysctl.DefaultPath)
	if e != nil {
		return &api.GetSysctlResponse{
			Error:      e.Error(),
			Value:      0,
			StatusCode: api.ReturnStatusCodes_INTERNAL_ERROR,
		}, nil
	}

	// Get the sysctl
	val, e1 := client.Get(req.Name)
	if e1 != nil {
		return &api.GetSysctlResponse{
			Error:      e1.Error(),
			Value:      0,
			StatusCode: api.ReturnStatusCodes_INVALID_FIELD_ERROR,
		}, nil
	}

	// Calculate an integer from the sysctl return string.
	if value, e2 := strconv.ParseInt(val, 10, 64); e != nil {
		return &api.GetSysctlResponse{
			Error:      e2.Error(),
			Value:      0,
			StatusCode: api.ReturnStatusCodes_INVALID_FIELD_ERROR,
		}, nil
	} else {
		return &api.GetSysctlResponse{
			Error:      "",
			Value:      value,
			StatusCode: api.ReturnStatusCodes_OK,
		}, nil
	}
}

func (s *NetlinkNodeServer) SetSysctl(ctx context.Context, req *api.SetSysctlRequest) (resp *api.SetSysctlResponse, err error) {

	// Set up a new client for each request for now.
	client, e := sysctl.NewClient(sysctl.DefaultPath)
	if e != nil {
		return &api.SetSysctlResponse{
			Error:      e.Error(),
			Value:      0,
			StatusCode: api.ReturnStatusCodes_INTERNAL_ERROR,
		}, nil
	}

	// Set the sysctl value and return the result.
	if e1 := client.Set(req.Name, fmt.Sprintf("%d", req.Value)); e != nil {
		return &api.SetSysctlResponse{
			Error:      e1.Error(),
			Value:      req.Value,
			StatusCode: api.ReturnStatusCodes_INVALID_FIELD_ERROR,
		}, nil
	} else {
		return &api.SetSysctlResponse{
			Error:      "",
			Value:      req.Value,
			StatusCode: api.ReturnStatusCodes_OK,
		}, nil
	}
}
