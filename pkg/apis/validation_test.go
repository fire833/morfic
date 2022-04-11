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

package apis

import (
	"reflect"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestValidateTypeMeta(t *testing.T) {
	type args struct {
		in                *metav1.TypeMeta
		desiredKind       string
		desiredAPIVersion string
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
			if got := ValidateTypeMeta(tt.args.in, tt.args.desiredKind, tt.args.desiredAPIVersion); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValidateTypeMeta() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateObjectMeta(t *testing.T) {
	type args struct {
		in *metav1.ObjectMeta
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
			if got := ValidateObjectMeta(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValidateObjectMeta() = %v, want %v", got, tt.want)
			}
		})
	}
}
