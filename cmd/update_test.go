package cmd

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUpdate(t *testing.T) {
	buffer := &bytes.Buffer{}
	log.SetOutput(buffer)

	t.Run("returns 0 and a message when local version is latest", func(t *testing.T) {
		buffer.Reset()
		ts := httptest.NewServer(getHandlerFor(`{"tag_name":"v1.0.0"}`))
		defer ts.Close()

		status := Update("v1.0.0", ts.URL)

		assert.Equal(t, 0, status)
		assert.Contains(t, buffer.String(), "Current version is latest version.")
	})

	t.Run("--downloads install script if newer version available", func(t *testing.T) {
		t.Skip()
	})

	t.Run("returns 0 and downloads install script if newer version available", func(t *testing.T) {
		t.Skip()
	})

	t.Run("returns 1 and message when error at request", func(t *testing.T) {
		buffer.Reset()
		ts := httptest.NewServer(getHandlerFor("", 500))
		defer ts.Close()

		status := Update("v1.0.0", ts.URL)

		assert.Equal(t, 1, status)
		assert.Contains(t, buffer.String(), "Error at retrieving latest version.")
	})
}

func TestGetLatestTag(t *testing.T) {
	t.Run("returns latest tag when request is successfully", func(t *testing.T) {
		ts := httptest.NewServer(getHandlerFor(`{"tag_name":"v1.2.0"}`))
		defer ts.Close()

		tag := getLatestTag(ts.URL)

		assert.Equal(t, "v1.2.0", tag)
	})

	t.Run("returns empty string when", func(t *testing.T) {

		t.Run("HTTP protocol error", func(t *testing.T) {
			tag := getLatestTag("xxx")

			assert.Empty(t, tag)
		})

		t.Run("response status code is not 200", func(t *testing.T) {
			ts := httptest.NewServer(getHandlerFor("", 500))
			defer ts.Close()

			tag := getLatestTag(ts.URL)

			assert.Empty(t, tag)
		})

		t.Run("response body is empty", func(t *testing.T) {
			ts := httptest.NewServer(getHandlerFor(""))
			defer ts.Close()

			tag := getLatestTag(ts.URL)

			assert.Empty(t, tag)
		})
	})
}

func getHandlerFor(resBody string, statusCode ...int) http.HandlerFunc {
	sc := 200
	if len(statusCode) > 0 {
		sc = statusCode[0]
	}

	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(sc)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(resBody))
	}
}
