#!/usr/bin/env bash

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

# To use your own boilerplate text append:
#   --go-header-file "${SCRIPT_ROOT}/hack/custom-boilerplate.go.txt"

# # export GOPATH=$(go env GOPATH)

declare apiGroups="addresses/v1alpha1,authentication/v1,certificates/v1alpha1,config/v1,dns/v1alpha1,firewall/v1alpha1,interfaces/v1,metrics/v1alpha1,nat/v1alpha1,neighbors/v1alpha1,routes/v1,services/v1alpha1,vpn/v1alpha1"

# ./update-deepcopy.sh

# ./update-clients.sh $apiGroups

# ./update-listers.sh $apiGroups

# ./update-informers.sh $apiGroups

GOPATH="${PWD}/../pkg" ./generate-internal-groups.sh \
all \
github.com/fire833/morfic/pkg/client github.com/fire833/morfic/pkg/apis github.com/fire833/morfic/pkg/apis \
"addresses:v1alpha1 authentication:v1 certificates:v1alpha1 config:v1 dns:v1alpha1 firewall:v1alpha1 interfaces:v1 metrics:v1alpha1 nat:v1alpha1 neighbors:v1alpha1 routes:v1 services:v1alpha1 vpn:v1alpha1" \
--go-header-file boilerplate.go.txt

# GOPATH="${PWD}/../pkg" ./generate-groups.sh \
# all \
# github.com/fire833/morfic/pkg/client github.com/fire833/morfic/pkg/apis \
# "addresses:v1alpha1 authentication:v1 certificates:v1alpha1 config:v1 dns:v1alpha1 firewall:v1alpha1 interfaces:v1 metrics:v1alpha1 nat:v1alpha1 neighbors:v1alpha1 routes:v1 services:v1alpha1 vpn:v1alpha1" \
# --go-header-file boilerplate.go.txt

# # wait
