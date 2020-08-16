package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	t.Run("Test valid return code", func(t *testing.T) {
		rc := RunCmd([]string{"bash", "-c", "exit 5"}, Environment{})
		require.Equal(t, 5, rc)
	})
}
