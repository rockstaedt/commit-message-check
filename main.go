package main

import (
	"flag"
	"fmt"
	"github.com/rockstaedt/txtreader"
	"log"
	"os"
	"rockstaedt/commit-message-check/cmd"
	"rockstaedt/commit-message-check/internal/model"
	"rockstaedt/commit-message-check/util"
)

var version string

func main() {
	flag.Usage = func() {
		util.PrintManual(os.Stderr)
	}

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

	cwd, err := os.Getwd()
	if err != nil {
		log.Printf("[ERROR]\t Could not determine working directory: %q", err.Error())
		os.Exit(1)
	}

	gitPath := fmt.Sprintf("%s/.git", cwd)
	_, err = os.Stat(gitPath)
	if err != nil {
		log.Println("[ERROR]\t No git repository could be found.")
		os.Exit(2)
	}

	var status int
	switch os.Args[1] {
	case "setup":
		status = cmd.Setup(gitPath)
	case "uninstall":
		status = cmd.Uninstall(gitPath)
	case "update":
		config := &model.Config{
			Version:       version,
			TagUrl:        "https://api.github.com/repos/rockstaedt/commit-message-check/releases/latest",
			BinaryBaseUrl: "https://github.com/rockstaedt/commit-message-check/releases/latest/download/",
			DownloadPath:  cwd,
		}

		status = cmd.Update(config)

		if status > 0 {
			log.Println("[ERROR]\t Could not update commit-message-check.")
			break
		}
		log.Printf("[SUCCESS]\t Updated commit-message-check successfully to %s", config.LatestVersion)
	case "validate":
		commitLines, err := txtreader.GetLinesFromTextFile(os.Args[2])
		if err != nil {
			log.Printf("[ERROR]\t Could not read commit message lines: %q", err.Error())
			status = 3
		}

		status = cmd.Validate(commitLines)
	default:
		fmt.Printf("Unknown subcommand %q. Please check manual with -h flag.\n", os.Args[1])
		status = 4
	}

	os.Exit(status)
}
