package service

import (
	"fmt"

	"github.com/kkntzw/bookmark/internal/domain/entity"
	"github.com/kkntzw/bookmark/internal/domain/repository"
)

// ブックマークに関するドメインサービスのインターフェース。
type Bookmark interface {
	// ブックマークが存在するか確認する。
	Exists(bookmark *entity.Bookmark) (bool, error)
}

// ブックマークに関するドメインサービスの具象型。
type bookmarkService struct {
	repository repository.Bookmark // リポジトリ
}

// ブックマークに関するドメインサービスを生成する。
func NewBookmarkService(repository repository.Bookmark) Bookmark {
	return &bookmarkService{
		repository: repository,
	}
}

// ブックマークが存在するか確認する。
//
// nilを指定した場合はエラーを返却する。
// ブックマークの検索に失敗した場合はエラーを返却する。
func (s *bookmarkService) Exists(bookmark *entity.Bookmark) (bool, error) {
	if bookmark == nil {
		return false, fmt.Errorf("argument \"bookmark\" is nil")
	}
	id := bookmark.ID()
	object, err := s.repository.FindByID(&id)
	if err != nil {
		return false, fmt.Errorf("failed at repository.FindByID: %w", err)
	}
	exists := object != nil
	return exists, nil
}
