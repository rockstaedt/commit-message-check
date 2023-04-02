package cmd

import "rockstaedt/commit-message-check/internal/model"

type Handler struct {
	Config model.Config
}

func NewHandler(config model.Config) *Handler {
	return &Handler{Config: config}
}

func (h *Handler) Run() {

	if h.Config.Command == "setup" {
		Setup(h.Config.GitPath)
	}

	Uninstall(h.Config.GitPath)
}
