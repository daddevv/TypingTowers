package phase

// PhaseHandler defines the API for game phase/state logic.
type PhaseHandler interface {
	Update(dt float64)
}

// PhaseHandler manages game phase/state logic.
type Handler struct {
	// Add fields for phase state here
}

// NewHandler creates a new PhaseHandler.
func NewHandler() *Handler {
	return &Handler{}
}

// Update updates the phase state.
func (h *Handler) Update(dt float64) {
	// Placeholder: update phase state here in the future.
}
