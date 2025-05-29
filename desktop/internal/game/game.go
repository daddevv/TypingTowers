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
	Level             world.Level
	Player            entity.Entity
	Mobs              []entity.Entity
	Projectiles       []*entity.Projectile
	InputHandler      *InputHandler
	MobSpawner        *entity.MobSpawner
	LastUpdate        time.Time
	LevelComplete     bool // True if level is complete
	CurrentWave       int  // Current wave index (0-based)
	WaveEnemyDefeated int  // Number of enemies defeated in current wave
	ShowWaveToast     bool
	WaveToastTimer    float64
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
		Level:             opts.Level,
		Player:            player,
		Mobs:              entity.EmptyList(),
		Projectiles:       make([]*entity.Projectile, 0),
		InputHandler:      inputHandler,
		MobSpawner:        mobSpawner,
		LastUpdate:        time.Now(),
		LevelComplete:     false,
		CurrentWave:       0,
		WaveEnemyDefeated: 0,
		ShowWaveToast:     true,
		WaveToastTimer:    3.5, // longer toast duration
		// Ensure no mobs spawn at level start
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
	newMob := g.MobSpawner.Update(deltaTime, g.WaveEnemyDefeated)
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
			if beachballMob, ok := mob.(*entity.BeachballMob); ok && (beachballMob.Dead || allLettersInactive(beachballMob)) {
				beachballMob.Dead = true
				g.WaveEnemyDefeated++
				g.MobSpawner.SpeedUpOverTime(g.WaveEnemyDefeated)
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
			g.ShowWaveToast = true
			g.WaveToastTimer = 3.5
			return nil // Pause wave progression until toast is gone
		}
	}
	if g.CurrentWave >= len(g.Level.Waves) {
		g.LevelComplete = true
	}
	if g.ShowWaveToast {
		g.WaveToastTimer -= 1.0 / 60.0 // assuming 60 FPS
		if g.WaveToastTimer <= 0 {
			g.ShowWaveToast = false
			// Only allow mob spawning after toast disappears
			g.MobSpawner.NextSpawnTime = 0.5 // small delay after toast
		}
		return nil // Prevent wave from progressing/spawning while toast is up
	}
	// Prevent mob spawning if toast is up
	if g.ShowWaveToast {
		return nil
	}
	return nil
}

