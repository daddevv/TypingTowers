package enemy

import (
	"td/internal/math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Mob struct {
	Name             string        // Name of the mob
	Sprite           *ebiten.Image // Sprite image for the mob
	Health           int           // Health points of the mob
	Target           *math.Vec2    // Target position for the mob to move towards
	Position         *math.Vec2    // Position of the mob in the game world
	Velocity         *math.Vec2    // Velocity vector of the mob
	Acceleration     *math.Vec2    // Acceleration vector for the mob
	AccelerationRate float64       // Rate of acceleration for the mob
	ActionRange      float64       // Range within which the mob can attack
	AttackDamage     int           // Damage dealt by the mob when it attacks
	AttackCooldown   int           // Cooldown time between attacks
	AttackTick       int           // Current tick for attack cooldown
}

func NewMob(name string, health, posX, posY, velX, velY, targetX, targetY int) *Mob {
	sprite, _, err := ebitenutil.NewImageFromFile("assets/images/mob/mob_beach_ball.png")
	if err != nil {
		panic(err) // Handle error appropriately in production code
	}

	return &Mob{
		Name:             name,
		Sprite:           sprite,
		Health:           health,
		Target:           math.NewVec2(float64(targetX), float64(targetY)), // Initialize target position to current position
		Position:         math.NewVec2(float64(posX), float64(posY)),
		Velocity:         math.NewVec2(float64(velX), float64(velY)),
		Acceleration:     math.NewVec2(0.0, 0.0), // Default X acceleration
		AccelerationRate: 0.1,                    // Default acceleration rate
		ActionRange:      100.0,                  // Default action range
		AttackDamage:     10,                     // Default attack damage
		AttackCooldown:   60,                     // Default attack cooldown in ticks
		AttackTick:       0,                      // Initialize attack tick
	}
}

func (m *Mob) Update() {
	m.Position.X += m.Velocity.X
	m.Position.Y += m.Velocity.Y

	// Update the mob's velocity based on acceleration
	m.Velocity.X += m.Acceleration.X
	m.Velocity.Y += m.Acceleration.Y

	// Clamp the mob's velocity to a maximum value
	maxSpeed := 5.0
	if m.Velocity.Magnitude() > maxSpeed {
		m.Velocity = m.Velocity.Normalize().Scale(maxSpeed)
	}

	// Update the mob's attack cooldown
	if m.AttackTick > 0 {
		m.AttackTick--
	}
}

func (m *Mob) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(m.Position.X, m.Position.Y) // Position the mob at its current position
	screen.DrawImage(m.Sprite, opts)                // Draw the mob sprite
}
