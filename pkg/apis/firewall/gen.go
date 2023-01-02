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

package firewall

//go:generate deepcopy-gen --input-dirs=../firewall -h ../../../hack/boilerplate.go.txt -O zz_generated.deepcopy
//go:generate openapi-gen --input-dirs=../firewall -h ../../../hack/boilerplate.go.txt -O zz_generated.openapi -p ../firewall
//go:generate defaulter-gen --input-dirs ../firewall -h ../../../hack/boilerplate.go.txt -O zz_generated.defaults
