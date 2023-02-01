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

	hookBytes, _ := os.Create(fmt.Sprintf("%s/hooks/commit-msg", gitPath))
	_ = writeCommitMsgHook(hookBytes)

	return 0
}

func writeCommitMsgHook(writer io.Writer) error {
	_, err := fmt.Fprint(writer, "#!/bin/sh\n\n./commit-message-check validate $1\n")
	if err != nil {
		return err
	}

	return nil
}
