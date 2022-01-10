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
func (r *bookmarkRepository) NextID() *bookmark.ID {
	uuid, _ := uuid.NewRandom()
	id, _ := bookmark.NewID(uuid.String())
	return id
}

// ブックマークを保存する。
//
// nilを指定した場合はエラーを返却する。
//
// 複製したインスタンスをストレージに保存する。
func (r *bookmarkRepository) Save(bookmark *bookmark.Bookmark) error {
	if bookmark == nil {
		return fmt.Errorf("argument \"bookmark\" is nil")
	}
	r.store[bookmark.ID()] = *bookmark.DeepCopy()
	return nil
}

// IDからブックマークを検索する。
//
// 該当するブックマークが存在する場合は複製したインスタンスを返却する。
// nilを指定した場合はエラーを返却する。
// 該当するブックマークが存在しない場合はnilを返却する。
func (r *bookmarkRepository) FindByID(id *bookmark.ID) (*bookmark.Bookmark, error) {
	if id == nil {
		return nil, fmt.Errorf("argument \"id\" is nil")
	}
	bookmark, ok := r.store[*id]
	if !ok {
		return nil, nil
	}
	return bookmark.DeepCopy(), nil
}
