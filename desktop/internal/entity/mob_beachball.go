package entity

import (
	"fmt"
	"math/rand"
	"td/internal/ui"

	"github.com/hajimehoshi/ebiten/v2"
)

type BeachballMob struct {
	Pos            ui.Location
	Speed          float64
	Sprite         *ebiten.Image // Current frame to draw
	MoveAnimation  *ui.Animation
	TargetY float64 // Target Y position for the mob
}

func NewBeachballMob() *BeachballMob {
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
	}
	mob.Sprite = mob.MoveAnimation.Update()
	fmt.Println("Beachball mob created at position:", mob.Pos.X, mob.Pos.Y)
	return mob
}

func (mob *BeachballMob) Draw(screen *ebiten.Image) {
	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Scale(3*float64(screen.Bounds().Dx())/1920, 3*float64(screen.Bounds().Dy())/1080) // Scale the mob image
	opts.GeoM.Translate(
		mob.Pos.X* float64(screen.Bounds().Dx()),
		mob.Pos.Y* float64(screen.Bounds().Dy()),
	) // Position based on screen size
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
	fmt.Println("Beachball mob updated to position:", mob.Pos.X, mob.Pos.Y)
	return nil
}

func (mob *BeachballMob) GetPosition() ui.Location {
	return mob.Pos
}

func (mob *BeachballMob) SetPosition(x, y float64) {
	mob.Pos.X = x
	mob.Pos.Y = y
}
