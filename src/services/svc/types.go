package svc

import (
	"sync"
)

// Describes a service and its associated state to be containerized/run on the host.
type ServiceDescriptor struct {
	// Mutex if API needs to configure the state of this service.
	m sync.Mutex
	// Describes the state of this service, can either be
	// 1 (running), 0 (disabled/not-running), or -1 (failed/exited).
	State int

	// "Static" configuration for a service to be controlled by the control plane.
	Config *ServiceConfiguration
	// More state to be added here as the API expands and as I learn more about CRI operations.
}

// The attached configuration
type ServiceConfiguration struct {
	// If the "static" configuration is to be changed by the API.
	m sync.Mutex

	// Specific image name and attached SHA hash of the image that you want run on your host.
	ImageID string `json:"image_id"`
	// Ports to be exposed by this service.
	ExposedPorts []int `json:"ports_exposed"`
}
