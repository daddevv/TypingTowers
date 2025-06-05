package sprite

// SpriteHandler defines the API for sprite/image logic and helpers.
type SpriteHandler interface {
	Update(dt float64)
}

// SpriteHandler manages sprite/image logic and helpers.
type Handler struct {
	// Add fields for sprite state here
}

// NewHandler creates a new SpriteHandler.
func NewHandler() *Handler {
	return &Handler{}
}

// Update updates the sprite state.
func (h *Handler) Update(dt float64) {
	// Placeholder: update sprite/image state here in the future.
}
