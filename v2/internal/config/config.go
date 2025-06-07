package config

// DefaultConfig provides the default game configuration.
var DefaultConfig = GameConfig{
	ScreenWidth:  800,
	ScreenHeight: 600,
	Fullscreen:   false,
	Volume:       1.0, // Default volume level
}

// ConfigFile defines the name of the configuration file.
const ConfigFile = "config.json"

// GameConfig holds the configuration settings for the game.
type GameConfig struct {
	ScreenWidth  int     // ScreenWidth defines the width of the game screen in pixels.
	ScreenHeight int     // ScreenHeight defines the height of the game screen in pixels.
	Fullscreen   bool    // Fullscreen defines whether the game is in fullscreen mode.
	Volume       float64 // Volume defines the audio volume level.
}

func ConfigFilePath() string {
	return ConfigFile
}

// LoadConfig loads the game configuration from a file.
func LoadConfig(filePath string) (*GameConfig, error) {
	// Implementation goes here
	return &DefaultConfig, nil
}

// SaveConfig saves the game configuration to a file.
func SaveConfig(cfg *GameConfig) error {
	// Implementation goes here
	// This function should serialize the cfg to a file at ConfigFilePath()
	return nil
}
