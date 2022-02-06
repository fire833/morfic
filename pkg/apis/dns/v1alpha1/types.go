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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DSNRecordList represents a list of DNS records.
type DNSRecordList struct {
	metav1.TypeMeta `json:",inline" yaml:",inline"`

	// Standard object metadata.
	// Utilizes the Kubernetes metadata object spec for now.
	metav1.ObjectMeta `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	// Items specifies the array of DNS records.
	Items []DNSRecord `json:"items" yaml:"items"`
}

type DNSRecord struct {
}

// DynamicDNSProviderList describes a list of dynamic DNS providers.
type DynamicDNSProviderList struct {
	metav1.TypeMeta `json:",inline" yaml:",inline"`

	// Standard object metadata.
	// Utilizes the Kubernetes metadata object spec for now.
	metav1.ObjectMeta `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	// Items specifies the array of dynamic DNS providers.
	Items []DynamicDNSProvider `json:"items" yaml:"items"`
}

type DynamicDNSProvider struct {
}
