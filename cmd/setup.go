package cmd

import (
	"log"
	"rockstaedt/commit-message-check/util"
)

func Setup(gitPath string) int {
	err := util.WalkHookDirs(gitPath, util.CreateHook)
	if err != nil {
		log.Println("[ERROR]\t Could not create commit-msg script.")
		return 1
	}

	log.Println("[SUCCESS]\t commit-message-check successfully installed.")
	return 0
}
