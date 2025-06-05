package game

import "github.com/hajimehoshi/ebiten/v2"

// OrcGrunt represents a basic enemy foot soldier.
type OrcGrunt struct {
	CombatUnit
}

// NewOrcGrunt creates a new orc grunt at the given position.
func NewOrcGrunt(x, y float64) *OrcGrunt {
	w, h := ImgMobA.Bounds().Dx(), ImgMobA.Bounds().Dy()
	return &OrcGrunt{
		CombatUnit: CombatUnit{
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
		},
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

// Frame satisfies the Entity interface for OrcGrunt.
func (o *OrcGrunt) Frame() *ebiten.Image { return o.frame }
