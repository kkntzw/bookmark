// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: bookmark.proto

package pb

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

// BookmarkerClient is the client API for Bookmarker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookmarkerClient interface {
	// ブックマークを作成する。
	//
	// 作成に成功した場合は OK を返却する。
	// 無効な引数を指定した場合は INVALID_ARGUMENT を返却する。
	// サーバエラーが発生した場合は INTERNAL を返却する。
	CreateBookmark(ctx context.Context, in *CreateBookmarkRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type bookmarkerClient struct {
	cc grpc.ClientConnInterface
}

func NewBookmarkerClient(cc grpc.ClientConnInterface) BookmarkerClient {
	return &bookmarkerClient{cc}
}

func (c *bookmarkerClient) CreateBookmark(ctx context.Context, in *CreateBookmarkRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/bookmark.Bookmarker/CreateBookmark", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookmarkerServer is the server API for Bookmarker service.
// All implementations must embed UnimplementedBookmarkerServer
// for forward compatibility
type BookmarkerServer interface {
	// ブックマークを作成する。
	//
	// 作成に成功した場合は OK を返却する。
	// 無効な引数を指定した場合は INVALID_ARGUMENT を返却する。
	// サーバエラーが発生した場合は INTERNAL を返却する。
	CreateBookmark(context.Context, *CreateBookmarkRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedBookmarkerServer()
}

// UnimplementedBookmarkerServer must be embedded to have forward compatible implementations.
type UnimplementedBookmarkerServer struct {
}

func (UnimplementedBookmarkerServer) CreateBookmark(context.Context, *CreateBookmarkRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBookmark not implemented")
}
func (UnimplementedBookmarkerServer) mustEmbedUnimplementedBookmarkerServer() {}

// UnsafeBookmarkerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookmarkerServer will
// result in compilation errors.
type UnsafeBookmarkerServer interface {
	mustEmbedUnimplementedBookmarkerServer()
}

func RegisterBookmarkerServer(s grpc.ServiceRegistrar, srv BookmarkerServer) {
	s.RegisterService(&Bookmarker_ServiceDesc, srv)
}

func _Bookmarker_CreateBookmark_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBookmarkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookmarkerServer).CreateBookmark(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bookmark.Bookmarker/CreateBookmark",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookmarkerServer).CreateBookmark(ctx, req.(*CreateBookmarkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Bookmarker_ServiceDesc is the grpc.ServiceDesc for Bookmarker service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Bookmarker_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bookmark.Bookmarker",
	HandlerType: (*BookmarkerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBookmark",
			Handler:    _Bookmarker_CreateBookmark_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bookmark.proto",
}
