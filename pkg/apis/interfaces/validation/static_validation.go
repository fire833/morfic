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

package validation

import (
	"errors"

	interfaces "github.com/fire833/vroute/pkg/apis/interfaces"
)

var (
	// When a mac address is invalid.
	InvalidMACAddressError = errors.New("invalid mac address")
)

// Statically Validates a LinkList object coming into the API
// and returns a list of errors with the object.
func ValidateLinkList(in *interfaces.LinkList) []error {
	var retErrors []error
	return retErrors
}

// Statically Validates an Link object coming into the API
// and returns a list of errors with the object.
func ValidateLink(in *interfaces.Link) []error {
	var retErrors []error
	return retErrors
}

// Statically Validates an LinkSpec object coming into the API
// and returns a list of errors with the object.
func ValidateLinkSpec(in *interfaces.LinkSpec) []error {
	var retErrors []error
	return retErrors
}
