package main

import (
	"fmt"
	"os"
	"rockstaedt/commit-message-check/src"
)

func main() {
	fmt.Println("[INFO]\t Validating commit message...")
	cm, err := src.CreateCommitMessageFrom(os.Args[1])
	if err != nil {
		panic(err)
	}

	if cm.ValidateSubject() == false {
		fmt.Println("[ERROR]\t Subject line too long. Please fix.")
		os.Exit(1)
	}

	os.Exit(0)
}
