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
		gitPath := t.TempDir()
		status := Setup(gitPath)

		assert.Equal(t, status, 0)
	})

	t.Run("returns 1 if git repo is not initialized and logs it", func(t *testing.T) {
		buffer.Reset()
		status := Setup("/no_existing_git")

		assert.Equal(t, status, 1)
		assert.Contains(t, buffer.String(), "[ERROR]\t No git repository could be found.")
	})
}
