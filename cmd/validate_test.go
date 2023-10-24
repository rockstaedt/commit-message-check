package cmd

import (
	"bytes"
	"github.com/TwiN/go-color"
	"github.com/stretchr/testify/assert"
	"os"
	"rockstaedt/commit-message-check/internal/model"
	"testing"
)

func TestValidate(t *testing.T) {
	buffer := &bytes.Buffer{}

	t.Run("returns 0 on success", func(t *testing.T) {
		buffer.Reset()
		testFile := t.TempDir() + "/text.txt"
		err := os.WriteFile(testFile, []byte("i am a short commit msg"), 0666)
		assert.Nil(t, err)
		handler := NewHandler(model.Config{CommitMsgFile: testFile})
		handler.Writer = buffer

		handler.Run("validate")

		assert.Equal(t, 0, buffer.Len())
	})

	t.Run("returns 0 when soft limit exceeds and logs a warning", func(t *testing.T) {
		buffer.Reset()
		testFile := t.TempDir() + "/text.txt"
		err := os.WriteFile(testFile, []byte("i am two characters more thaaaaaaaaaaaaaaaaaaaaan 50"), 0666)
		assert.Nil(t, err)
		handler := NewHandler(model.Config{CommitMsgFile: testFile})
		handler.Writer = buffer

		status := handler.validate()

		assert.Equal(t, status, 0)
		assert.Contains(t, buffer.String(), color.Yellow+"Your subject exceeds the soft limit of 50 chars by 2 chars.")
		assert.Contains(t, buffer.String(), "i am two characters more thaaaaaaaaaaaaaaaaaaaaan "+color.Yellow+"50")
	})

	t.Run("asks user for abort when commit message too long", func(t *testing.T) {
		testFile := t.TempDir() + "/text.txt"
		content := "waaaaaaaaaaaaaaaaaaaaaaaaaay tooooooooooooooooooo" +
			"looooooooooooooooooooooong"
		err := os.WriteFile(testFile, []byte(content), 0666)
		assert.Nil(t, err)
		handler := NewHandler(model.Config{CommitMsgFile: testFile})
		handler.Writer = buffer

		t.Run("user confirms abort", func(t *testing.T) {
			buffer.Reset()
			reader := bytes.NewReader([]byte("y"))
			handler.Reader = reader

			status := handler.Run("validate")

			assert.Contains(t, buffer.String(), color.Red+"Subject line too long. Do you want to abort? (y/n)")
			assert.Equal(t, 1, status)
		})
		t.Run("user declines abort", func(t *testing.T) {
			buffer.Reset()
			reader := bytes.NewReader([]byte("n"))
			handler.Reader = reader

			status := handler.Run("validate")

			assert.Equal(t, 0, status)
		})
	})

	t.Run("returns 1 when error at reading file", func(t *testing.T) {
		buffer.Reset()
		handler := NewHandler(model.Config{CommitMsgFile: "/no_file"})
		handler.Writer = buffer

		status := handler.Run("validate")

		want := `Could not read commit message: "file not found"`
		assert.Contains(t, buffer.String(), want)
		assert.NotContains(t, buffer.String(), "Valid commit")
		assert.Equal(t, 1, status)
	})
}
