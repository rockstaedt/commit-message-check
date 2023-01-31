package cmd

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestSetup(t *testing.T) {
	buffer := &bytes.Buffer{}
	log.SetOutput(buffer)

	t.Run("returns 0 and init hook script", func(t *testing.T) {
		gitPath := t.TempDir()
		hookPath := fmt.Sprintf("%s/hooks", gitPath)
		err := os.Mkdir(hookPath, os.ModePerm)
		assert.Nil(t, err)

		status := Setup(gitPath)

		assert.Equal(t, 0, status)
		filePath := fmt.Sprintf("%s/commit-msg", hookPath)
		assert.FileExists(t, filePath)
		contentBytes, err := os.ReadFile(filePath)
		assert.Nil(t, err)
		assert.Contains(t, string(contentBytes), "commit-message-check validate")
	})

	t.Run("returns 1 if git repo is not initialized and logs it", func(t *testing.T) {
		buffer.Reset()

		status := Setup("/no_existing_git")

		assert.Equal(t, 1, status)
		assert.Contains(t, buffer.String(), "[ERROR]\t No git repository could be found.")
	})
}

func TestWriteCommitMsgHook(t *testing.T) {
	buffer := &bytes.Buffer{}

	t.Run("marks file as shell script", func(t *testing.T) {
		buffer.Reset()

		_ = writeCommitMsgHook(buffer)

		assert.Contains(t, buffer.String(), "#!/bin/sh\n\n")
	})

	t.Run("executes commit-message-check", func(t *testing.T) {
		buffer.Reset()

		_ = writeCommitMsgHook(buffer)

		assert.Contains(t, buffer.String(), "./commit-message-check validate $1")
	})
}
