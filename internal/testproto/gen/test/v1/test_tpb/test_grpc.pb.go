// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: test/v1/topic/test.proto

package test_tpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TestTopicClient is the client API for TestTopic service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TestTopicClient interface {
	Test(ctx context.Context, in *TestMessage, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type testTopicClient struct {
	cc grpc.ClientConnInterface
}

func NewTestTopicClient(cc grpc.ClientConnInterface) TestTopicClient {
	return &testTopicClient{cc}
}

func (c *testTopicClient) Test(ctx context.Context, in *TestMessage, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/test.v1.topic.TestTopic/Test", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestTopicServer is the server API for TestTopic service.
// All implementations must embed UnimplementedTestTopicServer
// for forward compatibility
type TestTopicServer interface {
	Test(context.Context, *TestMessage) (*emptypb.Empty, error)
	mustEmbedUnimplementedTestTopicServer()
}

// UnimplementedTestTopicServer must be embedded to have forward compatible implementations.
type UnimplementedTestTopicServer struct {
}

func (UnimplementedTestTopicServer) Test(context.Context, *TestMessage) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Test not implemented")
}
func (UnimplementedTestTopicServer) mustEmbedUnimplementedTestTopicServer() {}

// UnsafeTestTopicServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TestTopicServer will
// result in compilation errors.
type UnsafeTestTopicServer interface {
	mustEmbedUnimplementedTestTopicServer()
}

func RegisterTestTopicServer(s grpc.ServiceRegistrar, srv TestTopicServer) {
	s.RegisterService(&TestTopic_ServiceDesc, srv)
}

func _TestTopic_Test_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestTopicServer).Test(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.v1.topic.TestTopic/Test",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestTopicServer).Test(ctx, req.(*TestMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// TestTopic_ServiceDesc is the grpc.ServiceDesc for TestTopic service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TestTopic_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "test.v1.topic.TestTopic",
	HandlerType: (*TestTopicServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Test",
			Handler:    _TestTopic_Test_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test/v1/topic/test.proto",
}

// GreetingRequestTopicClient is the client API for GreetingRequestTopic service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreetingRequestTopicClient interface {
	Greeting(ctx context.Context, in *GreetingMessage, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type greetingRequestTopicClient struct {
	cc grpc.ClientConnInterface
}

func NewGreetingRequestTopicClient(cc grpc.ClientConnInterface) GreetingRequestTopicClient {
	return &greetingRequestTopicClient{cc}
}

func (c *greetingRequestTopicClient) Greeting(ctx context.Context, in *GreetingMessage, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/test.v1.topic.GreetingRequestTopic/Greeting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreetingRequestTopicServer is the server API for GreetingRequestTopic service.
// All implementations must embed UnimplementedGreetingRequestTopicServer
// for forward compatibility
type GreetingRequestTopicServer interface {
	Greeting(context.Context, *GreetingMessage) (*emptypb.Empty, error)
	mustEmbedUnimplementedGreetingRequestTopicServer()
}

// UnimplementedGreetingRequestTopicServer must be embedded to have forward compatible implementations.
type UnimplementedGreetingRequestTopicServer struct {
}

func (UnimplementedGreetingRequestTopicServer) Greeting(context.Context, *GreetingMessage) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Greeting not implemented")
}
func (UnimplementedGreetingRequestTopicServer) mustEmbedUnimplementedGreetingRequestTopicServer() {}

// UnsafeGreetingRequestTopicServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreetingRequestTopicServer will
// result in compilation errors.
type UnsafeGreetingRequestTopicServer interface {
	mustEmbedUnimplementedGreetingRequestTopicServer()
}

func RegisterGreetingRequestTopicServer(s grpc.ServiceRegistrar, srv GreetingRequestTopicServer) {
	s.RegisterService(&GreetingRequestTopic_ServiceDesc, srv)
}

func _GreetingRequestTopic_Greeting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreetingMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetingRequestTopicServer).Greeting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.v1.topic.GreetingRequestTopic/Greeting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetingRequestTopicServer).Greeting(ctx, req.(*GreetingMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// GreetingRequestTopic_ServiceDesc is the grpc.ServiceDesc for GreetingRequestTopic service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GreetingRequestTopic_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "test.v1.topic.GreetingRequestTopic",
	HandlerType: (*GreetingRequestTopicServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Greeting",
			Handler:    _GreetingRequestTopic_Greeting_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test/v1/topic/test.proto",
}

// GreetingResponseTopicClient is the client API for GreetingResponseTopic service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreetingResponseTopicClient interface {
	Response(ctx context.Context, in *ResponseMessage, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type greetingResponseTopicClient struct {
	cc grpc.ClientConnInterface
}

func NewGreetingResponseTopicClient(cc grpc.ClientConnInterface) GreetingResponseTopicClient {
	return &greetingResponseTopicClient{cc}
}

func (c *greetingResponseTopicClient) Response(ctx context.Context, in *ResponseMessage, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/test.v1.topic.GreetingResponseTopic/Response", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreetingResponseTopicServer is the server API for GreetingResponseTopic service.
// All implementations must embed UnimplementedGreetingResponseTopicServer
// for forward compatibility
type GreetingResponseTopicServer interface {
	Response(context.Context, *ResponseMessage) (*emptypb.Empty, error)
	mustEmbedUnimplementedGreetingResponseTopicServer()
}

// UnimplementedGreetingResponseTopicServer must be embedded to have forward compatible implementations.
type UnimplementedGreetingResponseTopicServer struct {
}

func (UnimplementedGreetingResponseTopicServer) Response(context.Context, *ResponseMessage) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Response not implemented")
}
func (UnimplementedGreetingResponseTopicServer) mustEmbedUnimplementedGreetingResponseTopicServer() {}

// UnsafeGreetingResponseTopicServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreetingResponseTopicServer will
// result in compilation errors.
type UnsafeGreetingResponseTopicServer interface {
	mustEmbedUnimplementedGreetingResponseTopicServer()
}

func RegisterGreetingResponseTopicServer(s grpc.ServiceRegistrar, srv GreetingResponseTopicServer) {
	s.RegisterService(&GreetingResponseTopic_ServiceDesc, srv)
}

func _GreetingResponseTopic_Response_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResponseMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetingResponseTopicServer).Response(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.v1.topic.GreetingResponseTopic/Response",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetingResponseTopicServer).Response(ctx, req.(*ResponseMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// GreetingResponseTopic_ServiceDesc is the grpc.ServiceDesc for GreetingResponseTopic service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GreetingResponseTopic_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "test.v1.topic.GreetingResponseTopic",
	HandlerType: (*GreetingResponseTopicServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Response",
			Handler:    _GreetingResponseTopic_Response_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test/v1/topic/test.proto",
}