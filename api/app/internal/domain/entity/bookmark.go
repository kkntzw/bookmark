package entity

import "fmt"

// ブックマークを表すエンティティ。
type Bookmark struct {
	id   ID    // ID
	name Name  // ブックマーク名
	uri  URI   // URI
	tags []Tag // タグ一覧
}

// ブックマークを表すエンティティを生成する。
//
// nilを指定した場合はエラーを返却する。
//
// 複製したスライスをフィールドに設定する。
func NewBookmark(id *ID, name *Name, uri *URI, tags []Tag) (*Bookmark, error) {
	if id == nil {
		return nil, fmt.Errorf("argument \"id\" is nil")
	}
	if name == nil {
		return nil, fmt.Errorf("argument \"name\" is nil")
	}
	if uri == nil {
		return nil, fmt.Errorf("argument \"uri\" is nil")
	}
	if tags == nil {
		return nil, fmt.Errorf("argument \"tags\" is nil")
	}
	return &Bookmark{*id, *name, *uri, append([]Tag{}, tags...)}, nil
}

// フィールド id を返却する。
func (b *Bookmark) ID() ID {
	return b.id
}

// フィールド name を返却する。
func (b *Bookmark) Name() Name {
	return b.name
}

// フィールド uri を返却する。
func (b *Bookmark) URI() URI {
	return b.uri
}

// フィールド tags を返却する。
//
// 複製したスライスを返却する。
func (b *Bookmark) Tags() []Tag {
	return append([]Tag{}, b.tags...)
}

// ブックマーク名を変更する。
//
// nilを指定した場合はエラーを返却する。
func (b *Bookmark) Rename(name *Name) error {
	if name == nil {
		return fmt.Errorf("argument \"name\" is nil")
	}
	b.name = *name
	return nil
}

// インスタンスをディープコピーする。
func (b Bookmark) DeepCopy() *Bookmark {
	copy := &b
	copy.tags = append([]Tag{}, b.tags...)
	return copy
}
