package game

import (
	"image/color"
	"math"
	"math/rand"
	"unicode"

	"github.com/hajimehoshi/ebiten/v2"
)

// Tower represents a stationary auto-firing tower.
type Tower struct {
	BaseEntity
	cooldown     int
	rate         int
	rangeDst     float64
	game         *Game
	rangeImg     *ebiten.Image
	ammo         []struct{}
	ammoCapacity int
	reloadTime   int
	reloadTimer  int
	reloading    bool
	reloadLetter rune
	jammed       bool
}

// NewTower creates a new Tower at the given position.
func NewTower(g *Game, x, y float64) *Tower {
	w, h := ImgTower.Bounds().Dx(), ImgTower.Bounds().Dy()
	t := &Tower{
		BaseEntity: BaseEntity{
			pos:          Point{x, y},
			width:        w,
			height:       h,
			frame:        ImgTower,
			frameAnchorX: float64(w) / 2,
			frameAnchorY: float64(h) / 2,
			static:       true,
		},
		rate:         100,
		rangeDst:     500,
		game:         g,
		ammoCapacity: 5,
		ammo:         make([]struct{}, 5),
		reloadTime:   0,
		jammed:       false,
	}
	t.rangeImg = generateRangeImage(t.rangeDst)
	if g.cfg != nil {
		t.ApplyConfig(*g.cfg)
	}
	return t
}

func randomReloadLetter() rune {
	// randomly return either 'f' or 'j' for reload prompts
	if rand.Intn(2) == 0 {
		return 'f'
	}
	return 'j'
}

// ApplyConfig updates tower parameters based on the provided config.
func (t *Tower) ApplyConfig(cfg Config) {
	rate := 100
	if cfg.F > 0 {
		rate = int(float64(rate) / cfg.F)
		if rate < 1 {
			rate = 1
		}
	}
	t.rate = rate
}

// Update handles tower firing logic.
func (t *Tower) Update() {
	typed := t.game.input.TypedChars()

	if t.jammed {
		if t.game.input.Backspace() {
			t.jammed = false
		}
		return
	}

	if !t.reloading && len(t.ammo) < t.ammoCapacity {
		t.startReload()
	}

	if t.reloading {
		if t.reloadTimer > 0 {
			t.reloadTimer--
		} else {
			for _, r := range typed {
				if unicode.ToLower(r) == t.reloadLetter {
					t.ammo = append(t.ammo, struct{}{})
					if len(t.ammo) >= t.ammoCapacity {
						t.reloading = false
						t.cooldown = t.rate
					} else {
						t.reloadLetter = randomReloadLetter()
					}
					t.reloadTimer = t.reloadTime
					break
				} else {
					t.jammed = true
					t.reloading = false
					break
				}
			}
		}
	}

	if t.cooldown > 0 {
		t.cooldown--
	}
	if t.cooldown > 0 || t.reloadTimer > 0 {
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
	if target != nil && len(t.ammo) > 0 {
		p := NewProjectile(t.pos.X, t.pos.Y, target)
		t.game.projectiles = append(t.game.projectiles, p)
		t.cooldown = t.rate
		t.ammo = t.ammo[1:]
	}
}

func (t *Tower) startReload() {
	t.reloading = true
	t.reloadTimer = t.reloadTime
	t.cooldown = 0
	t.reloadLetter = randomReloadLetter()
}

// Draw renders the tower and its range indicator.
func (t *Tower) Draw(screen *ebiten.Image) {
	if t.rangeImg != nil {
		op := &ebiten.DrawImageOptions{}
		w, h := t.rangeImg.Bounds().Dx(), t.rangeImg.Bounds().Dy()
		op.GeoM.Translate(t.pos.X-float64(w)/2, t.pos.Y-float64(h)/2)
		screen.DrawImage(t.rangeImg, op)
	}
	t.BaseEntity.Draw(screen)
}

// generateRangeImage creates a semi-transparent circle representing the tower's range.
func generateRangeImage(radius float64) *ebiten.Image {
	r := int(radius)
	img := ebiten.NewImage(r*2, r*2)
	clr := color.RGBA{0, 255, 0, 80}
	rr := r * r
	inner := (r - 1) * (r - 1)
	for x := 0; x < r*2; x++ {
		for y := 0; y < r*2; y++ {
			dx := x - r
			dy := y - r
			d := dx*dx + dy*dy
			if d <= rr && d >= inner {
				img.Set(x, y, clr)
			}
		}
	}
	return img
}
