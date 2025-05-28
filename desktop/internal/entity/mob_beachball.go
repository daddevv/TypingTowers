package entity

import (
	"image/color"
	"math/rand"
	"td/internal/ui"
	"td/internal/utils"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type BeachballMob struct {
	Pos            ui.Location
	Speed          float64
	Sprite         *ebiten.Image // Current frame to draw
	MoveAnimation  *ui.Animation
	TargetY float64 // Target Y position for the mob
	Letters	   []string // Letters to display above the mob for player to type
	LetterCount int // Number of letters to type
	LetterIndex int // Index of the current letter to type
}

func NewBeachballMob(count int, possible []string) *BeachballMob {
	moveAnimation, err := ui.NewAnimation("assets/images/mob/mob_beachball_sheet.png", 1, 7, 48, 48, 6)
	if err != nil {
		return nil
	}
	initialY := rand.Float64() * 0.3 + 0.6
	mob := &BeachballMob{
		Pos:            ui.Location{X: 1, Y: float64(initialY)},
		Speed:          0.001,
		MoveAnimation:  moveAnimation,
		TargetY:        0.85,
		Letters:        utils.GenerateRandomLetters(count, possible),
		LetterCount:    count,
		LetterIndex:    0,
	}
	mob.Sprite = mob.MoveAnimation.Update()
	return mob
}

func (mob *BeachballMob) Draw(screen *ebiten.Image) {
	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Scale(3*float64(screen.Bounds().Dx())/1920, 3*float64(screen.Bounds().Dy())/1080)
	opts.GeoM.Translate(
		mob.Pos.X*float64(screen.Bounds().Dx()),
		mob.Pos.Y*float64(screen.Bounds().Dy()),
	)

	// --- Draw the target word above the mob, centered and proportional ---
	// Avoid per-frame string concatenation
	word := ""
	for _, letter := range mob.Letters {
		word += letter
	}
	font := ui.Font("Mob", 48*float64(screen.Bounds().Dx())/1920)
	letterSpacing := 0.015

	baseX := mob.Pos.X + float64(mob.LetterIndex)*letterSpacing
	baseY := mob.Pos.Y - 0.05

	for i := 0; i < mob.LetterCount; i++ {
		letter := mob.Letters[i]
		letterX := (baseX + float64(i)*letterSpacing) * float64(screen.Bounds().Dx())
		letterY := baseY * float64(screen.Bounds().Dy())

		optsText := &text.DrawOptions{}
		optsText.GeoM.Translate(letterX, letterY)
		if i == mob.LetterIndex {
			optsText.ColorScale.ScaleWithColor(color.RGBA{255, 0, 0, 255})
		} else {
			optsText.ColorScale.ScaleWithColor(color.White)
		}
		text.Draw(screen, letter, font, optsText)
	}

	screen.DrawImage(mob.Sprite, &opts)
}

func (mob *BeachballMob) Update() error {
	// Update the position of the beachball mob
	// For now, we just move it downwards at a constant speed
	mob.Pos.X -= mob.Speed // Move left
	if mob.Pos.Y < mob.TargetY {
		mob.Pos.Y += (mob.TargetY - mob.Pos.Y)*0.005 * mob.Pos.X // Move towards the ground
	} else if mob.Pos.Y > mob.TargetY {
		mob.Pos.Y -= (mob.Pos.Y - mob.TargetY) * 0.005 * mob.Pos.X // Move towards the ground
	}

	mob.Sprite = mob.MoveAnimation.Update()
	return nil
}

func (mob *BeachballMob) GetPosition() ui.Location {
	return mob.Pos
}

func (mob *BeachballMob) SetPosition(x, y float64) {
	mob.Pos.X = x
	mob.Pos.Y = y
}
