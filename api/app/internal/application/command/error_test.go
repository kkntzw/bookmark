package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvalidCommandError_エラー状態を表す(t *testing.T) {
	params := []struct {
		err      error
		expected string
	}{
		{err: &InvalidCommandError{}, expected: "command is invalid"},
		{err: &InvalidCommandError{Args: []string{}}, expected: "command is invalid"},
		{err: &InvalidCommandError{Args: []string{"A"}}, expected: "command is invalid: A"},
		{err: &InvalidCommandError{Args: []string{"A", "B"}}, expected: "command is invalid: A, B"},
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
