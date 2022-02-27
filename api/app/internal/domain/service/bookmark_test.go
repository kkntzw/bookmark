package service

import (
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kkntzw/bookmark/internal/domain/entity"
	"github.com/kkntzw/bookmark/test/helper"
	mock_repository "github.com/kkntzw/bookmark/test/mock/domain/repository"
	"github.com/stretchr/testify/assert"
)

func TestNewBookmarkService(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	t.Run("implementing service.Bookmark", func(t *testing.T) {
		t.Parallel()
		// given
		repository := mock_repository.NewMockBookmark(ctrl)
		// when
		object := NewBookmarkService(repository)
		// then
		assert.NotNil(t, object)
		interfaceObject := (*Bookmark)(nil)
		assert.Implements(t, interfaceObject, object)
	})
	t.Run("fields", func(t *testing.T) {
		t.Parallel()
		// given
		repository := mock_repository.NewMockBookmark(ctrl)
		abstractService := NewBookmarkService(repository)
		// when
		concreteService, ok := abstractService.(*bookmarkService)
		actualRepository := concreteService.repository
		// then
		assert.True(t, ok)
		expectedRepository := repository
		assert.Exactly(t, expectedRepository, actualRepository)
	})
}

func TestBookmark_Exists(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cases := map[string]struct {
		prepare        func(*mock_repository.MockBookmark)
		bookmark       *entity.Bookmark
		expectedExists bool
		expectedErr    error
	}{
		"existing bookmark": {
			func(repository *mock_repository.MockBookmark) {
				repository.EXPECT().FindByID(helper.ToID(t, "1")).Return(helper.ToBookmark(t, "1", "Example", "https://example.com"), nil)
			},
			helper.ToBookmark(t, "1", "Example", "https://example.com"),
			true,
			nil,
		},
		"non-existing bookmark": {
			func(repository *mock_repository.MockBookmark) {
				repository.EXPECT().FindByID(helper.ToID(t, "1")).Return(nil, nil)
			},
			helper.ToBookmark(t, "1", "Example", "https://example.com"),
			false,
			nil,
		},
		"nil bookmark": {
			func(repository *mock_repository.MockBookmark) {},
			nil,
			false,
			errors.New("argument \"bookmark\" is nil"),
		},
		"failed at repository.FindByID": {
			func(repository *mock_repository.MockBookmark) {
				repository.EXPECT().FindByID(helper.ToID(t, "1")).Return(nil, errors.New("some error"))
			},
			helper.ToBookmark(t, "1", "Example", "https://example.com"),
			false,
			fmt.Errorf("failed at repository.FindByID: %w", errors.New("some error")),
		},
	}
	for name, tc := range cases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			repository := mock_repository.NewMockBookmark(ctrl)
			tc.prepare(repository)
			// given
			service := NewBookmarkService(repository)
			// when
			actualExists, actualErr := service.Exists(tc.bookmark)
			// then
			assert.Exactly(t, tc.expectedExists, actualExists)
			assert.Exactly(t, tc.expectedErr, actualErr)
		})
	}
}
