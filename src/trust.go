/*
*	Copyright (C) 2021  Kendall Tauser
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

package src

import (
	"crypto/rand"
	"crypto/x509"
	"sync"
	"time"
)

var SharedToken *SharedTrustToken

// Certificate that validates requests between
// subprocesses during runtime. Is valid for 5 years.
var RuntimeCA x509.Certificate

// The local certificate loaded
var LocalCert *x509.Certificate

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
func GenerateRuntimeTrustToken() {
	token := NewToken()
	// Set the global token
	SharedToken = token

	// TODO generate runtime CA.

	return
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
	return time.Now().Sub(token.tokenResetTime)
}

func (token *SharedTrustToken) GetToken() string {
	return token.rpcToken
}
