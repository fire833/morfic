package cri

import (
	"context"

	criapi "k8s.io/cri-api/pkg/apis/runtime/v1"
)

func stuff() {
	resp, err := RuntimeClient.ContainerStats(context.Background(), &criapi.ContainerStatsRequest{})

	RuntimeClient.CreateContainer(context.Background(), &criapi.CreateContainerRequest{
		Config: &criapi.ContainerConfig{
			Linux: &criapi.LinuxContainerConfig{},
		},
	})
}

type Service interface {
	Start()
}
