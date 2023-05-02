package cmd

import (
	"io"
	"log"
	"rockstaedt/commit-message-check/internal/model"
)

type Handler struct {
	Config model.Config
	Writer io.Writer
}

func NewHandler(config model.Config) *Handler {
	return &Handler{Config: config}
}

func (h *Handler) Run(command string) int {
	var status int

	switch command {
	case "setup":
		status = h.setup()
	case "uninstall":
		status = h.uninstall()
	case "update":
		status = h.update()
	case "validate":
		status = h.validate()
	default:
		log.Println("Unknown subcommand. Please check manual.")
		return 1
	}

	return status
}

func (h *Handler) notify(message string) {
	h.Writer.Write([]byte(message))
}
