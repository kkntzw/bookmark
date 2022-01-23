package inmemory

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/kkntzw/bookmark/internal/domain/entity"
	"github.com/kkntzw/bookmark/internal/domain/repository"
)

// ブックマークの永続化を担うリポジトリの具象型。
type bookmarkRepository struct {
	store map[entity.ID]entity.Bookmark // ストレージ
}

// ブックマークの永続化を担うリポジトリを生成する。
func NewBookmarkRepository() repository.Bookmark {
	return &bookmarkRepository{
		store: make(map[entity.ID]entity.Bookmark),
	}
}

// IDを生成する。
//
// バージョン4のUUIDを16進表記で生成する。
func (r *bookmarkRepository) NextID() *entity.ID {
	uuid, _ := uuid.NewRandom()
	id, _ := entity.NewID(uuid.String())
	return id
}

// ブックマークを保存する。
//
// nilを指定した場合はエラーを返却する。
//
// 複製したインスタンスをストレージに保存する。
func (r *bookmarkRepository) Save(bookmark *entity.Bookmark) error {
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
func (r *bookmarkRepository) FindByID(id *entity.ID) (*entity.Bookmark, error) {
	if id == nil {
		return nil, fmt.Errorf("argument \"id\" is nil")
	}
	bookmark, ok := r.store[*id]
	if !ok {
		return nil, nil
	}
	return bookmark.DeepCopy(), nil
}
