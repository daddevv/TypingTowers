package game

import (
	"td/internal/content"
	"td/internal/world"
)

type GameOptions struct {
	GameMode   GameMode
	Level      world.Level
	MobConfigs []content.MobConfig // Add mob configs to options
}

func NewGameOptions() *GameOptions {
	return &GameOptions{
		GameMode:   ENDLESS,
		Level:      world.Level{}, // Initialize with an empty level
		// MobConfigs: content.GetAllMobConfigs(), // Load all mob configs from content
	}
}