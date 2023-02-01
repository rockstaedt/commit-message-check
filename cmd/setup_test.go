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

	gitPath := t.TempDir()
	hookPath := fmt.Sprintf("%s/hooks", gitPath)

	setup := func() {
		err := os.Mkdir(hookPath, os.ModePerm)
		assert.Nil(t, err)
	}

	t.Run("returns 0 and init hook script", func(t *testing.T) {
		setup()

		status := Setup(gitPath)

		assert.Equal(t, 0, status)
		filePath := fmt.Sprintf("%s/commit-msg", hookPath)
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

	t.Run("returns 3 when error at writing hook script and logs it", func(t *testing.T) {
		t.Skip()
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

		assert.Contains(t, buffer.String(), "./commit-message-check validate $1\n")
	})

	t.Run("returns any error", func(t *testing.T) {
		errBuffer := mocks.FakeWriter{}

		err := writeCommitMsgHook(errBuffer)

		assert.NotNil(t, err)
	})
}
