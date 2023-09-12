// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.3
// source: proto/lib/lib.proto

package lib

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

// LibClient is the client API for Lib service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LibClient interface {
	GetAuthors(ctx context.Context, in *BookRequest, opts ...grpc.CallOption) (*AuthorsResponse, error)
	GetBooks(ctx context.Context, in *AuthorRequest, opts ...grpc.CallOption) (*BooksResponse, error)
}

type libClient struct {
	cc grpc.ClientConnInterface
}

func NewLibClient(cc grpc.ClientConnInterface) LibClient {
	return &libClient{cc}
}

func (c *libClient) GetAuthors(ctx context.Context, in *BookRequest, opts ...grpc.CallOption) (*AuthorsResponse, error) {
	out := new(AuthorsResponse)
	err := c.cc.Invoke(ctx, "/Lib/GetAuthors", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libClient) GetBooks(ctx context.Context, in *AuthorRequest, opts ...grpc.CallOption) (*BooksResponse, error) {
	out := new(BooksResponse)
	err := c.cc.Invoke(ctx, "/Lib/GetBooks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LibServer is the server API for Lib service.
// All implementations must embed UnimplementedLibServer
// for forward compatibility
type LibServer interface {
	GetAuthors(context.Context, *BookRequest) (*AuthorsResponse, error)
	GetBooks(context.Context, *AuthorRequest) (*BooksResponse, error)
	mustEmbedUnimplementedLibServer()
}

// UnimplementedLibServer must be embedded to have forward compatible implementations.
type UnimplementedLibServer struct {
}

func (UnimplementedLibServer) GetAuthors(context.Context, *BookRequest) (*AuthorsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthors not implemented")
}
func (UnimplementedLibServer) GetBooks(context.Context, *AuthorRequest) (*BooksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBooks not implemented")
}
func (UnimplementedLibServer) mustEmbedUnimplementedLibServer() {}

// UnsafeLibServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LibServer will
// result in compilation errors.
type UnsafeLibServer interface {
	mustEmbedUnimplementedLibServer()
}

func RegisterLibServer(s grpc.ServiceRegistrar, srv LibServer) {
	s.RegisterService(&Lib_ServiceDesc, srv)
}

func _Lib_GetAuthors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibServer).GetAuthors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Lib/GetAuthors",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibServer).GetAuthors(ctx, req.(*BookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Lib_GetBooks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibServer).GetBooks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Lib/GetBooks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibServer).GetBooks(ctx, req.(*AuthorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Lib_ServiceDesc is the grpc.ServiceDesc for Lib service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Lib_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Lib",
	HandlerType: (*LibServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAuthors",
			Handler:    _Lib_GetAuthors_Handler,
		},
		{
			MethodName: "GetBooks",
			Handler:    _Lib_GetBooks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/lib/lib.proto",
}