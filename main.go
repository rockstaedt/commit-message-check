package main

import (
	"log"
	"os"
	"rockstaedt/commit-message-check/src"
)

func main() {
	log.Println("[INFO]\t Validating commit message...")
	cm, err := src.CreateCommitMessageFrom(os.Args[1])
	if err != nil {
		log.Printf("[ERROR]\t Could not create object: %q", err.Error())
		os.Exit(1)
	}

	if cm.ValidateSubject() == false {
		log.Println("[ERROR]\t Subject line too long. Please fix.")
		os.Exit(1)
	}

	os.Exit(0)
}
