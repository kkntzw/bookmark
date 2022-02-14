package service

import (
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kkntzw/bookmark/internal/domain/entity"
	mock_repository "github.com/kkntzw/bookmark/test/mock/domain/repository"
	"github.com/stretchr/testify/assert"
)

func ToID(t *testing.T, v string) *entity.ID {
	t.Helper()
	id, err := entity.NewID(v)
	if err != nil {
		t.Fatal(err)
	}
	return id
}

func ToBookmark(t *testing.T, iv, nv, uv string, tvs ...string) *entity.Bookmark {
	t.Helper()
	id, err := entity.NewID(iv)
	if err != nil {
		t.Fatal(err)
	}
	name, err := entity.NewName(nv)
	if err != nil {
		t.Fatal(err)
	}
	uri, err := entity.NewURI(uv)
	if err != nil {
		t.Fatal(err)
	}
	tags := make([]entity.Tag, len(tvs))
	for i, tv := range tvs {
		tag, err := entity.NewTag(tv)
		if err != nil {
			t.Fatal(err)
		}
		tags[i] = *tag
	}
	bookmark, err := entity.NewBookmark(id, name, uri, tags)
	if err != nil {
		t.Fatal(err)
	}
	return bookmark
}

func TestNewBookmarkService(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repository := mock_repository.NewMockBookmark(ctrl)
	{
		t.Run("implementing bookmark service", func(t *testing.T) {
			t.Parallel()
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
	existingBookmark := ToBookmark(t, "1", "Example A", "https://foo.example.com")
	nonExistingBookmark := ToBookmark(t, "2", "Example B", "https://bar.example.com")
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repository := mock_repository.NewMockBookmark(ctrl)
	cases := map[string]struct {
		prepare        func()
		bookmark       *entity.Bookmark
		expectedExists bool
		expectedErr    error
	}{
		"existing bookmark": {
			func() {
				repository.EXPECT().FindByID(ToID(t, "1")).Return(existingBookmark, nil)
			},
			existingBookmark,
			true,
			nil,
		},
		"non-existing bookmark": {
			func() {
				repository.EXPECT().FindByID(ToID(t, "2")).Return(nil, nil)
			},
			nonExistingBookmark,
			false,
			nil,
		},
		"nil bookmark": {
			func() {},
			nil,
			false,
			errors.New("argument \"bookmark\" is nil"),
		},
		"failed at repository.FindByID": {
			func() {
				repository.EXPECT().FindByID(ToID(t, "1")).Return(nil, errors.New("some error"))
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
			tc.prepare()
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
