package cmd

import (
	"log"
	"os"
	"rockstaedt/commit-message-check/util"
)

func Setup(gitPath string) int {
	_, err := os.Stat(gitPath)
	if err != nil {
		log.Println("[ERROR]\t No git repository could be found.")
		return 1
	}

	err = util.WalkHookDirs(gitPath, util.CreateHook)
	if err != nil {
		log.Println("[ERROR]\t Could not create commit-msg script.")
		return 2
	}

	log.Println("[SUCCESS]\t commit-message-check successfully installed.")
	return 0
}
