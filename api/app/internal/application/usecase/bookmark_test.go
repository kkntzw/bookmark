package usecase

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kkntzw/bookmark/internal/application/command"
	"github.com/kkntzw/bookmark/internal/application/dto"
	"github.com/kkntzw/bookmark/internal/domain/entity"
	sample_dto "github.com/kkntzw/bookmark/test/data/application/dto"
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
	_, errName := entity.NewName("")
	_, errUri := entity.NewURI("")
	_, errTag := entity.NewTag("")
	params := []struct {
		cmd      *command.RegisterBookmark
		expected error
	}{
		{
			cmd:      nil,
			expected: fmt.Errorf("argument \"cmd\" is nil"),
		},
		{
			cmd:      &command.RegisterBookmark{Name: "", URI: "", Tags: []string{""}},
			expected: &command.InvalidCommandError{Args: map[string]error{"Name": errName, "URI": errUri, "Tags": errTag}},
		},
	}
	for _, p := range params {
		// given
		repository := mock_repository.NewMockBookmark(ctrl)
		service := mock_service.NewMockBookmark(ctrl)
		usecase := NewBookmarkUsecase(repository, service)
		cmd := p.cmd
		// when
		actual := usecase.Register(cmd)
		// then
		assert.Exactly(t, p.expected, actual)
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
	service.EXPECT().Exists(sample_entity.Bookmark()).Return(false, fmt.Errorf("some error"))
	usecase := NewBookmarkUsecase(repository, service)
	cmd := &command.RegisterBookmark{
		Name: "example",
		URI:  "https://example.com",
		Tags: []string{"1", "2", "3"},
	}
	// when
	err := usecase.Register(cmd)
	// then
	errString := "failed at service.Exists: some error"
	assert.EqualError(t, err, errString)
}

func TestRegister_リポジトリの操作に失敗した場合はエラーを返却する(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	// given
	repository := mock_repository.NewMockBookmark(ctrl)
	repository.EXPECT().NextID().Return(sample_entity.BookmarkID())
	repository.EXPECT().Save(sample_entity.Bookmark()).Return(fmt.Errorf("some error"))
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
	errString := "failed at repository.Save: some error"
	assert.EqualError(t, err, errString)
}

func TestList_ブックマークが存在する場合はdto_Bookmark型のインスタンスが含まれたスライスを返却する(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	// given
	repository := mock_repository.NewMockBookmark(ctrl)
	repository.EXPECT().FindAll().Return([]entity.Bookmark{
		*sample_entity.BookmarkA(),
		*sample_entity.BookmarkB(),
		*sample_entity.BookmarkC(),
	}, nil)
	service := mock_service.NewMockBookmark(ctrl)
	usecase := NewBookmarkUsecase(repository, service)
	// when
	actual, err := usecase.List()
	// then
	expected := []dto.Bookmark{
		sample_dto.BookmarkA(),
		sample_dto.BookmarkB(),
		sample_dto.BookmarkC(),
	}
	assert.ElementsMatch(t, expected, actual)
	assert.NoError(t, err)
}

func TestList_ブックマークが存在しない場合は空のスライスを返却する(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	// given
	repository := mock_repository.NewMockBookmark(ctrl)
	repository.EXPECT().FindAll().Return([]entity.Bookmark{}, nil)
	service := mock_service.NewMockBookmark(ctrl)
	usecase := NewBookmarkUsecase(repository, service)
	// when
	object, err := usecase.List()
	// then
	assert.NotNil(t, object)
	assert.Empty(t, object)
	assert.NoError(t, err)
}

func TestList_リポジトリの操作に失敗した場合はエラーを返却する(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	// given
	repository := mock_repository.NewMockBookmark(ctrl)
	repository.EXPECT().FindAll().Return(nil, fmt.Errorf("some error"))
	service := mock_service.NewMockBookmark(ctrl)
	usecase := NewBookmarkUsecase(repository, service)
	// when
	object, err := usecase.List()
	// then
	assert.Nil(t, object)
	errString := "failed at repository.FindAll: some error"
	assert.EqualError(t, err, errString)
}

