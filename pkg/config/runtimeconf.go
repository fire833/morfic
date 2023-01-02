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

package config

import (
	"fmt"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

var (
	// The global unmarshalled runtime configuration for use within the control plane for it's internal operations
	// and defaults. This tries to offload as many constants/static data to a file so it can be edited by the user,
	// rather than keeping the data wrapped up in the binary.
	CPRF *ControlPlaneRuntimeConfig

	// Default locations where runtime configs are loaded in from on the host for the controlplane.
	// By default, the package installs a default config at /etc/router-cp/runtime.conf.
	runtimePaths []string = []string{"/etc/morfic/runtime.conf"}
)

type ControlPlaneRuntimeConfig struct {
	// Define a custom CRI gRPC socket, other than the ones that are defined by default already.
	// This socket location will take precendence over the default cri-o/containerd/dockershim socket paths
	// baked into the binary.
	CRISocket string `json:"cri_socket" yaml:"criSocket" toml:"criSocket"`
	// Another location for the keystore to be loaded in by the control plane. Will take precedence over any of the other
	// config locations that are baked into the binary.
	KeyStorePath string `json:"key_store" yaml:"keyStorePath" toml:"keyStorePath"`
	// Location for the priviledged node controller process to create a unix socket for IPC communication.
	// Defaults at runtime to /var/run/morfic/morfic_node.sock
	NodeControllerSocket string `json:"node_controller_socket" yaml:"nodeControllerSocket" toml:"nodeControllerSocket"`
}

func LoadRuntimeConfig() {

	for _, path := range runtimePaths {
		if file, err := os.ReadFile(path); err != nil {
			continue
		} else {
			c := &ControlPlaneRuntimeConfig{}

			setDefaultValues(c) // Sets defaults in the data structure.

			// Default right now to having the main configuration file be TOML, to follow suit with CRI-O and other container
			// applications/frameworks.
			if _, err := toml.Decode(string(file), c); err != nil {
				log.Default().Printf("Error with loading potential default configuration file %s: %s", path, err)
				continue
			}

			CPRF = c
			return
		}
	}

	fmt.Printf("No configuration file found in any of the default locations, exiting.")
	os.Exit(1)

}

func setDefaultValues(c *ControlPlaneRuntimeConfig) {
	c.NodeControllerSocket = "/var/run/morfic/morfic_node.sock"
}
