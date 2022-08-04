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
	r "crypto/rand"
	"math/big"
	"math/rand"
	"testing"

	"github.com/fire833/morfic/pkg/ipcapi/v1alpha1"
)

func init() {

}

func TestGenerateError(t *testing.T) {
	n, _ := r.Int(r.Reader, big.NewInt(999999999999999999))

	// Set up seed
	rand.Seed(n.Int64())

	for i := 0; i < 50; i++ {

		buf := make([]byte, 64)
		rand.Read(buf)

		e := NewError(string(buf), v1alpha1.ReturnStatusCodes_OK)

		if e.Error() != "OK: "+string(buf) {
			t.Fail()

			t.Logf("expected OK: %s, got: %s", string(buf), e.Error())
		}
	}
}
