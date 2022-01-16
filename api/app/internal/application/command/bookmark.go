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
	args := []string{}
	_, err = entity.NewName(cmd.Name)
	if err != nil {
		args = append(args, "Name")
	}
	_, err = entity.NewURI(cmd.URI)
	if err != nil {
		args = append(args, "URI")
	}
	for _, v := range cmd.Tags {
		_, err = entity.NewTag(v)
		if err != nil {
			args = append(args, "Tags")
			break
		}
	}
	if len(args) > 0 {
		return &InvalidCommandError{Args: args}
	}
	return nil
}
