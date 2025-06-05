package game

// Event is a generic interface for all events.
type Event interface{}

// EntityEvent represents events related to entities (e.g., spawn, death).
type EntityEvent struct {
	Type    string
	Payload interface{}
}

// UIEvent represents UI-related events (e.g., notification, panel toggle).
type UIEvent struct {
	Type    string
	Payload interface{}
}

// TechEvent represents tech tree unlocks or changes.
type TechEvent struct {
	Type    string
	Payload interface{}
}

// TowerEvent represents tower-specific events (e.g., upgrade, fire).
type TowerEvent struct {
	Type    string
	Payload interface{}
}

// PhaseEvent represents game phase/state transitions.
type PhaseEvent struct {
	Type    string
	Payload interface{}
}

// ContentEvent represents asset/content loading events.
type ContentEvent struct {
	Type    string
	Payload interface{}
}

// SpriteEvent represents sprite/image-related events.
type SpriteEvent struct {
	Type    string
	Payload interface{}
}

// Each handler should expose a channel for its event type for pub/sub (T-007).
// Example: EntityHandler exposes EntityEvents chan Event, etc.

// Example: Add more event types as needed for other modules.
