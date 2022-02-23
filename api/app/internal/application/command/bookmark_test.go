package command

import (
	"testing"

	"github.com/kkntzw/bookmark/test/helper"
	"github.com/stretchr/testify/assert"
)

func TestRegisterBookmark_Validate(t *testing.T) {
	t.Parallel()
	cases := map[string]struct {
		cmd         *RegisterBookmark
		expectedErr error
	}{
		"valid arguments (nil tags)": {
			&RegisterBookmark{"Example", "https://example.com", nil},
			nil,
		},
		"valid arguments (empty tags)": {
			&RegisterBookmark{"Example", "https://example.com", []string{}},
			nil,
		},
		"valid arguments (1 tag)": {
			&RegisterBookmark{"Example", "https://example.com", []string{"foo"}},
			nil,
		},
		"valid arguments (2 tags)": {
			&RegisterBookmark{"Example", "https://example.com", []string{"foo", "bar"}},
			nil,
		},
		"valid arguments (3 tags)": {
			&RegisterBookmark{"Example", "https://example.com", []string{"foo", "bar", "baz"}},
			nil,
		},
		"invalid name": {
			&RegisterBookmark{"", "https://example.com", []string{"foo", "bar", "baz"}},
			&InvalidCommandError{map[string]error{"Name": helper.ToErrName(t, "")}},
		},
		"invalid uri": {
			&RegisterBookmark{"Example", "", []string{"foo", "bar", "baz"}},
			&InvalidCommandError{map[string]error{"URI": helper.ToErrURI(t, "")}},
		},
		"invalid tags": {
			&RegisterBookmark{"Example", "https://example.com", []string{"foo", "", "baz"}},
			&InvalidCommandError{map[string]error{"Tags": helper.ToErrTag(t, "")}},
		},
		"invalid arguments": {
			&RegisterBookmark{"", "", []string{""}},
			&InvalidCommandError{map[string]error{"Name": helper.ToErrName(t, ""), "URI": helper.ToErrURI(t, ""), "Tags": helper.ToErrTag(t, "")}},
		},
	}
	for name, tc := range cases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			// when
			actualErr := tc.cmd.Validate()
			// then
			assert.Exactly(t, tc.expectedErr, actualErr)
		})
	}
}

func TestUpdateBookmark_Validate(t *testing.T) {
	t.Parallel()
	cases := map[string]struct {
		cmd         *UpdateBookmark
		expectedErr error
	}{
		"valid arguments": {
			&UpdateBookmark{"1", "Example", "https://example.com"},
			nil,
		},
		"invalid id": {
			&UpdateBookmark{"", "Example", "https://example.com"},
			&InvalidCommandError{map[string]error{"ID": helper.ToErrID(t, "")}},
		},
		"invalid name": {
			&UpdateBookmark{"1", "", "https://example.com"},
			&InvalidCommandError{map[string]error{"Name": helper.ToErrName(t, "")}},
		},
		"invalid uri": {
			&UpdateBookmark{"1", "Example", ""},
			&InvalidCommandError{map[string]error{"URI": helper.ToErrURI(t, "")}},
		},
		"invalid arguments": {
			&UpdateBookmark{"", "", ""},
			&InvalidCommandError{map[string]error{"ID": helper.ToErrID(t, ""), "Name": helper.ToErrName(t, ""), "URI": helper.ToErrURI(t, "")}},
		},
	}
	for name, tc := range cases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			// when
			actualErr := tc.cmd.Validate()
			// then
			assert.Exactly(t, tc.expectedErr, actualErr)
		})
	}
}
