package world

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Level struct {
	Name            string
	Difficulty      string
	Biome           Biome
	PossibleLetters []string
	Background      *ebiten.Image
}

func NewLevel(name, difficulty string, possibleLetters []string) *Level {
	var biome Biome
	var background *ebiten.Image
	var err error
	switch name {
	case "World 1":
		biome = BEACH
		background, _, err = ebitenutil.NewImageFromFile("assets/images/background/beach.png")
		if err != nil {
			panic("Failed to load background image: " + err.Error())
		}
	}
	return &Level{
		Name:            name,
		Difficulty:      difficulty,
		Biome:           biome,
		PossibleLetters: possibleLetters,
		Background:      background,
	}
}

// DrawBackground draws the level's background on the provided screen.
// It scales the background image to fill the screen dimensions.
func (l *Level) DrawBackground(screen *ebiten.Image) {
	if l.Background == nil {
		return
	}
	// Draw background at 1920x1080, no scaling to window size
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(1920.0/float64(l.Background.Bounds().Dx()), 1080.0/float64(l.Background.Bounds().Dy()))
	screen.DrawImage(l.Background, opts)
	// Drawing logic for the level can be implemented here.
	// This could include drawing the background, entities, etc.
}
