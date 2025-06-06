package econ

// EconomyHandler defines the API for economy management like resources and experience.
type EconomyHandler interface {
	Update(dt float64)
}

// Handler manages resource and economy state.
type Handler struct {
	// Add fields for economy state here
}

// NewHandler creates a new economy handler.
func NewHandler() *Handler {
	return &Handler{}
}

// Update updates the economy state.
func (h *Handler) Update(dt float64) {
	// Placeholder: update economy state here in the future.
}
