#!/bin/bash

echo "Rebuilding addresses API..."
informer-gen --input-dirs github.com/fire833/vroute/pkg/apis/addresses -h boilerplate.go.txt --listers-package github.com/fire833/vroute/pkg/client/listers -p addresses -o ../pkg/client/informers &
echo "Rebuilding authentication API..."
informer-gen --input-dirs github.com/fire833/vroute/pkg/apis/authentication -h boilerplate.go.txt --listers-package github.com/fire833/vroute/pkg/client/listers -p authetication -o ../pkg/client/informers &
echo "Rebuilding certificates API..."
informer-gen --input-dirs github.com/fire833/vroute/pkg/apis/certificates -h boilerplate.go.txt --listers-package github.com/fire833/vroute/pkg/client/listers -p certificates -o ../pkg/client/informers &
echo "Rebuilding config API..."
informer-gen --input-dirs github.com/fire833/vroute/pkg/apis/config -h boilerplate.go.txt --listers-package github.com/fire833/vroute/pkg/client/listers -p config -o ../pkg/client/informers &
echo "Rebuilding dns API..."
informer-gen --input-dirs github.com/fire833/vroute/pkg/apis/dns -h boilerplate.go.txt --listers-package github.com/fire833/vroute/pkg/client/listers -p dns -o ../pkg/client/informers & 
echo "Rebuilding firewall API..."
informer-gen --input-dirs github.com/fire833/vroute/pkg/apis/firewall -h boilerplate.go.txt --listers-package github.com/fire833/vroute/pkg/client/listers -p firewall -o ../pkg/client/informers & 
echo "Rebuilding interfaces API..."
informer-gen --input-dirs github.com/fire833/vroute/pkg/apis/interfaces -h boilerplate.go.txt --listers-package github.com/fire833/vroute/pkg/client/listers -p interfaces -o ../pkg/client/informers &
echo "Rebuilding metrics API..."
informer-gen --input-dirs github.com/fire833/vroute/pkg/apis/metrics -h boilerplate.go.txt --listers-package github.com/fire833/vroute/pkg/client/listers -p metrics -o ../pkg/client/informers &
echo "Rebuilding nat API..."
informer-gen --input-dirs github.com/fire833/vroute/pkg/apis/nat -h boilerplate.go.txt --listers-package github.com/fire833/vroute/pkg/client/listers -p nat -o ../pkg/client/informers &
echo "Rebuilding neighbors API..."
informer-gen --input-dirs github.com/fire833/vroute/pkg/apis/neighbors -h boilerplate.go.txt --listers-package github.com/fire833/vroute/pkg/client/listers -p neighbors -o ../pkg/client/informers &
echo "Rebuilding routes API..."
informer-gen --input-dirs github.com/fire833/vroute/pkg/apis/routes -h boilerplate.go.txt --listers-package github.com/fire833/vroute/pkg/client/listers -p routes -o ../pkg/client/informers &
echo "Rebuilding services API..."
informer-gen --input-dirs github.com/fire833/vroute/pkg/apis/services -h boilerplate.go.txt --listers-package github.com/fire833/vroute/pkg/client/listers -p services -o ../pkg/client/informers &
echo "Rebuilding vpn API..."
informer-gen --input-dirs github.com/fire833/vroute/pkg/apis/vpn -h boilerplate.go.txt --listers-package github.com/fire833/vroute/pkg/client/listers -p vpn -o ../pkg/client/informers &

wait