package cmd

import (
	"rockstaedt/commit-message-check/internal/model"
	"testing"
)

func TestHandler(t *testing.T) {

	t.Run("executes uninstall command", func(t *testing.T) {
		config := model.Config{Command: "uninstall", GitPath: "/:"}
		myHandler := NewHandler(config)

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
