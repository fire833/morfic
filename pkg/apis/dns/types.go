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

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type RecordType int

const (
	ARecord RecordType = iota
	AAAARecord
	CNAMERecord
	TXTRecord
	SRVRecord
	MXRecord
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen=true

// DNSRecordList represents a list of DNS records.
type DNSRecordList struct {
	metav1.TypeMeta `json:",inline" yaml:",inline"`

	// Standard object metadata.
	// Utilizes the Kubernetes metadata object spec for now.
	metav1.ObjectMeta `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	// Items specifies the array of DNS records.
	Items []DNSRecord `json:"items" yaml:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen=true

// DNSRecord represents a desired DNS record for some DNS server.
type DNSRecord struct {
	metav1.TypeMeta `json:",inline" yaml:",inline"`

	// Standard object metadata.
	// Utilizes the Kubernetes metadata object spec for now.
	metav1.ObjectMeta `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	Spec DNSRecordSpec `json:"spec" yaml:"spec"`

	Status DNSRecordStatus `json:"status" yaml:"status"`
}

// DNSRecordSpec specifies the desired specification for a DNS record.
type DNSRecordSpec struct {
	Type  string `json:"type" yaml:"type"`
	Host  string `json:"host" yaml:"host"`
	Value string `json:"value" yaml:"value"`
	TTL   uint   `json:"ttl" yaml:"ttl"`
}

// DNSRecordStatus describes the current status of a DNS record.
type DNSRecordStatus struct {
	Deployed bool `json:"deployed" yaml:"deployed"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen=true

// DNSProviderList describes a list of DNS providers.
type DNSProviderList struct {
	metav1.TypeMeta `json:",inline" yaml:",inline"`

	// Standard object metadata.
	// Utilizes the Kubernetes metadata object spec for now.
	metav1.ObjectMeta `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	// Items specifies the array of DNS providers.
	Items []DNSProvider `json:"items" yaml:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen=true

// DNSProvider represents a provider for implementing DNS records.
type DNSProvider struct {
	metav1.TypeMeta `json:",inline" yaml:",inline"`

	// Standard object metadata.
	// Utilizes the Kubernetes metadata object spec for now.
	metav1.ObjectMeta `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	// Spec is the desired spec of this DNS provider.
	Spec DNSProviderSpec `json:"spec" yaml:"spec"`

	// Status is the current state of this DNS provider on the host.
	//
	// Should not be filled out by the user, will be filled/managed
	// by the server. Can be read by user at runtime.
	Status DNSProviderStatus `json:"status" yaml:"status"`
}

// DNSProviderSpec specifies the desired spec for a DNS provider.
type DNSProviderSpec struct {
}

// DNSProviderStatus represents the current status for a DNS provider.
type DNSProviderStatus struct {
}
