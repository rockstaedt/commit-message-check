package cmd

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestSetup(t *testing.T) {
	buffer := &bytes.Buffer{}
	log.SetOutput(buffer)

	t.Run("returns 0 and init hook script", func(t *testing.T) {
		status := Setup([]string{"src", "internal", ".git"})

		assert.Equal(t, status, 0)
	})

	t.Run("returns 1 if git repo is not initialized and logs it", func(t *testing.T) {
		buffer.Reset()
		status := Setup([]string{"src", "internal"})

		assert.Equal(t, status, 1)
		assert.Contains(t, buffer.String(), "[ERROR]\t No git repository could be found.")
	})
}

func TestHasGitRepository(t *testing.T) {

	t.Run("returns true .git dir could be found", func(t *testing.T) {
		got := hasGitRepository([]string{"tmp", "src", ".git"})

		assert.True(t, got)
	})

	t.Run("return false when no git repository could be found", func(t *testing.T) {
		got := hasGitRepository([]string{"tmp", "src"})

		assert.False(t, got)
	})
}
