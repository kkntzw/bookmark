package dto

import (
	"testing"

	sample_entity "github.com/kkntzw/bookmark/test/data/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewBookmark_Bookmark型のインスタンスを返却する(t *testing.T) {
	// given
	entity := sample_entity.Bookmark()
	// when
	actual := NewBookmark(*entity)
	// then
	expected := Bookmark{
		ID:   "f81d4fae-7dec-11d0-a765-00a0c91e6bf6",
		Name: "example",
		URI:  "https://example.com",
		Tags: []string{"1", "2", "3"},
	}
	assert.Exactly(t, expected, actual)
}
