package cmd

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
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
		buffer.Reset()
		config := model.Config{Command: "update"}
		myHandler := NewHandler(config)

		myHandler.Run()

		assert.Contains(t, buffer.String(), "Error at retrieving")
	})

	t.Run("executes validate command", func(t *testing.T) {
		buffer.Reset()
		dir := t.TempDir()
		testFile := dir + "/text.txt"
		err := os.WriteFile(testFile, []byte("i am a commit msg"), 0666)
		assert.Nil(t, err)
		config := model.Config{Command: "validate", CommitMsg: testFile}
		myHandler := NewHandler(config)

		myHandler.Run()

		assert.Contains(t, buffer.String(), "Valid commit message")
	})

	t.Run("prints warning when any other command", func(t *testing.T) {
		t.Skip()
	})
}
