package game

import (
	"errors"
	"fmt"
	"image/color"
	"sync"
	"td/internal/entity"
	"td/internal/ui"
	"td/internal/world"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Game struct {
	Level        world.Level
	Player       entity.Entity
	Mobs         []entity.Entity
	Projectiles  []*entity.Projectile
	InputHandler *InputHandler
	MobSpawner   *entity.MobSpawner
	LastUpdate   time.Time
	Score        int // Number of mobs defeated
}

func NewGame(opts GameOptions) *Game {
	player := entity.NewPlayer()
	inputHandler := NewInputHandler(player.GetPosition())
	mobSpawner := entity.NewMobSpawner(opts.Level.PossibleLetters)
	
	return &Game{
		Level:        opts.Level,
		Player:       player,
		Mobs:         entity.EmptyList(),
		Projectiles:  make([]*entity.Projectile, 0),
		InputHandler: inputHandler,
		MobSpawner:   mobSpawner,
		LastUpdate:   time.Now(),
		Score:        0,
	}
}

func (g *Game) Update() error {
	// Calculate delta time for smooth timing
	now := time.Now()
	deltaTime := now.Sub(g.LastUpdate).Seconds()
	g.LastUpdate = now

	// Handle pause
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return errors.New("pause")
	}

	// Optional: Allow manual spawning for testing (keep space bar functionality)
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		// Force spawn a new beachball mob when space is pressed
		mob := g.MobSpawner.ForceSpawn()
		g.Mobs = append(g.Mobs, mob)
	}

	// Update mob spawner and potentially spawn new mobs
	newMob := g.MobSpawner.Update(deltaTime)
	if newMob != nil {
		g.Mobs = append(g.Mobs, newMob)
	}

	// Update input handler with current player position
	g.InputHandler.SetPlayerPosition(g.Player.GetPosition())

	// Process input and create projectiles (pass current projectiles for reservation logic)
	newProjectiles := g.InputHandler.ProcessInput(g.Mobs, g.Projectiles)
	g.Projectiles = append(g.Projectiles, newProjectiles...)

	// --- Parallel mob updates ---
	var wg sync.WaitGroup
	errCh := make(chan error, len(g.Mobs))
	for m := range g.Mobs {
		wg.Add(1)
		go func(mob entity.Entity) {
			defer wg.Done()
			if err := mob.Update(); err != nil {
				errCh <- err
			}
		}(g.Mobs[m])
	}
	wg.Wait()
	close(errCh)
	for err := range errCh {
		if err != nil {
			return err
		}
	}

	// Update projectiles
	for _, projectile := range g.Projectiles {
		if err := projectile.Update(); err != nil {
			return err
		}
	}

	// Check projectile-mob collisions
	g.checkProjectileCollisions()

	// --- Collision and despawn logic ---
	// Remove mobs that go off-screen (X < -100) or are dead
	activeMobs := g.Mobs[:0]
	for _, mob := range g.Mobs {
		pos := mob.GetPosition()
		if pos.X > -100 {
			activeMobs = append(activeMobs, mob)
		} else {
			// Check if mob was defeated (all letters typed) vs just went off screen
			if beachballMob, ok := mob.(*entity.BeachballMob); ok && beachballMob.Dead {
				g.Score++
				g.MobSpawner.SpeedUpOverTime(g.Score)
			}
		}
	}
	g.Mobs = activeMobs

	// Remove inactive projectiles and decrement pending projectiles for missed shots
	activeProjectiles := g.Projectiles[:0]
	for _, projectile := range g.Projectiles {
		if projectile.IsActive() {
			activeProjectiles = append(activeProjectiles, projectile)
		} else {
			// If projectile became inactive but didn't deal damage, it missed
			if !projectile.DamageDealt {
				if beachballMob, ok := projectile.TargetMob.(*entity.BeachballMob); ok {
					beachballMob.PendingProjectiles--
					
					// If this mob is pending death and has no more pending projectiles, start death animation
					if beachballMob.PendingDeath && beachballMob.PendingProjectiles <= 0 {
						beachballMob.StartDeath()
					}
				}
			}
		}
	}
	g.Projectiles = activeProjectiles

	return nil
}

// checkProjectileCollisions checks for collisions between projectiles and mobs
// Projectiles now only provide visual feedback - letter states are advanced immediately in input handler
func (g *Game) checkProjectileCollisions() {
	for _, projectile := range g.Projectiles {
		if !projectile.IsActive() || projectile.DamageDealt {
			continue
		}
		
		for _, mob := range g.Mobs {
			mobPos := mob.GetPosition()
			projPos := projectile.GetPosition()
			
			// Simple collision detection - check if projectile is within mob bounds
			// Assuming mob is roughly 48x48 pixels (sprite size * scale)
			mobSize := 48.0 * 3.0 // sprite size * scale factor
			if projPos.X >= mobPos.X && projPos.X <= mobPos.X+mobSize &&
				projPos.Y >= mobPos.Y && projPos.Y <= mobPos.Y+mobSize {
				
				// Collision detected - deactivate projectile (letter states already advanced)
				projectile.Deactivate()
				projectile.DamageDealt = true
				
				// Decrement pending projectiles counter for this mob
				if beachballMob, ok := mob.(*entity.BeachballMob); ok {
					beachballMob.PendingProjectiles--
					
					// If this mob is pending death and has no more pending projectiles, start death animation
					if beachballMob.PendingDeath && beachballMob.PendingProjectiles <= 0 {
						beachballMob.StartDeath()
					}
				}
				break
			}
		}
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Level.DrawBackground(screen)
	
	// Draw score in top left corner
	scoreStr := fmt.Sprintf("Score: %d", g.Score)
	opts := &text.DrawOptions{}
	opts.GeoM.Translate(20, 50)                    // position on screen
	opts.ColorScale.ScaleWithColor(color.White)    // text color
	font := ui.Font("Game-Bold", 32)               // use the font source to get the font
	text.Draw(screen, scoreStr, font, opts)

	entities := append(g.Mobs, g.Player)
	// TODO: Sort entities by Z-index (smallest Y first) if needed
	for _, entity := range entities {
		entity.Draw(screen)
	}
	
	// Draw projectiles
	for _, projectile := range g.Projectiles {
		projectile.Draw(screen)
	}
}
