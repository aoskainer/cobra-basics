package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGuessgameCmd(t *testing.T) {
	cases := []struct {
		name       string
		args       []string
		wantOutput string
	}{
		{
			name:       "Test case 1",
			args:       []string{},
			wantOutput: "Expected output for test case 1",
		},
		{
			name:       "Test case 2",
			args:       []string{"arg1", "arg2"},
			wantOutput: "Expected output for test case 2",
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			cmd := guessgameCmd
			cmd.SetArgs(tt.args)
			err := cmd.Execute()
			assert.NoError(t, err)
		})
	}
}
