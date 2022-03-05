# Use to quickly build/update the go API for the node/service proto file.

node:
	protoc --go_out=. --go-grpc_out=. node_control.proto
	sed -i 's/package __/package v1alpha1/g' node_control_grpc.pb.go
	sed -i 's/package _/package v1alpha1/g' node_control.pb.go

service: 	setup
	protoc --go_out=. --go-grpc_out=. service_control.proto


setup:
	export PATH="${PATH}:$(go env GOPATH)/bin"