package entity

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// OrcGrunt represents a basic enemy foot soldier.
type OrcGrunt struct {
	BaseEntity
	hp     int     // current hit points
	damage int     // melee damage dealt on contact
	speed  float64 // movement speed in pixels/sec
	alive  bool    // whether the grunt is active
}

// NewOrcGrunt creates a new orc grunt at the given position.
func NewOrcGrunt(x, y float64) *OrcGrunt {
	w, h := ImgMobA.Bounds().Dx(), ImgMobA.Bounds().Dy()
	return &OrcGrunt{
		BaseEntity: BaseEntity{
			pos:          Point{x, y},
			width:        w,
			height:       h,
			frame:        ImgMobA,
			frameAnchorX: float64(w) / 2,
			frameAnchorY: float64(h) / 2,
		},
		hp:     5,
		damage: 1,
		speed:  20,
		alive:  true,
	}
}

// Update moves the grunt to the left and checks for death.
func (o *OrcGrunt) Update(dt float64) error {
	if !o.alive {
		return nil
	}
	o.pos.X -= o.speed * dt
	if o.hp <= 0 {
		o.alive = false
	}
	return nil
}

// Alive reports whether the grunt is still active.
func (o *OrcGrunt) Alive() bool { return o.alive }

// Damage applies damage to the grunt.
func (o *OrcGrunt) Damage(amount int) {
	if !o.alive {
		return
	}
	o.hp -= amount
	if o.hp <= 0 {
		o.alive = false
	}
}

// Health returns the grunt's current HP.
func (o *OrcGrunt) Health() int { return o.hp }

// AttackDamage returns the damage this grunt deals in melee.
func (o *OrcGrunt) AttackDamage() int { return o.damage }

// Frame satisfies the Entity interface for OrcGrunt.
func (o *OrcGrunt) Frame() *ebiten.Image { return o.frame }
