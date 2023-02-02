package cmd

import (
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func Setup(gitPath string) int {
	_, err := os.Stat(gitPath)
	if err != nil {
		log.Println("[ERROR]\t No git repository could be found.")
		return 1
	}

	err = filepath.WalkDir(gitPath, func(p string, d fs.DirEntry, e error) error {
		if e != nil {
			return e
		}

		if d.Name() == "hooks" {
			file, err := os.Create(fmt.Sprintf("%s/commit-msg", p))
			if err != nil {
				return err
			}

			_ = file.Chmod(os.ModePerm)

			writeCommitMsgHook(file)
		}

		return nil
	})
	if err != nil {
		log.Println("[ERROR]\t Could not create commit-msg script.")
		return 2
	}

	return 0
}

func writeCommitMsgHook(writer io.Writer) {
	_, err := fmt.Fprint(writer, "#!/bin/sh\n\n./commit-message-check validate $1\n")
	if err != nil {
		log.Printf("[ERROR]\t Could not write commit-msg script: %s", err)
	}
}
