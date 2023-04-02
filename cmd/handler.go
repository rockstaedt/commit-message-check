package cmd

type Handler struct {
	Command string
}

func NewHandler(command string) *Handler {
	return &Handler{Command: command}
}

func (h *Handler) Run() {

}
