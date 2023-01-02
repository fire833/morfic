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

func TestValidateLinkList(t *testing.T) {
	type args struct {
		in *net.LinkList
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
			if got := ValidateLinkList(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValidateLinkList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateLink(t *testing.T) {
	type args struct {
		in *net.Link
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
			if got := ValidateLink(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValidateLink() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateLinkSpec(t *testing.T) {
	type args struct {
		in *net.LinkSpec
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
			if got := ValidateLinkSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValidateLinkSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}
