package game

import (
	"td/internal/entity"
	"td/internal/world"
)

type GameConfig struct {
	Name string `json:"level_name"` // Name of the level to load
	Background string `json:"level_background"` // Background image for the level
	Level world.Level `json:"level"` // World level configuration
	Waves []WaveConfig `json:"waves"` // List of waves in the level
	Spawner entity.BaseSpawner `json:"spawner"` // Configuration for the spawner
}

func NewGameConfig() *GameConfig {
	return &GameConfig{
		Name: "Default Level",
		Background: "default_background.png",
		Waves: []WaveConfig{
			{
				PossibleLetters: []string{"A", "B", "C"},
				MinLength: 3,
				MaxLength: 5,
				PossibleEnemies: []entity.Mob{},
				EnemyCount: 10,
			},
		},
		Level: world.Level{
			Name: "Default Level",
			WorldNumber: 1,
			LevelNumber: 1,
			World: world.BEACH,
			StartingLetters: []string{"A", "B", "C", "D"},
			Waves: nil,
			EnemyCount: 10,
			MobChances: nil,
			LevelCompleteScore: 1000,
			Background: "default_background.png",
		},
		Spawner: entity.BaseSpawner{
			MinSpawnInterval: 1.0,
			MaxSpawnInterval: 3.0,
			NextSpawnTime: 0.0,
			ElapsedTime: 0.0,
			LetterPool: entity.NewDefaultLetterPool(),
			WordPool: []string{"ABCD", "EFGH", "IJKL"},
			RandomWeight: 0.5,
			WordWeight: 0.5,
			MinLetterCount: 3,
			MaxLetterCount: 5,
			MobTypeConfigs: nil,
			ConfigDelay: 0.0,
		},
	}
}