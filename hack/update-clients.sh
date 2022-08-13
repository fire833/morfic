#!/bin/bash

#    Copyright (C) 2022  Kendall Tauser
#
#    This program is free software; you can redistribute it and/or modify
#    it under the terms of the GNU General Public License as published by
#    the Free Software Foundation; either version 2 of the License, or
#    (at your option) any later version.
#
#    This program is distributed in the hope that it will be useful,
#    but WITHOUT ANY WARRANTY; without even the implied warranty of
#    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
#    GNU General Public License for more details.
#
#    You should have received a copy of the GNU General Public License along
#    with this program; if not, write to the Free Software Foundation, Inc.,
#    51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA.

echo "Updating API clients..."

client-gen -i \
github.com/fire833/morfic/pkg/apis/net/v1alpha1,\
github.com/fire833/morfic/pkg/apis/authentication/v1alpha1,\
github.com/fire833/morfic/pkg/apis/certificates/v1alpha1,\
github.com/fire833/morfic/pkg/apis/dns/v1alpha1,\
github.com/fire833/morfic/pkg/apis/firewall/v1alpha1,\
github.com/fire833/morfic/pkg/apis/services/v1alpha1,\
github.com/fire833/morfic/pkg/apis/vpn/v1alpha1, \
github.com/fire833/morfic/pkg/apis/net,\
github.com/fire833/morfic/pkg/apis/authentication,\
github.com/fire833/morfic/pkg/apis/certificates,\
github.com/fire833/morfic/pkg/apis/dns,\
github.com/fire833/morfic/pkg/apis/firewall,\
github.com/fire833/morfic/pkg/apis/services,\
github.com/fire833/morfic/pkg/apis/vpn \
-h boilerplate.go.txt -o ../vendor -p github.com/fire833/morfic/pkg/client -n ../pkg/client/clientset --fake-clientset

# client-gen --input "${1}" -h boilerplate.go.txt -o ../pkg/client -p clientset -n morfic --fake-clientset
