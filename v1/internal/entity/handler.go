package entity

// EntityHandler defines the API for entity logic and state management.
type EntityHandler interface {
	Update(dt float64)
}

// EntityHandler manages all entity logic and state.
type Handler struct {
	// Add fields for entity state here
}

// NewHandler creates a new EntityHandler.
func NewHandler() *Handler {
	return &Handler{}
}

// Update updates all entities managed by the handler.
func (h *Handler) Update(dt float64) {
	// Placeholder: update all entities here in the future.
	// Example: for _, e := range h.entities { e.Update(dt) }
}
