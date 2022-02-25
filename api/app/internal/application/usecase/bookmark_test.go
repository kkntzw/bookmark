package usecase

import (
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kkntzw/bookmark/internal/application/command"
	"github.com/kkntzw/bookmark/internal/application/dto"
	"github.com/kkntzw/bookmark/internal/domain/entity"
	"github.com/kkntzw/bookmark/test/helper"
	mock_repository "github.com/kkntzw/bookmark/test/mock/domain/repository"
	mock_service "github.com/kkntzw/bookmark/test/mock/domain/service"
	"github.com/stretchr/testify/assert"
)

func TestNewBookmarkUsecase(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	t.Run("implementing usecase.Bookmark", func(t *testing.T) {
		t.Parallel()
		// given
		repository := mock_repository.NewMockBookmark(ctrl)
		service := mock_service.NewMockBookmark(ctrl)
		// when
		object := NewBookmarkUsecase(repository, service)
		// then
		assert.NotNil(t, object)
		interfaceObject := (*Bookmark)(nil)
		assert.Implements(t, interfaceObject, object)
	})
	t.Run("fields", func(t *testing.T) {
		t.Parallel()
		// given
		repository := mock_repository.NewMockBookmark(ctrl)
		service := mock_service.NewMockBookmark(ctrl)
		abstractUsecase := NewBookmarkUsecase(repository, service)
		// when
		concreteUsecase, ok := abstractUsecase.(*bookmarkUsecase)
		actualRepository := concreteUsecase.repository
		actualService := concreteUsecase.service
		// then
		assert.True(t, ok)
		expectedRepository := repository
		assert.Exactly(t, expectedRepository, actualRepository)
		expectedService := service
		assert.Exactly(t, expectedService, actualService)
	})
}

func TestBookmark_Register(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cases := map[string]struct {
		prepare     func(*mock_repository.MockBookmark, *mock_service.MockBookmark)
		cmd         *command.RegisterBookmark
		expectedErr error
	}{
		"non-nil command": {
			func(repository *mock_repository.MockBookmark, service *mock_service.MockBookmark) {
				repository.EXPECT().NextID().Return(helper.ToID(t, "1"))
				repository.EXPECT().Save(helper.ToBookmark(t, "1", "Example", "https://example.com", "foo", "bar")).Return(nil)
				service.EXPECT().Exists(helper.ToBookmark(t, "1", "Example", "https://example.com", "foo", "bar")).Return(false, nil)
			},
			&command.RegisterBookmark{Name: "Example", URI: "https://example.com", Tags: []string{"foo", "bar"}},
			nil,
		},
		"nil command": {
			func(repository *mock_repository.MockBookmark, service *mock_service.MockBookmark) {},
			nil,
			errors.New("argument \"cmd\" is nil"),
		},
		"invalid command": {
			func(repository *mock_repository.MockBookmark, service *mock_service.MockBookmark) {},
			&command.RegisterBookmark{Name: "Example", URI: "https://example.com", Tags: []string{""}},
			&command.InvalidCommandError{Args: map[string]error{"Tags": helper.ToErrTag(t, "")}},
		},
		"duplicate bookmark": {
			func(repository *mock_repository.MockBookmark, service *mock_service.MockBookmark) {
				repository.EXPECT().NextID().Return(helper.ToID(t, "1"))
				service.EXPECT().Exists(helper.ToBookmark(t, "1", "Example", "https://example.com", "foo", "bar")).Return(true, nil)
			},
			&command.RegisterBookmark{Name: "Example", URI: "https://example.com", Tags: []string{"foo", "bar"}},
			errors.New("bookmark already exists"),
		},
		"failed at service.Exists": {
			func(repository *mock_repository.MockBookmark, service *mock_service.MockBookmark) {
				repository.EXPECT().NextID().Return(helper.ToID(t, "1"))
				service.EXPECT().Exists(helper.ToBookmark(t, "1", "Example", "https://example.com", "foo", "bar")).Return(false, errors.New("some error"))
			},
			&command.RegisterBookmark{Name: "Example", URI: "https://example.com", Tags: []string{"foo", "bar"}},
			fmt.Errorf("failed at service.Exists: %w", errors.New("some error")),
		},
		"failed at repository.Save": {
			func(repository *mock_repository.MockBookmark, service *mock_service.MockBookmark) {
				repository.EXPECT().NextID().Return(helper.ToID(t, "1"))
				repository.EXPECT().Save(helper.ToBookmark(t, "1", "Example", "https://example.com", "foo", "bar")).Return(errors.New("some error"))
				service.EXPECT().Exists(helper.ToBookmark(t, "1", "Example", "https://example.com", "foo", "bar")).Return(false, nil)
			},
			&command.RegisterBookmark{Name: "Example", URI: "https://example.com", Tags: []string{"foo", "bar"}},
			fmt.Errorf("failed at repository.Save: %w", errors.New("some error")),
		},
	}
	for name, tc := range cases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			repository := mock_repository.NewMockBookmark(ctrl)
			service := mock_service.NewMockBookmark(ctrl)
			tc.prepare(repository, service)
			// given
			usecase := NewBookmarkUsecase(repository, service)
			// when
			actualErr := usecase.Register(tc.cmd)
			// then
			assert.Exactly(t, tc.expectedErr, actualErr)
		})
	}
}

