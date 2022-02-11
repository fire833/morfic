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

package dns

/*
	DNSProvider is an interface that can be utilized by the control plane to
	configure/manage DNS records on an authoritative DNS server.
*/
type DNSProvider interface {

	/*
		Initialize is called by the control plane when a DNS provider is being
		started up, either on control plane bootup or by user request. Similar
		to the PersistenceProvider and SecretPersistenceProvider interfaces, this
		method should be used by the provider to initialize any connections that
		are needed for future operations, opening file descriptors for files that
		could be needed in the future, etc. If there is an error with initalizing
		this provider, the appropriate error should be returned and the control
		plane will assume the provider is dead.
	*/
	Initialize() error

	GetRecords() []Record

	PutRecord(Record) error

	DeleteRecord(Record) error

	UpdateRecord(Record, Record) error
}

type Record struct {
	Type  string `json:"type" yaml:"type"`
	Host  string `json:"host" yaml:"host"`
	Value string `json:"value" yaml:"value"`
	TTL   uint   `json:"ttl" yaml:"ttl"`
}
