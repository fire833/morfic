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

package apis

import (
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ()

func ValidateTypeMeta(in *metav1.TypeMeta, desiredKind, desiredAPIVersion string) []error {
	var retErrors []error
	if in.Kind != desiredKind {
		retErrors = append(retErrors, fmt.Errorf("desired kind %s does not match actual kind of object: %s", desiredKind, in.Kind))
	}

	if in.APIVersion != desiredAPIVersion {
		retErrors = append(retErrors, fmt.Errorf("desired APIVersion %s does not match actual APIVersion of object: %s", desiredAPIVersion, in.APIVersion))
	}

	return retErrors
}

func ValidateObjectMeta(in *metav1.ObjectMeta) []error {
	var retErrors []error
	return retErrors
}
