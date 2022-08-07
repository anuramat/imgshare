// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: api/api.proto

package api

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

// BotDBClient is the client API for BotDB service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BotDBClient interface {
	CreateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
	ReadUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
	UpdateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
	DeleteUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
	CreateImage(ctx context.Context, in *ImageAuthRequest, opts ...grpc.CallOption) (*Image, error)
	ReadImage(ctx context.Context, in *Image, opts ...grpc.CallOption) (*Image, error)
	GetRandomImage(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Image, error)
	SetDescriptionImage(ctx context.Context, in *ImageAuthRequest, opts ...grpc.CallOption) (*Image, error)
	UpvoteImage(ctx context.Context, in *ImageAuthRequest, opts ...grpc.CallOption) (*Image, error)
	DownvoteImage(ctx context.Context, in *ImageAuthRequest, opts ...grpc.CallOption) (*Image, error)
	DeleteImage(ctx context.Context, in *ImageAuthRequest, opts ...grpc.CallOption) (*Empty, error)
	// HW-2 requirement
	GetAllImages(ctx context.Context, in *Page, opts ...grpc.CallOption) (*Images, error)
}

type botDBClient struct {
	cc grpc.ClientConnInterface
}

func NewBotDBClient(cc grpc.ClientConnInterface) BotDBClient {
	return &botDBClient{cc}
}

func (c *botDBClient) CreateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/api.BotDB/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botDBClient) ReadUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/api.BotDB/ReadUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botDBClient) UpdateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/api.BotDB/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botDBClient) DeleteUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/api.BotDB/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botDBClient) CreateImage(ctx context.Context, in *ImageAuthRequest, opts ...grpc.CallOption) (*Image, error) {
	out := new(Image)
	err := c.cc.Invoke(ctx, "/api.BotDB/CreateImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botDBClient) ReadImage(ctx context.Context, in *Image, opts ...grpc.CallOption) (*Image, error) {
	out := new(Image)
	err := c.cc.Invoke(ctx, "/api.BotDB/ReadImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botDBClient) GetRandomImage(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Image, error) {
	out := new(Image)
	err := c.cc.Invoke(ctx, "/api.BotDB/GetRandomImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botDBClient) SetDescriptionImage(ctx context.Context, in *ImageAuthRequest, opts ...grpc.CallOption) (*Image, error) {
	out := new(Image)
	err := c.cc.Invoke(ctx, "/api.BotDB/SetDescriptionImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botDBClient) UpvoteImage(ctx context.Context, in *ImageAuthRequest, opts ...grpc.CallOption) (*Image, error) {
	out := new(Image)
	err := c.cc.Invoke(ctx, "/api.BotDB/UpvoteImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botDBClient) DownvoteImage(ctx context.Context, in *ImageAuthRequest, opts ...grpc.CallOption) (*Image, error) {
	out := new(Image)
	err := c.cc.Invoke(ctx, "/api.BotDB/DownvoteImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botDBClient) DeleteImage(ctx context.Context, in *ImageAuthRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/api.BotDB/DeleteImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botDBClient) GetAllImages(ctx context.Context, in *Page, opts ...grpc.CallOption) (*Images, error) {
	out := new(Images)
	err := c.cc.Invoke(ctx, "/api.BotDB/GetAllImages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BotDBServer is the server API for BotDB service.
// All implementations must embed UnimplementedBotDBServer
// for forward compatibility
type BotDBServer interface {
	CreateUser(context.Context, *User) (*User, error)
	ReadUser(context.Context, *User) (*User, error)
	UpdateUser(context.Context, *User) (*User, error)
	DeleteUser(context.Context, *User) (*User, error)
	CreateImage(context.Context, *ImageAuthRequest) (*Image, error)
	ReadImage(context.Context, *Image) (*Image, error)
	GetRandomImage(context.Context, *Empty) (*Image, error)
	SetDescriptionImage(context.Context, *ImageAuthRequest) (*Image, error)
	UpvoteImage(context.Context, *ImageAuthRequest) (*Image, error)
	DownvoteImage(context.Context, *ImageAuthRequest) (*Image, error)
	DeleteImage(context.Context, *ImageAuthRequest) (*Empty, error)
	// HW-2 requirement
	GetAllImages(context.Context, *Page) (*Images, error)
	mustEmbedUnimplementedBotDBServer()
}

// UnimplementedBotDBServer must be embedded to have forward compatible implementations.
type UnimplementedBotDBServer struct {
}

func (UnimplementedBotDBServer) CreateUser(context.Context, *User) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedBotDBServer) ReadUser(context.Context, *User) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadUser not implemented")
}
func (UnimplementedBotDBServer) UpdateUser(context.Context, *User) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedBotDBServer) DeleteUser(context.Context, *User) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedBotDBServer) CreateImage(context.Context, *ImageAuthRequest) (*Image, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateImage not implemented")
}
func (UnimplementedBotDBServer) ReadImage(context.Context, *Image) (*Image, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadImage not implemented")
}
func (UnimplementedBotDBServer) GetRandomImage(context.Context, *Empty) (*Image, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRandomImage not implemented")
}
func (UnimplementedBotDBServer) SetDescriptionImage(context.Context, *ImageAuthRequest) (*Image, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetDescriptionImage not implemented")
}
func (UnimplementedBotDBServer) UpvoteImage(context.Context, *ImageAuthRequest) (*Image, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpvoteImage not implemented")
}
func (UnimplementedBotDBServer) DownvoteImage(context.Context, *ImageAuthRequest) (*Image, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownvoteImage not implemented")
}
func (UnimplementedBotDBServer) DeleteImage(context.Context, *ImageAuthRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteImage not implemented")
}
func (UnimplementedBotDBServer) GetAllImages(context.Context, *Page) (*Images, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllImages not implemented")
}
func (UnimplementedBotDBServer) mustEmbedUnimplementedBotDBServer() {}

// UnsafeBotDBServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BotDBServer will
// result in compilation errors.
type UnsafeBotDBServer interface {
	mustEmbedUnimplementedBotDBServer()
}

func RegisterBotDBServer(s grpc.ServiceRegistrar, srv BotDBServer) {
	s.RegisterService(&BotDB_ServiceDesc, srv)
}

func _BotDB_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotDBServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.BotDB/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotDBServer).CreateUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _BotDB_ReadUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotDBServer).ReadUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.BotDB/ReadUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotDBServer).ReadUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _BotDB_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotDBServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.BotDB/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotDBServer).UpdateUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _BotDB_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotDBServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.BotDB/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotDBServer).DeleteUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _BotDB_CreateImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImageAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotDBServer).CreateImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.BotDB/CreateImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotDBServer).CreateImage(ctx, req.(*ImageAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BotDB_ReadImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Image)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotDBServer).ReadImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.BotDB/ReadImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotDBServer).ReadImage(ctx, req.(*Image))
	}
	return interceptor(ctx, in, info, handler)
}

