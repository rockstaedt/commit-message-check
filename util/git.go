package util

import (
	"fmt"
	"os"
)

func IsGitRepo(path string) bool {
	_, err := os.Stat(fmt.Sprintf("%s/.git", path))
	if err != nil {
		return false
	}

	return true
}
