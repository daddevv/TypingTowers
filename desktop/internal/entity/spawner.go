package entity

import "td/internal/content"

type BaseSpawner struct {
	// Core timing
	MinSpawnInterval float64 `json:"min_spawn_interval"` // Minimum frames between spawns
	MaxSpawnInterval float64 `json:"max_spawn_interval"` // Maximum frames between spawns
	NextSpawnTime    float64 `json:"next_spawn_time"`    // Time until next spawn (in seconds)
	ElapsedTime      float64 `json:"elapsed_time"`      // Total elapsed time since spawner creation

	// Spawn configuration
	LetterPool   LetterPool `json:"letter_pool"`   // LetterPool for dynamic letter sets
	WordPool     []string   `json:"word_pool"`     // Pool of specific words that can be spawned
	RandomWeight float64    `json:"random_weight"` // Weight for spawning random letter mobs (0.0-1.0)
	WordWeight   float64    `json:"word_weight"`   // Weight for spawning word mobs (0.0-1.0)

	// Letter count for random mobs
	MinLetterCount int `json:"min_letter_count"` // Minimum letters for random mobs
	MaxLetterCount int `json:"max_letter_count"` // Maximum letters for random mobs

	// Per-mob-type letter count config
	MobTypeConfigs map[string]MobConfig `json:"mob_type_configs"` // Config for each mob type
	ConfigDelay    float64 `json:"config_delay"`    // Delay in frames before the spawner starts
}

// NewBaseSpawner creates a new BaseSpawner with default settings and a LetterPool.
func NewBaseSpawnerWithConfigs(letterPool LetterPool, mobConfigs []content.MobConfig) *BaseSpawner {
	mobConfigMap := make(map[string]MobConfig)
	for _, mob := range mobConfigs {
		mobConfigMap[mob.Name] = MobConfig{
			MinLetters: mob.MinLetters,
			MaxLetters: mob.MaxLetters,
		}
	}
	spawner := &BaseSpawner{
		MinSpawnInterval: 100.0,
		MaxSpawnInterval: 200.0,
		LetterPool:       letterPool,
		RandomWeight:     1.0,
		WordWeight:       0.0,
		MinLetterCount:   2,
		MaxLetterCount:   4,
		MobTypeConfigs:   mobConfigMap,
		ConfigDelay:      30.0, // Start after 30 frames
	}
	return spawner
}