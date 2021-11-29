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

package config

import (
	"errors"
	"sync"

	"github.com/fire833/router-cp/src/wg"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

// Describes a Wireguard tunnel interface configured by the control plane.
type WireguardTun struct {
	// Mutex if the API needs to mutate the state of this link.
	m sync.Mutex

	DevName string         `json:"device_name"`
	Config  wgtypes.Config `json:"interface_config"`
}

type wgConfig struct {
	PrivateKey   string               `json:"private_key"`
	ListenPort   int                  `json:"listen_port"`
	FirewallMark int                  `json:"firewall_mark"`
	ReplacePeers bool                 `json:"replace_peers"`
	Peers        []wgtypes.PeerConfig `json:"peer_configs"`
}

type wgPeer struct {
	PublicKey string `json:"public_key"`
}

// Checks that the tunnel device exists in the kernel. Returns nil if the device exists,
// returns an error if it does or if there is an error acquiring the device status/existence.
func (tun *WireguardTun) CheckDeviceExists() error {
	if wg.WgClient == nil {
		return errors.New("Wireguard client doesn't exist.")
	}

	if _, e := wg.WgClient.Device(tun.DevName); e != nil {
		return e
	} else {
		return nil
	}
}

// Configures the device state based on the configuration, assuming it exists.
func (tun *WireguardTun) SetStatus() error {
	if err := tun.CheckDeviceExists(); err != nil {
		return err
	}

	if err := wg.WgClient.ConfigureDevice(tun.DevName, tun.Config); err != nil {
		return err
	}

	return nil
}
