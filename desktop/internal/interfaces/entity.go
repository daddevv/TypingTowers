package i

import "td/internal/physics"

type Entity interface {
	// ID returns the unique identifier for the entity.
	ID() string
	// Type returns the type of the entity.
	Type() string
	// Position returns the position of the entity as a Vec2.
	Position() physics.Vec2
	// SetPosition sets the position of the entity.
	SetPosition(pos physics.Vec2)
}
