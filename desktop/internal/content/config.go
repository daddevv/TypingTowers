// Package content provides types for content configuration (levels, mobs, worlds).
package content

// LevelConfig defines the structure for a level configuration.
type LevelConfig struct {
	Name            string   `json:"name"`
	Difficulty      string   `json:"difficulty"`
	Biome           string   `json:"biome"`
	PossibleLetters []string `json:"possibleLetters"`
	Background      string   `json:"background"`
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
