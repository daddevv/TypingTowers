package entity

import (
	"github.com/daddevv/type-defense/internal/core"
	"github.com/hajimehoshi/ebiten/v2"
)

// Entity interface defines the contract for all entities in the game.
type Entity interface {
	Frame() *ebiten.Image              // Get the current image frame of the entity (static or animated)
	Update(dt float64) error           // Update the entity state using delta time
	Draw(screen *ebiten.Image)         // Draw the entity to the screen
	Position() (x, y float64)          // Get the current position of the entity
	Bounds() (x, y, width, height int) // Get the bounding box of the entity
	Hitbox() (x, y, width, height int) // Get the hitbox of the entity
	Static() bool                      // Check if the entity is static (does not move)
	Alive() bool                       // Check if the entity is alive (for entities that can die)
	ApplyDamage(amount int)            // Apply damage to the entity
	Health() int                       // Get the current health of the entity (if applicable)
	Destroy()                          // Clean up resources when the entity is no longer needed
}

// BaseEntity provides common fields and methods for all entities.
type BaseEntity struct {
	Pos           core.Point    // Position of the entity
	Width, Height int           // Size of the entity
	Sprite        *ebiten.Image // Current image frame of the entity
	FrameAnchorX  float64       // Anchor point for the frame (for positioning)
	FrameAnchorY  float64       // Anchor point for the frame (for positioning)
	Static        bool          // Whether the entity is static or not
}

// Update updates the entity's state. Override this method in derived entities to implement specific behavior.
func (e *BaseEntity) Update(dt float64) error {
	return nil
}

// Destroy releases resources.
func (e *BaseEntity) Destroy() {
	e.Sprite = nil
}

// Draw draws the entity to the screen at its current position.
// The frame is drawn at the position (X, Y) with the anchor point adjusted by FrameAnchorX and FrameAnchorY.
func (e *BaseEntity) Draw(screen *ebiten.Image) {
	if e.Sprite != nil {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(e.Pos.X)-e.FrameAnchorX, float64(e.Pos.Y)-e.FrameAnchorY)
		screen.DrawImage(e.Sprite, op)
	}
}

// Bounds returns the bounding box of the entity.
func (e *BaseEntity) Bounds() (x, y, width, height int) {
	return int(e.Pos.X), int(e.Pos.Y), e.Width, e.Height
}

// Hitbox returns the hitbox of the entity (by default, same as bounds).
func (e *BaseEntity) Hitbox() (x, y, width, height int) {
	return e.Bounds()
}
