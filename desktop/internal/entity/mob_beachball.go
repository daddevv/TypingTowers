package entity

import (
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
	initialY := rand.Float64()*300 + 600 // Y in px, between 600 and 900
	m := &BeachballMob{
		Mob: Mob{
			Position:       ui.Location{X: 1920, Y: initialY}, // start off right edge
			Speed:          2.0, // px per frame
			MoveAnimation:  moveAnimation,
			MoveTarget:     ui.Location{X: 100, Y: 900}, // px
			Letters:        make([]Letter, count),
			WordWidth:      0.0,
			IdleAnimation:  nil,
			Dead:           false,
			DeathTimer:     0,
		},
	}
	font := ui.Font("Mob", 32)
	for i := range m.Mob.Letters {
		if i == 0 {
			m.Mob.Letters[i] = NewLetter(GetLetterImage([]rune(possible[i])[0], LetterTarget, font), LetterTarget)
		} else {
			m.Mob.Letters[i] = NewLetter(GetLetterImage([]rune(possible[i])[0], LetterActive, font), LetterActive)
		}
	}
	// Calculate the total width of the word formed by the letters (in px)
	for _, letter := range m.Mob.Letters {
		m.WordWidth += float64(letter.Sprite.Bounds().Dx())
	}
	m.WordWidth += float64(len(m.Mob.Letters)-1) * 24 // 24px spacing
	m.Sprite = m.MoveAnimation.Update()
	return m
}

// Draw renders the BeachballMob on the given screen.
func (m *BeachballMob) Draw(screen *ebiten.Image) {
	// Draw mob sprite at absolute position (1920x1080 canvas)
	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Scale(3, 3)
	opts.GeoM.Translate(m.Position.X, m.Position.Y)
	screen.DrawImage(m.Sprite, &opts)

	// Draw the target word above the mob, using cached images
	letterSpacing := 24.0 // px
	baseX := m.Position.X + float64(m.Sprite.Bounds().Dx())*1.5 - m.WordWidth/2.0
	baseY := m.Position.Y - 60.0 // 60px above mob

	for i := 0; i < len(m.Letters); i++ {
		img := m.Letters[i].Sprite
		letterX := baseX + float64(i)*letterSpacing + float64(i)*float64(img.Bounds().Dx())
		letterY := baseY
		imgOpts := &ebiten.DrawImageOptions{}
		imgOpts.GeoM.Translate(letterX, letterY)
		screen.DrawImage(img, imgOpts)

		// Debug: Draw red rectangle around letter bounds
		bounds := img.Bounds()
		debugRect := ebiten.NewImage(bounds.Dx(), bounds.Dy())
		debugRect.Fill(color.RGBA{255, 0, 0, 0})
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
			mob.Position.X = -9999
		}
		return nil
	}
	// Move left at constant speed in px
	mob.Position.X -= mob.Speed
	// Move Y towards MoveTarget.Y
	if mob.Position.Y < mob.MoveTarget.Y {
		mob.Position.Y += (mob.MoveTarget.Y - mob.Position.Y) * 0.005
	} else if mob.Position.Y > mob.MoveTarget.Y {
		mob.Position.Y -= (mob.Position.Y - mob.MoveTarget.Y) * 0.005
	}
	mob.Sprite = mob.MoveAnimation.Update()

	allInactive := true
	for _, letter := range mob.Letters {
		if letter.State != LetterInactive {
			allInactive = false
			break
		}
	}
	if allInactive && !mob.Dead {
		mob.Dead = true
		mob.DeathTimer = 60
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
