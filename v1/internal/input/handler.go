package input

// InputHandler defines the API for input handling logic and state management.
type InputHandler interface {
	Update(dt float64)
}

// InputHandler manages all input logic and state.
type Handler struct {
	// Add fields for input state here
}

// NewHandler creates a new InputHandler.
func NewHandler() *Handler {
	return &Handler{}
}

// Update updates the input state.
func (h *Handler) Update(dt float64) {
	// Placeholder: update all input states here in the future.
}
