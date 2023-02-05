package cmd

import (
	"log"
	"rockstaedt/commit-message-check/util"
)

func Uninstall(gitPath string) int {
	err := util.WalkHookDirs(gitPath, util.DeleteHook)
	if err != nil {
		log.Println("[ERROR]\t Could not delete commit-msg hook.")
		return 1
	}

	log.Println("[SUCCESS]\t commit-message-check successfully uninstalled.")
	return 0
}
