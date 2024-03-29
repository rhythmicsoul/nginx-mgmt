// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package controller

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

// AddServiceClient is the client API for AddService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AddServiceClient interface {
	NewAgentToken(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AgentToken, error)
	BiT(ctx context.Context, opts ...grpc.CallOption) (AddService_BiTClient, error)
}

type addServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAddServiceClient(cc grpc.ClientConnInterface) AddServiceClient {
	return &addServiceClient{cc}
}

func (c *addServiceClient) NewAgentToken(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AgentToken, error) {
	out := new(AgentToken)
	err := c.cc.Invoke(ctx, "/controller.AddService/NewAgentToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addServiceClient) BiT(ctx context.Context, opts ...grpc.CallOption) (AddService_BiTClient, error) {
	stream, err := c.cc.NewStream(ctx, &AddService_ServiceDesc.Streams[0], "/controller.AddService/BiT", opts...)
	if err != nil {
		return nil, err
	}
	x := &addServiceBiTClient{stream}
	return x, nil
}

type AddService_BiTClient interface {
	Send(*Test) error
	Recv() (*Test, error)
	grpc.ClientStream
}

type addServiceBiTClient struct {
	grpc.ClientStream
}

func (x *addServiceBiTClient) Send(m *Test) error {
	return x.ClientStream.SendMsg(m)
}

func (x *addServiceBiTClient) Recv() (*Test, error) {
	m := new(Test)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AddServiceServer is the server API for AddService service.
// All implementations must embed UnimplementedAddServiceServer
// for forward compatibility
type AddServiceServer interface {
	NewAgentToken(context.Context, *Empty) (*AgentToken, error)
	BiT(AddService_BiTServer) error
	mustEmbedUnimplementedAddServiceServer()
}

// UnimplementedAddServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAddServiceServer struct {
}

func (UnimplementedAddServiceServer) NewAgentToken(context.Context, *Empty) (*AgentToken, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewAgentToken not implemented")
}
func (UnimplementedAddServiceServer) BiT(AddService_BiTServer) error {
	return status.Errorf(codes.Unimplemented, "method BiT not implemented")
}
func (UnimplementedAddServiceServer) mustEmbedUnimplementedAddServiceServer() {}

// UnsafeAddServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AddServiceServer will
// result in compilation errors.
type UnsafeAddServiceServer interface {
	mustEmbedUnimplementedAddServiceServer()
}

func RegisterAddServiceServer(s grpc.ServiceRegistrar, srv AddServiceServer) {
	s.RegisterService(&AddService_ServiceDesc, srv)
}

func _AddService_NewAgentToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddServiceServer).NewAgentToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/controller.AddService/NewAgentToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddServiceServer).NewAgentToken(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddService_BiT_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AddServiceServer).BiT(&addServiceBiTServer{stream})
}

type AddService_BiTServer interface {
	Send(*Test) error
	Recv() (*Test, error)
	grpc.ServerStream
}

type addServiceBiTServer struct {
	grpc.ServerStream
}

func (x *addServiceBiTServer) Send(m *Test) error {
	return x.ServerStream.SendMsg(m)
}

func (x *addServiceBiTServer) Recv() (*Test, error) {
	m := new(Test)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AddService_ServiceDesc is the grpc.ServiceDesc for AddService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AddService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "controller.AddService",
	HandlerType: (*AddServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewAgentToken",
			Handler:    _AddService_NewAgentToken_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "BiT",
			Handler:       _AddService_BiT_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/controller/controller.proto",
}
