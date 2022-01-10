package bookmark

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewName_正当な値を受け取るとName型のインスタンスを返却する(t *testing.T) {
	params := []struct {
		v        string
		expected *Name
	}{
		{v: "EXAMPLE 0", expected: &Name{"EXAMPLE 0"}},
		{v: "example 9", expected: &Name{"example 9"}},
		{v: "れい　０", expected: &Name{"れい　０"}},
		{v: "レイ　９", expected: &Name{"レイ　９"}},
		{v: "例", expected: &Name{"例"}},
	}
	for _, p := range params {
		// given
		v := p.v
		// when
		actual, err := NewName(v)
		// then
		assert.Exactly(t, p.expected, actual)
		assert.NoError(t, err)
	}
}

func TestNewName_不正な値を受け取るとエラーを返却する(t *testing.T) {
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
		object, err := NewName(v)
		// then
		assert.Nil(t, object)
		assert.EqualError(t, err, p.errString)
	}
}

func TestEquals_同じ値を持つName型のインスタンスは等しい(t *testing.T) {
	// given
	x, _ := NewName("example")
	y, _ := NewName("example")
	// when
	same := x == y
	equiv := *x == *y
	// then
	assert.False(t, same)
	assert.True(t, equiv)
}

func TestEquals_異なる値を持つName型のインスタンスは等しくない(t *testing.T) {
	// given
	x, _ := NewName("example")
	y, _ := NewName("Example")
	// when
	same := x == y
	equiv := *x == *y
	// then
	assert.False(t, same)
	assert.False(t, equiv)
}

func TestString_Name型からstring型に変換する(t *testing.T) {
	// given
	name, _ := NewName("example")
	// when
	actual := name.String()
	// then
	expected := "example"
	assert.Exactly(t, expected, actual)
}

func TestCopy_同じ値で異なるポインタを持つName型のインスタンスを返却する(t *testing.T) {
	// given
	name, _ := NewName("example")
	// when
	copy := name.Copy()
	// then
	assert.Exactly(t, name, copy)
	assert.NotSame(t, name, copy)
}
