#!/bin/bash

echo "Rebuilding addresses API..."
informer-gen --input-dirs ../pkg/apis/addresses -h boilerplate.go.txt -O zz_generated.informer -o ../pkg/client -p ../pkg/client/informers &
echo "Rebuilding authentication API..."
informer-gen --input-dirs ../pkg/apis/authentication -h boilerplate.go.txt -O zz_generated.informer -o ../pkg/client -p ../pkg/client/informers &
echo "Rebuilding certificates API..."
informer-gen --input-dirs ../pkg/apis/certificates -h boilerplate.go.txt -O zz_generated.informer --versioned-clientset-package github.com/fire833/vroute/pkg/client/vroute -o github.com/fire833/vroute/pkg/client/informers --single-directory --listers-package github.com/fire833/vroute/pkg/client/listers -p ../pkg/client/informers &
echo "Rebuilding config API..."
informer-gen --input-dirs ../pkg/apis/config -h boilerplate.go.txt -O zz_generated.informer -o ../pkg/client -p ../pkg/client/informers &
echo "Rebuilding dns API..."
informer-gen --input-dirs ../pkg/apis/dns -h boilerplate.go.txt -O zz_generated.informer -o ../pkg/client -p ../pkg/client/informers & 
echo "Rebuilding firewall API..."
informer-gen --input-dirs ../pkg/apis/firewall -h boilerplate.go.txt -O zz_generated.informer -o ../pkg/client -p ../pkg/client/informers & 
echo "Rebuilding interfaces API..."
informer-gen --input-dirs ../pkg/apis/interfaces -h boilerplate.go.txt -O zz_generated.informer -o ../pkg/client -p ../pkg/client/informers &
echo "Rebuilding metrics API..."
informer-gen --input-dirs ../pkg/apis/metrics -h boilerplate.go.txt -O zz_generated.informer -o ../pkg/client -p ../pkg/client/informers &
echo "Rebuilding nat API..."
informer-gen --input-dirs ../pkg/apis/nat -h boilerplate.go.txt -O zz_generated.informer -o ../pkg/client -p ../pkg/client/informers &
echo "Rebuilding neighbors API..."
informer-gen --input-dirs ../pkg/apis/neighbors -h boilerplate.go.txt -O zz_generated.informer -o ../pkg/client -p ../pkg/client/informers &
echo "Rebuilding routes API..."
informer-gen --input-dirs ../pkg/apis/routes -h boilerplate.go.txt -O zz_generated.informer -o ../pkg/client -p ../pkg/client/informers &
echo "Rebuilding services API..."
informer-gen --input-dirs ../pkg/apis/services -h boilerplate.go.txt -O zz_generated.informer -o ../pkg/client -p ../pkg/client/informers &
echo "Rebuilding vpn API..."
informer-gen --input-dirs ../pkg/apis/vpn -h boilerplate.go.txt -O zz_generated.informer -o ../pkg/client -p github.com/fire833/vroute/pkg/client/informers &

wait