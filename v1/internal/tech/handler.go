package tech

// TechHandler defines the API for tech tree logic and state management.
type TechHandler interface {
	Update(dt float64)
}

// TechHandler manages all tech tree logic and state.
type Handler struct {
	// Add fields for tech state here
}

// NewHandler creates a new TechHandler.
func NewHandler() *Handler {
	return &Handler{}
}

// Update updates the tech state.
func (h *Handler) Update(dt float64) {
	// Placeholder: update tech tree state here in the future.
}
