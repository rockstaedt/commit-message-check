package cmd

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestUninstall(t *testing.T) {
	buffer := &bytes.Buffer{}
	log.SetOutput(buffer)

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

	t.Run("returns 0 and", func(t *testing.T) {

		t.Run("removes all occurrences of commit-msg", func(t *testing.T) {
			path := createDirs()

			status := Uninstall(path)

			assert.Equal(t, 0, status)
			assert.NoFileExists(t, fmt.Sprintf("%s/hooks/commit-msg", path))
			assert.FileExists(t, fmt.Sprintf("%s/xyz/commit-msg", path))
		})

		t.Run("logs a success message", func(t *testing.T) {
			_ = Uninstall("")

			assert.Contains(t, buffer.String(), "[SUCCESS]\t Deleted all hook files.")
		})
	})

	t.Run("returns 1 on any error", func(t *testing.T) {
		errPath := t.TempDir()
		err := os.Mkdir(fmt.Sprintf("%s/hooks", errPath), 0000)
		assert.Nil(t, err)

		status := Uninstall(errPath)

		assert.Equal(t, 1, status)
	})
}
