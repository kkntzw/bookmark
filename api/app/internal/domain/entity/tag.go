package entity

import (
	"fmt"
	"unicode"
)

// タグを表す値オブジェクト。
type Tag struct {
	value string
}

// タグを検証する。
//
// 文字列長が0の場合はエラーを返却する。
// 制御文字を含む場合はエラーを返却する。
// 空白以外の文字を含まない場合はエラーを返却する。
func validateTag(s string) error {
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

// タグを表す値オブジェクトを生成する。
//
// 不正な値を指定した場合はエラーを返却する。
func NewTag(v string) (*Tag, error) {
	if err := validateTag(v); err != nil {
		return nil, err
	}
	return &Tag{v}, nil
}

// 値を取得する。
func (tag *Tag) Value() string {
	return tag.value
}
