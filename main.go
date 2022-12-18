package main

import (
	"log"
	"os"
	"rockstaedt/commit-message-check/src"
	"rockstaedt/commit-message-check/src/utils"
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
		os.Exit(1)
	}

	if cm.ValidateSubject() == false {
		log.Println("[ERROR]\t Abort commit. Subject line too long. Please fix.")
		os.Exit(1)
	}

	os.Exit(0)
}
