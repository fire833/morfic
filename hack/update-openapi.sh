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

echo "Rebuilding addresses API..."
openapi-gen --input-dirs ../pkg/apis/addresses -h boilerplate.go.txt -O zz_generated.openapi &
echo "Rebuilding authentication API..."
openapi-gen --input-dirs ../pkg/apis/authentication -h boilerplate.go.txt -O zz_generated.openapi &
echo "Rebuilding certificates API..."
openapi-gen --input-dirs ../pkg/apis/certificates -h boilerplate.go.txt -O zz_generated.openapi &
echo "Rebuilding config API..."
openapi-gen --input-dirs ../pkg/apis/config -h boilerplate.go.txt -O zz_generated.openapi &
echo "Rebuilding dns API..."
openapi-gen --input-dirs ../pkg/apis/dns -h boilerplate.go.txt -O zz_generated.openapi &
echo "Rebuilding firewall API..."
openapi-gen --input-dirs ../pkg/apis/firewall -h boilerplate.go.txt -O zz_generated.openapi &
echo "Rebuilding interfaces API..."
openapi-gen --input-dirs ../pkg/apis/interfaces -h boilerplate.go.txt -O zz_generated.openapi &
echo "Rebuilding metrics API..."
openapi-gen --input-dirs ../pkg/apis/metrics -h boilerplate.go.txt -O zz_generated.openapi &
echo "Rebuilding nat API..."
openapi-gen --input-dirs ../pkg/apis/nat -h boilerplate.go.txt -O zz_generated.openapi &
echo "Rebuilding neighbors API..."
openapi-gen --input-dirs ../pkg/apis/neighbors -h boilerplate.go.txt -O zz_generated.openapi &
echo "Rebuilding routes API..."
openapi-gen --input-dirs ../pkg/apis/routes -h boilerplate.go.txt -O zz_generated.openapi &
echo "Rebuilding services API..."
openapi-gen --input-dirs ../pkg/apis/services -h boilerplate.go.txt -O zz_generated.openapi &
echo "Rebuilding vpn API..."
openapi-gen --input-dirs ../pkg/apis/vpn -h boilerplate.go.txt -O zz_generated.openapi &

wait
