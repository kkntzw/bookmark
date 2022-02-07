package command

import (
	"testing"

	"github.com/kkntzw/bookmark/internal/domain/entity"
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
	_, errName := entity.NewName("")
	_, errUri := entity.NewURI("")
	_, errTag := entity.NewTag("")
	params := []struct {
		cmd      *RegisterBookmark
		expected error
	}{
		{
			cmd:      &RegisterBookmark{Name: "", URI: "https://example.com", Tags: []string{"A", "B", "C"}},
			expected: &InvalidCommandError{Args: map[string]error{"Name": errName}},
		},
		{
			cmd:      &RegisterBookmark{Name: "example", URI: "", Tags: []string{"A", "B", "C"}},
			expected: &InvalidCommandError{Args: map[string]error{"URI": errUri}},
		},
		{
			cmd:      &RegisterBookmark{Name: "example", URI: "https://example.com", Tags: []string{"A", "", "C"}},
			expected: &InvalidCommandError{Args: map[string]error{"Tags": errTag}},
		},
		{
			cmd:      &RegisterBookmark{Name: "", URI: "", Tags: []string{""}},
			expected: &InvalidCommandError{Args: map[string]error{"Name": errName, "URI": errUri, "Tags": errTag}},
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

func TestUpdateBookmarkValidate_正当な場合はnilを返却する(t *testing.T) {
	// given
	cmd := &UpdateBookmark{
		ID:   "f81d4fae-7dec-11d0-a765-00a0c91e6bf6",
		Name: "EXAMPLE",
		URI:  "http://example.com",
	}
	// when
	err := cmd.Validate()
	// then
	assert.NoError(t, err)
}

func TestUpdateBookmarkValidate_不正な場合はInvalidCommandErrorを返却する(t *testing.T) {
	_, errId := entity.NewID("")
	_, errName := entity.NewName("")
	_, errUri := entity.NewURI("")
	params := []struct {
		cmd      *UpdateBookmark
		expected error
	}{
		{
			cmd:      &UpdateBookmark{ID: "", Name: "EXAMPLE", URI: "http://example.com"},
			expected: &InvalidCommandError{Args: map[string]error{"ID": errId}},
		},
		{
			cmd:      &UpdateBookmark{ID: "f81d4fae-7dec-11d0-a765-00a0c91e6bf6", Name: "", URI: "http://example.com"},
			expected: &InvalidCommandError{Args: map[string]error{"Name": errName}},
		},
		{
			cmd:      &UpdateBookmark{ID: "f81d4fae-7dec-11d0-a765-00a0c91e6bf6", Name: "EXAMPLE", URI: ""},
			expected: &InvalidCommandError{Args: map[string]error{"URI": errUri}},
		},
		{
			cmd:      &UpdateBookmark{ID: "", Name: "", URI: ""},
			expected: &InvalidCommandError{Args: map[string]error{"ID": errId, "Name": errName, "URI": errUri}},
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
