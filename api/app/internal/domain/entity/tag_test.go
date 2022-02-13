package entity

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTag(t *testing.T) {
	t.Parallel()
	cases := map[string]struct {
		v           string
		expectedTag *Tag
		expectedErr error
	}{
		"non-empty string":           {"Hello, 世界", &Tag{"Hello, 世界"}, nil},
		"empty string":               {"", nil, errors.New("string length is 0")},
		"contains control character": {"Hello,\u0000世界", nil, errors.New("contains control character: U+0000 (index: 6)")},
		"blank string":               {" ", nil, errors.New("blank string")},
	}
	for name, tc := range cases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			// when
			actualTag, actualErr := NewTag(tc.v)
			// then
			assert.Exactly(t, tc.expectedTag, actualTag)
			assert.Exactly(t, tc.expectedErr, actualErr)
		})
	}
}

func TestTag_Equals(t *testing.T) {
	t.Parallel()
	cases := map[string]struct {
		xv            string
		yv            string
		expectedSame  bool
		expectedEquiv bool
	}{
		"equivalent value":     {"Hello, 世界", "Hello, 世界", false, true},
		"non-equivalent value": {"Hello, 世界", "Hello, World", false, false},
	}
	for name, tc := range cases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			// given
			x, _ := NewTag(tc.xv)
			y, _ := NewTag(tc.yv)
			// when
			actualSame := x == y
			actualEquiv := *x == *y
			// then
			assert.Exactly(t, tc.expectedSame, actualSame)
			assert.Exactly(t, tc.expectedEquiv, actualEquiv)
		})
	}
}

func TestTag_Value(t *testing.T) {
	t.Parallel()
	// given
	tag, _ := NewTag("Hello, 世界")
	// when
	actualValue := tag.Value()
	// then
	expectedValue := "Hello, 世界"
	assert.Exactly(t, expectedValue, actualValue)
}
