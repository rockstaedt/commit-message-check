package cmd

import "rockstaedt/commit-message-check/util"

func Uninstall(gitPath string) int {

	_ = util.WalkHookDirs(gitPath, util.DeleteHook)

	return 0
}
