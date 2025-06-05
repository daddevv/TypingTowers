package ui

// UiHandler defines the API for UI logic and state management.
type UiHandler interface {
	Update(dt float64)
}

// UIHandler manages all UI logic and state.
type Handler struct {
	// Add fields for UI state here
}

// NewHandler creates a new UIHandler.
func NewHandler() *Handler {
	return &Handler{}
}

// Update updates the UI state.
func (h *Handler) Update(dt float64) {
	// Placeholder: update UI state here in the future.
}
