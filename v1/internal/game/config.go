package game

import (
	"encoding/json"
	"os"
)

// ConfigFile is the default path for configuration data.
const ConfigFile = "config.json"

// Config holds tunable parameters for balancing and upgrades.
type Config struct {
	A float64
	B float64
	C float64
	D float64
	E float64
	F float64 // tower fire rate multiplier
	G float64
	H float64
	I float64
	J float64 // base starting health
	K float64
	L float64
	M float64
	N float64 // mob health growth per wave
	O float64
	P float64
	Q float64
	R float64
	S float64
	T float64
	U float64
	V float64
	W float64
	X float64
	Y float64
	Z float64

	// Explicitly named parameters used throughout the game. All values are
	// optional when loading from JSON; zero values fall back to defaults.
	TowerDamage       int     `json:"tower_damage"`
	TowerRange        float64 `json:"tower_range"`
	TowerFireRate     int     `json:"tower_fire_rate"`
	TowerReloadRate   int     `json:"tower_reload_rate"` // in frames
	TowerAmmoCapacity int     `json:"tower_ammo_capacity"`
	TowerProjectiles  int     `json:"tower_projectiles_per_shot"`
	TowerBounce       int     `json:"tower_bounce_count"`

	ProjectileSpeed float64 `json:"projectile_speed"`

	BaseHealth int `json:"base_health"`

	MobSpeed       float64 `json:"mob_speed"`
	MobBaseHealth  int     `json:"mob_base_health"`
	MobsPerWave    int     `json:"mobs_per_wave_base"`
	MobsPerWaveInc int     `json:"mobs_per_wave_growth"`
	SpawnInterval  int     `json:"spawn_interval"`
}

// DefaultConfig provides baseline parameters used when a new game starts.
var DefaultConfig = Config{
	A: 1,
	B: 1,
	C: 1,
	D: 1,
	E: 1,
	F: 1,
	G: 1,
	H: 1,
	I: 1,
	J: 10,
	K: 1,
	L: 1,
	M: 1,
	N: 0.5,
	O: 1,
	P: 1,
	Q: 1,
	R: 1,
	S: 1,
	T: 1,
	U: 1,
	V: 1,
	W: 1,
	X: 1,
	Y: 1,
	Z: 1,

	// Defaults for the new explicit parameters
	TowerDamage:       1,
	TowerRange:        500,
	TowerFireRate:     100,
	TowerAmmoCapacity: 5,
	TowerProjectiles:  1,
	TowerBounce:       0,

	ProjectileSpeed: 5.0,

	BaseHealth: 10,

	MobSpeed:       1.0,
	MobBaseHealth:  1,
	MobsPerWave:    3,
	MobsPerWaveInc: 3,
	SpawnInterval:  60,
}

// LoadConfig reads configuration values from the given JSON file.
// If the file cannot be read or parsed, DefaultConfig is returned.
func LoadConfig(path string) (Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return DefaultConfig, err
	}
	cfg := DefaultConfig
	if err := json.Unmarshal(data, &cfg); err != nil {
		return DefaultConfig, err
	}
	return cfg, nil
}
