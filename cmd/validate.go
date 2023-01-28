package cmd

import (
	"log"
	"rockstaedt/commit-message-check/internal/model"
)

const (
	softLimit = 50
	hardLimit = 72
)

func Validate(commitLines []string) int {
	log.Println("[INFO]\t Validate commit message.")

	cm := model.CreateCommitMessageFrom(commitLines)

	numOfExceedingChars := cm.ValidateSubject()
	if numOfExceedingChars == 0 {
		log.Println("[SUCCESS]\t Valid commit message.")
		return 0
	}

	if numOfExceedingChars > (hardLimit - softLimit) {
		log.Println("[ERROR]\t Abort commit. Subject line too long. Please fix.")
		return 1
	}

	log.Printf("[WARN]\t Your subject exceeds the soft limit of 50 chars by %d chars.", numOfExceedingChars)

	return 0
}
