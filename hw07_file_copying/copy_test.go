package main

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	var newFile *os.File
	newFilePath := "out.txt"
	testFilePath := "testdata/input.txt"
	testFile, err := os.OpenFile(testFilePath, os.O_RDONLY, os.FileMode(0o755))
	require.NoError(t, err)
	testFileInfo, err := testFile.Stat()
	require.NoError(t, err)

	defer func() {
		newFile, err = os.OpenFile(newFilePath, os.O_RDONLY, os.FileMode(0o755))
		if err == nil {
			_ = os.Remove(newFile.Name())
		}
	}()
	t.Run("offset boundary values", func(t *testing.T) {
		err = Copy(testFilePath, newFilePath, testFileInfo.Size(), 0)
		require.NoError(t, err)

		newFile, err = os.OpenFile(newFilePath, os.O_RDONLY, os.FileMode(0o755))
		newFileInfo, err := newFile.Stat()
		require.NoError(t, err)
		require.Equal(t, newFileInfo.Size(), int64(0))

		err = os.Remove(newFilePath)
		require.NoError(t, err)

		err = Copy(testFilePath, newFilePath, testFileInfo.Size()-1, 2)
		require.NoError(t, err)

		newFile, err = os.OpenFile(newFilePath, os.O_RDONLY, os.FileMode(0o755))
		require.NoError(t, err)
		newFileInfo, err = newFile.Stat()
		require.NoError(t, err)
		require.Equal(t, newFileInfo.Size(), int64(1))

		err = os.Remove(newFilePath)
		require.NoError(t, err)
	})

	t.Run("invalid fromPath", func(t *testing.T) {
		err = Copy("testdata", newFilePath, 0, 0)
		require.Truef(t, errors.Is(err, ErrFileNotFound), "actual err - %v", err)

		err = Copy("dsaa/a", newFilePath, 0, 0)
		require.Truef(t, errors.Is(err, ErrFileNotFound), "actual err - %v", err)
	})
}
