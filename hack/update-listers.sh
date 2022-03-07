#!/bin/bash

lister-gen -i \
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
-h boilerplate.go.txt -o ../pkg/client -p listers

# echo "Rebuilding addresses API..."
# lister-gen --input-dirs github.com/fire833/vroute/pkg/apis/addresses -h boilerplate.go.txt -p addresses -o ../pkg/client/listers &
# echo "Rebuilding authentication API..."
# lister-gen --input-dirs github.com/fire833/vroute/pkg/apis/authentication -h boilerplate.go.txt -p authentication -o ../pkg/client/listers &
# echo "Rebuilding certificates API..."
# lister-gen --input-dirs github.com/fire833/vroute/pkg/apis/certificates -h boilerplate.go.txt -p certificates -o ../pkg/client/listers &
# echo "Rebuilding config API..."
# lister-gen --input-dirs github.com/fire833/vroute/pkg/apis/config -h boilerplate.go.txt -p config -o ../pkg/client/listers &
# echo "Rebuilding dns API..."
# lister-gen --input-dirs github.com/fire833/vroute/pkg/apis/dns -h boilerplate.go.txt -p dns -o ../pkg/client/listers & 
# echo "Rebuilding firewall API..."
# lister-gen --input-dirs github.com/fire833/vroute/pkg/apis/firewall -h boilerplate.go.txt -p firewall -o ../pkg/client/listers & 
# echo "Rebuilding interfaces API..."
# lister-gen --input-dirs github.com/fire833/vroute/pkg/apis/interfaces -h boilerplate.go.txt -p interfaces -o ../pkg/client/listers &
# echo "Rebuilding metrics API..."
# lister-gen --input-dirs github.com/fire833/vroute/pkg/apis/metrics -h boilerplate.go.txt -p metrics -o ../pkg/client/listers &
# echo "Rebuilding nat API..."
# lister-gen --input-dirs github.com/fire833/vroute/pkg/apis/nat -h boilerplate.go.txt -p nat -o ../pkg/client/listers &
# echo "Rebuilding neighbors API..."
# lister-gen --input-dirs github.com/fire833/vroute/pkg/apis/neighbors -h boilerplate.go.txt -p neighbors -o ../pkg/client/listers &
# echo "Rebuilding routes API..."
# lister-gen --input-dirs github.com/fire833/vroute/pkg/apis/routes -h boilerplate.go.txt -p routes -o ../pkg/client/listers &
# echo "Rebuilding services API..."
# lister-gen --input-dirs github.com/fire833/vroute/pkg/apis/services -h boilerplate.go.txt -p services -o ../pkg/client/listers &
# echo "Rebuilding vpn API..."
# lister-gen --input-dirs github.com/fire833/vroute/pkg/apis/vpn -h boilerplate.go.txt -p vpn -o ../pkg/client/listers &

# wait
