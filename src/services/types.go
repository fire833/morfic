/*
*	Copyright (C) 2021  Kendall Tauser
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

package services

import (
	"sync"

	criapi "k8s.io/cri-api/pkg/apis/runtime/v1"
)

// Describes a service and its associated state to be containerized/run on the host.
type ServiceDescriptor struct {
	// Mutex if API/control-plane and status controller needs to configure the state of this service.
	M sync.Mutex
	// Describes the state of this service, can either be
	// 1 (running), 0 (disabled/not-running), or -1 (failed/exited).
	State int
	// Internal container state
	CState criapi.ContainerState
	// Map of command names inside of container mapped to the container ID of each of the running containers.
	Id map[string]string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the sandbox to which this container belongs.
	PodSandboxId string `protobuf:"bytes,2,opt,name=pod_sandbox_id,json=podSandboxId,proto3" json:"pod_sandbox_id,omitempty"`
	// Creation time of the container in nanoseconds.
	CreatedAt int64 `protobuf:"varint,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// "Static" configuration for a service to be read by the control plane for initializing a service,
	// and parts can be updated by the API during runtime.
	Config *ServiceConfiguration
	// More state to be added here as the API expands and as I learn more about CRI operations.
}

// The attached configuration for describing a service/pod to be run on the host.
// Is designed to be similar to a Deployment yaml file for Kubernetes, but with a more paired-down API.
type ServiceConfiguration struct {
	// If the "static" configuration is to be changed by the API.
	M sync.Mutex
	// Name of the service for identification and other things.
	Name          string   `json:"name" yaml:"name"`
	Configuration *Svcconf `json:"configuration" yaml:"configuration"`

	SandboxConfig    *criapi.PodSandboxConfig  `json:"pod_sandbox"`
	ContainerConfigs []*criapi.ContainerConfig `json:"containers"`
}

type Svcconf struct {
	Meta map[string]interface{} `json:"metadata" yaml:"metadata"`
	Spec *ServiceSpec           `json:"spec" yaml:"spec"`
}

type ServiceSpec struct {
	Hostname    string   `json:"hostname" yaml:"hostname"`
	Nameservers []string `json:"nameservers" yaml:"nameservers"`
}
