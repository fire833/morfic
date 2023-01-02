#    Copyright (C) 2023 Kendall Tauser
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

.PHONY: all
all:

.PHONY: binary
binary:
	@echo "Building control plane from source..."
	GOOS=linux GOARCH=amd64 ${GO} build -o bin/morfic -ldflags "-X 'main.version=${VERSION}' -X 'main.commit=${COMMIT}' -X 'main.buildTime=${DATE}'" cmd/morfic/main.go
	@# GOOS=windows GOARCH=amd64 ${GO} build -o bin/morfic.exe -ldflags "-X 'main.version=${VERSION}' -X 'main.commit=${COMMIT}' -X 'main.buildTime=${DATE}'" cmd/morfic/main.go 
	GOOS=linux GOARCH=arm64 ${GO} build -o bin/morfic-linux-arm64 -ldflags "-X 'main.version=${VERSION}' -X 'main.commit=${COMMIT}' -X 'main.buildTime=${DATE}'" cmd/morfic/main.go 
	GOOS=linux GOARCH=arm ${GO} build -o bin/morfic-linux-arm -ldflags "-X 'main.version=${VERSION}' -X 'main.commit=${COMMIT}' -X 'main.buildTime=${DATE}'" cmd/morfic/main.go 
	@echo "Success!"

.PHONY: codegen
codegen:	ipcapigen

.PHONY: ipcapigen
ipcapigen:
	make -C pkg/ipcapi/v1alpha1

.PHONY: test
test:
	@echo "Testing morfic packages..."
	${GO} test -v -cover ./pkg/...

.PHONY: vet
vet:
	@echo "Vetting morfic source..."
	${GO} vet -v ./...

.PHONY: fmt
fmt:
	@echo "Formatting morfic source..."
	${GO} fmt ./...
