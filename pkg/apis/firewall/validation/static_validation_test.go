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

package validation

import (
	"reflect"
	"testing"

	"github.com/fire833/vroute/pkg/apis/firewall"
)

func TestValidateFirewallTableList(t *testing.T) {
	type args struct {
		in *firewall.TableList
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
			if got := ValidateFirewallTableList(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValidateFirewallTableList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateFirewallTable(t *testing.T) {
	type args struct {
		in *firewall.Table
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
			if got := ValidateFirewallTable(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValidateFirewallTable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateFirewallTableSpec(t *testing.T) {
	type args struct {
		in *firewall.FirewallTableSpec
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
			if got := ValidateFirewallTableSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValidateFirewallTableSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateFirewallSet(t *testing.T) {
	type args struct {
		in *firewall.Set
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
			if got := ValidateFirewallSet(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValidateFirewallSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateFirewallChain(t *testing.T) {
	type args struct {
		in *firewall.Chain
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
			if got := ValidateFirewallChain(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValidateFirewallChain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateFirewallObject(t *testing.T) {
	type args struct {
		in *firewall.Object
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
			if got := ValidateFirewallObject(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValidateFirewallObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateFirewallRule(t *testing.T) {
	type args struct {
		in *firewall.Rule
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
			if got := ValidateFirewallRule(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValidateFirewallRule() = %v, want %v", got, tt.want)
			}
		})
	}
}
