package server

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kkntzw/bookmark/internal/application/command"
	"github.com/kkntzw/bookmark/internal/presentation/pb"
	mock_usecase "github.com/kkntzw/bookmark/test/mock/application/usecase"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestNewBookmarkServer_server_Bookmark型のインスタンスを返却する(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	// given
	usecase := mock_usecase.NewMockBookmark(ctrl)
	// when
	object := NewBookmarkServer(usecase)
	// then
	interfaceObject := (*pb.BookmarkerServer)(nil)
	assert.Implements(t, interfaceObject, object)
	assert.NotNil(t, object)
}

func TestNewBookmarkServer_戻り値は初期化済みのフィールドusecaseを持つ(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	// given
	usecase := mock_usecase.NewMockBookmark(ctrl)
	abstract := NewBookmarkServer(usecase)
	// when
	concrete, ok := abstract.(*bookmarkServer)
	// then
	assert.True(t, ok)
	expectedUsecase := usecase
	assert.Exactly(t, expectedUsecase, concrete.usecase)
}

func TestCreateBookmark_正当な値を受け取るとnilを返却する(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	// given
	usecase := mock_usecase.NewMockBookmark(ctrl)
	cmd := &command.RegisterBookmark{
		Name: "example",
		URI:  "https://example.com",
		Tags: []string{"1", "2", "3"},
	}
	usecase.EXPECT().Register(cmd).Return(nil)
	server := NewBookmarkServer(usecase)
	ctx := context.TODO()
	req := &pb.CreateBookmarkRequest{
		BookmarkName: "example",
		Uri:          "https://example.com",
		Tags: []*pb.Tag{
			{TagName: "1"},
			{TagName: "2"},
			{TagName: "3"},
		},
	}
	// when
	actual, err := server.CreateBookmark(ctx, req)
	// then
	expected := &emptypb.Empty{}
	assert.Exactly(t, expected, actual)
	assert.NoError(t, err)
}

func TestCreateBookmark_不正な値を受け取るとInvalidArgumentを返却する(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	// given
	usecase := mock_usecase.NewMockBookmark(ctrl)
	server := NewBookmarkServer(usecase)
	ctx := context.TODO()
	req := (*pb.CreateBookmarkRequest)(nil)
	// when
	object, actual := server.CreateBookmark(ctx, req)
	// then
	assert.Nil(t, object)
	expected := status.Error(codes.InvalidArgument, "argument \"req\" is nil")
	assert.Exactly(t, expected, actual)
}

func TestCreateBookmark_コマンドが不正な場合はInvalidArgumentを返却する(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	// given
	usecase := mock_usecase.NewMockBookmark(ctrl)
	cmd := &command.RegisterBookmark{
		Name: "",
		URI:  "",
		Tags: []string{""},
	}
	usecase.EXPECT().Register(cmd).Return(&command.InvalidCommandError{})
	server := NewBookmarkServer(usecase)
	ctx := context.TODO()
	req := &pb.CreateBookmarkRequest{BookmarkName: "", Uri: "", Tags: []*pb.Tag{{TagName: ""}}}
	// when
	object, actual := server.CreateBookmark(ctx, req)
	// then
	assert.Nil(t, object)
	expected := status.Error(codes.InvalidArgument, "request is invalid")
	assert.Exactly(t, expected, actual)
}

func TestCreateBookmark_ブックマーク登録中にエラーが発生した場合はInternalを返却する(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	// given
	usecase := mock_usecase.NewMockBookmark(ctrl)
	cmd := &command.RegisterBookmark{
		Name: "example",
		URI:  "https://example.com",
		Tags: []string{"1", "2", "3"},
	}
	usecase.EXPECT().Register(cmd).Return(fmt.Errorf("some error"))
	server := NewBookmarkServer(usecase)
	ctx := context.TODO()
	req := &pb.CreateBookmarkRequest{
		BookmarkName: "example",
		Uri:          "https://example.com",
		Tags: []*pb.Tag{
			{TagName: "1"},
			{TagName: "2"},
			{TagName: "3"},
		},
	}
	// when
	object, actual := server.CreateBookmark(ctx, req)
	// then
	assert.Nil(t, object)
	expected := status.Error(codes.Internal, "server error")
	assert.Exactly(t, expected, actual)
}
