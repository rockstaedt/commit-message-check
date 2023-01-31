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

	_ = os.WriteFile(fmt.Sprintf("%s/hooks/commit-msg", gitPath), nil, os.ModePerm)

	return 0
}

func writeCommitMsgHook(writer io.Writer) error {
	return nil
}
