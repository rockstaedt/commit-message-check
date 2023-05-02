package cmd

import (
	"rockstaedt/commit-message-check/util"
)

func (h *Handler) setup() int {
	err := util.WalkHookDirs(h.Config.GitPath, util.CreateHook)
	if err != nil {
		h.notify("Could not create commit-msg script.", "red")
		return 1
	}

	h.notify("commit-message-check successfully installed.", "green")
	return 0
}
