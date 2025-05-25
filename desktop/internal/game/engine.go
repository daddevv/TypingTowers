package game

import (
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

type Engine struct {
	Width   int
	Height  int
	Title   string
	Version string
	Scene   CurrentScene
	Input   *InputState
	Menu    *Menu
	Game    *Game
}

func NewEngine(width, height int, title, version string) *Engine {
	return &Engine{
		Width:   width,
		Height:  height,
		Title:   title,
		Version: version,
		Scene:   NewCurrentScene(),
		Menu:  NewMenu(),
		Game:  NewGame(),
	}
}

func (e *Engine) Update() error {
	switch e.Scene {
	case MAINMENU:
		selection, err := e.Menu.Update()
		if err != nil {
			return err
		}
		if selection != "" {
			switch selection {
			case "Start Game":
				e.Scene = GAME
			case "Options":
				e.Scene = OPTIONS
			case "Quit":
				os.Exit(0)
			}
		}
		return nil
	case GAME:
		err := e.Game.Update()
		if err != nil {
			return err
		}
		return nil
	default:
		// Handle default logic
	}

	return nil
}

func (g *Engine) Draw(screen *ebiten.Image) {
	switch g.Scene {
	case MAINMENU:
		g.Menu.Draw(screen)
	case GAME:
		g.Game.Draw(screen)
	case GAMEOVER:
		// Handle game over drawing
	default:
		// Handle default drawing
	}
}

func (g *Engine) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.Width, g.Height
}
