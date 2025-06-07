package entity

import "github.com/daddevv/type-defense/internal/event"

// EntityHandler defines the API for entity logic and state management.
type EntityHandler interface {
	Update(dt float64)
}

// EntityHandler manages all entity logic and state.
type Handler struct {
	// Events is the outbound channel for entity-related events.
	// Consumers can subscribe via the central event bus.
	Events chan event.Event
}

// NewHandler creates a new EntityHandler.
func NewHandler() *Handler {
	return &Handler{Events: make(chan event.Event, 8)}
}

// Update updates all entities managed by the handler.
func (h *Handler) Update(dt float64) {
	// Placeholder: update all entities here in the future.
	// Example: for _, e := range h.entities { e.Update(dt) }
}
