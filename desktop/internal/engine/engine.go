package engine

import (
	"fmt"
	"os"
	"path/filepath"
	"td/internal/content"
	"td/internal/game"
	"td/internal/state"
	"td/internal/ui"

	"github.com/hajimehoshi/ebiten/v2"
	lua "github.com/yuin/gopher-lua"
)

// Engine represents the main game engine that manages the game state, menus, and rendering.
type Engine struct {
	Version 	 string // Version of the game engine, used for debugging and logging
	State   	 state.EngineState // enum: current screen to display
	Screen  	 *ebiten.Image // Internal buffer for drawing at 1920x1080
	MainMenu 	 *ui.MainMenu // Main menu instance
	// PauseMenu 	 *ui.PauseMenu // Pause menu instance
	Game    	 *game.Game // Current active game, can be nil if no game is active
	GameConfig 	 *game.GameConfig // Configuration for the game
	L      	 	 *lua.LState // Lua state for scripting and game logic
}

// NewEngine initializes a new game engine instance.
func NewEngine(version string) *Engine {
	// Load all content configs at engine startup
	err := content.LoadContentConfigs()
	if err != nil {
		panic("Failed to load content configs: " + err.Error())
	}

	// Initialize Lua state
	l := lua.NewState()
	l.SetGlobal("example", l.NewFunction(func(L *lua.LState) int {
		fmt.Println("'example' called with parameters:", L.ToString(1), L.ToNumber(2), L.ToNumber(3))
		return 0
	}))

	return &Engine{
		Version:      version,
		Screen:      ebiten.NewImage(1920, 1080),
		MainMenu:   nil,
		// PauseMenu:  nil,
		Game:       nil,
		GameConfig: nil,
		L:         l,
	}
}

// Update calls the appropriate update method based on the current state of the engine.
func (e *Engine) Update() error {
	switch e.State {
	case state.MENU_MAIN:
		if e.MainMenu == nil {
			e.MainMenu = ui.NewMainMenu(e.L)
		}
		if selection, err := e.MainMenu.Update(); err != nil {
			return fmt.Errorf("failed to update main menu: %w", err)
		} else if selection != nil {
			e.handleMainMenuSelection(*selection)
		}
	case state.GAME_PLAYING:
		if e.Game == nil || e.GameConfig == nil {
			return fmt.Errorf("game is not initialized")
		}
	}
	return nil
}

// Draw renders the current state to the screen.
func (e *Engine) Draw(screen *ebiten.Image) {
	// Draw everything to the internal 1920x1080 screen
	e.Screen.Clear()
	switch e.State {
	case state.MENU_MAIN:
		e.MainMenu.Draw(e.Screen)
	case state.GAME_PLAYING:
		e.Game.Draw(e.Screen)
	case state.GAME_PAUSED:
		e.State = state.MENU_MAIN // Temporarily set to GAME_PLAYING for drawing
		// e.PauseMenu.Draw(e.Screen)
	}
	// Now scale the internal screen to the window size
	w, h := screen.Bounds().Dx(), screen.Bounds().Dy()
	scaleX := float64(w) / 1920.0
	scaleY := float64(h) / 1080.0
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(scaleX, scaleY)
	screen.DrawImage(e.Screen, opts)
}

// Layout defines the size of the internal canvas used by the engine.
// This is always 1920x1080, regardless of the actual window size.
func (g *Engine) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	// Always use 1920x1080 for the internal canvas
	return 1920, 1080
}

// loadLuaPlugins loads all Lua scripts from the specified directory into the Lua state.
func loadLuaPlugins(L *lua.LState, dir string) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return
	}
	for _, f := range files {
		if filepath.Ext(f.Name()) == ".lua" {
			script, err := os.ReadFile(filepath.Join(dir, f.Name()))
			if err == nil {
				L.DoString(string(script))
			}
		}
	}
}

// handleMainMenuSelection processes the selection made in the main menu.
func (e *Engine) handleMainMenuSelection(selection ui.MainMenuOption) {
	switch selection {
	case ui.MainMenuStartGame:
		if e.GameConfig == nil {
			e.GameConfig = game.NewGameConfig()
		}
		e.Game = game.NewGame(e.GameConfig, e.L)
		e.State = state.GAME_PLAYING
	case ui.MainMenuOptions:
		// Show options menu
		fmt.Println("Options menu not implemented yet")
	case ui.MainMenuQuit:
		// Quit the game
		fmt.Println("Quitting game...")
		os.Exit(0)
	}
}