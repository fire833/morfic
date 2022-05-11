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

package cloudflare

import "github.com/fire833/morfic/pkg/dns"

type CloudflareDNSProvider struct {
}

func (dns *CloudflareDNSProvider) Initialize() error {
	return nil
}

func (dns *CloudflareDNSProvider) GetRecords() []dns.Record {
	return nil
}

func (dns *CloudflareDNSProvider) PutRecord(in dns.Record) error {
	return nil
}

func (dns *CloudflareDNSProvider) DeleteRecord(in dns.Record) error {
	return nil
}

func (dns *CloudflareDNSProvider) UpdateRecord(in dns.Record, update dns.Record) error {
	return nil
}
