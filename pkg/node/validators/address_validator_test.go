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
	"fmt"
	"math/big"
	"math/rand"
	"net"

	crand "crypto/rand"

	"testing"

	api "github.com/fire833/vroute/pkg/apis/ipcapi/v1alpha1"
)

func init() {
	i, _ := crand.Int(crand.Reader, big.NewInt(999999999999999999))
	rand.Seed(i.Int64()) // Init a random seed for every test.
}

var (
	ipbyteLengths   = []int{1, 2, 3, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32}
	macbytesLengths = []int{1, 2, 3, 4, 5, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31}
)

func genBytesBasedOnType(t api.IPType) []byte {
	switch t {
	case api.IPType_IPV4:
		{
			data := make([]byte, net.IPv4len)
			rand.Read(data)

			return data
		}
	case api.IPType_IPV6:
		{
			data := make([]byte, net.IPv6len)
			rand.Read(data)

			return data
		}
	}

	return nil
}

func genNumberofBytes(n int) []byte {
	data := make([]byte, n)
	rand.Read(data)

	return data
}

func genValidIPAddress() *api.IPAddress {
	b := rand.Intn(2) // Either a 0 or 1

	ip := &api.IPAddress{}

	if b == 0 { // Gen an IPv4 address
		ip.Type = api.IPType_IPV4
		ip.Address = genBytesBasedOnType(api.IPType_IPV4)

		return ip
	} else if b == 1 {
		ip.Type = api.IPType_IPV6
		ip.Address = genBytesBasedOnType(api.IPType_IPV6)

		return ip
	} else {
		fmt.Println("Invalid number")
	}

	return ip
}

func genInvalidIPAddress() *api.IPAddress {
	b := rand.Intn(2) // Between 0-1 to decide what invalid type to generate.

	ip := &api.IPAddress{}

	switch b {
	case 0:
		{ // Generate an invalid type to address mapping, (ie IPv4 address with IPv6 type)
			b1 := rand.Intn(2)
			if b1 == 0 {
				ip.Type = api.IPType_IPV4
				ip.Address = genBytesBasedOnType(api.IPType_IPV6)

				return ip
			} else if b1 == 1 {
				ip.Type = api.IPType_IPV6
				ip.Address = genBytesBasedOnType(api.IPType_IPV4)

				return ip
			}
		}
	case 1:
		{ // Generate an invalidly sized address byte array, so would be between 1-3 bytes, 5-15, and >17 bytes.
			b1 := rand.Intn(2)
			if b1 == 0 {
				ip.Type = api.IPType_IPV4
				ip.Address = genNumberofBytes(ipbyteLengths[rand.Intn(30)])

				return ip
			} else if b1 == 1 {
				ip.Type = api.IPType_IPV6
				ip.Address = genNumberofBytes(ipbyteLengths[rand.Intn(30)])

				return ip
			}
		}
	}

	return ip
}

func TestValidateIPAddress(t *testing.T) {

	for i := 0; i < 500; i++ {
		addr := genValidIPAddress()

		if e := ValidateIPAddress(addr); e != nil {
			t.Fail()
			t.Logf("failed parsing: %v, should be valid, but ruled as invalid.", addr.Address)
		}
	}

	for i := 0; i < 500; i++ {
		addr := genInvalidIPAddress()

		if e := ValidateIPAddress(addr); e == nil {
			t.Fail()
			t.Logf("failed parsing: %v, should be invalid, but ruled as valid.", addr.Address)
		}
	}
}

func genValidMacAddress() *api.MACAddress {
	data := make([]byte, MACAddressLength)
	rand.Read(data)

	mac := &api.MACAddress{
		Address: data,
	}

	return mac
}

func genInvalidMacAddress() *api.MACAddress {
	data := make([]byte, macbytesLengths[rand.Intn(30)])
	rand.Read(data)

	mac := &api.MACAddress{
		Address: data,
	}

	return mac
}

func TestValidateMACAddress(t *testing.T) {
	for i := 0; i < 500; i++ {
		addr := genValidMacAddress()

		if e := ValidateMACAddress(addr); e != nil {
			t.Fail()
			t.Logf("labaled valid mac as an invalid mac.")
		}
	}

	for i := 0; i < 500; i++ {
		addr := genInvalidMacAddress()

		if e := ValidateMACAddress(addr); e == nil {
			t.Fail()
			t.Logf("labaled invalid mac as an valid mac.")
		}
	}
}
