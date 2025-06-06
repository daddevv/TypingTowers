package game

import (
	"github.com/daddevv/type-defense/internal/entity"
	"math"
)

// calcIntercept returns a normalized direction vector from the shooter position
// to where the projectile should aim in order to intercept the moving target.
func calcIntercept(px, py float64, target Enemy, speed float64) (float64, float64) {
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
	vx, vy float64
	speed  float64
	target Enemy
	alive  bool

	damage int
	bounce int
	game   *Game
}

// NewProjectile creates a new projectile aimed at the target.
func NewProjectile(g *Game, x, y float64, target Enemy, dmg int, speed float64, bounce int) *Projectile {
	vx, vy := calcIntercept(x, y, target, speed)
	w, h := ImgProjectile.Bounds().Dx(), ImgProjectile.Bounds().Dy()
	return &Projectile{
		BaseEntity: entity.BaseEntity{
			pos:          entity.Point{x, y},
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
		damage: dmg,
		bounce: bounce,
		game:   g,
	}
}

// Update moves the projectile and checks collision.
func (p *Projectile) Update(dt float64) {
	p.pos.X += p.vx * p.speed * dt
	p.pos.Y += p.vy * p.speed * dt
	if p.target != nil && p.target.Alive() {
		tx, ty := p.target.Position()
		dx := tx - p.pos.X
		dy := ty - p.pos.Y
		if math.Hypot(dx, dy) < 16 {
			if p.target.Alive() {
				p.target.Damage(p.damage)
				if p.bounce > 0 && p.game != nil {
					p.bounce--
					// pick new target: closest alive mob
					var next Enemy
					dist := math.MaxFloat64
					for _, m := range p.game.mobs {
						if m.Alive() && m != p.target {
							mx, my := m.Position()
							dx := mx - p.pos.X
							dy := my - p.pos.Y
							d := math.Hypot(dx, dy)
							if d < dist {
								dist = d
								next = m
							}
						}
					}
					if next != nil {
						p.target = next
						p.vx, p.vy = calcIntercept(p.pos.X, p.pos.Y, next, p.speed)
						return
					}
				}
				p.alive = false
			}
		}
	}
	if p.pos.X < -10 || p.pos.X > 1930 || p.pos.Y < -10 || p.pos.Y > 1090 {
		p.alive = false
	}
}
