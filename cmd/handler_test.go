package cmd

import (
	"bytes"
	"errors"
	"github.com/TwiN/go-color"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"log"
	"os"
	"rockstaedt/commit-message-check/internal/model"
	"rockstaedt/commit-message-check/testdata/mocks"
	"testing"
)

func TestRun(t *testing.T) {
	buffer := &bytes.Buffer{}
	log.SetOutput(buffer)

	t.Run("executes uninstall command", func(t *testing.T) {
		buffer.Reset()
		myHandler := NewHandler(model.Config{GitPath: "/"})
		myHandler.Writer = buffer

		status := myHandler.Run("uninstall")

		assert.Contains(t, buffer.String(), "Could not delete")
		assert.True(t, status > 0)
	})

	t.Run("executes setup command", func(t *testing.T) {
		buffer.Reset()
		protectedPath := t.TempDir() + "/fake"
		err := os.Mkdir(protectedPath, 0000)
		assert.Nil(t, err)
		myHandler := NewHandler(model.Config{GitPath: protectedPath})
		myHandler.Writer = buffer

		status := myHandler.Run("setup")

		assert.Contains(t, buffer.String(), "Could not create")
		assert.True(t, status > 0)
	})

	t.Run("executes update command", func(t *testing.T) {
		buffer.Reset()
		myHandler := NewHandler(model.Config{})
		myHandler.Writer = buffer

		status := myHandler.Run("update")

		assert.Contains(t, buffer.String(), "Error at retrieving")
		assert.True(t, status > 0)
	})

	t.Run("executes validate command", func(t *testing.T) {

		buffer.Reset()
		testFile := t.TempDir() + "/text.txt"
		err := os.WriteFile(testFile, []byte("i am a commit msg"), 0666)
		assert.Nil(t, err)
		myHandler := NewHandler(model.Config{CommitMsgFile: testFile})

		myHandler.Run("validate")

		assert.Contains(t, buffer.String(), "Valid commit message")
	})

	t.Run("prints warning when any other command", func(t *testing.T) {
		buffer.Reset()
		myHandler := NewHandler(model.Config{})

		status := myHandler.Run("unknown")

		want := "Unknown subcommand. Please check manual."
		assert.Contains(t, buffer.String(), want)
		assert.Equal(t, 1, status)
	})
}

func TestNotify(t *testing.T) {
	fwm := &mocks.FakeWriterMock{}
	handler := NewHandler(model.Config{})
	handler.Writer = fwm

	t.Run("writes a message to the writer", func(t *testing.T) {
		fwm.ResetCalls()
		fwm.On("Write", mock.Anything).Return(1, nil)

		handler.notify("I am a message")

		fwm.AssertCalled(t, "Write", []byte("I am a message\n"))
	})

	t.Run("colorize text in", func(t *testing.T) {

		t.Run("green", func(t *testing.T) {
			fwm.ResetCalls()
			fwm.On("Write", mock.Anything).Return(1, nil)

			handler.notify("I am a message", "green")

			fwm.AssertCalled(t, "Write", []byte(color.Green+"I am a message"+color.Reset+"\n"))
		})

		t.Run("red", func(t *testing.T) {
			fwm.ResetCalls()
			fwm.On("Write", mock.Anything).Return(1, nil)

			handler.notify("I am a message", "red")

			fwm.AssertCalled(t, "Write", []byte(color.Red+"I am a message"+color.Reset+"\n"))
		})

		t.Run("yellow", func(t *testing.T) {
			fwm.ResetCalls()
			fwm.On("Write", mock.Anything).Return(1, nil)

			handler.notify("I am", "yellow")

			fwm.AssertCalled(t, "Write", []byte(color.Yellow+"I am"+color.Reset+"\n"))
		})
	})

	t.Run("handles error at writing", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		log.SetOutput(buffer)
		fwm.ResetCalls()
		fwm.On("Write", mock.Anything).Return(0, errors.New("error at writing"))

		handler.notify("this causes an error")

		assert.Contains(t, buffer.String(), "Error at writing!")
	})
}
