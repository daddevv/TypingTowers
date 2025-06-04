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
	
	// Separate ammo for shooting vs reload bullets being added
	shootingAmmo int // ammo available for firing projectiles
	reloadBullets int // bullets being added through typing (separate from shooting ammo)
	ammoCapacity int
	
	damage       int
	projectiles  int
	bounce       int
	reloadTime   int
	reloadTimer  int
	reloading    bool
	reloadLetter rune
	nextReloadLetter rune // preview of next letter
	jammed       bool
	jammedLetter rune // preserve letter when jammed

	lastFiredFrame int // track last frame when tower fired
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
		rate:         DefaultConfig.TowerFireRate,
		rangeDst:     DefaultConfig.TowerRange,
		game:         g,
		ammoCapacity: DefaultConfig.TowerAmmoCapacity,
		shootingAmmo: DefaultConfig.TowerAmmoCapacity, // start with full ammo
		reloadBullets: 0, // no bullets being reloaded initially
		damage:       DefaultConfig.TowerDamage,
		projectiles:  DefaultConfig.TowerProjectiles,
		bounce:       DefaultConfig.TowerBounce,
		reloadTime:   DefaultConfig.TowerReloadRate,
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
	if cfg.TowerFireRate > 0 {
		t.rate = cfg.TowerFireRate
	}
	if cfg.F > 0 {
		r := int(float64(t.rate) / cfg.F)
		if r < 1 {
			r = 1
		}
		t.rate = r
	}
	if cfg.TowerRange > 0 {
		t.rangeDst = cfg.TowerRange
		t.rangeImg = generateRangeImage(t.rangeDst)
	}
	if cfg.TowerAmmoCapacity > 0 {
		oldCapacity := t.ammoCapacity
		t.ammoCapacity = cfg.TowerAmmoCapacity
		// Adjust shooting ammo proportionally
		if oldCapacity > 0 {
			ratio := float64(t.shootingAmmo) / float64(oldCapacity)
			t.shootingAmmo = int(ratio * float64(t.ammoCapacity))
		} else {
			t.shootingAmmo = t.ammoCapacity
		}
	}
	if cfg.TowerDamage > 0 {
		t.damage = cfg.TowerDamage
	}
	if cfg.TowerProjectiles > 0 {
		t.projectiles = cfg.TowerProjectiles
	}
	if cfg.TowerBounce >= 0 {
		t.bounce = cfg.TowerBounce
	}
}

// Update handles tower firing logic.
func (t *Tower) Update() {
	typed := t.game.input.TypedChars()

	// Handle jam clearing
	if t.jammed {
		if t.game.input.Backspace() {
			t.jammed = false
			// Restore the letter that was being typed when jammed
			t.reloadLetter = t.jammedLetter
		}
		// Jammed towers can still fire, just can't reload
	}

	// Handle firing cooldown first - this must decrement before anything else
	if t.cooldown > 0 {
		t.cooldown--
	}

	// Start reloading if not at capacity and not already reloading
	// Only start reload if we're not in firing cooldown
	if !t.reloading && t.shootingAmmo < t.ammoCapacity && t.cooldown == 0 {
		t.startReload()
	}

	// Handle reload typing (only if not jammed)
	if t.reloading && !t.jammed {
		if t.reloadTimer > 0 {
			t.reloadTimer--
		} else {
			for _, r := range typed {
				if unicode.ToLower(r) == t.reloadLetter {
					// Successfully typed reload letter
					t.reloadBullets++
					t.shootingAmmo++
					
					if t.shootingAmmo >= t.ammoCapacity {
						// Fully reloaded - but don't override firing cooldown
						t.reloading = false
						// Only set cooldown if we're not already in firing cooldown
						if t.cooldown == 0 {
							t.cooldown = t.rate
						}
					} else {
						// Set up next reload letter
						t.reloadLetter = t.nextReloadLetter
						t.nextReloadLetter = randomReloadLetter()
					}
					t.reloadTimer = t.reloadTime
					break
				} else {
					// Wrong letter - jam the tower
					t.jammed = true
					t.jammedLetter = t.reloadLetter // preserve current letter
					break
				}
			}
		}
	}

	// Early return if cooldown or reload timer is active
	if t.cooldown > 0 || t.reloadTimer > 0 {
		return
	}

	// Prevent firing more than once per frame
	currentFrame := t.game.currentFrame
	if t.lastFiredFrame == currentFrame {
		return
	}

	// Only fire if we have ammo
	if t.shootingAmmo <= 0 {
		return
	}

	// Find all targets in range, sorted by distance
	type mobDist struct {
		m *Mob
		d float64
	}
	var targets []mobDist
	for _, m := range t.game.mobs {
		if !m.alive {
			continue
		}
		dx := m.pos.X - t.pos.X
		dy := m.pos.Y - t.pos.Y
		d := math.Hypot(dx, dy)
		if d < t.rangeDst {
			targets = append(targets, mobDist{m, d})
		}
	}
	
	// No targets, no firing
	if len(targets) == 0 {
		return
	}
	
	// Sort by distance ascending
	if len(targets) > 1 {
		// Simple insertion sort (small N)
		for i := 1; i < len(targets); i++ {
			j := i
			for j > 0 && targets[j].d < targets[j-1].d {
				targets[j], targets[j-1] = targets[j-1], targets[j]
				j--
			}
		}
	}

	// Determine how many shots to fire - limited by ammo, targets, and projectiles setting
	shots := t.projectiles
	if shots < 1 {
		shots = 1
	}
	if shots > t.shootingAmmo {
		shots = t.shootingAmmo
	}
	if shots > len(targets) {
		shots = len(targets)
	}

	// Fire at the closest unique targets, one projectile per mob
	speed := DefaultConfig.ProjectileSpeed
	if t.game != nil && t.game.cfg != nil && t.game.cfg.ProjectileSpeed > 0 {
		speed = t.game.cfg.ProjectileSpeed
	}
	
	shotsFired := 0
	for i := 0; i < shots && i < len(targets); i++ {
		targetMob := targets[i].m
		if targetMob == nil || !targetMob.alive {
			continue
		}
		
		p := NewProjectile(t.game, t.pos.X, t.pos.Y, targetMob, t.damage, speed, t.bounce)
		t.game.projectiles = append(t.game.projectiles, p)
		t.shootingAmmo--
		shotsFired++
	}
	
	// Set cooldown only if we actually fired
	if shotsFired > 0 {
		t.cooldown = t.rate
		t.lastFiredFrame = currentFrame
	}
}

func (t *Tower) startReload() {
	t.reloading = true
	t.reloadTimer = t.reloadTime
	// Don't reset cooldown to 0 - preserve firing cooldown
	t.reloadLetter = randomReloadLetter()
	t.nextReloadLetter = randomReloadLetter()
	t.reloadBullets = 0
}

// GetAmmoStatus returns current ammo and capacity for HUD display
func (t *Tower) GetAmmoStatus() (int, int) {
	return t.shootingAmmo, t.ammoCapacity
}

// GetReloadStatus returns reload state info for HUD
func (t *Tower) GetReloadStatus() (bool, rune, rune, int, bool) {
	return t.reloading, t.reloadLetter, t.nextReloadLetter, t.reloadTimer, t.jammed
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
