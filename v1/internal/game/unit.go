package game

import (
	"math"
)

// Footman represents a basic friendly melee unit.
type Footman struct {
	BaseEntity
	speed    float64
	damage   int
	health   int
	game     *Game
	target   Enemy
	cooldown float64
}

// NewFootman creates a new footman instance.
func NewFootman(g *Game, x, y float64) *Footman {
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
		speed:  1.0,
		damage: 1,
		health: 5,
		game:   g,
	}
}

// Alive reports whether the footman is alive.
func (f *Footman) Alive() bool { return f.health > 0 }

// Damage applies damage to the footman.
func (f *Footman) Damage(d int) {
	f.health -= d
}

// closestEnemy returns the nearest alive enemy.
func (f *Footman) closestEnemy() Enemy {
	var target Enemy
	dist := math.MaxFloat64
	for _, m := range f.game.mobs {
		if !m.Alive() {
			continue
		}
		mx, my := m.Position()
		d := math.Hypot(mx-f.pos.X, my-f.pos.Y)
		if d < dist {
			dist = d
			target = m
		}
	}
	return target
}

// Update moves the footman toward the nearest enemy and attacks when close.
func (f *Footman) Update(dt float64) error {
	if !f.Alive() {
		return nil
	}
	if f.cooldown > 0 {
		f.cooldown -= dt
		if f.cooldown < 0 {
			f.cooldown = 0
		}
	}
	if f.target == nil || !f.target.Alive() {
		f.target = f.closestEnemy()
	}
	if f.target == nil {
		return nil
	}
	tx, ty := f.target.Position()
	dx := tx - f.pos.X
	dy := ty - f.pos.Y
	dist := math.Hypot(dx, dy)
	if dist > 0 {
		f.pos.X += dx / dist * f.speed * dt
		f.pos.Y += dy / dist * f.speed * dt
	}
	if dist < 16 && f.cooldown <= 0 {
		f.target.Damage(f.damage)
		f.cooldown = 1.0
	}
	return nil
}
