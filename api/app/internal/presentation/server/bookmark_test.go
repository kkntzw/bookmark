package server

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kkntzw/bookmark/internal/application/command"
	"github.com/kkntzw/bookmark/internal/application/dto"
	"github.com/kkntzw/bookmark/internal/presentation/pb"
	mock_usecase "github.com/kkntzw/bookmark/test/mock/application/usecase"
	mock_pb "github.com/kkntzw/bookmark/test/mock/presentation/pb"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func ToCreateBookmarkRequest(t *testing.T, name, uri string, tagNames ...string) *pb.CreateBookmarkRequest {
	t.Helper()
	tags := make([]*pb.Tag, len(tagNames))
	for i, tagName := range tagNames {
		tags[i] = &pb.Tag{TagName: tagName}
	}
	req := &pb.CreateBookmarkRequest{
		BookmarkName: name,
		Uri:          uri,
		Tags:         tags,
	}
	return req
}

func ToUpdateBookmarkRequest(t *testing.T, id, name, uri string) *pb.UpdateBookmarkRequest {
	t.Helper()
	req := &pb.UpdateBookmarkRequest{
		BookmarkId:   id,
		BookmarkName: name,
		Uri:          uri,
	}
	return req
}

func ToBookmark(t *testing.T, id, name, uri string, tagNames ...string) *pb.Bookmark {
	t.Helper()
	tags := make([]*pb.Tag, len(tagNames))
	for i, tagName := range tagNames {
		tags[i] = &pb.Tag{TagName: tagName}
	}
	bookmark := &pb.Bookmark{
		BookmarkId:   id,
		BookmarkName: name,
		Uri:          uri,
		Tags:         tags,
	}
	return bookmark
}

func TestNewBookmarkServer(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	{
		t.Run("implementing bookmark server", func(t *testing.T) {
			t.Parallel()
			// given
			usecase := mock_usecase.NewMockBookmark(ctrl)
			// when
			object := NewBookmarkServer(usecase)
			// then
			assert.NotNil(t, object)
			interfaceObject := (*pb.BookmarkerServer)(nil)
			assert.Implements(t, interfaceObject, object)
		})
	}
	{
		t.Run("fields", func(t *testing.T) {
			t.Parallel()
			// given
			usecase := mock_usecase.NewMockBookmark(ctrl)
			abstractServer := NewBookmarkServer(usecase)
			// when
			concreteServer, ok := abstractServer.(*bookmarkServer)
			actualUsecase := concreteServer.usecase
			// then
			assert.True(t, ok)
			expectedUsecase := usecase
			assert.Exactly(t, expectedUsecase, actualUsecase)
		})
	}
}

func TestBookmark_CreateBookmark(t *testing.T) {
	t.Parallel()
	cmd := &command.RegisterBookmark{Name: "Example A", URI: "https://foo.example.com", Tags: []string{"1-A", "1-B", "1-C"}}
	req := ToCreateBookmarkRequest(t, "Example A", "https://foo.example.com", "1-A", "1-B", "1-C")
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cases := map[string]struct {
		prepare          func(*mock_usecase.MockBookmark)
		req              *pb.CreateBookmarkRequest
		expectedResponse *emptypb.Empty
		expectedErr      error
	}{
		"non-nil request": {
			func(usecase *mock_usecase.MockBookmark) {
				usecase.EXPECT().Register(cmd).Return(nil)
			},
			req,
			&emptypb.Empty{},
			nil,
		},
		"nil request": {
			func(usecase *mock_usecase.MockBookmark) {},
			nil,
			nil,
			status.Error(codes.InvalidArgument, "argument \"req\" is nil"),
		},
		"invalid request": {
			func(usecase *mock_usecase.MockBookmark) {
				usecase.
					EXPECT().
					Register(&command.RegisterBookmark{Name: "Example A", URI: "", Tags: []string{}}).
					Return(&command.InvalidCommandError{Args: map[string]error{"URI": errors.New("some error")}})
			},
			ToCreateBookmarkRequest(t, "Example A", ""),
			nil,
			status.Error(codes.InvalidArgument, "request is invalid"),
		},
		"failed at usecase.Register": {
			func(usecase *mock_usecase.MockBookmark) {
				usecase.EXPECT().Register(cmd).Return(errors.New("some error"))
			},
			req,
			nil,
			status.Error(codes.Internal, "server error"),
		},
	}
	for name, tc := range cases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			usecase := mock_usecase.NewMockBookmark(ctrl)
			tc.prepare(usecase)
			// given
			server := NewBookmarkServer(usecase)
			ctx := context.TODO()
			// when
			actualResponse, actualErr := server.CreateBookmark(ctx, tc.req)
			// then
			assert.Exactly(t, tc.expectedResponse, actualResponse)
			assert.Exactly(t, tc.expectedErr, actualErr)
		})
	}
}

