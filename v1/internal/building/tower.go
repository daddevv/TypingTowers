package entity

import (
	"image/color"
	"math/rand"

	"github.com/daddevv/type-defense/internal/assets"
	"github.com/daddevv/type-defense/internal/config"
	"github.com/daddevv/type-defense/internal/core"
	"github.com/daddevv/type-defense/internal/entity"
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
	entity.BaseEntity
	CooldownTimer core.CooldownTimer // Use proper timer instead of float64
	Rate          float64       // seconds between shots
	RangeDst      float64
	RangeImg      *ebiten.Image

	// Type of tower (basic, sniper, rapid-fire)
	TowerType TowerType

	// Two-queue ammo system
	AmmoQueue    []bool // ready-to-fire ammunition (true = loaded, false = empty)
	ReloadQueue  []rune // letters that need to be typed to reload
	AmmoCapacity int    // maximum size of AmmoQueue
	Damage       int
	Projectiles  int
	Bounce       int
	Level        int
	Jammed       bool
	JammedLetter rune // preserve letter when jammed
	Foresight    int  // number of reload letters to preview

	// Advanced reload mechanics
	ReloadSeq       []rune // optional fixed reload sequence
	ReloadIdx       int    // index into ReloadSeq
	ChallengeWord   []rune // special challenge sequence
	ChallengeIdx    int
	ChallengeActive bool
	BonusTimer      core.CooldownTimer // Use timer for bonus duration
	DamageBonus     int
}

// NewTower creates a basic tower at the given position.
func NewTower(x, y float64) *Tower {
	return NewTowerWithTypeAndLevel(x, y, TowerBasic, 1)
}

// NewTowerWithLevel creates a basic Tower at the given position and level.
func NewTowerWithLevel(x, y float64, level int) *Tower {
	return NewTowerWithTypeAndLevel(x, y, TowerBasic, level)
}

// NewTowerWithType creates a tower of the specified type at the given position.
func NewTowerWithType(x, y float64, tt TowerType) *Tower {
	return NewTowerWithTypeAndLevel(x, y, tt, 1)
}

// NewTowerWithTypeAndLevel creates a tower of the specified type and level.
func NewTowerWithTypeAndLevel(x, y float64, tt TowerType, level int) *Tower {
	if level < 1 {
		level = 1
	}
	if level > 5 {
		level = 5
	}
	w, h := assets.ImgTower.Bounds().Dx(), assets.ImgTower.Bounds().Dy()

	// Get config from game if available, otherwise use default
	cfg := config.DefaultConfig
	t := &Tower{
		BaseEntity: entity.BaseEntity{
			Position:          core.Point{X: x, Y: y},
			Width:        w,
			Height:       h,
			Frame:        assets.ImgTower,
			FrameAnchorX: float64(w) / 2,
			FrameAnchorY: float64(h) / 2,
			Static:       true,
		},
		CooldownTimer: core.NewCooldownTimer(cfg.TowerFireRate * 3.0), // Much slower firing
		RangeDst:      cfg.TowerRange,
		AmmoCapacity:  cfg.TowerAmmoCapacity,
		Damage:        cfg.TowerDamage,
		Projectiles:   cfg.TowerProjectiles,
		Bounce:        cfg.TowerBounce,
		Level:         level,
		Jammed:        false,
		Foresight:     5,
		DamageBonus:   0,
		TowerType:     tt,
		BonusTimer:    core.NewCooldownTimer(5.0),
	}

	// Initialize ammo queue with full capacity
	t.AmmoQueue = make([]bool, t.AmmoCapacity)
	for i := range t.AmmoQueue {
		t.AmmoQueue[i] = true
	}
	t.ReloadQueue = make([]rune, 0)

	// Generate range image
	t.RangeImg = generateRangeImage(t.RangeDst)

	// Apply tower type-specific stats AFTER config application
	// to ensure tower types maintain their unique characteristics
	switch tt {
	case TowerSniper:
		t.Damage *= 3
		t.RangeDst *= 2.0                                           // Ensure sniper has significantly longer range
		t.CooldownTimer.SetInterval(t.CooldownTimer.Interval * 2.5) // slower fire rate
		t.Rate *= 2.5                                               // update rate for display/upgrades
		t.AmmoCapacity = 3
	case TowerRapid:
		if t.Damage > 1 {
			t.Damage /= 2
		}
		t.RangeDst *= 0.7
		t.CooldownTimer.SetInterval(t.CooldownTimer.Interval * 0.4) // faster fire rate
		t.Rate *= 0.4                                               // update rate for display/upgrades
		t.AmmoCapacity = 6
	}

	// Apply level upgrades after all type-specific modifications
	t.applyLevel()

	// Regenerate range image with the final range distance
	t.RangeImg = generateRangeImage(t.RangeDst)

	// Ensure ammo capacity is consistent with the queue size
	if len(t.AmmoQueue) != t.AmmoCapacity {
		newAmmoQueue := make([]bool, t.AmmoCapacity)
		for i := 0; i < t.AmmoCapacity && i < len(t.AmmoQueue); i++ {
			newAmmoQueue[i] = t.AmmoQueue[i]
		}
		for i := len(t.AmmoQueue); i < t.AmmoCapacity; i++ {
			newAmmoQueue[i] = true
		}
		t.AmmoQueue = newAmmoQueue
	}

	return t
}

