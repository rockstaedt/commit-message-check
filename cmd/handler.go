package cmd

import "rockstaedt/commit-message-check/internal/model"

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
	}
}
