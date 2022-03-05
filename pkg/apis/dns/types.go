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

// RecordType defines the type of a DNS record.
type RecordType int

const (
	ARecord RecordType = iota
	AAAARecord
	CNAMERecord
	TXTRecord
	SRVRecord
	MXRecord
)

// ProviderStatus defines the status of a provider from the perspective of the control plane.
type ProviderStatus int

// Available statuses for provider.
const (
	Available ProviderStatus = iota
	PluginError
	Unavailable
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
	// Type specifies the type of this DNS entry. The supported values
	// are notated in the above enum for types.
	Type RecordType `json:"type" yaml:"type"`

	// Specify the hostname that is to be controlled via this record.
	// Should be formatted as <subdomain_to_be_used>.<domain>.<tld>.
	Host string `json:"host" yaml:"host"`

	// The value for the record.
	Value string `json:"value" yaml:"value"`

	// Time to live for the record, cannot be negative.
	//
	// +optional
	TTL uint `json:"ttl" yaml:"ttl"`
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
	// Specify the name for this DNS provider for reference by other
	// resources and/or the control plane itself.
	Name string `json:"name" yaml:"name"`

	// Specify a cloudflare provider as the backend for this provider object.
	//
	// +optional
	Cloudflare CloudflareProviderSpec `json:"cloudflare" yaml:"cloudflare"`

	// Specify a route53 provider as the backend for this provider object.
	//
	// +optional
	Route53 Route53ProviderSpec `json:"route53" yaml:"route53"`
}

// CloudflareProviderSpec represents the specifications required for using
// cloudflare to update/maintain DNS records.
type CloudflareProviderSpec struct {
	// Specify the secret name to pull credentials from for this backend.
	//
	// This secrets should contain a key cloudflare.api.id and cloudflare.api.key,
	// corresponding to theapi key id and key secret id that are needed to authenticate
	// with the cloudflare API.
	//
	// cloudflare.api.zone is another key that can be used to store the zone id for cloudflare
	// instead of specifying the ZoneId field within this spec object. The ZoneId field will take
	// precedence over any keys that are found in secretName, however.
	SecretName string `json:"secretName" yaml:"secretName"`

	// Specify the zone Id that this provider will manage with the provided API. This field will take
	// priority over an id that was specified in the secret secretName references.
	//
	// +optional
	ZoneId string `json:"zone" yaml:"zone"`
}

// Route53ProviderSpec represents the specifications required for using
// Route 53 from AWS to update/maintain DNS records.
type Route53ProviderSpec struct {
}

// DNSProviderStatus represents the current status for a DNS provider.
type DNSProviderStatus struct {
	// Enabled defines if the provider is enabled and has been verified as working by the control plane.
	Enabled bool `json:"enabled" yaml:"enabled"`

	// Status Defines what the current status of the provider.
	Status ProviderStatus `json:"status" yaml:"status"`

	// StatusString represents a more verbose string status of the provider.
	//
	// +optional
	StatusString string `json:"statusString" yaml:"statusString"`
}