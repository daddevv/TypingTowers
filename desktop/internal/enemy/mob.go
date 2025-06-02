package enemy

import (
	"td/internal/physics"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Mob struct {
	Name             string        // Name of the mob
	Sprite           *ebiten.Image // Sprite image for the mob
	Health           int           // Health points of the mob
	Target           *physics.Vec2    // Target position for the mob to move towards
	Position         *physics.Vec2    // Position of the mob in the game world
	Velocity         *physics.Vec2    // Velocity vector of the mob
	Acceleration     *physics.Vec2    // Acceleration vector for the mob
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
		Target:           physics.NewVec2(float64(targetX), float64(targetY)), // Initialize target position to current position
		Position:         physics.NewVec2(float64(posX), float64(posY)),
		Velocity:         physics.NewVec2(float64(velX), float64(velY)),
		Acceleration:     physics.NewVec2(0.0, 0.0), // Default X acceleration
		AccelerationRate: 0.5,                    // Default acceleration rate
		ActionRange:      10.0,                  // Default action range
		AttackDamage:     10,                     // Default attack damage
		AttackCooldown:   60,                     // Default attack cooldown in ticks
		AttackTick:       0,                      // Initialize attack tick
	}
}

func (m *Mob) Update(target *physics.Vec2) {
	mobFeet := m.Position.Add(physics.NewVec2(float64(m.Sprite.Bounds().Dx())/2, float64(m.Sprite.Bounds().Dy()))) // Adjust mob position to feet level

	// Update the mob's target position
	m.Target = target
	// Calculate the direction vector towards the target
	direction := m.Target.Subtract(mobFeet)
	// Normalize the direction vector to get the unit vector
	directionMagnitude := direction.Magnitude()
	if directionMagnitude > 0 {
		direction = direction.Normalize()
	}
	// Update the mob's acceleration towards the target
	m.Acceleration = direction.Scale(m.AccelerationRate)
	// Update the mob's position based on its velocity
	// Move the mob towards the target position
	if directionMagnitude < m.ActionRange {
		// If within action range, set velocity to zero
		m.Velocity = physics.NewVec2(0, 0)
	} else {
		// Otherwise, move towards the target
		m.Velocity = direction.Scale(m.AccelerationRate)
	}
	// Update the mob's position based on its velocity
	m.Position.X += m.Velocity.X
	m.Position.Y += m.Velocity.Y

	// Update the mob's velocity based on acceleration
	m.Velocity.X += m.Acceleration.X
	m.Velocity.Y += m.Acceleration.Y

	// Clamp the mob's velocity to a maximum value
	maxSpeed := 10.0
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
