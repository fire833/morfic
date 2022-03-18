#!/bin/bash

# export GOPATH=$(go env GOPATH)

GOPATH=$(go env GOPATH) ./generate-internal-groups.sh \
client,lister,informer \
github.com/fire833/vroute/pkg/client github.com/fire833/vroute/pkg/apis github.com/fire833/vroute/pkg/apis \
"addresses:v1alpha1 authentication:v1 certificates:v1alpha1 config:v1 dns:v1alpha1 firewall:v1alpha1 interfaces:v1alpha1 metrics:v1alpha1 nat:v1alpha1 neighbors:v1alpha1 routes:v1 services:v1alpha1 vpn:v1alpha1" \
--go-header-file boilerplate.go.txt \
# --output-base ../pkg/client

# GOPATH=$(go env GOPATH) ./generate-groups.sh \
# all \
# github.com/fire833/vroute/pkg/client github.com/fire833/vroute/pkg/apis \
# "addresses:v1alpha1 authentication:v1 certificates:v1alpha1 config:v1 dns:v1alpha1 firewall:v1alpha1 interfaces:v1alpha1 metrics:v1alpha1 nat:v1alpha1 neighbors:v1alpha1 routes:v1 services:v1alpha1 vpn:v1alpha1" \
# --go-header-file boilerplate.go.txt \

# wait
