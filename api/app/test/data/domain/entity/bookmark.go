package sample_entity

import (
	"github.com/kkntzw/bookmark/internal/domain/entity"
)

// entity.Bookmark型のサンプルインスタンス。
func Bookmark() *entity.Bookmark {
	id, _ := entity.NewID("f81d4fae-7dec-11d0-a765-00a0c91e6bf6")
	name, _ := entity.NewName("example")
	uri, _ := entity.NewURI("https://example.com")
	tag1, _ := entity.NewTag("1")
	tag2, _ := entity.NewTag("2")
	tag3, _ := entity.NewTag("3")
	tags := []entity.Tag{*tag1, *tag2, *tag3}
	bookmark, _ := entity.NewBookmark(id, name, uri, tags)
	return bookmark
}

// entity.ID型のサンプルインスタンス。
func BookmarkID() *entity.ID {
	id, _ := entity.NewID("f81d4fae-7dec-11d0-a765-00a0c91e6bf6")
	return id
}
