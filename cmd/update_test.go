package cmd

import (
	"bytes"
	"fmt"
	"github.com/TwiN/go-color"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"rockstaedt/commit-message-check/internal/model"
	"runtime"
	"testing"
)

func TestUpdate(t *testing.T) {
	buffer := &bytes.Buffer{}

	t.Run("returns 0 and", func(t *testing.T) {

		t.Run("logs a message when local version is latest", func(t *testing.T) {
			buffer.Reset()
			ts := httptest.NewServer(getHandlerFor(`{"tag_name":"v1.0.0"}`))
			defer ts.Close()
			handler := NewHandler(model.Config{Version: "v1.0.0", TagUrl: ts.URL})
			handler.Writer = buffer

			status := handler.update()

			assert.Equal(t, 0, status)
			assert.Contains(t, buffer.String(), "Current version is latest version.")
		})

		t.Run("downloads install script if newer version available", func(t *testing.T) {
			buffer.Reset()
			tsTag := httptest.NewServer(getHandlerFor(`{"tag_name":"v1.1.0"}`))
			defer tsTag.Close()
			tsBinary := httptest.NewServer(getHandlerFor("i am a binary"))
			defer tsBinary.Close()
			tempDir := t.TempDir()
			handler := NewHandler(model.Config{
				Version:       "v1.0.0",
				BinaryBaseUrl: tsBinary.URL,
				TagUrl:        tsTag.URL,
				DownloadPath:  tempDir,
			})
			handler.Writer = buffer

			status := handler.update()

			assert.FileExists(t, tempDir+"/commit-message-check")
			assert.Equal(t, 0, status)
			assert.Contains(t, buffer.String(), "updated successfully to")
		})
	})

	t.Run("returns 1 and message when error at request", func(t *testing.T) {

		t.Run("for tag", func(t *testing.T) {
			buffer.Reset()
			ts := httptest.NewServer(getHandlerFor("", 500))
			defer ts.Close()
			handler := NewHandler(model.Config{Version: "v1.0.0", TagUrl: ts.URL, DownloadPath: ""})
			handler.Writer = buffer

			status := handler.update()

			assert.Equal(t, 1, status)
			assert.Contains(t, buffer.String(), color.Red+"Error at retrieving latest version.")
		})

		t.Run("for binary", func(t *testing.T) {
			buffer.Reset()
			tsTag := httptest.NewServer(getHandlerFor(`{"tag_name":"v1.1.0"}`))
			defer tsTag.Close()
			tsBinary := httptest.NewServer(getHandlerFor("", 500))
			defer tsBinary.Close()
			tempDir := t.TempDir()
			handler := NewHandler(model.Config{
				Version:       "v1.0.0",
				BinaryBaseUrl: tsBinary.URL,
				TagUrl:        tsTag.URL,
				DownloadPath:  tempDir,
			})
			handler.Writer = buffer

			status := handler.update()

			assert.Equal(t, 1, status)
			assert.Contains(t, buffer.String(), color.Red+"Error while downloading binary.")
		})
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

func TestDownloadScript(t *testing.T) {

	getProtectedPath := func(t *testing.T) string {
		tempDir := t.TempDir()
		protectedPath := tempDir + "/protected"
		err := os.Mkdir(protectedPath, 0000)
		assert.Nil(t, err)

		return protectedPath
	}

	t.Run("returns success message and writes downloaded binary content to file", func(t *testing.T) {
		tempDir := t.TempDir()
		err := os.WriteFile(tempDir+"/dummy", []byte("i am a go binary"), os.ModePerm)
		assert.Nil(t, err)
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			targetUrl := fmt.Sprintf("/commit-message-check-v1.1.1-%s-%s", runtime.GOOS, runtime.GOARCH)
			if r.URL.String() == targetUrl {
				http.ServeFile(w, r, tempDir+"/dummy")
			}
		}))
		defer ts.Close()
		config := &model.Config{LatestVersion: "v1.1.1", DownloadPath: tempDir, BinaryBaseUrl: ts.URL}

		msg := downloadScript(config)

		contentBytes, err := os.ReadFile(tempDir + "/commit-message-check")
		assert.Nil(t, err)
		assert.Contains(t, string(contentBytes), "i am a go binary")
		wantedUpdateMsg := "commit-message-check updated successfully to v1.1.1"
		assert.Contains(t, msg, wantedUpdateMsg)
	})

	t.Run("returns empty string when", func(t *testing.T) {

		t.Run("error at creating file", func(t *testing.T) {
			config := &model.Config{DownloadPath: getProtectedPath(t)}

			msg := downloadScript(config)

			assert.Equal(t, "", msg)
		})

		t.Run("http protocol error", func(t *testing.T) {
			tempDir := t.TempDir()
			config := &model.Config{DownloadPath: tempDir, BinaryBaseUrl: "/xxx"}

			msg := downloadScript(config)

			assert.Equal(t, "", msg)
		})

		t.Run("request not successfully", func(t *testing.T) {
			tempDir := t.TempDir()
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
				w.WriteHeader(500)
			}))
			defer ts.Close()
			config := &model.Config{DownloadPath: tempDir, BinaryBaseUrl: ts.URL}

			msg := downloadScript(config)

			assert.Equal(t, "", msg)
		})
	})
}
