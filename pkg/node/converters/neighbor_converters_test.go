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

func TestConvertAPINeighborToNetlinkNeighborNew(t *testing.T) {
	type args struct {
		neigh *api.Neighbor
	}
	tests := []struct {
		name string
		args args
		want *netlink.Neigh
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertAPINeighborToNetlinkNeighborNew(tt.args.neigh); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertAPINeighborToNetlinkNeighborNew() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertAPINeighborToNetlinkNeighbor(t *testing.T) {
	type args struct {
		neigh *api.Neighbor
	}
	tests := []struct {
		name string
		args args
		want *rtnetlink.NeighMessage
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertAPINeighborToNetlinkNeighbor(tt.args.neigh); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertAPINeighborToNetlinkNeighbor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertNetlinkNeighborNewToAPINeighbor(t *testing.T) {
	type args struct {
		neigh *netlink.Neigh
	}
	tests := []struct {
		name string
		args args
		want *api.Neighbor
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertNetlinkNeighborNewToAPINeighbor(tt.args.neigh); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertNetlinkNeighborNewToAPINeighbor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertNetlinkNeighborToAPINeighbor(t *testing.T) {
	type args struct {
		neigh *rtnetlink.NeighMessage
	}
	tests := []struct {
		name string
		args args
		want *api.Neighbor
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertNetlinkNeighborToAPINeighbor(tt.args.neigh); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertNetlinkNeighborToAPINeighbor() = %v, want %v", got, tt.want)
			}
		})
	}
}
