#!/bin/bash

echo "Rebuilding addresses API..."
deepcopy-gen --input-dirs ../pkg/apis/addresses -h boilerplate.go.txt -O zz_generated.deepcopy &
echo "Rebuilding authentication API..."
deepcopy-gen --input-dirs ../pkg/apis/authentication -h boilerplate.go.txt -O zz_generated.deepcopy &
echo "Rebuilding certificates API..."
deepcopy-gen --input-dirs ../pkg/apis/certificates -h boilerplate.go.txt -O zz_generated.deepcopy &
echo "Rebuilding config API..."
deepcopy-gen --input-dirs ../pkg/apis/config -h boilerplate.go.txt -O zz_generated.deepcopy &
echo "Rebuilding dns API..."
deepcopy-gen --input-dirs ../pkg/apis/dns -h boilerplate.go.txt -O zz_generated.deepcopy &
echo "Rebuilding firewall API..."
deepcopy-gen --input-dirs ../pkg/apis/firewall -h boilerplate.go.txt -O zz_generated.deepcopy &
echo "Rebuilding interfaces API..."
deepcopy-gen --input-dirs ../pkg/apis/interfaces -h boilerplate.go.txt -O zz_generated.deepcopy & 
echo "Rebuilding metrics API..."
deepcopy-gen --input-dirs ../pkg/apis/metrics -h boilerplate.go.txt -O zz_generated.deepcopy &
echo "Rebuilding nat API..."
deepcopy-gen --input-dirs ../pkg/apis/nat -h boilerplate.go.txt -O zz_generated.deepcopy &
echo "Rebuilding neighbors API..."
deepcopy-gen --input-dirs ../pkg/apis/neighbors -h boilerplate.go.txt -O zz_generated.deepcopy &
echo "Rebuilding routes API..."
deepcopy-gen --input-dirs ../pkg/apis/routes -h boilerplate.go.txt -O zz_generated.deepcopy &
echo "Rebuilding services API..."
deepcopy-gen --input-dirs ../pkg/apis/services -h boilerplate.go.txt -O zz_generated.deepcopy &
echo "Rebuilding vpn API..."
deepcopy-gen --input-dirs ../pkg/apis/vpn -h boilerplate.go.txt -O zz_generated.deepcopy &

wait
