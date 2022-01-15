package usecase

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kkntzw/bookmark/internal/application/command"
	sample_entity "github.com/kkntzw/bookmark/test/data/domain/entity"
	mock_repository "github.com/kkntzw/bookmark/test/mock/domain/repository"
	mock_service "github.com/kkntzw/bookmark/test/mock/domain/service"
	"github.com/stretchr/testify/assert"
)

func TestNewBookmarkUsecase_usecase_Bookmark型のインスタンスを返却する(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	// given
	repository := mock_repository.NewMockBookmark(ctrl)
	service := mock_service.NewMockBookmark(ctrl)
	// when
	object := NewBookmarkUsecase(repository, service)
	// then
	interfaceObject := (*Bookmark)(nil)
	assert.Implements(t, interfaceObject, object)
	assert.NotNil(t, object)
}

func TestNewBookmarkUsecase_戻り値は初期化済みのフィールドrepositoryとserviceを持つ(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	// given
	repository := mock_repository.NewMockBookmark(ctrl)
	service := mock_service.NewMockBookmark(ctrl)
	abstract := NewBookmarkUsecase(repository, service)
	// when
	concrete, ok := abstract.(*bookmarkUsecase)
	// then
	assert.True(t, ok)
	expectedRepository := repository
	assert.Exactly(t, expectedRepository, concrete.repository)
	expectedService := service
	assert.Exactly(t, expectedService, concrete.service)
}

func TestRegister_正当な値を受け取るとnilを返却する(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	// given
	repository := mock_repository.NewMockBookmark(ctrl)
	repository.EXPECT().NextID().Return(sample_entity.BookmarkID())
	repository.EXPECT().Save(sample_entity.Bookmark()).Return(nil)
	service := mock_service.NewMockBookmark(ctrl)
	service.EXPECT().Exists(sample_entity.Bookmark()).Return(false, nil)
	usecase := NewBookmarkUsecase(repository, service)
	cmd := &command.RegisterBookmark{
		Name: "example",
		URI:  "https://example.com",
		Tags: []string{"1", "2", "3"},
	}
	// when
	err := usecase.Register(cmd)
	// then
	assert.NoError(t, err)
}

func TestRegister_不正な値を受け取るとエラーを返却する(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	params := []struct {
		name      string
		uri       string
		tags      []string
		errString string
	}{
		{name: "", uri: "https://example.com", tags: []string{"1"}, errString: "command \"Name\" is invalid"},
		{name: "example", uri: "", tags: []string{"1"}, errString: "command \"URI\" is invalid"},
		{name: "example", uri: "https://example.com", tags: []string{""}, errString: "command \"Tags\" is invalid"},
	}
	for _, p := range params {
		// given
		repository := mock_repository.NewMockBookmark(ctrl)
		repository.EXPECT().NextID().Return(sample_entity.BookmarkID())
		service := mock_service.NewMockBookmark(ctrl)
		usecase := NewBookmarkUsecase(repository, service)
		cmd := &command.RegisterBookmark{
			Name: p.name,
			URI:  p.uri,
			Tags: p.tags,
		}
		// when
		err := usecase.Register(cmd)
		// then
		assert.EqualError(t, err, p.errString)
	}
}

func TestRegister_ブックマークが重複して存在する場合はエラーを返却する(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	// given
	repository := mock_repository.NewMockBookmark(ctrl)
	repository.EXPECT().NextID().Return(sample_entity.BookmarkID())
	service := mock_service.NewMockBookmark(ctrl)
	service.EXPECT().Exists(sample_entity.Bookmark()).Return(true, nil)
	usecase := NewBookmarkUsecase(repository, service)
	cmd := &command.RegisterBookmark{
		Name: "example",
		URI:  "https://example.com",
		Tags: []string{"1", "2", "3"},
	}
	// when
	err := usecase.Register(cmd)
	// then
	errString := "bookmark already exists"
	assert.EqualError(t, err, errString)
}

func TestRegister_ブックマーク重複確認中にエラーが発生した場合はエラーを返却する(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	// given
	repository := mock_repository.NewMockBookmark(ctrl)
	repository.EXPECT().NextID().Return(sample_entity.BookmarkID())
	service := mock_service.NewMockBookmark(ctrl)
	service.EXPECT().Exists(sample_entity.Bookmark()).Return(false, errors.New("some error"))
	usecase := NewBookmarkUsecase(repository, service)
	cmd := &command.RegisterBookmark{
		Name: "example",
		URI:  "https://example.com",
		Tags: []string{"1", "2", "3"},
	}
	// when
	err := usecase.Register(cmd)
	// then
	errString := "some error"
	assert.EqualError(t, err, errString)
}

func TestRegister_ブックマーク保存中にエラーが発生した場合はエラーを返却する(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	// given
	repository := mock_repository.NewMockBookmark(ctrl)
	repository.EXPECT().NextID().Return(sample_entity.BookmarkID())
	repository.EXPECT().Save(sample_entity.Bookmark()).Return(errors.New("some error"))
	service := mock_service.NewMockBookmark(ctrl)
	service.EXPECT().Exists(sample_entity.Bookmark()).Return(false, nil)
	usecase := NewBookmarkUsecase(repository, service)
	cmd := &command.RegisterBookmark{
		Name: "example",
		URI:  "https://example.com",
		Tags: []string{"1", "2", "3"},
	}
	// when
	err := usecase.Register(cmd)
	// then
	errString := "some error"
	assert.EqualError(t, err, errString)
}
