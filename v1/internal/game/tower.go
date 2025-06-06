package game

import (
	"image/color"
	"math"
	"math/rand"
	"unicode"

	"github.com/hajimehoshi/ebiten/v2"
)

// TowerType represents the variety of tower with unique stats.
type TowerType int

const (
	TowerBasic TowerType = iota
	TowerSniper
	TowerRapid
)

// Tower represents a stationary auto-firing tower.
type Tower struct {
	BaseEntity
	cooldownTimer CooldownTimer // Use proper timer instead of float64
	rate          float64       // seconds between shots
	rangeDst      float64
	game          *Game
	rangeImg      *ebiten.Image

	// Type of tower (basic, sniper, rapid-fire)
	towerType TowerType

	// Two-queue ammo system
	ammoQueue    []bool // ready-to-fire ammunition (true = loaded, false = empty)
	reloadQueue  []rune // letters that need to be typed to reload
	ammoCapacity int    // maximum size of ammoQueue
	damage       int
	projectiles  int
	bounce       int
	level        int
	jammed       bool
	jammedLetter rune // preserve letter when jammed
	foresight    int  // number of reload letters to preview

	// Advanced reload mechanics
	reloadSeq       []rune // optional fixed reload sequence
	reloadIdx       int    // index into reloadSeq
	challengeWord   []rune // special challenge sequence
	challengeIdx    int
	challengeActive bool
	bonusTimer      CooldownTimer // Use timer for bonus duration
	damageBonus     int
}

// NewTower creates a basic tower at the given position.
func NewTower(g *Game, x, y float64) *Tower {
	return NewTowerWithTypeAndLevel(g, x, y, TowerBasic, 1)
}

// NewTowerWithLevel creates a basic Tower at the given position and level.
func NewTowerWithLevel(g *Game, x, y float64, level int) *Tower {
	return NewTowerWithTypeAndLevel(g, x, y, TowerBasic, level)
}

// NewTowerWithType creates a tower of the specified type at the given position.
func NewTowerWithType(g *Game, x, y float64, tt TowerType) *Tower {
	return NewTowerWithTypeAndLevel(g, x, y, tt, 1)
}

// NewTowerWithTypeAndLevel creates a tower of the specified type and level.
func NewTowerWithTypeAndLevel(g *Game, x, y float64, tt TowerType, level int) *Tower {
	if level < 1 {
		level = 1
	}
	if level > 5 {
		level = 5
	}
	w, h := ImgTower.Bounds().Dx(), ImgTower.Bounds().Dy()

	// Get config from game if available, otherwise use default
	cfg := DefaultConfig
	if g != nil && g.cfg != nil {
		cfg = *g.cfg
	}

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
		cooldownTimer: NewCooldownTimer(cfg.TowerFireRate * 3.0), // Much slower firing
		rangeDst:      cfg.TowerRange,
		game:          g,
		ammoCapacity:  cfg.TowerAmmoCapacity,
		damage:        cfg.TowerDamage,
		projectiles:   cfg.TowerProjectiles,
		bounce:        cfg.TowerBounce,
		level:         level,
		jammed:        false,
		foresight:     5,
		damageBonus:   0,
		towerType:     tt,
		bonusTimer:    NewCooldownTimer(5.0),
	}

	// Initialize ammo queue with full capacity
	t.ammoQueue = make([]bool, t.ammoCapacity)
	for i := range t.ammoQueue {
		t.ammoQueue[i] = true
	}
	t.reloadQueue = make([]rune, 0)

	// Generate range image
	t.rangeImg = generateRangeImage(t.rangeDst)

	// Apply game config if available
	if g.cfg != nil {
		t.ApplyConfig(*g.cfg)
	}

	// Apply tower type-specific stats AFTER config application
	// to ensure tower types maintain their unique characteristics
	switch tt {
	case TowerSniper:
		t.damage *= 3
		t.rangeDst *= 2.0                                           // Ensure sniper has significantly longer range
		t.cooldownTimer.SetInterval(t.cooldownTimer.interval * 2.5) // slower fire rate
		t.rate *= 2.5                                               // update rate for display/upgrades
		t.ammoCapacity = 3
	case TowerRapid:
		if t.damage > 1 {
			t.damage /= 2
		}
		t.rangeDst *= 0.7
		t.cooldownTimer.SetInterval(t.cooldownTimer.interval * 0.4) // faster fire rate
		t.rate *= 0.4                                               // update rate for display/upgrades
		t.ammoCapacity = 6
	}

	// Apply level upgrades after all type-specific modifications
	t.applyLevel()

	// Regenerate range image with the final range distance
	t.rangeImg = generateRangeImage(t.rangeDst)

	// Ensure ammo capacity is consistent with the queue size
	if len(t.ammoQueue) != t.ammoCapacity {
		newAmmoQueue := make([]bool, t.ammoCapacity)
		for i := 0; i < t.ammoCapacity && i < len(t.ammoQueue); i++ {
			newAmmoQueue[i] = t.ammoQueue[i]
		}
		for i := len(t.ammoQueue); i < t.ammoCapacity; i++ {
			newAmmoQueue[i] = true
		}
		t.ammoQueue = newAmmoQueue
	}

	return t
}

// applyLevel adjusts base stats according to the tower's level.
func (t *Tower) applyLevel() {
	switch t.level {
	case 2:
		t.damage++
		t.rangeDst += 50
		t.rate *= 0.9
		t.ammoCapacity++
	case 3:
		t.damage += 2
		t.rangeDst += 100
		t.rate *= 0.8
		t.ammoCapacity += 2
	case 4:
		t.damage += 3
		t.rangeDst += 150
		t.rate *= 0.7
		t.ammoCapacity += 3
	case 5:
		t.damage += 4
		t.rangeDst += 200
		t.rate *= 0.6
		t.ammoCapacity += 4
	}
	// Note: range image will be regenerated in the main constructor
}

