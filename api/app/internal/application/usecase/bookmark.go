package usecase

import (
	"fmt"

	"github.com/kkntzw/bookmark/internal/application/command"
	"github.com/kkntzw/bookmark/internal/application/dto"
	"github.com/kkntzw/bookmark/internal/domain/entity"
	"github.com/kkntzw/bookmark/internal/domain/repository"
	"github.com/kkntzw/bookmark/internal/domain/service"
)

// ブックマークに関するユースケースのインターフェース。
type Bookmark interface {
	// ブックマークを登録する。
	Register(*command.RegisterBookmark) error

	// ブックマークを一覧取得する。
	List() ([]dto.Bookmark, error)

	// ブックマークを更新する。
	Update(*command.UpdateBookmark) error

	// ブックマークを削除する。
	Delete(*command.DeleteBookmark) error
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
// nilを指定した場合はエラーを返却する。
// 不正なコマンドを指定した場合はエラーを返却する。
// ブックマークの存在確認に失敗した場合はエラーを返却する。
// ブックマークが存在する場合はエラーを返却する。
// ブックマークの保存に失敗した場合はエラーを返却する。
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

// ブックマークを一覧取得する。
//
// ブックマークの検索に失敗した場合はエラーを返却する。
func (u *bookmarkUsecase) List() ([]dto.Bookmark, error) {
	entities, err := u.repository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed at repository.FindAll: %w", err)
	}
	bookmarks := make([]dto.Bookmark, len(entities))
	for i, entity := range entities {
		bookmarks[i] = dto.NewBookmark(entity)
	}
	return bookmarks, nil
}

// ブックマークを更新する。
//
// nilを指定した場合はエラーを返却する。
// 不正なコマンドを指定した場合はエラーを返却する。
// ブックマークの検索に失敗した場合はエラーを返却する。
// ブックマークが存在しない場合はエラーを返却する。
// ブックマークの保存に失敗した場合はエラーを返却する。
func (u *bookmarkUsecase) Update(cmd *command.UpdateBookmark) error {
	if cmd == nil {
		return fmt.Errorf("argument \"cmd\" is nil")
	}
	if err := cmd.Validate(); err != nil {
		return err
	}
	id, _ := entity.NewID(cmd.ID)
	bookmark, err := u.repository.FindByID(id)
	if err != nil {
		return fmt.Errorf("failed at repository.FindByID: %w", err)
	}
	if bookmark == nil {
		return fmt.Errorf("bookmark does not exist")
	}
	name, _ := entity.NewName(cmd.Name)
	bookmark.Rename(name)
	uri, _ := entity.NewURI(cmd.URI)
	bookmark.RewriteURI(uri)
	if err := u.repository.Save(bookmark); err != nil {
		return fmt.Errorf("failed at repository.Save: %w", err)
	}
	return nil
}

// ブックマークを削除する。
//
// nilを指定した場合はエラーを返却する。
// 不正なコマンドを指定した場合はエラーを返却する。
// ブックマークの検索に失敗した場合はエラーを返却する。
// ブックマークが存在しない場合はエラーを返却する。
// ブックマークの削除に失敗した場合はエラーを返却する。
func (u *bookmarkUsecase) Delete(cmd *command.DeleteBookmark) error {
	if cmd == nil {
		return fmt.Errorf("argument \"cmd\" is nil")
	}
	if err := cmd.Validate(); err != nil {
		return err
	}
	id, _ := entity.NewID(cmd.ID)
	bookmark, err := u.repository.FindByID(id)
	if err != nil {
		return fmt.Errorf("failed at repository.FindByID: %w", err)
	}
	if bookmark == nil {
		return fmt.Errorf("bookmark does not exist")
	}
	if err := u.repository.Delete(bookmark); err != nil {
		return fmt.Errorf("failed at repository.Delete: %w", err)
	}
	return nil
}
