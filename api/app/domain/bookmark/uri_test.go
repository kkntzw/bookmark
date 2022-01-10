package bookmark

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewURI_正当な値を受け取るとURI型のインスタンスを返却する(t *testing.T) {
	params := []struct {
		v        string
		expected *URI
	}{
		{
			v: "https://example.com:443/foo?q=bar#baz",
			expected: &URI{url.URL{
				Scheme:   "https",
				Host:     "example.com:443",
				Path:     "/foo",
				RawQuery: "q=bar",
				Fragment: "baz",
			}},
		},
		{
			v: "tel:+1-201-555-0123",
			expected: &URI{url.URL{
				Scheme: "tel",
				Opaque: "+1-201-555-0123",
			}},
		},
	}
	for _, p := range params {
		// given
		v := p.v
		// when
		actual, err := NewURI(v)
		// then
		assert.Exactly(t, p.expected, actual)
		assert.NoError(t, err)
	}
}

func TestNewURI_不正な値を受け取るとエラーを返却する(t *testing.T) {
	params := []struct {
		v         string
		errString string
	}{
		{v: "", errString: "string length is 0"},
		{v: "\u0009\u000A\u000B\u000C\u000D\u0020\u0085\u00A0", errString: "blank string"},
		{v: "\u001F", errString: "invalid format: \u001F"},
		{v: "\u007F", errString: "invalid format: \u007F"},
	}
	for _, p := range params {
		// given
		v := p.v
		// when
		object, err := NewURI(v)
		// then
		assert.Nil(t, object)
		assert.EqualError(t, err, p.errString)
	}
}

func TestEquals_同じ値を持つURI型のインスタンスは等しい(t *testing.T) {
	// given
	x, _ := NewURI("https://example.com")
	y, _ := NewURI("https://example.com")
	// when
	same := x == y
	equiv := *x == *y
	// then
	assert.False(t, same)
	assert.True(t, equiv)
}

func TestEquals_異なる値を持つURI型のインスタンスは等しくない(t *testing.T) {
	// given
	x, _ := NewURI("https://example.com")
	y, _ := NewURI("http://example.com")
	// when
	same := x == y
	equiv := *x == *y
	// then
	assert.False(t, same)
	assert.False(t, equiv)
}

func TestString_URI型からstring型に変換する(t *testing.T) {
	// given
	uri, _ := NewURI("https://example.com")
	// when
	actual := uri.String()
	// then
	expected := "https://example.com"
	assert.Exactly(t, expected, actual)
}

func TestCopy_同じ値で異なるポインタを持つURI型のインスタンスを返却する(t *testing.T) {
	// given
	uri, _ := NewURI("https://example.com")
	// when
	copy := uri.Copy()
	// then
	assert.Exactly(t, uri, copy)
	assert.NotSame(t, uri, copy)
}
