package server

import (
	"context"

	"github.com/kkntzw/bookmark/internal/application/command"
	"github.com/kkntzw/bookmark/internal/application/usecase"
	"github.com/kkntzw/bookmark/internal/presentation/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// ブックマークに関するgRPCサーバの具象型
type bookmarkServer struct {
	usecase usecase.Bookmark // ユースケース
	pb.UnimplementedBookmarkerServer
}

// ブックマークに関するgRPCサーバを生成する。
func NewBookmarkServer(usecase usecase.Bookmark) pb.BookmarkerServer {
	return &bookmarkServer{
		usecase: usecase,
	}
}

// ブックマークを作成する。
//
// 作成に成功した場合は OK を返却する。
// 無効な引数を指定した場合は INVALID_ARGUMENT を返却する。
// サーバエラーが発生した場合は INTERNAL を返却する。z
func (s *bookmarkServer) CreateBookmark(ctx context.Context, req *pb.CreateBookmarkRequest) (*emptypb.Empty, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "argument \"req\" is nil")
	}
	name := req.BookmarkName
	uri := req.Uri
	tags := make([]string, len(req.Tags))
	for i, tag := range req.Tags {
		tags[i] = tag.TagName
	}
	cmd := &command.RegisterBookmark{Name: name, URI: uri, Tags: tags}
	if err := cmd.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, "request is invalid")
	}
	if err := s.usecase.Register(cmd); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &emptypb.Empty{}, nil
}
