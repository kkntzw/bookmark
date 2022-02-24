package server

import (
	"context"
	"errors"

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
// ブックマークの作成に成功した場合は OK を返却する。
// nilを指定した場合は INVALID_ARGUMENT を返却する。
// 不正なリクエストを指定した場合は INVALID_ARGUMENT を返却する。
// ブックマークの作成に失敗した場合は INTERNAL を返却する。
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
	err := s.usecase.Register(cmd)
	var icerr *command.InvalidCommandError
	if errors.As(err, &icerr) {
		return nil, status.Error(codes.InvalidArgument, "request is invalid")
	}
	if err != nil {
		return nil, status.Error(codes.Internal, "server error")
	}
	return &emptypb.Empty{}, nil
}

// ブックマークを一覧取得する。
//
// ブックマークの一覧取得に成功した場合は OK を返却する。
// nilを指定した場合は INVALID_ARGUMENT を返却する。
// ブックマークの一覧取得に失敗した場合は INTERNAL を返却する。
// ストリームの送信に失敗した場合は INTERNAL を返却する。
func (s *bookmarkServer) ListBookmarks(req *emptypb.Empty, stream pb.Bookmarker_ListBookmarksServer) error {
	if req == nil {
		return status.Error(codes.InvalidArgument, "argument \"req\" is nil")
	}
	bookmarks, err := s.usecase.List()
	if err != nil {
		return status.Error(codes.Internal, "server error")
	}
	for _, bookmark := range bookmarks {
		tags := make([]*pb.Tag, len(bookmark.Tags))
		for i, tag := range bookmark.Tags {
			tags[i] = &pb.Tag{TagName: tag}
		}
		res := &pb.Bookmark{
			BookmarkId:   bookmark.ID,
			BookmarkName: bookmark.Name,
			Uri:          bookmark.URI,
			Tags:         tags,
		}
		if err := stream.Send(res); err != nil {
			return status.Error(codes.Internal, "response failed")
		}
	}
	return nil
}

// ブックマークを更新する。
//
// ブックマークの更新に成功した場合は OK を返却する。
// nilを指定した場合は INVALID_ARGUMENT を返却する。
// 不正なリクエストを指定した場合は INVALID_ARGUMENT を返却する。
// ブックマークの更新に失敗した場合は INTERNAL を返却する。
func (s *bookmarkServer) UpdateBookmark(ctx context.Context, req *pb.UpdateBookmarkRequest) (*emptypb.Empty, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "argument \"req\" is nil")
	}
	id := req.BookmarkId
	name := req.BookmarkName
	uri := req.Uri
	cmd := &command.UpdateBookmark{ID: id, Name: name, URI: uri}
	err := s.usecase.Update(cmd)
	var icerr *command.InvalidCommandError
	if errors.As(err, &icerr) {
		return nil, status.Error(codes.InvalidArgument, "request is invalid")
	}
	if err != nil {
		return nil, status.Error(codes.Internal, "server error")
	}
	return &emptypb.Empty{}, nil
}
