package main

import (
	"log"
	"os"
	"rockstaedt/commit-message-check/src"
	"rockstaedt/commit-message-check/src/utils"
)

const (
	softLimit = 50
	hardLimit = 72
)

func main() {
	log.Println("[INFO]\t Validating commit message...")
	commitLines, err := utils.GetLinesFromTextFile(os.Args[1])
	if err != nil {
		log.Printf("[ERROR]\t Could not read commit message lines: %q", err.Error())
		os.Exit(1)
	}
	cm, err := src.CreateCommitMessageFrom(commitLines)
	if err != nil {
		log.Printf("[ERROR]\t Could not create object: %q", err.Error())
		os.Exit(2)
	}

	numOfExceedingChars := cm.ValidateSubject()
	if numOfExceedingChars == 0 {
		os.Exit(0)
	}

	if numOfExceedingChars > (hardLimit - softLimit) {
		log.Println("[ERROR]\t Abort commit. Subject line too long. Please fix.")
		os.Exit(3)
	}

	log.Printf("[WARN]\t Your subject exceeds the soft limit of 50 chars by %d chars.", numOfExceedingChars)

	os.Exit(0)
}
