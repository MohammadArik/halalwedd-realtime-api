// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package serverManaging

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

// ServerHandlingClient is the client API for ServerHandling service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServerHandlingClient interface {
	PublishServer(ctx context.Context, in *PublishServerReq, opts ...grpc.CallOption) (*PublishServerRes, error)
}

type serverHandlingClient struct {
	cc grpc.ClientConnInterface
}

func NewServerHandlingClient(cc grpc.ClientConnInterface) ServerHandlingClient {
	return &serverHandlingClient{cc}
}

func (c *serverHandlingClient) PublishServer(ctx context.Context, in *PublishServerReq, opts ...grpc.CallOption) (*PublishServerRes, error) {
	out := new(PublishServerRes)
	err := c.cc.Invoke(ctx, "/halalwedd.ServerHandling/publishServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServerHandlingServer is the server API for ServerHandling service.
// All implementations must embed UnimplementedServerHandlingServer
// for forward compatibility
type ServerHandlingServer interface {
	PublishServer(context.Context, *PublishServerReq) (*PublishServerRes, error)
	mustEmbedUnimplementedServerHandlingServer()
}

// UnimplementedServerHandlingServer must be embedded to have forward compatible implementations.
type UnimplementedServerHandlingServer struct {
}

func (UnimplementedServerHandlingServer) PublishServer(context.Context, *PublishServerReq) (*PublishServerRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishServer not implemented")
}
func (UnimplementedServerHandlingServer) mustEmbedUnimplementedServerHandlingServer() {}

// UnsafeServerHandlingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServerHandlingServer will
// result in compilation errors.
type UnsafeServerHandlingServer interface {
	mustEmbedUnimplementedServerHandlingServer()
}

func RegisterServerHandlingServer(s grpc.ServiceRegistrar, srv ServerHandlingServer) {
	s.RegisterService(&ServerHandling_ServiceDesc, srv)
}

func _ServerHandling_PublishServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishServerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerHandlingServer).PublishServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/halalwedd.ServerHandling/publishServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerHandlingServer).PublishServer(ctx, req.(*PublishServerReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ServerHandling_ServiceDesc is the grpc.ServiceDesc for ServerHandling service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServerHandling_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "halalwedd.ServerHandling",
	HandlerType: (*ServerHandlingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "publishServer",
			Handler:    _ServerHandling_PublishServer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protofiles/serverManaging.proto",
}

// EdgeServerConnectivityClient is the client API for EdgeServerConnectivity service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EdgeServerConnectivityClient interface {
	PingServer(ctx context.Context, in *PingReq, opts ...grpc.CallOption) (*PingRes, error)
	VerifyServer(ctx context.Context, in *VerifyReq, opts ...grpc.CallOption) (*VerifyRes, error)
	DataUpdate(ctx context.Context, in *DataUpdateReq, opts ...grpc.CallOption) (*DataUpdateRes, error)
}

type edgeServerConnectivityClient struct {
	cc grpc.ClientConnInterface
}

func NewEdgeServerConnectivityClient(cc grpc.ClientConnInterface) EdgeServerConnectivityClient {
	return &edgeServerConnectivityClient{cc}
}

func (c *edgeServerConnectivityClient) PingServer(ctx context.Context, in *PingReq, opts ...grpc.CallOption) (*PingRes, error) {
	out := new(PingRes)
	err := c.cc.Invoke(ctx, "/halalwedd.EdgeServerConnectivity/pingServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *edgeServerConnectivityClient) VerifyServer(ctx context.Context, in *VerifyReq, opts ...grpc.CallOption) (*VerifyRes, error) {
	out := new(VerifyRes)
	err := c.cc.Invoke(ctx, "/halalwedd.EdgeServerConnectivity/verifyServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *edgeServerConnectivityClient) DataUpdate(ctx context.Context, in *DataUpdateReq, opts ...grpc.CallOption) (*DataUpdateRes, error) {
	out := new(DataUpdateRes)
	err := c.cc.Invoke(ctx, "/halalwedd.EdgeServerConnectivity/dataUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EdgeServerConnectivityServer is the server API for EdgeServerConnectivity service.
// All implementations must embed UnimplementedEdgeServerConnectivityServer
// for forward compatibility
type EdgeServerConnectivityServer interface {
	PingServer(context.Context, *PingReq) (*PingRes, error)
	VerifyServer(context.Context, *VerifyReq) (*VerifyRes, error)
	DataUpdate(context.Context, *DataUpdateReq) (*DataUpdateRes, error)
	mustEmbedUnimplementedEdgeServerConnectivityServer()
}

// UnimplementedEdgeServerConnectivityServer must be embedded to have forward compatible implementations.
type UnimplementedEdgeServerConnectivityServer struct {
}

func (UnimplementedEdgeServerConnectivityServer) PingServer(context.Context, *PingReq) (*PingRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PingServer not implemented")
}
func (UnimplementedEdgeServerConnectivityServer) VerifyServer(context.Context, *VerifyReq) (*VerifyRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyServer not implemented")
}
func (UnimplementedEdgeServerConnectivityServer) DataUpdate(context.Context, *DataUpdateReq) (*DataUpdateRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DataUpdate not implemented")
}
func (UnimplementedEdgeServerConnectivityServer) mustEmbedUnimplementedEdgeServerConnectivityServer() {
}

// UnsafeEdgeServerConnectivityServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EdgeServerConnectivityServer will
// result in compilation errors.
type UnsafeEdgeServerConnectivityServer interface {
	mustEmbedUnimplementedEdgeServerConnectivityServer()
}

func RegisterEdgeServerConnectivityServer(s grpc.ServiceRegistrar, srv EdgeServerConnectivityServer) {
	s.RegisterService(&EdgeServerConnectivity_ServiceDesc, srv)
}

func _EdgeServerConnectivity_PingServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EdgeServerConnectivityServer).PingServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/halalwedd.EdgeServerConnectivity/pingServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EdgeServerConnectivityServer).PingServer(ctx, req.(*PingReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _EdgeServerConnectivity_VerifyServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EdgeServerConnectivityServer).VerifyServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/halalwedd.EdgeServerConnectivity/verifyServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EdgeServerConnectivityServer).VerifyServer(ctx, req.(*VerifyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _EdgeServerConnectivity_DataUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EdgeServerConnectivityServer).DataUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/halalwedd.EdgeServerConnectivity/dataUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EdgeServerConnectivityServer).DataUpdate(ctx, req.(*DataUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

// EdgeServerConnectivity_ServiceDesc is the grpc.ServiceDesc for EdgeServerConnectivity service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EdgeServerConnectivity_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "halalwedd.EdgeServerConnectivity",
	HandlerType: (*EdgeServerConnectivityServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "pingServer",
			Handler:    _EdgeServerConnectivity_PingServer_Handler,
		},
		{
			MethodName: "verifyServer",
			Handler:    _EdgeServerConnectivity_VerifyServer_Handler,
		},
		{
			MethodName: "dataUpdate",
			Handler:    _EdgeServerConnectivity_DataUpdate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protofiles/serverManaging.proto",
}
