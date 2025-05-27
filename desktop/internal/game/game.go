package game

import (
	"errors"
	"td/internal/entity"
	"td/internal/world"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var fontSource *text.GoTextFaceSource

type Game struct {
	Width  int
	Height int
	Level  world.Level
	Player entity.Entity
	Mobs []entity.Entity
}

func NewGame(opts GameOptions) *Game {
	return &Game{
		Width:      opts.Width,
		Height:     opts.Height,
		Level:      opts.Level,
		Player:     entity.NewPlayer(),
		Mobs: entity.EmptyList(),
	}
}

func (g *Game) Update() error {
	// Handle pause
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return errors.New("pause")
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		// Spawn a new beachball mob when space is pressed
		mob := entity.NewBeachballMob()
		if mob != nil {
			g.Mobs = append(g.Mobs, mob)
		} else {
			return errors.New("failed to create new beachball mob")
		}
	}

	for m := range g.Mobs {
		if err := g.Mobs[m].Update(); err != nil {
			return err
		}
	}

	// --- Collision and despawn logic ---
	// Example: Remove mobs that go off-screen (X < -100)
	activeMobs := g.Mobs[:0]
	for _, mob := range g.Mobs {
		pos := mob.GetPosition()
		if pos.X > 200 {
			activeMobs = append(activeMobs, mob)
		}
	}
	g.Mobs = activeMobs

	// TODO: Add collision checks between mobs, player, projectiles, etc.
	// If collision detected:
	//   mob.StartDeath()
	//   player.StartDeath()
	//   etc.

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Level.DrawBackground(screen)
    // scoreStr := fmt.Sprintf("Score: %d", g.Score)
    // opts := &text.DrawOptions{}
    // opts.GeoM.Translate(10, 30)                  // position on screen
    // opts.ColorScale.ScaleWithColor(color.White)  // text color
    // font := ui.Font("Game-Bold", 48)			 // use the font source to get the font
    // text.Draw(screen, scoreStr, font, opts)

	entities := append(g.Mobs, g.Player)
	// TODO: Sort entities by Z-index (smallest Y first) if needed
	for _, entity := range entities {
		entity.Draw(screen)
	}
}
