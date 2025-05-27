package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Entity interface {
	Draw(screen *ebiten.Image)
	Update() error
	GetPosition() Location
	SetPosition(x, y float64)
	StartDeath()
}
