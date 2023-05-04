package cmd

import (
	"bytes"
	"fmt"
	"github.com/TwiN/go-color"
	"github.com/stretchr/testify/assert"
	"os"
	"rockstaedt/commit-message-check/internal/model"
	"testing"
)

func TestSetup(t *testing.T) {
	buffer := &bytes.Buffer{}

	t.Run("returns 0 and", func(t *testing.T) {

		fakeHandler := func() *Handler {
			path := t.TempDir()
			err := os.Mkdir(fmt.Sprintf("%s/hooks", path), os.ModePerm)
			assert.Nil(t, err)

			handler := NewHandler(model.Config{GitPath: path})
			handler.Writer = buffer

			return handler
		}()

		t.Run("creates commit-msg script in hook folder", func(t *testing.T) {
			status := fakeHandler.setup()

			assert.Equal(t, 0, status)
			assert.FileExists(t, fmt.Sprintf("%s/hooks/commit-msg", fakeHandler.Config.GitPath))
		})

		t.Run("logs a success message", func(t *testing.T) {
			buffer.Reset()

			_ = fakeHandler.setup()

			assert.Contains(t, buffer.String(), color.Green+"commit-message-check successfully installed.")
		})
	})

	t.Run("returns 1 when error at walking hooks and logs it", func(t *testing.T) {
		buffer.Reset()
		errPath := t.TempDir()
		err := os.Mkdir(fmt.Sprintf("%s/hooks", errPath), 0000)
		assert.Nil(t, err)
		handler := NewHandler(model.Config{GitPath: errPath})
		handler.Writer = buffer

		status := handler.setup()

		assert.Equal(t, 1, status)
		assert.Contains(t, buffer.String(), color.Red+"Could not create commit-msg script.")
	})
}
