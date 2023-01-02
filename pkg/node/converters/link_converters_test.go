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

package converters

import (
	"reflect"
	"testing"

	api "github.com/fire833/morfic/pkg/ipcapi/v1alpha1"
	"github.com/jsimonetti/rtnetlink"
	"github.com/vishvananda/netlink"
)

func TestConvertAPILinkToNetlinkLinkNew(t *testing.T) {
	type args struct {
		link *api.Link
	}
	tests := []struct {
		name string
		args args
		want *netlink.Link
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertAPILinkToNetlinkLinkNew(tt.args.link); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertAPILinkToNetlinkLinkNew() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertNetlinkLinkNewToAPILink(t *testing.T) {
	type args struct {
		link *netlink.Link
	}
	tests := []struct {
		name string
		args args
		want *api.Link
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertNetlinkLinkNewToAPILink(tt.args.link); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertNetlinkLinkNewToAPILink() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertAPILinkToNetlinkLink(t *testing.T) {
	type args struct {
		link *api.Link
	}
	tests := []struct {
		name string
		args args
		want *rtnetlink.LinkMessage
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertAPILinkToNetlinkLink(tt.args.link); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertAPILinkToNetlinkLink() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertNetlinkLinkAPILink(t *testing.T) {
	type args struct {
		link *rtnetlink.LinkMessage
	}
	tests := []struct {
		name string
		args args
		want *api.Link
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertNetlinkLinkAPILink(tt.args.link); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertNetlinkLinkAPILink() = %v, want %v", got, tt.want)
			}
		})
	}
}
