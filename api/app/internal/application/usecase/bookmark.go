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
func (u *bookmarkUsecase) Register(command *command.RegisterBookmark) error {
	id := u.repository.NextID()
	name, err := entity.NewName(command.Name)
	if err != nil {
		return fmt.Errorf("command \"Name\" is invalid")
	}
	uri, err := entity.NewURI(command.URI)
	if err != nil {
		return fmt.Errorf("command \"URI\" is invalid")
	}
	tags := make([]entity.Tag, len(command.Tags))
	for i, v := range command.Tags {
		tag, err := entity.NewTag(v)
		if err != nil {
			return fmt.Errorf("command \"Tags\" is invalid")
		}
		tags[i] = *tag
	}
	bookmark, _ := entity.NewBookmark(id, name, uri, tags)
	exists, err := u.service.Exists(bookmark)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("bookmark already exists")
	}
	return u.repository.Save(bookmark)
}
