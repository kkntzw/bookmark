package entity

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewID(t *testing.T) {
	t.Parallel()
	cases := map[string]struct {
		v           string
		expectedId  *ID
		expectedErr error
	}{
		"\"\" EMPTY STRING":                   {"", nil, errors.New("string length is 0")},
		"\" \" U+0020 SPACE":                  {"\u0020", nil, errors.New("contains invalid rune: '\u0020' (index: 0)")},
		"\",\" U+002C COMMA":                  {"\u002C", nil, errors.New("contains invalid rune: '\u002C' (index: 0)")},
		"\"-\" U+002D HYPHEN-MINUS":           {"\u002D", &ID{"\u002D"}, nil},
		"\".\" U+002E FULL STOP":              {"\u002E", nil, errors.New("contains invalid rune: '\u002E' (index: 0)")},
		"\"/\" U+002F SOLIDUS":                {"\u002F", nil, errors.New("contains invalid rune: '\u002F' (index: 0)")},
		"\"0\" U+0030 DIGIT ZERO":             {"\u0030", &ID{"\u0030"}, nil},
		"\"9\" U+0039 DIGIT NINE":             {"\u0039", &ID{"\u0039"}, nil},
		"\":\" U+003A COLON":                  {"\u003A", nil, errors.New("contains invalid rune: '\u003A' (index: 0)")},
		"\"@\" U+0040 COMMERCIAL AT":          {"\u0040", nil, errors.New("contains invalid rune: '\u0040' (index: 0)")},
		"\"A\" U+0041 LATIN CAPITAL LETTER A": {"\u0041", &ID{"\u0041"}, nil},
		"\"Z\" U+005A LATIN CAPITAL LETTER Z": {"\u005A", &ID{"\u005A"}, nil},
		"\"[\" U+005B LEFT SQUARE BRACKET":    {"\u005B", nil, errors.New("contains invalid rune: '\u005B' (index: 0)")},
		"\"`\" U+0060 GRAVE ACCENT":           {"\u0060", nil, errors.New("contains invalid rune: '\u0060' (index: 0)")},
		"\"a\" U+0061 LATIN SMALL LETTER A":   {"\u0061", &ID{"\u0061"}, nil},
		"\"z\" U+007A LATIN SMALL LETTER Z":   {"\u007A", &ID{"\u007A"}, nil},
		"\"{\" U+007B LEFT CURLY BRACKET":     {"\u007B", nil, errors.New("contains invalid rune: '\u007B' (index: 0)")},
	}
	for name, tc := range cases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			// when
			actualId, actualErr := NewID(tc.v)
			// then
			assert.Exactly(t, tc.expectedId, actualId)
			assert.Exactly(t, tc.expectedErr, actualErr)
		})
	}
}

func TestID_Equals(t *testing.T) {
	t.Parallel()
	cases := map[string]struct {
		xv            string
		yv            string
		expectedSame  bool
		expectedEquiv bool
	}{
		"equivalent value":     {"00a0c91e6bf6", "00a0c91e6bf6", false, true},
		"non-equivalent value": {"00a0c91e6bf6", "00A0C91E6bF6", false, false},
	}
	for name, tc := range cases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			// given
			x, _ := NewID(tc.xv)
			y, _ := NewID(tc.yv)
			// when
			actualSame := x == y
			actualEquiv := *x == *y
			// then
			assert.Exactly(t, tc.expectedSame, actualSame)
			assert.Exactly(t, tc.expectedEquiv, actualEquiv)
		})
	}
}

func TestID_Value(t *testing.T) {
	t.Parallel()
	// given
	id, _ := NewID("00a0c91e6bf6")
	// when
	actualValue := id.Value()
	// then
	expectedValue := "00a0c91e6bf6"
	assert.Exactly(t, expectedValue, actualValue)
}
