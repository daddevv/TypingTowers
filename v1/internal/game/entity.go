package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Entity interface defines the contract for all entities in the game.
type Entity interface {
	Frame() *ebiten.Image              // Get the current image frame of the entity (static or animated)
	Update() error                     // Update the entity state, return an error if something goes wrong
	Draw(screen *ebiten.Image)         // Draw the entity to the screen
	Position() (x, y int)              // Get the current position of the entity
	Bounds() (x, y, width, height int) // Get the bounding box of the entity
	Hitbox() (x, y, width, height int) // Get the hitbox of the entity
	Static() bool                      // Check if the entity is static (does not move)
	StaticEntity() *StaticEntity       // Get the static entity representation if applicable
	DynamicEntity() *DynamicEntity     // Get the dynamic entity representation if applicable
	Destroy()                          // Clean up resources when the entity is no longer needed
}

// BaseEntity provides common fields and methods for all entities.
type BaseEntity struct {
	X, Y          int           // Position of the entity
	Width, Height int           // Size of the entity
	frame         *ebiten.Image // Current image frame of the entity
	frameAnchorX  float64       // Anchor point for the frame (for positioning)
	frameAnchorY  float64       // Anchor point for the frame (for positioning)
}

// Frame returns the current image frame.
func (e *BaseEntity) Frame() *ebiten.Image {
	return e.frame
}

// Update does nothing for BaseEntity.
func (e *BaseEntity) Update() error {
	return nil
}

// Destroy releases resources.
func (e *BaseEntity) Destroy() {
	e.frame = nil
}

// Draw draws the entity at its position.
func (e *BaseEntity) Draw(screen *ebiten.Image) {
	if e.frame != nil {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(e.X)-e.frameAnchorX, float64(e.Y)-e.frameAnchorY)
		screen.DrawImage(e.frame, op)
	}
}

// Position returns the entity's position.
func (e *BaseEntity) Position() (x, y int) {
	return e.X, e.Y
}

// Bounds returns the bounding box of the entity.
func (e *BaseEntity) Bounds() (x, y, width, height int) {
	return e.X, e.Y, e.Width, e.Height
}

// Hitbox returns the hitbox of the entity (by default, same as bounds).
func (e *BaseEntity) Hitbox() (x, y, width, height int) {
	return e.Bounds()
}

// Static returns false for BaseEntity.
func (e *BaseEntity) Static() bool {
	return false
}

// StaticEntity returns nil for BaseEntity.
func (e *BaseEntity) StaticEntity() *StaticEntity {
	return nil
}

// StaticEntity represents a non-moving entity.
type StaticEntity struct {
	BaseEntity
}

// Static returns true for StaticEntity.
func (e *StaticEntity) Static() bool {
	return true
}

// StaticEntity returns itself.
func (e *StaticEntity) StaticEntity() *StaticEntity {
	return e
}

// DynamicEntity represents a moving entity.
func (e *StaticEntity) DynamicEntity() *DynamicEntity {
	return nil
}

// DynamicEntity represents an entity that can move.
type DynamicEntity struct {
	BaseEntity
	VelX, VelY float64 // Velocity components
}

// Update updates the position based on velocity.
func (e *DynamicEntity) Update() error {
	e.X += int(e.VelX)
	e.Y += int(e.VelY)
	return nil
}

// Static returns false for DynamicEntity.
func (e *DynamicEntity) Static() bool {
	return false
}

// StaticEntity returns nil for DynamicEntity.
func (e *DynamicEntity) StaticEntity() *StaticEntity {
	return nil
}

// DynamicEntity returns itself.
func (e *DynamicEntity) DynamicEntity() *DynamicEntity {
	return e
}
