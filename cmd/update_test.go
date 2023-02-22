package cmd

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUpdate(t *testing.T) {

	t.Run("returns 0 and", func(t *testing.T) {

		t.Run("requests latest release version", func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			}))
			defer ts.Close()

			status := Update("v1.0.0", ts.URL)

			assert.Equal(t, 0, status)
		})

		t.Run("compares latest version with binary version", func(t *testing.T) {
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
