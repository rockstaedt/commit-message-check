package cmd

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUpdate(t *testing.T) {

	t.Run("returns 0 and a message when local version is latest", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			_, err := fmt.Fprintln(w, `{"tag_name":"v1.0.0"}`)
			assert.Nil(t, err)
		}))
		defer ts.Close()

		status := Update("v1.0.0", ts.URL)

		assert.Equal(t, 0, status)
	})

	t.Run("--downloads install script if newer version available", func(t *testing.T) {
		t.Skip()
	})

	t.Run("returns 1 and no message when a newer version was found", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			_, err := fmt.Fprintln(w, `{"tag_name":"v1.2.0"}`)
			assert.Nil(t, err)
		}))
		defer ts.Close()

		status := Update("v1.0.0", ts.URL)

		assert.Equal(t, 1, status)
	})

	t.Run("returns 1 on any error", func(t *testing.T) {
		t.Skip()
	})

}
