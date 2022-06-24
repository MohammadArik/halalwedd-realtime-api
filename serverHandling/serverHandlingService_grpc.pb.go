// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package serverHandling

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
	PublishServer(ctx context.Context, in *ServerInfo, opts ...grpc.CallOption) (*ServerInfoResponse, error)
}

type serverHandlingClient struct {
	cc grpc.ClientConnInterface
}

func NewServerHandlingClient(cc grpc.ClientConnInterface) ServerHandlingClient {
	return &serverHandlingClient{cc}
}

func (c *serverHandlingClient) PublishServer(ctx context.Context, in *ServerInfo, opts ...grpc.CallOption) (*ServerInfoResponse, error) {
	out := new(ServerInfoResponse)
	err := c.cc.Invoke(ctx, "/halalwedd_managing_server.ServerHandling/publishServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServerHandlingServer is the server API for ServerHandling service.
// All implementations must embed UnimplementedServerHandlingServer
// for forward compatibility
type ServerHandlingServer interface {
	PublishServer(context.Context, *ServerInfo) (*ServerInfoResponse, error)
	mustEmbedUnimplementedServerHandlingServer()
}

// UnimplementedServerHandlingServer must be embedded to have forward compatible implementations.
type UnimplementedServerHandlingServer struct {
}

func (UnimplementedServerHandlingServer) PublishServer(context.Context, *ServerInfo) (*ServerInfoResponse, error) {
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
	in := new(ServerInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerHandlingServer).PublishServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/halalwedd_managing_server.ServerHandling/publishServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerHandlingServer).PublishServer(ctx, req.(*ServerInfo))
	}
	return interceptor(ctx, in, info, handler)
}

// ServerHandling_ServiceDesc is the grpc.ServiceDesc for ServerHandling service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServerHandling_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "halalwedd_managing_server.ServerHandling",
	HandlerType: (*ServerHandlingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "publishServer",
			Handler:    _ServerHandling_PublishServer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protofiles/serverHandlingService.proto",
}
