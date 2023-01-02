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

package comparator

import (
	"testing"

	api "github.com/fire833/morfic/pkg/ipcapi/v1alpha1"
)

func TestCompareLinkAddress(t *testing.T) {
	type args struct {
		addr1 *api.LinkAddress
		addr2 *api.LinkAddress
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompareLinkAddress(tt.args.addr1, tt.args.addr2); got != tt.want {
				t.Errorf("CompareLinkAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompareIPAddress(t *testing.T) {
	type args struct {
		addr1 *api.IPAddress
		addr2 *api.IPAddress
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompareIPAddress(tt.args.addr1, tt.args.addr2); got != tt.want {
				t.Errorf("CompareIPAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompareLinkMACAddress(t *testing.T) {
	type args struct {
		addr1 *api.LinkMACAddress
		addr2 *api.LinkMACAddress
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompareLinkMACAddress(tt.args.addr1, tt.args.addr2); got != tt.want {
				t.Errorf("CompareLinkMACAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompareMACAddress(t *testing.T) {
	type args struct {
		addr1 *api.MACAddress
		addr2 *api.MACAddress
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompareMACAddress(tt.args.addr1, tt.args.addr2); got != tt.want {
				t.Errorf("CompareMACAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompareCIDR(t *testing.T) {
	type args struct {
		cidr1 *api.IPCIDR
		cidr2 *api.IPCIDR
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompareCIDR(tt.args.cidr1, tt.args.cidr2); got != tt.want {
				t.Errorf("CompareCIDR() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompareSubnetMask(t *testing.T) {
	type args struct {
		mask1 *api.IPMask
		mask2 *api.IPMask
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompareSubnetMask(tt.args.mask1, tt.args.mask2); got != tt.want {
				t.Errorf("CompareSubnetMask() = %v, want %v", got, tt.want)
			}
		})
	}
}
