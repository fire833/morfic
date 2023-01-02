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

# Use to quickly build/update the go API for the node/service proto file.

GO		=	$(shell which go)
PROTOC	=	$(shell which protoc)
SED		=	$(shell which sed)

.PHONY: node
node:
	@echo "Rebuilding node control IPC API..."
	@${PROTOC} --go_out=. --go-grpc_out=. node_control.proto
	@${SED} -i 's/package __/package v1alpha1/g' node_control_grpc.pb.go
	@${SED} -i 's/package __/package v1alpha1/g' node_control.pb.go
	@echo "Success!"

.PHONY: service
service: 	setup
	@echo "Rebuilding service control IPC API..."
	@${PROTOC} --go_out=. --go-grpc_out=. service_control.proto
	@echo "Success!"

.PHONY: setup
setup:
	@echo "Installing protoc and grpc plugin..."
	${GO} install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	${GO} install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

.PHONY: env
env:
	@echo "Adding environment variables..."
	$(shell export PATH="${PATH}:$(go env GOPATH)/bin")