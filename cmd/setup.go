package cmd

import (
	"fmt"
	"io"
	"log"
	"os"
)

func Setup(gitPath string) int {
	_, err := os.Stat(gitPath)
	if err != nil {
		log.Println("[ERROR]\t No git repository could be found.")
		return 1
	}

	hookBytes, err := os.Create(fmt.Sprintf("%s/hooks/commit-msg", gitPath))
	if err != nil {
		log.Println("[ERROR]\t Could not create commit-msg script.")
		return 2
	}

	return writeCommitMsgHook(hookBytes)
}

func writeCommitMsgHook(writer io.Writer) int {
	_, err := fmt.Fprint(writer, "#!/bin/sh\n\n./commit-message-check validate $1\n")
	if err != nil {
		log.Println("[ERROR]\t Could not write to commit-msg script.")
		return 3
	}

	return 0
}
