package cmd

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestUpdate(t *testing.T) {
	buffer := &bytes.Buffer{}
	log.SetOutput(buffer)

	t.Run("returns 0 and", func(t *testing.T) {

		t.Run("logs a message when local version is latest", func(t *testing.T) {
			buffer.Reset()
			ts := httptest.NewServer(getHandlerFor(`{"tag_name":"v1.0.0"}`))
			defer ts.Close()
			config := &UpdateConfig{Version: "v1.0.0", TagUrl: ts.URL}

			status := Update(config)

			assert.Equal(t, 0, status)
			assert.Contains(t, buffer.String(), "Current version is latest version.")
		})

		t.Run("downloads install script if newer version available", func(t *testing.T) {
			ts := httptest.NewServer(getHandlerFor(`{"tag_name":"v1.1.0"}`))
			defer ts.Close()
			tempDir := t.TempDir()
			config := &UpdateConfig{Version: "v1.0.0", TagUrl: ts.URL, DownloadPath: tempDir}

			_ = Update(config)

			assert.FileExists(t, tempDir+"/commit-message-check")
		})
	})

	t.Run("returns 1 and message when error at request", func(t *testing.T) {
		buffer.Reset()
		ts := httptest.NewServer(getHandlerFor("", 500))
		defer ts.Close()
		config := &UpdateConfig{Version: "v1.0.0", TagUrl: ts.URL, DownloadPath: ""}

		status := Update(config)

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

func TestDownloadScript(t *testing.T) {

	getProtectedPath := func(t *testing.T) string {
		tempDir := t.TempDir()
		protectedPath := tempDir + "/protected"
		err := os.Mkdir(protectedPath, 0000)
		assert.Nil(t, err)

		return protectedPath
	}

	t.Run("returns 0 and downloads binary", func(t *testing.T) {
		tempDir := t.TempDir()
		err := os.WriteFile(tempDir+"/dummy", []byte("i am a go binary"), os.ModePerm)
		assert.Nil(t, err)
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, tempDir+"/dummy")
		}))
		defer ts.Close()
		config := &UpdateConfig{DownloadPath: tempDir, BinaryUrl: ts.URL}

		status := downloadScript(config)

		assert.Equal(t, 0, status)
		contentBytes, err := os.ReadFile(tempDir + "/commit-message-check")
		assert.Nil(t, err)
		assert.Contains(t, string(contentBytes), "i am a go binary")
	})

	t.Run("returns 1 when error at creating file", func(t *testing.T) {
		config := &UpdateConfig{DownloadPath: getProtectedPath(t)}

		status := downloadScript(config)

		assert.Equal(t, 1, status)
	})

	t.Run("return 2 when http protocol error", func(t *testing.T) {
		tempDir := t.TempDir()
		config := &UpdateConfig{DownloadPath: tempDir, BinaryUrl: "/xxx"}

		status := downloadScript(config)

		assert.Equal(t, 2, status)
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
