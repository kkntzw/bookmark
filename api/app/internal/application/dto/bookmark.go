package dto

import (
	"github.com/kkntzw/bookmark/internal/domain/entity"
)

// ブックマークを表すDTO。
type Bookmark struct {
	ID   string   // ID
	Name string   // ブックマーク名
	URI  string   // URI
	Tags []string // タグ一覧
}

// ブックマークを表すエンティティからDTOを生成する。
func NewBookmark(entity entity.Bookmark) Bookmark {
	id := entity.ID()
	name := entity.Name()
	uri := entity.URI()
	tags := make([]string, len(entity.Tags()))
	for i, tag := range entity.Tags() {
		tags[i] = tag.Value()
	}
	return Bookmark{id.Value(), name.Value(), uri.String(), tags}
}
