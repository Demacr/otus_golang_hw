package main

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	t.Run("Ignore file containing = in name", func(t *testing.T) {
		tempDir, err := ioutil.TempDir(".", "testdir_")
		require.NoError(t, err)
		defer os.RemoveAll(tempDir)

		_, err = ioutil.TempFile(tempDir, "test=ignored")
		require.NoError(t, err)

		env, err := ReadDir(tempDir)
		require.NoError(t, err)
		require.Empty(t, env)
	})
}
