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

package converters

import (
	"reflect"
	"testing"

	api "github.com/fire833/morfic/pkg/apis/ipcapi/v1alpha1"
	"github.com/jsimonetti/rtnetlink"
	"github.com/vishvananda/netlink"
)

func TestConvertAPIRouteToNetlinkRouteNew(t *testing.T) {
	type args struct {
		route *api.Route
	}
	tests := []struct {
		name string
		args args
		want *netlink.Route
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertAPIRouteToNetlinkRouteNew(tt.args.route); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertAPIRouteToNetlinkRouteNew() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertAPIRouteToNetlinkRoute(t *testing.T) {
	type args struct {
		route *api.Route
	}
	tests := []struct {
		name string
		args args
		want *rtnetlink.RouteMessage
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertAPIRouteToNetlinkRoute(tt.args.route); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertAPIRouteToNetlinkRoute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertNetlinkRouteNewToAPIRoute(t *testing.T) {
	type args struct {
		route *netlink.Route
	}
	tests := []struct {
		name string
		args args
		want *api.Route
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertNetlinkRouteNewToAPIRoute(tt.args.route); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertNetlinkRouteNewToAPIRoute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertNetlinkRouteToAPIRoute(t *testing.T) {
	type args struct {
		route *rtnetlink.RouteMessage
	}
	tests := []struct {
		name string
		args args
		want *api.Route
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertNetlinkRouteToAPIRoute(tt.args.route); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertNetlinkRouteToAPIRoute() = %v, want %v", got, tt.want)
			}
		})
	}
}
