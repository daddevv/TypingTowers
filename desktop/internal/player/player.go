package player

import (
	"td/internal/sprite"
)

const (
	defaultInitialHealth = 100 // Default initial health for the player
)

// Player represents the player character in the game.
type Player struct {
	Sprite map[string]sprite.Sprite // Player's sprite, can be animated or static
	Health int                      // Player's health
}

// NewPlayer creates a new Player instance with the given sprite and initial health.
func NewPlayer(sprite map[string]sprite.Sprite) *Player {
	return &Player{
		Sprite: sprite,
		Health: defaultInitialHealth,
	}
}
