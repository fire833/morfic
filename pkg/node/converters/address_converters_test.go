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
	"net"
	"reflect"
	"testing"

	api "github.com/fire833/morfic/pkg/ipcapi/v1alpha1"
	"github.com/jsimonetti/rtnetlink"
	"github.com/vishvananda/netlink"
)

func TestConvertAPIAddrtoNetlinkAddrNew(t *testing.T) {
	type args struct {
		addr *api.LinkAddress
	}
	tests := []struct {
		name string
		args args
		want *netlink.Addr
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertAPIAddrtoNetlinkAddrNew(tt.args.addr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertAPIAddrtoNetlinkAddrNew() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertAPIAddrtoNetlinkAddr(t *testing.T) {
	type args struct {
		addr *api.IPCIDR
	}
	tests := []struct {
		name string
		args args
		want *rtnetlink.AddressMessage
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertAPIAddrtoNetlinkAddr(tt.args.addr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertAPIAddrtoNetlinkAddr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertNetlinkAddrNewToAPIAddr(t *testing.T) {
	type args struct {
		addr *netlink.Addr
	}
	tests := []struct {
		name string
		args args
		want *api.LinkAddress
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertNetlinkAddrNewToAPIAddr(tt.args.addr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertNetlinkAddrNewToAPIAddr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertNetlinkAddrToAPIAddr(t *testing.T) {
	type args struct {
		addr *rtnetlink.AddressMessage
	}
	tests := []struct {
		name string
		args args
		want *api.LinkAddress
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertNetlinkAddrToAPIAddr(tt.args.addr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertNetlinkAddrToAPIAddr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertAPICDRtoNetIP(t *testing.T) {
	type args struct {
		cidr *api.IPCIDR
	}
	tests := []struct {
		name string
		args args
		want *net.IPNet
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertAPICDRtoNetIP(tt.args.cidr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertAPICDRtoNetIP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertAPIAddrtoNetIP(t *testing.T) {
	type args struct {
		addr *api.IPAddress
	}
	tests := []struct {
		name string
		args args
		want net.IP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertAPIAddrtoNetIP(tt.args.addr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertAPIAddrtoNetIP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertAPIMasktoNetMask(t *testing.T) {
	type args struct {
		mask *api.IPMask
	}
	tests := []struct {
		name string
		args args
		want net.IPMask
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertAPIMasktoNetMask(tt.args.mask); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertAPIMasktoNetMask() = %v, want %v", got, tt.want)
			}
		})
	}
}
