package cmd

import "testing"

func TestUpdate(t *testing.T) {

	t.Run("returns 0 and", func(t *testing.T) {

		t.Run("requests latest release tag", func(t *testing.T) {
			t.Skip()
		})

		t.Run("compares latest tag with local tag", func(t *testing.T) {
			t.Skip()
		})

		t.Run("downloads install script if newer version available", func(t *testing.T) {
			t.Skip()
		})
	})

	t.Run("returns 1 on any error", func(t *testing.T) {
		t.Skip()
	})

}
