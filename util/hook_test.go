package util

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"log"
	"os"
	"rockstaedt/commit-message-check/testdata/mocks"
	"testing"
)

func TestWalkHookDirs(t *testing.T) {

	createDirs := func() string {
		path := t.TempDir()
		err := os.Mkdir(fmt.Sprintf("%s/hooks", path), os.ModePerm)
		assert.Nil(t, err)
		err = os.Mkdir(fmt.Sprintf("%s/no-hook", path), os.ModePerm)
		assert.Nil(t, err)
		err = os.MkdirAll(fmt.Sprintf("%s/nested/level/hooks", path), os.ModePerm)
		assert.Nil(t, err)

		return path
	}

	fakeDo := func(path, exePath string) error {
		log.Printf("running in %s\n", path)

		return nil
	}

	t.Run("runs do function only in hook dirs", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		log.SetOutput(buffer)
		path := createDirs()

		_ = WalkHookDirs(path, fakeDo)

		assert.Contains(t, buffer.String(), fmt.Sprintf("%s/hooks", path))
		assert.Contains(t, buffer.String(), fmt.Sprintf("%s/nested/level/hooks", path))
		assert.NotContains(t, buffer.String(), fmt.Sprintf("%s/no-hook", path))
	})

	t.Run("returns any error", func(t *testing.T) {

		t.Run("on walking", func(t *testing.T) {
			path := t.TempDir()
			err := os.Chmod(path, 0000)
			assert.Nil(t, err)

			err = WalkHookDirs(path, fakeDo)

			assert.Contains(t, err.Error(), "permission denied")
		})

		t.Run("on running do func", func(t *testing.T) {
			path := createDirs()
			wantedErr := errors.New("error at doing")
			errDo := func(path, exePath string) error {
				return wantedErr
			}

			err := WalkHookDirs(path, errDo)

			assert.ErrorIs(t, err, wantedErr)
		})
	})
}

func TestCreateHook(t *testing.T) {
	hookPath := t.TempDir()

	t.Run("creates hook script", func(t *testing.T) {
		_ = CreateHook(hookPath, "root")

		assert.FileExists(t, fmt.Sprintf("%s/commit-msg", hookPath))
	})

	t.Run("makes script executable", func(t *testing.T) {
		_ = CreateHook(hookPath, "root")

		info, err := os.Stat(fmt.Sprintf("%s/commit-msg", hookPath))
		assert.Nil(t, err)
		assert.Equal(t, info.Mode(), os.ModePerm)
	})

	t.Run("fills content", func(t *testing.T) {
		_ = CreateHook(hookPath, "root")

		contentBytes, err := os.ReadFile(fmt.Sprintf("%s/commit-msg", hookPath))
		assert.Nil(t, err)
		assert.Contains(t, string(contentBytes), `"root/commit-message-check" validate`)
	})

	t.Run("returns any error", func(t *testing.T) {
		path := t.TempDir()
		protectedPath := fmt.Sprintf("%s/protected_dir", path)
		err := os.Mkdir(protectedPath, 0000)
		assert.Nil(t, err)

		err = CreateHook(protectedPath, "root")

		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "permission denied")
	})
}

func TestDeleteHook(t *testing.T) {
	hookPath := t.TempDir()

	t.Run("deletes hook", func(t *testing.T) {
		_, err := os.Create(fmt.Sprintf("%s/commit-msg", hookPath))
		assert.Nil(t, err)

		_ = DeleteHook(hookPath, "")

		assert.NoFileExists(t, fmt.Sprintf("%s/commit-msg", hookPath))
	})

	t.Run("returns any error", func(t *testing.T) {
		err := DeleteHook(hookPath, "")

		assert.Contains(t, err.Error(), "no such file")
	})
}

func TestWriteContent(t *testing.T) {
	buffer := &bytes.Buffer{}

	t.Run("marks file as shell script and returns 0", func(t *testing.T) {
		buffer.Reset()

		writeContent(buffer, "usr/tmp")

		assert.Contains(t, buffer.String(), "#!/bin/sh\n\n")
		assert.Contains(t, buffer.String(), "exec < /dev/tty\n\n")
	})

	t.Run("executes commit-message-check with root path and quotes path to handle spaces", func(t *testing.T) {
		buffer.Reset()

		writeContent(buffer, "usr/tmp")

		assert.Contains(t, buffer.String(), `"usr/tmp/commit-message-check" validate $1`)
	})

	t.Run("inserts a blank line at the end of file", func(t *testing.T) {
		buffer.Reset()

		writeContent(buffer, "usr/tmp")

		assert.Contains(t, buffer.String(), "$1\n")
	})

	t.Run("logs any error", func(t *testing.T) {
		log.SetOutput(buffer)
		fwm := &mocks.FakeWriterMock{}
		fwm.On("Write", mock.Anything).Return(0, errors.New("error at writing"))

		writeContent(fwm, "usr/tmp")

		assert.Contains(t, buffer.String(), "[ERROR]\t Could not write commit-msg script: error at writing")
	})
}