// applyLevel adjusts base stats according to the tower's level.
func (t *Tower) applyLevel() {
	switch t.Level {
	case 2:
		t.Damage++
		t.RangeDst += 50
		t.Rate *= 0.9
		t.AmmoCapacity++
	case 3:
		t.Damage += 2
		t.RangeDst += 100
		t.Rate *= 0.8
		t.AmmoCapacity += 2
	case 4:
		t.Damage += 3
		t.RangeDst += 150
		t.Rate *= 0.7
		t.AmmoCapacity += 3
	case 5:
		t.Damage += 4
		t.RangeDst += 200
		t.Rate *= 0.6
		t.AmmoCapacity += 4
	}
	// Note: range image will be regenerated in the main constructor
}

func (t *Tower) randomReloadLetter() rune {
	if len(t.ReloadSeq) > 0 {
		r := t.ReloadSeq[t.ReloadIdx%len(t.ReloadSeq)]
		t.ReloadIdx++
		return r
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
		t.Damage = int(float64(t.Damage) * mod.DamageMult)
	}
	if mod.RangeMult != 0 {
		t.RangeDst *= mod.RangeMult
		t.RangeImg = generateRangeImage(t.RangeDst)
	}
	if mod.FireRateMult != 0 {
		t.Rate *= mod.FireRateMult
		if t.Rate < 0.01 {
			t.Rate = 0.01
		}
	}
	if mod.AmmoAdd != 0 {
		t.AmmoCapacity += mod.AmmoAdd
		if t.AmmoCapacity < 1 {
			t.AmmoCapacity = 1
		}
		newAmmo := make([]bool, t.AmmoCapacity)
		for i := 0; i < t.AmmoCapacity && i < len(t.AmmoQueue); i++ {
			newAmmo[i] = t.AmmoQueue[i]
		}
		for i := len(t.AmmoQueue); i < t.AmmoCapacity; i++ {
			newAmmo[i] = true
		}
		t.AmmoQueue = newAmmo
	}
}

// SetReloadSequence sets a fixed reload sequence for the tower.
func (t *Tower) SetReloadSequence(seq []rune) {
	t.ReloadSeq = seq
	t.ReloadIdx = 0
}

// StartReloadChallenge activates a special reload challenge word.
func (t *Tower) StartReloadChallenge(word string) {
	if word == "" {
		return
	}
	t.ChallengeWord = []rune(word)
	t.ChallengeIdx = 0
	t.ChallengeActive = true
}

// ApplyConfig updates tower parameters based on the provided config.
func (t *Tower) ApplyConfig(cfg config.Config) {
	if cfg.TowerFireRate > 0 {
		t.Rate = cfg.TowerFireRate / 1000.0 // convert ms to seconds
	}
	if cfg.F > 0 {
		r := t.Rate / cfg.F
		if r < 0.01 {
			r = 0.01
		}
		t.Rate = r
	}
	if cfg.TowerRange > 0 {
		t.RangeDst = cfg.TowerRange
		t.RangeImg = generateRangeImage(t.RangeDst)
	}
	if cfg.TowerAmmoCapacity > 0 {
		t.AmmoCapacity = cfg.TowerAmmoCapacity
		newAmmoQueue := make([]bool, t.AmmoCapacity)
		for i := 0; i < t.AmmoCapacity && i < len(t.AmmoQueue); i++ {
			newAmmoQueue[i] = t.AmmoQueue[i]
		}
		for i := len(t.AmmoQueue); i < t.AmmoCapacity; i++ {
			newAmmoQueue[i] = true
		}
		t.AmmoQueue = newAmmoQueue
	}
	if cfg.TowerDamage > 0 {
		t.Damage = cfg.TowerDamage
	}
	if cfg.TowerProjectiles > 0 {
		t.Projectiles = cfg.TowerProjectiles
	}
	if cfg.TowerBounce >= 0 {
		t.Bounce = cfg.TowerBounce
	}
}

