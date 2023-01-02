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

package validation

import (
	"errors"
	n "net"

	"github.com/fire833/morfic/pkg/apis/net"
)

var (
	// When an IPSpec has an invalid 'type' field.
	ErrInvalidIPType = errors.New("ip type is not set to either 'IPv4' or 'IPv6'")
	// When an IPSpec has an invalid address that can't be parsed.
	ErrInvalidIP = errors.New("ip address is not correctly formatted")
)

// Statically Validates an AddressList object coming into the API
// and returns a list of errors with the object.
func ValidateAddressList(in *net.AddressList) []error {
	var retErrors []error

	for _, address := range in.Items {
		retErrors = append(retErrors, ValidateAddress(&address)...)
	}

	return retErrors
}

// Statically Validates an Address object coming into the API
// and returns a list of errors with the object.
func ValidateAddress(in *net.Address) []error {
	var retErrors []error
	retErrors = append(retErrors, ValidateAddressSpec(&in.Spec)...)
	return retErrors
}

// Statically validates an AddressSpec object coming into the API
// and returns a list of errors with the object.
func ValidateAddressSpec(in *net.AddressSpec) []error {
	var retErrors []error
	retErrors = append(retErrors, ValidateIPSpec(&in.IP)...)
	return retErrors
}

func ValidateIPSpec(in *net.IPSpec) []error {
	var retErrors []error

	if in.Type != "IPv4" && in.Type != "IPv6" {
		retErrors = append(retErrors, ErrInvalidIPType)
	}

	if ip := n.ParseIP(in.Address); ip == nil {
		retErrors = append(retErrors, ErrInvalidIP)
	}

	return retErrors
}
