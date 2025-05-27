package world

import (
	"td/internal/entity"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Level struct {
	Name            string
	Difficulty      string
	Biome           Biome
	PossibleMobs    []entity.Entity
	PossibleLetters []string
	Background      *ebiten.Image
}

func NewLevel(name, difficulty string, possibleLetters []string) *Level {
	var biome Biome
	var background *ebiten.Image
	var possibleMobs []entity.Entity
	var err error
	switch name {
	case "World 1":
		biome = BEACH
		possibleMobs = entity.NewMobList("beach")
		background, _, err = ebitenutil.NewImageFromFile("assets/images/background/beach.png")
		if err != nil {
			panic("Failed to load background image: " + err.Error())
		}
	}
	return &Level{
		Name:            name,
		Difficulty:      difficulty,
		Biome:           biome,
		PossibleMobs:    possibleMobs,
		PossibleLetters: possibleLetters,
		Background:      background,
	}
}

func (l *Level) DrawBackground(screen *ebiten.Image) {
	if l.Background != nil {
		screen.DrawImage(l.Background, nil)
	}
	// Drawing logic for the level can be implemented here.
	// This could include drawing the background, entities, etc.
}
