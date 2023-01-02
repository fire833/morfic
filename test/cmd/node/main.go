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

package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/fire833/morfic/pkg/config"
	"github.com/fire833/morfic/pkg/node"
)

func main() {

	testConfig() // set up the configuration struct.

	cert, err := ephemeralCert()
	if err != nil {
		fmt.Printf("Unable to create node server cert: %v\n", err)

		os.Exit(1)
	}

	if e := node.BeginNodeServer(cert); e != nil {
		fmt.Printf("Unable to start node grpc server: %v\n", e)

		os.Exit(1)
	}

}

func testConfig() {
	config.CPRF = new(config.ControlPlaneRuntimeConfig)

	config.CPRF.CRISocket = "/var/run/morfic/morfic_node.sock"
}

func ephemeralCert() (*tls.Certificate, error) {

	pub, priv, e := ed25519.GenerateKey(rand.Reader)
	if e != nil {
		return nil, e
	}

	templ := &x509.Certificate{
		SerialNumber:       big.NewInt(1),
		SignatureAlgorithm: x509.PureEd25519,
		IsCA:               false, // This is the CA.
		NotBefore:          time.Now(),
		NotAfter:           time.Now().Local().AddDate(0, 0, 1),
		PublicKey:          pub,
		PublicKeyAlgorithm: x509.Ed25519,
	}

	cert, e1 := x509.CreateCertificate(rand.Reader, templ, templ, pub, priv)
	if e1 != nil {
		return nil, e1
	}

	data := pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: cert,
	})

	finalcert, e2 := tls.X509KeyPair(data, priv)
	if e2 != nil {
		return nil, e2
	}

	return &finalcert, nil
}
