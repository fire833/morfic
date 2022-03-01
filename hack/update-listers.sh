#!/bin/bash

echo "Rebuilding addresses API..."
lister-gen --input-dirs ../pkg/apis/addresses -h boilerplate.go.txt -O zz_generated.lister
echo "Rebuilding authentication API..."
lister-gen --input-dirs ../pkg/apis/authentication -h boilerplate.go.txt -O zz_generated.lister
echo "Rebuilding config API..."
lister-gen --input-dirs ../pkg/apis/config -h boilerplate.go.txt -O zz_generated.lister
echo "Rebuilding dns API..."
lister-gen --input-dirs ../pkg/apis/dns -h boilerplate.go.txt -O zz_generated.lister
echo "Rebuilding firewall API..."
lister-gen --input-dirs ../pkg/apis/firewall -h boilerplate.go.txt -O zz_generated.lister
echo "Rebuilding interfaces API..."
lister-gen --input-dirs ../pkg/apis/interfaces -h boilerplate.go.txt -O zz_generated.lister
echo "Rebuilding metrics API..."
lister-gen --input-dirs ../pkg/apis/metrics -h boilerplate.go.txt -O zz_generated.lister
echo "Rebuilding nat API..."
lister-gen --input-dirs ../pkg/apis/nat -h boilerplate.go.txt -O zz_generated.lister
echo "Rebuilding neighbors API..."
lister-gen --input-dirs ../pkg/apis/neighbors -h boilerplate.go.txt -O zz_generated.lister
echo "Rebuilding routes API..."
lister-gen --input-dirs ../pkg/apis/routes -h boilerplate.go.txt -O zz_generated.lister
echo "Rebuilding services API..."
lister-gen --input-dirs ../pkg/apis/services -h boilerplate.go.txt -O zz_generated.lister
echo "Rebuilding vpn API..."
lister-gen --input-dirs ../pkg/apis/vpn -h boilerplate.go.txt -O zz_generated.lister
