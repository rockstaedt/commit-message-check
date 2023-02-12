package util

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestPrintManual(t *testing.T) {
	buffer := &bytes.Buffer{}
	log.SetOutput(buffer)

	t.Run("prints intro sentence", func(t *testing.T) {
		buffer.Reset()

		PrintManual(buffer)

		assert.Contains(t, buffer.String(), "Manual for commit-message-check:\n")
	})

	t.Run("prints flag section with", func(t *testing.T) {

		t.Run("heading", func(t *testing.T) {
			buffer.Reset()

			PrintManual(buffer)

			assert.Contains(t, buffer.String(), "- Flags:\n")
		})

		t.Run("help for version flag", func(t *testing.T) {
			buffer.Reset()

			PrintManual(buffer)

			assert.Contains(t, buffer.String(), "\tv\t\tShows the current version of the executable.")
		})
	})

	t.Run("prints subcommand section with", func(t *testing.T) {

		t.Run("heading", func(t *testing.T) {
			buffer.Reset()

			PrintManual(buffer)

			assert.Contains(t, buffer.String(), "- Subcommands:\n")
		})

		t.Run("help for", func(t *testing.T) {

			t.Run("setup", func(t *testing.T) {
				buffer.Reset()

				PrintManual(buffer)

				msg := "\tsetup\t\tInstalls the commit-msg script in every hook directory.\n"
				assert.Contains(t, buffer.String(), msg)
			})

			t.Run("uninstall", func(t *testing.T) {
				buffer.Reset()

				PrintManual(buffer)

				msg := "\tuninstall\tRemoves all commit-msg scripts.\n"
				assert.Contains(t, buffer.String(), msg)
			})

			t.Run("validate", func(t *testing.T) {
				buffer.Reset()

				PrintManual(buffer)

				msg := "\tvalidate <PATH>\tValidates the commit message written in <PATH>.\n"
				assert.Contains(t, buffer.String(), msg)
			})
		})
	})
}
