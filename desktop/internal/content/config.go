// Package content provides types for content configuration (levels, mobs, worlds).
package content

// MobSpawnChance defines the spawn chance for a specific mob type.
type MobSpawnChance struct {
	Type   string  `json:"type"`
	Chance float64 `json:"chance"` // 0.0-1.0, e.g. 1.0 = 100%
}

// LevelWaveConfig defines a wave in a level.
type LevelWaveConfig struct {
	WaveNumber      int      `json:"waveNumber"`
	PossibleLetters []string `json:"possibleLetters"`
	EnemyCount      int      `json:"enemyCount"` // Number of enemies to defeat in this wave
	MobChances      []MobSpawnChance `json:"mobChances"` // Mob spawn % for this wave
}

// LevelConfig defines the structure for a level configuration.
type LevelConfig struct {
	Name            string            `json:"name"` // Fun display name
	WorldNumber     int               `json:"worldNumber"`
	LevelNumber     int               `json:"levelNumber"`
	World           string            `json:"world"` // Reference to WorldConfig.Name
	StartingLetters []string          `json:"startingLetters"`
	Waves           []LevelWaveConfig `json:"waves"`
	LevelCompleteScore int            `json:"levelCompleteScore"` // Score required to complete the level
}

// MobConfig defines the structure for a mob configuration.
type MobConfig struct {
	Type           string  `json:"type"`
	SpriteSheet    string  `json:"spriteSheet"`
	FrameRows      int     `json:"frameRows"`
	FrameCols      int     `json:"frameCols"`
	FrameWidth     int     `json:"frameWidth"`
	FrameHeight    int     `json:"frameHeight"`
	FrameDuration  int     `json:"frameDuration"`
	DefaultSpeed   float64 `json:"defaultSpeed"`
	LetterFont     string  `json:"letterFont"`
	LetterFontSize int     `json:"letterFontSize"`
	MinLetters     int     `json:"minLetters"`
	MaxLetters     int     `json:"maxLetters"`
}

// WorldConfig defines the structure for a world configuration.
type WorldConfig struct {
	Name       string `json:"name"`
	Background string `json:"background"`
}
