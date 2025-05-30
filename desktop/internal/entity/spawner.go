package entity

type SpawnerConfig struct {
	Delay float64 `json:"delay"` // Delay in frames before the spawner starts
	MinSpawnInterval float64 `json:"min_spawn_interval"` // Minimum frames between spawns
	MaxSpawnInterval float64 `json:"max_spawn_interval"` // Maximum frames between spawns
}
	