package game

import "math"

// Projectile represents a moving projectile toward a target.
type Projectile struct {
	BaseEntity
	vx, vy float64
	speed  float64
	target *Mob
	alive  bool
}

// NewProjectile creates a new projectile aimed at the target.
func NewProjectile(x, y float64, target *Mob) *Projectile {
	dx := target.pos.X - x
	dy := target.pos.Y - y
	dist := math.Hypot(dx, dy)
	vx, vy := dx/dist, dy/dist
	w, h := ImgProjectile.Bounds().Dx(), ImgProjectile.Bounds().Dy()
	return &Projectile{
		BaseEntity: BaseEntity{
			pos:          Point{x, y},
			width:        w,
			height:       h,
			frame:        ImgProjectile,
			frameAnchorX: float64(w) / 2,
			frameAnchorY: float64(h) / 2,
		},
		vx:     vx,
		vy:     vy,
		speed:  5,
		target: target,
		alive:  true,
	}
}

// Update moves the projectile and checks collision.
func (p *Projectile) Update() {
	p.pos.X += p.vx * p.speed
	p.pos.Y += p.vy * p.speed
	if p.target != nil && p.target.alive {
		dx := p.target.pos.X - p.pos.X
		dy := p.target.pos.Y - p.pos.Y
		if math.Hypot(dx, dy) < 16 {
			p.target.alive = false
			p.alive = false
		}
	}
	if p.pos.X < -10 || p.pos.X > 1930 || p.pos.Y < -10 || p.pos.Y > 1090 {
		p.alive = false
	}
}
