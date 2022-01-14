// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: proto/cryptocurrency.proto

package pb

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

// CryptocurrencyElectionClient is the client API for CryptocurrencyElection service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CryptocurrencyElectionClient interface {
	CreateNew(ctx context.Context, in *CreateCryptocurrency, opts ...grpc.CallOption) (*Cryptocurrency, error)
	FindCryptocurrencies(ctx context.Context, in *GetCryptocurrencyParams, opts ...grpc.CallOption) (*CryptocurrencyList, error)
	FindById(ctx context.Context, in *CryptocurrencyId, opts ...grpc.CallOption) (*Cryptocurrency, error)
	DeleteById(ctx context.Context, in *CryptocurrencyId, opts ...grpc.CallOption) (*CryptocurrencyMessage, error)
	UpdateById(ctx context.Context, in *UpdateCryptocurrency, opts ...grpc.CallOption) (*CryptocurrencyMessage, error)
	UpvoteById(ctx context.Context, in *CryptocurrencyId, opts ...grpc.CallOption) (*Cryptocurrency, error)
	DownvoteById(ctx context.Context, in *CryptocurrencyId, opts ...grpc.CallOption) (*Cryptocurrency, error)
}

type cryptocurrencyElectionClient struct {
	cc grpc.ClientConnInterface
}

func NewCryptocurrencyElectionClient(cc grpc.ClientConnInterface) CryptocurrencyElectionClient {
	return &cryptocurrencyElectionClient{cc}
}

