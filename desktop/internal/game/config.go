package game

import (
	"td/internal/entity"
)

type GameConfig struct {
	Name string `json:"level_name"` // Name of the level to load
	Background string `json:"level_background"` // Background image for the level
	Waves []WaveConfig `json:"waves"` // List of waves in the level
	Spawner entity.SpawnerConfig `json:"spawner"` // Configuration for the spawner
}