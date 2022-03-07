#!/bin/bash

informer-gen -i \
github.com/fire833/vroute/pkg/apis/addresses,\
github.com/fire833/vroute/pkg/apis/authentication,\
github.com/fire833/vroute/pkg/apis/certificates,\
github.com/fire833/vroute/pkg/apis/config,\
github.com/fire833/vroute/pkg/apis/dns,\
github.com/fire833/vroute/pkg/apis/firewall,\
github.com/fire833/vroute/pkg/apis/interfaces,\
github.com/fire833/vroute/pkg/apis/metrics,\
github.com/fire833/vroute/pkg/apis/nat,\
github.com/fire833/vroute/pkg/apis/neighbors,\
github.com/fire833/vroute/pkg/apis/routes,\
github.com/fire833/vroute/pkg/apis/services,\
github.com/fire833/vroute/pkg/apis/vpn \
-h boilerplate.go.txt --listers-package github.com/fire833/vroute/pkg/client/listers --single-directory -o ../pkg/client -p informers

# echo "Rebuilding addresses API..."
# informer-gen --input-dirs github.com/fire833/vroute/pkg/apis/addresses -h boilerplate.go.txt --listers-package github.com/fire833/vroute/pkg/client/listers --single-directory -p addresses -o ../pkg/client/informers &
# echo "Rebuilding authentication API..."
# informer-gen --input-dirs github.com/fire833/vroute/pkg/apis/authentication -h boilerplate.go.txt --listers-package github.com/fire833/vroute/pkg/client/listers --single-directory -p authetication -o ../pkg/client/informers &
# echo "Rebuilding certificates API..."
# informer-gen --input-dirs github.com/fire833/vroute/pkg/apis/certificates -h boilerplate.go.txt --listers-package github.com/fire833/vroute/pkg/client/listers --single-directory -p certificates -o ../pkg/client/informers &
# echo "Rebuilding config API..."
# informer-gen --input-dirs github.com/fire833/vroute/pkg/apis/config -h boilerplate.go.txt --listers-package github.com/fire833/vroute/pkg/client/listers --single-directory -p config -o ../pkg/client/informers &
# echo "Rebuilding dns API..."
# informer-gen --input-dirs github.com/fire833/vroute/pkg/apis/dns -h boilerplate.go.txt --listers-package github.com/fire833/vroute/pkg/client/listers --single-directory -p dns -o ../pkg/client/informers & 
# echo "Rebuilding firewall API..."
# informer-gen --input-dirs github.com/fire833/vroute/pkg/apis/firewall -h boilerplate.go.txt --listers-package github.com/fire833/vroute/pkg/client/listers --single-directory -p firewall -o ../pkg/client/informers & 
# echo "Rebuilding interfaces API..."
# informer-gen --input-dirs github.com/fire833/vroute/pkg/apis/interfaces -h boilerplate.go.txt --listers-package github.com/fire833/vroute/pkg/client/listers --single-directory -p interfaces -o ../pkg/client/informers &
# echo "Rebuilding metrics API..."
# informer-gen --input-dirs github.com/fire833/vroute/pkg/apis/metrics -h boilerplate.go.txt --listers-package github.com/fire833/vroute/pkg/client/listers --single-directory -p metrics -o ../pkg/client/informers &
# echo "Rebuilding nat API..."
# informer-gen --input-dirs github.com/fire833/vroute/pkg/apis/nat -h boilerplate.go.txt --listers-package github.com/fire833/vroute/pkg/client/listers --single-directory -p nat -o ../pkg/client/informers &
# echo "Rebuilding neighbors API..."
# informer-gen --input-dirs github.com/fire833/vroute/pkg/apis/neighbors -h boilerplate.go.txt --listers-package github.com/fire833/vroute/pkg/client/listers --single-directory -p neighbors -o ../pkg/client/informers &
# echo "Rebuilding routes API..."
# informer-gen --input-dirs github.com/fire833/vroute/pkg/apis/routes -h boilerplate.go.txt --listers-package github.com/fire833/vroute/pkg/client/listers --single-directory -p routes -o ../pkg/client/informers &
# echo "Rebuilding services API..."
# informer-gen --input-dirs github.com/fire833/vroute/pkg/apis/services -h boilerplate.go.txt --listers-package github.com/fire833/vroute/pkg/client/listers --single-directory -p services -o ../pkg/client/informers &
# echo "Rebuilding vpn API..."
# informer-gen --input-dirs github.com/fire833/vroute/pkg/apis/vpn -h boilerplate.go.txt --listers-package github.com/fire833/vroute/pkg/client/listers --single-directory -p vpn -o ../pkg/client/informers &

# wait
