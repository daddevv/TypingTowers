package content

// ContentHandler defines the API for asset/content loading and resources.
type ContentHandler interface {
	Update(dt float64)
}

// ContentHandler manages asset/content loading and resources.
type Handler struct {
	// Add fields for content/resource state here
}

// NewHandler creates a new ContentHandler.
func NewHandler() *Handler {
	return &Handler{}
}

// Update updates the content/resource state.
func (h *Handler) Update(dt float64) {
	// Placeholder: update content/resource state here in the future.
}
