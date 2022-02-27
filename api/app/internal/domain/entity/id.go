package entity

import (
	"fmt"
)

// IDを表す値オブジェクト。
type ID struct {
	value string
}

// IDを検証する。
//
// RegExp: `^[-0-9A-Za-z]+$`
func validateID(s string) error {
	if len(s) == 0 {
		return fmt.Errorf("string length is 0")
	}
	for i, r := range s {
		if (r != '-') && (r < '0' || '9' < r) && (r < 'A' || 'Z' < r) && (r < 'a' || 'z' < r) {
			return fmt.Errorf("contains invalid rune: '%c' (index: %d)", r, i)
		}
	}
	return nil
}

// IDを表す値オブジェクトを生成する。
//
// 文字列長が0の場合はエラーを返却する。
// ハイフン、半角英数字以外の文字を含む場合はエラーを返却する。
func NewID(v string) (*ID, error) {
	if err := validateID(v); err != nil {
		return nil, err
	}
	return &ID{v}, nil
}

// 値を取得する。
func (id *ID) Value() string {
	return id.value
}
