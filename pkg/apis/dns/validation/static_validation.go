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
	"errors"
	"fmt"
	"net"
	"regexp"

	"github.com/fire833/morfic/pkg/apis"
	"github.com/fire833/morfic/pkg/apis/dns"
)

var ()

func ValidateDNSRecordList(in *dns.DNSRecordList) []error {
	var retErrors []error
	retErrors = append(retErrors, apis.ValidateTypeMeta(&in.TypeMeta, "DNSRecordList", "dns.morfic.io/v1alpha1")...)
	for _, record := range in.Items {
		retErrors = append(retErrors, ValidateDNSRecord(&record)...)
	}
	return retErrors
}

func ValidateDNSRecord(in *dns.DNSRecord) []error {
	var retErrors []error
	retErrors = append(retErrors, apis.ValidateTypeMeta(&in.TypeMeta, "DNSRecord", "dns.morfic.io/v1alpha1")...)
	retErrors = append(retErrors, apis.ValidateObjectMeta(&in.ObjectMeta)...)
	retErrors = append(retErrors, ValidateDNSRecordSpec(&in.Spec)...)
	return retErrors
}

func ValidateDNSRecordSpec(in *dns.DNSRecordSpec) []error {
	var retErrors []error

	retErrors = append(retErrors, ValidateDomain(in.Host, in.Domain)...)

	if in.TTL > 2147483647 { // Max value in RFC 1034
		retErrors = append(retErrors, fmt.Errorf("invalid TTL length: %v, must be less than 2147483647 seconds", in.TTL))
	}

	switch in.Type {
	case dns.ARecord: // Validate that the value is an IPv4 value.
		{
			if ip := net.ParseIP(in.Value); len(ip) != 4 {
				retErrors = append(retErrors, fmt.Errorf("invalid IPv4 address for A record type: %s", in.Value))
			}
		}
	case dns.AAAARecord:
		{
			if ip := net.ParseIP(in.Value); len(ip) != 16 {
				retErrors = append(retErrors, fmt.Errorf("invalid IPv6 address for AAAA record type: %s", in.Value))
			}
		}
	}
	// TODO addsupport for validating more entry types.

	return retErrors
}

func ValidateDNSProviderList(in *dns.DNSProviderList) []error {
	var retErrors []error
	for _, provider := range in.Items {
		retErrors = append(retErrors, ValidateDNSProvider(&provider)...)
	}
	return retErrors
}

func ValidateDNSProvider(in *dns.DNSProvider) []error {
	var retErrors []error
	retErrors = append(retErrors, apis.ValidateTypeMeta(&in.TypeMeta, "DNSProvider", "dns.morfic.io/v1alpha1")...)
	retErrors = append(retErrors, apis.ValidateObjectMeta(&in.ObjectMeta)...)
	retErrors = append(retErrors, ValidateDNSProviderSpec(&in.Spec)...)
	return retErrors
}

func ValidateDNSProviderSpec(in *dns.DNSProviderSpec) []error {
	var retErrors []error
	retErrors = append(retErrors, ValidateDNSProviderCloudflareSpec(&in.Cloudflare)...)
	retErrors = append(retErrors, ValidateDNSProviderRoute53Spec(&in.Route53)...)
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

var (
	domainHostRegex  = regexp.MustCompile(`^([a-zA-Z0-9-\.]{1,63})\.([a-zA-Z0-9-]{1,63})\.([a-z-A-Z]{2,6})$`)
	errInvalidDomain = errors.New("invalid domain/hostname, does not conform to semantics")
)

func ValidateDomain(host, domain string) []error {
	var retErrors []error

	str := host + "." + domain

	if !domainHostRegex.Match([]byte(str)) {
		retErrors = append(retErrors, errInvalidDomain)
	}
	return retErrors
}
