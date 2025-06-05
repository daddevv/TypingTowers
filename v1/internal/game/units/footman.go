package game

import "github.com/hajimehoshi/ebiten/v2"

// Footman represents a simple melee unit spawned from the Barracks.
type Footman struct {
	CombatUnit
}

// NewFootman creates a Footman at the given position.
func NewFootman(x, y float64) *Footman {
	w, h := ImgFootman.Bounds().Dx(), ImgFootman.Bounds().Dy()
	return &Footman{
		CombatUnit: CombatUnit{
			BaseEntity: BaseEntity{
				pos:          Point{x, y},
				width:        w,
				height:       h,
				frame:        ImgFootman,
				frameAnchorX: float64(w) / 2,
				frameAnchorY: float64(h) / 2,
			},
			hp:     10,
			damage: 1,
			speed:  50,
			alive:  true,
		},
	}
}

// Update moves the Footman to the right.
// Update moves the Footman to the right and checks if it is still alive.
func (f *Footman) Update(dt float64) error {
	if !f.alive {
		return nil
	}
	f.pos.X += f.speed * dt
	if f.hp <= 0 {
		f.alive = false
	}
	return nil
}

// Frame satisfies the Entity interface for Footman.
func (f *Footman) Frame() *ebiten.Image { return f.frame }
