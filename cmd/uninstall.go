package cmd

import (
	"rockstaedt/commit-message-check/util"
)

func (h *Handler) uninstall() int {
	err := util.WalkHookDirs(h.Config.GitPath, util.DeleteHook)
	if err != nil {
		h.notify("Could not delete commit-msg hook.", "red")
		return 1
	}

	h.notify("commit-message-check successfully uninstalled.", "green")
	return 0
}
