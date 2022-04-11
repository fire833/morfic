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

	"github.com/fire833/vroute/pkg/apis/dns"
)

func TestValidateDNSRecordList(t *testing.T) {
	type args struct {
		in *dns.DNSRecordList
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
			if got := ValidateDNSRecordList(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValidateDNSRecordList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateDNSRecord(t *testing.T) {
	type args struct {
		in *dns.DNSRecord
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
			if got := ValidateDNSRecord(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValidateDNSRecord() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateDNSRecordSpec(t *testing.T) {
	type args struct {
		in *dns.DNSRecordSpec
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
			if got := ValidateDNSRecordSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValidateDNSRecordSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateDNSProviderList(t *testing.T) {
	type args struct {
		in *dns.DNSProviderList
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
			if got := ValidateDNSProviderList(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValidateDNSProviderList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateDNSProvider(t *testing.T) {
	type args struct {
		in *dns.DNSProvider
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
			if got := ValidateDNSProvider(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValidateDNSProvider() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateDNSProviderSpec(t *testing.T) {
	type args struct {
		in *dns.DNSProviderSpec
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
			if got := ValidateDNSProviderSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValidateDNSProviderSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateDNSProviderCloudflareSpec(t *testing.T) {
	type args struct {
		in *dns.CloudflareProviderSpec
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
			if got := ValidateDNSProviderCloudflareSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValidateDNSProviderCloudflareSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateDNSProviderRoute53Spec(t *testing.T) {
	type args struct {
		in *dns.Route53ProviderSpec
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
			if got := ValidateDNSProviderRoute53Spec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValidateDNSProviderRoute53Spec() = %v, want %v", got, tt.want)
			}
		})
	}
}
