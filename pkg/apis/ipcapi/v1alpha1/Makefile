# Use to quickly build/update the go API for the node/service proto file.

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
	export PATH="${PATH}:$(go env GOPATH)/bin"