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

package comparator

import (
	api "github.com/fire833/morfic/pkg/apis/ipcapi/v1alpha1"
)

// Returns true if the link objects are equivalent. Otherwise returns false.
func CompareLink(link1 *api.Link, link2 *api.Link) bool {

	if link1.Name != link2.Name {
		return false
	}

	if link1.Index != link2.Index {
		return false
	}

	if link1.Mtu != link2.Mtu {
		return false
	}

	if link1.Type != link2.Type {
		return false
	}

	if !CompareLinkAttributes(link1.Attributes, link2.Attributes) {
		return false
	}

	return true
}

// Returns true if the link attributes are equivalent. Otherwise returns false.
func CompareLinkAttributes(attr1 *api.LinkAttributes, attr2 *api.LinkAttributes) bool {

	if attr1.Up != attr2.Up {
		return false
	}

	if attr1.NoArp != attr2.NoArp {
		return false
	}

	if attr1.Multicast != attr2.Multicast {
		return false
	}

	if attr1.Dynamic != attr2.Dynamic {
		return false
	}

	return true
}
