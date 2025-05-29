package entity

import (
	"fmt"
	"math/rand"
	"td/internal/ui"

	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// BeachballMob represents a mobile entity that moves across the screen,
// displaying letters that the player must type.
type BeachballMob struct {
	Mob
}

// NewBeachballMob creates a new BeachballMob with the given letter count
// and possible letters to choose from.
func NewBeachballMob(count int, possible []string) *BeachballMob {
	moveAnimation, err := ui.NewAnimation("assets/images/mob/mob_beachball_sheet.png", 1, 7, 48, 48, 6)
	if err != nil {
		return nil
	}
	initialY := rand.Float64() * 0.3 + 0.6
	m := &BeachballMob{
		Mob: Mob{
			Position:            ui.Location{X: 1, Y: float64(initialY)},
			Speed:          0.001,
			MoveAnimation:  moveAnimation,
			MoveTarget:     ui.Location{X: 0.05, Y: float64(0.85)},
			Letters:        make([]Letter, count),
			WordWidth:    	0.0, // Width will be calculated later
			IdleAnimation:  nil, // No idle animation for this mob
			Dead:           false,
			DeathTimer:     0,
		},
	}
	winWidth, _ := ebiten.WindowSize()
	font := ui.Font("Mob", 32 * (float64(winWidth) / 1920))
	for i := range m.Mob.Letters {
		if i == 0 {
			m.Mob.Letters[i] = NewLetter(GetLetterImage([]rune(possible[i])[0], LetterTarget, font), LetterTarget)
		} else {
			m.Mob.Letters[i] = NewLetter(GetLetterImage([]rune(possible[i])[0], LetterActive, font), LetterActive)
		}
	}
	// Calculate the total width of the word formed by the letters
	for _, letter := range m.Mob.Letters {
		m.WordWidth += float64(letter.Sprite.Bounds().Dx())
	}
	m.WordWidth = m.WordWidth / float64(1920) // Scale width based on window size
	m.WordWidth += float64(len(m.Mob.Letters)-1) * 0.015 * (float64(winWidth) / 1920) // Add spacing between letters
	fmt.Println("Word width:", m.WordWidth)
	m.Sprite = m.MoveAnimation.Update()
	return m
}

// Draw renders the BeachballMob on the given screen.
func (m *BeachballMob) Draw(screen *ebiten.Image) {
	scalingFactor := float64(screen.Bounds().Dx()) / 1920.0

	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Scale(3*scalingFactor, 3*scalingFactor) // Scale the mob size
	opts.GeoM.Translate(
		m.Position.X*float64(screen.Bounds().Dx()),
		m.Position.Y*float64(screen.Bounds().Dy()),
	)
	screen.DrawImage(m.Sprite, &opts)

	// Draw the target word above the mob, using cached images
	letterSpacing := 0.015
	baseX := m.Position.X + float64(m.Sprite.Bounds().Dx())/1920.0/2.0 - m.WordWidth/2.0
	baseY := m.Position.Y - 0.05
	fmt.Println(m.Position.X, m.Sprite.Bounds().Dx(), scalingFactor, m.WordWidth)

	for i := 0; i < len(m.Letters); i++ {
		img := m.Letters[i].Sprite
		// Convert normalized coordinates to pixel coordinates
		letterX := (baseX + float64(i)*letterSpacing) * float64(screen.Bounds().Dx())
		letterY := baseY * float64(screen.Bounds().Dy())
		imgOpts := &ebiten.DrawImageOptions{}
		imgOpts.GeoM.Translate(letterX, letterY)
		screen.DrawImage(img, imgOpts)
		fmt.Printf("Drawing letter %x at (%f, %f)\n", m.Letters[i], letterX, letterY)

		// Debug: Draw red rectangle around letter bounds
		bounds := img.Bounds()
		debugRect := ebiten.NewImage(bounds.Dx(), bounds.Dy())
		debugRect.Fill(color.RGBA{255, 0, 0, 0}) // Transparent fill
		// Draw only the border
		for x := 0; x < bounds.Dx(); x++ {
			debugRect.Set(x, 0, color.RGBA{255, 0, 0, 255})
			debugRect.Set(x, bounds.Dy()-1, color.RGBA{255, 0, 0, 255})
		}
		for y := 0; y < bounds.Dy(); y++ {
			debugRect.Set(0, y, color.RGBA{255, 0, 0, 255})
			debugRect.Set(bounds.Dx()-1, y, color.RGBA{255, 0, 0, 255})
		}
		debugOpts := &ebiten.DrawImageOptions{}
		debugOpts.GeoM.Translate(letterX, letterY)
		screen.DrawImage(debugRect, debugOpts)
	}

}

// Update advances the BeachballMob's position and updates its animation state.
func (mob *BeachballMob) Update() error {
	if mob.Dead {
		mob.DeathTimer--
		if mob.DeathTimer <= 0 {
			// Mark for removal by setting X off-screen
			mob.Position.X = -9999
		}
		return nil
	}
	// Update the position of the beachball mob
	// For now, we just move it downwards at a constant speed
	mob.Position.X -= mob.Speed // Move left
	if mob.Position.Y < mob.MoveTarget.Y {
		mob.Position.Y += (mob.MoveTarget.Y - mob.Position.Y)*0.005 * mob.Position.X // Move towards the ground
	} else if mob.Position.Y > mob.MoveTarget.Y {
		mob.Position.Y -= (mob.Position.Y - mob.MoveTarget.Y) * 0.005 * mob.Position.X // Move towards the ground
	}
	mob.Sprite = mob.MoveAnimation.Update()

	// Check if all letters are INACTIVE
	allInactive := true
	for _, letter := range mob.Letters {
		if letter.State != LetterInactive {
			allInactive = false
			break
		}
	}
	if allInactive && !mob.Dead {
		mob.Dead = true
		mob.DeathTimer = 60 // 1 second at 60fps
	}
	return nil
}

// GetPosition returns the current position of the BeachballMob.
func (mob *BeachballMob) GetPosition() ui.Location {
	return mob.Position
}

// SetPosition sets the BeachballMob's position to the given coordinates.
func (mob *BeachballMob) SetPosition(x, y float64) {
	mob.Position.X = x
	mob.Position.Y = y
}
