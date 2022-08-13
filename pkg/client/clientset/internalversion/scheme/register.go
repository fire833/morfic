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

// Code generated by client-gen. DO NOT EDIT.

package scheme

import (
	authentication "github.com/fire833/morfic/pkg/apis/authentication/install"
	certificates "github.com/fire833/morfic/pkg/apis/certificates/install"
	dns "github.com/fire833/morfic/pkg/apis/dns/install"
	firewall "github.com/fire833/morfic/pkg/apis/firewall/install"
	net "github.com/fire833/morfic/pkg/apis/net/install"
	services "github.com/fire833/morfic/pkg/apis/services/install"
	vpn "github.com/fire833/morfic/pkg/apis/vpn/install"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
)

var Scheme = runtime.NewScheme()
var Codecs = serializer.NewCodecFactory(Scheme)
var ParameterCodec = runtime.NewParameterCodec(Scheme)

func init() {
	v1.AddToGroupVersion(Scheme, schema.GroupVersion{Version: "v1"})
	Install(Scheme)
}

// Install registers the API group and adds types to a scheme
func Install(scheme *runtime.Scheme) {
	authentication.Install(scheme)
	certificates.Install(scheme)
	dns.Install(scheme)
	firewall.Install(scheme)
	net.Install(scheme)
	services.Install(scheme)
	vpn.Install(scheme)
}
