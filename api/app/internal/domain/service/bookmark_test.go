package service

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kkntzw/bookmark/internal/domain/entity"
	sample_entity "github.com/kkntzw/bookmark/test/data/domain/entity"
	mock_repository "github.com/kkntzw/bookmark/test/mock/domain/repository"
	"github.com/stretchr/testify/assert"
)

func TestNewBookmarkService_service_Bookmark型のインスタンスを返却する(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	// given
	repository := mock_repository.NewMockBookmark(ctrl)
	// when
	object := NewBookmarkService(repository)
	// then
	interfaceObject := (*Bookmark)(nil)
	assert.Implements(t, interfaceObject, object)
	assert.NotNil(t, object)
}

func TestNewBookmarkService_戻り値は初期化済みのフィールドrepositoryを持つ(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	// given
	repository := mock_repository.NewMockBookmark(ctrl)
	abstract := NewBookmarkService(repository)
	// when
	concrete, ok := abstract.(*bookmarkService)
	// then
	assert.True(t, ok)
	expected := repository
	assert.Exactly(t, expected, concrete.repository)
}

func TestExists_ブックマークが存在する場合はtrueを返却する(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	// given
	repository := mock_repository.NewMockBookmark(ctrl)
	repository.EXPECT().FindByID(sample_entity.BookmarkID()).Return(sample_entity.Bookmark(), nil)
	service := NewBookmarkService(repository)
	bookmark := sample_entity.Bookmark()
	// when
	exists, err := service.Exists(bookmark)
	// then
	assert.True(t, exists)
	assert.NoError(t, err)
}

func TestExists_ブックマークが存在しない場合はfalseを返却する(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	// given
	repository := mock_repository.NewMockBookmark(ctrl)
	repository.EXPECT().FindByID(sample_entity.BookmarkID()).Return(nil, nil)
	service := NewBookmarkService(repository)
	bookmark := sample_entity.Bookmark()
	// when
	exists, err := service.Exists(bookmark)
	// then
	assert.False(t, exists)
	assert.NoError(t, err)
}

func TestExists_不正な値を受け取るとエラーを返却する(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	// given
	repository := mock_repository.NewMockBookmark(ctrl)
	service := NewBookmarkService(repository)
	bookmark := (*entity.Bookmark)(nil)
	// when
	exists, err := service.Exists(bookmark)
	// then
	assert.False(t, exists)
	errString := "argument \"bookmark\" is nil"
	assert.EqualError(t, err, errString)
}

func TestExists_リポジトリでエラーが発生した場合はエラーを返却する(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	// given
	repository := mock_repository.NewMockBookmark(ctrl)
	repository.EXPECT().FindByID(sample_entity.BookmarkID()).Return(nil, fmt.Errorf("some error"))
	service := NewBookmarkService(repository)
	bookmark := sample_entity.Bookmark()
	// when
	exists, err := service.Exists(bookmark)
	// then
	assert.False(t, exists)
	errString := "failed at repository.FindByID: some error"
	assert.EqualError(t, err, errString)
}
