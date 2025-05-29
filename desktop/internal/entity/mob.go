package entity

import (
	"td/internal/ui"

	"github.com/hajimehoshi/ebiten/v2"
)

type Mob struct {
	Sprite   *ebiten.Image // Current frame to draw
	Position ui.Location // Position of the mob on the screen
	MoveTarget ui.Location // Target position for the mob to move towards
	Speed    float64      // Speed of the mob's movement
	IdleAnimation *ui.Animation // Animation for idle state
	MoveAnimation *ui.Animation // Animation for moving state
	Letters  []Letter      // Letters to display above the mob for player to type
	WordWidth float64      // Width of the word formed by letters
	Dead     bool         // Whether the mob is dead (for removal)
	DeathTimer int        // Frames left for death animation
}