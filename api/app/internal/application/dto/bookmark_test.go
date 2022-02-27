package dto

import (
	"testing"

	"github.com/kkntzw/bookmark/internal/domain/entity"
	"github.com/kkntzw/bookmark/test/helper"
	"github.com/stretchr/testify/assert"
)

func TestNewBookmark(t *testing.T) {
	t.Parallel()
	cases := map[string]struct {
		entity           entity.Bookmark
		expectedBookmark Bookmark
	}{
		"valid entity (empty tags)": {
			*helper.ToBookmark(t, "1", "Example", "https://example.com"),
			Bookmark{"1", "Example", "https://example.com", []string{}},
		},
		"valid entity (3 tags)": {
			*helper.ToBookmark(t, "1", "Example", "https://example.com", "foo", "bar", "baz"),
			Bookmark{"1", "Example", "https://example.com", []string{"foo", "bar", "baz"}},
		},
	}
	for name, tc := range cases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			// when
			actualBookmark := NewBookmark(tc.entity)
			// then
			assert.Exactly(t, tc.expectedBookmark, actualBookmark)
		})
	}
}
