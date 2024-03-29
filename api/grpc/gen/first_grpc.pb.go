// Code generated by protoc-gen-go-app. DO NOT EDIT.
// versions:
// - protoc-gen-go-app v1.2.0
// - protoc             v4.25.3
// source: first.proto

package timergrpc

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the app package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// APITimerClient is the client API for APITimer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type APITimerClient interface {
	Timer(ctx context.Context, in *TimerRequest, opts ...grpc.CallOption) (*TimerResponse, error)
	CallbackConfig(ctx context.Context, in *CallbackConfigRequest, opts ...grpc.CallOption) (*CallbackConfigResponse, error)
}

type aPITimerClient struct {
	cc grpc.ClientConnInterface
}

func NewAPITimerClient(cc grpc.ClientConnInterface) APITimerClient {
	return &aPITimerClient{cc}
}

func (c *aPITimerClient) Timer(ctx context.Context, in *TimerRequest, opts ...grpc.CallOption) (*TimerResponse, error) {
	out := new(TimerResponse)
	err := c.cc.Invoke(ctx, "/api.APITimer/Timer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPITimerClient) CallbackConfig(ctx context.Context, in *CallbackConfigRequest, opts ...grpc.CallOption) (*CallbackConfigResponse, error) {
	out := new(CallbackConfigResponse)
	err := c.cc.Invoke(ctx, "/api.APITimer/CallbackConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// APITimerServer is the server API for APITimer service.
// All implementations must embed UnimplementedAPITimerServer
// for forward compatibility
type APITimerServer interface {
	Timer(context.Context, *TimerRequest) (*TimerResponse, error)
	CallbackConfig(context.Context, *CallbackConfigRequest) (*CallbackConfigResponse, error)
	mustEmbedUnimplementedAPITimerServer()
}

// UnimplementedAPITimerServer must be embedded to have forward compatible implementations.
type UnimplementedAPITimerServer struct {
}

func (UnimplementedAPITimerServer) Timer(context.Context, *TimerRequest) (*TimerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Timer not implemented")
}
func (UnimplementedAPITimerServer) CallbackConfig(context.Context, *CallbackConfigRequest) (*CallbackConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CallbackConfig not implemented")
}
func (UnimplementedAPITimerServer) mustEmbedUnimplementedAPITimerServer() {}

// UnsafeAPITimerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to APITimerServer will
// result in compilation errors.
type UnsafeAPITimerServer interface {
	mustEmbedUnimplementedAPITimerServer()
}

func RegisterAPITimerServer(s grpc.ServiceRegistrar, srv APITimerServer) {
	s.RegisterService(&APITimer_ServiceDesc, srv)
}

func _APITimer_Timer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TimerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APITimerServer).Timer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.APITimer/Timer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APITimerServer).Timer(ctx, req.(*TimerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _APITimer_CallbackConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CallbackConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APITimerServer).CallbackConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.APITimer/CallbackConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APITimerServer).CallbackConfig(ctx, req.(*CallbackConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// APITimer_ServiceDesc is the grpc.ServiceDesc for APITimer service.
// It's only intended for direct use with app.RegisterService,
// and not to be introspected or modified (even as a copy)
var APITimer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.APITimer",
	HandlerType: (*APITimerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Timer",
			Handler:    _APITimer_Timer_Handler,
		},
		{
			MethodName: "CallbackConfig",
			Handler:    _APITimer_CallbackConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "first.proto",
}
