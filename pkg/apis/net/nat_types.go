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

package net

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen=true
// +genclient
// +genclient:nonNamespaced
// NATConfiguration describes a configuration for NAT masqurading on a system.
type NATConfiguration struct {
	metav1.TypeMeta `json:",inline" yaml:",inline"`

	// Standard object metadata.
	// Utilizes the Kubernetes metadata object spec for now.
	metav1.ObjectMeta `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	// Spec is the desired spec of this node's NAT config.
	Spec NATConfigurationSpec `json:"spec" yaml:"spec"`

	Status NATConfigurationStatus `json:"status" yaml:"status"`
}

type NATConfigurationSpec struct {
}

type NATConfigurationStatus struct {
}
