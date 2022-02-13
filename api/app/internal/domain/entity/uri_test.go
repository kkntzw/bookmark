package entity

import (
	"errors"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewURI(t *testing.T) {
	t.Parallel()
	cases := map[string]struct {
		v           string
		expectedUri *URI
		expectedErr error
	}{
		"valid format": {
			"https://example.com:443/foo?q=bar#baz",
			&URI{url.URL{Scheme: "https", Host: "example.com:443", Path: "/foo", RawQuery: "q=bar", Fragment: "baz"}},
			nil,
		},
		"invalid format": {
			"://example.com",
			nil,
			errors.New("invalid format: ://example.com"),
		},
		"empty string": {
			"",
			nil,
			errors.New("string length is 0"),
		},
		"contains control character": {
			"https://example.com\u0000",
			nil,
			errors.New("contains control character: U+0000 (index: 19)"),
		},
		"blank string": {
			" ",
			nil,
			errors.New("blank string"),
		},
	}
	for name, tc := range cases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			// when
			actualUri, actualErr := NewURI(tc.v)
			// then
			assert.Exactly(t, tc.expectedUri, actualUri)
			assert.Exactly(t, tc.expectedErr, actualErr)
		})
	}
}

func TestURI_Equals(t *testing.T) {
	t.Parallel()
	cases := map[string]struct {
		xv            string
		yv            string
		expectedSame  bool
		expectedEquiv bool
	}{
		"equivalent value":     {"https://example.com", "https://example.com", false, true},
		"non-equivalent value": {"https://example.com", "http://example.com", false, false},
	}
	for name, tc := range cases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			// given
			x, _ := NewURI(tc.xv)
			y, _ := NewURI(tc.yv)
			// when
			actualSame := x == y
			actualEquiv := *x == *y
			// then
			assert.Exactly(t, tc.expectedSame, actualSame)
			assert.Exactly(t, tc.expectedEquiv, actualEquiv)
		})
	}
}

func TestURI_Value(t *testing.T) {
	t.Parallel()
	// given
	uri, _ := NewURI("https://example.com")
	// when
	actualValue := uri.Value()
	// then
	expectedValue := url.URL{Scheme: "https", Host: "example.com"}
	assert.Exactly(t, expectedValue, actualValue)
}

func TestURI_String(t *testing.T) {
	t.Parallel()
	// given
	uri, _ := NewURI("https://example.com")
	// when
	actualString := uri.String()
	// then
	expectedString := "https://example.com"
	assert.Exactly(t, expectedString, actualString)
}
