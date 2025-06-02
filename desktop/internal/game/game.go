package game

import (
	"fmt"
	"image"
	"image/color"
	"math"
	"math/rand"
	"slices"
	"td/internal/physics"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct {
	Mobs []*Mob // Slice to hold active mobs
	SpawnTicks int // Ticks for spawning new mobs
	SpawnInterval int // Interval between spawns
	AllySpawnPoints []*physics.Vec2 // Points where mobs can spawn
	EnemySpawnPoints []*physics.Vec2 // Points where enemy mobs can spawn
}

// Mob represents an autonomous entity in the game, such as an enemy or NPC.
type Mob struct {
	Name             string
	Team              Team
	Sprite           *ebiten.Image
	Health           int
	Target           *physics.Vec2
	Position         *physics.Vec2
	Velocity         *physics.Vec2
	Acceleration     *physics.Vec2
	AccelerationRate float64
	ActionRange      float64
	AttackDamage     int
	AttackCooldown   int
	AttackTick       int
	HitBox           image.Rectangle // Add hitbox for collision
	IdleTimer      int           // Timer for changing idle direction
	IdleDuration   int           // How long to keep current idle direction
	IdleDirection  *physics.Vec2 // Current idle direction
}

// NewMob creates a new Mob instance with the specified parameters.
func NewMob(name string, team Team, health, posX, posY, velX, velY int) *Mob {
	sprite := ebiten.NewImage(16, 16) // Create a new image for the mob sprite
	switch team {
	case TeamPlayer:
		vector.StrokeCircle(sprite, 8, 8, 8, 1, color.RGBA{R: 0, G: 0, B: 255, A: 255}, false) // Blue circle for player team
	case TeamEnemy:
		vector.StrokeCircle(sprite, 8, 8, 8, 1, color.RGBA{R: 255, G: 0, B: 0, A: 255}, false) // Red circle for enemy team
	default:
		vector.StrokeCircle(sprite, 8, 8, 8, 1, color.RGBA{R: 128, G: 128, B: 128, A: 255}, false)
	}
	hitBox := image.Rect(2, 2, 14, 14) // Smaller than sprite for better collision
	return &Mob{
		Name:             name,
		Team:             team,
		Sprite:           sprite,
		Health:           health,
		Target:           physics.NewVec2(float64(posX), float64(posY)), // Initialize target position to current position
		Position:         physics.NewVec2(float64(posX), float64(posY)),
		Velocity:         physics.NewVec2(float64(velX), float64(velY)),
		Acceleration:     physics.NewVec2(0.0, 0.0), // Default X acceleration
		AccelerationRate: 0.4,                // Increased for more aggressive movement
		ActionRange:      26.0,               // Reduced to ~10 pixels (16px sprite + 10px = 26px)
		AttackDamage:     15,                 // Increased damage
		AttackCooldown:   20,                 // Reduced cooldown for faster attacks
		AttackTick:       0,                  // Initialize attack tick
		HitBox:           hitBox,
		IdleTimer:    0,
		IdleDuration: 0,
		IdleDirection: physics.NewVec2(0, 0),
	}
}

// BoidSettings holds the parameters for boid behavior and combat.
type BoidSettings struct {
	SeparationDist   float64
	CohesionDist     float64
	AlignmentDist    float64
	SeparationWeight float64
	CohesionWeight   float64
	AlignmentWeight  float64
	MaxSpeed         float64
	MaxForce         float64
	NeighborRadius   float64

	EnemyFrontCohesionWeight float64
	EnemyFrontCohesionDist   float64

	CombatStoppingDist   float64
	EnemyRepulsionDist   float64
	CombatRepulsionForce float64
}

var (
	BoidParams = &BoidSettings{
		SeparationDist:   40.0,
		CohesionDist:     200.0,
		AlignmentDist:    120.0,
		SeparationWeight: 0.9,
		CohesionWeight:   0.8,
		AlignmentWeight:  1.5,
		MaxSpeed:         1.5,
		MaxForce:         0.5,
		NeighborRadius:   250.0,

		EnemyFrontCohesionWeight: 1.9,
		EnemyFrontCohesionDist:   350.0,

		CombatStoppingDist:   35.0,
		EnemyRepulsionDist:   20.0,
		CombatRepulsionForce: 1.0,
	}
	boidParamNames = []string{
		"SeparationDist", "CohesionDist", "AlignmentDist",
		"SeparationWeight", "CohesionWeight", "AlignmentWeight",
		"MaxSpeed", "MaxForce", "NeighborRadius",
		"EnemyFrontCohesionWeight", "EnemyFrontCohesionDist",
		"CombatStoppingDist", "EnemyRepulsionDist", "CombatRepulsionForce",
	}
	boidParamPtrs = []*float64{
		&BoidParams.SeparationDist, &BoidParams.CohesionDist, &BoidParams.AlignmentDist,
		&BoidParams.SeparationWeight, &BoidParams.CohesionWeight, &BoidParams.AlignmentWeight,
		&BoidParams.MaxSpeed, &BoidParams.MaxForce, &BoidParams.NeighborRadius,
		&BoidParams.EnemyFrontCohesionWeight, &BoidParams.EnemyFrontCohesionDist,
		&BoidParams.CombatStoppingDist, &BoidParams.EnemyRepulsionDist, &BoidParams.CombatRepulsionForce,
	}
	selectedBoidParam = 0
)

// LocateActionableTarget finds the nearest enemy mob within the game.
func (m *Mob) LocateActionableTarget(mobs []*Mob) *Mob {
	nearest := m.TargetNearestEnemy(mobs) // Find the nearest enemy mob
	if nearest != nil && m.Position.Distance(nearest.Position) <= m.ActionRange {
		return nearest // Return the nearest enemy if within action range
	}
	return nil // Return nil if no enemy is found within action range
}

// limitForce limits the magnitude of a force vector to maxForce
func limitForce(force *physics.Vec2, maxForce float64) *physics.Vec2 {
	if force.Magnitude() > maxForce {
		return force.Normalize().Scale(maxForce)
	}
	return force
}

// UpdateBoid updates the mob's state using improved boid behavior (separation, alignment, cohesion, and enemy-front cohesion).
func (m *Mob) UpdateBoid(mobs []*Mob, inCombat bool) {
	// Reduce boid influence during combat
	var forceMultiplier float64 = 1.0
	if inCombat {
		forceMultiplier = 0.3 // Greatly reduce boid forces during combat
	}

	var (
		separation = physics.NewVec2(0, 0)
		alignment  = physics.NewVec2(0, 0)
		cohesion   = physics.NewVec2(0, 0)
		countSep, countAli, countCoh int
	)

	// New: enemy-front cohesion
	var (
		enemyFrontSum   = physics.NewVec2(0, 0)
		enemyFrontCount int
		closestEnemy    *Mob
		closestDist     = 1e9
	)

	// Calculate boid forces only with same team members
	for _, other := range mobs {
		if other == nil || other == m || other.Health <= 0 {
			continue
		}

		dist := m.Position.Distance(other.Position)

		// Only consider nearby boids for flocking
		if other.Team == m.Team {
			if dist > BoidParams.NeighborRadius {
				continue
			}
			// Separation: steer away from nearby boids (increased distance for better spacing)
			if dist < BoidParams.SeparationDist && dist > 0 {
				diff := m.Position.Subtract(other.Position)
				// Weight by distance - closer boids have more influence
				diff = diff.Normalize().Scale(BoidParams.SeparationDist / dist)
				separation = separation.Add(diff)
				countSep++
			}

			// Alignment: steer towards average velocity of neighbors
			if dist < BoidParams.AlignmentDist {
				alignment = alignment.Add(other.Velocity)
				countAli++
			}

			// Cohesion: steer towards average position of neighbors
			if dist < BoidParams.CohesionDist {
				cohesion = cohesion.Add(other.Position)
				countCoh++
			}
		} else {
			// New: enemy-front cohesion
			if dist < BoidParams.EnemyFrontCohesionDist {
				enemyFrontSum = enemyFrontSum.Add(other.Position)
				enemyFrontCount++
			}
			// Track closest enemy for fallback
			if dist < closestDist {
				closestDist = dist
				closestEnemy = other
			}
		}
	}

	// Calculate separation force
	var separationForce *physics.Vec2
	if countSep > 0 {
		separation = separation.Scale(1.0 / float64(countSep))
		if separation.Magnitude() > 0 {
			separationForce = separation.Normalize().Scale(BoidParams.MaxSpeed)
			separationForce = separationForce.Subtract(m.Velocity)
			separationForce = limitForce(separationForce, BoidParams.MaxForce)
		} else {
			separationForce = physics.NewVec2(0, 0)
		}
	} else {
		separationForce = physics.NewVec2(0, 0)
	}

	// Calculate alignment force
	var alignmentForce *physics.Vec2
	if countAli > 0 {
		alignment = alignment.Scale(1.0 / float64(countAli))
		if alignment.Magnitude() > 0 {
			alignmentForce = alignment.Normalize().Scale(BoidParams.MaxSpeed)
			alignmentForce = alignmentForce.Subtract(m.Velocity)
			alignmentForce = limitForce(alignmentForce, BoidParams.MaxForce)
		} else {
			alignmentForce = physics.NewVec2(0, 0)
		}
	} else {
		alignmentForce = physics.NewVec2(0, 0)
	}

	// Calculate cohesion force
	var cohesionForce *physics.Vec2
	if countCoh > 0 {
		cohesion = cohesion.Scale(1.0 / float64(countCoh))
		cohesion = cohesion.Subtract(m.Position)
		if cohesion.Magnitude() > 0 {
			cohesionForce = cohesion.Normalize().Scale(BoidParams.MaxSpeed)
			cohesionForce = cohesionForce.Subtract(m.Velocity)
			cohesionForce = limitForce(cohesionForce, BoidParams.MaxForce)
		} else {
			cohesionForce = physics.NewVec2(0, 0)
		}
	} else {
		cohesionForce = physics.NewVec2(0, 0)
	}

	// New: calculate enemy-front cohesion force
	var enemyFrontCohesionForce *physics.Vec2
	if enemyFrontCount > 0 {
		enemyFrontAvg := enemyFrontSum.Scale(1.0 / float64(enemyFrontCount))
		// Move toward the average position of the enemy front, but only if not in combat
		enemyFrontVec := enemyFrontAvg.Subtract(m.Position)
		if enemyFrontVec.Magnitude() > 0 {
			enemyFrontCohesionForce = enemyFrontVec.Normalize().Scale(BoidParams.MaxSpeed)
			enemyFrontCohesionForce = enemyFrontCohesionForce.Subtract(m.Velocity)
			enemyFrontCohesionForce = limitForce(enemyFrontCohesionForce, BoidParams.MaxForce)
		} else {
			enemyFrontCohesionForce = physics.NewVec2(0, 0)
		}
	} else if closestEnemy != nil {
		// Fallback: move toward closest enemy if no front detected
		enemyFrontVec := closestEnemy.Position.Subtract(m.Position)
		if enemyFrontVec.Magnitude() > 0 {
			enemyFrontCohesionForce = enemyFrontVec.Normalize().Scale(BoidParams.MaxSpeed)
			enemyFrontCohesionForce = enemyFrontCohesionForce.Subtract(m.Velocity)
			enemyFrontCohesionForce = limitForce(enemyFrontCohesionForce, BoidParams.MaxForce)
		} else {
			enemyFrontCohesionForce = physics.NewVec2(0, 0)
		}
	} else {
		enemyFrontCohesionForce = physics.NewVec2(0, 0)
	}

	// Combine all forces with weights and combat multiplier
	boidForce := separationForce.Scale(BoidParams.SeparationWeight * forceMultiplier).
		Add(alignmentForce.Scale(BoidParams.AlignmentWeight * forceMultiplier)).
		Add(cohesionForce.Scale(BoidParams.CohesionWeight * forceMultiplier)).
		Add(enemyFrontCohesionForce.Scale(BoidParams.EnemyFrontCohesionWeight * forceMultiplier))

	// Apply the combined boid force
	m.Acceleration = m.Acceleration.Add(boidForce)
}

// Update updates the mob's position and state based on its target.
func (m *Mob) Update(mobs []*Mob) {
	// Reset acceleration
	m.Acceleration = physics.NewVec2(0, 0)

	// Find nearest enemy and determine combat state
	targetMob := m.TargetNearestEnemy(mobs)
	hasEnemy := targetMob != nil
	inCombat := hasEnemy && m.Position.Distance(targetMob.Position) < BoidParams.CombatStoppingDist

	// Apply boid behavior with reduced influence during combat
	m.UpdateBoid(mobs, inCombat)

	// Strong enemy repulsion to prevent armies pushing through each other
	for _, other := range mobs {
		if other == nil || other == m || other.Team == m.Team || other.Health <= 0 {
			continue
		}

		dist := m.Position.Distance(other.Position)
		if dist < BoidParams.EnemyRepulsionDist && dist > 0 {
			// Strong repulsion from enemies when too close but not attacking
			offset := m.Position.Subtract(other.Position)
			repulse := offset.Normalize().Scale(BoidParams.CombatRepulsionForce * (BoidParams.EnemyRepulsionDist / dist))
			m.Acceleration = m.Acceleration.Add(repulse)
		}
	}

	// Stronger anti-stacking for same team mobs
	for _, other := range mobs {
		if other == nil || other == m || other.Team != m.Team || other.Health <= 0 {
			continue
		}
		if m.HitBox.Overlaps(other.HitBox) {
			// Push away from each other - increased force for better spacing
			offset := m.Position.Subtract(other.Position)
			if offset.Magnitude() == 0 {
				// If positions are identical, create random offset
				offset = physics.NewVec2(rand.Float64()-0.5, rand.Float64()-0.5)
			}
			repulse := offset.Normalize().Scale(0.5) // Increased repulsion for better spacing
			m.Acceleration = m.Acceleration.Add(repulse)
		}
	}

	// Combat behavior - highest priority
	if hasEnemy {
		m.Target = targetMob.Position
		dist := m.Position.Distance(targetMob.Position)

		if dist < m.ActionRange {
			// In attack range: stop completely for precise combat
			m.Velocity = physics.NewVec2(0, 0)
			m.Acceleration = physics.NewVec2(0, 0)
			// Reset idle state when enemy found
			m.IdleTimer = 0
			m.IdleDuration = 0
			m.IdleDirection = physics.NewVec2(0, 0)
		} else if dist < BoidParams.CombatStoppingDist {
			// Close to enemy but not in attack range: slow approach with strong stopping
			dir := m.Target.Subtract(m.Position)
			if dir.Magnitude() > 0 {
				// Very gentle approach when close
				seekForce := dir.Normalize().Scale(BoidParams.MaxSpeed * 0.3)
				seekForce = seekForce.Subtract(m.Velocity)
				seekForce = limitForce(seekForce, BoidParams.MaxForce*0.5)
				m.Acceleration = m.Acceleration.Add(seekForce)
			}
			// Strong damping to prevent overshooting
			m.Velocity = m.Velocity.Scale(0.7)
		} else {
			// Far from enemy: aggressive seeking
			dir := m.Target.Subtract(m.Position)
			if dir.Magnitude() > 0 {
				seekForce := dir.Normalize().Scale(BoidParams.MaxSpeed * 1.2)
				seekForce = seekForce.Subtract(m.Velocity)
				seekForce = limitForce(seekForce, BoidParams.MaxForce*1.5)
				seekForce = seekForce.Scale(m.AccelerationRate)
				m.Acceleration = m.Acceleration.Add(seekForce)
			}
		}
		// Reset idle state when enemy found
		m.IdleTimer = 0
		m.IdleDuration = 0
		m.IdleDirection = physics.NewVec2(0, 0)
	} else {
		// Idle wandering: more restless movement when no enemies
		if m.IdleTimer <= 0 {
			angle := rand.Float64() * 2 * math.Pi
			mag := 0.2 + rand.Float64()*0.4 // More aggressive idle movement
			m.IdleDirection = physics.NewVec2(math.Cos(angle), math.Sin(angle)).Scale(mag)
			m.IdleDuration = 45 + rand.Intn(90) // Shorter duration, more frequent direction changes
			m.IdleTimer = m.IdleDuration
		}

		// Apply idle wandering force
		idleForce := limitForce(m.IdleDirection, BoidParams.MaxForce*0.7) // Stronger idle force
		m.Acceleration = m.Acceleration.Add(idleForce)
		m.IdleTimer--
	}

	// Update velocity with acceleration
	m.Velocity = m.Velocity.Add(m.Acceleration)

	// Limit velocity to max speed
	if m.Velocity.Magnitude() > BoidParams.MaxSpeed {
		m.Velocity = m.Velocity.Normalize().Scale(BoidParams.MaxSpeed)
	}

	// Update position
	m.Position = m.Position.Add(m.Velocity)

	// Boundary collision with bounce-back
	screenMargin := 32.0
	if m.Position.X < screenMargin {
		m.Position.X = screenMargin
		m.Velocity.X = math.Abs(m.Velocity.X) // Bounce off left wall
	}
	if m.Position.Y < screenMargin {
		m.Position.Y = screenMargin
		m.Velocity.Y = math.Abs(m.Velocity.Y) // Bounce off top wall
	}
	if m.Position.X > 1920-float64(m.Sprite.Bounds().Dx())-screenMargin {
		m.Position.X = 1920 - float64(m.Sprite.Bounds().Dx()) - screenMargin
		m.Velocity.X = -math.Abs(m.Velocity.X) // Bounce off right wall
	}
	if m.Position.Y > 1080-float64(m.Sprite.Bounds().Dy())-screenMargin {
		m.Position.Y = 1080 - float64(m.Sprite.Bounds().Dy()) - screenMargin
		m.Velocity.Y = -math.Abs(m.Velocity.Y) // Bounce off bottom wall
	}

	// Update hitbox position
	m.HitBox = image.Rect(
		int(m.Position.X)+2, int(m.Position.Y)+2,
		int(m.Position.X)+14, int(m.Position.Y)+14,
	)

	// Attack behavior with reduced range - mobs need to be very close to attack
	if hasEnemy && m.Position.Distance(targetMob.Position) < m.ActionRange {
		if m.AttackTick == 0 {
			// Attack when very close (about 10 pixels apart)
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
		SpawnTicks: 0,         // No mobs spawned initially
		SpawnInterval: 5,
		AllySpawnPoints: []*physics.Vec2{
			physics.NewVec2(100, 100),   // Example spawn points for allies
			physics.NewVec2(200, 100),
			physics.NewVec2(300, 100),
			physics.NewVec2(400, 100),
			physics.NewVec2(500, 100),
			physics.NewVec2(600, 100),
			physics.NewVec2(700, 100),
			physics.NewVec2(800, 100),
			physics.NewVec2(900, 100),
			physics.NewVec2(1000, 100),
			physics.NewVec2(1100, 100),
			physics.NewVec2(1200, 100),
			physics.NewVec2(1300, 100),
			physics.NewVec2(1400, 100),
			physics.NewVec2(1500, 100),
			physics.NewVec2(1600, 100),
			physics.NewVec2(1700, 100),
			physics.NewVec2(1800, 100),
			physics.NewVec2(1900, 100),
		},
		EnemySpawnPoints: []*physics.Vec2{
			physics.NewVec2(100, 900),   // Example spawn points for enemies
			physics.NewVec2(200, 900),
			physics.NewVec2(300, 900),
			physics.NewVec2(400, 900),
			physics.NewVec2(500, 900),
			physics.NewVec2(600, 900),
			physics.NewVec2(700, 900),
			physics.NewVec2(800, 900),
			physics.NewVec2(900, 900),
			physics.NewVec2(1000, 900),
			physics.NewVec2(1100, 900),
			physics.NewVec2(1200, 900),
			physics.NewVec2(1300, 900),
			physics.NewVec2(1400, 900),
			physics.NewVec2(1500, 900),
			physics.NewVec2(1600, 900),
			physics.NewVec2(1700, 900),
			physics.NewVec2(1800, 900),
			physics.NewVec2(1900, 900),
		},
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

	if inpututil.IsKeyJustPressed(ebiten.KeyC) {
		g.ClearMobs() // Clear all mobs from the game
	}

	// Handle boid parameter adjustment
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
		selectedBoidParam++
		if selectedBoidParam >= len(boidParamNames) {
			selectedBoidParam = 0 // Wrap around to the first parameter
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
		selectedBoidParam--
		if selectedBoidParam < 0 {
			selectedBoidParam = len(boidParamNames) - 1 // Wrap around to the last parameter
		}
	}

	// +/- to adjust selected parameter
	if inpututil.IsKeyJustPressed(ebiten.KeyEqual) || inpututil.IsKeyJustPressed(ebiten.KeyKPAdd) || ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		*boidParamPtrs[selectedBoidParam] += 0.01
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyMinus) || inpututil.IsKeyJustPressed(ebiten.KeyKPSubtract) || ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		*boidParamPtrs[selectedBoidParam] -= 0.01
		if *boidParamPtrs[selectedBoidParam] < 0 {
			*boidParamPtrs[selectedBoidParam] = 0
		}
	}

	// Spawn new mobs periodically
	if g.SpawnTicks <= 0 {
		if len(g.Mobs) >= 100 {
			return nil // Stop spawning if there are already 100 mobs
		}
		enemySpawnPoint := g.EnemySpawnPoints[rand.Intn(len(g.EnemySpawnPoints))] // Randomly select a spawn point
		allySpawnPoint := g.AllySpawnPoints[rand.Intn(len(g.AllySpawnPoints))] // Randomly select a spawn point for allies
		newEnemyMob := NewMob(fmt.Sprintf("Mob%d", len(g.Mobs)+1), TeamEnemy, 100, int(enemySpawnPoint.X), int(enemySpawnPoint.Y), 0, 0)
		g.AddMob(newEnemyMob) // Add a new enemy mob to the game

		newAllyMob := NewMob(fmt.Sprintf("Mob%d", len(g.Mobs)+1), TeamPlayer, 100, int(allySpawnPoint.X), int(allySpawnPoint.Y), 0, 0)
		g.AddMob(newAllyMob) // Add a new mob to the game

		g.SpawnTicks = g.SpawnInterval // Reset spawn ticks
	} else {
		g.SpawnTicks-- // Decrement spawn ticks
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
		// Draw debug info
		debugInfo := g.PrintBoidParams()
		ebitenutil.DebugPrint(screen, debugInfo)
	}

	enemyCount := 0
	allyCount := 0
	for _, mob := range g.Mobs {
		if mob.Team == TeamEnemy {
			enemyCount++ // Count enemy mobs
		} else if mob.Team == TeamPlayer {
			allyCount++ // Count ally mobs
		}
	}

	allyPercent := 0
	if len(g.Mobs) > 0 {
		allyPercent = int(float64(allyCount) / float64(len(g.Mobs)) * 100) // Calculate percentage of ally mobs
	}
	enemyPercent := 0
	if len(g.Mobs) > 0 {
		enemyPercent = int(float64(enemyCount) / float64(len(g.Mobs)) * 100) // Calculate percentage of enemy mobs
	}

	ebitenutil.DebugPrintAt(screen, "Ally Mobs: "+fmt.Sprint(allyCount), 10, 500) // Display the number of ally mobs
	ebitenutil.DebugPrintAt(screen, "Enemy Mobs: "+fmt.Sprint(enemyCount), 10, 530) // Display the number of enemy mobs
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Ally Percentage: %d%%", allyPercent), 100, 500) // Display the percentage of ally mobs
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Enemy Percentage: %d%%", enemyPercent), 100, 530) // Display the percentage of enemy mobs
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

func (g *Game) GetMobAtPosition(pos *physics.Vec2) *Mob {
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

// Add a function to print boid params for debug overlay
func (g *Game) PrintBoidParams() string {
	s := "Boid Params (select: 1-9,0,q,w,e,r | +/- to adjust):\n"
	for i, name := range boidParamNames {
		val := *boidParamPtrs[i]
		cursor := " "
		if i == selectedBoidParam {
			cursor = ">"
		}
		s += cursor + name + ": " + fmt.Sprintf("%.2f", val) + "\n"
	}
	return s
}
