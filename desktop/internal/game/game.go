package game

import (
	"image"
	"image/color"
	maths "math"
	"math/rand"
	"slices"
	"td/internal/math"
	"time"
	"unsafe"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct {
	Mobs []*Mob // Slice to hold active mobs
}

// Mob represents an autonomous entity in the game, such as an enemy or NPC.
type Mob struct {
	Name             string
	Team              Team
	Sprite           *ebiten.Image
	Health           int
	Target           *math.Vec2
	Position         *math.Vec2
	Velocity         *math.Vec2
	Acceleration     *math.Vec2
	AccelerationRate float64
	ActionRange      float64
	AttackDamage     int
	AttackCooldown   int
	AttackTick       int
	HitBox           image.Rectangle // Add hitbox for collision
	IdleTimer      int           // Timer for changing idle direction
	IdleDuration   int           // How long to keep current idle direction
	IdleDirection  *math.Vec2   // Current idle direction
}

// NewMob creates a new Mob instance with the specified parameters.
func NewMob(name string, team Team, health, posX, posY, velX, velY int) *Mob {
	sprite := ebiten.NewImage(64, 64)
	switch team {
	case TeamPlayer:
		vector.DrawFilledCircle(sprite, 32, 32, 32, color.RGBA{R: 0, G: 0, B: 255, A: 255}, false)
	case TeamEnemy:
		vector.DrawFilledCircle(sprite, 32, 32, 32, color.RGBA{R: 255, G: 0, B: 0, A: 255}, false)
	default:
		vector.DrawFilledCircle(sprite, 32, 32, 32, color.RGBA{R: 128, G: 128, B: 128, A: 255}, false)
	}
	hitBox := image.Rect(0, 0, 48, 48) // Smaller than sprite for better collision
	return &Mob{
		Name:             name,
		Team:             team,
		Sprite:           sprite,
		Health:           health,
		Target:           math.NewVec2(float64(posX), float64(posY)), // Initialize target position to current position
		Position:         math.NewVec2(float64(posX), float64(posY)),
		Velocity:         math.NewVec2(float64(velX), float64(velY)),
		Acceleration:     math.NewVec2(0.0, 0.0), // Default X acceleration
		AccelerationRate: 0.2,                // Default acceleration rate
		ActionRange:      48.0,                 // Default action range
		AttackDamage:     10,                   // Default attack damage
		AttackCooldown:   60,                  // Default attack cooldown in ticks
		AttackTick:       0,                    // Initialize attack tick
		HitBox:           hitBox,
		IdleTimer:    0,
		IdleDuration: 0,
		IdleDirection: math.NewVec2(0, 0),
	}
}

// Boid parameters
const (
	boidSeparationDist = 60.0
	boidCohesionDist   = 120.0
	boidAlignmentDist  = 120.0
	boidSeparationWeight = 1.2
	boidCohesionWeight   = 0.5
	boidAlignmentWeight  = 0.5
	boidMaxSpeed         = 3.0
	boidMaxForce         = 0.2
)

// // Update updates the mob's position and state based on its target.
// func (m *Mob) Update(target *math.Vec2) {
// 	mobFeet := m.Position.Add(math.NewVec2(float64(m.Sprite.Bounds().Dx())/2, float64(m.Sprite.Bounds().Dy()))) // Adjust mob position to feet level
// 	// Update the mob's target position
// 	m.Target = target
// 	// Calculate the direction vector towards the target
// 	direction := m.Target.Subtract(mobFeet)
// 	// Normalize the direction vector to get the unit vector
// 	directionMagnitude := direction.Magnitude()
// 	if directionMagnitude > 0 {
// 		direction = direction.Normalize() // Normalize the direction vector
// 	}
// 	// Update the mob's acceleration towards the target
// 	m.Acceleration = direction.Scale(m.AccelerationRate) // Scale the direction by the acceleration rate
// 	// Update the mob's velocity based on acceleration
// 	if directionMagnitude < m.ActionRange {
// 		// If within action range, set velocity to zero
// 		m.Velocity = math.NewVec2(0, 0)
// 	} else {
// 		// Otherwise, move towards the target
// 		m.Velocity = m.Velocity.Add(m.Acceleration) // Update velocity with acceleration
// 	}
// 	// Update the mob's position based on its velocity
// 	m.Position = m.Position.Add(m.Velocity) // Move the mob towards the target position
// 	// Ensure the mob's position does not exceed the game boundaries
// 	if m.Position.X < 0 {
// 		m.Position.X = 0 // Prevent moving out of bounds on the left
// 	}
// 	if m.Position.Y < 0 {
// 		m.Position.Y = 0 // Prevent moving out of bounds on the top
// 	}
// 	if m.Position.X > 1920-float64(m.Sprite.Bounds().Dx()) {
// 		m.Position.X = 1920 - float64(m.Sprite.Bounds().Dx()) // Prevent moving out of bounds on the right
// 	}
// 	if m.Position.Y > 1080-float64(m.Sprite.Bounds().Dy()) {
// 		m.Position.Y = 1080 - float64(m.Sprite.Bounds().Dy()) // Prevent moving out of bounds on the bottom
// 	}
// }

// LocateActionableTarget finds the nearest enemy mob within the game.
func (m *Mob) LocateActionableTarget(mobs []*Mob) *Mob {
	nearest := m.TargetNearestEnemy(mobs) // Find the nearest enemy mob
	if nearest != nil && m.Position.Distance(nearest.Position) <= m.ActionRange {
		return nearest // Return the nearest enemy if within action range
	}
	return nil // Return nil if no enemy is found within action range
}

// UpdateBoid updates the mob's state using boid behavior (separation, alignment, cohesion).
func (m *Mob) UpdateBoid(mobs []*Mob) {
	// Boid: separation, alignment, cohesion (only with same team)
	var (
		separation = math.NewVec2(0, 0)
		alignment  = math.NewVec2(0, 0)
		cohesion   = math.NewVec2(0, 0)
		countSep, countAli, countCoh int
	)
	for _, other := range mobs {
		if other == nil || other == m || other.Team != m.Team {
			continue
		}
		dist := m.Position.Distance(other.Position)
		if dist < boidSeparationDist && dist > 0 {
			diff := m.Position.Subtract(other.Position).Normalize().Scale(1.0 / dist)
			separation = separation.Add(diff)
			countSep++
		}
		if dist < boidAlignmentDist {
			alignment = alignment.Add(other.Velocity)
			countAli++
		}
		if dist < boidCohesionDist {
			cohesion = cohesion.Add(other.Position)
			countCoh++
		}
	}
	if countSep > 0 {
		separation = separation.Scale(1.0 / float64(countSep))
	}
	if countAli > 0 {
		alignment = alignment.Scale(1.0 / float64(countAli))
		alignment = alignment.Normalize().Scale(boidMaxSpeed)
		alignment = alignment.Subtract(m.Velocity)
	}
	if countCoh > 0 {
		cohesion = cohesion.Scale(1.0 / float64(countCoh))
		cohesion = cohesion.Subtract(m.Position)
		cohesion = cohesion.Normalize().Scale(boidMaxSpeed)
		cohesion = cohesion.Subtract(m.Velocity)
	}
	// Combine
	boidForce := separation.Scale(boidSeparationWeight).
		Add(alignment.Scale(boidAlignmentWeight)).
		Add(cohesion.Scale(boidCohesionWeight))
	// Limit force
	if boidForce.Magnitude() > boidMaxForce {
		boidForce = boidForce.Normalize().Scale(boidMaxForce)
	}
	m.Acceleration = m.Acceleration.Add(boidForce)
}

// Update updates the mob's position and state based on its target.
func (m *Mob) Update(mobs []*Mob) {
	// Boid behavior (per team)
	m.UpdateBoid(mobs)

	// Prevent stacking: add repulsion if hitboxes overlap (per team)
	for _, other := range mobs {
		if other == nil || other == m || other.Team != m.Team {
			continue
		}
		if m.HitBox.Overlaps(other.HitBox) {
			// Push away from each other
			offset := m.Position.Subtract(other.Position)
			if offset.Magnitude() == 0 {
				offset = math.NewVec2(rand.Float64()-0.5, rand.Float64()-0.5)
			}
			repulse := offset.Normalize().Scale(0.5)
			m.Acceleration = m.Acceleration.Add(repulse)
		}
	}

	// Seek nearest enemy only if one is actionable (within a reasonable distance)
	targetMob := m.TargetNearestEnemy(mobs)
	hasEnemy := targetMob != nil

	if hasEnemy {
		m.Target = targetMob.Position
		dist := m.Position.Distance(targetMob.Position)
		if dist < m.ActionRange {
			// In action range: stationary while battling
			m.Acceleration = math.NewVec2(0, 0)
			m.Velocity = math.NewVec2(0, 0)
			// Reset idle state when enemy found
			m.IdleTimer = 0
			m.IdleDuration = 0
			m.IdleDirection = math.NewVec2(0, 0)
		} else {
			// Move toward target
			dir := m.Target.Subtract(m.Position)
			if dir.Magnitude() > 0 {
				dir = dir.Normalize()
			}
			seekForce := dir.Scale(m.AccelerationRate)
			m.Acceleration = m.Acceleration.Add(seekForce)
			// Reset idle state when enemy found
			m.IdleTimer = 0
			m.IdleDuration = 0
			m.IdleDirection = math.NewVec2(0, 0)
		}
	} else {
		// Idle wandering: change direction every N ticks
		if m.IdleTimer <= 0 {
			rand.Seed(time.Now().UnixNano() + int64(uintptr(unsafe.Pointer(m))))
			angle := rand.Float64() * 2 * 3.1415926535
			mag := 0.2 + rand.Float64()*0.5 // random acceleration magnitude
			m.IdleDirection = math.NewVec2(maths.Cos(angle), maths.Sin(angle)).Scale(mag)
			m.IdleDuration = 60 + rand.Intn(120) // 1-3 seconds at 60fps
			m.IdleTimer = m.IdleDuration
		}
		m.Acceleration = m.Acceleration.Add(m.IdleDirection)
		m.IdleTimer--
	}

	// Update velocity and clamp
	m.Velocity = m.Velocity.Add(m.Acceleration)
	if m.Velocity.Magnitude() > boidMaxSpeed {
		m.Velocity = m.Velocity.Normalize().Scale(boidMaxSpeed)
	}
	m.Position = m.Position.Add(m.Velocity)
	m.Acceleration = math.NewVec2(0, 0)

	// Update hitbox position
	m.HitBox = image.Rect(
		int(m.Position.X)+8, int(m.Position.Y)+8,
		int(m.Position.X)+8+32, int(m.Position.Y)+8+32,
	)

	// Attack if in range and there is an enemy
	if hasEnemy && m.Position.Distance(targetMob.Position) < m.ActionRange {
		if m.AttackTick == 0 {
			targetMob.Health -= m.AttackDamage
			m.AttackTick = m.AttackCooldown
		}
	}
	if m.AttackTick > 0 {
		m.AttackTick--
	}
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
		if mob != nil && mob.Health > 0 {
			mob.Update(g.Mobs)
		}
	}
	// Remove dead mobs
	for i := len(g.Mobs) - 1; i >= 0; i-- {
		if g.Mobs[i] != nil && g.Mobs[i].Health <= 0 {
			g.RemoveMob(g.Mobs[i])
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
			// Draw hitbox for debug
			// Uncomment for debug:
			// hitbox := mob.HitBox
			// vector.StrokeRect(screen, float32(hitbox.Min.X), float32(hitbox.Min.Y), float32(hitbox.Dx()), float32(hitbox.Dy()), 2, color.RGBA{0,255,0,255}, false)
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

func (m *Mob) TargetNearestEnemy(mobs []*Mob) *Mob {
	var nearest *Mob
	minDistance := 1e9
	for _, mob := range mobs {
		if mob != nil && mob.Team != m.Team && mob.Health > 0 {
			distance := m.Position.Distance(mob.Position)
			if distance < minDistance {
				minDistance = distance
				nearest = mob
			}
		}
	}
	return nearest
}
