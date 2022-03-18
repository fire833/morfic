#!/bin/bash

# export GOPATH=$(go env GOPATH)

GOPATH=$(go env GOPATH) ./generate-internal-groups.sh \
deepcopy,client,lister,informer \
github.com/fire833/vroute/pkg/client github.com/fire833/vroute/pkg/apis github.com/fire833/vroute/pkg/apis \
"addresses.vroute.io:v1alpha1 authentication.vroute.io:v1 certificates.vroute.io:v1alpha1 config.vroute.io:v1 dns.vroute.io:v1alpha1 firewall.vroute.io:v1alpha1 interfaces.vroute.io:v1alpha1 metrics.vroute.io:v1alpha1 nat.vroute.io:v1alpha1 neighbors.vroute.io:v1alpha1 routes.vroute.io:v1 services.vroute.io:v1alpha1 vpn.vroute.io:v1alpha1" \
--go-header-file boilerplate.go.txt \
# --output-base ../pkg/client

# wait
