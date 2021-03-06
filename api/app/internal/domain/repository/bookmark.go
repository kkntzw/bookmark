package repository

import (
	"github.com/kkntzw/bookmark/internal/domain/entity"
)

// ブックマークの永続化を担うリポジトリのインターフェース。
type Bookmark interface {
	// IDを生成する。
	NextID() *entity.ID

	// ブックマークを保存する。
	Save(bookmark *entity.Bookmark) error

	// ブックマーク一覧を検索する。
	//
	// ブックマークが存在しない場合は空のスライスを返却する。
	FindAll() ([]entity.Bookmark, error)

	// IDからブックマークを検索する。
	//
	// 該当するブックマークが存在しない場合はnilを返却する。
	FindByID(id *entity.ID) (*entity.Bookmark, error)

	// ブックマークを削除する。
	Delete(bookmark *entity.Bookmark) error
}
