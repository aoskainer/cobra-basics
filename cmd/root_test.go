package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRootCmd(t *testing.T) {
	cmd := rootCmd
	args := []string{"--help"}

	cmd.SetArgs(args)
	err := cmd.Execute()

	assert.NoError(t, err)
}
