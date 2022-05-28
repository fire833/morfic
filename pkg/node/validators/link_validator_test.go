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

package validators

import (
	"testing"

	api "github.com/fire833/morfic/pkg/apis/ipcapi/v1alpha1"
)

func TestValidateLink(t *testing.T) {
	type args struct {
		in *api.Link
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateLink(tt.args.in); (err != nil) != tt.wantErr {
				t.Errorf("ValidateLink() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateInterfaceName(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateInterfaceName(tt.args.in); (err != nil) != tt.wantErr {
				t.Errorf("ValidateInterfaceName() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateLinkFilterElems(t *testing.T) {
	type args struct {
		in map[string]string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateLinkFilterElems(tt.args.in); (err != nil) != tt.wantErr {
				t.Errorf("ValidateLinkFilterElems() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
