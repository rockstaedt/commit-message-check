package util

import (
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

type DoPathFunc func(path string) error

func WalkHookDirs(gitPath string, do DoPathFunc) error {
	return filepath.WalkDir(gitPath, func(p string, d fs.DirEntry, e error) error {
		if e != nil {
			return e
		}

		if d.Name() == "hooks" {
			err := do(p)
			if err != nil {
				return err
			}
		}

		return nil
	})
}

func CreateHook(path string) error {
	file, err := os.Create(fmt.Sprintf("%s/commit-msg", path))
	if err != nil {
		return err
	}

	_ = file.Chmod(os.ModePerm)

	writeContent(file)

	return nil
}

func writeContent(writer io.Writer) {
	_, err := fmt.Fprint(writer, "#!/bin/sh\n\n./commit-message-check validate $1\n")
	if err != nil {
		log.Printf("[ERROR]\t Could not write commit-msg script: %s", err)
	}
}
