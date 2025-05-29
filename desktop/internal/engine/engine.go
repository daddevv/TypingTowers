package engine

import (
	"os"
	"td/internal/game"
	"td/internal/ui"
	"td/internal/world"

	"github.com/hajimehoshi/ebiten/v2"
)

type Engine struct {
	Width   int
	Height  int
	Version string
	State   EngineState
	Screen  *ebiten.Image
	Menu    *ui.MainMenu
	Lobby   *ui.LobbyMenu
	Game    *game.Game
}

func NewEngine(width, height int, version string) *Engine {
	// Load all content configs at engine startup
	err := LoadAllContent()
	if err != nil {
		panic("Failed to load content configs: " + err.Error())
	}
	return &Engine{
		Version: version,
		Height:  1080, // Always use 1920x1080 for internal canvas
		Width:   1920,
		Screen:  ebiten.NewImage(1920, 1080),
		State:   MAIN_MENU,
		Menu:    ui.InitializeMainMenu(),
		Lobby:   nil,
		Game:    nil,
	}
}

func (e *Engine) Update() error {
	switch e.State {
	case MAIN_MENU:
		state, err := e.Menu.Update()
		if err != nil {
			return err
		}
		switch state {
		case "Start Game":
			if e.Game == nil {
				e.Game = game.NewGame(game.GameOptions{
					Level:    *world.NewLevel("World 1", "normal", []string{"f", "g", "h", "j"}),
					GameMode: game.ENDLESS,
					MobConfigs: LoadedMobs, // Pass loaded mob configs
				})
			}
			e.State = GAME_PLAYING
		case "Options":
			e.State = GAME_SETTINGS_MENU
		case "Quit":
			os.Exit(0)
		}
		return nil
	case GAME_PLAYING:
		err := e.Game.Update()
		if err != nil {
			if err.Error() == "pause" {
				e.State = MAIN_MENU
			} else {
				return err
			}
		}
		return nil
	default:
		// Handle default logic
	}

	return nil
}

func (g *Engine) Draw(screen *ebiten.Image) {
	// Draw everything to the internal 1920x1080 screen
	g.Screen.Clear()
	switch g.State {
	case MAIN_MENU:
		g.Menu.Draw(g.Screen)
	case GAME_PLAYING:
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

func (g *Engine) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	// Always use 1920x1080 for the internal canvas
	return 1920, 1080
}
