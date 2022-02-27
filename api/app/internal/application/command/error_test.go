package command

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvalidCommandError_Error(t *testing.T) {
	t.Parallel()
	cases := map[string]struct {
		args              map[string]error
		expectedErrString string
	}{
		"nil args": {
			nil,
			"command is invalid",
		},
		"empty args": {
			map[string]error{},
			"command is invalid",
		},
		"1 arg": {
			map[string]error{"A": errors.New("some error")},
			"command is invalid: [A: some error]",
		},
		"2 args": {
			map[string]error{"A": errors.New("some error"), "B": errors.New("some error")},
			"command is invalid: [A: some error, B: some error]",
		},
	}
	for name, tc := range cases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			// given
			err := &InvalidCommandError{tc.args}
			// when
			actualErrString := err.Error()
			// then
			assert.Exactly(t, tc.expectedErrString, actualErrString)
		})
	}
}
