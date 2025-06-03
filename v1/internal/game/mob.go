package game

import "math"

// Mob represents a basic enemy moving left.
type Mob struct {
	BaseEntity
	speed      float64
	animTicker int
	alive      bool
	vx, vy     float64
	target     *Base
}

// NewMob returns a new mob at the given position.
func NewMob(x, y float64, target *Base) *Mob {
	w, h := ImgMobA.Bounds().Dx(), ImgMobA.Bounds().Dy()
	return &Mob{
		BaseEntity: BaseEntity{
			pos:          Point{x, y},
			width:        w,
			height:       h,
			frame:        ImgMobA,
			frameAnchorX: float64(w) / 2,
			frameAnchorY: float64(h) / 2,
		},
		speed:  1,
		alive:  true,
		target: target,
	}
}

// Update moves the mob and handles animation.
func (m *Mob) Update() {
	if m.target != nil {
		dx := m.target.pos.X - m.pos.X
		dy := m.target.pos.Y - m.pos.Y
		dist := math.Hypot(dx, dy)
		if dist != 0 {
			m.vx = dx / dist * m.speed
			m.vy = dy / dist * m.speed
		}
	}
	m.pos.X += m.vx
	m.pos.Y += m.vy
	m.animTicker++
	if m.animTicker%30 < 15 {
		m.frame = ImgMobA
	} else {
		m.frame = ImgMobB
	}
}

// Velocity returns the mob's current velocity components.
func (m *Mob) Velocity() (vx, vy float64) {
	return m.vx, m.vy
}
