package repository

import (
	"github.com/kkntzw/bookmark/internal/domain/bookmark"
)

// ブックマークの永続化を担うリポジトリのインターフェース。
type Bookmark interface {
	// IDを生成する。
	NextID() *bookmark.ID

	// IDからブックマークを検索する。
	//
	// 該当するブックマークが存在しない場合はnilを返却する。
	FindByID(id *bookmark.ID) (*bookmark.Bookmark, error)

	// ブックマークを保存する。
	Save(bookmark *bookmark.Bookmark) error
}
