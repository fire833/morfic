# Use to quickly build/update the go API for the node/service proto file.

GO		=	$(shell which go)

node:
	@echo "Rebuilding node control IPC API..."
	@protoc --go_out=. --go-grpc_out=. node_control.proto
	@sed -i 's/package __/package v1alpha1/g' node_control_grpc.pb.go
	@sed -i 's/package _/package v1alpha1/g' node_control.pb.go
	@echo "Success!"

service: 	setup
	@echo "Rebuilding service control IPC API..."
	@protoc --go_out=. --go-grpc_out=. service_control.proto
	@echo "Success!"

setup:
	
	@echo "Installing protoc and grpc plugin..."
	${GO} install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	${GO} install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	
	export PATH="${PATH}:$(go env GOPATH)/bin"