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
