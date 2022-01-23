package command

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvalidCommandError_エラー状態を表す(t *testing.T) {
	params := []struct {
		err      error
		expected string
	}{
		{
			err:      &InvalidCommandError{},
			expected: "command is invalid",
		},
		{
			err:      &InvalidCommandError{Args: map[string]error{}},
			expected: "command is invalid",
		},
		{
			err:      &InvalidCommandError{Args: map[string]error{"A": fmt.Errorf("some error")}},
			expected: "command is invalid: [A: some error]",
		},
		{
			err:      &InvalidCommandError{Args: map[string]error{"A": fmt.Errorf("some error"), "B": fmt.Errorf("some error")}},
			expected: "command is invalid: [A: some error, B: some error]",
		},
	}
	for _, p := range params {
		// given
		err := p.err
		// when
		actual := err.Error()
		// then
		expected := p.expected
		assert.Exactly(t, expected, actual)
	}
}
