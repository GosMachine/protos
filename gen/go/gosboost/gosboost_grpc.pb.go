// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.3
// source: gosboost/gosboost.proto

package gosboostv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	GosBoostDB_AddAccounts_FullMethodName = "/gosboost.GosBoostDB/AddAccounts"
)

// GosBoostDBClient is the client API for GosBoostDB service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GosBoostDBClient interface {
	AddAccounts(ctx context.Context, in *AddAccountsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type gosBoostDBClient struct {
	cc grpc.ClientConnInterface
}

func NewGosBoostDBClient(cc grpc.ClientConnInterface) GosBoostDBClient {
	return &gosBoostDBClient{cc}
}

func (c *gosBoostDBClient) AddAccounts(ctx context.Context, in *AddAccountsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, GosBoostDB_AddAccounts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GosBoostDBServer is the server API for GosBoostDB service.
// All implementations must embed UnimplementedGosBoostDBServer
// for forward compatibility.
type GosBoostDBServer interface {
	AddAccounts(context.Context, *AddAccountsRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedGosBoostDBServer()
}

// UnimplementedGosBoostDBServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGosBoostDBServer struct{}

func (UnimplementedGosBoostDBServer) AddAccounts(context.Context, *AddAccountsRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAccounts not implemented")
}
func (UnimplementedGosBoostDBServer) mustEmbedUnimplementedGosBoostDBServer() {}
func (UnimplementedGosBoostDBServer) testEmbeddedByValue()                    {}

// UnsafeGosBoostDBServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GosBoostDBServer will
// result in compilation errors.
type UnsafeGosBoostDBServer interface {
	mustEmbedUnimplementedGosBoostDBServer()
}

func RegisterGosBoostDBServer(s grpc.ServiceRegistrar, srv GosBoostDBServer) {
	// If the following call pancis, it indicates UnimplementedGosBoostDBServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&GosBoostDB_ServiceDesc, srv)
}

func _GosBoostDB_AddAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddAccountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GosBoostDBServer).AddAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GosBoostDB_AddAccounts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GosBoostDBServer).AddAccounts(ctx, req.(*AddAccountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GosBoostDB_ServiceDesc is the grpc.ServiceDesc for GosBoostDB service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GosBoostDB_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gosboost.GosBoostDB",
	HandlerType: (*GosBoostDBServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddAccounts",
			Handler:    _GosBoostDB_AddAccounts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gosboost/gosboost.proto",
}

const (
	GosBoost_TwitchFollowChannel_FullMethodName = "/gosboost.GosBoost/TwitchFollowChannel"
)

// GosBoostClient is the client API for GosBoost service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GosBoostClient interface {
	TwitchFollowChannel(ctx context.Context, in *TwitchFollowChannelRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type gosBoostClient struct {
	cc grpc.ClientConnInterface
}

func NewGosBoostClient(cc grpc.ClientConnInterface) GosBoostClient {
	return &gosBoostClient{cc}
}

func (c *gosBoostClient) TwitchFollowChannel(ctx context.Context, in *TwitchFollowChannelRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, GosBoost_TwitchFollowChannel_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GosBoostServer is the server API for GosBoost service.
// All implementations must embed UnimplementedGosBoostServer
// for forward compatibility.
type GosBoostServer interface {
	TwitchFollowChannel(context.Context, *TwitchFollowChannelRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedGosBoostServer()
}

// UnimplementedGosBoostServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGosBoostServer struct{}

func (UnimplementedGosBoostServer) TwitchFollowChannel(context.Context, *TwitchFollowChannelRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TwitchFollowChannel not implemented")
}
func (UnimplementedGosBoostServer) mustEmbedUnimplementedGosBoostServer() {}
func (UnimplementedGosBoostServer) testEmbeddedByValue()                  {}

// UnsafeGosBoostServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GosBoostServer will
// result in compilation errors.
type UnsafeGosBoostServer interface {
	mustEmbedUnimplementedGosBoostServer()
}

func RegisterGosBoostServer(s grpc.ServiceRegistrar, srv GosBoostServer) {
	// If the following call pancis, it indicates UnimplementedGosBoostServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&GosBoost_ServiceDesc, srv)
}

func _GosBoost_TwitchFollowChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TwitchFollowChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GosBoostServer).TwitchFollowChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GosBoost_TwitchFollowChannel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GosBoostServer).TwitchFollowChannel(ctx, req.(*TwitchFollowChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GosBoost_ServiceDesc is the grpc.ServiceDesc for GosBoost service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GosBoost_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gosboost.GosBoost",
	HandlerType: (*GosBoostServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TwitchFollowChannel",
			Handler:    _GosBoost_TwitchFollowChannel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gosboost/gosboost.proto",
}