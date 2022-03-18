#!/bin/bash

echo "Updating API clients..."

client-gen -i \
"github.com/fire833/vroute/pkg/apis/addresses/v1alpha1,\
github.com/fire833/vroute/pkg/apis/authentication/v1,\
github.com/fire833/vroute/pkg/apis/certificates/v1alpha1,\
github.com/fire833/vroute/pkg/apis/config/v1,\
github.com/fire833/vroute/pkg/apis/dns/v1alpha1,\
github.com/fire833/vroute/pkg/apis/firewall/v1alpha1,\
github.com/fire833/vroute/pkg/apis/interfaces/v1,\
github.com/fire833/vroute/pkg/apis/metrics/v1alpha1,\
github.com/fire833/vroute/pkg/apis/nat/v1alpha1,\
github.com/fire833/vroute/pkg/apis/neighbors/v1alpha1,\
github.com/fire833/vroute/pkg/apis/routes/v1,\
github.com/fire833/vroute/pkg/apis/services/v1alpha1,\
github.com/fire833/vroute/pkg/apis/vpn/v1alpha1, \
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
github.com/fire833/vroute/pkg/apis/vpn" \
-h boilerplate.go.txt -o "../vendor" -p "github.com/fire833/client-go/vroute" # -n clientset

# mv ../pkg/client/clientset/clientset/* ../pkg/client/clientset

# rm -rf ../pkg/client/clientset/clientset