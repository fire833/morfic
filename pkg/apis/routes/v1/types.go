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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type RouteTable struct {
	metav1.TypeMeta

	metav1.ObjectMeta

	Spec RouteTableSpec

	Status RouteTableStatus
}

type RouteTableSpec struct {
}

type RouteTableStatus struct {
}

type Route struct {
	metav1.TypeMeta

	metav1.ObjectMeta

	Spec RouteSpec

	Status RouteStatus
}

type RouteSpec struct {
}

type RouteStatus struct {
}
