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
	"io/fs"
	"os"
	"path/filepath"

	"github.com/fire833/vroute/src/config/persist"
)

var (
	serviceLocations []string = []string{"/etc/vroute/services/"}

	// Initialized by loadSavedConfigs()
	serviceDir string

	// Unmarshalled service configuration files that are loaded from disk, they will be loaded into service descriptors that will bundle
	// up additional runtime state information.
	// Initialized by loadSavedConfigs()
	ServiceConfigs []*ServiceConfiguration
)

func WatchServiceDescriptorFiles() {

	loadSavedConfigs()

	for {

		files := readAllConfigs()

	}
}

func loadSavedConfigs() {

	var f []fs.DirEntry
	// Find the first directory of possible options that
	for _, service := range serviceLocations {
		files, e := os.ReadDir(service)
		if e != nil {
			continue
		}

		// Set the current runtime directory location.
		serviceDir = service
		f = files
		break
	}

	ServiceConfigs = make([]*ServiceConfiguration, len(f))

	for _, s := range f {
		i, e := s.Info()
		if e != nil {
			continue
		}

		sd := &ServiceConfiguration{}

		if err := persist.Unwrap(filepath.Join(serviceDir, i.Name()), sd); err != nil {
			continue
		}

		// Save all the configs to be loaded
		ServiceConfigs = append(ServiceConfigs, sd)

	}
	return
}

// Used by the listener to constrantly compare configs to update them.
func readAllConfigs() []string {

	configs, e := os.ReadDir(serviceDir)
	if e != nil {
		return nil
	}

	var readConfigs []string

	for _, conf := range configs {
		if conf.IsDir() {
			continue
		}

		file, e := os.ReadFile(filepath.Join(serviceDir, conf.Name()))
		if e != nil {
			continue
		}

		readConfigs = append(readConfigs, string(file))
	}

	return readConfigs
}
