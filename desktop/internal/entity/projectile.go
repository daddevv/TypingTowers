package entity

import (
	"td/internal/ui"

	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

// Projectile represents a projectile shot at a mob when a letter is typed correctly.
type Projectile struct {
	Position    ui.Location // Current position of the projectile
	Target      ui.Location // Target position (usually a mob's position)
	Speed       float64     // Speed of projectile movement
	Sprite      *ebiten.Image // Visual representation of the projectile
	Active      bool        // Whether the projectile is still active
	DamageDealt bool        // Whether damage has been dealt to prevent multiple hits
	TargetChar  rune        // The letter this projectile is intended to hit
	TargetMob   Entity      // The mob this projectile is homing in on
	// (Future: add hit/miss logic or accuracy here)
}

// NewProjectile creates a new projectile from start position to target position.
func NewProjectile(start ui.Location, target ui.Location, mob Entity) *Projectile {
	// Create a simple white circle sprite for the projectile
	sprite := ebiten.NewImage(8, 8)
	sprite.Fill(color.RGBA{255, 255, 255, 255})
	
	return &Projectile{
		Position:    start,
		Target:      target,
		Speed:       30.0, // pixels per frame
		Sprite:      sprite,
		Active:      true,
		DamageDealt: false,
		TargetChar:  0, // Set by input handler after creation
		TargetMob:   mob,
	}
}

// Update moves the projectile towards its target.
func (p *Projectile) Update() error {
	if !p.Active {
		return nil
	}

	// If homing, update target position to mob's current position
	if p.TargetMob != nil {
		p.Target = p.TargetMob.GetPosition()
	}

	// Calculate direction vector from position to target
	dx := p.Target.X - p.Position.X
	dy := p.Target.Y - p.Position.Y

	// Calculate distance to target
	distance := math.Sqrt(dx*dx + dy*dy)
	if distance < 1.0 { // Close enough to target
		p.Active = false
		return nil
	}

	// Normalize direction and apply speed
	if distance > 0 {
		moveX := (dx / distance) * p.Speed
		moveY := (dy / distance) * p.Speed

		// Clamp movement if we're closer than speed to the target
		if math.Abs(moveX) > math.Abs(dx) {
			moveX = dx
		}
		if math.Abs(moveY) > math.Abs(dy) {
			moveY = dy
		}

		p.Position.X += moveX
		p.Position.Y += moveY
	}

	// For now, projectiles always hit if they reach the target (100% hit rate)
	// (Future: add miss chance or accuracy logic here)

	return nil
}

// Draw renders the projectile on the screen.
func (p *Projectile) Draw(screen *ebiten.Image) {
	if !p.Active {
		return
	}
	
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(p.Position.X, p.Position.Y)
	screen.DrawImage(p.Sprite, opts)
}

// GetPosition returns the current position of the projectile.
func (p *Projectile) GetPosition() ui.Location {
	return p.Position
}

// SetPosition sets the projectile's position.
func (p *Projectile) SetPosition(x, y float64) {
	p.Position.X = x
	p.Position.Y = y
}

// IsActive returns whether the projectile is still active.
func (p *Projectile) IsActive() bool {
	return p.Active
}

// Deactivate marks the projectile as inactive.
func (p *Projectile) Deactivate() {
	p.Active = false
}
