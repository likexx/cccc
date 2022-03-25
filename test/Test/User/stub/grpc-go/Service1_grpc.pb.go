// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package stubs

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

// Service1Client is the client API for Service1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type Service1Client interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type service1Client struct {
	cc grpc.ClientConnInterface
}

func NewService1Client(cc grpc.ClientConnInterface) Service1Client {
	return &service1Client{cc}
}

func (c *service1Client) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/Test.User.Service1/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Service1Server is the server API for Service1 service.
// All implementations must embed UnimplementedService1Server
// for forward compatibility
type Service1Server interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	mustEmbedUnimplementedService1Server()
}

// UnimplementedService1Server must be embedded to have forward compatible implementations.
type UnimplementedService1Server struct {
}

func (UnimplementedService1Server) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedService1Server) mustEmbedUnimplementedService1Server() {}

// UnsafeService1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to Service1Server will
// result in compilation errors.
type UnsafeService1Server interface {
	mustEmbedUnimplementedService1Server()
}

func RegisterService1Server(s grpc.ServiceRegistrar, srv Service1Server) {
	s.RegisterService(&Service1_ServiceDesc, srv)
}

func _Service1_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Service1Server).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Test.User.Service1/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Service1Server).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Service1_ServiceDesc is the grpc.ServiceDesc for Service1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Test.User.Service1",
	HandlerType: (*Service1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Service1_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Service1.proto",
}
