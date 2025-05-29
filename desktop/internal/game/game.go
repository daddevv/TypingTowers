package game

import (
	"errors"
	"sync"
	"td/internal/entity"
	"td/internal/world"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	Level        world.Level
	Player       entity.Entity
	Mobs         []entity.Entity
	Projectiles  []*entity.Projectile
	InputHandler *InputHandler
}

func NewGame(opts GameOptions) *Game {
	player := entity.NewPlayer()
	inputHandler := NewInputHandler(player.GetPosition())
	
	return &Game{
		Level:        opts.Level,
		Player:       player,
		Mobs:         entity.EmptyList(),
		Projectiles:  make([]*entity.Projectile, 0),
		InputHandler: inputHandler,
	}
}

func (g *Game) Update() error {
	// Handle pause
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return errors.New("pause")
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		// Spawn a new beachball mob when space is pressed
		mob := entity.NewBeachballMob(3, g.Level.PossibleLetters)
		if mob != nil {
			g.Mobs = append(g.Mobs, mob)
		} else {
			return errors.New("failed to create new beachball mob")
		}
	}

	// Update input handler with current player position
	g.InputHandler.SetPlayerPosition(g.Player.GetPosition())
	
	// Process input and create projectiles
	newProjectiles := g.InputHandler.ProcessInput(g.Mobs)
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
		}
	}
	g.Mobs = activeMobs

	// Remove inactive projectiles
	activeProjectiles := g.Projectiles[:0]
	for _, projectile := range g.Projectiles {
		if projectile.IsActive() {
			activeProjectiles = append(activeProjectiles, projectile)
		}
	}
	g.Projectiles = activeProjectiles

	return nil
}

// checkProjectileCollisions checks for collisions between projectiles and mobs
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
				
				// Collision detected - deactivate projectile
				projectile.Deactivate()
				projectile.DamageDealt = true
				break
			}
		}
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Level.DrawBackground(screen)
	// scoreStr := fmt.Sprintf("Score: %d", g.Score)
	// opts := &text.DrawOptions{}
	// opts.GeoM.Translate(10, 30)                  // position on screen
	// opts.ColorScale.ScaleWithColor(color.White)  // text color
	// font := ui.Font("Game-Bold", 48)			 // use the font source to get the font
	// text.Draw(screen, scoreStr, font, opts)

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
