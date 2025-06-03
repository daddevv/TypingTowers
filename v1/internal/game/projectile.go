package game

import "math"

// calcIntercept returns a normalized direction vector from the shooter position
// to where the projectile should aim in order to intercept the moving target.
func calcIntercept(px, py float64, target *Mob, speed float64) (float64, float64) {
	tx, ty := target.pos.X, target.pos.Y
	tvx, tvy := target.Velocity()
	rx := tx - px
	ry := ty - py
	a := tvx*tvx + tvy*tvy - speed*speed
	b := 2 * (rx*tvx + ry*tvy)
	c := rx*rx + ry*ry
	disc := b*b - 4*a*c
	var t float64
	if disc >= 0 && math.Abs(a) > 1e-6 {
		sqrt := math.Sqrt(disc)
		t1 := (-b - sqrt) / (2 * a)
		t2 := (-b + sqrt) / (2 * a)
		t = t1
		if t < 0 || (t2 >= 0 && t2 < t) {
			t = t2
		}
		if t < 0 {
			t = 0
		}
	} else if math.Abs(b) > 1e-6 {
		t = -c / b
		if t < 0 {
			t = 0
		}
	}
	ix := tx + tvx*t
	iy := ty + tvy*t
	dx := ix - px
	dy := iy - py
	d := math.Hypot(dx, dy)
	if d == 0 {
		return 0, 0
	}
	return dx / d, dy / d
}

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
	speed := 5.0
	vx, vy := calcIntercept(x, y, target, speed)
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
		speed:  speed,
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
