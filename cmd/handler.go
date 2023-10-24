package cmd

import (
	"github.com/TwiN/go-color"
	"io"
	"log"
	"rockstaedt/commit-message-check/internal/model"
)

type Handler struct {
	Config model.Config
	Writer io.Writer
	Reader io.Reader
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

func (h *Handler) notify(message string, txtColor ...string) {
	if len(txtColor) > 0 && txtColor[0] == "green" {
		message = color.InGreen(message)
	}

	if len(txtColor) > 0 && txtColor[0] == "red" {
		message = color.InRed(message)
	}

	if len(txtColor) > 0 && txtColor[0] == "yellow" {
		message = color.InYellow(message)
	}

	_, err := h.Writer.Write([]byte(message + "\n"))
	if err != nil {
		log.Println("Error at writing!")
	}
}
