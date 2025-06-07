package core

import (
	"typingtowers/internal/config"
	"typingtowers/internal/tiles"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Config *config.GameConfig // Config holds the game configuration.
	Tiles  *tiles.TileMap     // Tiles holds the tile map for the game.
	// Add other fields as necessary, such as game state, player data, etc.
}

// NewGame initializes a new Game instance with the default configuration.
func NewGame() *Game {
	return NewGameWithConfig(&config.DefaultConfig)
}

// NewGameWithConfig initializes a new Game instance with the provided configuration.
func NewGameWithConfig(cfg *config.GameConfig) *Game {
	return &Game{
		Config: cfg,
		Tiles:  tiles.NewTileMap(cfg.TileMapWidth, cfg.TileMapHeight, cfg.TileMapTopMargin, cfg.TileMapLeftMargin, cfg.TileSize),
		// Initialize other fields as necessary.
	}
}

// Update updates the game state. This method will be called on each frame.
func (g *Game) Update() error {
	if g.Tiles.Tiles[0][0] == nil {
		// Initialize the tile map if it is not already initialized.
		g.Tiles.InitializeTiles()
	}
	// Update game logic here.
	// This is a placeholder for the game update logic.
	// For example, you can update player positions, check for collisions, etc.
	return nil
}

// Draw draws the game on the screen. This method will be called on each frame.
func (g *Game) Draw(screen *ebiten.Image) {
	// Draw the game state to the screen.
	// This is a placeholder for the game drawing logic.
	// For example, you can draw the player, enemies, background, etc.
	g.Tiles.Draw(screen)
}

// Layout returns the layout of the game screen.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	// Return the game screen dimensions based on the configuration.
	// This is a placeholder for the layout logic.
	return g.Config.ScreenWidth, g.Config.ScreenHeight
}

// IsFullscreen returns whether the game is in fullscreen mode.
func (g *Game) IsFullscreen() bool {
	return g.Config.Fullscreen
}

// SetFullscreen sets the fullscreen mode for the game.
func (g *Game) SetFullscreen(fullscreen bool) {
	g.Config.Fullscreen = fullscreen
	// Additional logic to handle fullscreen mode can be added here.
}

// Volume returns the current volume level of the game.
func (g *Game) Volume() float64 {
	return g.Config.Volume
}

// SetVolume sets the volume level for the game.
func (g *Game) SetVolume(volume float64) {
	g.Config.Volume = volume
	// Additional logic to handle volume changes can be added here.
}

// SaveConfig saves the current game configuration to a file.
func (g *Game) SaveConfig() error {
	return config.SaveConfig(g.Config)
}

// LoadConfig loads the game configuration from a file.
func (g *Game) LoadConfig() error {
	cfg, err := config.LoadConfig(config.ConfigFilePath())
	if err != nil {
		return err
	}
	g.Config = cfg
	return nil
}
