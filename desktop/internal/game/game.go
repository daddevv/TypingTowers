package game

import (
	"errors"
	"fmt"
	"image/color"
	"os"
	"path/filepath"
	"sync"
	"td/internal/entity"
	"td/internal/ui"
	"td/internal/world"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	lua "github.com/yuin/gopher-lua"
)

type Game struct {
	Frame *ui.BaseScreen
	Level             world.Level
	Player            *entity.Player // Use concrete type for health access
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

	// Stats
	Misses            int
	HighestStreak     int
	CurrentStreak     int
	LuckyHits         int // Secondary mob hit instead of front
	GameOver          bool
	L                 *lua.LState // Lua VM for scripting
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
	L := lua.NewState()
	game := &Game{
		Frame:             ui.NewBaseScreen(1920, 1080),
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
		WaveToastTimer:    3.5,
		Misses:            0,
		HighestStreak:     0,
		CurrentStreak:     0,
		LuckyHits:         0,
		L:                 L,
	}
	RegisterGameAPI(L, game)
	game.loadLuaPlugins("lua") // <-- changed from "plugins" to "lua"
	return game
}

func (g *Game) Update() error {
	if g.GameOver {
		// Handle input for game over menu
		if inpututil.IsKeyJustPressed(ebiten.KeyR) {
			return errors.New("restart_level")
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyM) {
			return errors.New("return_to_menu")
		}
		return nil
	}
	if g.LevelComplete {
		// Allow returning to menu after level complete
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			return errors.New("return_to_menu")
		}
		return nil
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

	// --- Mob cleanup and wave logic ---
	activeMobs := g.Mobs[:0]
	for _, mob := range g.Mobs {
		// Remove mobs that are dead or have exceeded grace period after defeat
		if mobIntf, ok := mob.(interface{ IsDead() bool }); ok && mobIntf.IsDead() {
			continue
		}
		activeMobs = append(activeMobs, mob)
	}
	g.Mobs = activeMobs

	// Wave progression: if all mobs defeated for this wave, advance wave and show toast
	if g.CurrentWave < len(g.Level.Waves) {
		wave := g.Level.Waves[g.CurrentWave]
		if g.WaveEnemyDefeated >= wave.EnemyCount && len(g.Mobs) == 0 {
			g.CurrentWave++
			g.WaveEnemyDefeated = 0
			g.ShowWaveToast = true
			g.WaveToastTimer = 2.5 // seconds
			return nil // Pause game logic while toast is up
		}
	}
	if g.CurrentWave >= len(g.Level.Waves) {
		g.LevelComplete = true
	}
	if g.ShowWaveToast {
		g.WaveToastTimer -= 1.0 / 60.0 // assuming 60 FPS
		if g.WaveToastTimer <= 0 {
			g.ShowWaveToast = false
		}
		return nil // Prevent mob spawning while toast is up
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

// drawKeyboardIndicators draws a QWERTY keyboard layout (unshifted and shifted) with possible letters highlighted.
func drawKeyboardIndicators(screen *ebiten.Image, possible map[string]bool) {
	font := ui.Font("Game-Bold", 24)
	startX := 18.0
	startY := 18.0
	rowStep := 38.0
	colStep := 38.0

	// QWERTY layout (unshifted)
	rows := [][]string{
		{"q", "w", "e", "r", "t", "y", "u", "i", "o", "p"},
		{"a", "s", "d", "f", "g", "h", "j", "k", "l"},
		{"z", "x", "c", "v", "b", "n", "m"},
	}
	rowOffsets := []float64{0, colStep * 0.6, colStep * 1.5} // offset 2nd and 3rd rows
	// Shifted layout (uppercase and symbols)
	shiftRows := [][]string{
		{"Q", "W", "E", "R", "T", "Y", "U", "I", "O", "P"},
		{"A", "S", "D", "F", "G", "H", "J", "K", "L"},
		{"Z", "X", "C", "V", "B", "N", "M"},
	}
	shiftRowOffsets := rowOffsets
	// Number row (unshifted and shifted)
	numRow := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	shiftNumRow := []string{"!", "@", "#", "$", "%", "^", "&", "*", "(", ")"}
	// Symbol row (unshifted and shifted)
	symbolRow := []string{"-", "=", "[", "]", ";", "'", ",", ".", "/"}
	shiftSymbolRow := []string{"_", "+", "{", "}", ":", "\"", "<", ">", "?"}

	// Expanded background quad for both keyboards
	quadWidth := 13*colStep + 48.0 // wider for offsets and symbols
	quadHeight := 2*startY + 2*rowStep + 10*rowStep + 32.0 // much taller for both keyboards
	quad := ebiten.NewImage(int(quadWidth), int(quadHeight))
	quad.Fill(color.RGBA{30, 30, 30, 220})
	quadOpts := &ebiten.DrawImageOptions{}
	quadOpts.GeoM.Translate(startX-9, startY-9)
	screen.DrawImage(quad, quadOpts)

	// Draw number row (unshifted)
	for i, ch := range numRow {
		x := startX + float64(i)*colStep
		y := startY
		opts := &text.DrawOptions{}
		opts.GeoM.Translate(x, y)
		if possible[ch] {
			opts.ColorScale.ScaleWithColor(color.White)
		} else {
			opts.ColorScale.ScaleWithColor(color.RGBA{120, 120, 120, 255})
		}
		text.Draw(screen, ch, font, opts)
	}
	// Draw QWERTY rows (unshifted)
	for rowIdx, row := range rows {
		offset := rowOffsets[rowIdx]
		for colIdx, ch := range row {
			x := startX + offset + float64(colIdx)*colStep
			y := startY + float64(rowIdx+1)*rowStep
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
	// Draw symbol row (unshifted)
	for i, ch := range symbolRow {
		x := startX + colStep*2.5 + float64(i)*colStep
		y := startY + 4*rowStep
		opts := &text.DrawOptions{}
		opts.GeoM.Translate(x, y)
		if possible[ch] {
			opts.ColorScale.ScaleWithColor(color.White)
		} else {
			opts.ColorScale.ScaleWithColor(color.RGBA{120, 120, 120, 255})
		}
		text.Draw(screen, ch, font, opts)
	}

	// Draw shifted (shift) keyboard below
	shiftY := startY + 5*rowStep + 24
	// Draw number row (shifted)
	for i, ch := range shiftNumRow {
		x := startX + float64(i)*colStep
		y := shiftY
		opts := &text.DrawOptions{}
		opts.GeoM.Translate(x, y)
		if possible[ch] {
			opts.ColorScale.ScaleWithColor(color.White)
		} else {
			opts.ColorScale.ScaleWithColor(color.RGBA{120, 120, 120, 255})
		}
		text.Draw(screen, ch, font, opts)
	}
	// Draw QWERTY rows (shifted)
	for rowIdx, row := range shiftRows {
		offset := shiftRowOffsets[rowIdx]
		for colIdx, ch := range row {
			x := startX + offset + float64(colIdx)*colStep
			y := shiftY + float64(rowIdx+1)*rowStep
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
	// Draw symbol row (shifted)
	for i, ch := range shiftSymbolRow {
		x := startX + colStep*2.5 + float64(i)*colStep
		y := shiftY + 4*rowStep
		opts := &text.DrawOptions{}
		opts.GeoM.Translate(x, y)
		if possible[ch] {
			opts.ColorScale.ScaleWithColor(color.White)
		} else {
			opts.ColorScale.ScaleWithColor(color.RGBA{120, 120, 120, 255})
		}
		text.Draw(screen, ch, font, opts)
	}

	// Add labels
	labelFont := ui.Font("Game-Bold", 20)
	labelOpts := &text.DrawOptions{}
	labelOpts.GeoM.Translate(startX, startY-16)
	labelOpts.ColorScale.ScaleWithColor(color.RGBA{200, 200, 255, 255})
	text.Draw(screen, "Keyboard (unshifted)", labelFont, labelOpts)
	labelOpts2 := &text.DrawOptions{}
	labelOpts2.GeoM.Translate(startX, shiftY-16)
	labelOpts2.ColorScale.ScaleWithColor(color.RGBA{200, 200, 255, 255})
	text.Draw(screen, "Keyboard (shift)", labelFont, labelOpts2)
}

// drawMobChances draws the mob spawn chances box on the right side of the HUD.
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
	drawKeyboardIndicators(screen, possible)
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

	// Draw wave announcement toast
	if g.ShowWaveToast && g.CurrentWave < len(g.Level.Waves) {
		wave := g.Level.Waves[g.CurrentWave]
		msg := fmt.Sprintf("Wave %d!", wave.WaveNumber)
		font := ui.Font("Game-Bold", 64)
		w := textWidth(font, msg)
		opts := &text.DrawOptions{}
		opts.GeoM.Translate(960-float64(w)/2, 200)
		opts.ColorScale.ScaleWithColor(color.RGBA{255, 255, 0, 255})
		text.Draw(screen, msg, font, opts)
	}

	// --- Table-driven HUD via Lua ---
	hudTable := g.L.GetGlobal("HUD")
	if tbl, ok := hudTable.(*lua.LTable); ok {
		// Health
		if health := tbl.RawGetString("health"); health.Type() == lua.LTTable {
			h := health.(*lua.LTable)
			x := int(getLuaNumber(h, "x", 30))
			y := int(getLuaNumber(h, "y", 900))
			fontName := getLuaString(h, "font", "Game-Bold")
			fontSize := getLuaNumber(h, "fontSize", 32)
			colorArr := getLuaColor(h, "color", color.RGBA{255, 80, 80, 255})
			format := getLuaString(h, "format", "Health: %d")
			healthStr := fmt.Sprintf(format, g.Player.Health)
			healthFont := ui.Font(fontName, fontSize)
			healthOpts := &text.DrawOptions{}
			healthOpts.GeoM.Translate(float64(x), float64(y))
			healthOpts.ColorScale.ScaleWithColor(colorArr)
			text.Draw(screen, healthStr, healthFont, healthOpts)
		}
		// Stats
		if stats := tbl.RawGetString("stats"); stats.Type() == lua.LTTable {
			s := stats.(*lua.LTable)
			x := int(getLuaNumber(s, "x", 30))
			y := int(getLuaNumber(s, "y", 940))
			fontName := getLuaString(s, "font", "Game-Bold")
			fontSize := getLuaNumber(s, "fontSize", 24)
			colorArr := getLuaColor(s, "color", color.White)
			fields := s.RawGetString("fields")
			if fieldsTbl, ok := fields.(*lua.LTable); ok {
				statsFont := ui.Font(fontName, fontSize)
				row := 0
				fieldsTbl.ForEach(func(_, v lua.LValue) {
					if field, ok := v.(*lua.LTable); ok {
						name := getLuaString(field, "name", "")
						key := getLuaString(field, "key", "")
						val := 0
						switch key {
						case "Misses":
							val = g.Misses
						case "HighestStreak":
							val = g.HighestStreak
						case "LuckyHits":
							val = g.LuckyHits
						}
						label := fmt.Sprintf("%s: %d", name, val)
						opts := &text.DrawOptions{}
						opts.GeoM.Translate(float64(x), float64(y+row*30))
						opts.ColorScale.ScaleWithColor(colorArr)
						text.Draw(screen, label, statsFont, opts)
						row++
					}
				})
			}
		}
	} else {
		// fallback to hardcoded HUD if no HUD table
		// Draw player health
		healthFont := ui.Font("Game-Bold", 32)
		healthStr := fmt.Sprintf("Health: %d", g.Player.Health)
		healthOpts := &text.DrawOptions{}
		healthOpts.GeoM.Translate(30, 900)
		healthOpts.ColorScale.ScaleWithColor(color.RGBA{255, 80, 80, 255})
		text.Draw(screen, healthStr, healthFont, healthOpts)
		// Draw stats near player (bottom left)
		statsFont := ui.Font("Game-Bold", 24)
		statsY := 940.0
		missesOpts := &text.DrawOptions{}
		missesOpts.GeoM.Translate(30, statsY)
		text.Draw(screen, fmt.Sprintf("Misses: %d", g.Misses), statsFont, missesOpts)
		streakOpts := &text.DrawOptions{}
		streakOpts.GeoM.Translate(30, statsY+30)
		text.Draw(screen, fmt.Sprintf("Highest Streak: %d", g.HighestStreak), statsFont, streakOpts)
		luckyOpts := &text.DrawOptions{}
		luckyOpts.GeoM.Translate(30, statsY+60)
		text.Draw(screen, fmt.Sprintf("Lucky Hits: %d", g.LuckyHits), statsFont, luckyOpts)
	}

	// Game Over overlay
	if g.GameOver {
		msg := "Game Over!"
		msgFont := ui.Font("Game-Bold", 64)
		msgWidth := textWidth(msgFont, msg)
		msgBoxWidth := msgWidth + 120
		msgBoxHeight := 220.0
		msgBoxX := 1920.0/2 - msgBoxWidth/2
		msgBoxY := 400.0
		msgQuad := ebiten.NewImage(int(msgBoxWidth), int(msgBoxHeight))
		msgQuad.Fill(color.RGBA{30, 30, 30, 240})
		msgQuadOpts := &ebiten.DrawImageOptions{}
		msgQuadOpts.GeoM.Translate(msgBoxX, msgBoxY)
		screen.DrawImage(msgQuad, msgQuadOpts)
		msgOpts := &text.DrawOptions{}
		msgOpts.GeoM.Translate(1920/2-msgWidth/2, msgBoxY+60)
		msgOpts.ColorScale.ScaleWithColor(color.RGBA{255, 80, 80, 255})
		text.Draw(screen, msg, msgFont, msgOpts)
		// Draw buttons
		btnFont := ui.Font("Game-Bold", 36)
		btn1 := "[R] Restart Level"
		btn2 := "[M] Main Menu"
		btn1Width := textWidth(btnFont, btn1)
		btn2Width := textWidth(btnFont, btn2)
		btn1Opts := &text.DrawOptions{}
		btn1Opts.GeoM.Translate(1920/2-btn1Width/2, msgBoxY+120)
		btn1Opts.ColorScale.ScaleWithColor(color.White)
		text.Draw(screen, btn1, btnFont, btn1Opts)
		btn2Opts := &text.DrawOptions{}
		btn2Opts.GeoM.Translate(1920/2-btn2Width/2, msgBoxY+170)
		btn2Opts.ColorScale.ScaleWithColor(color.White)
		text.Draw(screen, btn2, btnFont, btn2Opts)
	}
}

func (g *Game) loadLuaPlugins(dir string) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return
	}
	for _, f := range files {
		if filepath.Ext(f.Name()) == ".lua" {
			script, err := os.ReadFile(filepath.Join(dir, f.Name()))
			if err == nil {
				g.L.DoString(string(script))
			}
		}
	}
}

func (g *Game) CallLuaFunction(funcName string, args ...lua.LValue) (lua.LValue, error) {
	fn := g.L.GetGlobal(funcName)
	if fn.Type() == lua.LTFunction {
		err := g.L.CallByParam(lua.P{
			Fn:      fn,
			NRet:    1,
			Protect: true,
		}, args...)
		if err != nil {
			return lua.LNil, err
		}
		ret := g.L.Get(-1)
		g.L.Pop(1)
		return ret, nil
	}
	return lua.LNil, fmt.Errorf("function %s not found", funcName)
}

// Helper functions for extracting Lua table fields
func getLuaString(tbl *lua.LTable, key, def string) string {
	v := tbl.RawGetString(key)
	if str, ok := v.(lua.LString); ok {
		return string(str)
	}
	return def
}
func getLuaNumber(tbl *lua.LTable, key string, def float64) float64 {
	v := tbl.RawGetString(key)
	if num, ok := v.(lua.LNumber); ok {
		return float64(num)
	}
	return def
}
func getLuaColor(tbl *lua.LTable, key string, def color.Color) color.Color {
	v := tbl.RawGetString(key)
	if arr, ok := v.(*lua.LTable); ok {
		var rgba [4]uint8
		for i := 1; i <= 4; i++ {
			val := arr.RawGetInt(i)
			if num, ok := val.(lua.LNumber); ok {
				rgba[i-1] = uint8(num)
			}
		}
		return color.RGBA{rgba[0], rgba[1], rgba[2], rgba[3]}
	}
	return def
}
