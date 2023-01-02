/*
*	Copyright (C) 2023 Kendall Tauser
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

// Rule represents a firewall rule.
type Rule struct {
	Name      string `json:"name" yaml:"name"`
	Index     int    `json:"index" yaml:"index"`
	Statement string `json:"statement" yaml:"statement"`
	Handle    string `json:"handle" yaml:"handle"`
}

// Chain represents a firewall chain of rules.
type Chain struct {
	Rules []Rule `json:"rules" yaml:"rules"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen=true
// +genclient
// +genclient:nonNamespaced
// TableList repesents a list of firewall tables.
type TableList struct {
	metav1.TypeMeta `json:",inline" yaml:",inline"`

	// +optional
	metav1.ListMeta

	// Items represents the array of firewall tables.
	Items []Table `json:"items" yaml:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen=true
// +genclient
// +genclient:nonNamespaced
// Table represents a firewall table.
type Table struct {
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

	// Family represents the family of data that this table will filter on.
	Family string `json:"family" yaml:"family"`

	// Sets represents an array of sets that should be a part of this table.
	Sets []Set `json:"sets" yaml:"sets"`

	// Chains represents an array of chain objects that are a part of this table.
	Chains []Chain `json:"chains" yaml:"chains"`

	// Objects represents a set of stateful objects that are members of this table.
	Objects []Object `json:"objects" yaml:"objects"`
}

// FirewallTableStatus specifies the current status for a firewall table on the host.
type FirewallTableStatus struct {
	Enabled bool `json:"enabled"`
}

// Set reresents a set of objects within the host firewall configuration.
type Set struct {
}

// Object represents a stateful object.
type Object struct {
}
