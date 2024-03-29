// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.17.3
// source: mp/v1/mini_app.proto

package v1

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

const (
	MiniAppService_LoginByCode_FullMethodName    = "/mp.v1.MiniAppService/LoginByCode"
	MiniAppService_GetPhoneNumber_FullMethodName = "/mp.v1.MiniAppService/GetPhoneNumber"
	MiniAppService_ListMiniApp_FullMethodName    = "/mp.v1.MiniAppService/ListMiniApp"
	MiniAppService_CreateMiniApp_FullMethodName  = "/mp.v1.MiniAppService/CreateMiniApp"
	MiniAppService_GetMiniApp_FullMethodName     = "/mp.v1.MiniAppService/GetMiniApp"
	MiniAppService_UpdateMiniApp_FullMethodName  = "/mp.v1.MiniAppService/UpdateMiniApp"
	MiniAppService_DeleteMiniApp_FullMethodName  = "/mp.v1.MiniAppService/DeleteMiniApp"
)

// MiniAppServiceClient is the client API for MiniAppService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiniAppServiceClient interface {
	// 通过code获取openid并登录
	LoginByCode(ctx context.Context, in *CodeLoginReq, opts ...grpc.CallOption) (*LoginReply, error)
	// 通过code获取用户手机号
	GetPhoneNumber(ctx context.Context, in *GetPhoneNumberReq, opts ...grpc.CallOption) (*GetPhoneNumberReply, error)
	// RPC
	ListMiniApp(ctx context.Context, in *ListAppReq, opts ...grpc.CallOption) (*ListAppReply, error)
	CreateMiniApp(ctx context.Context, in *CreateAppReq, opts ...grpc.CallOption) (*MiniApp, error)
	GetMiniApp(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*MiniApp, error)
	UpdateMiniApp(ctx context.Context, in *UpdateAppReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteMiniApp(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type miniAppServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMiniAppServiceClient(cc grpc.ClientConnInterface) MiniAppServiceClient {
	return &miniAppServiceClient{cc}
}

func (c *miniAppServiceClient) LoginByCode(ctx context.Context, in *CodeLoginReq, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := c.cc.Invoke(ctx, MiniAppService_LoginByCode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *miniAppServiceClient) GetPhoneNumber(ctx context.Context, in *GetPhoneNumberReq, opts ...grpc.CallOption) (*GetPhoneNumberReply, error) {
	out := new(GetPhoneNumberReply)
	err := c.cc.Invoke(ctx, MiniAppService_GetPhoneNumber_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *miniAppServiceClient) ListMiniApp(ctx context.Context, in *ListAppReq, opts ...grpc.CallOption) (*ListAppReply, error) {
	out := new(ListAppReply)
	err := c.cc.Invoke(ctx, MiniAppService_ListMiniApp_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *miniAppServiceClient) CreateMiniApp(ctx context.Context, in *CreateAppReq, opts ...grpc.CallOption) (*MiniApp, error) {
	out := new(MiniApp)
	err := c.cc.Invoke(ctx, MiniAppService_CreateMiniApp_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *miniAppServiceClient) GetMiniApp(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*MiniApp, error) {
	out := new(MiniApp)
	err := c.cc.Invoke(ctx, MiniAppService_GetMiniApp_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *miniAppServiceClient) UpdateMiniApp(ctx context.Context, in *UpdateAppReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, MiniAppService_UpdateMiniApp_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *miniAppServiceClient) DeleteMiniApp(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, MiniAppService_DeleteMiniApp_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiniAppServiceServer is the server API for MiniAppService service.
// All implementations must embed UnimplementedMiniAppServiceServer
// for forward compatibility
type MiniAppServiceServer interface {
	// 通过code获取openid并登录
	LoginByCode(context.Context, *CodeLoginReq) (*LoginReply, error)
	// 通过code获取用户手机号
	GetPhoneNumber(context.Context, *GetPhoneNumberReq) (*GetPhoneNumberReply, error)
	// RPC
	ListMiniApp(context.Context, *ListAppReq) (*ListAppReply, error)
	CreateMiniApp(context.Context, *CreateAppReq) (*MiniApp, error)
	GetMiniApp(context.Context, *IdReq) (*MiniApp, error)
	UpdateMiniApp(context.Context, *UpdateAppReq) (*emptypb.Empty, error)
	DeleteMiniApp(context.Context, *IdReq) (*emptypb.Empty, error)
	mustEmbedUnimplementedMiniAppServiceServer()
}

// UnimplementedMiniAppServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMiniAppServiceServer struct {
}

func (UnimplementedMiniAppServiceServer) LoginByCode(context.Context, *CodeLoginReq) (*LoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginByCode not implemented")
}
func (UnimplementedMiniAppServiceServer) GetPhoneNumber(context.Context, *GetPhoneNumberReq) (*GetPhoneNumberReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPhoneNumber not implemented")
}
func (UnimplementedMiniAppServiceServer) ListMiniApp(context.Context, *ListAppReq) (*ListAppReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMiniApp not implemented")
}
func (UnimplementedMiniAppServiceServer) CreateMiniApp(context.Context, *CreateAppReq) (*MiniApp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMiniApp not implemented")
}
func (UnimplementedMiniAppServiceServer) GetMiniApp(context.Context, *IdReq) (*MiniApp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMiniApp not implemented")
}
func (UnimplementedMiniAppServiceServer) UpdateMiniApp(context.Context, *UpdateAppReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMiniApp not implemented")
}
func (UnimplementedMiniAppServiceServer) DeleteMiniApp(context.Context, *IdReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMiniApp not implemented")
}
func (UnimplementedMiniAppServiceServer) mustEmbedUnimplementedMiniAppServiceServer() {}

// UnsafeMiniAppServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MiniAppServiceServer will
// result in compilation errors.
type UnsafeMiniAppServiceServer interface {
	mustEmbedUnimplementedMiniAppServiceServer()
}

func RegisterMiniAppServiceServer(s grpc.ServiceRegistrar, srv MiniAppServiceServer) {
	s.RegisterService(&MiniAppService_ServiceDesc, srv)
}

func _MiniAppService_LoginByCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CodeLoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiniAppServiceServer).LoginByCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MiniAppService_LoginByCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiniAppServiceServer).LoginByCode(ctx, req.(*CodeLoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MiniAppService_GetPhoneNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPhoneNumberReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiniAppServiceServer).GetPhoneNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MiniAppService_GetPhoneNumber_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiniAppServiceServer).GetPhoneNumber(ctx, req.(*GetPhoneNumberReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MiniAppService_ListMiniApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAppReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiniAppServiceServer).ListMiniApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MiniAppService_ListMiniApp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiniAppServiceServer).ListMiniApp(ctx, req.(*ListAppReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MiniAppService_CreateMiniApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAppReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiniAppServiceServer).CreateMiniApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MiniAppService_CreateMiniApp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiniAppServiceServer).CreateMiniApp(ctx, req.(*CreateAppReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MiniAppService_GetMiniApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiniAppServiceServer).GetMiniApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MiniAppService_GetMiniApp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiniAppServiceServer).GetMiniApp(ctx, req.(*IdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MiniAppService_UpdateMiniApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAppReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiniAppServiceServer).UpdateMiniApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MiniAppService_UpdateMiniApp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiniAppServiceServer).UpdateMiniApp(ctx, req.(*UpdateAppReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MiniAppService_DeleteMiniApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiniAppServiceServer).DeleteMiniApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MiniAppService_DeleteMiniApp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiniAppServiceServer).DeleteMiniApp(ctx, req.(*IdReq))
	}
	return interceptor(ctx, in, info, handler)
}

// MiniAppService_ServiceDesc is the grpc.ServiceDesc for MiniAppService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MiniAppService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mp.v1.MiniAppService",
	HandlerType: (*MiniAppServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LoginByCode",
			Handler:    _MiniAppService_LoginByCode_Handler,
		},
		{
			MethodName: "GetPhoneNumber",
			Handler:    _MiniAppService_GetPhoneNumber_Handler,
		},
		{
			MethodName: "ListMiniApp",
			Handler:    _MiniAppService_ListMiniApp_Handler,
		},
		{
			MethodName: "CreateMiniApp",
			Handler:    _MiniAppService_CreateMiniApp_Handler,
		},
		{
			MethodName: "GetMiniApp",
			Handler:    _MiniAppService_GetMiniApp_Handler,
		},
		{
			MethodName: "UpdateMiniApp",
			Handler:    _MiniAppService_UpdateMiniApp_Handler,
		},
		{
			MethodName: "DeleteMiniApp",
			Handler:    _MiniAppService_DeleteMiniApp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mp/v1/mini_app.proto",
}
