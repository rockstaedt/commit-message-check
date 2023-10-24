package cmd

import (
	"fmt"
	"github.com/TwiN/go-color"
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
		h.notify("Subject line too long. Do you want to abort? (y/n)", "red")

		var decision string
		if _, err := fmt.Fscanln(h.Reader, &decision); err != nil {
			h.notify("Could not read user input.", "red")
			return 1
		}

		if decision == "y" {
			return 1
		}

		return 0
	}

	message := fmt.Sprintf("Your subject exceeds the soft limit of 50 chars by %d chars.", numOfExceedingChars)
	h.notify(message, "yellow")
	h.notify(cm.Subject[:softLimit] + color.InYellow(cm.Subject[softLimit:]))

	return 0
}
