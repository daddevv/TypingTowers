package ally

import (
	"github.com/daddevv/type-defense/internal/assets"
	"github.com/daddevv/type-defense/internal/core"
	"github.com/daddevv/type-defense/internal/entity"
)

// Footman represents a simple melee unit spawned from the Barracks with basic
// combat stats.
type Footman struct {
	entity.BaseEntity
	Hp     int     // current hit points
	Damage int     // melee damage dealt when attacking
	Speed  float64 // movement speed in pixels/sec
	Alive  bool    // whether the unit is active
}

// NewFootman creates a Footman at the given position.
func NewFootman(x, y float64) *Footman {
	w, h := assets.ImgFootman.Bounds().Dx(), assets.ImgFootman.Bounds().Dy()
	return &Footman{
		BaseEntity: entity.BaseEntity{
			Pos:          core.Point{X: x, Y: y},
			Width:        w,
			Height:       h,
			Sprite:       assets.ImgFootman,
			FrameAnchorX: float64(w) / 2,
			FrameAnchorY: float64(h) / 2,
		},
		Hp:     10,
		Damage: 1,
		Speed:  50,
		Alive:  true,
	}
}

// Update moves the Footman to the right.
// Update moves the Footman to the right and checks if it is still alive.
func (f *Footman) Update(dt float64) error {
	if !f.Alive {
		return nil
	}
	f.Pos.X += f.Speed * dt
	if f.Hp <= 0 {
		f.Alive = false
	}
	return nil
}

// Damage reduces the Footman's HP.
func (f *Footman) ApplyDamage(amount int) {
	if !f.Alive {
		return
	}
	f.Hp -= amount
	if f.Hp <= 0 {
		f.Alive = false
	}
}

// Health returns the Footman's current HP.
func (f *Footman) Health() int { return f.Hp }

// AttackDamage returns the damage dealt by the Footman when attacking.
func (f *Footman) AttackDamage() int {
	return f.Damage
}
