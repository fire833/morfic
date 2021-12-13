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
	"context"
	"time"

	"github.com/fire833/vroute/src/cri"
	criapi "k8s.io/cri-api/pkg/apis/runtime/v1"
)

type ServiceController interface {
	Start() error
	Kill() error
	Update() error
	Stats()
}

func (sd *ServiceDescriptor) Start() error {
	sd.M.Lock()
	defer sd.M.Unlock()

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*5))
	defer ctx.Done()
	defer cancel()

	if resp, err := cri.RuntimeClient.RunPodSandbox(ctx, &criapi.RunPodSandboxRequest{
		Config:         sd.Config.SandboxConfig,
		RuntimeHandler: "", // Add a handler here for future requests.
	}, nil); err != nil {
		return err
	} else {
		sd.PodSandboxId = resp.GetPodSandboxId()
	}

	for _, container := range sd.Config.ContainerConfigs {

		ctx1, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*5))
		defer ctx.Done()
		defer cancel()

		if resp, err := cri.RuntimeClient.CreateContainer(ctx1, &criapi.CreateContainerRequest{
			PodSandboxId:  sd.PodSandboxId,
			SandboxConfig: sd.Config.SandboxConfig,
			Config:        container,
		}, nil); err != nil {
			return err // Just return error for now, but eventually need to log individual erros and report them for handling upstream.
		} else {
			sd.Id[container.Command[0]] = resp.GetContainerId()

			ctx, cancel := context.WithTimeout(ctx1, time.Duration(time.Second*5))
			defer ctx.Done()
			defer cancel()

			if e := startContainer(resp.GetContainerId(), ctx); e != nil {
				return err // Just return error for now, but eventually need to log individual erros and report them for handling upstream.
			}
		}

	}

	sd.State = 1

	return nil

}

func startContainer(id string, ctx context.Context) error {
	_, err := cri.RuntimeClient.StartContainer(ctx, &criapi.StartContainerRequest{
		ContainerId: id,
	}, nil)

	return err
}

func (sd *ServiceDescriptor) Kill() error {
	return nil
}