func (c *cryptocurrencyElectionClient) CreateNew(ctx context.Context, in *CreateCryptocurrency, opts ...grpc.CallOption) (*Cryptocurrency, error) {
	out := new(Cryptocurrency)
	err := c.cc.Invoke(ctx, "/CryptocurrencyElection/CreateNew", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptocurrencyElectionClient) FindCryptocurrencies(ctx context.Context, in *GetCryptocurrencyParams, opts ...grpc.CallOption) (*CryptocurrencyList, error) {
	out := new(CryptocurrencyList)
	err := c.cc.Invoke(ctx, "/CryptocurrencyElection/FindCryptocurrencies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptocurrencyElectionClient) FindById(ctx context.Context, in *CryptocurrencyId, opts ...grpc.CallOption) (*Cryptocurrency, error) {
	out := new(Cryptocurrency)
	err := c.cc.Invoke(ctx, "/CryptocurrencyElection/FindById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptocurrencyElectionClient) DeleteById(ctx context.Context, in *CryptocurrencyId, opts ...grpc.CallOption) (*CryptocurrencyMessage, error) {
	out := new(CryptocurrencyMessage)
	err := c.cc.Invoke(ctx, "/CryptocurrencyElection/DeleteById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptocurrencyElectionClient) UpdateById(ctx context.Context, in *UpdateCryptocurrency, opts ...grpc.CallOption) (*CryptocurrencyMessage, error) {
	out := new(CryptocurrencyMessage)
	err := c.cc.Invoke(ctx, "/CryptocurrencyElection/UpdateById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptocurrencyElectionClient) UpvoteById(ctx context.Context, in *CryptocurrencyId, opts ...grpc.CallOption) (*Cryptocurrency, error) {
	out := new(Cryptocurrency)
	err := c.cc.Invoke(ctx, "/CryptocurrencyElection/UpvoteById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptocurrencyElectionClient) DownvoteById(ctx context.Context, in *CryptocurrencyId, opts ...grpc.CallOption) (*Cryptocurrency, error) {
	out := new(Cryptocurrency)
	err := c.cc.Invoke(ctx, "/CryptocurrencyElection/DownvoteById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CryptocurrencyElectionServer is the server API for CryptocurrencyElection service.
// All implementations must embed UnimplementedCryptocurrencyElectionServer
// for forward compatibility
type CryptocurrencyElectionServer interface {
	CreateNew(context.Context, *CreateCryptocurrency) (*Cryptocurrency, error)
	FindCryptocurrencies(context.Context, *GetCryptocurrencyParams) (*CryptocurrencyList, error)
	FindById(context.Context, *CryptocurrencyId) (*Cryptocurrency, error)
	DeleteById(context.Context, *CryptocurrencyId) (*CryptocurrencyMessage, error)
	UpdateById(context.Context, *UpdateCryptocurrency) (*CryptocurrencyMessage, error)
	UpvoteById(context.Context, *CryptocurrencyId) (*Cryptocurrency, error)
	DownvoteById(context.Context, *CryptocurrencyId) (*Cryptocurrency, error)
	mustEmbedUnimplementedCryptocurrencyElectionServer()
}

// UnimplementedCryptocurrencyElectionServer must be embedded to have forward compatible implementations.
type UnimplementedCryptocurrencyElectionServer struct {
}

func (UnimplementedCryptocurrencyElectionServer) CreateNew(context.Context, *CreateCryptocurrency) (*Cryptocurrency, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNew not implemented")
}
func (UnimplementedCryptocurrencyElectionServer) FindCryptocurrencies(context.Context, *GetCryptocurrencyParams) (*CryptocurrencyList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindCryptocurrencies not implemented")
}
func (UnimplementedCryptocurrencyElectionServer) FindById(context.Context, *CryptocurrencyId) (*Cryptocurrency, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindById not implemented")
}
func (UnimplementedCryptocurrencyElectionServer) DeleteById(context.Context, *CryptocurrencyId) (*CryptocurrencyMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteById not implemented")
}
func (UnimplementedCryptocurrencyElectionServer) UpdateById(context.Context, *UpdateCryptocurrency) (*CryptocurrencyMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateById not implemented")
}
func (UnimplementedCryptocurrencyElectionServer) UpvoteById(context.Context, *CryptocurrencyId) (*Cryptocurrency, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpvoteById not implemented")
}
func (UnimplementedCryptocurrencyElectionServer) DownvoteById(context.Context, *CryptocurrencyId) (*Cryptocurrency, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownvoteById not implemented")
}
func (UnimplementedCryptocurrencyElectionServer) mustEmbedUnimplementedCryptocurrencyElectionServer() {
}

// UnsafeCryptocurrencyElectionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CryptocurrencyElectionServer will
// result in compilation errors.
type UnsafeCryptocurrencyElectionServer interface {
	mustEmbedUnimplementedCryptocurrencyElectionServer()
}

func RegisterCryptocurrencyElectionServer(s grpc.ServiceRegistrar, srv CryptocurrencyElectionServer) {
	s.RegisterService(&CryptocurrencyElection_ServiceDesc, srv)
}

func _CryptocurrencyElection_CreateNew_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCryptocurrency)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptocurrencyElectionServer).CreateNew(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CryptocurrencyElection/CreateNew",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptocurrencyElectionServer).CreateNew(ctx, req.(*CreateCryptocurrency))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptocurrencyElection_FindCryptocurrencies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCryptocurrencyParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptocurrencyElectionServer).FindCryptocurrencies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CryptocurrencyElection/FindCryptocurrencies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptocurrencyElectionServer).FindCryptocurrencies(ctx, req.(*GetCryptocurrencyParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptocurrencyElection_FindById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CryptocurrencyId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptocurrencyElectionServer).FindById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CryptocurrencyElection/FindById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptocurrencyElectionServer).FindById(ctx, req.(*CryptocurrencyId))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptocurrencyElection_DeleteById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CryptocurrencyId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptocurrencyElectionServer).DeleteById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CryptocurrencyElection/DeleteById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptocurrencyElectionServer).DeleteById(ctx, req.(*CryptocurrencyId))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptocurrencyElection_UpdateById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCryptocurrency)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptocurrencyElectionServer).UpdateById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CryptocurrencyElection/UpdateById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptocurrencyElectionServer).UpdateById(ctx, req.(*UpdateCryptocurrency))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptocurrencyElection_UpvoteById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CryptocurrencyId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptocurrencyElectionServer).UpvoteById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CryptocurrencyElection/UpvoteById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptocurrencyElectionServer).UpvoteById(ctx, req.(*CryptocurrencyId))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptocurrencyElection_DownvoteById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CryptocurrencyId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptocurrencyElectionServer).DownvoteById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CryptocurrencyElection/DownvoteById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptocurrencyElectionServer).DownvoteById(ctx, req.(*CryptocurrencyId))
	}
	return interceptor(ctx, in, info, handler)
}

// CryptocurrencyElection_ServiceDesc is the grpc.ServiceDesc for CryptocurrencyElection service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CryptocurrencyElection_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "CryptocurrencyElection",
	HandlerType: (*CryptocurrencyElectionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNew",
			Handler:    _CryptocurrencyElection_CreateNew_Handler,
		},
		{
			MethodName: "FindCryptocurrencies",
			Handler:    _CryptocurrencyElection_FindCryptocurrencies_Handler,
		},
		{
			MethodName: "FindById",
			Handler:    _CryptocurrencyElection_FindById_Handler,
		},
		{
			MethodName: "DeleteById",
			Handler:    _CryptocurrencyElection_DeleteById_Handler,
		},
		{
			MethodName: "UpdateById",
			Handler:    _CryptocurrencyElection_UpdateById_Handler,
		},
		{
			MethodName: "UpvoteById",
			Handler:    _CryptocurrencyElection_UpvoteById_Handler,
		},
		{
			MethodName: "DownvoteById",
			Handler:    _CryptocurrencyElection_DownvoteById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/cryptocurrency.proto",
}