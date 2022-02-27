package entity

import (
	"fmt"
	"net/url"
	"unicode"
)

// URIを表す値オブジェクト。
type URI struct {
	value url.URL
}

// URIを検証する。
func validateURI(s string) error {
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

// URIを表す値オブジェクトを生成する。
//
// 文字列長が0の場合はエラーを返却する。
// 制御文字を含む場合はエラーを返却する。
// 空白以外の文字を含まない場合はエラーを返却する。
// URIの解析に失敗した場合はエラーを返却する。
func NewURI(v string) (*URI, error) {
	if err := validateURI(v); err != nil {
		return nil, err
	}
	u, err := url.Parse(v)
	if err != nil {
		return nil, fmt.Errorf("invalid format: %s", v)
	}
	return &URI{*u}, nil
}

// 値を取得する。
func (uri *URI) Value() url.URL {
	return uri.value
}

// string型の値を取得する。
func (uri *URI) String() string {
	return uri.value.String()
}
