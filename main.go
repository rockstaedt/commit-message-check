package main

import (
	"flag"
	"fmt"
	"github.com/rockstaedt/txtreader"
	"log"
	"os"
	"rockstaedt/commit-message-check/src"
)

const (
	softLimit = 50
	hardLimit = 72
)

var version string

func main() {
	versionPtr := flag.Bool("v", false, "Prints the current version of the executable")
	flag.Parse()

	if *versionPtr {
		fmt.Println(version)
		os.Exit(0)
	}

	log.Println("[INFO]\t Validating commit message...")
	commitLines, err := txtreader.GetLinesFromTextFile(os.Args[1])
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
