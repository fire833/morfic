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

package morfic

// Generate listers
//go:generate lister-gen --input-dirs github.com/fire833/morfic/pkg/apis/net,github.com/fire833/morfic/pkg/apis/vpn,github.com/fire833/morfic/pkg/apis/firewall,github.com/fire833/morfic/pkg/apis/dns,github.com/fire833/morfic/pkg/apis/services --go-header-file hack/boilerplate.go.txt --output-package pkg/client/listers

// Genarate clientset
//go:generate client-gen --input-dirs github.com/fire833/morfic/pkg/apis/net,github.com/fire833/morfic/pkg/apis/vpn,github.com/fire833/morfic/pkg/apis/firewall,github.com/fire833/morfic/pkg/apis/dns,github.com/fire833/morfic/pkg/apis/services --input "net/v1alpha1,vpn/v1alpha1,dns/v1alpha1,firewall/v1alpha1,services/v1alpha1" --fake-clientset --clientset-name versioned --input-base github.com/fire833/morfic/pkg/apis --go-header-file hack/boilerplate.go.txt --output-package pkg/client/clientset

// Generate informers
//go:generate informer-gen --input-dirs github.com/fire833/morfic/pkg/apis/net,github.com/fire833/morfic/pkg/apis/vpn,github.com/fire833/morfic/pkg/apis/firewall,github.com/fire833/morfic/pkg/apis/dns,github.com/fire833/morfic/pkg/apis/services --go-header-file hack/boilerplate.go.txt --output-package pkg/client/informers --listers-package github.com/fire833/morfic/pkg/client/listers --versioned-clientset-package github.com/fire833/morfic/pkg/client/clientset/versioned --output-package pkg/client/informers
