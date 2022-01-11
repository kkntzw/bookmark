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
// RegExp: `^[-0-9a-z]+$`
//
// 文字列長が0の場合はエラーを返却する。
// ハイフン、半角数字、半角英小文字(\u002D\u0030-\u0039\u0061-\u007A)以外の文字を含む場合はエラーを返却する。
func validateID(s string) error {
	if len(s) == 0 {
		return fmt.Errorf("string length is 0")
	}
	for i, r := range s {
		if (r != '-') && (r < '0' || '9' < r) && (r < 'a' || 'z' < r) {
			return fmt.Errorf("contains invalid rune: '%c' (index: %d)", r, i)
		}
	}
	return nil
}

// IDを表す値オブジェクトを生成する。
//
// 不正な値を指定した場合はエラーを返却する。
func NewID(v string) (*ID, error) {
	if err := validateID(v); err != nil {
		return nil, err
	}
	return &ID{v}, nil
}

// string型に変換する。
func (id *ID) String() string {
	return id.value
}

// インスタンスを複製する。
func (id ID) Copy() *ID {
	return &id
}
