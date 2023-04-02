package cmd

import (
	"testing"
)

func TestHandler(t *testing.T) {

	t.Run("executes uninstall command", func(t *testing.T) {
		myHandler := NewHandler("uninstall")

		myHandler.Run()
	})

	t.Run("executes setup command", func(t *testing.T) {
		t.Skip()
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