func TestUpdate_正当な値を受け取るとnilを返却する(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	// given
	repository := mock_repository.NewMockBookmark(ctrl)
	repository.EXPECT().FindByID(sample_entity.BookmarkID()).Return(sample_entity.Bookmark(), nil)
	repository.EXPECT().Save(sample_entity.ModifiedBookmark()).Return(nil)
	service := mock_service.NewMockBookmark(ctrl)
	usecase := NewBookmarkUsecase(repository, service)
	cmd := &command.UpdateBookmark{
		ID:   "f81d4fae-7dec-11d0-a765-00a0c91e6bf6",
		Name: "EXAMPLE",
		URI:  "http://example.com",
	}
	// when
	err := usecase.Update(cmd)
	// then
	assert.NoError(t, err)
}

func TestUpdate_不正な値を受け取るとエラーを返却する(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	_, errId := entity.NewID("")
	_, errName := entity.NewName("")
	_, errUri := entity.NewURI("")
	params := []struct {
		cmd      *command.UpdateBookmark
		expected error
	}{
		{
			cmd:      nil,
			expected: fmt.Errorf("argument \"cmd\" is nil"),
		},
		{
			cmd:      &command.UpdateBookmark{ID: "", Name: "", URI: ""},
			expected: &command.InvalidCommandError{Args: map[string]error{"ID": errId, "Name": errName, "URI": errUri}},
		},
	}
	for _, p := range params {
		// given
		repository := mock_repository.NewMockBookmark(ctrl)
		service := mock_service.NewMockBookmark(ctrl)
		usecase := NewBookmarkUsecase(repository, service)
		cmd := p.cmd
		// when
		actual := usecase.Update(cmd)
		// then
		assert.Exactly(t, p.expected, actual)
	}
}

func TestUpdate_ブックマークの検索に失敗した場合はエラーを返却する(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	// given
	repository := mock_repository.NewMockBookmark(ctrl)
	repository.EXPECT().FindByID(sample_entity.BookmarkID()).Return(nil, fmt.Errorf("some error"))
	service := mock_service.NewMockBookmark(ctrl)
	usecase := NewBookmarkUsecase(repository, service)
	cmd := &command.UpdateBookmark{
		ID:   "f81d4fae-7dec-11d0-a765-00a0c91e6bf6",
		Name: "EXAMPLE",
		URI:  "http://example.com",
	}
	// when
	err := usecase.Update(cmd)
	// then
	errString := "failed at repository.FindByID: some error"
	assert.EqualError(t, err, errString)
}

func TestUpdate_ブックマークが存在しない場合はエラーを返却する(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	// given
	repository := mock_repository.NewMockBookmark(ctrl)
	repository.EXPECT().FindByID(sample_entity.BookmarkID()).Return(nil, nil)
	service := mock_service.NewMockBookmark(ctrl)
	usecase := NewBookmarkUsecase(repository, service)
	cmd := &command.UpdateBookmark{
		ID:   "f81d4fae-7dec-11d0-a765-00a0c91e6bf6",
		Name: "EXAMPLE",
		URI:  "http://example.com",
	}
	// when
	err := usecase.Update(cmd)
	// then
	errString := "bookmark does not exist"
	assert.EqualError(t, err, errString)
}

func TestUpdate_ブックマークの保存に失敗した場合はエラーを返却する(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	// given
	repository := mock_repository.NewMockBookmark(ctrl)
	repository.EXPECT().FindByID(sample_entity.BookmarkID()).Return(sample_entity.Bookmark(), nil)
	repository.EXPECT().Save(sample_entity.ModifiedBookmark()).Return(fmt.Errorf("some error"))
	service := mock_service.NewMockBookmark(ctrl)
	usecase := NewBookmarkUsecase(repository, service)
	cmd := &command.UpdateBookmark{
		ID:   "f81d4fae-7dec-11d0-a765-00a0c91e6bf6",
		Name: "EXAMPLE",
		URI:  "http://example.com",
	}
	// when
	err := usecase.Update(cmd)
	// then
	errString := "failed at repository.Save: some error"
	assert.EqualError(t, err, errString)
}
