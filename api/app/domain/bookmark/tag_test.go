package bookmark

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTag_正当な値を受け取るとTag型のインスタンスを返却する(t *testing.T) {
	params := []struct {
		v        string
		expected *Tag
	}{
		{v: "EXAMPLE 0", expected: &Tag{"EXAMPLE 0"}},
		{v: "example 9", expected: &Tag{"example 9"}},
		{v: "れい　０", expected: &Tag{"れい　０"}},
		{v: "レイ　９", expected: &Tag{"レイ　９"}},
		{v: "例", expected: &Tag{"例"}},
	}
	for _, p := range params {
		// given
		v := p.v
		// when
		actual, err := NewTag(v)
		// then
		assert.Exactly(t, p.expected, actual)
		assert.NoError(t, err)
	}
}

func TestNewTag_不正な値を受け取るとエラーを返却する(t *testing.T) {
	params := []struct {
		v         string
		errString string
	}{
		{v: "", errString: "string length is 0"},
		{v: "\u0000", errString: "contains control character: U+0000 (index: 0)"},
		{v: "\u001F", errString: "contains control character: U+001F (index: 0)"},
		{v: "\u007F", errString: "contains control character: U+007F (index: 0)"},
		{v: "\u0020\u0085\u00A0", errString: "blank string"},
	}
	for _, p := range params {
		// given
		v := p.v
		// when
		object, err := NewTag(v)
		// then
		assert.Nil(t, object)
		assert.EqualError(t, err, p.errString)
	}
}

func TestEquals_同じ値を持つTag型のインスタンスは等しい(t *testing.T) {
	// given
	x, _ := NewTag("example")
	y, _ := NewTag("example")
	// when
	same := x == y
	equiv := *x == *y
	// then
	assert.False(t, same)
	assert.True(t, equiv)
}

func TestEquals_異なる値を持つTag型のインスタンスは等しくない(t *testing.T) {
	// given
	x, _ := NewTag("example")
	y, _ := NewTag("Example")
	// when
	same := x == y
	equiv := *x == *y
	// then
	assert.False(t, same)
	assert.False(t, equiv)
}

func TestString_Tag型からstring型に変換する(t *testing.T) {
	// given
	tag, _ := NewTag("example")
	// when
	actual := tag.String()
	// then
	expected := "example"
	assert.Exactly(t, expected, actual)
}

func TestCopy_同じ値を持つ異なるポインタのTag型のインスタンスを返却する(t *testing.T) {
	// given
	tag, _ := NewTag("example")
	// when
	copy := tag.Copy()
	// then
	assert.Exactly(t, tag, copy)
	assert.NotSame(t, tag, copy)
}
