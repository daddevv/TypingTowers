package mob

import (
	"github.com/daddevv/type-defense/internal/assets"
	"github.com/daddevv/type-defense/internal/core"
	"github.com/daddevv/type-defense/internal/entity"
)

// OrcGrunt represents a basic enemy foot soldier.
type OrcGrunt struct {
	entity.BaseEntity
	Hp     int     // current hit points
	Damage int     // melee damage dealt on contact
	Speed  float64 // movement speed in pixels/sec
	Alive  bool    // whether the grunt is active
}

// NewOrcGrunt creates a new orc grunt at the given position.
func NewOrcGrunt(x, y float64) *OrcGrunt {
	w, h := assets.ImgMobA.Bounds().Dx(), assets.ImgMobA.Bounds().Dy()
	return &OrcGrunt{
		BaseEntity: entity.BaseEntity{
			Position:     core.Point{X: x, Y: y},
			Width:       w,
			Height:      h,
			Frame:       assets.ImgMobA,
			FrameAnchorX: float64(w) / 2,
			FrameAnchorY: float64(h) / 2,
		},
		Hp:     5,
		Damage: 1,
		Speed:  20,
		Alive:  true,
	}
}

// Update moves the grunt to the left and checks for death.
func (o *OrcGrunt) Update(dt float64) error {
	if !o.Alive {
		return nil
	}
	o.Position.X -= o.Speed * dt
	if o.Hp <= 0 {
		o.Alive = false
	}
	return nil
}

// ApplyDamage applies damage to the grunt.
func (o *OrcGrunt) ApplyDamage(amount int) {
	if !o.Alive {
		return
	}
	o.Hp -= amount
	if o.Hp <= 0 {
		o.Alive = false
	}
}

// Health returns the grunt's current HP.
func (o *OrcGrunt) Health() int { return o.Hp }

// AttackDamage returns the damage this grunt deals in melee.
func (o *OrcGrunt) AttackDamage() int { return o.Damage }
