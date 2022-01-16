package command

import (
	"strings"
)

// 不正コマンドを表すエラー。
type InvalidCommandError struct {
	Args []string // 不正な引数一覧
}

// エラー状態を表す。
//
// Args が空の場合は "command is invalid" を出力する。
// Args に値が存在する場合は値を連結して "command is invalid: A, B" を出力する。
func (e *InvalidCommandError) Error() string {
	text := "command is invalid"
	if e.Args != nil && len(e.Args) > 0 {
		text += ": "
		text += strings.Join(e.Args[:], ", ")
	}
	return text
}
