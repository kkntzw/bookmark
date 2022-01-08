package bookmark

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
// U+002D(-) または U+0030-U+0039(0-9) または U+0061-U+007A(a-z) 以外の文字を含む場合はエラーを返却する。
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
