package mob

import (
	"math"

	"github.com/daddevv/type-defense/internal/assets"
	"github.com/daddevv/type-defense/internal/core"
	"github.com/daddevv/type-defense/internal/entity"
	"github.com/daddevv/type-defense/internal/structure"
)

// Mob represents a basic enemy moving left.
type Mob struct {
	entity.BaseEntity
	Speed      float64
	AnimTicker float64
	Alive      bool
	VX, VY     float64
	Target     *structure.Base
	Health     int

	Armor       int
	Shield      int
	Burst       float64
	BurstTimer  core.CooldownTimer // Use proper timer for burst cooldown
	BurstActive core.CooldownTimer // Use proper timer for burst duration
	MobType     entity.MobType
}

// NewMob returns a new mob at the given position.
func NewMob(x, y float64, target *structure.Base, hp int, speed float64) *Mob {
	w, h := assets.ImgMobA.Bounds().Dx(), assets.ImgMobA.Bounds().Dy()
	return &Mob{
		BaseEntity: entity.BaseEntity{
			Position:          core.Point{X: x, Y: y},
			Width:        w,
			Height:       h,
			Frame:        assets.ImgMobA,
			FrameAnchorX: float64(w) / 2,
			FrameAnchorY: float64(h) / 2,
		},
		Speed:   speed,
		Alive:   true,
		Target:  target,
		Health:  hp,
		Armor:   0,
		Shield:  0,
		Burst:   0,
		MobType: entity.MobBasic,
	}
}

// NewArmoredMob creates a mob with armor reducing incoming damage.
func NewArmoredMob(x, y float64, target *structure.Base, hp, armor int, speed float64) *Mob {
	m := NewMob(x, y, target, hp, speed)
	m.Armor = armor
	m.MobType = entity.MobArmored
	return m
}

// NewFastMob creates a mob with periodic speed bursts.
func NewFastMob(x, y float64, target *structure.Base, hp int, speed, burst float64) *Mob {
	m := NewMob(x, y, target, hp, speed*0.3) // Much slower base movement
	m.Burst = burst
	m.BurstTimer = core.NewCooldownTimer(4.0)  // Longer cooldown between bursts
	m.BurstActive = core.NewCooldownTimer(1.0) // Burst lasts 1 second
	m.BurstActive.Remaining = 0                 // Start not in burst
	m.MobType = entity.MobFast
	return m
}

// NewBossMob creates a tough boss enemy.
func NewBossMob(x, y float64, target *structure.Base, hp int, speed float64) *Mob {
	m := NewMob(x, y, target, hp, speed)
	m.MobType = entity.MobBoss
	return m
}

// Update moves the mob and handles animation.
func (m *Mob) Update(dt float64) error {
	spd := m.Speed * 0.5 // Slow down all mobs significantly

	// Handle burst mechanics for fast mobs
	if m.Burst > 0 {
		if !m.BurstActive.Ready() {
			// Currently in burst phase
			spd = m.Speed * m.Burst * 0.5 // Apply burst but still slower
			m.BurstActive.Tick(dt)
		} else if m.BurstTimer.Tick(dt) {
			// Cooldown finished, start new burst
			m.BurstActive.Reset()
		}
	}

	// Calculate velocity based on target direction and current speed
	if m.Target != nil {
		dx := m.Target.Position.X - m.Position.X
		dy := m.Target.Position.Y - m.Position.Y
		dist := math.Hypot(dx, dy)
		if dist != 0 {
			m.VX = dx / dist * spd
			m.VY = dy / dist * spd
		}
	}

	// Update position
	m.Position.X += m.VX * dt
	m.Position.Y += m.VY * dt

	// Handle animation (slower)
	m.AnimTicker += dt
	if int(m.AnimTicker/0.5)%2 == 0 { // Slower animation
		m.Frame = assets.ImgMobA
	} else {
		m.Frame = assets.ImgMobB
	}
	return nil
}

// Velocity returns the mob's current velocity components.
func (m *Mob) Velocity() (vx, vy float64) {
	return m.VX, m.VY
}

// Damage applies damage considering armor and shields.
func (m *Mob) Damage(d int) {
	if !m.Alive {
		return
	}
	if m.Shield > 0 {
		m.Shield -= d
		if m.Shield < 0 {
			d = -m.Shield
			m.Shield = 0
		} else {
			return
		}
	}
	if m.Armor > 0 {
		d -= m.Armor
		if d < 0 {
			d = 0
		}
	}
	m.Health -= d
	if m.Health <= 0 {
		m.Alive = false
	}
}

// Type returns the mob type.
func (m *Mob) Type() entity.MobType { return m.MobType }
