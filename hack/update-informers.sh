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

echo "Updating API informers..."

informer-gen -i \
github.com/fire833/morfic/pkg/apis/addresses/v1alpha1,\
github.com/fire833/morfic/pkg/apis/authentication/v1,\
github.com/fire833/morfic/pkg/apis/certificates/v1alpha1,\
github.com/fire833/morfic/pkg/apis/config/v1,\
github.com/fire833/morfic/pkg/apis/dns/v1alpha1,\
github.com/fire833/morfic/pkg/apis/firewall/v1alpha1,\
github.com/fire833/morfic/pkg/apis/interfaces/v1,\
github.com/fire833/morfic/pkg/apis/metrics/v1alpha1,\
github.com/fire833/morfic/pkg/apis/nat/v1alpha1,\
github.com/fire833/morfic/pkg/apis/neighbors/v1alpha1,\
github.com/fire833/morfic/pkg/apis/routes/v1,\
github.com/fire833/morfic/pkg/apis/services/v1alpha1,\
github.com/fire833/morfic/pkg/apis/vpn/v1alpha1, \
github.com/fire833/morfic/pkg/apis/addresses,\
github.com/fire833/morfic/pkg/apis/authentication,\
github.com/fire833/morfic/pkg/apis/certificates,\
github.com/fire833/morfic/pkg/apis/config,\
github.com/fire833/morfic/pkg/apis/dns,\
github.com/fire833/morfic/pkg/apis/firewall,\
github.com/fire833/morfic/pkg/apis/interfaces,\
github.com/fire833/morfic/pkg/apis/metrics,\
github.com/fire833/morfic/pkg/apis/nat,\
github.com/fire833/morfic/pkg/apis/neighbors,\
github.com/fire833/morfic/pkg/apis/routes,\
github.com/fire833/morfic/pkg/apis/services,\
github.com/fire833/morfic/pkg/apis/vpn \
-h boilerplate.go.txt --listers-package github.com/fire833/morfic/pkg/client/listers \
--single-directory -o ../pkg/client -p informers \
--versioned-clientset-package github.com/fire833/morfic/pkg/client/clientset

# Move all of the apis out to the outer directory.
# mv ../pkg/client/informers/apis/* ../pkg/client/informers

# Remove the apis directory.
# rm -rf ../pkg/client/informers/apis

# echo "Rebuilding addresses API..."
# informer-gen --input-dirs github.com/fire833/morfic/pkg/apis/addresses -h boilerplate.go.txt --listers-package github.com/fire833/morfic/pkg/client/listers --single-directory -p addresses -o ../pkg/client/informers &
# echo "Rebuilding authentication API..."
# informer-gen --input-dirs github.com/fire833/morfic/pkg/apis/authentication -h boilerplate.go.txt --listers-package github.com/fire833/morfic/pkg/client/listers --single-directory -p authetication -o ../pkg/client/informers &
# echo "Rebuilding certificates API..."
# informer-gen --input-dirs github.com/fire833/morfic/pkg/apis/certificates -h boilerplate.go.txt --listers-package github.com/fire833/morfic/pkg/client/listers --single-directory -p certificates -o ../pkg/client/informers &
# echo "Rebuilding config API..."
# informer-gen --input-dirs github.com/fire833/morfic/pkg/apis/config -h boilerplate.go.txt --listers-package github.com/fire833/morfic/pkg/client/listers --single-directory -p config -o ../pkg/client/informers &
# echo "Rebuilding dns API..."
# informer-gen --input-dirs github.com/fire833/morfic/pkg/apis/dns -h boilerplate.go.txt --listers-package github.com/fire833/morfic/pkg/client/listers --single-directory -p dns -o ../pkg/client/informers & 
# echo "Rebuilding firewall API..."
# informer-gen --input-dirs github.com/fire833/morfic/pkg/apis/firewall -h boilerplate.go.txt --listers-package github.com/fire833/morfic/pkg/client/listers --single-directory -p firewall -o ../pkg/client/informers & 
# echo "Rebuilding interfaces API..."
# informer-gen --input-dirs github.com/fire833/morfic/pkg/apis/interfaces -h boilerplate.go.txt --listers-package github.com/fire833/morfic/pkg/client/listers --single-directory -p interfaces -o ../pkg/client/informers &
# echo "Rebuilding metrics API..."
# informer-gen --input-dirs github.com/fire833/morfic/pkg/apis/metrics -h boilerplate.go.txt --listers-package github.com/fire833/morfic/pkg/client/listers --single-directory -p metrics -o ../pkg/client/informers &
# echo "Rebuilding nat API..."
# informer-gen --input-dirs github.com/fire833/morfic/pkg/apis/nat -h boilerplate.go.txt --listers-package github.com/fire833/morfic/pkg/client/listers --single-directory -p nat -o ../pkg/client/informers &
# echo "Rebuilding neighbors API..."
# informer-gen --input-dirs github.com/fire833/morfic/pkg/apis/neighbors -h boilerplate.go.txt --listers-package github.com/fire833/morfic/pkg/client/listers --single-directory -p neighbors -o ../pkg/client/informers &
# echo "Rebuilding routes API..."
# informer-gen --input-dirs github.com/fire833/morfic/pkg/apis/routes -h boilerplate.go.txt --listers-package github.com/fire833/morfic/pkg/client/listers --single-directory -p routes -o ../pkg/client/informers &
# echo "Rebuilding services API..."
# informer-gen --input-dirs github.com/fire833/morfic/pkg/apis/services -h boilerplate.go.txt --listers-package github.com/fire833/morfic/pkg/client/listers --single-directory -p services -o ../pkg/client/informers &
# echo "Rebuilding vpn API..."
# informer-gen --input-dirs github.com/fire833/morfic/pkg/apis/vpn -h boilerplate.go.txt --listers-package github.com/fire833/morfic/pkg/client/listers --single-directory -p vpn -o ../pkg/client/informers &

# wait
