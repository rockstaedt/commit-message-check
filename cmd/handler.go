package cmd

import "rockstaedt/commit-message-check/internal/model"

type Handler struct {
	Config model.Config
}

func NewHandler(config model.Config) *Handler {
	return &Handler{Config: config}
}

func (h *Handler) Run() {
	Uninstall(h.Config.GitPath)
}
