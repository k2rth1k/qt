// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package quick_trade

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

// QuickTradeClient is the client API for QuickTrade service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QuickTradeClient interface {
	HelloWorld(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*Message, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*User, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Refresh(ctx context.Context, in *RefreshRequest, opts ...grpc.CallOption) (*RefreshResponse, error)
	Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*EmptyMessage, error)
}

type quickTradeClient struct {
	cc grpc.ClientConnInterface
}

func NewQuickTradeClient(cc grpc.ClientConnInterface) QuickTradeClient {
	return &quickTradeClient{cc}
}

func (c *quickTradeClient) HelloWorld(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/quick_trade.QuickTrade/HelloWorld", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quickTradeClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/quick_trade.QuickTrade/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quickTradeClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/quick_trade.QuickTrade/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quickTradeClient) Refresh(ctx context.Context, in *RefreshRequest, opts ...grpc.CallOption) (*RefreshResponse, error) {
	out := new(RefreshResponse)
	err := c.cc.Invoke(ctx, "/quick_trade.QuickTrade/Refresh", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quickTradeClient) Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*EmptyMessage, error) {
	out := new(EmptyMessage)
	err := c.cc.Invoke(ctx, "/quick_trade.QuickTrade/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QuickTradeServer is the server API for QuickTrade service.
// All implementations should embed UnimplementedQuickTradeServer
// for forward compatibility
type QuickTradeServer interface {
	HelloWorld(context.Context, *EmptyMessage) (*Message, error)
	CreateUser(context.Context, *CreateUserRequest) (*User, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Refresh(context.Context, *RefreshRequest) (*RefreshResponse, error)
	Logout(context.Context, *LogoutRequest) (*EmptyMessage, error)
}

// UnimplementedQuickTradeServer should be embedded to have forward compatible implementations.
type UnimplementedQuickTradeServer struct {
}

func (UnimplementedQuickTradeServer) HelloWorld(context.Context, *EmptyMessage) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HelloWorld not implemented")
}
func (UnimplementedQuickTradeServer) CreateUser(context.Context, *CreateUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedQuickTradeServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedQuickTradeServer) Refresh(context.Context, *RefreshRequest) (*RefreshResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Refresh not implemented")
}
func (UnimplementedQuickTradeServer) Logout(context.Context, *LogoutRequest) (*EmptyMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}

// UnsafeQuickTradeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QuickTradeServer will
// result in compilation errors.
type UnsafeQuickTradeServer interface {
	mustEmbedUnimplementedQuickTradeServer()
}

func RegisterQuickTradeServer(s grpc.ServiceRegistrar, srv QuickTradeServer) {
	s.RegisterService(&QuickTrade_ServiceDesc, srv)
}

func _QuickTrade_HelloWorld_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuickTradeServer).HelloWorld(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/quick_trade.QuickTrade/HelloWorld",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuickTradeServer).HelloWorld(ctx, req.(*EmptyMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuickTrade_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuickTradeServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/quick_trade.QuickTrade/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuickTradeServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuickTrade_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuickTradeServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/quick_trade.QuickTrade/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuickTradeServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuickTrade_Refresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuickTradeServer).Refresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/quick_trade.QuickTrade/Refresh",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuickTradeServer).Refresh(ctx, req.(*RefreshRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuickTrade_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuickTradeServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/quick_trade.QuickTrade/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuickTradeServer).Logout(ctx, req.(*LogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// QuickTrade_ServiceDesc is the grpc.ServiceDesc for QuickTrade service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var QuickTrade_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "quick_trade.QuickTrade",
	HandlerType: (*QuickTradeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HelloWorld",
			Handler:    _QuickTrade_HelloWorld_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _QuickTrade_CreateUser_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _QuickTrade_Login_Handler,
		},
		{
			MethodName: "Refresh",
			Handler:    _QuickTrade_Refresh_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _QuickTrade_Logout_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "quick-trade.proto",
}
