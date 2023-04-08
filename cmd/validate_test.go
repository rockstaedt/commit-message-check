package cmd

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"rockstaedt/commit-message-check/internal/model"
	"testing"
)

func TestValidate(t *testing.T) {
	buffer := &bytes.Buffer{}
	log.SetOutput(buffer)

	t.Run("returns 0 on success", func(t *testing.T) {
		buffer.Reset()
		testFile := t.TempDir() + "/text.txt"
		err := os.WriteFile(testFile, []byte("i am a short commit msg"), 0666)
		assert.Nil(t, err)
		handler := NewHandler(model.Config{CommitMsgFile: testFile})

		handler.Run("validate")

		assert.Contains(t, buffer.String(), "Valid commit message")
	})

	t.Run("returns 0 when soft limit exceeds and logs a warning", func(t *testing.T) {
		buffer.Reset()
		testFile := t.TempDir() + "/text.txt"
		err := os.WriteFile(testFile, []byte("i am two characters more thaaaaaaaaaaaaaaaaaaaaan 50"), 0666)
		assert.Nil(t, err)
		handler := NewHandler(model.Config{CommitMsgFile: testFile})

		status := handler.validate()

		assert.Equal(t, status, 0)
		assert.Contains(t, buffer.String(), "[WARN]\t Your subject exceeds the soft limit of 50 chars by 2 chars.")
	})

	t.Run("returns 1 when commit message too long", func(t *testing.T) {
		buffer.Reset()
		testFile := t.TempDir() + "/text.txt"
		content := "waaaaaaaaaaaaaaaaaaaaaaaaaay tooooooooooooooooooo" +
			"looooooooooooooooooooooong"
		err := os.WriteFile(testFile, []byte(content), 0666)
		assert.Nil(t, err)
		myHandler := NewHandler(model.Config{CommitMsgFile: testFile})

		status := myHandler.Run("validate")

		assert.Contains(t, buffer.String(), "Abort commit")
		assert.Equal(t, 1, status)
	})

	t.Run("returns 1 when error at reading file", func(t *testing.T) {
		buffer.Reset()
		myHandler := NewHandler(model.Config{CommitMsgFile: "/no_file"})

		status := myHandler.Run("validate")

		want := `Could not read commit message: "file not found"`
		assert.Contains(t, buffer.String(), want)
		assert.NotContains(t, buffer.String(), "Valid commit")
		assert.Equal(t, 1, status)
	})
}
