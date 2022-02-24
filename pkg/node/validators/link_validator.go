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

package validators

import (
	api "github.com/fire833/vroute/pkg/apis/ipcapi/v1alpha1"
)

// wrapper for validating the integrity of a link object.
// returns the status code for the link at this stage.
func ValidateLink(in *api.Link) error {

	if e := ValidateInterfaceName(in.GetName()); e != nil {
		return e
	}

	for _, addr := range in.Address {
		if e := ValidateLinkAddress(addr); e != nil {
			return e
		} else {
			continue
		}
	}

	return nil
}

func ValidateInterfaceName(in string) error {
	if len([]byte(in)) >= 16 {
		return NewError("invalid link name length (must be less than or equal to 16 bytes)", api.ReturnStatusCodes_INVALID_FIELD_ERROR)
	}
	return nil
}