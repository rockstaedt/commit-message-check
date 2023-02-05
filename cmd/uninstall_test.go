package cmd

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestUninstall(t *testing.T) {

	createDirs := func() string {
		path := t.TempDir()
		err := os.Mkdir(fmt.Sprintf("%s/hooks", path), os.ModePerm)
		assert.Nil(t, err)
		_, err = os.Create(fmt.Sprintf("%s/hooks/commit-msg", path))
		assert.Nil(t, err)
		err = os.Mkdir(fmt.Sprintf("%s/xyz", path), os.ModePerm)
		assert.Nil(t, err)
		_, err = os.Create(fmt.Sprintf("%s/xyz/commit-msg", path))
		assert.Nil(t, err)

		return path
	}

	t.Run("returns 0 and removes all occurrences of commit-msg", func(t *testing.T) {
		path := createDirs()

		status := Uninstall(path)

		assert.Equal(t, 0, status)
		assert.NoFileExists(t, fmt.Sprintf("%s/hooks/commit-msg", path))
		assert.FileExists(t, fmt.Sprintf("%s/xyz/commit-msg", path))
	})

	t.Run("returns 1 on any error", func(t *testing.T) {
		t.Skip()
	})
}
