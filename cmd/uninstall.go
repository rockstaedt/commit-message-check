package cmd

import (
	"log"
	"rockstaedt/commit-message-check/util"
)

func Uninstall(gitPath string) int {

	_ = util.WalkHookDirs(gitPath, util.DeleteHook)

	log.Println("[SUCCESS]\t Deleted all hook files.")
	return 0
}
