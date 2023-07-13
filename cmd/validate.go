package cmd

import (
	"fmt"
	"github.com/rockstaedt/txtreader"
	"rockstaedt/commit-message-check/internal/model"
)

const (
	softLimit = 50
	hardLimit = 72
)

func (h *Handler) validate() int {
	commitLines, err := txtreader.GetLinesFromTextFile(h.Config.CommitMsgFile)
	if err != nil {
		h.notify(fmt.Sprintf("Could not read commit message: %q", err.Error()), "red")
		return 1
	}

	cm := model.CreateCommitMessageFrom(commitLines)

	numOfExceedingChars := cm.ValidateSubject()
	if numOfExceedingChars == 0 {
		return 0
	}

	if numOfExceedingChars > (hardLimit - softLimit) {
		h.notify("Abort commit. Subject line too long. Please fix.", "red")
		return 1
	}

	message := fmt.Sprintf("Your subject exceeds the soft limit of 50 chars by %d chars.", numOfExceedingChars)
	h.notify(message, "yellow")

	return 0
}
