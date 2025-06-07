package config

// DefaultConfig provides the default game configuration.
var DefaultConfig = GameConfig{
	ScreenWidth:  800,
	ScreenHeight: 600,
	Fullscreen:   false,
	Volume:       1.0, // Default volume level

	TileMapWidth:      20, // Default tile map width in tiles
	TileMapHeight:     15, // Default tile map height in tiles
	TileMapTopMargin:  10, // Default top margin of the tile map in pixels
	TileMapLeftMargin: 10, // Default left margin of the tile map in pixels
	TileSize:          32, // Default size of each tile in pixels
}

// ConfigFile defines the name of the configuration file.
const ConfigFile = "config.json"

// GameConfig holds the configuration settings for the game.
type GameConfig struct {
	ScreenWidth  int     // ScreenWidth defines the width of the game screen in pixels.
	ScreenHeight int     // ScreenHeight defines the height of the game screen in pixels.
	Fullscreen   bool    // Fullscreen defines whether the game is in fullscreen mode.
	Volume       float64 // Volume defines the audio volume level.

	TileMapWidth      int // TileMapWidth defines the width of the tile map in tiles.
	TileMapHeight     int // TileMapHeight defines the height of the tile map in tiles.
	TileMapTopMargin  int // TileMapTopMargin defines the top margin of the tile map in pixels.
	TileMapLeftMargin int // TileMapLeftMargin defines the left margin of the tile map in pixels.
	TileSize          int // TileSize defines the size of each tile in pixels.
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
