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

set -o errexit
set -o nounset
set -o pipefail

SCRIPT_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
CODEGEN_PKG=${CODEGEN_PKG:-$(cd "${SCRIPT_ROOT}"; ls -d -1 ./vendor/k8s.io/code-generator 2>/dev/null || echo ../code-generator)}

# generate the code with:
# --output-base    because this script should also be able to run inside the vendor dir of
#                  k8s.io/kubernetes. The output-base is needed for the generators to output into the vendor dir
#                  instead of the $GOPATH directly. For normal projects this can be dropped.
bash "generate-groups.sh" all \
  github.com/fire833/vroute/pkg/generated github.com/fire833/vroute/pkg/apis \
  "addresses:v1alpha1 authentication:v1 certificates:v1alpha1 config:v1 dns:v1alpha1 firewall:v1alpha1 interfaces:v1alpha1 metrics:v1alpha1 nat:v1alpha1 neighbors:v1alpha1 routes:v1 services:v1alpha1 vpn:v1alpha1" \
  --output-base "$(dirname "${BASH_SOURCE[0]}")/../../.." \
  --go-header-file "${SCRIPT_ROOT}"/hack/boilerplate.go.txt

# bash "generate-internal-groups.sh" "deepcopy,defaulter,conversion,openapi" \
#   github.com/fire833/vroute/pkg/generated github.com/fire833/vroute/pkg/apis github.com/fire833/vroute/pkg/apis \
#   "addresses:v1alpha1 authentication:v1 certificates:v1alpha1 config:v1 dns:v1alpha1 firewall:v1alpha1 interfaces:v1alpha1 metrics:v1alpha1 nat:v1alpha1 neighbors:v1alpha1 routes:v1 services:v1alpha1 vpn:v1alpha1" \
#   --output-base "$(dirname "${BASH_SOURCE[0]}")/../../.." \
#   --go-header-file "${SCRIPT_ROOT}/hack/boilerplate.go.txt"

# To use your own boilerplate text append:
#   --go-header-file "${SCRIPT_ROOT}/hack/custom-boilerplate.go.txt"

# # export GOPATH=$(go env GOPATH)

# GOPATH=$(go env GOPATH) ./generate-internal-groups.sh \
# deepcopy,client,lister,informer \
# github.com/fire833/vroute/pkg/client github.com/fire833/vroute/pkg/apis github.com/fire833/vroute/pkg/apis \
# "addresses:v1alpha1 authentication:v1 certificates:v1alpha1 config:v1 dns:v1alpha1 firewall:v1alpha1 interfaces:v1alpha1 metrics:v1alpha1 nat:v1alpha1 neighbors:v1alpha1 routes:v1 services:v1alpha1 vpn:v1alpha1" \
# --go-header-file boilerplate.go.txt \
# # --output-base ../pkg/client

# # GOPATH=$(go env GOPATH) ./generate-groups.sh \
# # all \
# # github.com/fire833/vroute/pkg/client github.com/fire833/vroute/pkg/apis \
# # "addresses:v1alpha1 authentication:v1 certificates:v1alpha1 config:v1 dns:v1alpha1 firewall:v1alpha1 interfaces:v1alpha1 metrics:v1alpha1 nat:v1alpha1 neighbors:v1alpha1 routes:v1 services:v1alpha1 vpn:v1alpha1" \
# # --go-header-file boilerplate.go.txt \

# # wait
