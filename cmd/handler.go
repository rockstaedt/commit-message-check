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

func (h *Handler) Run() int {

	var status int
	switch h.Config.Command {
	case "setup":
		status = Setup(h.Config.GitPath)
	case "uninstall":
		status = Uninstall(h.Config.GitPath)
	case "update":
		status = Update(&h.Config)
	case "validate":
		commitLines, err := txtreader.GetLinesFromTextFile(h.Config.CommitMsg)
		if err != nil {
			log.Printf("Could not read commit message: %q", err.Error())
			return 3
		}

		Validate(commitLines)
	default:
		log.Println("Unknown subcommand. Please check manual.")
		return 4
	}

	return status
}
