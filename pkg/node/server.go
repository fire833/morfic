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

package node

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"math/big"
	"net"
	"os"
	"time"

	"github.com/fire833/vroute/pkg"
	"github.com/fire833/vroute/pkg/apis/ipcapi/v1alpha1"
	"github.com/fire833/vroute/pkg/config"
	"github.com/fire833/vroute/pkg/node/netlink"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	NodeControllerServer *grpc.Server

	LocalCert    *x509.Certificate
	LocalCertKey ed25519.PrivateKey
)

// Begins the node grpc server. Should never exit.
func BeginNodeServer(cert *tls.Certificate) error {

	log.Printf("creating listener for node grpc server")
	// create the unix bind listener.
	l, e := net.Listen("unix", config.CPRF.NodeControllerSocket)
	if e != nil {
		fmt.Printf("Unable to start node server: %s", e.Error())
		os.Exit(1)
	}

	log.Printf("creating credentials for node grpc server")
	creds := credentials.NewServerTLSFromCert(cert)

	s := grpc.NewServer(grpc.Creds(creds))

	log.Printf("registering services with node grpc server")
	// Register the services
	handlers := netlink.NetlinkNodeServer{}

	s.RegisterService(&v1alpha1.NodeFirewallControllerService_ServiceDesc, handlers)
	s.RegisterService(&v1alpha1.NodeControllerService_ServiceDesc, handlers)

	NodeControllerServer = s
	// This will basically be the last main function for the process.
	return s.Serve(l)
}

func CreateServerCert() (*tls.Certificate, error) {
	pub, priv, e := ed25519.GenerateKey(rand.Reader)
	if e != nil {
		return nil, e
	}

	templ := &x509.Certificate{
		SerialNumber:       big.NewInt(1),
		SignatureAlgorithm: x509.PureEd25519,
		IsCA:               false,
		NotBefore:          time.Now(),
		NotAfter:           time.Now().Local().AddDate(5, 0, 0), // Default to making self-signed cert last for five years, so I don't have to create renewal logic.
		PublicKey:          pub,
		PublicKeyAlgorithm: x509.Ed25519,
	}

	cert, e1 := x509.CreateCertificate(rand.Reader, templ, pkg.RuntimeCA, pub, pkg.RuntimeCAKey)
	if e1 != nil {
		return nil, e1
	}

	finalcert, e2 := tls.X509KeyPair(cert, priv)
	if e2 != nil {
		return nil, e2
	}

	return &finalcert, nil
}
