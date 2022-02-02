package sample_dto

import (
	"github.com/kkntzw/bookmark/internal/application/dto"
)

// dto.Bookmark型のサンプルインスタンス。
func Bookmark() dto.Bookmark {
	return dto.Bookmark{
		ID:   "f81d4fae-7dec-11d0-a765-00a0c91e6bf6",
		Name: "example",
		URI:  "https://example.com",
		Tags: []string{"1", "2", "3"},
	}
}

// dto.Bookmark型のサンプルインスタンスA。
func BookmarkA() dto.Bookmark {
	return dto.Bookmark{
		ID:   "f8ddce3a-0e87-4f3b-9f5d-148ba3125e42",
		Name: "example A",
		URI:  "https://example.com/foo",
		Tags: []string{},
	}
}

// dto.Bookmark型のサンプルインスタンスB。
func BookmarkB() dto.Bookmark {
	return dto.Bookmark{
		ID:   "7a5c72ca-6e7d-4592-abb7-363ecac0d847",
		Name: "example B",
		URI:  "https://example.com/bar",
		Tags: []string{"B-1"},
	}
}

// dto.Bookmark型のサンプルインスタンスC。
func BookmarkC() dto.Bookmark {
	return dto.Bookmark{
		ID:   "14bce21c-c9f1-43d2-a399-3e42954400f2",
		Name: "example C",
		URI:  "https://example.com/baz",
		Tags: []string{"C-1", "C-2"},
	}
}
