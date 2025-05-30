package world

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type MobChance struct {
	Type   string  // Type of the mob (e.g., "Goblin", "Orc")
	Chance float64 // Probability of this mob appearing in the wave
}

type Wave struct {
	WaveNumber      int
	PossibleLetters []string
	EnemyCount      int
	MobChances      []MobChance
}

type Level struct {
	Name               string
	WorldNumber        int
	LevelNumber        int
	World              string
	StartingLetters    []string
	Waves              []Wave
	EnemyCount      int
	MobChances      []MobChance
	LevelCompleteScore int
	Background         *ebiten.Image
}

func NewLevel(name string, worldNumber, levelNumber int, world string, startingLetters []string, waves []Wave, levelCompleteScore int, background *ebiten.Image) *Level {
	return &Level{
		Name:               name,
		WorldNumber:        worldNumber,
		LevelNumber:        levelNumber,
		World:              world,
		StartingLetters:    startingLetters,
		Waves:              waves,
		LevelCompleteScore: levelCompleteScore,
		Background:         background,
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