func TestBookmark_List(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cases := map[string]struct {
		prepare           func(*mock_repository.MockBookmark)
		expectedBookmarks []dto.Bookmark
		expectedErr       error
	}{
		"3 bookmarks": {
			func(repository *mock_repository.MockBookmark) {
				repository.EXPECT().FindAll().Return(
					[]entity.Bookmark{
						*helper.ToBookmark(t, "1", "Example A", "https://foo.example.com"),
						*helper.ToBookmark(t, "2", "Example B", "https://bar.example.com", "2-A"),
						*helper.ToBookmark(t, "3", "Example C", "https://baz.example.com", "3-A", "3-B"),
					},
					nil,
				)
			},
			[]dto.Bookmark{
				{ID: "1", Name: "Example A", URI: "https://foo.example.com", Tags: []string{}},
				{ID: "2", Name: "Example B", URI: "https://bar.example.com", Tags: []string{"2-A"}},
				{ID: "3", Name: "Example C", URI: "https://baz.example.com", Tags: []string{"3-A", "3-B"}},
			},
			nil,
		},
		"no bookmarks": {
			func(repository *mock_repository.MockBookmark) {
				repository.EXPECT().FindAll().Return([]entity.Bookmark{}, nil)
			},
			[]dto.Bookmark{},
			nil,
		},
		"failed at repository.FindAll": {
			func(repository *mock_repository.MockBookmark) {
				repository.EXPECT().FindAll().Return(nil, errors.New("some error"))
			},
			nil,
			fmt.Errorf("failed at repository.FindAll: %w", errors.New("some error")),
		},
	}
	for name, tc := range cases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			repository := mock_repository.NewMockBookmark(ctrl)
			service := mock_service.NewMockBookmark(ctrl)
			tc.prepare(repository)
			// given
			usecase := NewBookmarkUsecase(repository, service)
			// when
			actualBookmarks, actualErr := usecase.List()
			// then
			assert.Exactly(t, tc.expectedBookmarks, actualBookmarks)
			assert.Exactly(t, tc.expectedErr, actualErr)
		})
	}
}

func TestBookmark_Update(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cases := map[string]struct {
		prepare     func(*mock_repository.MockBookmark)
		cmd         *command.UpdateBookmark
		expectedErr error
	}{
		"non-nil command": {
			func(repository *mock_repository.MockBookmark) {
				repository.EXPECT().FindByID(helper.ToID(t, "1")).Return(helper.ToBookmark(t, "1", "Example", "http://example.com", "foo", "bar", "baz"), nil)
				repository.EXPECT().Save(helper.ToBookmark(t, "1", "EXAMPLE", "https://example.com", "foo", "bar", "baz")).Return(nil)
			},
			&command.UpdateBookmark{ID: "1", Name: "EXAMPLE", URI: "https://example.com"},
			nil,
		},
		"nil command": {
			func(repository *mock_repository.MockBookmark) {},
			nil,
			errors.New("argument \"cmd\" is nil"),
		},
		"invalid command": {
			func(repository *mock_repository.MockBookmark) {},
			&command.UpdateBookmark{ID: "1", Name: "EXAMPLE", URI: ""},
			&command.InvalidCommandError{Args: map[string]error{"URI": helper.ToErrURI(t, "")}},
		},
		"non-existent bookmark": {
			func(repository *mock_repository.MockBookmark) {
				repository.EXPECT().FindByID(helper.ToID(t, "1")).Return(nil, nil)
			},
			&command.UpdateBookmark{ID: "1", Name: "EXAMPLE", URI: "https://example.com"},
			errors.New("bookmark does not exist"),
		},
		"failed at repository.FindByID": {
			func(repository *mock_repository.MockBookmark) {
				repository.EXPECT().FindByID(helper.ToID(t, "1")).Return(nil, errors.New("some error"))
			},
			&command.UpdateBookmark{ID: "1", Name: "EXAMPLE", URI: "https://example.com"},
			fmt.Errorf("failed at repository.FindByID: %w", errors.New("some error")),
		},
		"failed at repository.Save": {
			func(repository *mock_repository.MockBookmark) {
				repository.EXPECT().FindByID(helper.ToID(t, "1")).Return(helper.ToBookmark(t, "1", "Example", "http://example.com", "foo", "bar", "baz"), nil)
				repository.EXPECT().Save(helper.ToBookmark(t, "1", "EXAMPLE", "https://example.com", "foo", "bar", "baz")).Return(errors.New("some error"))
			},
			&command.UpdateBookmark{ID: "1", Name: "EXAMPLE", URI: "https://example.com"},
			fmt.Errorf("failed at repository.Save: %w", errors.New("some error")),
		},
	}
	for name, tc := range cases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			repository := mock_repository.NewMockBookmark(ctrl)
			service := mock_service.NewMockBookmark(ctrl)
			tc.prepare(repository)
			// given
			usecase := NewBookmarkUsecase(repository, service)
			// when
			actualErr := usecase.Update(tc.cmd)
			// then
			assert.Exactly(t, tc.expectedErr, actualErr)
		})
	}
}
