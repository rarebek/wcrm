// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: protos/user_service/user_service.proto

package __

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

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	CreateOwner(ctx context.Context, in *Owner, opts ...grpc.CallOption) (*GetOwnerRequest, error)
	GetOwner(ctx context.Context, in *GetOwnerRequest, opts ...grpc.CallOption) (*Owner, error)
	UpdateOwner(ctx context.Context, in *Owner, opts ...grpc.CallOption) (*Owner, error)
	DeleteOwner(ctx context.Context, in *GetOwnerRequest, opts ...grpc.CallOption) (*DeletedOwner, error)
	ListOwner(ctx context.Context, in *GetAllOwnerRequest, opts ...grpc.CallOption) (*GetAllOwnerResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) CreateOwner(ctx context.Context, in *Owner, opts ...grpc.CallOption) (*GetOwnerRequest, error) {
	out := new(GetOwnerRequest)
	err := c.cc.Invoke(ctx, "/UserService/CreateOwner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetOwner(ctx context.Context, in *GetOwnerRequest, opts ...grpc.CallOption) (*Owner, error) {
	out := new(Owner)
	err := c.cc.Invoke(ctx, "/UserService/GetOwner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateOwner(ctx context.Context, in *Owner, opts ...grpc.CallOption) (*Owner, error) {
	out := new(Owner)
	err := c.cc.Invoke(ctx, "/UserService/UpdateOwner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteOwner(ctx context.Context, in *GetOwnerRequest, opts ...grpc.CallOption) (*DeletedOwner, error) {
	out := new(DeletedOwner)
	err := c.cc.Invoke(ctx, "/UserService/DeleteOwner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ListOwner(ctx context.Context, in *GetAllOwnerRequest, opts ...grpc.CallOption) (*GetAllOwnerResponse, error) {
	out := new(GetAllOwnerResponse)
	err := c.cc.Invoke(ctx, "/UserService/ListOwner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	CreateOwner(context.Context, *Owner) (*GetOwnerRequest, error)
	GetOwner(context.Context, *GetOwnerRequest) (*Owner, error)
	UpdateOwner(context.Context, *Owner) (*Owner, error)
	DeleteOwner(context.Context, *GetOwnerRequest) (*DeletedOwner, error)
	ListOwner(context.Context, *GetAllOwnerRequest) (*GetAllOwnerResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) CreateOwner(context.Context, *Owner) (*GetOwnerRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOwner not implemented")
}
func (UnimplementedUserServiceServer) GetOwner(context.Context, *GetOwnerRequest) (*Owner, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOwner not implemented")
}
func (UnimplementedUserServiceServer) UpdateOwner(context.Context, *Owner) (*Owner, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOwner not implemented")
}
func (UnimplementedUserServiceServer) DeleteOwner(context.Context, *GetOwnerRequest) (*DeletedOwner, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOwner not implemented")
}
func (UnimplementedUserServiceServer) ListOwner(context.Context, *GetAllOwnerRequest) (*GetAllOwnerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOwner not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_CreateOwner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Owner)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateOwner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserService/CreateOwner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateOwner(ctx, req.(*Owner))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetOwner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOwnerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetOwner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserService/GetOwner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetOwner(ctx, req.(*GetOwnerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateOwner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Owner)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateOwner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserService/UpdateOwner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateOwner(ctx, req.(*Owner))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteOwner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOwnerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteOwner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserService/DeleteOwner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteOwner(ctx, req.(*GetOwnerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ListOwner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllOwnerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ListOwner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserService/ListOwner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ListOwner(ctx, req.(*GetAllOwnerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOwner",
			Handler:    _UserService_CreateOwner_Handler,
		},
		{
			MethodName: "GetOwner",
			Handler:    _UserService_GetOwner_Handler,
		},
		{
			MethodName: "UpdateOwner",
			Handler:    _UserService_UpdateOwner_Handler,
		},
		{
			MethodName: "DeleteOwner",
			Handler:    _UserService_DeleteOwner_Handler,
		},
		{
			MethodName: "ListOwner",
			Handler:    _UserService_ListOwner_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/user_service/user_service.proto",
}
