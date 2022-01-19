package entity

import (
	"fmt"
	"net/url"
)

// URIを表す値オブジェクト。
type URI struct {
	value url.URL
}

// URIを検証する。
//
// RegExp: `^[\s\v\c85\cA0]*$`
//
// 文字列長が0の場合はエラーを返却する。
// 空白(\u0009-\u000D\u0020\u0085\u00A0)以外の文字を含まない場合はエラーを返却する。
func validateURI(s string) error {
	if len(s) == 0 {
		return fmt.Errorf("string length is 0")
	}
	for _, r := range s {
		if (r < '\u0009' || '\u000D' < r) && (r != '\u0020') && (r != '\u0085') && (r != '\u00A0') {
			return nil
		}
	}
	return fmt.Errorf("blank string")
}

// URIを表す値オブジェクトを生成する。
//
// 不正な値を指定した場合はエラーを返却する。
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

// インスタンスを複製する。
func (uri URI) Copy() *URI {
	return &uri
}
