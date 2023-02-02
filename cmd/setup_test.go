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
		err = os.MkdirAll(fmt.Sprintf("%s/modules/my_submodule/hooks", gitPath), os.ModePerm)
		assert.Nil(t, err)
		err = os.MkdirAll(fmt.Sprintf("%s/modules/joined/submodule2/hooks", gitPath), os.ModePerm)
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

	t.Run("returns 0 and init hook scripts also in submodules", func(t *testing.T) {
		gitPath := createGitStructure()

		status := Setup(gitPath)

		assert.Equal(t, 0, status)
		assert.FileExists(t, fmt.Sprintf("%s/modules/my_submodule/hooks/commit-msg", gitPath))
		assert.FileExists(t, fmt.Sprintf("%s/modules/joined/submodule2/hooks/commit-msg", gitPath))
	})

	t.Run("returns 1 when git repo is not initialized and logs it", func(t *testing.T) {
		buffer.Reset()

		status := Setup("/no_existing_git")

		assert.Equal(t, 1, status)
		assert.Contains(t, buffer.String(), "[ERROR]\t No git repository could be found.")
	})

	t.Run("returns 2 when error at creating hook script and logs it", func(t *testing.T) {
		buffer.Reset()
		errPath := t.TempDir()
		err := os.Mkdir(fmt.Sprintf("%s/not_readable", errPath), 0000)
		assert.Nil(t, err)
		err = os.Mkdir(fmt.Sprintf("%s/hooks", errPath), 0000)
		assert.Nil(t, err)

		status := Setup(errPath)

		assert.Equal(t, 2, status)
		assert.Contains(t, buffer.String(), "[ERROR]\t Could not create commit-msg script.")
	})

	t.Run("returns 2 when .git not readable", func(t *testing.T) {
		errPath := t.TempDir()
		err := os.Chmod(errPath, 0000)
		assert.Nil(t, err)

		status := Setup(errPath)

		assert.Equal(t, 2, status)
	})
}

func TestWriteCommitMsgHook(t *testing.T) {
	buffer := &bytes.Buffer{}

	t.Run("marks file as shell script and returns 0", func(t *testing.T) {
		buffer.Reset()

		writeCommitMsgHook(buffer)

		assert.Contains(t, buffer.String(), "#!/bin/sh\n\n")
	})

	t.Run("executes commit-message-check", func(t *testing.T) {
		buffer.Reset()
		
		writeCommitMsgHook(buffer)

		assert.Contains(t, buffer.String(), "./commit-message-check validate $1\n")
	})

	t.Run("logs any error", func(t *testing.T) {
		log.SetOutput(buffer)
		errBuffer := mocks.FakeWriter{}

		writeCommitMsgHook(errBuffer)

		assert.Contains(t, buffer.String(), "[ERROR]\t Could not write commit-msg script: error at writing")
	})
}