func (t *Tower) randomReloadLetter() rune {
	if len(t.reloadSeq) > 0 {
		r := t.reloadSeq[t.reloadIdx%len(t.reloadSeq)]
		t.reloadIdx++
		return r
	}
	if t.game != nil {
		return t.game.randomReloadLetter()
	}
	// fallback if game is nil
	if rand.Intn(2) == 0 {
		return 'f'
	}
	return 'j'
}

// ApplyModifiers multiplies tower stats according to the provided modifiers.
func (t *Tower) ApplyModifiers(mod TowerModifiers) {
	if mod.DamageMult != 0 {
		t.damage = int(float64(t.damage) * mod.DamageMult)
	}
	if mod.RangeMult != 0 {
		t.rangeDst *= mod.RangeMult
		t.rangeImg = generateRangeImage(t.rangeDst)
	}
	if mod.FireRateMult != 0 {
		t.rate *= mod.FireRateMult
		if t.rate < 0.01 {
			t.rate = 0.01
		}
	}
	if mod.AmmoAdd != 0 {
		t.ammoCapacity += mod.AmmoAdd
		if t.ammoCapacity < 1 {
			t.ammoCapacity = 1
		}
		newAmmo := make([]bool, t.ammoCapacity)
		for i := 0; i < t.ammoCapacity && i < len(t.ammoQueue); i++ {
			newAmmo[i] = t.ammoQueue[i]
		}
		for i := len(t.ammoQueue); i < t.ammoCapacity; i++ {
			newAmmo[i] = true
		}
		t.ammoQueue = newAmmo
	}
}

// SetReloadSequence sets a fixed reload sequence for the tower.
func (t *Tower) SetReloadSequence(seq []rune) {
	t.reloadSeq = seq
	t.reloadIdx = 0
}

// StartReloadChallenge activates a special reload challenge word.
func (t *Tower) StartReloadChallenge(word string) {
	if word == "" {
		return
	}
	t.challengeWord = []rune(word)
	t.challengeIdx = 0
	t.challengeActive = true
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

	if !t.challengeActive && len(t.reloadQueue) == 0 && rand.Float64() < 0.05 {
		t.StartReloadChallenge("bonus")
	}
}

// Update handles tower firing logic.
func (t *Tower) Update(dt float64) {
	typed := t.game.input.TypedChars()

	if !t.bonusTimer.Ready() {
		t.bonusTimer.Tick(dt)
	}

	// Handle jam clearing
	if t.jammed {
		if t.game.input.Backspace() {
			t.jammed = false
		}
		// Jammed towers can still fire, just can't reload
	}

	// Handle active challenge before reload typing
	if t.challengeActive {
		for _, r := range typed {
			if unicode.ToLower(r) == unicode.ToLower(t.challengeWord[t.challengeIdx]) {
				t.challengeIdx++
				if t.challengeIdx >= len(t.challengeWord) {
					t.challengeActive = false
					t.challengeIdx = 0
					t.bonusTimer.Reset()
					if t.game != nil {
						t.game.typing.Record(true)
					}
				}
			} else {
				t.challengeIdx = 0
				if t.game != nil {
					t.game.typing.Record(false)
				}
			}
		}
		// letters used for challenge shouldn't also be used for reload
		return
	}

	// Handle firing cooldown first
	t.cooldownTimer.Tick(dt)

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
				if t.game != nil {
					t.game.typing.Record(true)
				}
				break
			} else if len(t.reloadQueue) > 0 {
				// Wrong letter - jam the tower
				if t.game != nil {
					t.game.typing.Record(false)
					t.game.MistypeFeedback()
				}
				t.jammed = true
				t.jammedLetter = t.reloadQueue[0] // preserve current letter
				break
			}
		}
	}

	// Early return if cooldown is not ready
	if !t.cooldownTimer.Ready() {
		return
	}

	// Only fire if we have ammo
	if t.getAvailableAmmo() <= 0 {
		return
	}

	// Find all targets in range, sorted by distance
	type mobDist struct {
		m Enemy
		d float64
	}
	var targets []mobDist
	for _, m := range t.game.mobs {
		if !m.Alive() {
			continue
		}
		mx, my := m.Position()
		dx := mx - t.pos.X
		dy := my - t.pos.Y
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
		if targetMob == nil || !targetMob.Alive() {
			continue
		}

		if t.consumeAmmo() {
			dmg := t.damage
			if t.bonusTimer.Ready() {
				dmg += t.damageBonus
			}
			p := NewProjectile(t.game, t.pos.X, t.pos.Y, targetMob, dmg, speed, t.bounce)
			t.game.projectiles = append(t.game.projectiles, p)
			shotsFired++
		}
	}

	// Set cooldown only if we actually fired
	if shotsFired > 0 {
		mult := 1.0
		if t.game != nil {
			mult = t.game.typing.RateMultiplier()
			if t.game.sound != nil {
				t.game.sound.PlayBeep()
			}
		}
		t.cooldownTimer.SetInterval(t.cooldownTimer.interval * mult)
		t.cooldownTimer.Reset()
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

	previewQueue := make([]rune, 0, t.foresight)
	for i := 0; i < len(t.reloadQueue) && i < t.foresight; i++ {
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

// UpgradeForesight increases how many reload letters are previewed
func (t *Tower) UpgradeForesight(increase int) {
	if increase <= 0 {
		return
	}
	t.foresight += increase
	if t.foresight > 10 {
		t.foresight = 10
	}
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
