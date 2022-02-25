package entity

import (
	"fmt"
	"unicode"
)

// ブックマーク名を表す値オブジェクト。
type Name struct {
	value string
}

// ブックマーク名を検証する。
func validateName(s string) error {
	if len(s) == 0 {
		return fmt.Errorf("string length is 0")
	}
	blank := true
	for i, r := range s {
		if unicode.IsControl(r) {
			return fmt.Errorf("contains control character: %U (index: %d)", r, i)
		}
		if !unicode.IsSpace(r) {
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
// 文字列長が0の場合はエラーを返却する。
// 制御文字を含む場合はエラーを返却する。
// 空白以外の文字を含まない場合はエラーを返却する。
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
