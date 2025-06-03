package game

import "math"

// Tower represents a stationary auto-firing tower.
type Tower struct {
	BaseEntity
	cooldown int
	rate     int
	rangeDst float64
	game     *Game
}

// NewTower creates a new Tower at the given position.
func NewTower(g *Game, x, y float64) *Tower {
	w, h := ImgTower.Bounds().Dx(), ImgTower.Bounds().Dy()
	return &Tower{
		BaseEntity: BaseEntity{
			pos:          Point{x, y},
			width:        w,
			height:       h,
			frame:        ImgTower,
			frameAnchorX: float64(w) / 2,
			frameAnchorY: float64(h) / 2,
			static:       true,
		},
		rate:     60,
		rangeDst: 300,
		game:     g,
	}
}

// Update handles tower firing logic.
func (t *Tower) Update() {
	if t.cooldown > 0 {
		t.cooldown--
	}
	if t.cooldown > 0 {
		return
	}
	var target *Mob
	dist := math.MaxFloat64
	for _, m := range t.game.mobs {
		dx := m.pos.X - t.pos.X
		dy := m.pos.Y - t.pos.Y
		d := math.Hypot(dx, dy)
		if d < t.rangeDst && d < dist {
			dist = d
			target = m
		}
	}
	if target != nil {
		p := NewProjectile(t.pos.X, t.pos.Y, target)
		t.game.projectiles = append(t.game.projectiles, p)
		t.cooldown = t.rate
	}
}
