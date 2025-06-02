package sprite

import (
	"image"
	"td/internal/math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite interface {
	GetFrame() *ebiten.Image
	GetFrameRect() image.Rectangle
	GetPosition() math.Vec2
	SetPosition(pos math.Vec2)
	Update(deltaTime float64)
	Draw(canvas *ebiten.Image, opts *ebiten.DrawImageOptions)
}
