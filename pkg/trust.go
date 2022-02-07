/*
*	Copyright (C) 2022 Kendall Tauser
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

package pkg

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/x509"
	"fmt"
	"math/big"
	"os"
	"sync"
	"time"
)

var SharedToken *SharedTrustToken

// Certificate that validates requests between
// subprocesses during runtime. Is valid for 5 years.
var RuntimeCA *x509.Certificate
var RuntimeCAKey ed25519.PrivateKey

// The local certificate loaded for use by the grpc client
// for authenticating with the grpc node server.
var LocalCert *x509.Certificate
var LocalCertKey ed25519.PrivateKey

const (
	tokenByteSize = 64
)

type SharedTrustToken struct {
	// Sync for this structure on updates.
	m *sync.Mutex
	// bootToken string
	// Token used for verifying requests between rpc endpoints between different runtime processes.
	rpcToken string
	// Describes the time that this structure was initiated.
	restTime time.Time
	// Describes the time that this token was last updated.
	tokenResetTime time.Time
}

// Used to generate shared tokens in memory during runtime for authenticating subprocesses to each other
// during RPC calls to the unix endpoints.
//
// This method should be called only once by the main vroute command function, as it is used for all of
// setup with runtime CA as well as well as the runtime trust token.
func GenerateRuntimeTrustAnchors() {
	token := NewToken()
	// Set the global token
	SharedToken = token

	pub, priv, e := ed25519.GenerateKey(rand.Reader)
	if e != nil {
		// Fail out if we can't create this key, because then bootstrapping won't work.
		fmt.Printf("Unable to create runtime CA keypair: %v", e)
		os.Exit(1)
	}

	RuntimeCAKey = priv

	RuntimeCA = &x509.Certificate{
		SerialNumber:       big.NewInt(1),
		SignatureAlgorithm: x509.PureEd25519,
		IsCA:               true, // This is the CA.
		NotBefore:          time.Now(),
		NotAfter:           time.Now().Local().AddDate(5, 0, 0), // Default to making self-signed cert last for five years, so I don't have to create renewal logic.
		PublicKey:          pub,
		PublicKeyAlgorithm: x509.Ed25519,
	}

	pub2, priv2, e1 := ed25519.GenerateKey(rand.Reader)
	if e1 != nil {
		// Fail out if we can't create this key, because then bootstrapping won't work.
		fmt.Printf("Unable to create runtime certificate keypair: %v", e1)
		os.Exit(1)
	}

	LocalCertKey = priv2

	lcert, e2 := x509.CreateCertificate(rand.Reader, &x509.Certificate{
		SerialNumber:       big.NewInt(1),
		SignatureAlgorithm: x509.PureEd25519,
		IsCA:               true, // This is the CA.
		NotBefore:          time.Now(),
		NotAfter:           time.Now().Local().AddDate(5, 0, 0), // Default to making self-signed cert last for five years, so I don't have to create renewal logic.
		PublicKey:          pub2,
		PublicKeyAlgorithm: x509.Ed25519,
	}, RuntimeCA, pub2, RuntimeCAKey)
	if e2 != nil {
		// Fail out if we can't create this key, because then bootstrapping won't work.
		fmt.Printf("Unable to create runtime certificate: %v", e2)
		os.Exit(1)
	}

	LocalCert, _ = x509.ParseCertificate(lcert) // Ignore the error since we are passing in a cert that was just successfully generated.

}

func NewToken() *SharedTrustToken {
	token := make([]byte, tokenByteSize)
	rand.Reader.Read(token)

	time := time.Now()

	tstruct := &SharedTrustToken{
		m:              new(sync.Mutex),
		rpcToken:       string(token),
		restTime:       time, // Should not be mutated ever again.
		tokenResetTime: time,
	}

	// tstruct.m.Unlock()
	return tstruct
}

// Updates the token within the data structure and updates the new timestamp for
// when the token was generated. Used by the updater thread to keep the running
// processes synched up with good data for keeping trust with RPC calls.
func (token *SharedTrustToken) UpdateToken() {
	token.m.Lock()
	defer token.m.Unlock()

	t := make([]byte, tokenByteSize)
	rand.Reader.Read(t)

	token.rpcToken = string(t)
	token.tokenResetTime = time.Now()

}

// Returns uptime of the process token structure.
func (token *SharedTrustToken) GetTokenUptime() time.Duration {
	return time.Since(token.tokenResetTime)
}

func (token *SharedTrustToken) GetToken() string {
	return token.rpcToken
}
