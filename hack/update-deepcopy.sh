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

echo "Updating API deepcopy..."

echo "Rebuilding addresses API..."
deepcopy-gen --input-dirs ../pkg/apis/addresses -h boilerplate.go.txt -O zz_generated.deepcopy &
echo "Rebuilding authentication API..."
deepcopy-gen --input-dirs ../pkg/apis/authentication -h boilerplate.go.txt -O zz_generated.deepcopy &
echo "Rebuilding certificates API..."
deepcopy-gen --input-dirs ../pkg/apis/certificates -h boilerplate.go.txt -O zz_generated.deepcopy &
echo "Rebuilding config API..."
deepcopy-gen --input-dirs ../pkg/apis/config -h boilerplate.go.txt -O zz_generated.deepcopy &
echo "Rebuilding dns API..."
deepcopy-gen --input-dirs ../pkg/apis/dns -h boilerplate.go.txt -O zz_generated.deepcopy &
echo "Rebuilding firewall API..."
deepcopy-gen --input-dirs ../pkg/apis/firewall -h boilerplate.go.txt -O zz_generated.deepcopy &
echo "Rebuilding interfaces API..."
deepcopy-gen --input-dirs ../pkg/apis/interfaces -h boilerplate.go.txt -O zz_generated.deepcopy & 
echo "Rebuilding metrics API..."
deepcopy-gen --input-dirs ../pkg/apis/metrics -h boilerplate.go.txt -O zz_generated.deepcopy &
echo "Rebuilding nat API..."
deepcopy-gen --input-dirs ../pkg/apis/nat -h boilerplate.go.txt -O zz_generated.deepcopy &
echo "Rebuilding neighbors API..."
deepcopy-gen --input-dirs ../pkg/apis/neighbors -h boilerplate.go.txt -O zz_generated.deepcopy &
echo "Rebuilding routes API..."
deepcopy-gen --input-dirs ../pkg/apis/routes -h boilerplate.go.txt -O zz_generated.deepcopy &
echo "Rebuilding services API..."
deepcopy-gen --input-dirs ../pkg/apis/services -h boilerplate.go.txt -O zz_generated.deepcopy &
echo "Rebuilding vpn API..."
deepcopy-gen --input-dirs ../pkg/apis/vpn -h boilerplate.go.txt -O zz_generated.deepcopy &

wait
