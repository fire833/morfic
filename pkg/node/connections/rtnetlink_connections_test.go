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

package connections

import (
	"fmt"
	"math/big"
	"math/rand"
	"sync"
	"testing"
	"time"

	crand "crypto/rand"
)

func init() {
	seed, _ := crand.Int(crand.Reader, big.NewInt(99999999999))
	rand.Seed(seed.Int64()) // Set a new seed for math random.
}

func TestConnectionLoad(t *testing.T) {
	NetlinkPool = NewConnPool(rand.Intn(10))

	workerWait := new(sync.WaitGroup)
	workerNums := rand.Intn(25) + 1
	workerWait.Add(workerNums)

	// Start a random number of workers that will get connections from the pool.
	for i := 0; i < workerNums; i++ {
		go func(wg *sync.WaitGroup) {
			defer wg.Done()

			tstart := time.Now()
			c, e := NetlinkPool.GetConn()
			fmt.Printf("Acuired new conn fd, time: %v\n", time.Since(tstart))
			if e != nil {
				t.Logf("Error in acquiring worker thread netlink socket.")
			}

			time.Sleep(time.Second * time.Duration(rand.Intn(4)+1))
			NetlinkPool.ReturnConn(c)

		}(workerWait)
	}

	workerWait.Wait()

}
