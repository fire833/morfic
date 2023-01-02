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

package app

import (
	"github.com/fire833/morfic/cmd/morfic-apiserver/app/options"
	"github.com/fire833/morfic/pkg"
	"github.com/spf13/cobra"

	genericapiserver "k8s.io/apiserver/pkg/server"
	"k8s.io/klog/v2"
)

const (
	name string = "morflet"
)

func NewMorfletServerCommand() *cobra.Command {
	opts := options.NewServerRunOptions()
	cmd := &cobra.Command{
		Use:   name,
		Short: "Run the morfic privileged node agent",
		Long:  ``,

		RunE: func(cmd *cobra.Command, args []string) error {
			return Run(opts, genericapiserver.SetupSignalHandler())
		},
	}

	return cmd
}

func Run(opts *options.ServerRunOptions, stopCh <-chan struct{}) error {
	// List application version for debugging.
	klog.Infof(pkg.NewVersionString(name))

	// genericapiserver.NewConfig()

	return nil
}
