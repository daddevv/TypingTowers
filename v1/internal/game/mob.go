package game

import (
	"github.com/daddevv/type-defense/internal/entity"
	"math"
)

// Mob represents a basic enemy moving left.
type Mob struct {
	entity.BaseEntity
	speed      float64
	animTicker float64
	alive      bool
	vx, vy     float64
	target     *Base
	health     int

	armor       int
	shield      int
	burst       float64
	burstTimer  CooldownTimer // Use proper timer for burst cooldown
	burstActive CooldownTimer // Use proper timer for burst duration
	mobType     MobType
}

// NewMob returns a new mob at the given position.
func NewMob(x, y float64, target *Base, hp int, speed float64) *Mob {
	w, h := ImgMobA.Bounds().Dx(), ImgMobA.Bounds().Dy()
	return &Mob{
		BaseEntity: entity.BaseEntity{
			pos:          entity.Point{x, y},
			width:        w,
			height:       h,
			frame:        ImgMobA,
			frameAnchorX: float64(w) / 2,
			frameAnchorY: float64(h) / 2,
		},
		speed:   speed,
		alive:   true,
		target:  target,
		health:  hp,
		mobType: MobBasic,
	}
}

// NewArmoredMob creates a mob with armor reducing incoming damage.
func NewArmoredMob(x, y float64, target *Base, hp, armor int, speed float64) *Mob {
	m := NewMob(x, y, target, hp, speed)
	m.armor = armor
	m.mobType = MobArmored
	return m
}

// NewFastMob creates a mob with periodic speed bursts.
func NewFastMob(x, y float64, target *Base, hp int, speed, burst float64) *Mob {
	m := NewMob(x, y, target, hp, speed*0.3) // Much slower base movement
	m.burst = burst
	m.burstTimer = NewCooldownTimer(4.0)  // Longer cooldown between bursts
	m.burstActive = NewCooldownTimer(1.0) // Burst lasts 1 second
	m.burstActive.remaining = 0           // Start not in burst
	m.mobType = MobFast
	return m
}

// NewBossMob creates a tough boss enemy.
func NewBossMob(x, y float64, target *Base, hp int, speed float64) *Mob {
	m := NewMob(x, y, target, hp, speed)
	m.mobType = MobBoss
	return m
}

// Update moves the mob and handles animation.
func (m *Mob) Update(dt float64) error {
	spd := m.speed * 0.5 // Slow down all mobs significantly

	// Handle burst mechanics for fast mobs
	if m.burst > 0 {
		if !m.burstActive.Ready() {
			// Currently in burst phase
			spd = m.speed * m.burst * 0.5 // Apply burst but still slower
			m.burstActive.Tick(dt)
		} else if m.burstTimer.Tick(dt) {
			// Cooldown finished, start new burst
			m.burstActive.Reset()
		}
	}

	// Calculate velocity based on target direction and current speed
	if m.target != nil {
		dx := m.target.pos.X - m.pos.X
		dy := m.target.pos.Y - m.pos.Y
		dist := math.Hypot(dx, dy)
		if dist != 0 {
			m.vx = dx / dist * spd
			m.vy = dy / dist * spd
		}
	}

	// Update position
	m.pos.X += m.vx * dt
	m.pos.Y += m.vy * dt

	// Handle animation (slower)
	m.animTicker += dt
	if int(m.animTicker/0.5)%2 == 0 { // Slower animation
		m.frame = ImgMobA
	} else {
		m.frame = ImgMobB
	}
	return nil
}

// Velocity returns the mob's current velocity components.
func (m *Mob) Velocity() (vx, vy float64) {
	return m.vx, m.vy
}

// Alive reports whether the mob is still active.
func (m *Mob) Alive() bool { return m.alive }

// Damage applies damage considering armor and shields.
func (m *Mob) Damage(d int) {
	if !m.alive {
		return
	}
	if m.shield > 0 {
		m.shield -= d
		if m.shield < 0 {
			d = -m.shield
			m.shield = 0
		} else {
			return
		}
	}
	if m.armor > 0 {
		d -= m.armor
		if d < 0 {
			d = 0
		}
	}
	m.health -= d
	if m.health <= 0 {
		m.alive = false
	}
}

// Type returns the mob type.
func (m *Mob) Type() MobType { return m.mobType }
