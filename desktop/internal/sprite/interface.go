package sprite

import (
	"image"
	"td/internal/physics"

	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite interface {
	GetFrame() *ebiten.Image
	GetFrameRect() image.Rectangle
	GetPosition() physics.Vec2
	SetPosition(pos physics.Vec2)
	Update(deltaTime float64)
	Draw(canvas *ebiten.Image, opts *ebiten.DrawImageOptions)
}
