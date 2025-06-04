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
	cooldown float64 // seconds
	rate     float64 // seconds between shots
	rangeDst float64
	game     *Game
	rangeImg *ebiten.Image

	// Two-queue ammo system
	ammoQueue    []bool // ready-to-fire ammunition (true = loaded, false = empty)
	reloadQueue  []rune // letters that need to be typed to reload
	ammoCapacity int    // maximum size of ammoQueue
	damage       int
	projectiles  int
	bounce       int
	jammed       bool
	jammedLetter rune // preserve letter when jammed
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
		rate:         DefaultConfig.TowerFireRate, // seconds
		rangeDst:     DefaultConfig.TowerRange,
		game:         g,
		ammoCapacity: DefaultConfig.TowerAmmoCapacity,
		damage:       DefaultConfig.TowerDamage,
		projectiles:  DefaultConfig.TowerProjectiles,
		bounce:       DefaultConfig.TowerBounce,
		jammed:       false,
	}

	// Initialize ammo queue with full capacity
	t.ammoQueue = make([]bool, t.ammoCapacity)
	for i := range t.ammoQueue {
		t.ammoQueue[i] = true
	}
	t.reloadQueue = make([]rune, 0)

	t.rangeImg = generateRangeImage(t.rangeDst)
	if g.cfg != nil {
		t.ApplyConfig(*g.cfg)
	}
	return t
}

func (t *Tower) randomReloadLetter() rune {
	if t.game != nil {
		return t.game.randomReloadLetter()
	}
	// fallback if game is nil
	if rand.Intn(2) == 0 {
		return 'f'
	}
	return 'j'
}

// ApplyConfig updates tower parameters based on the provided config.
func (t *Tower) ApplyConfig(cfg Config) {
	if cfg.TowerFireRate > 0 {
		t.rate = cfg.TowerFireRate / 1000.0 // convert ms to seconds
	}
	if cfg.F > 0 {
		r := t.rate / cfg.F
		if r < 0.01 {
			r = 0.01
		}
		t.rate = r
	}
	if cfg.TowerRange > 0 {
		t.rangeDst = cfg.TowerRange
		t.rangeImg = generateRangeImage(t.rangeDst)
	}
	if cfg.TowerAmmoCapacity > 0 {
		t.ammoCapacity = cfg.TowerAmmoCapacity
		newAmmoQueue := make([]bool, t.ammoCapacity)
		for i := 0; i < t.ammoCapacity && i < len(t.ammoQueue); i++ {
			newAmmoQueue[i] = t.ammoQueue[i]
		}
		for i := len(t.ammoQueue); i < t.ammoCapacity; i++ {
			newAmmoQueue[i] = true
		}
		t.ammoQueue = newAmmoQueue
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

// getAvailableAmmo counts ready-to-fire ammunition
func (t *Tower) getAvailableAmmo() int {
	count := 0
	for _, loaded := range t.ammoQueue {
		if loaded {
			count++
		}
	}
	return count
}

// consumeAmmo removes one ammunition from the queue (first available)
func (t *Tower) consumeAmmo() bool {
	for i := range t.ammoQueue {
		if t.ammoQueue[i] {
			t.ammoQueue[i] = false
			return true
		}
	}
	return false
}

// fillReloadQueue adds random letters to reload queue for empty ammo slots
func (t *Tower) fillReloadQueue() {
	emptySlots := 0
	for _, loaded := range t.ammoQueue {
		if !loaded {
			emptySlots++
		}
	}

	// Add letters to reload queue to match empty slots
	for len(t.reloadQueue) < emptySlots {
		t.reloadQueue = append(t.reloadQueue, t.randomReloadLetter())
	}
}

// Update handles tower firing logic.
func (t *Tower) Update(dt float64) {
	typed := t.game.input.TypedChars()

	// Handle jam clearing
	if t.jammed {
		if t.game.input.Backspace() {
			t.jammed = false
		}
		// Jammed towers can still fire, just can't reload
	}

	// Handle firing cooldown first - this must decrement before anything else
	if t.cooldown > 0 {
		t.cooldown -= dt
		if t.cooldown < 0 {
			t.cooldown = 0
		}
	}

	// Always ensure reload queue is populated for empty ammo slots
	t.fillReloadQueue()

	// Handle reload typing (only if not jammed and reload queue has letters)
	if !t.jammed && len(t.reloadQueue) > 0 {
		for _, r := range typed {
			if len(t.reloadQueue) > 0 && unicode.ToLower(r) == t.reloadQueue[0] {
				// Successfully typed the first letter in reload queue
				t.reloadQueue = t.reloadQueue[1:]

				// Add ammo to first empty slot
				for i := range t.ammoQueue {
					if !t.ammoQueue[i] {
						t.ammoQueue[i] = true
						break
					}
				}
				break
			} else if len(t.reloadQueue) > 0 {
				// Wrong letter - jam the tower
				t.jammed = true
				t.jammedLetter = t.reloadQueue[0] // preserve current letter
				break
			}
		}
	}

	// Early return if cooldown or reload timer is active
	if t.cooldown > 0 {
		return
	}

	// Only fire if we have ammo
	if t.getAvailableAmmo() <= 0 {
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
	availableAmmo := t.getAvailableAmmo()
	if shots > availableAmmo {
		shots = availableAmmo
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

		if t.consumeAmmo() {
			p := NewProjectile(t.game, t.pos.X, t.pos.Y, targetMob, t.damage, speed, t.bounce)
			t.game.projectiles = append(t.game.projectiles, p)
			shotsFired++
		}
	}

	// Set cooldown only if we actually fired
	if shotsFired > 0 {
		t.cooldown = t.rate
	}
}

// GetAmmoStatus returns current ammo and capacity for HUD display
func (t *Tower) GetAmmoStatus() (int, int) {
	return t.getAvailableAmmo(), t.ammoCapacity
}

// GetReloadStatus returns reload state info for HUD
func (t *Tower) GetReloadStatus() (bool, rune, []rune, float64, bool) {
	reloading := len(t.reloadQueue) > 0
	var currentLetter rune
	if t.jammed {
		currentLetter = t.jammedLetter
	} else if len(t.reloadQueue) > 0 {
		currentLetter = t.reloadQueue[0]
	}

	// Return up to 5 letters for preview (future "Foresight" upgrade)
	previewQueue := make([]rune, 0, 5)
	for i := 0; i < len(t.reloadQueue) && i < 5; i++ {
		previewQueue = append(previewQueue, t.reloadQueue[i])
	}

	return reloading, currentLetter, previewQueue, 0, t.jammed // timer always 0 now
}

// UpgradeAmmoCapacity increases the tower's ammunition capacity
func (t *Tower) UpgradeAmmoCapacity(increase int) {
	if increase <= 0 {
		return
	}

	oldCapacity := t.ammoCapacity
	t.ammoCapacity += increase

	// Expand ammo queue and fill new slots with ammo
	newAmmoQueue := make([]bool, t.ammoCapacity)
	copy(newAmmoQueue, t.ammoQueue)
	for i := oldCapacity; i < t.ammoCapacity; i++ {
		newAmmoQueue[i] = true // New slots start loaded
	}
	t.ammoQueue = newAmmoQueue
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
