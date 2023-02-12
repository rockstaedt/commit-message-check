package model

import (
	"github.com/stretchr/testify/assert"
	"regexp"
	"strings"
	"testing"
)

func TestNewCommitMessage(t *testing.T) {
	validCommitMsgLines := []string{
		"I am a valid Subject with less than 50 characters",
		"",
		"At this point comes the descriptive explanation what awesome stuff this",
		"commit brings. And for this reason the line width increases up to",
		"72 chars :)",
	}

	t.Run("create new commit message object from File", func(t *testing.T) {
		cm := CreateCommitMessageFrom(validCommitMsgLines)

		assert.Equal(t, "I am a valid Subject with less than 50 characters", cm.Subject)
		assert.Len(t, cm.Body, 3)
		assert.False(t, cm.InvalidBody)
	})

	t.Run("Handles empty commit message", func(t *testing.T) {
		var emptyCommitMsgLines []string
		cm := CreateCommitMessageFrom(emptyCommitMsgLines)

		assert.Equal(t, "", cm.Subject)
		assert.Len(t, cm.Body, 0)
	})

	t.Run("Handles no linebreak between subject and body", func(t *testing.T) {
		invalidCommitMsgLines := []string{
			"subject line",
			"body line 1",
			"body line 2",
		}
		cm := CreateCommitMessageFrom(invalidCommitMsgLines)

		t.Run("sets body correct", func(t *testing.T) {
			assert.Equal(t, "subject line", cm.Subject)
			assert.Len(t, cm.Body, 2)
		})

		t.Run("indicates body is malformed", func(t *testing.T) {
			assert.True(t, cm.InvalidBody)
		})
	})

	t.Run("Validates subject line", func(t *testing.T) {
		getDescFrom := func(subject string) string {
			re := regexp.MustCompile(`\.+|(#\d+ - )`)
			return strings.TrimSpace(re.ReplaceAllString(subject, " "))
		}

		testcases := []struct {
			subject     string
			expectation int
		}{
			{"more than................72....................................characters", 23},
			{"more than................50................less than 72 characters", 16},
			{"#1301 - more than........50..............through ID prefix", 0},
			{"short subject line", 0},
		}

		for _, tc := range testcases {
			t.Run(getDescFrom(tc.subject), func(t *testing.T) {
				cm := CreateCommitMessageFrom([]string{tc.subject})

				aboveSoftLimit := cm.ValidateSubject()

				assert.Equal(t, tc.expectation, aboveSoftLimit)
			})
		}
	})
}
