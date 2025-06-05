package game

import "github.com/hajimehoshi/ebiten/v2"

// Footman represents a simple melee unit spawned from the Barracks.
type Footman struct {
	BaseEntity
	speed float64
	alive bool
}

// NewFootman creates a Footman at the given position.
func NewFootman(x, y float64) *Footman {
	w, h := ImgFootman.Bounds().Dx(), ImgFootman.Bounds().Dy()
	return &Footman{
		BaseEntity: BaseEntity{
			pos:          Point{x, y},
			width:        w,
			height:       h,
			frame:        ImgFootman,
			frameAnchorX: float64(w) / 2,
			frameAnchorY: float64(h) / 2,
		},
		speed: 50,
		alive: true,
	}
}

// Update moves the Footman to the right.
func (f *Footman) Update(dt float64) error {
	f.pos.X += f.speed * dt
	return nil
}

// Alive reports whether the Footman is still active.
func (f *Footman) Alive() bool { return f.alive }

// Frame satisfies the Entity interface for Footman.
func (f *Footman) Frame() *ebiten.Image { return f.frame }
