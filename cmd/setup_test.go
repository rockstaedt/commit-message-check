package cmd

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"rockstaedt/commit-message-check/testdata/mocks"
	"testing"
)

func TestSetup(t *testing.T) {
	buffer := &bytes.Buffer{}
	log.SetOutput(buffer)

	createGitStructure := func() string {
		gitPath := t.TempDir()
		err := os.Mkdir(fmt.Sprintf("%s/hooks", gitPath), os.ModePerm)
		assert.Nil(t, err)

		return gitPath
	}

	t.Run("returns 0 and init hook script", func(t *testing.T) {
		gitPath := createGitStructure()

		status := Setup(gitPath)

		assert.Equal(t, 0, status)
		filePath := fmt.Sprintf("%s/hooks/commit-msg", gitPath)
		assert.FileExists(t, filePath)
		contentBytes, err := os.ReadFile(filePath)
		assert.Nil(t, err)
		assert.Contains(t, string(contentBytes), "commit-message-check validate")
	})

	t.Run("returns 1 when git repo is not initialized and logs it", func(t *testing.T) {
		buffer.Reset()

		status := Setup("/no_existing_git")

		assert.Equal(t, 1, status)
		assert.Contains(t, buffer.String(), "[ERROR]\t No git repository could be found.")
	})

	t.Run("returns 2 when error at creating hook script and logs it", func(t *testing.T) {
		buffer.Reset()

		status := Setup(t.TempDir())

		assert.Equal(t, 2, status)
		assert.Contains(t, buffer.String(), "[ERROR]\t Could not create commit-msg script.")
	})
}

func TestWriteCommitMsgHook(t *testing.T) {
	buffer := &bytes.Buffer{}

	t.Run("marks file as shell script and returns 0", func(t *testing.T) {
		buffer.Reset()

		status := writeCommitMsgHook(buffer)

		assert.Equal(t, 0, status)
		assert.Contains(t, buffer.String(), "#!/bin/sh\n\n")
	})

	t.Run("executes commit-message-check and returns 0", func(t *testing.T) {
		buffer.Reset()

		status := writeCommitMsgHook(buffer)

		assert.Equal(t, 0, status)
		assert.Contains(t, buffer.String(), "./commit-message-check validate $1\n")
	})

	t.Run("logs any error and returns 3", func(t *testing.T) {
		buffer.Reset()
		log.SetOutput(buffer)
		errBuffer := mocks.FakeWriter{}

		status := writeCommitMsgHook(errBuffer)

		assert.Equal(t, 3, status)
		assert.Contains(t, buffer.String(), "[ERROR]\t Could not write to commit-msg script.")
	})
}
