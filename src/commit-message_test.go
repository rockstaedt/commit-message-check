package src

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCommitMessage(t *testing.T) {
	t.Run("create new commit message object from File", func(t *testing.T) {
		cm, err := CreateCommitMessageFrom("../testdata/cm.txt")
		assert.Nil(t, err)

		assert.Equal(t, "commit subject", cm.Subject)
		assert.Len(t, cm.Body, 2)
		assert.False(t, cm.InvalidBody)
	})

	t.Run("Handles empty commit message", func(t *testing.T) {
		cm, err := CreateCommitMessageFrom("../testdata/cm-empty.txt")
		assert.Nil(t, err)

		assert.Equal(t, "", cm.Subject)
		assert.Len(t, cm.Body, 0)
	})

	t.Run("Handles no linebreak between subject and body", func(t *testing.T) {
		cm, err := CreateCommitMessageFrom("../testdata/cm-malformed.txt")
		assert.Nil(t, err)

		t.Run("sets body correct", func(t *testing.T) {
			assert.Equal(t, "commit subject", cm.Subject)
			assert.Len(t, cm.Body, 2)
		})

		t.Run("indicates body is malformed", func(t *testing.T) {
			assert.True(t, cm.InvalidBody)
		})
	})

	t.Run("Handles error if file was not found", func(t *testing.T) {
		_, err := CreateCommitMessageFrom("unknown_file")

		assert.Error(t, err)
	})

	t.Run("Validates subject line", func(t *testing.T) {
		testcases := []struct {
			desc             string
			inputFile        string
			wantedValidation bool
		}{
			{"more than 50 characters", "cm-long.txt", false},
			{"50 characters or less", "cm.txt", true},
		}

		for _, tc := range testcases {
			t.Run(tc.desc, func(t *testing.T) {
				cm, _ := CreateCommitMessageFrom(fmt.Sprintf("../testdata/%s", tc.inputFile))

				isValid := cm.ValidateSubject()

				assert.Equal(t, tc.wantedValidation, isValid)
			})
		}
	})
}
