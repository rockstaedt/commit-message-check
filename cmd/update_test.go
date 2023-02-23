package cmd

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUpdate(t *testing.T) {

	getHandlerFor := func(resBody string, statusCode ...int) http.HandlerFunc {
		sc := 200
		if len(statusCode) > 0 {
			sc = statusCode[0]
		}

		return func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(sc)
			w.Header().Set("Content-Type", "application/json")
			_, err := w.Write([]byte(resBody))
			assert.Nil(t, err)
		}
	}

	t.Run("returns 0 and a message when local version is latest", func(t *testing.T) {
		ts := httptest.NewServer(getHandlerFor(`{"tag_name":"v1.0.0"}`))
		defer ts.Close()

		status := Update("v1.0.0", ts.URL)

		assert.Equal(t, 0, status)
	})

	t.Run("--downloads install script if newer version available", func(t *testing.T) {
		t.Skip()
	})

	t.Run("return 1 when response status code is not 200", func(t *testing.T) {
		ts := httptest.NewServer(getHandlerFor("v1.1.1", 500))
		defer ts.Close()

		status := Update("v1.0.0", ts.URL)

		assert.Equal(t, 1, status)
	})

	t.Run("returns 2 when response body is empty", func(t *testing.T) {
		ts := httptest.NewServer(getHandlerFor(""))
		defer ts.Close()

		status := Update("v1.0.0", ts.URL)

		assert.Equal(t, 2, status)
	})

	t.Run("returns 3 and no message when a newer version was found", func(t *testing.T) {
		ts := httptest.NewServer(getHandlerFor(`{"tag_name":"v1.2.0"}`))
		defer ts.Close()

		status := Update("v1.0.0", ts.URL)

		assert.Equal(t, 3, status)
	})
}
