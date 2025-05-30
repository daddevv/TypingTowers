package entity

import (
	"math/rand"
	"td/internal/ui"
	"td/internal/utils"
	"time"

	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// BeachballMob represents a mobile entity that moves across the screen,
// displaying letters that the player must type.
type BeachballMob struct {
	MobBase
}

// Ensure BeachballMob implements Mob interface
var _ Mob = (*BeachballMob)(nil)

func (m *BeachballMob) GetLetters() []Letter {
	return m.Letters
}

func (m *BeachballMob) IsDead() bool {
	return m.Dead
}

func (m *BeachballMob) IsPendingDeath() bool {
	return m.PendingDeath
}

func (m *BeachballMob) StartDeath() {
	if m.PendingDeath && !m.Dead {
		m.Dead = true
		m.PendingDeath = false
		m.DeathTimer = 60
	}
}

func (m *BeachballMob) IncrementPendingProjectiles() {
	m.PendingProjectiles++
}

func (m *BeachballMob) GetPendingProjectiles() int {
	return m.PendingProjectiles
}

// NewBeachballMob creates a new BeachballMob with random letters from possible choices.
func NewBeachballMob(count int, possible []string) *BeachballMob {
	return NewBeachballMobWithLetters(utils.GenerateRandomLetters(count, possible))
}

// NewBeachballMobWithWord creates a new BeachballMob with letters from a specific word.
func NewBeachballMobWithWord(word string) *BeachballMob {
	letters := make([]string, len(word))
	for i, char := range word {
		letters[i] = string(char)
	}
	return NewBeachballMobWithLetters(letters)
}

// NewBeachballMobWithLetters creates a new BeachballMob with the given letters.
func NewBeachballMobWithLetters(letters []string) *BeachballMob {
	moveAnimation, err := ui.NewAnimation("assets/images/mob/mob_beachball_sheet.png", 1, 7, 48, 48, 6)
	if err != nil {
		return nil
	}
	initialY := rand.Float64()*300 + 600 // Y in px, between 600 and 900
	m := &BeachballMob{
		MobBase: MobBase{
			Position:       ui.Location{X: 1920, Y: initialY}, // start off right edge
			Speed:          2.0, // px per frame
			MoveAnimation:  moveAnimation,
			MoveTarget:     ui.Location{X: 100, Y: 900}, // px
			Letters:        make([]Letter, len(letters)),
			WordWidth:      0.0,
			IdleAnimation:  nil,
			Dead:           false,
			PendingDeath:   false,
			DeathTimer:     0,
			PendingProjectiles: 0,
		},
	}
	font := ui.Font("Mob", 48)
	for i := range m.MobBase.Letters {
		char := []rune(letters[i])[0]
		if i == 0 {
			m.MobBase.Letters[i] = NewLetter(GetLetterImage(char, LetterTarget, font), LetterTarget, char)
		} else {
			m.MobBase.Letters[i] = NewLetter(GetLetterImage(char, LetterActive, font), LetterActive, char)
		}
	}
	// Calculate the total width of the word formed by the letters (in px)
	for _, letter := range m.MobBase.Letters {
		m.WordWidth += float64(letter.Sprite.Bounds().Dx())
	}
	m.WordWidth += float64(len(m.MobBase.Letters)-1) * 24 // 24px spacing
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
	letterSpacing := 10.0 // px
	baseX := m.Position.X + float64(m.Sprite.Bounds().Dx())*1.5 - m.WordWidth/2.0
	baseY := m.Position.Y - 40.0 // 40px above mob (was 50px, moved down 10px)

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

	// Check if all letters are inactive and DefeatedAt is not set
	allInactive := true
	for _, l := range mob.Letters {
		if l.State != LetterInactive {
			allInactive = false
			break
		}
	}
	if allInactive && mob.DefeatedAt == 0 {
		mob.DefeatedAt = float64(time.Now().UnixNano()) / 1e9
	}
	// If defeated and grace period passed, and no pending projectiles, mark as dead
	if mob.DefeatedAt > 0 && !mob.Dead {
		if mob.PendingProjectiles == 0 && mob.DeathTimer <= 0 {
			if float64(time.Now().UnixNano())/1e9-mob.DefeatedAt > 2.0 {
				mob.Dead = true
			}
		}
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

func (m *BeachballMob) AdvanceLetterState(char rune) {
	targetIndex := -1
	for i, letter := range m.Letters {
		if letter.State == LetterTarget && letter.Character == char {
			m.Letters[i].State = LetterInactive
			m.Letters[i].Sprite = GetLetterImage(letter.Character, LetterInactive, ui.Font("Mob", 48))
			targetIndex = i
			break
		}
	}
	if targetIndex >= 0 && targetIndex+1 < len(m.Letters) {
		nextIndex := targetIndex + 1
		m.Letters[nextIndex].State = LetterTarget
		m.Letters[nextIndex].Sprite = GetLetterImage(m.Letters[nextIndex].Character, LetterTarget, ui.Font("Mob", 48))
	}
}
