package structure

// TowerHandler defines the API for tower logic and state management.
type TowerHandler interface {
	Update(dt float64)
}

// TowerHandler manages all tower logic and state.
type Handler struct {
	// Add fields for tower state here
}

// NewHandler creates a new TowerHandler.
func NewHandler() *Handler {
	return &Handler{}
}

// Update updates the tower state.
func (h *Handler) Update(dt float64) {
	// Placeholder: update all towers here in the future.
}
