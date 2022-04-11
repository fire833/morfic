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
	"github.com/fire833/vroute/pkg/apis/dns"
)

func ValidateDNSRecordList(in *dns.DNSRecordList) []error {
	var retErrors []error
	for _, record := range in.Items {
		retErrors = append(retErrors, ValidateDNSRecord(&record)...)
	}
	return retErrors
}

func ValidateDNSRecord(in *dns.DNSRecord) []error {
	var retErrors []error
	return retErrors
}

func ValidateDNSRecordSpec(in *dns.DNSRecordSpec) []error {
	var retErrors []error
	return retErrors
}

func ValidateDNSProviderList(in *dns.DNSProviderList) []error {
	var retErrors []error
	return retErrors
}

func ValidateDNSProvider(in *dns.DNSProvider) []error {
	var retErrors []error
	return retErrors
}

func ValidateDNSProviderSpec(in *dns.DNSProviderSpec) []error {
	var retErrors []error
	return retErrors
}

func ValidateDNSProviderCloudflareSpec(in *dns.CloudflareProviderSpec) []error {
	var retErrors []error
	return retErrors
}

func ValidateDNSProviderRoute53Spec(in *dns.Route53ProviderSpec) []error {
	var retErrors []error
	return retErrors
}
