package dto

import (
	"testing"

	"github.com/kkntzw/bookmark/internal/domain/entity"
	"github.com/stretchr/testify/assert"
)

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

func TestNewBookmark(t *testing.T) {
	t.Parallel()
	// given
	entity := ToBookmark(t, "1", "Example", "https://example.com", "1-A", "1-B", "1-C")
	// when
	actualBookmark := NewBookmark(*entity)
	// then
	expectedBookmark := Bookmark{"1", "Example", "https://example.com", []string{"1-A", "1-B", "1-C"}}
	assert.Exactly(t, expectedBookmark, actualBookmark)
}
