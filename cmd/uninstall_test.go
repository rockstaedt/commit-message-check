package cmd

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"rockstaedt/commit-message-check/internal/model"
	"testing"
)

func TestUninstall(t *testing.T) {
	buffer := &bytes.Buffer{}
	log.SetOutput(buffer)

	createFakeHandlerWithDirs := func() *Handler {
		path := t.TempDir()
		err := os.Mkdir(fmt.Sprintf("%s/hooks", path), os.ModePerm)
		assert.Nil(t, err)
		_, err = os.Create(fmt.Sprintf("%s/hooks/commit-msg", path))
		assert.Nil(t, err)
		err = os.Mkdir(fmt.Sprintf("%s/xyz", path), os.ModePerm)
		assert.Nil(t, err)
		_, err = os.Create(fmt.Sprintf("%s/xyz/commit-msg", path))
		assert.Nil(t, err)

		return NewHandler(model.Config{GitPath: path})
	}

	t.Run("returns 0 and", func(t *testing.T) {

		t.Run("removes all occurrences of commit-msg", func(t *testing.T) {
			handler := createFakeHandlerWithDirs()

			status := handler.uninstall()

			path := handler.Config.GitPath
			assert.Equal(t, 0, status)
			assert.NoFileExists(t, fmt.Sprintf("%s/hooks/commit-msg", path))
			assert.FileExists(t, fmt.Sprintf("%s/xyz/commit-msg", path))
		})

		t.Run("logs a success message", func(t *testing.T) {
			buffer.Reset()
			handler := createFakeHandlerWithDirs()

			_ = handler.uninstall()

			assert.Contains(t, buffer.String(), "[SUCCESS]\t commit-message-check successfully uninstalled.")
		})
	})

	t.Run("returns 1 on any error", func(t *testing.T) {
		buffer.Reset()
		errPath := t.TempDir()
		err := os.Mkdir(fmt.Sprintf("%s/hooks", errPath), 0000)
		assert.Nil(t, err)
		handler := NewHandler(model.Config{GitPath: errPath})

		status := handler.uninstall()

		assert.Equal(t, 1, status)
		assert.Contains(t, buffer.String(), "[ERROR]\t Could not delete commit-msg hook.")
	})
}
