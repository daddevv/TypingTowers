package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Engine struct {
	Width   int
	Height  int
	Title   string
	Version string
	Game    *Game
}

func NewEngine(width, height int, title, version string) *Engine {
	return &Engine{
		Width:   width,
		Height:  height,
		Title:   title,
		Version: version,
		Game:    NewGame(),
	}
}

func (g *Engine) Update() error {
	err := g.Game.Update()
	if err != nil {
		return err
	}
	return nil
}

func (g *Engine) Draw(screen *ebiten.Image) {
	g.Game.Draw(screen)
}

func (g *Engine) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.Width, g.Height
}
