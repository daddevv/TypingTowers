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
	LevelComplete bool // True if level is complete
	CurrentWave   int // Current wave index (0-based)
	WaveEnemyDefeated int // Number of enemies defeated in current wave
}

func NewGame(opts GameOptions) *Game {
	player := entity.NewPlayer()
	inputHandler := NewInputHandler(player.GetPosition())
	// Use a LetterPool for endless mode (letters expand over time)
	letterPool := entity.NewDefaultLetterPool()
	var mobSpawner *entity.MobSpawner
	if opts.MobConfigs != nil && len(opts.MobConfigs) > 0 {
		mobSpawner = entity.NewMobSpawnerWithConfigs(letterPool, opts.MobConfigs)
	} else {
		mobSpawner = entity.NewMobSpawner(letterPool)
	}
	return &Game{
		Level:        opts.Level,
		Player:       player,
		Mobs:         entity.EmptyList(),
		Projectiles:  make([]*entity.Projectile, 0),
		InputHandler: inputHandler,
		MobSpawner:   mobSpawner,
		LastUpdate:   time.Now(),
		Score:        0,
		LevelComplete: false,
		CurrentWave:   0,
		WaveEnemyDefeated: 0,
	}
}

func (g *Game) Update() error {
	if g.LevelComplete {
		return nil // No updates if level is complete
	}
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
	newMob := g.MobSpawner.Update(deltaTime, g.Score)
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
	activeMobs := g.Mobs[:0]
	for _, mob := range g.Mobs {
		pos := mob.GetPosition()
		if pos.X > 200 {
			activeMobs = append(activeMobs, mob)
		} else {
			// Check if mob was defeated (all letters typed) vs just went off screen
			if beachballMob, ok := mob.(*entity.BeachballMob); ok && beachballMob.Dead {
				g.Score++
				g.WaveEnemyDefeated++
				g.MobSpawner.SpeedUpOverTime(g.Score)
			}
		}
	}
	g.Mobs = activeMobs

	// Wave/level progression logic
	if g.CurrentWave < len(g.Level.Waves) {
		wave := g.Level.Waves[g.CurrentWave]
		if g.WaveEnemyDefeated >= wave.EnemyCount {
			g.CurrentWave++
			g.WaveEnemyDefeated = 0
			// Optionally: update possible letters for next wave, etc.
		}
	}
	if g.Score >= g.Level.LevelCompleteScore {
		g.LevelComplete = true
	}
	return nil
}

// checkProjectileCollisions checks for collisions between projectiles and mobs
// Projectiles now only provide visual feedback - letter states are advanced immediately in input handler
func (g *Game) checkProjectileCollisions() {
	for _, projectile := range g.Projectiles {
		if !projectile.IsActive() || projectile.DamageDealt {
			continue
		}

		// Only check collision with the projectile's intended target mob
		if projectile.TargetMob == nil {
			continue
		}
		mob := projectile.TargetMob
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
		}
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Level.DrawBackground(screen)

	// Draw level name at top center
	levelName := g.Level.Name
	font := ui.Font("Game-Bold", 48)
	nameOpts := &text.DrawOptions{}
	nameOpts.GeoM.Translate(1920/2-300, 60)
	nameOpts.ColorScale.ScaleWithColor(color.RGBA{255, 255, 0, 255})
	text.Draw(screen, levelName, font, nameOpts)

	// Draw world/level number
	levelNumStr := fmt.Sprintf("World %d - Level %d", g.Level.WorldNumber, g.Level.LevelNumber)
	levelNumFont := ui.Font("Game-Bold", 32)
	levelNumOpts := &text.DrawOptions{}
	levelNumOpts.GeoM.Translate(1920/2-200, 120)
	levelNumOpts.ColorScale.ScaleWithColor(color.White)
	text.Draw(screen, levelNumStr, levelNumFont, levelNumOpts)

	// Draw score in top left corner
	scoreStr := fmt.Sprintf("Score: %d", g.Score)
	scoreOpts := &text.DrawOptions{}
	scoreOpts.GeoM.Translate(20, 50)
	scoreOpts.ColorScale.ScaleWithColor(color.White)
	fontSmall := ui.Font("Game-Bold", 32)
	text.Draw(screen, scoreStr, fontSmall, scoreOpts)

	// Draw wave info
	if g.CurrentWave < len(g.Level.Waves) {
		wave := g.Level.Waves[g.CurrentWave]
		waveStr := fmt.Sprintf("Wave %d: %d/%d enemies defeated", wave.WaveNumber, g.WaveEnemyDefeated, wave.EnemyCount)
		waveFont := ui.Font("Game-Bold", 32)
		waveOpts := &text.DrawOptions{}
		waveOpts.GeoM.Translate(20, 100)
		waveOpts.ColorScale.ScaleWithColor(color.White)
		text.Draw(screen, waveStr, waveFont, waveOpts)
	}

	entities := append(g.Mobs, g.Player)
	// TODO: Sort entities by Z-index (smallest Y first) if needed
	for _, entity := range entities {
		entity.Draw(screen)
	}

	// Draw projectiles
	for _, projectile := range g.Projectiles {
		projectile.Draw(screen)
	}

	// Draw level complete message if needed
	if g.LevelComplete {
		msg := "Level Complete!"
		msgFont := ui.Font("Game-Bold", 64)
		msgOpts := &text.DrawOptions{}
		msgOpts.GeoM.Translate(1920/2-320, 540)
		msgOpts.ColorScale.ScaleWithColor(color.RGBA{0,255,0,255})
		text.Draw(screen, msg, msgFont, msgOpts)
	}
}
