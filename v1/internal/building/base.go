package building

import (
	"github.com/daddevv/type-defense/internal/assets"
	"github.com/daddevv/type-defense/internal/core"
	"github.com/daddevv/type-defense/internal/entity"
)

// Base represents the player's base that mobs try to destroy.
const BaseStartingHealth = 10

// Base represents the player's base that mobs try to destroy.
type Base struct {
	entity.BaseEntity
	Hp int
}

// NewBase creates a new base at the given position.
func NewBase(x, y float64, hp int) *Base {
	w, h := assets.ImgBase.Bounds().Dx(), assets.ImgBase.Bounds().Dy()
	return &Base{
		BaseEntity: entity.BaseEntity{
			Position:     core.Point{X: x, Y: y},
			Width:        w,
			Height:       h,
			Frame:        assets.ImgBase,
			FrameAnchorX: float64(w) / 2,
			FrameAnchorY: float64(h) / 2,
			Static:       true,
		},
		Hp: hp,
	}
}

// Update updates the base state. Currently does nothing.
func (b *Base) Update(dt float64) {}

// ApplyDamage reduces the base's health by the given amount.
func (b *Base) ApplyDamage(amount int) {
	b.Hp -= amount
}

// Alive reports whether the base still has health remaining.
func (b *Base) Alive() bool {
	return b.Hp > 0
}

// Health returns the current health of the base.
func (b *Base) Health() int {
	return b.Hp
}
