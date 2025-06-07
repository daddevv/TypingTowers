package game

import (
	"image/color"
	"log"

	"github.com/daddevv/type-defense/internal/assets"
	"github.com/daddevv/type-defense/internal/event"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type DebugGame struct {
	screen    *ebiten.Image
	input     *Input
	bus       *event.EventBus
	typedLine string
}

func NewDebugGame() *DebugGame {
	ebiten.SetWindowTitle("TypingTowers Debug")
	ebiten.SetWindowSize(1920, 1080)
	g := &DebugGame{
		screen: ebiten.NewImage(1920, 1080),
		input:  NewInput(),
		bus:    event.NewEventBus(),
	}
	return g
}

func (g *DebugGame) Update() error {
	g.input.Update()
	for _, r := range g.input.TypedChars() {
		g.bus.Publish(string(event.WordLetterTyped), event.TypingEvent{Rune: r})
		log.Println(string(r))
		g.typedLine += string(r)
	}
	if g.input.Backspace() && len(g.typedLine) > 0 {
		g.typedLine = g.typedLine[:len(g.typedLine)-1]
	}
	if g.input.Quit() {
		return ebiten.Termination
	}
	return nil
}

func (g *DebugGame) Draw(screen *ebiten.Image) {
	g.screen.Fill(color.Black)
	opts := &text.DrawOptions{}
	opts.GeoM.Translate(20, 40)
	opts.ColorScale.ScaleWithColor(color.White)
	text.Draw(g.screen, g.typedLine, assets.BoldFont, opts)
	screen.DrawImage(g.screen, nil)
}

func (g *DebugGame) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 1920, 1080
}
