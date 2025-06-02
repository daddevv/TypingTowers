package game

import (
	"image/color"
	"td/internal/math"

	"slices"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct {
	Mobs []*Mob // Slice to hold active mobs
}

// Mob represents an autonomous entity in the game, such as an enemy or NPC.
type Mob struct {
	Name          string  // Name of the mob
	Team Team
	Sprite        *ebiten.Image // Sprite image for the mob
	Health        int     // Health points of the mob
	Target        *math.Vec2   // Target position for the mob to move towards
	Position      *math.Vec2   // Position of the mob in the game world
	Velocity      *math.Vec2   // Velocity vector of the mob
	Acceleration  *math.Vec2   // Acceleration vector for the mob
	AccelerationRate float64 // Rate of acceleration for the mob
	ActionRange   float64 // Range within which the mob can attack
	AttackDamage  int     // Damage dealt by the mob when it attacks
	AttackCooldown int    // Cooldown time between attacks
	AttackTick    int     // Current tick for attack cooldown
}

// NewMob creates a new Mob instance with the specified parameters.
func NewMob(name string, team Team, health, posX, posY, velX, velY int) *Mob {
	sprite := ebiten.NewImage(64, 64) // Create a new image for the mob sprite
	switch team {
	case TeamPlayer:
		vector.DrawFilledCircle(sprite, 32, 32, 32, color.RGBA{R: 0, G: 0, B: 255, A: 255}, false)
	case TeamEnemy:
		vector.DrawFilledCircle(sprite, 32, 32, 32, color.RGBA{R: 255, G: 0, B: 0, A: 255}, false)
	default:
		vector.DrawFilledCircle(sprite, 32, 32, 32, color.RGBA{R: 128, G: 128, B: 128, A: 255}, false)
	}
	return &Mob{
		Name:          name,
		Team:          team,
		Sprite:        sprite,
		Health:        health,
		Target:        math.NewVec2(float64(posX), float64(posY)), // Initialize target position to current position
		Position:      math.NewVec2(float64(posX), float64(posY)),
		Velocity:      math.NewVec2(float64(velX), float64(velY)),
		Acceleration:  math.NewVec2(0.0, 0.0), // Default X acceleration
		AccelerationRate: 0.1,                // Default acceleration rate
		ActionRange:   20.0,                 // Default action range
		AttackDamage:  10,                   // Default attack damage
		AttackCooldown: 60,                  // Default attack cooldown in ticks
		AttackTick:    0,                    // Initialize attack tick
	}
}

// Update updates the mob's position and state based on its target.
func (m *Mob) Update(target *math.Vec2) {
	mobFeet := m.Position.Add(math.NewVec2(float64(m.Sprite.Bounds().Dx())/2, float64(m.Sprite.Bounds().Dy()))) // Adjust mob position to feet level
	// Update the mob's target position
	m.Target = target
	// Calculate the direction vector towards the target
	direction := m.Target.Subtract(mobFeet)
	// Normalize the direction vector to get the unit vector
	directionMagnitude := direction.Magnitude()
	if directionMagnitude > 0 {
		direction = direction.Normalize() // Normalize the direction vector
	}
	// Update the mob's acceleration towards the target
	m.Acceleration = direction.Scale(m.AccelerationRate) // Scale the direction by the acceleration rate
	// Update the mob's velocity based on acceleration
	if directionMagnitude < m.ActionRange {
		// If within action range, set velocity to zero
		m.Velocity = math.NewVec2(0, 0)
	} else {
		// Otherwise, move towards the target
		m.Velocity = m.Velocity.Add(m.Acceleration) // Update velocity with acceleration
	}
	// Update the mob's position based on its velocity
	m.Position = m.Position.Add(m.Velocity) // Move the mob towards the target position
	// Ensure the mob's position does not exceed the game boundaries
	if m.Position.X < 0 {
		m.Position.X = 0 // Prevent moving out of bounds on the left
	}
	if m.Position.Y < 0 {
		m.Position.Y = 0 // Prevent moving out of bounds on the top
	}
	if m.Position.X > 1920-float64(m.Sprite.Bounds().Dx()) {
		m.Position.X = 1920 - float64(m.Sprite.Bounds().Dx()) // Prevent moving out of bounds on the right
	}
	if m.Position.Y > 1080-float64(m.Sprite.Bounds().Dy()) {
		m.Position.Y = 1080 - float64(m.Sprite.Bounds().Dy()) // Prevent moving out of bounds on the bottom
	}
}

// LocateActionableTarget finds the nearest enemy mob within the game.
func (m *Mob) LocateActionableTarget(mobs []*Mob) *Mob {
	nearest := m.TargetNearestEnemy(mobs) // Find the nearest enemy mob
	if nearest != nil && m.Position.Distance(nearest.Position) <= m.ActionRange {
		return nearest // Return the nearest enemy if within action range
	}
	return nil // Return nil if no enemy is found within action range
}

// TargetNearestEnemy sets the mob's target to the nearest enemy mob.
func (m *Mob) TargetNearestEnemy(mobs []*Mob) *Mob {
	var nearest *Mob
	minDistance := float64(1<<63 - 1) // Initialize to a very large number
	for _, mob := range mobs {
		if mob != nil && mob.Team != m.Team{ // Only consider mobs that are not on the same team
			distance := m.Position.Distance(mob.Position) // Calculate distance to the mob
			if distance < minDistance {
				minDistance = distance // Update the minimum distance
				nearest = mob // Update the nearest
			}
		}
	}
	return nearest // Return the nearest enemy mob
}

// Team represents the team affiliation of a mob.
type Team int
const (
	TeamPlayer Team = iota // Player team
	TeamEnemy               // Enemy team
)

// NewGame creates a new Game instance.
func NewGame() *Game {
	return &Game{
		// Initialize game state here
		Mobs: make([]*Mob, 0), // Initialize the slice of mobs
	}
}

// Update updates the game state for a single frame.
func (g *Game) Update() error {
	// Update game logic here, such as player movement, enemy AI, etc.
	for _, mob := range g.Mobs {
		if mob != nil {
			// Update each mob's state, e.g., position, health, etc.
			target := mob.TargetNearestEnemy(g.Mobs) // Find the nearest enemy for the mob
			if target != nil {
				mob.Update(target.Position) // Update the mob's position towards the target
			}
			// Handle mob actions, such as attacking or moving
			if mob.AttackTick > 0 {
				mob.AttackTick-- // Decrease the attack cooldown tick
			}
			if mob.Health <= 0 {
				g.RemoveMob(mob) // Remove the mob if its health is zero or below
			}
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return ebiten.Termination // Exit the game if Escape is pressed
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyP) {
		// Handle P key press, e.g., pause the game or perform an action
		ally := NewMob("New Mob", TeamPlayer, 100, 100, 100, 0, 0) // Create a new mob
		g.AddMob(ally) // Add the new mob to the game
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyE) {
		enemy := NewMob("Enemy Mob", TeamEnemy, 100, 900, 900, 0, 0) // Create a new enemy mob
		g.AddMob(enemy) // Add the new enemy mob to the game
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyC) {
		g.ClearMobs() // Clear all mobs from the game
	}

	return nil
}

// Draw draws the game state to the screen.
func (g *Game) Draw(screen *ebiten.Image) {
	// Render the game state to the screen here
	// This could include drawing the player, enemies, background, etc.
	for _, mob := range g.Mobs {
		if mob != nil {
			opts := &ebiten.DrawImageOptions{}
			opts.GeoM.Translate(mob.Position.X, mob.Position.Y) // Set the position of the mob
			screen.DrawImage(mob.Sprite, opts) // Draw the mob's sprite at its position
		}
	}
}

func (g *Game) AddMob(mob *Mob) {
	g.Mobs = append(g.Mobs, mob) // Add a new mob to the game
}

func (g *Game) RemoveMob(mob *Mob) {
	for i, m := range g.Mobs {
		if m == mob {
			g.Mobs = slices.Delete(g.Mobs, i, i+1) // Remove the mob from the game
			break
		}
	}
}

func (g *Game) GetMobs() []*Mob {
	return g.Mobs // Return the list of mobs in the game
}

func (g *Game) ClearMobs() {
	g.Mobs = make([]*Mob, 0) // Clear the list of mobs
}

func (g *Game) GetMobByName(name string) *Mob {
	for _, mob := range g.Mobs {
		if mob.Name == name {
			return mob // Return the first mob with the matching name
		}
	}
	return nil // Return nil if no mob with the given name is found
}

func (g *Game) GetMobAtPosition(pos *math.Vec2) *Mob {
	for _, mob := range g.Mobs {
		if mob.Position.Equals(pos) {
			return mob // Return the first mob at the specified position
		}
	}
	return nil // Return nil if no mob is found at the position
}
