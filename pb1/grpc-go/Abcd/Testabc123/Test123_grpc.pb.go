// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package Testabc123

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

// Test123Client is the client API for Test123 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type Test123Client interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type test123Client struct {
	cc grpc.ClientConnInterface
}

func NewTest123Client(cc grpc.ClientConnInterface) Test123Client {
	return &test123Client{cc}
}

func (c *test123Client) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/Abcd.Testabc123.Test123/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Test123Server is the server API for Test123 service.
// All implementations must embed UnimplementedTest123Server
// for forward compatibility
type Test123Server interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	mustEmbedUnimplementedTest123Server()
}

// UnimplementedTest123Server must be embedded to have forward compatible implementations.
type UnimplementedTest123Server struct {
}

func (UnimplementedTest123Server) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedTest123Server) mustEmbedUnimplementedTest123Server() {}

// UnsafeTest123Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to Test123Server will
// result in compilation errors.
type UnsafeTest123Server interface {
	mustEmbedUnimplementedTest123Server()
}

func RegisterTest123Server(s grpc.ServiceRegistrar, srv Test123Server) {
	s.RegisterService(&Test123_ServiceDesc, srv)
}

func _Test123_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Test123Server).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Abcd.Testabc123.Test123/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Test123Server).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Test123_ServiceDesc is the grpc.ServiceDesc for Test123 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Test123_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Abcd.Testabc123.Test123",
	HandlerType: (*Test123Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Test123_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Test123.proto",
}
