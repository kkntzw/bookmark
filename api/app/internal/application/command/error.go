package command

import (
	"fmt"
	"sort"
	"strings"
)

// 不正コマンドを表すエラー。
type InvalidCommandError struct {
	Args map[string]error // 不正な引数とエラー一覧
}

// エラー状態を表す。
//
// Args が空の場合は "command is invalid" を出力する。
// Args に値が存在する場合は値を辞書順に連結して "command is invalid: [A: some error, B: some error]" を出力する。
func (e *InvalidCommandError) Error() string {
	text := "command is invalid"
	if len(e.Args) > 0 {
		keys := []string{}
		for k := range e.Args {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		args := []string{}
		for _, k := range keys {
			args = append(args, fmt.Sprintf("%s: %s", k, e.Args[k]))
		}
		text = fmt.Sprintf("%s: [%s]", text, strings.Join(args[:], ", "))
	}
	return text
}
