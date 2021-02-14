// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package test

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

// PointClient is the client API for Point service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PointClient interface {
	SortPoints(ctx context.Context, in *Points, opts ...grpc.CallOption) (*Points, error)
}

type pointClient struct {
	cc grpc.ClientConnInterface
}

func NewPointClient(cc grpc.ClientConnInterface) PointClient {
	return &pointClient{cc}
}

func (c *pointClient) SortPoints(ctx context.Context, in *Points, opts ...grpc.CallOption) (*Points, error) {
	out := new(Points)
	err := c.cc.Invoke(ctx, "/decard.Point/SortPoints", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PointServer is the server API for Point service.
// All implementations must embed UnimplementedPointServer
// for forward compatibility
type PointServer interface {
	SortPoints(context.Context, *Points) (*Points, error)
	mustEmbedUnimplementedPointServer()
}

// UnimplementedPointServer must be embedded to have forward compatible implementations.
type UnimplementedPointServer struct {
}

func (UnimplementedPointServer) SortPoints(context.Context, *Points) (*Points, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SortPoints not implemented")
}
func (UnimplementedPointServer) mustEmbedUnimplementedPointServer() {}

// UnsafePointServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PointServer will
// result in compilation errors.
type UnsafePointServer interface {
	mustEmbedUnimplementedPointServer()
}

func RegisterPointServer(s grpc.ServiceRegistrar, srv PointServer) {
	s.RegisterService(&Point_ServiceDesc, srv)
}

func _Point_SortPoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Points)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PointServer).SortPoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/decard.Point/SortPoints",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PointServer).SortPoints(ctx, req.(*Points))
	}
	return interceptor(ctx, in, info, handler)
}

// Point_ServiceDesc is the grpc.ServiceDesc for Point service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Point_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "decard.Point",
	HandlerType: (*PointServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SortPoints",
			Handler:    _Point_SortPoints_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "decard.proto",
}