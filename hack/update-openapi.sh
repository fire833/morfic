#!/bin/bash

echo "Rebuilding addresses API..."
openapi-gen --input-dirs ../pkg/apis/addresses -h boilerplate.go.txt -O zz_generated.openapi &
echo "Rebuilding authentication API..."
openapi-gen --input-dirs ../pkg/apis/authentication -h boilerplate.go.txt -O zz_generated.openapi &
echo "Rebuilding config API..."
openapi-gen --input-dirs ../pkg/apis/config -h boilerplate.go.txt -O zz_generated.openapi &
echo "Rebuilding dns API..."
openapi-gen --input-dirs ../pkg/apis/dns -h boilerplate.go.txt -O zz_generated.openapi &
echo "Rebuilding firewall API..."
openapi-gen --input-dirs ../pkg/apis/firewall -h boilerplate.go.txt -O zz_generated.openapi &
echo "Rebuilding interfaces API..."
openapi-gen --input-dirs ../pkg/apis/interfaces -h boilerplate.go.txt -O zz_generated.openapi &
echo "Rebuilding metrics API..."
openapi-gen --input-dirs ../pkg/apis/metrics -h boilerplate.go.txt -O zz_generated.openapi &
echo "Rebuilding nat API..."
openapi-gen --input-dirs ../pkg/apis/nat -h boilerplate.go.txt -O zz_generated.openapi &
echo "Rebuilding neighbors API..."
openapi-gen --input-dirs ../pkg/apis/neighbors -h boilerplate.go.txt -O zz_generated.openapi &
echo "Rebuilding routes API..."
openapi-gen --input-dirs ../pkg/apis/routes -h boilerplate.go.txt -O zz_generated.openapi &
echo "Rebuilding services API..."
openapi-gen --input-dirs ../pkg/apis/services -h boilerplate.go.txt -O zz_generated.openapi &
echo "Rebuilding vpn API..."
openapi-gen --input-dirs ../pkg/apis/vpn -h boilerplate.go.txt -O zz_generated.openapi &

wait
