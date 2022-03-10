#!/bin/bash

echo "Updating API clients..."

client-gen -i \
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
-h boilerplate.go.txt -o ../pkg/client -p clientset -n clientset

mv ../pkg/client/clientset/clientset/* ../pkg/client/clientset

rm -rf ../pkg/client/clientset/clientset