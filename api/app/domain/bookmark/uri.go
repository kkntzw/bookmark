package bookmark

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
// 空文字(\c09\c0A\c0B\c0C\c0D\c20\c85\cA0) 以外の文字を含まない場合はエラーを返却する。
func validateURI(s string) error {
	if len(s) == 0 {
		return fmt.Errorf("string length is 0")
	}
	for _, r := range s {
		if (r < '\u0009' || '\u000D' < r) && (r != '\u0020') && (r != '\u0085') && (r != '\u00A0') {
			return nil
		}
	}
	return fmt.Errorf("empty string")
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

// string型に変換する。
func (uri *URI) String() string {
	return uri.value.String()
}

// インスタンスを複製する。
func (uri URI) Copy() *URI {
	return &uri
}
