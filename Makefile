#    Copyright (C) 2022  Kendall Tauser
#
#    This program is free software; you can redistribute it and/or modify
#    it under the terms of the GNU General Public License as published by
#    the Free Software Foundation; either version 2 of the License, or
#    (at your option) any later version.
#
#    This program is distributed in the hope that it will be useful,
#    but WITHOUT ANY WARRANTY; without even the implied warranty of
#    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
#    GNU General Public License for more details.
#
#    You should have received a copy of the GNU General Public License along
#    with this program; if not, write to the Free Software Foundation, Inc.,
#    51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA.

VERSION 	=	0.0.1

GO			=	$(shell which go)
COMMIT		=	$(shell git rev-parse HEAD)
DATE		=	$(shell date)

all:

binary:
	@echo "Building control plane from source..."
	@GOOS=linux GOARCH=amd64 ${GO} build -o bin/vroute cmd/vroute/main.go cmd/vroute/constants.go cmd/vroute/node.go cmd/vroute/api.go cmd/vroute/commands.go cmd/vroute/controller.go
	@# GOOS=windows GOARCH=amd64 ${GO} build -o bin/vroute.exe cmd/vroute/main.go cmd/vroute/constants.go cmd/vroute/node.go cmd/vroute/api.go cmd/vroute/commands.go cmd/vroute/controller.go
	@GOOS=linux GOARCH=arm64 ${GO} build -o bin/vroute-linux-arm64 cmd/vroute/main.go cmd/vroute/constants.go cmd/vroute/node.go cmd/vroute/api.go cmd/vroute/commands.go cmd/vroute/controller.go
	@GOOS=linux GOARCH=arm ${GO} build -o bin/vroute-linux-arm cmd/vroute/main.go cmd/vroute/constants.go cmd/vroute/node.go cmd/vroute/api.go cmd/vroute/commands.go cmd/vroute/controller.go
	@echo "Success!"

test:
	@echo "Testing vroute packages..."
	${GO} test -v ./...

vet:
	@echo "Vetting vroute source..."
	${GO} vet -v ./...

fmt:
	@echo "Formatting vroute source..."
	${GO} fmt ./...
