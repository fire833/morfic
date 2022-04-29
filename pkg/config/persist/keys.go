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

package persist

import (
	"fmt"
	"os"
)

var (
	// Default
	keylocations []string = []string{"/etc/vroute/keys.json", "/etc/vroute/keys.yaml", "/etc/vroute/keys.toml"}
)

type KeyStore struct {
	APIKeys []Key  `json:"keys" yaml:"keys" toml:"keys"`
	Users   []User `json:"users" yaml:"users" toml:"users"`
}

type Key struct {
	// API key, in argon2id format.
	Key string `json:"api_key" yaml:"apiKey" toml:"apikey"`
	// Unixnano expiration for this credential.
	Expiration int64 `json:"expire" yaml:"expire" toml:"expire"`
	// Effective priviledge level of this credential.
	PrivLevel int `json:"privilege" yaml:"privilege" toml:"privilege"`
}

type User struct {
	// Username for this user
	Username string `json:"user" yaml:"user" toml:"user"`
	// Password for this user, (in argon2 hashed format)
	PassKey string `json:"pass_key" yaml:"passKey" toml:"passkey"`
	// Unixnano expiration for this credential.
	Expiration int64 `json:"expire" yaml:"expire" toml:"expire"`
	// Effective priviledge level of this credential.
	PrivLevel int `json:"privilege" yaml:"privilege" toml:"privilege"`
}

func UnWrapConfig() {
	for _, location := range keylocations {
		if _, e := os.Stat(location); e != nil {
			continue
		}

		store := &KeyStore{}

		if err := Unwrap(location, store); err != nil {

		}
		return
	}

	fmt.Print("KeyStore file not found.")
	os.Exit(1)
}
