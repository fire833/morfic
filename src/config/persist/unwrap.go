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
	"encoding/json"
	"errors"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v3"
)

func Unwrap(file string, i interface{}) error {

	data, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	switch filepath.Ext(file) {
	case ".json":
		{
			return json.Unmarshal(data, i)
		}
	case ".yml", ".yaml":
		{
			return yaml.Unmarshal(data, i)
		}
	case ".toml":
		{
			return toml.Unmarshal(data, i)
		}
	default:
		{
			return errors.New("Unable to unmarshal this file (unknown extension).")
		}
	}

}
