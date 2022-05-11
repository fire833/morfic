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

import "github.com/fire833/morfic/pkg/apis/firewall"

func ValidateFirewallTableList(in *firewall.TableList) []error {
	var retErrors []error
	return retErrors
}

func ValidateFirewallTable(in *firewall.Table) []error {
	var retErrors []error
	return retErrors
}

func ValidateFirewallTableSpec(in *firewall.FirewallTableSpec) []error {
	var retErrors []error
	return retErrors
}

func ValidateFirewallSet(in *firewall.Set) []error {
	var retErrors []error
	return retErrors
}

func ValidateFirewallChain(in *firewall.Chain) []error {
	var retErrors []error
	return retErrors
}

func ValidateFirewallObject(in *firewall.Object) []error {
	var retErrors []error
	return retErrors
}

func ValidateFirewallRule(in *firewall.Rule) []error {
	var retErrors []error
	return retErrors
}
