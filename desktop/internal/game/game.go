package game

import (
	"bytes"
	"errors"
	"fmt"
	"image/color"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var fontSource *text.GoTextFaceSource

type Game struct {
	Background *ebiten.Image
	Width      int
	Height     int
	Score      int
	Player     Entity
	Mobs	   []Entity
}

func NewGame() *Game {
	// Initialize the game with a background image and player
	background, _, err := ebitenutil.NewImageFromFile("assets/bg_beach.png")
	if err != nil {
		panic(err)
	}

	// Read the .ttf file into a byte slice.
    data, err := os.ReadFile("assets/OpenDyslexicNerdFontPropo-Bold.otf")
    if err != nil {
        log.Fatal("unable to read font file:", err)
    }
    // Create the font source for text/v2.
    fontSource, err = text.NewGoTextFaceSource(bytes.NewReader(data))
    if err != nil {
        log.Fatal("failed to parse font:", err)
    }
	
	return &Game{
		Background: background,
		Width:      1920,
		Height:     1080,
		Score:      0,
		Player:     NewPlayer(),
		Mobs:       []Entity{},
	}
}

func (e *Game) Update() error {
	// 
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return errors.New("pause")
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		// Spawn a new beachball mob when space is pressed
		mob := NewBeachballMob()
		if mob != nil {
			e.Mobs = append(e.Mobs, mob)
		} else {
			return errors.New("failed to create new beachball mob")
		}
	}

	for m := range e.Mobs {
		if err := e.Mobs[m].Update(); err != nil {
			return err
		}
	}

	// --- Collision and despawn logic ---
	// Example: Remove mobs that go off-screen (X < -100)
	activeMobs := e.Mobs[:0]
	for _, mob := range e.Mobs {
		pos := mob.GetPosition()
		if pos.X > 200 {
			activeMobs = append(activeMobs, mob)
		} else {
			mob.StartDeath()
			// Optionally: increase score, play sound, etc.
		}
	}
	e.Mobs = activeMobs

	// TODO: Add collision checks between mobs, player, projectiles, etc.
	// If collision detected:
	//   mob.StartDeath()
	//   player.StartDeath()
	//   etc.

	return nil
}

func (e *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(e.Background, nil)

	// Prepare the string (e.g., with fmt.Sprintf or message.Printer)
    scoreStr := fmt.Sprintf("Score: %d", e.Score)

    // Configure draw options
    opts := &text.DrawOptions{}
    opts.GeoM.Translate(10, 30)                  // position on screen
    opts.ColorScale.ScaleWithColor(color.White)  // text color

    // Choose a text.Face implementation
    face := &text.GoTextFace{Source: fontSource, Size: 48} // Use the font source created earlier

    // Draw the score
    text.Draw(screen, scoreStr, face, opts)

	entities := append(e.Mobs, e.Player)
	// TODO: Sort entities by Z-index (smallest Y first) if needed
	for _, entity := range entities {
		entity.Draw(screen)
	}
}
