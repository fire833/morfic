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

package validation

import (
	"reflect"
	"testing"

	"github.com/fire833/morfic/pkg/apis/net"
)

func TestValidateAddressList(t *testing.T) {
	type args struct {
		in *net.AddressList
	}
	tests := []struct {
		name string
		args args
		want []error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateAddressList(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValidateAddressList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateAddress(t *testing.T) {
	type args struct {
		in *net.Address
	}
	tests := []struct {
		name string
		args args
		want []error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateAddress(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValidateAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateAddressSpec(t *testing.T) {
	type args struct {
		in *net.AddressSpec
	}
	tests := []struct {
		name string
		args args
		want []error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateAddressSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValidateAddressSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateIPSpec(t *testing.T) {
	type args struct {
		in *net.IPSpec
	}
	tests := []struct {
		name string
		args args
		want []error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateIPSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValidateIPSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}
