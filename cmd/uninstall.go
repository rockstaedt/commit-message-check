package cmd

import (
	"log"
	"rockstaedt/commit-message-check/util"
)

func (h *Handler) uninstall() int {
	err := util.WalkHookDirs(h.Config.GitPath, util.DeleteHook)
	if err != nil {
		log.Println("[ERROR]\t Could not delete commit-msg hook.")
		return 1
	}

	log.Println("[SUCCESS]\t commit-message-check successfully uninstalled.")
	return 0
}
