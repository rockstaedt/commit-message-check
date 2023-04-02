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

		status := myHandler.Run()

		assert.Contains(t, buffer.String(), "Could not delete")
		assert.True(t, status > 0)
	})

	t.Run("executes setup command", func(t *testing.T) {
		buffer.Reset()
		protectedPath := t.TempDir() + "/fake"
		err := os.Mkdir(protectedPath, 0000)
		assert.Nil(t, err)
		config := model.Config{Command: "setup", GitPath: protectedPath}
		myHandler := NewHandler(config)

		status := myHandler.Run()

		assert.Contains(t, buffer.String(), "Could not create")
		assert.True(t, status > 0)
	})

	t.Run("executes update command", func(t *testing.T) {
		buffer.Reset()
		config := model.Config{Command: "update"}
		myHandler := NewHandler(config)

		status := myHandler.Run()

		assert.Contains(t, buffer.String(), "Error at retrieving")
		assert.True(t, status > 0)
	})

	t.Run("executes validate command", func(t *testing.T) {

		t.Run("success", func(t *testing.T) {
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

		t.Run("error at reading file", func(t *testing.T) {
			buffer.Reset()
			config := model.Config{Command: "validate", CommitMsg: "/no_file"}
			myHandler := NewHandler(config)

			status := myHandler.Run()

			want := `Could not read commit message: "file not found"`
			assert.Contains(t, buffer.String(), want)
			assert.NotContains(t, buffer.String(), "Valid")
			assert.Equal(t, 3, status)
		})
	})

	t.Run("prints warning when any other command", func(t *testing.T) {
		buffer.Reset()
		config := model.Config{Command: "unknown"}
		myHandler := NewHandler(config)

		status := myHandler.Run()

		want := "Unknown subcommand. Please check manual."
		assert.Contains(t, buffer.String(), want)
		assert.Equal(t, 4, status)
	})
}
