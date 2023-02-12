package cmd

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestValidate(t *testing.T) {
	buffer := &bytes.Buffer{}
	log.SetOutput(buffer)

	t.Run("returns 0 if valid and logs a success", func(t *testing.T) {
		buffer.Reset()
		commitLines := []string{"i am shorter than 50 characters"}

		status := Validate(commitLines)

		assert.Equal(t, status, 0)
		assert.Contains(t, buffer.String(), "[INFO]\t Validate commit message.")
		assert.Contains(t, buffer.String(), "[SUCCESS]\t Valid commit message")
	})

	t.Run("returns 0 if soft limit exceeds and logs a warning", func(t *testing.T) {
		buffer.Reset()
		commitLines := []string{"i am two characters more thaaaaaaaaaaaaaaaaaaaaan 50"}

		status := Validate(commitLines)

		assert.Equal(t, status, 0)
		assert.Contains(t, buffer.String(), "[WARN]\t Your subject exceeds the soft limit of 50 chars by 2 chars.")
	})

	t.Run("returns 1 if length > 70 and logs an error", func(t *testing.T) {
		buffer.Reset()
		commitLines := []string{"i am way toooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooo long"}

		status := Validate(commitLines)

		assert.Equal(t, status, 1)
		assert.Contains(t, buffer.String(), "[ERROR]\t Abort commit. Subject line too long. Please fix.")
	})
}
