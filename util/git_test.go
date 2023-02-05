package util

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestIsGitRepo(t *testing.T) {

	t.Run("returns true if .git could be found", func(t *testing.T) {
		path := t.TempDir()
		err := os.Mkdir(fmt.Sprintf("%s/.git", path), os.ModePerm)
		assert.Nil(t, err)

		check := IsGitRepo(path)

		assert.True(t, check)
	})

	t.Run("returns false if .git could not be found", func(t *testing.T) {
		check := IsGitRepo(t.TempDir())

		assert.False(t, check)
	})
}
