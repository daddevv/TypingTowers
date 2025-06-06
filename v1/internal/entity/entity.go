package entity

import (
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
	Destroy()                          // Clean up resources when the entity is no longer needed
}

// BaseEntity provides common fields and methods for all entities.
type BaseEntity struct {
	pos           Point         // Position of the entity
	width, height int           // Size of the entity
	frame         *ebiten.Image // Current image frame of the entity
	frameAnchorX  float64       // Anchor point for the frame (for positioning)
	frameAnchorY  float64       // Anchor point for the frame (for positioning)
	static        bool          // Whether the entity is static or not
}

// Frame returns the current image frame.
func (e *BaseEntity) Frame() *ebiten.Image {
	return e.frame
}

// Update updates the entity's state. Override this method in derived entities to implement specific behavior.
func (e *BaseEntity) Update(dt float64) error {
	return nil
}

// Destroy releases resources.
func (e *BaseEntity) Destroy() {
	e.frame = nil
}

// Draw draws the entity to the screen at its current position.
// The frame is drawn at the position (X, Y) with the anchor point adjusted by frameAnchorX and frameAnchorY.
func (e *BaseEntity) Draw(screen *ebiten.Image) {
	if e.frame != nil {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(e.pos.X)-e.frameAnchorX, float64(e.pos.Y)-e.frameAnchorY)
		screen.DrawImage(e.frame, op)
	}
}

// Position returns the entity's position.
func (e *BaseEntity) Position() (x, y float64) {
	return e.pos.X, e.pos.Y
}

// Bounds returns the bounding box of the entity.
func (e *BaseEntity) Bounds() (x, y, width, height int) {
	return int(e.pos.X), int(e.pos.Y), e.width, e.height
}

// Hitbox returns the hitbox of the entity (by default, same as bounds).
func (e *BaseEntity) Hitbox() (x, y, width, height int) {
	return e.Bounds()
}

// Static returns whether the entity is static.
func (e *BaseEntity) Static() bool {
	return e.static
}
