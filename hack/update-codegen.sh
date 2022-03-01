#!/bin/bash

export GOPATH=$(go env GOPATH)

./generate-internal-groups.sh all \
github.com/fire833/vroute/pkg/client github.com/fire833/vroute/apis \
"certificates:v1alpha1" \ 
-h boilerplate.go.txt -v 2
