package cmd

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"rockstaedt/commit-message-check/internal/model"
	"testing"
)

func TestRun(t *testing.T) {
	buffer := &bytes.Buffer{}
	log.SetOutput(buffer)

	t.Run("executes uninstall command", func(t *testing.T) {
		buffer.Reset()
		myHandler := NewHandler(model.Config{GitPath: "/"})

		status := myHandler.Run("uninstall")

		assert.Contains(t, buffer.String(), "Could not delete")
		assert.True(t, status > 0)
	})

	t.Run("executes setup command", func(t *testing.T) {
		buffer.Reset()
		protectedPath := t.TempDir() + "/fake"
		err := os.Mkdir(protectedPath, 0000)
		assert.Nil(t, err)
		myHandler := NewHandler(model.Config{GitPath: protectedPath})

		status := myHandler.Run("setup")

		assert.Contains(t, buffer.String(), "Could not create")
		assert.True(t, status > 0)
	})

	t.Run("executes update command", func(t *testing.T) {
		buffer.Reset()
		myHandler := NewHandler(model.Config{})

		status := myHandler.Run("update")

		assert.Contains(t, buffer.String(), "Error at retrieving")
		assert.True(t, status > 0)
	})

	t.Run("executes validate command", func(t *testing.T) {

		buffer.Reset()
		testFile := t.TempDir() + "/text.txt"
		err := os.WriteFile(testFile, []byte("i am a commit msg"), 0666)
		assert.Nil(t, err)
		myHandler := NewHandler(model.Config{CommitMsgFile: testFile})

		myHandler.Run("validate")

		assert.Contains(t, buffer.String(), "Valid commit message")
	})

	t.Run("prints warning when any other command", func(t *testing.T) {
		buffer.Reset()
		myHandler := NewHandler(model.Config{})

		status := myHandler.Run("unknown")

		want := "Unknown subcommand. Please check manual."
		assert.Contains(t, buffer.String(), want)
		assert.Equal(t, 1, status)
	})
}
