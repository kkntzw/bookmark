package command

import (
	"github.com/kkntzw/bookmark/internal/domain/entity"
)

// ブックマーク登録用のコマンド。
type RegisterBookmark struct {
	Name string   // ブックマーク名
	URI  string   // URI
	Tags []string // タグ一覧
}

// コマンドの妥当性を検証する。
//
// コマンドが不正な場合は InvalidCommandError を返却する。
func (cmd *RegisterBookmark) Validate() error {
	var err error
	args := map[string]error{}
	_, err = entity.NewName(cmd.Name)
	if err != nil {
		args["Name"] = err
	}
	_, err = entity.NewURI(cmd.URI)
	if err != nil {
		args["URI"] = err
	}
	for _, v := range cmd.Tags {
		_, err = entity.NewTag(v)
		if err != nil {
			args["Tags"] = err
			break
		}
	}
	if len(args) > 0 {
		return &InvalidCommandError{Args: args}
	}
	return nil
}

// ブックマーク更新用のコマンド。
type UpdateBookmark struct {
	ID   string // ID
	Name string // ブックマーク名
	URI  string // URI
}

// コマンドの妥当性を検証する。
//
// コマンドが不正な場合は InvalidCommandError を返却する。
func (cmd *UpdateBookmark) Validate() error {
	args := map[string]error{}
	if _, err := entity.NewID(cmd.ID); err != nil {
		args["ID"] = err
	}
	if _, err := entity.NewName(cmd.Name); err != nil {
		args["Name"] = err
	}
	if _, err := entity.NewURI(cmd.URI); err != nil {
		args["URI"] = err
	}
	if len(args) > 0 {
		return &InvalidCommandError{Args: args}
	}
	return nil
}
