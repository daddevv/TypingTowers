package game

import "math"

// Mob represents a basic enemy moving left.
type Mob struct {
	BaseEntity
	speed      float64
	animTicker float64
	alive      bool
	vx, vy     float64
	target     *Base
	health     int
}

// NewMob returns a new mob at the given position.
func NewMob(x, y float64, target *Base, hp int, speed float64) *Mob {
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
		speed:  speed,
		alive:  true,
		target: target,
		health: hp,
	}
}

// Update moves the mob and handles animation.
func (m *Mob) Update(dt float64) {
	if m.target != nil {
		dx := m.target.pos.X - m.pos.X
		dy := m.target.pos.Y - m.pos.Y
		dist := math.Hypot(dx, dy)
		if dist != 0 {
			m.vx = dx / dist * m.speed
			m.vy = dy / dist * m.speed
		}
	}
	m.pos.X += m.vx * dt
	m.pos.Y += m.vy * dt
	m.animTicker += dt
	if int(m.animTicker/0.25)%2 == 0 {
		m.frame = ImgMobA
	} else {
		m.frame = ImgMobB
	}
}

// Velocity returns the mob's current velocity components.
func (m *Mob) Velocity() (vx, vy float64) {
	return m.vx, m.vy
}
