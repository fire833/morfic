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

package firewall

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen=true

// FirewallRuleList represents a list of local firewall rules.
type FirewallRuleList struct {
	metav1.TypeMeta `json:",inline" yaml:",inline"`

	// Standard object metadata.
	// Utilizes the Kubernetes metadata object spec for now.
	metav1.ObjectMeta `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	// Items represents the array of firewall rules.
	Items []FirewallRule `json:"items" yaml:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen=true

// FirewallRule represents a firewall rule.
type FirewallRule struct {
	metav1.TypeMeta `json:",inline" yaml:",inline"`

	// Standard object metadata.
	// Utilizes the Kubernetes metadata object spec for now.
	metav1.ObjectMeta `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	// Spec is the desired spec of this firewall rule object.
	Spec FirewallRuleSpec `json:"spec" yaml:"spec"`

	// Status is the current state of this firewall rule on the host.
	//
	// Should not be filled out by the user, will be filled/managed
	// by the server. Can be read by user at runtime.
	Status FirewallRuleStatus `json:"status" yaml:"status"`
}

// FirewallRuleSpec represents the desired spec for a firewall rule.
type FirewallRuleSpec struct {
}

// FirewallRuleStatus specifies the status of a firewall rule.
type FirewallRuleStatus struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen=true

// FirewallChainList represents a list of local firewall chains.
type FirewallChainList struct {
	metav1.TypeMeta `json:",inline" yaml:",inline"`

	// Standard object metadata.
	// Utilizes the Kubernetes metadata object spec for now.
	metav1.ObjectMeta `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	// Items represents the array of firewall chains.
	Items []FirewallChain `json:"items" yaml:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen=true

// FirewallChain represents a firewall chain of rules.
type FirewallChain struct {
	metav1.TypeMeta `json:",inline" yaml:",inline"`

	// Standard object metadata.
	// Utilizes the Kubernetes metadata object spec for now.
	metav1.ObjectMeta `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	// Spec is the desired spec of this firewall chain object.
	Spec FirewallChainSpec `json:"spec" yaml:"spec"`

	// Status is the current state of this firewall chain on the host.
	//
	// Should not be filled out by the user, will be filled/managed
	// by the server. Can be read by user at runtime.
	Status FirewallRuleStatus `json:"status" yaml:"status"`
}

// FirewallChainSpec represents the desired spec of a firewall chain.
type FirewallChainSpec struct {
}

// FirewallChainStatus specifies the status of a firewall chain on the host.
type FirewallChainStatus struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen=true

// FirewallTableList repesents a list of firewall tables.
type FirewallTableList struct {
	metav1.TypeMeta `json:",inline" yaml:",inline"`

	// Standard object metadata.
	// Utilizes the Kubernetes metadata object spec for now.
	metav1.ObjectMeta `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	// Items represents the array of firewall tables.
	Items []FirewallTable `json:"items" yaml:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen=true

// FirewallTable represents a firewall table.
type FirewallTable struct {
	metav1.TypeMeta `json:",inline" yaml:",inline"`

	// Standard object metadata.
	// Utilizes the Kubernetes metadata object spec for now.
	metav1.ObjectMeta `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	// Spec is the desired spec of this firewall table object.
	Spec FirewallTableSpec `json:"spec" yaml:"spec"`

	// Status is the current state of this firewall table on the host.
	//
	// Should not be filled out by the user, will be filled/managed
	// by the server. Can be read by user at runtime.
	Status FirewallTableStatus `json:"status" yaml:"status"`
}

// FirewallTableSpec represents the desired spec for a firewall table on the host.
type FirewallTableSpec struct {
}

// FirewallTableStatus specifies the current status for a firewall table on the host.
type FirewallTableStatus struct {
}
