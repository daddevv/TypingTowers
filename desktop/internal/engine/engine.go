package engine

import (
	"os"
	"td/internal/game"
	"td/internal/menu"
	"td/internal/world"

	"github.com/hajimehoshi/ebiten/v2"
)

type Engine struct {
	Width   int
	Height  int
	Version string
	State   EngineState
	Screen  *ebiten.Image
	Menu    *menu.MainMenu
	Lobby   *menu.LobbyMenu
	Game    *game.Game
}

func NewGame(width, height int, version string) *Engine {
	return &Engine{
		Version: version,
		Height:  height,
		Width:   width,
		Screen:  ebiten.NewImage(width, height),
		State:   MAIN_MENU,
		Menu:    menu.InitializeMainMenu(),
		Lobby:   nil,
		Game:    nil,
	}
}

func (e *Engine) Update() error {
	switch e.State {
	case MAIN_MENU:
		selection, err := e.Menu.Update()
		if err != nil {
			return err
		}
		if selection != "" {
			switch selection {
			case "Start Game":
				if e.Game == nil {
					e.Game = game.NewGame(game.GameOptions{
						PossibleLetters: []string{"A", "B", "C", "D"},
						Width:           e.Width,
						Height:          e.Height,
						Level:           *world.NewLevel("World 1", "normal", []string{"f", "g", "h", "j"}),
						GameMode:        game.ENDLESS,
					})
				}
				e.State = GAME_PLAYING
			case "Options":
				e.State = GAME_SETTINGS_MENU
			case "Quit":
				os.Exit(0)
			}
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
	switch g.State {
	case MAIN_MENU:
		g.Menu.Draw(screen)
	case GAME_PLAYING:
		g.Game.Draw(screen)
	default:
		// Handle default drawing
	}
}

func (g *Engine) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.Width, g.Height
}
