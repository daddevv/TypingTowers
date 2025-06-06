package projectile

import (
	"math"

	"github.com/daddevv/type-defense/internal/assets"
	"github.com/daddevv/type-defense/internal/core"
	"github.com/daddevv/type-defense/internal/entity"
	"github.com/daddevv/type-defense/internal/entity/enemy"
)

// calcIntercept returns a normalized direction vector from the shooter position
// to where the projectile should aim in order to intercept the moving target.
func calcIntercept(px, py float64, target enemy.Enemy, speed float64) (float64, float64) {
	tx, ty := target.Position()
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
	entity.BaseEntity
	Vx, Vy float64
	Speed  float64
	Target enemy.Enemy
	Alive  bool

	Damage int
	Bounce int
}

// NewProjectile creates a new projectile aimed at the target.
func NewProjectile(x, y float64, target enemy.Enemy, dmg int, speed float64, bounce int) *Projectile {
	vx, vy := calcIntercept(x, y, target, speed)
	w, h := assets.ImgProjectile.Bounds().Dx(), assets.ImgProjectile.Bounds().Dy()
	return &Projectile{
		BaseEntity: entity.BaseEntity{
			Pos:          core.Point{X: x, Y: y},
			Width:        w,
			Height:       h,
			Sprite:       assets.ImgProjectile,
			FrameAnchorX: float64(w) / 2,
			FrameAnchorY: float64(h) / 2,
		},
		Vx:     vx,
		Vy:     vy,
		Speed:  speed,
		Target: target,
		Alive:  true,
		Damage: dmg,
		Bounce: bounce,
	}
}

// Update moves the projectile
func (p *Projectile) Update(dt float64) {
	p.Pos.X += p.Vx * p.Speed * dt
	p.Pos.Y += p.Vy * p.Speed * dt
	if p.Target != nil && p.Target.Alive() {
		tx, ty := p.Target.Position()
		dx := tx - p.Pos.X
		dy := ty - p.Pos.Y
		if math.Hypot(dx, dy) < 16 {
			if p.Target.Alive() {
				p.Target.ApplyDamage(p.Damage)
				if p.Bounce > 0 {
					p.Bounce--
					p.Target = nil // reset target to allow bouncing
				} else {
					p.Destroy()
					// pick new target: closest alive mob
					// var next Enemy
					// dist := math.MaxFloat64
					// for _, m := range p.game.mobs {
					// 	if m.Alive() && m != p.Target {
					// 		mx, my := m.Position()
					// 		dx := mx - p.Position.X
					// 		dy := my - p.Position.Y
					// 		d := math.Hypot(dx, dy)
					// 		if d < dist {
					// 			dist = d
					// 			next = m
					// 		}
					// 	}
					// }
					// if next != nil {
					// 	p.Target = next
					// 	p.Vx, p.Vy = calcIntercept(p.Position.X, p.Position.Y, next, p.Speed)
					// 	return
					// }
				}
				p.Alive = false
			}
		}
	}
	if p.Pos.X < -10 || p.Pos.X > 1930 || p.Pos.Y < -10 || p.Pos.Y > 1090 {
		p.Destroy()
	}
}
