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

func TestValidateDomain(t *testing.T) {
	// var e []error

	validCombos := map[string]string{
		"drive":                             "google.com",
		"tree-sac0-0001":                    "backblaze.com",
		"ae2.3204.edge7.Amsterdam1":         "level3.net",
		"be2977.rcr21.mex02.atlas":          "cogentco.com",
		"maps":                              "google.com",
		"ldap":                              "ninjawanted.com",
		"BR2.Amsterdam1":                    "surf.net",
		"drive1":                            "stacksphere.com",
		"lala-123":                          "vast-and-endless-sea.com",
		"something":                         "isstackoverflowdownforeveryoneorjustme.com",
		"3":                                 "14159265358979323846264338327950288419716939937510582097494459.net",
		"testapp":                           "isitworking.xyz",
		"docs":                              "vroute.io",
		"gar25.dlstx.ip":                    "att.net",
		"play":                              "fubo.tv",
		"ip-101-130.amnet":                  "com.ni",
		"be2784.nr01.b042396-0.mex02.atlas": "cogentco.com",
		"70-138-128-1.lightspeed.hstntx":    "sbcglobal.net",
		"167-338-273-3.lightspeed.chiill":   "sbcglobal.net",
		"ae1.37.bar4.Minneapolis2":          "level3.net",
		"if-ae-47-2.tcore1.svq-singapore":   "as6453.net",
		"if-be-10-2.ecore2.svq-singapore":   "as6453.net",
		"if-ae-2-2.tcore2.dt8-dallas":       "as6453.net",
		"ae0-2.RT.IR9.AMS.NL":               "retn.net",
		"be2763.ccr31.dfw01.atlas":          "cogentco.com",
		"ip4":                               "gtt.net",
		"ae0.cr2-nyc4.ip4":                  "gtt.net",
	}

	for host, domain := range validCombos {
		errs := ValidateDomain(host, domain)
		if len(errs) != 0 {
			t.Fail()
			t.Logf("failed a valid host/domain combo: %s.%s", host, domain)
		}

		// t.Logf("%s.%s", host, domain)
	}

}

func TestValidateDNSRecordSpec(t *testing.T) {
	// var e []error

	type args struct {
		in *dns.DNSRecordSpec
	}
	tests := []struct {
		name string
		args args
		want []error
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateDNSRecordSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValidateDNSRecordSpec() = %v, want %v", got, tt.want)
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
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateDNSProviderSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValidateDNSProviderSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}