// getAvailableAmmo counts ready-to-fire ammunition
func (t *Tower) getAvailableAmmo() int {
	count := 0
	for _, loaded := range t.AmmoQueue {
		if loaded {
			count++
		}
	}
	return count
}

// consumeAmmo removes one ammunition from the queue (first available)
func (t *Tower) consumeAmmo() bool {
	for i := range t.AmmoQueue {
		if t.AmmoQueue[i] {
			t.AmmoQueue[i] = false
			return true
		}
	}
	return false
}

// FillReloadQueue adds random letters to reload queue for empty ammo slots
func (t *Tower) FillReloadQueue() {
	emptySlots := 0
	for _, loaded := range t.AmmoQueue {
		if !loaded {
			emptySlots++
		}
	}

	// Add letters to reload queue to match empty slots
	for len(t.ReloadQueue) < emptySlots {
		t.ReloadQueue = append(t.ReloadQueue, t.randomReloadLetter())
	}

	if !t.ChallengeActive && len(t.ReloadQueue) == 0 && rand.Float64() < 0.05 {
		t.StartReloadChallenge("bonus")
	}
}

// Update handles tower firing logic.
func (t *Tower) Update(dt float64) {
	// typed := t.game.input.TypedChars()

	if !t.BonusTimer.Ready() {
		t.BonusTimer.Tick(dt)
	}

	// Handle jam clearing
	if t.Jammed {
		// if t.game.input.Backspace() {
		// 	t.Jammed = false
		// }
		// Jammed towers can still fire, just can't reload
	}

	// Handle active challenge before reload typing
	if t.ChallengeActive {
		// for _, r := range typed {
		// 	if unicode.ToLower(r) == unicode.ToLower(t.ChallengeWord[t.ChallengeIdx]) {
		// 		t.ChallengeIdx++
		// 		if t.ChallengeIdx >= len(t.ChallengeWord) {
		// 			t.ChallengeActive = false
		// 			t.ChallengeIdx = 0
		// 			t.BonusTimer.Reset()
		// 			// if t.game != nil {
		// 			// 	t.game.typing.Record(true)
		// 			// }
		// 		}
		// 	} else {
		// 		t.ChallengeIdx = 0
		// 		// if t.game != nil {
		// 		// 	t.game.typing.Record(false)
		// 		// }
		// 	}
		// }
		// letters used for challenge shouldn't also be used for reload
		return
	}

	// Handle firing cooldown first
	t.CooldownTimer.Tick(dt)

	// Always ensure reload queue is populated for empty ammo slots
	t.FillReloadQueue()

	// Handle reload typing (only if not jammed and reload queue has letters)
	if !t.Jammed && len(t.ReloadQueue) > 0 {
		// for _, r := range typed {
		// 	if len(t.ReloadQueue) > 0 && unicode.ToLower(r) == t.ReloadQueue[0] {
		// 		// Successfully typed the first letter in reload queue
		// 		t.ReloadQueue = t.ReloadQueue[1:]

		// 		// Add ammo to first empty slot
		// 		for i := range t.AmmoQueue {
		// 			if !t.AmmoQueue[i] {
		// 				t.AmmoQueue[i] = true
		// 				break
		// 			}
		// 		}
		// 		if t.game != nil {
		// 			t.game.typing.Record(true)
		// 		}
		// 		break
		// 	} else if len(t.reloadQueue) > 0 {
		// 		// Wrong letter - jam the tower
		// 		if t.game != nil {
		// 			t.game.typing.Record(false)
		// 			t.game.MistypeFeedback()
		// 		}
		// 		t.jammed = true
		// 		t.jammedLetter = t.reloadQueue[0] // preserve current letter
		// 		break
		// 	}
		// }
	}

	// Early return if cooldown is not ready
	if !t.CooldownTimer.Ready() {
		return
	}

	// Only fire if we have ammo
	if t.getAvailableAmmo() <= 0 {
		return
	}

	// Find all targets in range, sorted by distance
	// type mobDist struct {
	// 	m Enemy
	// 	d float64
	// }
	// var targets []mobDist
	// for _, m := range t.game.mobs {
	// 	if !m.Alive() {
	// 		continue
	// 	}
	// 	mx, my := m.Position()
	// 	dx := mx - t.pos.X
	// 	dy := my - t.pos.Y
	// 	d := math.Hypot(dx, dy)
	// 	if d < t.rangeDst {
	// 		targets = append(targets, mobDist{m, d})
	// 	}
	// }

	// No targets, no firing
	// if len(targets) == 0 {
	// 	return
	// }

	// // Sort by distance ascending
	// if len(targets) > 1 {
	// 	// Simple insertion sort (small N)
	// 	for i := 1; i < len(targets); i++ {
	// 		j := i
	// 		for j > 0 && targets[j].d < targets[j-1].d {
	// 			targets[j], targets[j-1] = targets[j-1], targets[j]
	// 			j--
	// 		}
	// 	}
	// }

	// Determine how many shots to fire - limited by ammo, targets, and projectiles setting
	shots := t.Projectiles
	if shots < 1 {
		shots = 1
	}
	availableAmmo := t.getAvailableAmmo()
	if shots > availableAmmo {
		shots = availableAmmo
	}
	// if shots > len(targets) {
	// 	shots = len(targets)
	// }

	// Fire at the closest unique targets, one projectile per mob
	// speed := config.DefaultConfig.ProjectileSpeed
	// if t.game != nil && t.game.cfg != nil && t.game.cfg.ProjectileSpeed > 0 {
	// 	speed = t.game.cfg.ProjectileSpeed
	// }

	// shotsFired := 0
	// for i := 0; i < shots && i < len(targets); i++ {
	// 	targetMob := targets[i].m
	// 	if targetMob == nil || !targetMob.Alive() {
	// 		continue
	// 	}

	// 	if t.consumeAmmo() {
	// 		dmg := t.damage
	// 		if t.bonusTimer.Ready() {
	// 			dmg += t.damageBonus
	// 		}
	// 		p := NewProjectile(t.game, t.pos.X, t.pos.Y, targetMob, dmg, speed, t.bounce)
	// 		t.game.projectiles = append(t.game.projectiles, p)
	// 		shotsFired++
	// 	}
	// }

	// // Set cooldown only if we actually fired
	// if shotsFired > 0 {
	// 	mult := 1.0
	// 	if t.game != nil {
	// 		mult = t.game.typing.RateMultiplier()
	// 		if t.game.sound != nil {
	// 			t.game.sound.PlayBeep()
	// 		}
	// 	}
	// 	t.CooldownTimer.SetInterval(t.CooldownTimer.Interval() * mult)
	// 	t.CooldownTimer.Reset()
	// }
}

