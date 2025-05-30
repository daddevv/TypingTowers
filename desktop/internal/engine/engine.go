package engine

import (
	"fmt"
	"os"
	"path/filepath"
	"td/internal/game"
	"td/internal/menu"
	"td/internal/state"
	"td/internal/ui"
	"td/internal/world"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	lua "github.com/yuin/gopher-lua"
)

type Engine struct {
	isGameActive bool // Flag to indicate if the game is currently active
	WindowWidth   int
	WindowHeight  int
	Version string
	State   state.EngineState // enum: current screen to display
	Screen  *ebiten.Image // Internal buffer for drawing at 1920x1080
	Menu    ui.Screen
	Game    ui.Screen
	L      *lua.LState
}

func NewEngine(width, height int, version string) *Engine {
	// Load all content configs at engine startup
	err := LoadContentConfigs()
	if err != nil {
		panic("Failed to load content configs: " + err.Error())
	}

	// Initialize Lua state
	l := lua.NewState()
	l.SetGlobal("example", l.NewFunction(func(L *lua.LState) int {
		fmt.Println("'example' called with parameters:", L.ToString(1), L.ToNumber(2), L.ToNumber(3))
		return 0
	}))

	m := menu.NewMainMenu(l)
	return &Engine{
		isGameActive: false,
		Version:      version,
		WindowHeight: 1080, // Always use 1920x1080 for internal canvas
		WindowWidth:  1920,
		Screen:      ebiten.NewImage(1920, 1080),
		State:       state.MAIN_MENU,
		Menu:       m,
		Game:       nil,
		L:         l,
	}
}

// Update calls the appropriate update method based on the current state of the engine.
func (e *Engine) Update() error {
	// If the game is active, update the game screen
	if e.isGameActive { 
		if e.Game == nil {
			e.NewGame()
		}
		if err := e.Game.Update(); err != nil {
			return fmt.Errorf("failed to update game: %w", err)
		}
	// Otherwise, update the menu screen
	} else { 
		if err := e.Menu.Update(); err != nil {
			return fmt.Errorf("failed to update menu: %w", err)
		}
	}
	return nil
}

// Draw renders the current screen to the provided canvas.
func (g *Engine) Draw(screen *ebiten.Image) {
	// Draw everything to the internal 1920x1080 screen
	g.Screen.Clear()
	switch g.State {
	case state.MAIN_MENU:
		g.Menu.Draw(g.Screen)
	case state.GAME_PLAYING:
		g.Game.Draw(g.Screen)
	default:
		// Handle default drawing
	}
	// Now scale the internal screen to the window size
	w, h := screen.Bounds().Dx(), screen.Bounds().Dy()
	scaleX := float64(w) / 1920.0
	scaleY := float64(h) / 1080.0
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(scaleX, scaleY)
	screen.DrawImage(g.Screen, opts)
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

func (e *Engine) GoToMainMenu() {
	m := &menu.MainMenu{
		Options:     []menu.MainMenuOption{menu.StartGameOption, menu.OptionsOption, menu.QuitOption},
		ActiveOption: 0,
		L:       e.L,
	}
	// Reset game state
	if e.Game != nil {
		e.Game = nil
	}
	// Set the engine state to MAIN_MENU
	e.State = state.MAIN_MENU
	e.Menu = m
}

func (e *Engine) NewGame() {
	if e.Game != nil {
		return // Game already active
	}
	bg, _, _ := ebitenutil.NewImageFromFile("assets/images/background/beach.png")
	waves := []world.Wave{
		{
			WaveNumber:      1,
			PossibleLetters: []string{"f", "g", "h", "i"},
			EnemyCount:      10,
			MobChances: []world.MobChance{
				{Type: "beachball", Chance: 1.0},
			},
		},
		{
			WaveNumber:      2,
			PossibleLetters: []string{"f", "g", "h", "i", "r"},
			EnemyCount:      15,
			MobChances: []world.MobChance{
				{Type: "beachball", Chance: 1.0},
			},
		},
	}
	// Initialize a new game screen
	gameOpts := game.GameOptions{
		GameMode: game.ENDLESS, // Default to ENDLESS mode
		Level: *world.NewLevel("First Level", 1 , 1, "Beach", []string{"f", "g", "h", "j"}, waves, 10, bg),
	}
	gameScreen := game.NewGame(gameOpts)
	e.Game = gameScreen
	e.State = state.GAME_PLAYING

	// Load Lua plugins for the game
	loadLuaPlugins(e.L, "plugins")
}