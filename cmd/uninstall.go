package cmd

import (
	"log"
	"rockstaedt/commit-message-check/util"
)

func Uninstall(gitPath string) int {

	err := util.WalkHookDirs(gitPath, util.DeleteHook)
	if err != nil {
		return 1
	}

	log.Println("[SUCCESS]\t Deleted all hook files.")
	return 0
}
