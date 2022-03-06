#!/bin/bash

echo "Rebuilding addresses API..."
lister-gen --input-dirs ../pkg/apis/addresses -h boilerplate.go.txt -O zz_generated.lister -p ../pkg/client/listers &
echo "Rebuilding authentication API..."
lister-gen --input-dirs ../pkg/apis/authentication -h boilerplate.go.txt -O zz_generated.lister -p ../pkg/client/listers &
echo "Rebuilding certificates API..."
lister-gen --input-dirs ../pkg/apis/certificates -h boilerplate.go.txt -O zz_generated.lister -p ../pkg/client/listers &
echo "Rebuilding config API..."
lister-gen --input-dirs ../pkg/apis/config -h boilerplate.go.txt -O zz_generated.lister -p ../pkg/client/listers &
echo "Rebuilding dns API..."
lister-gen --input-dirs ../pkg/apis/dns -h boilerplate.go.txt -O zz_generated.lister -p ../pkg/client/listers & 
echo "Rebuilding firewall API..."
lister-gen --input-dirs ../pkg/apis/firewall -h boilerplate.go.txt -O zz_generated.lister -p ../pkg/client/listers & 
echo "Rebuilding interfaces API..."
lister-gen --input-dirs ../pkg/apis/interfaces -h boilerplate.go.txt -O zz_generated.lister -p ../pkg/client/listers &
echo "Rebuilding metrics API..."
lister-gen --input-dirs ../pkg/apis/metrics -h boilerplate.go.txt -O zz_generated.lister -p ../pkg/client/listers &
echo "Rebuilding nat API..."
lister-gen --input-dirs ../pkg/apis/nat -h boilerplate.go.txt -O zz_generated.lister -p ../pkg/client/listers &
echo "Rebuilding neighbors API..."
lister-gen --input-dirs ../pkg/apis/neighbors -h boilerplate.go.txt -O zz_generated.lister -p ../pkg/client/listers &
echo "Rebuilding routes API..."
lister-gen --input-dirs ../pkg/apis/routes -h boilerplate.go.txt -O zz_generated.lister -p ../pkg/client/listers &
echo "Rebuilding services API..."
lister-gen --input-dirs ../pkg/apis/services -h boilerplate.go.txt -O zz_generated.lister -p ../pkg/client/listers &
echo "Rebuilding vpn API..."
lister-gen --input-dirs ../pkg/apis/vpn -h boilerplate.go.txt -O zz_generated.lister -p ../pkg/client/listers &

wait