func TestBookmark_ListBookmarks(t *testing.T) {
	t.Parallel()
	bookmarks := []dto.Bookmark{
		{ID: "1", Name: "Example A", URI: "https://foo.example.com", Tags: []string{}},
		{ID: "2", Name: "Example B", URI: "https://bar.example.com", Tags: []string{"2-A"}},
		{ID: "3", Name: "Example C", URI: "https://baz.example.com", Tags: []string{"3-A", "3-B"}},
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cases := map[string]struct {
		prepare     func(*mock_usecase.MockBookmark, *mock_pb.MockBookmarker_ListBookmarksServer)
		req         *emptypb.Empty
		expectedErr error
	}{
		"non-nil request": {
			func(usecase *mock_usecase.MockBookmark, stream *mock_pb.MockBookmarker_ListBookmarksServer) {
				usecase.EXPECT().List().Return(bookmarks, nil)
				stream.EXPECT().Send(ToBookmark(t, "1", "Example A", "https://foo.example.com")).Return(nil)
				stream.EXPECT().Send(ToBookmark(t, "2", "Example B", "https://bar.example.com", "2-A")).Return(nil)
				stream.EXPECT().Send(ToBookmark(t, "3", "Example C", "https://baz.example.com", "3-A", "3-B")).Return(nil)
			},
			&emptypb.Empty{},
			nil,
		},
		"nil request": {
			func(usecase *mock_usecase.MockBookmark, stream *mock_pb.MockBookmarker_ListBookmarksServer) {},
			nil,
			status.Error(codes.InvalidArgument, "argument \"req\" is nil"),
		},
		"failed at usecase.List": {
			func(usecase *mock_usecase.MockBookmark, stream *mock_pb.MockBookmarker_ListBookmarksServer) {
				usecase.EXPECT().List().Return(nil, errors.New("some error"))
			},
			&emptypb.Empty{},
			status.Error(codes.Internal, "server error"),
		},
		"failed at stream.Send": {
			func(usecase *mock_usecase.MockBookmark, stream *mock_pb.MockBookmarker_ListBookmarksServer) {
				usecase.EXPECT().List().Return(bookmarks, nil)
				stream.EXPECT().Send(ToBookmark(t, "1", "Example A", "https://foo.example.com")).Return(nil)
				stream.EXPECT().Send(ToBookmark(t, "2", "Example B", "https://bar.example.com", "2-A")).Return(errors.New("some error"))
			},
			&emptypb.Empty{},
			status.Error(codes.Internal, "response failed"),
		},
	}
	for name, tc := range cases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			usecase := mock_usecase.NewMockBookmark(ctrl)
			stream := mock_pb.NewMockBookmarker_ListBookmarksServer(ctrl)
			tc.prepare(usecase, stream)
			// given
			server := NewBookmarkServer(usecase)
			// when
			actualErr := server.ListBookmarks(tc.req, stream)
			// then
			assert.Exactly(t, tc.expectedErr, actualErr)
		})
	}
}

func TestBookmark_UpdateBookmark(t *testing.T) {
	t.Parallel()
	cmd := &command.UpdateBookmark{
		ID:   "1",
		Name: "Example Bar",
		URI:  "https://foo.example.com/bar",
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cases := map[string]struct {
		prepare          func(*mock_usecase.MockBookmark)
		req              *pb.UpdateBookmarkRequest
		expectedResponse *emptypb.Empty
		expectedErr      error
	}{
		"non-nil request": {
			func(usecase *mock_usecase.MockBookmark) {
				usecase.EXPECT().Update(cmd).Return(nil)
			},
			ToUpdateBookmarkRequest(t, "1", "Example Bar", "https://foo.example.com/bar"),
			&emptypb.Empty{},
			nil,
		},
		"nil request": {
			func(usecase *mock_usecase.MockBookmark) {},
			nil,
			nil,
			status.Error(codes.InvalidArgument, "argument \"req\" is nil"),
		},
		"invalid request": {
			func(usecase *mock_usecase.MockBookmark) {
				usecase.EXPECT().Update(&command.UpdateBookmark{ID: "", Name: "", URI: ""}).Return(&command.InvalidCommandError{})
			},
			ToUpdateBookmarkRequest(t, "", "", ""),
			nil,
			status.Error(codes.InvalidArgument, "request is invalid"),
		},
		"failed at usecase.Update": {
			func(usecase *mock_usecase.MockBookmark) {
				usecase.EXPECT().Update(cmd).Return(errors.New("some error"))
			},
			ToUpdateBookmarkRequest(t, "1", "Example Bar", "https://foo.example.com/bar"),
			nil,
			status.Error(codes.Internal, "server error"),
		},
	}
	for name, tc := range cases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			usecase := mock_usecase.NewMockBookmark(ctrl)
			tc.prepare(usecase)
			// given
			server := NewBookmarkServer(usecase)
			ctx := context.TODO()
			// when
			actualResponse, actualErr := server.UpdateBookmark(ctx, tc.req)
			// then
			assert.Exactly(t, tc.expectedResponse, actualResponse)
			assert.Exactly(t, tc.expectedErr, actualErr)
		})
	}
}
