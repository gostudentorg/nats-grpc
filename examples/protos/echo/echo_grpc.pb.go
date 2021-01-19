// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package echo

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

// EchoClient is the client API for Echo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EchoClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	Echo(ctx context.Context, opts ...grpc.CallOption) (Echo_EchoClient, error)
}

type echoClient struct {
	cc grpc.ClientConnInterface
}

func NewEchoClient(cc grpc.ClientConnInterface) EchoClient {
	return &echoClient{cc}
}

func (c *echoClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/echo.Echo/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *echoClient) Echo(ctx context.Context, opts ...grpc.CallOption) (Echo_EchoClient, error) {
	stream, err := c.cc.NewStream(ctx, &Echo_ServiceDesc.Streams[0], "/echo.Echo/Echo", opts...)
	if err != nil {
		return nil, err
	}
	x := &echoEchoClient{stream}
	return x, nil
}

type Echo_EchoClient interface {
	Send(*EchoRequest) error
	Recv() (*EchoReply, error)
	grpc.ClientStream
}

type echoEchoClient struct {
	grpc.ClientStream
}

func (x *echoEchoClient) Send(m *EchoRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *echoEchoClient) Recv() (*EchoReply, error) {
	m := new(EchoReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EchoServer is the server API for Echo service.
// All implementations must embed UnimplementedEchoServer
// for forward compatibility
type EchoServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	Echo(Echo_EchoServer) error
	mustEmbedUnimplementedEchoServer()
}

// UnimplementedEchoServer must be embedded to have forward compatible implementations.
type UnimplementedEchoServer struct {
}

func (UnimplementedEchoServer) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedEchoServer) Echo(Echo_EchoServer) error {
	return status.Errorf(codes.Unimplemented, "method Echo not implemented")
}
func (UnimplementedEchoServer) mustEmbedUnimplementedEchoServer() {}

// UnsafeEchoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EchoServer will
// result in compilation errors.
type UnsafeEchoServer interface {
	mustEmbedUnimplementedEchoServer()
}

func RegisterEchoServer(s grpc.ServiceRegistrar, srv EchoServer) {
	s.RegisterService(&Echo_ServiceDesc, srv)
}

func _Echo_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/echo.Echo/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Echo_Echo_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EchoServer).Echo(&echoEchoServer{stream})
}

type Echo_EchoServer interface {
	Send(*EchoReply) error
	Recv() (*EchoRequest, error)
	grpc.ServerStream
}

type echoEchoServer struct {
	grpc.ServerStream
}

func (x *echoEchoServer) Send(m *EchoReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *echoEchoServer) Recv() (*EchoRequest, error) {
	m := new(EchoRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Echo_ServiceDesc is the grpc.ServiceDesc for Echo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Echo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "echo.Echo",
	HandlerType: (*EchoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Echo_SayHello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Echo",
			Handler:       _Echo_Echo_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "echo.proto",
}