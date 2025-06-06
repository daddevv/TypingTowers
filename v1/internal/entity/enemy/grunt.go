package enemy

import (
	"github.com/daddevv/type-defense/internal/assets"
	"github.com/daddevv/type-defense/internal/core"
	"github.com/daddevv/type-defense/internal/entity"
	"github.com/hajimehoshi/ebiten/v2"
)

// OrcGrunt represents a basic enemy foot soldier.
type OrcGrunt struct {
	entity.BaseEntity
	Hp      int     // current hit points
	Damage  int     // melee damage dealt on contact
	Speed   float64 // movement speed in pixels/sec
	IsAlive bool    // whether the grunt is active
}

// NewOrcGrunt creates a new orc grunt at the given position.
func NewOrcGrunt(x, y float64) *OrcGrunt {
	w, h := assets.ImgMobA.Bounds().Dx(), assets.ImgMobA.Bounds().Dy()
	return &OrcGrunt{
		BaseEntity: entity.BaseEntity{
			Pos:          core.Point{X: x, Y: y},
			Width:        w,
			Height:       h,
			Sprite:       assets.ImgMobA,
			FrameAnchorX: float64(w) / 2,
			FrameAnchorY: float64(h) / 2,
		},
		Hp:      5,
		Damage:  1,
		Speed:   20,
		IsAlive: true,
	}
}

// Update moves the grunt to the left and checks for death.
func (o *OrcGrunt) Update(dt float64) error {
	if !o.IsAlive {
		return nil
	}
	o.Pos.X -= o.Speed * dt
	if o.Hp <= 0 {
		o.IsAlive = false
	}
	return nil
}

// ApplyDamage applies damage to the grunt.
func (o *OrcGrunt) ApplyDamage(amount int) {
	if !o.IsAlive {
		return
	}
	o.Hp -= amount
	if o.Hp <= 0 {
		o.IsAlive = false
	}
}

// Health returns the grunt's current HP.
func (o *OrcGrunt) Health() int { return o.Hp }

// AttackDamage returns the damage this grunt deals in melee.
func (o *OrcGrunt) AttackDamage() int { return o.Damage }

// Alive returns whether the grunt is still active.
func (o *OrcGrunt) Alive() bool { return o.IsAlive }

// Frame returns the current sprite frame.
func (o *OrcGrunt) Frame() *ebiten.Image {
	return o.Sprite
}

// Position returns the current position of the grunt.
func (o *OrcGrunt) Position() (x, y float64) {
	return o.Pos.X, o.Pos.Y
}

// Bounds returns the bounding box of the grunt.
func (o *OrcGrunt) Bounds() (x, y, width, height int) {
	return int(o.Pos.X), int(o.Pos.Y), o.Width, o.Height
}

// Hitbox returns the hitbox of the grunt.
func (o *OrcGrunt) Hitbox() (x, y, width, height int) {
	return o.Bounds()
}

// Static returns whether the grunt is static (does not move).
func (o *OrcGrunt) Static() bool {
	return false
}

// Destroy cleans up resources used by the grunt.
func (o *OrcGrunt) Destroy() {
	o.Sprite = nil
	o.IsAlive = false // Mark as inactive to prevent further updates
}
