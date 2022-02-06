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

// entity.Bookmark型のサンプルインスタンス。
func ModifiedBookmark() *entity.Bookmark {
	id, _ := entity.NewID("f81d4fae-7dec-11d0-a765-00a0c91e6bf6")
	name, _ := entity.NewName("EXAMPLE")
	uri, _ := entity.NewURI("http://example.com")
	tag1, _ := entity.NewTag("1")
	tag2, _ := entity.NewTag("2")
	tag3, _ := entity.NewTag("3")
	tags := []entity.Tag{*tag1, *tag2, *tag3}
	bookmark, _ := entity.NewBookmark(id, name, uri, tags)
	return bookmark
}

// entity.Bookmark型のサンプルインスタンスA。
func BookmarkA() *entity.Bookmark {
	id, _ := entity.NewID("f8ddce3a-0e87-4f3b-9f5d-148ba3125e42")
	name, _ := entity.NewName("example A")
	uri, _ := entity.NewURI("https://example.com/foo")
	tags := []entity.Tag{}
	bookmark, _ := entity.NewBookmark(id, name, uri, tags)
	return bookmark
}

// entity.Bookmark型のサンプルインスタンスB。
func BookmarkB() *entity.Bookmark {
	id, _ := entity.NewID("7a5c72ca-6e7d-4592-abb7-363ecac0d847")
	name, _ := entity.NewName("example B")
	uri, _ := entity.NewURI("https://example.com/bar")
	tag1, _ := entity.NewTag("B-1")
	tags := []entity.Tag{*tag1}
	bookmark, _ := entity.NewBookmark(id, name, uri, tags)
	return bookmark
}

// entity.Bookmark型のサンプルインスタンスC。
func BookmarkC() *entity.Bookmark {
	id, _ := entity.NewID("14bce21c-c9f1-43d2-a399-3e42954400f2")
	name, _ := entity.NewName("example C")
	uri, _ := entity.NewURI("https://example.com/baz")
	tag1, _ := entity.NewTag("C-1")
	tag2, _ := entity.NewTag("C-2")
	tags := []entity.Tag{*tag1, *tag2}
	bookmark, _ := entity.NewBookmark(id, name, uri, tags)
	return bookmark
}

// entity.ID型のサンプルインスタンス。
func BookmarkID() *entity.ID {
	id, _ := entity.NewID("f81d4fae-7dec-11d0-a765-00a0c91e6bf6")
	return id
}
