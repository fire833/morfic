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

package node

import (
	"fmt"
	"testing"

	"github.com/fire833/morfic/pkg"
	"github.com/fire833/morfic/pkg/config"
)

func init() {
	config.CPRF = &config.ControlPlaneRuntimeConfig{}

	pkg.GenerateRuntimeTrustAnchors()
}

func TestStartGRPCServer(t *testing.T) {
	if cert, e := CreateServerCert(); e != nil {
		t.Fail()

		fmt.Printf("Unable to create server cert: %v\n", e)
	} else {
		if e1 := BeginNodeServer(cert); e1 != nil {
			t.Fail()

			fmt.Printf("Unable to start grpc server: %v\n", e1)
		}
	}
}
