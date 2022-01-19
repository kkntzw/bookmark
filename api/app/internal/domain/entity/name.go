package entity

import "fmt"

// ブックマーク名を表す値オブジェクト。
type Name struct {
	value string
}

// ブックマーク名を検証する。
//
// 文字列長が0の場合はエラーを返却する。
// 制御文字(\u0000-\u001F\u007F)を含む場合はエラーを返却する。
// 空白(\u0020\u0085\u00A0)以外の文字を含まない場合はエラーを返却する。
func validateName(s string) error {
	if len(s) == 0 {
		return fmt.Errorf("string length is 0")
	}
	blank := true
	for i, r := range s {
		if ('\u0000' <= r && r <= '\u001F') || (r == '\u007F') {
			return fmt.Errorf("contains control character: %U (index: %d)", r, i)
		}
		if (r != '\u0020') && (r != '\u0085') && (r != '\u00A0') {
			blank = false
		}
	}
	if blank {
		return fmt.Errorf("blank string")
	}
	return nil
}

// ブックマーク名を表す値オブジェクトを生成する。
//
// 不正な値を指定した場合はエラーを返却する。
func NewName(v string) (*Name, error) {
	if err := validateName(v); err != nil {
		return nil, err
	}
	return &Name{v}, nil
}

// 値を取得する。
func (name *Name) Value() string {
	return name.value
}

// インスタンスを複製する。
func (name Name) Copy() *Name {
	return &name
}
