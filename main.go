package main

import (
	"flag"
	"fmt"
	"github.com/rockstaedt/txtreader"
	"log"
	"os"
	"rockstaedt/commit-message-check/cmd"
)

var version string

func main() {
	var versionFlag bool
	flag.BoolVar(&versionFlag, "v", false, "Shows the current version of the executable.")

	flag.Parse()

	if versionFlag {
		fmt.Println(version)
		os.Exit(0)
	}

	if len(os.Args) == 1 {
		fmt.Println("No subcommands given. Please check manual.")
		os.Exit(1)
	}

	var status int
	cwd, err := os.Getwd()
	if err != nil {
		log.Printf("[ERROR]\t Could not determine working directory: %q", err.Error())
		status = 1
	}
	gitPath := fmt.Sprintf("%s/.git", cwd)
	switch os.Args[1] {
	case "setup":
		status = cmd.Setup(gitPath)
	case "uninstall":
		status = cmd.Uninstall(gitPath)
	case "validate":
		commitLines, err := txtreader.GetLinesFromTextFile(os.Args[2])
		if err != nil {
			log.Printf("[ERROR]\t Could not read commit message lines: %q", err.Error())
			status = 1
		}

		status = cmd.Validate(commitLines)
	default:
		fmt.Printf("Unknown subcommand %q. Please check manual.\n", os.Args[1])
		status = 1
	}

	os.Exit(status)
}