func _BotDB_GetRandomImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotDBServer).GetRandomImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.BotDB/GetRandomImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotDBServer).GetRandomImage(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _BotDB_SetDescriptionImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImageAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotDBServer).SetDescriptionImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.BotDB/SetDescriptionImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotDBServer).SetDescriptionImage(ctx, req.(*ImageAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BotDB_UpvoteImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImageAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotDBServer).UpvoteImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.BotDB/UpvoteImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotDBServer).UpvoteImage(ctx, req.(*ImageAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BotDB_DownvoteImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImageAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotDBServer).DownvoteImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.BotDB/DownvoteImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotDBServer).DownvoteImage(ctx, req.(*ImageAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BotDB_DeleteImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImageAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotDBServer).DeleteImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.BotDB/DeleteImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotDBServer).DeleteImage(ctx, req.(*ImageAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BotDB_GetAllImages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Page)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotDBServer).GetAllImages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.BotDB/GetAllImages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotDBServer).GetAllImages(ctx, req.(*Page))
	}
	return interceptor(ctx, in, info, handler)
}

// BotDB_ServiceDesc is the grpc.ServiceDesc for BotDB service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BotDB_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.BotDB",
	HandlerType: (*BotDBServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _BotDB_CreateUser_Handler,
		},
		{
			MethodName: "ReadUser",
			Handler:    _BotDB_ReadUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _BotDB_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _BotDB_DeleteUser_Handler,
		},
		{
			MethodName: "CreateImage",
			Handler:    _BotDB_CreateImage_Handler,
		},
		{
			MethodName: "ReadImage",
			Handler:    _BotDB_ReadImage_Handler,
		},
		{
			MethodName: "GetRandomImage",
			Handler:    _BotDB_GetRandomImage_Handler,
		},
		{
			MethodName: "SetDescriptionImage",
			Handler:    _BotDB_SetDescriptionImage_Handler,
		},
		{
			MethodName: "UpvoteImage",
			Handler:    _BotDB_UpvoteImage_Handler,
		},
		{
			MethodName: "DownvoteImage",
			Handler:    _BotDB_DownvoteImage_Handler,
		},
		{
			MethodName: "DeleteImage",
			Handler:    _BotDB_DeleteImage_Handler,
		},
		{
			MethodName: "GetAllImages",
			Handler:    _BotDB_GetAllImages_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/api.proto",
}
