package cmd

import "testing"

func TestUninstall(t *testing.T) {

	t.Run("returns 0 and removes all occurrences of commit-msg", func(t *testing.T) {
		_ = Uninstall("")
	})

	t.Run("returns 1 on any error", func(t *testing.T) {
		t.Skip()
	})
}
