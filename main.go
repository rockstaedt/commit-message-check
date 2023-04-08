package main

import (
	"flag"
	"fmt"
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

	var commitMsg string
	if len(os.Args) == 3 {
		commitMsg = os.Args[2]
	}

	config := model.Config{
		CommitMsg:     commitMsg,
		GitPath:       gitPath,
		Version:       version,
		TagUrl:        "https://api.github.com/repos/rockstaedt/commit-message-check/releases/latest",
		BinaryBaseUrl: "https://github.com/rockstaedt/commit-message-check/releases/latest/download/",
		DownloadPath:  cwd,
	}

	handler := cmd.NewHandler(config)

	os.Exit(handler.Run(os.Args[1]))
}
