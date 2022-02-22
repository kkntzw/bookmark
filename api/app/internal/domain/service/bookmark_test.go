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
	{
		t.Run("implementing bookmark service", func(t *testing.T) {
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
	}
	{
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
}

func TestBookmark_Exists(t *testing.T) {
	t.Parallel()
	existingBookmark := helper.ToBookmark(t, "1", "Example A", "https://foo.example.com")
	nonExistingBookmark := helper.ToBookmark(t, "2", "Example B", "https://bar.example.com")
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
				repository.EXPECT().FindByID(helper.ToID(t, "1")).Return(existingBookmark, nil)
			},
			existingBookmark,
			true,
			nil,
		},
		"non-existing bookmark": {
			func(repository *mock_repository.MockBookmark) {
				repository.EXPECT().FindByID(helper.ToID(t, "2")).Return(nil, nil)
			},
			nonExistingBookmark,
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
			existingBookmark,
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
