package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewID_正当な値を受け取るとID型のインスタンスを返却する(t *testing.T) {
	params := []struct {
		v        string
		expected *ID
	}{
		{v: "-", expected: &ID{"-"}},
		{v: "0", expected: &ID{"0"}},
		{v: "9", expected: &ID{"9"}},
		{v: "a", expected: &ID{"a"}},
		{v: "z", expected: &ID{"z"}},
	}
	for _, p := range params {
		// given
		v := p.v
		// when
		actual, err := NewID(v)
		// then
		assert.Exactly(t, p.expected, actual)
		assert.NoError(t, err)
	}
}

func TestNewID_不正な値を受け取るとエラーを返却する(t *testing.T) {
	params := []struct {
		v         string
		errString string
	}{
		{v: "", errString: "string length is 0"},
		{v: "\u0020", errString: "contains invalid rune: '\u0020' (index: 0)"},
		{v: "\u0029", errString: "contains invalid rune: '\u0029' (index: 0)"},
		{v: "\u0040", errString: "contains invalid rune: '\u0040' (index: 0)"},
		{v: "\u0060", errString: "contains invalid rune: '\u0060' (index: 0)"},
		{v: "\u007B", errString: "contains invalid rune: '\u007B' (index: 0)"},
	}
	for _, p := range params {
		// given
		v := p.v
		// when
		object, err := NewID(v)
		// then
		assert.Nil(t, object)
		assert.EqualError(t, err, p.errString)
	}
}

func TestEquals_同じ値を持つID型のインスタンスは等しい(t *testing.T) {
	// given
	x, _ := NewID("f81d4fae-7dec-11d0-a765-00a0c91e6bf6")
	y, _ := NewID("f81d4fae-7dec-11d0-a765-00a0c91e6bf6")
	// when
	same := x == y
	equiv := *x == *y
	// then
	assert.False(t, same)
	assert.True(t, equiv)
}

func TestEquals_異なる値を持つID型のインスタンスは等しくない(t *testing.T) {
	// given
	x, _ := NewID("f81d4fae-7dec-11d0-a765-00a0c91e6bf6")
	y, _ := NewID("f81d4fae-7dec-11d0-a765-00a0c91e6bf9")
	// when
	same := x == y
	equiv := *x == *y
	// then
	assert.False(t, same)
	assert.False(t, equiv)
}

func TestString_ID型からstring型に変換する(t *testing.T) {
	// given
	id, _ := NewID("f81d4fae-7dec-11d0-a765-00a0c91e6bf6")
	// when
	actual := id.String()
	// then
	expected := "f81d4fae-7dec-11d0-a765-00a0c91e6bf6"
	assert.Exactly(t, expected, actual)
}

func TestCopy_同じ値で異なるポインタを持つID型のインスタンスを返却する(t *testing.T) {
	// given
	id, _ := NewID("f81d4fae-7dec-11d0-a765-00a0c91e6bf6")
	// when
	copy := id.Copy()
	// then
	assert.Exactly(t, id, copy)
	assert.NotSame(t, id, copy)
}