func allLettersInactive(mob *entity.BeachballMob) bool {
	for _, l := range mob.Letters {
		if l.State != entity.LetterInactive {
			return false
		}
	}
	return true
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

func textWidth(face *text.GoTextFace, s string) float64 {
	// Fallback: estimate width as len(s) * size * 0.6 (approximate for monospace)
	return float64(len(s)) * face.Size * 0.6
}

func drawLetterIndicators(screen *ebiten.Image, possible map[string]bool) {
	// Stacked left-to-right, 13 per row, top left, with background quad
	font := ui.Font("Game-Bold", 28)
	startX := 18.0
	startY := 18.0
	rowStep := 32.0
	colStep := 38.0 // slightly tighter
	letters := []string{}
	// A-Z
	for i := 0; i < 26; i++ {
		letters = append(letters, fmt.Sprintf("%c", 'A'+i))
	}
	// a-z
	for i := 0; i < 26; i++ {
		letters = append(letters, fmt.Sprintf("%c", 'a'+i))
	}
	// 0-9 and symbols
	symbols := "0123456789!@#$%^&*()-_=+[]{};:'\",.<>/?|`~"
	for _, r := range symbols {
		letters = append(letters, string(r))
	}
	// Calculate grid size
	total := len(letters)
	cols := 13
	rows := (total + cols - 1) / cols
	quadWidth := float64(cols)*colStep + 18.0
	quadHeight := float64(rows)*rowStep + 18.0
	quad := ebiten.NewImage(int(quadWidth), int(quadHeight))
	quad.Fill(color.RGBA{30, 30, 30, 220})
	quadOpts := &ebiten.DrawImageOptions{}
	quadOpts.GeoM.Translate(startX-9, startY-9)
	screen.DrawImage(quad, quadOpts)
	for idx, ch := range letters {
		row := idx / cols
		col := idx % cols
		x := startX + float64(col)*colStep
		y := startY + float64(row)*rowStep
		opts := &text.DrawOptions{}
		opts.GeoM.Translate(x, y)
		if possible[ch] {
			opts.ColorScale.ScaleWithColor(color.White)
		} else {
			opts.ColorScale.ScaleWithColor(color.RGBA{120, 120, 120, 255})
		}
		text.Draw(screen, ch, font, opts)
	}
}

func drawMobChances(screen *ebiten.Image, mobChances []struct {
	Type   string
	Chance float64
}) {
	font := ui.Font("Game-Bold", 28)
	rowStep := 40.0
	paddingX := 24.0
	paddingY := 18.0
	maxLabelWidth := 0.0
	labels := make([]string, len(mobChances))
	for i, mob := range mobChances {
		labels[i] = fmt.Sprintf("%s: %d%%", mob.Type, int(mob.Chance*100))
		w := textWidth(font, labels[i])
		if w > maxLabelWidth {
			maxLabelWidth = w
		}
	}
	boxWidth := maxLabelWidth + paddingX*2
	boxHeight := float64(len(mobChances))*rowStep + paddingY*2
	boxX := 1920.0 - boxWidth - 18.0
	boxY := 30.0 // move closer to top
	quad := ebiten.NewImage(int(boxWidth), int(boxHeight))
	quad.Fill(color.RGBA{30, 30, 30, 220})
	quadOpts := &ebiten.DrawImageOptions{}
	quadOpts.GeoM.Translate(boxX, boxY)
	screen.DrawImage(quad, quadOpts)
	for i, label := range labels {
		opts := &text.DrawOptions{}
		opts.GeoM.Translate(boxX+paddingX, boxY+paddingY+float64(i)*rowStep)
		opts.ColorScale.ScaleWithColor(color.White)
		text.Draw(screen, label, font, opts)
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Level.DrawBackground(screen)

	// --- Draw header background quad (skinnier for just the text) ---
	headerWidth := 600.0
	headerHeight := 150.0
	headerX := 1920.0/2 - headerWidth/2
	headerY := 20.0
	headerQuad := ebiten.NewImage(int(headerWidth), int(headerHeight))
	headerQuad.Fill(color.RGBA{30, 30, 30, 220})
	headerQuadOpts := &ebiten.DrawImageOptions{}
	headerQuadOpts.GeoM.Translate(headerX, headerY)
	screen.DrawImage(headerQuad, headerQuadOpts)

	// Draw level name at top center (move yellow title higher)
	levelName := g.Level.Name
	fontTitle := ui.Font("Game-Bold", 48)
	nameWidth := textWidth(fontTitle, levelName)
	nameOpts := &text.DrawOptions{}
	nameOpts.GeoM.Translate(1920/2-nameWidth/2, headerY+8) // move up
	nameOpts.ColorScale.ScaleWithColor(color.RGBA{255, 255, 0, 255})
	text.Draw(screen, levelName, fontTitle, nameOpts)

	// Draw world/level number and wave info below title
	levelNumStr := fmt.Sprintf("World %d - Level %d", g.Level.WorldNumber, g.Level.LevelNumber)
	levelNumFont := ui.Font("Game-Bold", 32)
	levelNumWidth := textWidth(levelNumFont, levelNumStr)
	levelNumOpts := &text.DrawOptions{}
	levelNumOpts.GeoM.Translate(1920/2-levelNumWidth/2, headerY+60)
	levelNumOpts.ColorScale.ScaleWithColor(color.White)
	text.Draw(screen, levelNumStr, levelNumFont, levelNumOpts)

	if g.CurrentWave < len(g.Level.Waves) {
		wave := g.Level.Waves[g.CurrentWave]
		waveStr := fmt.Sprintf("Wave %d: %d/%d enemies defeated", wave.WaveNumber, g.WaveEnemyDefeated, wave.EnemyCount)
		waveFont := ui.Font("Game-Bold", 32)
		waveWidth := textWidth(waveFont, waveStr)
		waveOpts := &text.DrawOptions{}
		waveOpts.GeoM.Translate(1920/2-waveWidth/2, headerY+100)
		waveOpts.ColorScale.ScaleWithColor(color.White)
		text.Draw(screen, waveStr, waveFont, waveOpts)
	}

	// Draw left-side letter indicators
	possible := map[string]bool{}
	if g.CurrentWave < len(g.Level.Waves) {
		for _, l := range g.Level.Waves[g.CurrentWave].PossibleLetters {
			possible[l] = true
		}
	}
	drawLetterIndicators(screen, possible)
	// Draw right-side mob chances with background
	if g.CurrentWave < len(g.Level.Waves) {
		wave := g.Level.Waves[g.CurrentWave]
		mobChances := make([]struct {
			Type   string
			Chance float64
		}, len(wave.MobChances))
		for i, m := range wave.MobChances {
			mobChances[i] = struct {
				Type   string
				Chance float64
			}{m.Type, m.Chance}
		}
		drawMobChances(screen, mobChances)
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

	// --- Level Complete Toast ---
	if g.LevelComplete {
		msg := "Level Complete!"
		msgFont := ui.Font("Game-Bold", 64)
		msgWidth := textWidth(msgFont, msg)
		msgBoxWidth := msgWidth + 120
		msgBoxHeight := 120.0
		msgBoxX := 1920.0/2 - msgBoxWidth/2
		msgBoxY := 540.0
		msgQuad := ebiten.NewImage(int(msgBoxWidth), int(msgBoxHeight))
		msgQuad.Fill(color.RGBA{30, 30, 30, 220})
		msgQuadOpts := &ebiten.DrawImageOptions{}
		msgQuadOpts.GeoM.Translate(msgBoxX, msgBoxY)
		screen.DrawImage(msgQuad, msgQuadOpts)
		msgOpts := &text.DrawOptions{}
		msgOpts.GeoM.Translate(1920/2-msgWidth/2, msgBoxY+msgBoxHeight/2-10) // center vertically
		msgOpts.ColorScale.ScaleWithColor(color.RGBA{0, 255, 0, 255})
		text.Draw(screen, msg, msgFont, msgOpts)
	}

	// --- Wave Toast ---
	if g.ShowWaveToast && g.CurrentWave < len(g.Level.Waves) {
		wave := g.Level.Waves[g.CurrentWave]
		toastTitle := fmt.Sprintf("Wave %d", wave.WaveNumber)
		lettersStr := fmt.Sprintf("Letters: %v", wave.PossibleLetters)
		toastFont := ui.Font("Game-Bold", 48)
		lettersFont := ui.Font("Game-Bold", 36)
		titleWidth := textWidth(toastFont, toastTitle)
		lettersWidth := textWidth(lettersFont, lettersStr)
		toastWidth := titleWidth
		if lettersWidth > toastWidth {
			toastWidth = lettersWidth
		}
		paddingX := 60.0
		paddingY := 40.0
		boxWidth := toastWidth + paddingX*2
		boxHeight := 120.0 + paddingY*2
		boxX := 1920/2 - boxWidth/2
		boxY := 400
		quad := ebiten.NewImage(int(boxWidth), int(boxHeight))
		quad.Fill(color.RGBA{30, 30, 30, 220})
		quadOpts := &ebiten.DrawImageOptions{}
		quadOpts.GeoM.Translate(boxX, float64(boxY))
		screen.DrawImage(quad, quadOpts)
		titleOpts := &text.DrawOptions{}
		titleOpts.GeoM.Translate(1920/2-titleWidth/2, float64(boxY)+paddingY)
		titleOpts.ColorScale.ScaleWithColor(color.RGBA{255, 255, 0, 255})
		text.Draw(screen, toastTitle, toastFont, titleOpts)
		lettersOpts := &text.DrawOptions{}
		lettersOpts.GeoM.Translate(1920/2-lettersWidth/2, float64(boxY)+paddingY+60)
		lettersOpts.ColorScale.ScaleWithColor(color.White)
		text.Draw(screen, lettersStr, lettersFont, lettersOpts)
	}
}
