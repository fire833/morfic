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

echo "Updating API listers..."

lister-gen -i \
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
-h boilerplate.go.txt -o ../pkg/client -p listers

# mv ../pkg/client/listers/apis/* ../pkg/client/listers

# rm -rf ../pkg/client/listers/apis

# echo "Rebuilding addresses API..."
# lister-gen --input-dirs github.com/fire833/morfic/pkg/apis/addresses -h boilerplate.go.txt -p addresses -o ../pkg/client/listers &
# echo "Rebuilding authentication API..."
# lister-gen --input-dirs github.com/fire833/morfic/pkg/apis/authentication -h boilerplate.go.txt -p authentication -o ../pkg/client/listers &
# echo "Rebuilding certificates API..."
# lister-gen --input-dirs github.com/fire833/morfic/pkg/apis/certificates -h boilerplate.go.txt -p certificates -o ../pkg/client/listers &
# echo "Rebuilding config API..."
# lister-gen --input-dirs github.com/fire833/morfic/pkg/apis/config -h boilerplate.go.txt -p config -o ../pkg/client/listers &
# echo "Rebuilding dns API..."
# lister-gen --input-dirs github.com/fire833/morfic/pkg/apis/dns -h boilerplate.go.txt -p dns -o ../pkg/client/listers & 
# echo "Rebuilding firewall API..."
# lister-gen --input-dirs github.com/fire833/morfic/pkg/apis/firewall -h boilerplate.go.txt -p firewall -o ../pkg/client/listers & 
# echo "Rebuilding interfaces API..."
# lister-gen --input-dirs github.com/fire833/morfic/pkg/apis/interfaces -h boilerplate.go.txt -p interfaces -o ../pkg/client/listers &
# echo "Rebuilding metrics API..."
# lister-gen --input-dirs github.com/fire833/morfic/pkg/apis/metrics -h boilerplate.go.txt -p metrics -o ../pkg/client/listers &
# echo "Rebuilding nat API..."
# lister-gen --input-dirs github.com/fire833/morfic/pkg/apis/nat -h boilerplate.go.txt -p nat -o ../pkg/client/listers &
# echo "Rebuilding neighbors API..."
# lister-gen --input-dirs github.com/fire833/morfic/pkg/apis/neighbors -h boilerplate.go.txt -p neighbors -o ../pkg/client/listers &
# echo "Rebuilding routes API..."
# lister-gen --input-dirs github.com/fire833/morfic/pkg/apis/routes -h boilerplate.go.txt -p routes -o ../pkg/client/listers &
# echo "Rebuilding services API..."
# lister-gen --input-dirs github.com/fire833/morfic/pkg/apis/services -h boilerplate.go.txt -p services -o ../pkg/client/listers &
# echo "Rebuilding vpn API..."
# lister-gen --input-dirs github.com/fire833/morfic/pkg/apis/vpn -h boilerplate.go.txt -p vpn -o ../pkg/client/listers &

# wait
