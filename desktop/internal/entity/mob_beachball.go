package entity

import (
	"image"
	"math/rand"
	"td/internal/ui"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type BeachballMob struct {
	Pos            ui.Location
	Speed          float64
	CurrentFrame   int
	MoveAnimation  []*ebiten.Image
	AnimationDelay int           // Delay between frames (in ticks)
	animationTick  int           // Internal counter for animation
	Sprite         *ebiten.Image // Current frame to draw
}

func NewBeachballMob() *BeachballMob {
	spritesheet, _, err := ebitenutil.NewImageFromFile("assets/mob_beachball_sheet.png")
	if err != nil {
		return nil
	}

	initialY := rand.Intn(200) + 700
	frames := 7
	size := 48

	mob := &BeachballMob{
		Pos:            ui.Location{X: 1920, Y: float64(initialY)},
		Speed:          2,
		MoveAnimation:  make([]*ebiten.Image, frames),
		CurrentFrame:   0,
		AnimationDelay: 10, // 120 TPS / 7 frames ≈ 17 ticks per frame
		animationTick:  0,
	}

	for i := range frames {
		rect := image.Rect(i*size, 0, (i+1)*size, size)
		frame := spritesheet.SubImage(rect).(*ebiten.Image)
		mob.MoveAnimation[i] = frame
	}
	mob.Sprite = mob.MoveAnimation[0]
	return mob
}

func (mob *BeachballMob) Draw(screen *ebiten.Image) {
	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Scale(3, 3)
	opts.GeoM.Translate(mob.Pos.X, mob.Pos.Y)
	screen.DrawImage(mob.Sprite, &opts)
}

func (mob *BeachballMob) Update() error {
	// Update the position of the beachball mob
	// For now, we just move it downwards at a constant speed
	mob.Pos.X -= mob.Speed // Move left
	if mob.Pos.Y < 910 {
		mob.Pos.Y += (910 - mob.Pos.Y) * 3 / mob.Pos.X // Move towards the ground
	} else if mob.Pos.Y > 910 {
		mob.Pos.Y -= (mob.Pos.Y - 910) * 3 / mob.Pos.X // Move towards the ground
	}

	// Animate: advance frame every AnimationDelay ticks
	mob.animationTick++
	if mob.animationTick >= mob.AnimationDelay {
		mob.CurrentFrame = (mob.CurrentFrame + 1) % len(mob.MoveAnimation)
		mob.Sprite = mob.MoveAnimation[mob.CurrentFrame]
		mob.animationTick = 0
	}

	return nil
}

func (mob *BeachballMob) GetPosition() ui.Location {
	return mob.Pos
}

func (mob *BeachballMob) SetPosition(x, y float64) {
	mob.Pos.X = x
	mob.Pos.Y = y
}

func (mob *BeachballMob) StartDeath() {
	// TODO: Implement mob death logic (animation, removal, etc)
}
