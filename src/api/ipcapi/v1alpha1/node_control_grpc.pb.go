// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: node_control.proto

package v1alpha1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// NodeControllerServiceClient is the client API for NodeControllerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NodeControllerServiceClient interface {
	// Create a new link inside of the kernel by specifying its type and other attributes.
	CreateLink(ctx context.Context, in *CreateLinkRequest, opts ...grpc.CallOption) (*CreateLinkResponse, error)
	// Update a link status/attributes.
	UpdateLink(ctx context.Context, in *UpdateLinkRequest, opts ...grpc.CallOption) (*UpdateLinkResponse, error)
	// Remove a linnk inside of the kernel.
	DeleteLink(ctx context.Context, in *DeleteLinkRequest, opts ...grpc.CallOption) (*DeleteLinkResponse, error)
	// Get a link object returned from the kernel based on filters.
	GetLink(ctx context.Context, in *GetLinkRequest, opts ...grpc.CallOption) (*GetLinkResponse, error)
	// Returns all the links in the kernel.
	GetAllLinks(ctx context.Context, in *GetAllLinksRequest, opts ...grpc.CallOption) (*GetAllLinksResponse, error)
	// Add a static route inside of the kernel.
	CreateStaticRoute(ctx context.Context, in *CreateStaticRouteRequest, opts ...grpc.CallOption) (*CreateStaticRouteResponse, error)
	// Deletes a static route inside the kernel and returns the route object.
	DeleteStaticRoute(ctx context.Context, in *DeleteStaticRouteRequest, opts ...grpc.CallOption) (*DeleteStaticRouteResponse, error)
	// Updates a static route with new parameters.
	UpdateStaticRoute(ctx context.Context, in *UpdateStaticRouteRequest, opts ...grpc.CallOption) (*UpdateStaticRouteResponse, error)
	// Gets a route and returns an object based on the provided filter.
	GetRoute(ctx context.Context, in *GetRouteRequest, opts ...grpc.CallOption) (*GetRouteResponse, error)
	GetAllRoutes(ctx context.Context, in *GetAllRoutesRequest, opts ...grpc.CallOption) (*GetAllRoutesResponse, error)
	GetNeighbor(ctx context.Context, in *GetNeighborRequest, opts ...grpc.CallOption) (*GetNeighborResponse, error)
	GetAllNeighbors(ctx context.Context, in *GetAllNeighborsRequest, opts ...grpc.CallOption) (*GetAllNeighborsResponse, error)
	CreateNeighbor(ctx context.Context, in *CreateNeighborRequest, opts ...grpc.CallOption) (*CreateNeighborResponse, error)
	DeleteNeighbor(ctx context.Context, in *DeleteNeighborRequest, opts ...grpc.CallOption) (*DeleteNeighborRequest, error)
	UpdateNeighbor(ctx context.Context, in *UpdateNeighborRequest, opts ...grpc.CallOption) (*UpdateNeighborResponse, error)
	DeleteAddress(ctx context.Context, in *DeleteAddressRequest, opts ...grpc.CallOption) (*DeleteAddressResponse, error)
	UpdateAddress(ctx context.Context, in *UpdateAddressRequest, opts ...grpc.CallOption) (*UpdateAddressResponse, error)
}

type nodeControllerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeControllerServiceClient(cc grpc.ClientConnInterface) NodeControllerServiceClient {
	return &nodeControllerServiceClient{cc}
}

