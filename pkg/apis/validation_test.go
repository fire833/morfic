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

package apis

import (
	"fmt"
	"reflect"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestValidateTypeMeta(t *testing.T) {
	var e []error

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
		{
			name: "1",
			args: args{
				in: &metav1.TypeMeta{
					Kind:       "Route",
					APIVersion: "routes.vroute.io/v1alpha1",
				},
				desiredKind:       "Route",
				desiredAPIVersion: "routes.vroute.io/v1alpha1",
			},
			want: e,
		},
		{
			name: "2",
			args: args{
				in: &metav1.TypeMeta{
					Kind:       "Neighbor",
					APIVersion: "neighbors.vroute.io/v1alpha1",
				},
				desiredKind:       "Neighbor",
				desiredAPIVersion: "neighbors.vroute.io/v1alpha1",
			},
			want: e,
		},
		{
			name: "3",
			args: args{
				in: &metav1.TypeMeta{
					Kind:       "FirewallTable",
					APIVersion: "firewall.vroute.io/v1alpha1",
				},
				desiredKind:       "FirewallTable",
				desiredAPIVersion: "firewall.vroute.io/v1alpha1",
			},
			want: e,
		},
		{
			name: "4",
			args: args{
				in: &metav1.TypeMeta{
					Kind:       "Address",
					APIVersion: "addresses.vroute.io/v1alpha1",
				},
				desiredKind:       "Address",
				desiredAPIVersion: "addresses.vroute.io/v1alpha1",
			},
			want: e,
		},
		{
			name: "5",
			args: args{
				in: &metav1.TypeMeta{
					Kind:       "Address",
					APIVersion: "addresses.vroute.io/v1alpha1",
				},
				desiredKind:       "Route",
				desiredAPIVersion: "routes.vroute.io/v1alpha1",
			},
			want: []error{fmt.Errorf("desired kind %s does not match actual kind of object: %s", "Route", "Address"),
				fmt.Errorf("desired APIVersion %s does not match actual APIVersion of object: %s", "routes.vroute.io/v1alpha1", "addresses.vroute.io/v1alpha1")},
		},
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
