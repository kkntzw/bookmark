package inmemory

import (
	"errors"
	"testing"

	"github.com/kkntzw/bookmark/internal/domain/entity"
	"github.com/kkntzw/bookmark/internal/domain/repository"
	"github.com/kkntzw/bookmark/test/helper"
	"github.com/stretchr/testify/assert"
)

func TestNewBookmarkRepository(t *testing.T) {
	t.Parallel()
	t.Run("implementing repository.Bookmark", func(t *testing.T) {
		t.Parallel()
		// when
		object := NewBookmarkRepository()
		// then
		assert.NotNil(t, object)
		interfaceObject := (*repository.Bookmark)(nil)
		assert.Implements(t, interfaceObject, object)
	})
	t.Run("fields", func(t *testing.T) {
		t.Parallel()
		// given
		abstractRepository := NewBookmarkRepository()
		// when
		concreteRepository, ok := abstractRepository.(*bookmarkRepository)
		actualStore := concreteRepository.store
		// then
		assert.True(t, ok)
		expectedStore := map[entity.ID]entity.Bookmark{}
		assert.Exactly(t, expectedStore, actualStore)
	})
}

func TestBookmark_NextID(t *testing.T) {
	t.Parallel()
	// given
	repository := NewBookmarkRepository()
	// when
	id := repository.NextID()
	// then
	assert.NotNil(t, id)
	expectedType := &entity.ID{}
	assert.IsType(t, expectedType, id)
}

func TestBookmark_Save(t *testing.T) {
	t.Parallel()
	cases := map[string]struct {
		bookmark    *entity.Bookmark
		expectedErr error
	}{
		"non-nil bookmark": {
			helper.ToBookmark(t, "1", "Example", "https://example.com", "foo", "bar", "baz"),
			nil,
		},
		"nil bookmark": {
			nil,
			errors.New("argument \"bookmark\" is nil"),
		},
	}
	for name, tc := range cases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			// given
			repository := NewBookmarkRepository()
			// when
			actualErr := repository.Save(tc.bookmark)
			// then
			assert.Exactly(t, tc.expectedErr, actualErr)
		})
	}
}

func TestBookmark_FindAll(t *testing.T) {
	t.Parallel()
	cases := map[string]struct {
		prepare           func(repository.Bookmark)
		expectedBookmarks []entity.Bookmark
		expectedErr       error
	}{
		"stored bookmarks": {
			func(r repository.Bookmark) {
				r.Save(helper.ToBookmark(t, "1", "Example A", "https://foo.example.com"))
				r.Save(helper.ToBookmark(t, "2", "Example B", "https://bar.example.com"))
				r.Save(helper.ToBookmark(t, "3", "Example C", "https://baz.example.com"))
			},
			[]entity.Bookmark{
				*helper.ToBookmark(t, "1", "Example A", "https://foo.example.com"),
				*helper.ToBookmark(t, "2", "Example B", "https://bar.example.com"),
				*helper.ToBookmark(t, "3", "Example C", "https://baz.example.com"),
			},
			nil,
		},
		"unstored bookmarks": {
			func(r repository.Bookmark) {},
			[]entity.Bookmark{},
			nil,
		},
	}
	for name, tc := range cases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			// given
			repository := NewBookmarkRepository()
			tc.prepare(repository)
			// when
			actualBookmarks, actualErr := repository.FindAll()
			// then
			assert.ElementsMatch(t, tc.expectedBookmarks, actualBookmarks)
			assert.Exactly(t, tc.expectedErr, actualErr)
		})
	}
}

func TestBookmark_FindByID(t *testing.T) {
	t.Parallel()
	cases := map[string]struct {
		prepare          func(repository.Bookmark)
		id               *entity.ID
		expectedBookmark *entity.Bookmark
		expectedErr      error
	}{
		"id of stored bookmark": {
			func(r repository.Bookmark) {
				r.Save(helper.ToBookmark(t, "1", "Example", "https://example.com"))
			},
			helper.ToID(t, "1"),
			helper.ToBookmark(t, "1", "Example", "https://example.com"),
			nil,
		},
		"id of unstored bookmark": {
			func(r repository.Bookmark) {},
			helper.ToID(t, "1"),
			nil,
			nil,
		},
		"nil id": {
			func(r repository.Bookmark) {},
			nil,
			nil,
			errors.New("argument \"id\" is nil"),
		},
	}
	for name, tc := range cases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			// given
			repository := NewBookmarkRepository()
			tc.prepare(repository)
			// when
			actualBookmark, actualErr := repository.FindByID(tc.id)
			// then
			assert.Exactly(t, tc.expectedBookmark, actualBookmark)
			assert.Exactly(t, tc.expectedErr, actualErr)
		})
	}
}

func TestBookmark_Delete(t *testing.T) {
	t.Parallel()
	cases := map[string]struct {
		prepare     func(repository.Bookmark)
		bookmark    *entity.Bookmark
		expectedErr error
	}{
		"stored bookmark": {
			func(r repository.Bookmark) {
				r.Save(helper.ToBookmark(t, "1", "Example", "https://example.com", "foo", "bar", "baz"))
			},
			helper.ToBookmark(t, "1", "Example", "https://example.com", "foo", "bar", "baz"),
			nil,
		},
		"unstored bookmark": {
			func(r repository.Bookmark) {},
			helper.ToBookmark(t, "1", "Example", "https://example.com", "foo", "bar", "baz"),
			nil,
		},
		"nil bookmark": {
			func(r repository.Bookmark) {},
			nil,
			errors.New("argument \"bookmark\" is nil"),
		},
	}
	for name, tc := range cases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			// given
			repository := NewBookmarkRepository()
			tc.prepare(repository)
			// when
			actualErr := repository.Delete(tc.bookmark)
			// then
			assert.Exactly(t, tc.expectedErr, actualErr)
		})
	}
}