func (c *nodeControllerServiceClient) CreateLink(ctx context.Context, in *CreateLinkRequest, opts ...grpc.CallOption) (*CreateLinkResponse, error) {
	out := new(CreateLinkResponse)
	err := c.cc.Invoke(ctx, "/v1alpha1.NodeControllerService/CreateLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeControllerServiceClient) UpdateLink(ctx context.Context, in *UpdateLinkRequest, opts ...grpc.CallOption) (*UpdateLinkResponse, error) {
	out := new(UpdateLinkResponse)
	err := c.cc.Invoke(ctx, "/v1alpha1.NodeControllerService/UpdateLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeControllerServiceClient) DeleteLink(ctx context.Context, in *DeleteLinkRequest, opts ...grpc.CallOption) (*DeleteLinkResponse, error) {
	out := new(DeleteLinkResponse)
	err := c.cc.Invoke(ctx, "/v1alpha1.NodeControllerService/DeleteLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeControllerServiceClient) GetLink(ctx context.Context, in *GetLinkRequest, opts ...grpc.CallOption) (*GetLinkResponse, error) {
	out := new(GetLinkResponse)
	err := c.cc.Invoke(ctx, "/v1alpha1.NodeControllerService/GetLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeControllerServiceClient) GetAllLinks(ctx context.Context, in *GetAllLinksRequest, opts ...grpc.CallOption) (*GetAllLinksResponse, error) {
	out := new(GetAllLinksResponse)
	err := c.cc.Invoke(ctx, "/v1alpha1.NodeControllerService/GetAllLinks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeControllerServiceClient) CreateStaticRoute(ctx context.Context, in *CreateStaticRouteRequest, opts ...grpc.CallOption) (*CreateStaticRouteResponse, error) {
	out := new(CreateStaticRouteResponse)
	err := c.cc.Invoke(ctx, "/v1alpha1.NodeControllerService/CreateStaticRoute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeControllerServiceClient) DeleteStaticRoute(ctx context.Context, in *DeleteStaticRouteRequest, opts ...grpc.CallOption) (*DeleteStaticRouteResponse, error) {
	out := new(DeleteStaticRouteResponse)
	err := c.cc.Invoke(ctx, "/v1alpha1.NodeControllerService/DeleteStaticRoute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeControllerServiceClient) UpdateStaticRoute(ctx context.Context, in *UpdateStaticRouteRequest, opts ...grpc.CallOption) (*UpdateStaticRouteResponse, error) {
	out := new(UpdateStaticRouteResponse)
	err := c.cc.Invoke(ctx, "/v1alpha1.NodeControllerService/UpdateStaticRoute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeControllerServiceClient) GetRoute(ctx context.Context, in *GetRouteRequest, opts ...grpc.CallOption) (*GetRouteResponse, error) {
	out := new(GetRouteResponse)
	err := c.cc.Invoke(ctx, "/v1alpha1.NodeControllerService/GetRoute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeControllerServiceClient) GetAllRoutes(ctx context.Context, in *GetAllRoutesRequest, opts ...grpc.CallOption) (*GetAllRoutesResponse, error) {
	out := new(GetAllRoutesResponse)
	err := c.cc.Invoke(ctx, "/v1alpha1.NodeControllerService/GetAllRoutes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeControllerServiceClient) GetNeighbor(ctx context.Context, in *GetNeighborRequest, opts ...grpc.CallOption) (*GetNeighborResponse, error) {
	out := new(GetNeighborResponse)
	err := c.cc.Invoke(ctx, "/v1alpha1.NodeControllerService/GetNeighbor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeControllerServiceClient) GetAllNeighbors(ctx context.Context, in *GetAllNeighborsRequest, opts ...grpc.CallOption) (*GetAllNeighborsResponse, error) {
	out := new(GetAllNeighborsResponse)
	err := c.cc.Invoke(ctx, "/v1alpha1.NodeControllerService/GetAllNeighbors", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeControllerServiceClient) CreateNeighbor(ctx context.Context, in *CreateNeighborRequest, opts ...grpc.CallOption) (*CreateNeighborResponse, error) {
	out := new(CreateNeighborResponse)
	err := c.cc.Invoke(ctx, "/v1alpha1.NodeControllerService/CreateNeighbor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeControllerServiceClient) DeleteNeighbor(ctx context.Context, in *DeleteNeighborRequest, opts ...grpc.CallOption) (*DeleteNeighborRequest, error) {
	out := new(DeleteNeighborRequest)
	err := c.cc.Invoke(ctx, "/v1alpha1.NodeControllerService/DeleteNeighbor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeControllerServiceClient) UpdateNeighbor(ctx context.Context, in *UpdateNeighborRequest, opts ...grpc.CallOption) (*UpdateNeighborResponse, error) {
	out := new(UpdateNeighborResponse)
	err := c.cc.Invoke(ctx, "/v1alpha1.NodeControllerService/UpdateNeighbor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeControllerServiceClient) DeleteAddress(ctx context.Context, in *DeleteAddressRequest, opts ...grpc.CallOption) (*DeleteAddressResponse, error) {
	out := new(DeleteAddressResponse)
	err := c.cc.Invoke(ctx, "/v1alpha1.NodeControllerService/DeleteAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeControllerServiceClient) UpdateAddress(ctx context.Context, in *UpdateAddressRequest, opts ...grpc.CallOption) (*UpdateAddressResponse, error) {
	out := new(UpdateAddressResponse)
	err := c.cc.Invoke(ctx, "/v1alpha1.NodeControllerService/UpdateAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeControllerServiceServer is the server API for NodeControllerService service.
// All implementations must embed UnimplementedNodeControllerServiceServer
// for forward compatibility
type NodeControllerServiceServer interface {
	// Create a new link inside of the kernel by specifying its type and other attributes.
	CreateLink(context.Context, *CreateLinkRequest) (*CreateLinkResponse, error)
	// Update a link status/attributes.
	UpdateLink(context.Context, *UpdateLinkRequest) (*UpdateLinkResponse, error)
	// Remove a linnk inside of the kernel.
	DeleteLink(context.Context, *DeleteLinkRequest) (*DeleteLinkResponse, error)
	// Get a link object returned from the kernel based on filters.
	GetLink(context.Context, *GetLinkRequest) (*GetLinkResponse, error)
	// Returns all the links in the kernel.
	GetAllLinks(context.Context, *GetAllLinksRequest) (*GetAllLinksResponse, error)
	// Add a static route inside of the kernel.
	CreateStaticRoute(context.Context, *CreateStaticRouteRequest) (*CreateStaticRouteResponse, error)
	// Deletes a static route inside the kernel and returns the route object.
	DeleteStaticRoute(context.Context, *DeleteStaticRouteRequest) (*DeleteStaticRouteResponse, error)
	// Updates a static route with new parameters.
	UpdateStaticRoute(context.Context, *UpdateStaticRouteRequest) (*UpdateStaticRouteResponse, error)
	// Gets a route and returns an object based on the provided filter.
	GetRoute(context.Context, *GetRouteRequest) (*GetRouteResponse, error)
	GetAllRoutes(context.Context, *GetAllRoutesRequest) (*GetAllRoutesResponse, error)
	GetNeighbor(context.Context, *GetNeighborRequest) (*GetNeighborResponse, error)
	GetAllNeighbors(context.Context, *GetAllNeighborsRequest) (*GetAllNeighborsResponse, error)
	CreateNeighbor(context.Context, *CreateNeighborRequest) (*CreateNeighborResponse, error)
	DeleteNeighbor(context.Context, *DeleteNeighborRequest) (*DeleteNeighborRequest, error)
	UpdateNeighbor(context.Context, *UpdateNeighborRequest) (*UpdateNeighborResponse, error)
	DeleteAddress(context.Context, *DeleteAddressRequest) (*DeleteAddressResponse, error)
	UpdateAddress(context.Context, *UpdateAddressRequest) (*UpdateAddressResponse, error)
	mustEmbedUnimplementedNodeControllerServiceServer()
}

// UnimplementedNodeControllerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNodeControllerServiceServer struct {
}

func (UnimplementedNodeControllerServiceServer) CreateLink(context.Context, *CreateLinkRequest) (*CreateLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLink not implemented")
}
func (UnimplementedNodeControllerServiceServer) UpdateLink(context.Context, *UpdateLinkRequest) (*UpdateLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateLink not implemented")
}
func (UnimplementedNodeControllerServiceServer) DeleteLink(context.Context, *DeleteLinkRequest) (*DeleteLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteLink not implemented")
}
func (UnimplementedNodeControllerServiceServer) GetLink(context.Context, *GetLinkRequest) (*GetLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLink not implemented")
}
func (UnimplementedNodeControllerServiceServer) GetAllLinks(context.Context, *GetAllLinksRequest) (*GetAllLinksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllLinks not implemented")
}
func (UnimplementedNodeControllerServiceServer) CreateStaticRoute(context.Context, *CreateStaticRouteRequest) (*CreateStaticRouteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStaticRoute not implemented")
}
func (UnimplementedNodeControllerServiceServer) DeleteStaticRoute(context.Context, *DeleteStaticRouteRequest) (*DeleteStaticRouteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteStaticRoute not implemented")
}
func (UnimplementedNodeControllerServiceServer) UpdateStaticRoute(context.Context, *UpdateStaticRouteRequest) (*UpdateStaticRouteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStaticRoute not implemented")
}
func (UnimplementedNodeControllerServiceServer) GetRoute(context.Context, *GetRouteRequest) (*GetRouteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoute not implemented")
}
func (UnimplementedNodeControllerServiceServer) GetAllRoutes(context.Context, *GetAllRoutesRequest) (*GetAllRoutesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllRoutes not implemented")
}
func (UnimplementedNodeControllerServiceServer) GetNeighbor(context.Context, *GetNeighborRequest) (*GetNeighborResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNeighbor not implemented")
}
func (UnimplementedNodeControllerServiceServer) GetAllNeighbors(context.Context, *GetAllNeighborsRequest) (*GetAllNeighborsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllNeighbors not implemented")
}
func (UnimplementedNodeControllerServiceServer) CreateNeighbor(context.Context, *CreateNeighborRequest) (*CreateNeighborResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNeighbor not implemented")
}
func (UnimplementedNodeControllerServiceServer) DeleteNeighbor(context.Context, *DeleteNeighborRequest) (*DeleteNeighborRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNeighbor not implemented")
}
func (UnimplementedNodeControllerServiceServer) UpdateNeighbor(context.Context, *UpdateNeighborRequest) (*UpdateNeighborResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNeighbor not implemented")
}
func (UnimplementedNodeControllerServiceServer) DeleteAddress(context.Context, *DeleteAddressRequest) (*DeleteAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAddress not implemented")
}
func (UnimplementedNodeControllerServiceServer) UpdateAddress(context.Context, *UpdateAddressRequest) (*UpdateAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAddress not implemented")
}
func (UnimplementedNodeControllerServiceServer) mustEmbedUnimplementedNodeControllerServiceServer() {}

// UnsafeNodeControllerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NodeControllerServiceServer will
// result in compilation errors.
type UnsafeNodeControllerServiceServer interface {
	mustEmbedUnimplementedNodeControllerServiceServer()
}

func RegisterNodeControllerServiceServer(s grpc.ServiceRegistrar, srv NodeControllerServiceServer) {
	s.RegisterService(&NodeControllerService_ServiceDesc, srv)
}

func _NodeControllerService_CreateLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeControllerServiceServer).CreateLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha1.NodeControllerService/CreateLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeControllerServiceServer).CreateLink(ctx, req.(*CreateLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeControllerService_UpdateLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeControllerServiceServer).UpdateLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha1.NodeControllerService/UpdateLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeControllerServiceServer).UpdateLink(ctx, req.(*UpdateLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeControllerService_DeleteLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeControllerServiceServer).DeleteLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha1.NodeControllerService/DeleteLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeControllerServiceServer).DeleteLink(ctx, req.(*DeleteLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeControllerService_GetLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeControllerServiceServer).GetLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha1.NodeControllerService/GetLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeControllerServiceServer).GetLink(ctx, req.(*GetLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeControllerService_GetAllLinks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllLinksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeControllerServiceServer).GetAllLinks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha1.NodeControllerService/GetAllLinks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeControllerServiceServer).GetAllLinks(ctx, req.(*GetAllLinksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeControllerService_CreateStaticRoute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStaticRouteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeControllerServiceServer).CreateStaticRoute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha1.NodeControllerService/CreateStaticRoute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeControllerServiceServer).CreateStaticRoute(ctx, req.(*CreateStaticRouteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeControllerService_DeleteStaticRoute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteStaticRouteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeControllerServiceServer).DeleteStaticRoute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha1.NodeControllerService/DeleteStaticRoute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeControllerServiceServer).DeleteStaticRoute(ctx, req.(*DeleteStaticRouteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeControllerService_UpdateStaticRoute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStaticRouteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeControllerServiceServer).UpdateStaticRoute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha1.NodeControllerService/UpdateStaticRoute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeControllerServiceServer).UpdateStaticRoute(ctx, req.(*UpdateStaticRouteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeControllerService_GetRoute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRouteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeControllerServiceServer).GetRoute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha1.NodeControllerService/GetRoute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeControllerServiceServer).GetRoute(ctx, req.(*GetRouteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeControllerService_GetAllRoutes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRoutesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeControllerServiceServer).GetAllRoutes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha1.NodeControllerService/GetAllRoutes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeControllerServiceServer).GetAllRoutes(ctx, req.(*GetAllRoutesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeControllerService_GetNeighbor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNeighborRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeControllerServiceServer).GetNeighbor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha1.NodeControllerService/GetNeighbor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeControllerServiceServer).GetNeighbor(ctx, req.(*GetNeighborRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeControllerService_GetAllNeighbors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllNeighborsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeControllerServiceServer).GetAllNeighbors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha1.NodeControllerService/GetAllNeighbors",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeControllerServiceServer).GetAllNeighbors(ctx, req.(*GetAllNeighborsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeControllerService_CreateNeighbor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNeighborRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeControllerServiceServer).CreateNeighbor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha1.NodeControllerService/CreateNeighbor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeControllerServiceServer).CreateNeighbor(ctx, req.(*CreateNeighborRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeControllerService_DeleteNeighbor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNeighborRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeControllerServiceServer).DeleteNeighbor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha1.NodeControllerService/DeleteNeighbor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeControllerServiceServer).DeleteNeighbor(ctx, req.(*DeleteNeighborRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeControllerService_UpdateNeighbor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNeighborRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeControllerServiceServer).UpdateNeighbor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha1.NodeControllerService/UpdateNeighbor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeControllerServiceServer).UpdateNeighbor(ctx, req.(*UpdateNeighborRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeControllerService_DeleteAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeControllerServiceServer).DeleteAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha1.NodeControllerService/DeleteAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeControllerServiceServer).DeleteAddress(ctx, req.(*DeleteAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeControllerService_UpdateAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeControllerServiceServer).UpdateAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha1.NodeControllerService/UpdateAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeControllerServiceServer).UpdateAddress(ctx, req.(*UpdateAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NodeControllerService_ServiceDesc is the grpc.ServiceDesc for NodeControllerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NodeControllerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1alpha1.NodeControllerService",
	HandlerType: (*NodeControllerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateLink",
			Handler:    _NodeControllerService_CreateLink_Handler,
		},
		{
			MethodName: "UpdateLink",
			Handler:    _NodeControllerService_UpdateLink_Handler,
		},
		{
			MethodName: "DeleteLink",
			Handler:    _NodeControllerService_DeleteLink_Handler,
		},
		{
			MethodName: "GetLink",
			Handler:    _NodeControllerService_GetLink_Handler,
		},
		{
			MethodName: "GetAllLinks",
			Handler:    _NodeControllerService_GetAllLinks_Handler,
		},
		{
			MethodName: "CreateStaticRoute",
			Handler:    _NodeControllerService_CreateStaticRoute_Handler,
		},
		{
			MethodName: "DeleteStaticRoute",
			Handler:    _NodeControllerService_DeleteStaticRoute_Handler,
		},
		{
			MethodName: "UpdateStaticRoute",
			Handler:    _NodeControllerService_UpdateStaticRoute_Handler,
		},
		{
			MethodName: "GetRoute",
			Handler:    _NodeControllerService_GetRoute_Handler,
		},
		{
			MethodName: "GetAllRoutes",
			Handler:    _NodeControllerService_GetAllRoutes_Handler,
		},
		{
			MethodName: "GetNeighbor",
			Handler:    _NodeControllerService_GetNeighbor_Handler,
		},
		{
			MethodName: "GetAllNeighbors",
			Handler:    _NodeControllerService_GetAllNeighbors_Handler,
		},
		{
			MethodName: "CreateNeighbor",
			Handler:    _NodeControllerService_CreateNeighbor_Handler,
		},
		{
			MethodName: "DeleteNeighbor",
			Handler:    _NodeControllerService_DeleteNeighbor_Handler,
		},
		{
			MethodName: "UpdateNeighbor",
			Handler:    _NodeControllerService_UpdateNeighbor_Handler,
		},
		{
			MethodName: "DeleteAddress",
			Handler:    _NodeControllerService_DeleteAddress_Handler,
		},
		{
			MethodName: "UpdateAddress",
			Handler:    _NodeControllerService_UpdateAddress_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "node_control.proto",
}

// NodeFirewallControllerServiceClient is the client API for NodeFirewallControllerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NodeFirewallControllerServiceClient interface {
}

type nodeFirewallControllerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeFirewallControllerServiceClient(cc grpc.ClientConnInterface) NodeFirewallControllerServiceClient {
	return &nodeFirewallControllerServiceClient{cc}
}

// NodeFirewallControllerServiceServer is the server API for NodeFirewallControllerService service.
// All implementations must embed UnimplementedNodeFirewallControllerServiceServer
// for forward compatibility
type NodeFirewallControllerServiceServer interface {
	mustEmbedUnimplementedNodeFirewallControllerServiceServer()
}

// UnimplementedNodeFirewallControllerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNodeFirewallControllerServiceServer struct {
}

func (UnimplementedNodeFirewallControllerServiceServer) mustEmbedUnimplementedNodeFirewallControllerServiceServer() {
}

// UnsafeNodeFirewallControllerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NodeFirewallControllerServiceServer will
// result in compilation errors.
type UnsafeNodeFirewallControllerServiceServer interface {
	mustEmbedUnimplementedNodeFirewallControllerServiceServer()
}

func RegisterNodeFirewallControllerServiceServer(s grpc.ServiceRegistrar, srv NodeFirewallControllerServiceServer) {
	s.RegisterService(&NodeFirewallControllerService_ServiceDesc, srv)
}

// NodeFirewallControllerService_ServiceDesc is the grpc.ServiceDesc for NodeFirewallControllerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NodeFirewallControllerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1alpha1.NodeFirewallControllerService",
	HandlerType: (*NodeFirewallControllerServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "node_control.proto",
}