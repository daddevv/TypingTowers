package entity

import (
	"td/internal/ui"

	"github.com/hajimehoshi/ebiten/v2"
)

type Entity interface {
	Draw(screen *ebiten.Image)
	Update() error
	SetPosition(x, y float64)
	GetPosition() ui.Location
}

func EmptyList() []Entity {
	return make([]Entity, 0)
}
