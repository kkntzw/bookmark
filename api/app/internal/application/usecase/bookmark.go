package usecase

import (
	"fmt"

	"github.com/kkntzw/bookmark/internal/application/command"
	"github.com/kkntzw/bookmark/internal/domain/entity"
	"github.com/kkntzw/bookmark/internal/domain/repository"
	"github.com/kkntzw/bookmark/internal/domain/service"
)

// ブックマークに関するユースケースのインターフェース。
type Bookmark interface {
	// ブックマークを登録する。
	Register(*command.RegisterBookmark) error
}

// ブックマークに関するユースケースの具象型。
type bookmarkUsecase struct {
	repository repository.Bookmark // リポジトリ
	service    service.Bookmark    // ドメインサービス
}

// ブックマークに関するユースケースを生成する。
func NewBookmarkUsecase(repository repository.Bookmark, service service.Bookmark) Bookmark {
	return &bookmarkUsecase{
		repository: repository,
		service:    service,
	}
}

// ブックマークを登録する。
//
// 不正なコマンドを受け取るとエラーを返却する。
// ブックマークが重複して存在する場合はエラーを返却する。
// リポジトリの操作中にエラーが発生した場合はエラーを返却する。
func (u *bookmarkUsecase) Register(cmd *command.RegisterBookmark) error {
	if cmd == nil {
		return fmt.Errorf("argument \"cmd\" is nil")
	}
	if err := cmd.Validate(); err != nil {
		return err
	}
	id := u.repository.NextID()
	name, _ := entity.NewName(cmd.Name)
	uri, _ := entity.NewURI(cmd.URI)
	tags := make([]entity.Tag, len(cmd.Tags))
	for i, v := range cmd.Tags {
		tag, _ := entity.NewTag(v)
		tags[i] = *tag
	}
	bookmark, _ := entity.NewBookmark(id, name, uri, tags)
	exists, err := u.service.Exists(bookmark)
	if err != nil {
		return fmt.Errorf("failed at service.Exists: %w", err)
	}
	if exists {
		return fmt.Errorf("bookmark already exists")
	}
	if err := u.repository.Save(bookmark); err != nil {
		return fmt.Errorf("failed at repository.Save: %w", err)
	}
	return nil
}