// GetAmmoStatus returns current ammo and capacity for HUD display
func (t *Tower) GetAmmoStatus() (int, int) {
	return t.getAvailableAmmo(), t.AmmoCapacity
}

// GetReloadStatus returns reload state info for HUD
func (t *Tower) GetReloadStatus() (bool, rune, []rune, float64, bool) {
	reloading := len(t.ReloadQueue) > 0
	var currentLetter rune
	if t.Jammed {
		currentLetter = t.JammedLetter
	} else if len(t.ReloadQueue) > 0 {
		currentLetter = t.ReloadQueue[0]
	}

	previewQueue := make([]rune, 0, t.Foresight)
	for i := 0; i < len(t.ReloadQueue) && i < t.Foresight; i++ {
		previewQueue = append(previewQueue, t.ReloadQueue[i])
	}

	return reloading, currentLetter, previewQueue, 0, t.Jammed // timer always 0 now
}

// UpgradeAmmoCapacity increases the tower's ammunition capacity
func (t *Tower) UpgradeAmmoCapacity(increase int) {
	if increase <= 0 {
		return
	}

	oldCapacity := t.AmmoCapacity
	t.AmmoCapacity += increase

	// Expand ammo queue and fill new slots with ammo
	newAmmoQueue := make([]bool, t.AmmoCapacity)
	copy(newAmmoQueue, t.AmmoQueue)
	for i := oldCapacity; i < t.AmmoCapacity; i++ {
		newAmmoQueue[i] = true // New slots start loaded
	}
	t.AmmoQueue = newAmmoQueue
}

// UpgradeForesight increases how many reload letters are previewed
func (t *Tower) UpgradeForesight(increase int) {
	if increase <= 0 {
		return
	}
	t.Foresight += increase
	if t.Foresight > 10 {
		t.Foresight = 10
	}
}

// Draw renders the tower and its range indicator.
func (t *Tower) Draw(screen *ebiten.Image) {
	if t.RangeImg != nil {
		op := &ebiten.DrawImageOptions{}
		w, h := t.RangeImg.Bounds().Dx(), t.RangeImg.Bounds().Dy()
		op.GeoM.Translate(t.Position.X-float64(w)/2, t.Position.Y-float64(h)/2)
		screen.DrawImage(t.RangeImg, op)
	}
	t.BaseEntity.Draw(screen)
}

// generateRangeImage creates a semi-transparent circle representing the tower's range.
func generateRangeImage(radius float64) *ebiten.Image {
	r := int(radius)
	img := ebiten.NewImage(r*2, r*2)
	clr := color.RGBA{0, 255, 0, 30}
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
