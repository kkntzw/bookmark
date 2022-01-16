package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterBookmarkValidate_正当な場合はnilを返却する(t *testing.T) {
	params := []struct {
		cmd *RegisterBookmark
	}{
		{cmd: &RegisterBookmark{Name: "example", URI: "https://example.com", Tags: nil}},
		{cmd: &RegisterBookmark{Name: "example", URI: "https://example.com", Tags: []string{}}},
		{cmd: &RegisterBookmark{Name: "example", URI: "https://example.com", Tags: []string{"A", "B", "C"}}},
	}
	for _, p := range params {
		// given
		cmd := p.cmd
		// when
		err := cmd.Validate()
		// then
		assert.NoError(t, err)
	}
}

func TestRegisterBookmarkValidate_不正な場合はInvalidCommandErrorを返却する(t *testing.T) {
	params := []struct {
		cmd      *RegisterBookmark
		expected error
	}{
		{
			cmd: &RegisterBookmark{Name: "", URI: "https://example.com", Tags: []string{"A", "B", "C"}},
			expected: &InvalidCommandError{Args: []string{"Name"}},
		},
		{
			cmd: &RegisterBookmark{Name: "example", URI: "", Tags: []string{"A", "B", "C"}},
			expected: &InvalidCommandError{Args: []string{"URI"}},
		},
		{
			cmd: &RegisterBookmark{Name: "example", URI: "https://example.com", Tags: []string{"A", "", "C"}},
			expected: &InvalidCommandError{Args: []string{"Tags"}},
		},
	}
	for _, p := range params {
		// given
		cmd := p.cmd
		// when
		actual := cmd.Validate()
		// then
		expected := p.expected
		assert.Exactly(t, expected, actual)
	}
}
