package inmemory

import (
	"errors"
	"testing"

	"github.com/kkntzw/bookmark/internal/domain/entity"
	"github.com/kkntzw/bookmark/internal/domain/repository"
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

func TestNewBookmarkRepository(t *testing.T) {
	t.Parallel()
	{
		t.Run("implementing bookmark repository", func(t *testing.T) {
			t.Parallel()
			// when
			object := NewBookmarkRepository()
			// then
			assert.NotNil(t, object)
			interfaceObject := (*repository.Bookmark)(nil)
			assert.Implements(t, interfaceObject, object)
		})
	}
	{
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
		"non-nil bookmark": {ToBookmark(t, "1", "Example A", "https://foo.example.com"), nil},
		"nil bookmark":     {nil, errors.New("argument \"bookmark\" is nil")},
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
	bookmark1 := ToBookmark(t, "1", "Example A", "https://foo.example.com")
	bookmark2 := ToBookmark(t, "2", "Example B", "https://bar.example.com")
	bookmark3 := ToBookmark(t, "3", "Example C", "https://baz.example.com")
	cases := map[string]struct {
		prepare           func(repository.Bookmark)
		expectedBookmarks []entity.Bookmark
		expectedErr       error
	}{
		"stored bookmarks": {
			func(r repository.Bookmark) {
				r.Save(bookmark1)
				r.Save(bookmark2)
				r.Save(bookmark3)
			},
			[]entity.Bookmark{*bookmark1, *bookmark2, *bookmark3},
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
	bookmark := ToBookmark(t, "1", "Example A", "https://foo.example.com")
	cases := map[string]struct {
		prepare          func(repository.Bookmark)
		id               *entity.ID
		expectedBookmark *entity.Bookmark
		expectedErr      error
	}{
		"stored bookmark":   {func(r repository.Bookmark) { r.Save(bookmark) }, ToID(t, "1"), bookmark, nil},
		"unstored bookmark": {func(r repository.Bookmark) {}, ToID(t, "1"), nil, nil},
		"nil id":            {func(r repository.Bookmark) {}, nil, nil, errors.New("argument \"id\" is nil")},
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
