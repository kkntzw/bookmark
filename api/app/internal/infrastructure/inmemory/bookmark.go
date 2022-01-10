package inmemory

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/kkntzw/bookmark/internal/domain/bookmark"
)

// ブックマークの永続化を担うリポジトリの具象型。
type bookmarkRepository struct {
	store map[bookmark.ID]bookmark.Bookmark // ストレージ
}

// ブックマークの永続化を担うリポジトリを生成する。
func NewBookmarkRepository() bookmark.Repository {
	return &bookmarkRepository{
		store: make(map[bookmark.ID]bookmark.Bookmark),
	}
}

// IDを生成する。
//
// バージョン4のUUIDを16進表記で生成する。
func (b *bookmarkRepository) NextID() *bookmark.ID {
	uuid, _ := uuid.NewRandom()
	id, _ := bookmark.NewID(uuid.String())
	return id
}

// ブックマークを保存する。
//
// nilを指定した場合はエラーを返却する。
//
// 複製したインスタンスをストレージに保存する。
func (b *bookmarkRepository) Save(bookmark *bookmark.Bookmark) error {
	if bookmark == nil {
		return fmt.Errorf("argument \"bookmark\" is nil")
	}
	b.store[bookmark.ID()] = *bookmark.DeepCopy()
	return nil
}