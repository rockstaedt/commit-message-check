package cmd

import (
	"github.com/rockstaedt/txtreader"
	"log"
	"rockstaedt/commit-message-check/internal/model"
)

type Handler struct {
	Config model.Config
}

func NewHandler(config model.Config) *Handler {
	return &Handler{Config: config}
}

func (h *Handler) Run() {

	switch h.Config.Command {
	case "setup":
		Setup(h.Config.GitPath)
	case "uninstall":
		Uninstall(h.Config.GitPath)
	case "update":
		Update(&h.Config)
	case "validate":
		commitLines, err := txtreader.GetLinesFromTextFile(h.Config.CommitMsg)
		if err != nil {
			log.Printf("Could not read commit message: %q", err.Error())
			break
		}

		Validate(commitLines)
	}
}
