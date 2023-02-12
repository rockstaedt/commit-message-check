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

	t.Run("returns 0 and", func(t *testing.T) {

		createDirs := func() string {
			path := t.TempDir()
			err := os.Mkdir(fmt.Sprintf("%s/hooks", path), os.ModePerm)
			assert.Nil(t, err)

			return path
		}

		t.Run("creates commit-msg script in hook folder", func(t *testing.T) {
			path := createDirs()

			status := Setup(path)

			assert.Equal(t, 0, status)
			assert.FileExists(t, fmt.Sprintf("%s/hooks/commit-msg", path))
		})

		t.Run("logs a success message", func(t *testing.T) {
			buffer.Reset()
			path := createDirs()

			_ = Setup(path)

			assert.Contains(t, buffer.String(), "[SUCCESS]\t commit-message-check successfully installed.")
		})
	})

	t.Run("returns 1 when error at walking hooks and logs it", func(t *testing.T) {
		errPath := t.TempDir()
		err := os.Mkdir(fmt.Sprintf("%s/hooks", errPath), 0000)
		assert.Nil(t, err)

		status := Setup(errPath)

		assert.Equal(t, 1, status)
		assert.Contains(t, buffer.String(), "[ERROR]\t Could not create commit-msg script.")
	})
}
