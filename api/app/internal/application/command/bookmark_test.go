package command

import (
	"testing"

	"github.com/kkntzw/bookmark/internal/domain/entity"
	"github.com/stretchr/testify/assert"
)

func ToErrID(t *testing.T, v string) error {
	t.Helper()
	_, err := entity.NewID(v)
	if err == nil {
		t.Fatal()
	}
	return err
}

func ToErrName(t *testing.T, v string) error {
	t.Helper()
	_, err := entity.NewName(v)
	if err == nil {
		t.Fatal()
	}
	return err
}

func ToErrURI(t *testing.T, v string) error {
	t.Helper()
	_, err := entity.NewURI(v)
	if err == nil {
		t.Fatal()
	}
	return err
}

func ToErrTag(t *testing.T, v string) error {
	t.Helper()
	_, err := entity.NewTag(v)
	if err == nil {
		t.Fatal()
	}
	return err
}

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
		"valid arguments (0 tags)": {
			&RegisterBookmark{"Example", "https://example.com", []string{}},
			nil,
		},
		"valid arguments (1 tag)": {
			&RegisterBookmark{"Example", "https://example.com", []string{"1-A"}},
			nil,
		},
		"valid arguments (2 tags)": {
			&RegisterBookmark{"Example", "https://example.com", []string{"1-A", "1-B"}},
			nil,
		},
		"valid arguments (3 tags)": {
			&RegisterBookmark{"Example", "https://example.com", []string{"1-A", "1-B", "1-C"}},
			nil,
		},
		"invalid name": {
			&RegisterBookmark{"", "https://example.com", []string{"1-A", "1-B", "1-C"}},
			&InvalidCommandError{map[string]error{"Name": ToErrName(t, "")}},
		},
		"invalid uri": {
			&RegisterBookmark{"Example", "", []string{"1-A", "1-B", "1-C"}},
			&InvalidCommandError{map[string]error{"URI": ToErrURI(t, "")}},
		},
		"invalid tags": {
			&RegisterBookmark{"Example", "https://example.com", []string{"1-A", "", "1-C"}},
			&InvalidCommandError{map[string]error{"Tags": ToErrTag(t, "")}},
		},
		"invalid arguments": {
			&RegisterBookmark{"", "", []string{""}},
			&InvalidCommandError{map[string]error{"Name": ToErrName(t, ""), "URI": ToErrURI(t, ""), "Tags": ToErrTag(t, "")}},
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
			&InvalidCommandError{map[string]error{"ID": ToErrID(t, "")}},
		},
		"invalid name": {
			&UpdateBookmark{"1", "", "https://example.com"},
			&InvalidCommandError{map[string]error{"Name": ToErrName(t, "")}},
		},
		"invalid uri": {
			&UpdateBookmark{"1", "Example", ""},
			&InvalidCommandError{map[string]error{"URI": ToErrURI(t, "")}},
		},
		"invalid arguments": {
			&UpdateBookmark{"", "", ""},
			&InvalidCommandError{map[string]error{"ID": ToErrID(t, ""), "Name": ToErrName(t, ""), "URI": ToErrURI(t, "")}},
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
