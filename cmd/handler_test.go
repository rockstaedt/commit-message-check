package cmd

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"log"
	"rockstaedt/commit-message-check/internal/model"
	"testing"
)

func TestHandler(t *testing.T) {
	buffer := &bytes.Buffer{}
	log.SetOutput(buffer)

	t.Run("executes uninstall command", func(t *testing.T) {
		buffer.Reset()
		config := model.Config{Command: "uninstall", GitPath: "/"}
		myHandler := NewHandler(config)

		myHandler.Run()

		assert.Contains(t, buffer.String(), "Could not delete")
	})

	t.Run("executes setup command", func(t *testing.T) {
		buffer.Reset()
		config := model.Config{Command: "setup", GitPath: t.TempDir()}
		myHandler := NewHandler(config)

		myHandler.Run()

		assert.Contains(t, buffer.String(), "successfully")
	})

	t.Run("executes update command", func(t *testing.T) {
		t.Skip()
	})

	t.Run("executes validate command", func(t *testing.T) {
		t.Skip()
	})

	t.Run("prints warning when any other command", func(t *testing.T) {
		t.Skip()
	})
}
