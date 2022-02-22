package dto

import (
	"testing"

	"github.com/kkntzw/bookmark/test/helper"
	"github.com/stretchr/testify/assert"
)

func TestNewBookmark(t *testing.T) {
	t.Parallel()
	// given
	entity := helper.ToBookmark(t, "1", "Example", "https://example.com", "1-A", "1-B", "1-C")
	// when
	actualBookmark := NewBookmark(*entity)
	// then
	expectedBookmark := Bookmark{"1", "Example", "https://example.com", []string{"1-A", "1-B", "1-C"}}
	assert.Exactly(t, expectedBookmark, actualBookmark)
}
