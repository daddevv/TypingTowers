package game

import (
	"encoding/json"
	"errors"
	"fmt"
	"image/color"
	"math"
	"math/rand"
	"os"
	"path/filepath"
	"sort"
	"time"

	"strings"
	"unicode"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const jamFlashDuration = 0.15
const conveyorSpeed = 200.0 // pixels per second for queue slide
const letterWidth = 13.0    // approximate width of a character
const SaveVersion = 1

var (
	mousePressed bool
	clickedTileX int
	clickedTileY int
	houses       = make(map[string]struct{})

	// ErrSaveVersion indicates the save file version is incompatible.
	ErrSaveVersion = errors.New("save file version mismatch")
)

type savedTower struct {
	X            float64
	Y            float64
	Damage       int
	Range        float64
	Rate         float64
	AmmoCapacity int
}

type savedGame struct {
	Version  int
	Gold     int
	Food     int
	Wave     int
	BaseHP   int
	Towers   []savedTower
	Settings Settings
	Skills   []string
}

// Game represents the game state and implements ebiten.Game interface.
type Game struct {
	screen      *ebiten.Image
	input       InputHandler
	towers      []*Tower
	mobs        []Enemy
	projectiles []*Projectile
	base        *Base
	hud         *HUD
	gameOver    bool
	paused      bool
	resources   ResourcePool

	shopOpen bool

	selectedTower int
	shopCursor    int

	cfg *Config

	currentWave   int
	spawnInterval float64
	spawnTicker   float64
	mobsToSpawn   int

	letterPool   []rune
	unlockStage  int
	techTree     *TechTree
	skillTree    *SkillTree
	achievements []string
	towerMods    TowerModifiers
	wpmBonus     int
	autoCollect  bool
	hotkeys      bool

	unlockedSkills map[string]bool
	skillMenuOpen  bool
	statsPanelOpen bool
	skillCursor    int
	skillCategory  SkillCategory

	techMenuOpen bool
	searchBuffer string
	techCursor   int

	score           int
	gameOverHandled bool
	history         *PerformanceHistory

	cursorX int
	cursorY int

	lastUpdate time.Time

	typing TypingStats

	// Per-word metrics
	currentWord WordStat
	wordHistory []WordStat

	pauseCursor    int
	settingsOpen   bool
	settingsCursor int

	buildMenuOpen bool
	buildCursor   int

	upgradeMenuOpen bool
	upgradeCursor   int

	sound    *SoundManager
	settings Settings

	flashTimer float64

	// Persistence
	saveDir       string
	saveSlot      int
	lastWaveSaved int

	slotMenuOpen bool
	slotCursor   int
	slotModeSave bool

	// Building integration
	queue      *QueueManager
	farmer     *Farmer
	lumberjack *Lumberjack
	miner      *Miner
	barracks   *Barracks
	military   *Military

	// Typing state for the queue - jam indicator
	queueJam bool

	// Command mode for power users
	commandMode   bool
	commandBuffer string

	// Tower selection system
	towerSelectMode bool
	towerLabels     map[string]int // label -> tower index

	// Static word processing location
	wordProcessX float64
	wordProcessY float64
	// Visual offset for conveyor belt animation
	conveyorOffset float64

	// High level state
	phase     GamePhase
	prevPhase GamePhase
	mainMenu  *MainMenu
	preGame   *PreGame
	quit      bool
}

// Gold returns the player's current gold amount.
func (g *Game) Gold() int { return g.resources.GoldAmount() }

// AddGold increases the player's gold.
func (g *Game) AddGold(n int) { g.resources.AddGold(n) }

// Queue returns the global word queue manager.
func (g *Game) Queue() *QueueManager { return g.queue }

// WordHistory returns the slice of completed word statistics.
func (g *Game) WordHistory() []WordStat { return g.wordHistory }

// SpendGold attempts to deduct the given amount of gold and returns true on success.
func (g *Game) SpendGold(n int) bool { return g.resources.Gold.Spend(n) }

// currentSavePath returns the file path for the active save slot.
func (g *Game) currentSavePath() string {
	name := fmt.Sprintf("save_slot%d.json", g.saveSlot)
	return filepath.Join(g.saveDir, name)
}

// SetSaveSlot changes the active save slot for subsequent saves/loads.
func (g *Game) SetSaveSlot(slot int) {
	if slot < 1 {
		slot = 1
	} else if slot > 3 {
		slot = 3
	}
	g.saveSlot = slot
}

// NewGame creates a new instance of the Game.
func NewGame() *Game {
	return NewGameWithConfig(DefaultConfig)
}

// NewGameWithConfig creates a new instance of the Game using the provided configuration.
func NewGameWithConfig(cfg Config) *Game {
	return NewGameWithHistory(cfg, &PerformanceHistory{})
}

// NewGameWithHistory allows supplying an existing performance history when creating a game.
func NewGameWithHistory(cfg Config, hist *PerformanceHistory) *Game {
	ebiten.SetWindowTitle("TypingTowers")
	ebiten.SetWindowSize(1920/8, 1080/8)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetFullscreen(true)

	g := &Game{
		history:         hist,
		score:           0,
		gameOverHandled: false,
		screen:          ebiten.NewImage(1920, 1080),
		input:           NewInput(),
		paused:          false,
		resources:       ResourcePool{},
		shopOpen:        false,
		selectedTower:   0,
		shopCursor:      0,
		currentWave:     1,
		spawnInterval:   cfg.SpawnInterval * 4.0, // Much slower spawning
		spawnTicker:     0,
		mobsToSpawn:     cfg.MobsPerWave,
		cfg:             &cfg,
		mobs:            make([]Enemy, 0),
		projectiles:     make([]*Projectile, 0),
		letterPool:      make([]rune, 0),
		unlockStage:     0,
		techTree:        DefaultTechTree(),
		skillTree:       func() *SkillTree { t, _ := SampleSkillTree(); return t }(),
		unlockedSkills:  make(map[string]bool),
		achievements:    make([]string, 0),
		towerMods:       TowerModifiers{DamageMult: 1, RangeMult: 1, FireRateMult: 1},
		wpmBonus:        0,
		autoCollect:     false,
		hotkeys:         false,
		typing:          NewTypingStats(),
		cursorX:         2,
		cursorY:         16,
		sound:           NewSoundManager(),
		settings:        DefaultSettings(),
		buildMenuOpen:   false,
		buildCursor:     0,
		upgradeMenuOpen: false,
		upgradeCursor:   0,
		techMenuOpen:    false,
		skillMenuOpen:   false,
		statsPanelOpen:  false,
		searchBuffer:    "",
		techCursor:      0,
		skillCursor:     0,
		skillCategory:   SkillOffense,
		flashTimer:      0,
		saveDir:         ".",
		saveSlot:        1,
		lastWaveSaved:   0,
		slotMenuOpen:    false,
		slotCursor:      0,
		slotModeSave:    true,
		queue:           NewQueueManager(),
		farmer:          NewFarmer(),
		lumberjack:      NewLumberjack(),
		miner:           NewMiner(),
		barracks:        NewBarracks(),
		military:        NewMilitary(),
		wordProcessX:    400,
		wordProcessY:    900,
		conveyorOffset:  0,
		commandMode:     false,
		commandBuffer:   "",
		towerSelectMode: false,
		towerLabels:     make(map[string]int),
		phase:           PhaseMainMenu,
		mainMenu:        NewMainMenu(),
		preGame:         NewPreGame(),
		currentWord:     WordStat{},
		wordHistory:     make([]WordStat, 0),
	}
	if g.sound != nil {
		g.sound.StartMusic()
	}

	tx, ty := tilePosition(1, 16)
	hp := cfg.BaseHealth
	if hp == 0 {
		hp = int(cfg.J)
	}
	if hp <= 0 {
		hp = BaseStartingHealth
	}
	g.base = NewBase(float64(tx+32), float64(ty+16), hp)
	if g.queue != nil {
		g.queue.SetBase(g.base)
	}

	// Wire up shared systems
	g.queue.SetBase(g.base)
	g.farmer.SetQueue(g.queue)
	g.lumberjack.SetQueue(g.queue)
	g.miner.SetQueue(g.queue)
	g.barracks.SetQueue(g.queue)
	g.barracks.SetMilitary(g.military)

	tx, ty = tilePosition(2, 16)
	tower := NewTower(g, float64(tx+16), float64(ty+16))
	tower.ApplyModifiers(g.towerMods)
	g.towers = []*Tower{tower}
	if tree, err := SampleSkillTree(); err == nil {
		g.skillTree = tree
	}
	g.lastUpdate = time.Now()
	g.hud = NewHUD(g)
	return g
}

// Update updates the game state. This method is called every frame.
func (g *Game) Update() error {
	now := time.Now()
	dt := 1.0 / 60.0
	if !g.lastUpdate.IsZero() {
		dt = now.Sub(g.lastUpdate).Seconds()
	}
	g.lastUpdate = now
	g.input.Update()
	g.handleStatsPanelInput()
	g.handleSkillMenuInput()
	if g.skillMenuOpen {
		return nil
	}

	if !g.commandMode && g.input.Command() {
		g.commandMode = true
		g.commandBuffer = ""
		return nil
	}
	if g.commandMode {
		for _, r := range g.input.TypedChars() {
			if unicode.IsPrint(r) {
				g.commandBuffer += string(r)
			}
		}
		if g.input.Backspace() && len(g.commandBuffer) > 0 {
			g.commandBuffer = g.commandBuffer[:len(g.commandBuffer)-1]
		}
		if g.input.Enter() {
			g.executeCommand(strings.TrimSpace(g.commandBuffer))
			g.commandMode = false
			g.commandBuffer = ""
		}
		return nil
	}

	if g.input.Quit() {
		return ebiten.Termination
	}

	if g.phase == PhaseMainMenu {
		return g.mainMenu.Update(g, dt)
	}
	if g.phase == PhasePreGame {
		return g.preGame.Update(g, dt)
	}
	if g.phase == PhasePaused {
		return g.updatePaused(dt)
	}
	if g.phase == PhaseSettings {
		return g.updateSettings(dt)
	}

	// Animate conveyor belt offset
	if g.conveyorOffset > 0 {
		shift := conveyorSpeed * dt
		if shift > g.conveyorOffset {
			shift = g.conveyorOffset
		}
		g.conveyorOffset -= shift
	}

	if g.flashTimer > 0 {
		g.flashTimer -= dt
		if g.flashTimer < 0 {
			g.flashTimer = 0
		}
	}

	if g.upgradeMenuOpen {
		const (
			upgradeDamageCost    = 5
			upgradeRangeCost     = 5
			upgradeFireRateCost  = 5
			upgradeAmmoCost      = 10
			upgradeForesightCost = 5
			optionsCount         = 5
		)
		if g.input.Down() {
			g.upgradeCursor = (g.upgradeCursor + 1) % (optionsCount + 1)
		}
		if g.input.Up() {
			g.upgradeCursor = (g.upgradeCursor - 1 + optionsCount + 1) % (optionsCount + 1)
		}
		if len(g.towers) > 0 {
			tower := g.towers[g.selectedTower]
			purchase := func(opt int) bool {
				switch opt {
				case 0:
					if g.SpendGold(upgradeDamageCost) {
						tower.damage++
						return true
					}
				case 1:
					if g.SpendGold(upgradeRangeCost) {
						tower.rangeDst += 50
						tower.rangeImg = generateRangeImage(tower.rangeDst)
						return true
					}
				case 2:
					if g.SpendGold(upgradeFireRateCost) {
						if tower.rate > 10 {
							tower.rate -= 10
						}
						return true
					}
				case 3:
					if g.SpendGold(upgradeAmmoCost) {
						tower.UpgradeAmmoCapacity(2)
						return true
					}
				case 4:
					if g.SpendGold(upgradeForesightCost) {
						tower.UpgradeForesight(2)
						return true
					}
				}
				return false
			}

			if inpututil.IsKeyJustPressed(ebiten.Key1) {
				purchase(0)
			}
			if inpututil.IsKeyJustPressed(ebiten.Key2) {
				purchase(1)
			}
			if inpututil.IsKeyJustPressed(ebiten.Key3) {
				purchase(2)
			}
			if inpututil.IsKeyJustPressed(ebiten.Key4) {
				purchase(3)
			}
			if inpututil.IsKeyJustPressed(ebiten.Key5) {
				purchase(4)
			}

			if g.input.Enter() {
				if g.upgradeCursor < optionsCount {
					purchase(g.upgradeCursor)
				} else {
					g.upgradeMenuOpen = false
					g.upgradeCursor = 0
				}
			}
		}
		if g.input.SelectTower() {
			g.upgradeMenuOpen = false
		}
		return nil
	}

	if g.towerSelectMode {
		for _, r := range g.input.TypedChars() {
			label := strings.ToLower(string(r))
			if idx, ok := g.towerLabels[label]; ok {
				g.selectedTower = idx
				g.towerSelectMode = false
				g.upgradeMenuOpen = true
				g.upgradeCursor = 0
				break
			}
		}
		if g.input.SelectTower() {
			g.towerSelectMode = false
		}
		return nil
	}

	if !g.upgradeMenuOpen && !g.buildMenuOpen && !g.shopOpen && g.input.SelectTower() {
		g.enterTowerSelectMode()
		return nil
	}

	if g.buildMenuOpen {
		const optionsCount = 4
		if g.input.Down() {
			g.buildCursor = (g.buildCursor + 1) % optionsCount
		}
		if g.input.Up() {
			g.buildCursor = (g.buildCursor - 1 + optionsCount) % optionsCount
		}
		if inpututil.IsKeyJustPressed(ebiten.Key1) {
			g.buildTowerAtCursorType(TowerBasic)
			g.buildMenuOpen = false
		}
		if inpututil.IsKeyJustPressed(ebiten.Key2) {
			g.buildTowerAtCursorType(TowerSniper)
			g.buildMenuOpen = false
		}
		if inpututil.IsKeyJustPressed(ebiten.Key3) {
			g.buildTowerAtCursorType(TowerRapid)
			g.buildMenuOpen = false
		}
		if g.input.Enter() {
			switch g.buildCursor {
			case 0:
				g.buildTowerAtCursorType(TowerBasic)
			case 1:
				g.buildTowerAtCursorType(TowerSniper)
			case 2:
				g.buildTowerAtCursorType(TowerRapid)
			}
			g.buildMenuOpen = false
		}
		if g.input.Build() {
			g.buildMenuOpen = false
		}
		return nil
	}

	if !g.shopOpen {
		if g.input.Left() {
			g.cursorX--
		}
		if g.input.Right() {
			g.cursorX++
		}
		if g.input.Up() {
			g.cursorY--
		}
		if g.input.Down() {
			g.cursorY++
		}
		if g.cursorX < 0 {
			g.cursorX = 0
		}
		if g.cursorX > 59 {
			g.cursorX = 59
		}
		if g.cursorY < 0 {
			g.cursorY = 0
		}
		if g.cursorY > 33 {
			g.cursorY = 33
		}
		if g.input.Build() {
			g.buildMenuOpen = true
			g.buildCursor = 0
		}
	}

	if g.input.Reload() {
		if err := g.reloadConfig(ConfigFile); err != nil {
			fmt.Println("reload config:", err)
		}
	}

	g.handleSlotMenuInput()
	if g.slotMenuOpen {
		return nil
	}

	if g.input.Space() && g.phase == PhasePlaying {
		g.phase = PhasePaused
		g.paused = true
		g.pauseCursor = 0
	}

	if g.input.Quit() {
		return ebiten.Termination
	}

	if g.phase == PhaseGameOver {
		return nil
	}

	// ---- Global typing queue processing (letter by letter) ----
	if g.queue != nil {
		g.queue.Update(dt)
		if _, ok := g.queue.Peek(); ok {
			if g.queueJam {
				if g.input.Backspace() {
					g.queueJam = false
					g.queue.ResetProgress()
				}
			} else {
				for _, r := range g.input.TypedChars() {
					match, done, dq := g.queue.TryLetter(r)
					if match {
						g.conveyorOffset += letterWidth
						g.typing.Record(true)

						if g.currentWord.Text == "" {
							g.currentWord.Text = dq.Text
						}
						g.currentWord.Start()
						g.currentWord.Correct++

						if done {
							g.currentWord.Finish()
							g.wordHistory = append(g.wordHistory, g.currentWord)
							g.currentWord = WordStat{}

							switch dq.Source {
							case "Farmer":
								g.farmer.OnWordCompleted(dq.Text, &g.resources)
							case "Barracks":
								if unit := g.barracks.OnWordCompleted(dq.Text); unit != nil {
									g.military.AddUnit(unit)
								}
							}
						}
					} else {
						g.typing.Record(false)
						g.currentWord.Incorrect++
						g.MistypeFeedback()
						g.queueJam = true
					}
					break
				}
			}
		}
	}

	if len(g.towers) > 0 {
		if g.input.Down() {
			g.selectedTower = (g.selectedTower + 1) % len(g.towers)
		}
		if g.input.Up() {
			g.selectedTower = (g.selectedTower - 1 + len(g.towers)) % len(g.towers)
		}
	}

	if g.shopOpen {
		// Upgrade options and costs
		const (
			upgradeDamageCost    = 5
			upgradeRangeCost     = 5
			upgradeFireRateCost  = 5
			upgradeAmmoCost      = 10
			upgradeForesightCost = 5
			optionsCount         = 8
		)

		if g.input.Down() {
			g.shopCursor = (g.shopCursor + 1) % optionsCount
		}
		if g.input.Up() {
			g.shopCursor = (g.shopCursor - 1 + optionsCount) % optionsCount
		}

		if len(g.towers) > 0 {
			tower := g.towers[g.selectedTower]

			purchase := func(opt int) bool {
				switch opt {
				case 0:
					if g.SpendGold(upgradeDamageCost) {
						tower.damage++
						return true
					}
				case 1:
					if g.SpendGold(upgradeRangeCost) {
						tower.rangeDst += 50
						tower.rangeImg = generateRangeImage(tower.rangeDst)
						return true
					}
				case 2:
					if g.SpendGold(upgradeFireRateCost) {
						if tower.rate > 10 {
							tower.rate -= 10
						}
						return true
					}
				case 3:
					if g.SpendGold(upgradeAmmoCost) {
						tower.UpgradeAmmoCapacity(2)
						return true
					}
				case 4:
					if g.SpendGold(upgradeForesightCost) {
						tower.UpgradeForesight(2)
						return true
					}
				case 5:
					if g.farmer != nil {
						return g.farmer.UnlockNext(&g.resources)
					}
				case 6:
					if g.barracks != nil {
						return g.barracks.UnlockNext(&g.resources)
					}
				}
				return false
			}

			// Direct number keys
			if inpututil.IsKeyJustPressed(ebiten.Key1) {
				purchase(0)
			}
			if inpututil.IsKeyJustPressed(ebiten.Key2) {
				purchase(1)
			}
			if inpututil.IsKeyJustPressed(ebiten.Key3) {
				purchase(2)
			}
			if inpututil.IsKeyJustPressed(ebiten.Key4) {
				purchase(3)
			}
			if inpututil.IsKeyJustPressed(ebiten.Key5) {
				purchase(4)
			}
			if inpututil.IsKeyJustPressed(ebiten.Key6) {
				purchase(5)
			}
			if inpututil.IsKeyJustPressed(ebiten.Key7) {
				purchase(6)
			}

			if g.input.Enter() {
				if g.shopCursor < 7 {
					purchase(g.shopCursor)
				} else {
					g.shopOpen = false
					g.shopCursor = 0
					g.currentWave++
					g.startWave()
				}
			}
		}
		return nil
	}

	// Slow down mob spawning significantly
	if g.mobsToSpawn > 0 {
		g.spawnTicker += dt
		if g.spawnTicker >= g.spawnInterval {
			g.spawnTicker = 0
			g.spawnMob()
			g.mobsToSpawn--
		}
	} else if len(g.mobs) == 0 {
		if !g.shopOpen {
			g.shopOpen = true
			if g.lastWaveSaved != g.currentWave {
				g.saveGame(g.currentSavePath())
				g.lastWaveSaved = g.currentWave
			}
		}
	}

	// Update buildings and units
	if g.military != nil {
		// Combat resolution currently only handles OrcGrunts; none are
		// spawned yet so pass nil.
		g.military.Update(dt, nil)
	}
	if g.farmer != nil {
		if w := g.farmer.Update(dt); w != "" {
			g.farmer.OnWordCompleted(w, &g.resources)
		}
	}
	if g.lumberjack != nil {
		if w := g.lumberjack.Update(dt); w != "" {
			g.lumberjack.OnWordCompleted(w, &g.resources)
		}
	}
	if g.miner != nil {
		if w := g.miner.Update(dt); w != "" {
			g.miner.OnWordCompleted(w, &g.resources)
		}
	}
	if g.barracks != nil {
		if w := g.barracks.Update(dt); w != "" {
			g.barracks.OnWordCompleted(w)
		}
	}

	for _, t := range g.towers {
		t.Update(dt)
	}

	g.base.Update(dt)

	for i := 0; i < len(g.projectiles); {
		p := g.projectiles[i]
		p.Update(dt)
		if !p.alive {
			g.projectiles = append(g.projectiles[:i], g.projectiles[i+1:]...)
			continue
		}
		i++
	}

	for i := 0; i < len(g.mobs); {
		m := g.mobs[i]
		m.Update(dt)
		bx, by, bw, bh := g.base.Bounds()
		mx, my := m.Position()
		_, _, mw, _ := m.Bounds()
		dx := mx - float64(bx+bw/2)
		dy := my - float64(by+bh/2)
		if math.Hypot(dx, dy) < float64(mw/2+bw/2) {
			g.base.Damage(1)
			m.Damage(mw) // force kill
		}
		if !m.Alive() {
			g.mobs = append(g.mobs[:i], g.mobs[i+1:]...)
			mult := g.typing.ScoreMultiplier()
			reward := int(mult)
			if reward < 1 {
				reward = 1
			}
			g.AddGold(reward)
			g.score += reward
			continue
		}
		i++
	}

	if !g.base.Alive() {
		g.gameOver = true
		g.phase = PhaseGameOver
	}

	if g.gameOver && !g.gameOverHandled {
		if g.history != nil {
			g.history.Record(g.typing)
		}
		g.evaluatePerformanceAchievements()
		g.gameOverHandled = true
	}

	return nil
}

// Step advances the game state by dt seconds.
// This is used for testing conveyor animation.
func (g *Game) Step(dt float64) error {
	// For test: just call Update() after setting lastUpdate to simulate dt.
	g.lastUpdate = g.lastUpdate.Add(-time.Duration(dt * float64(time.Second)))
	return g.Update()
}

func (g *Game) updatePaused(dt float64) error {
	const optionsCount = 4
	if g.input.Down() {
		g.pauseCursor = (g.pauseCursor + 1) % optionsCount
	}
	if g.input.Up() {
		g.pauseCursor = (g.pauseCursor - 1 + optionsCount) % optionsCount
	}
	if g.input.Enter() {
		switch g.pauseCursor {
		case 0:
			g.phase = PhasePlaying
			g.paused = false
		case 1:
			g.Restart()
		case 2:
			return ebiten.Termination
		case 3:
			g.phase = PhaseSettings
			g.prevPhase = PhasePaused
			g.settingsCursor = 0
		}
	}
	return nil
}

func (g *Game) updateSettings(dt float64) error {
	if g.input.Down() {
		g.settingsCursor = (g.settingsCursor + 1) % 2
	}
	if g.input.Up() {
		g.settingsCursor = (g.settingsCursor - 1 + 2) % 2
	}
	if g.input.Enter() {
		switch g.settingsCursor {
		case 0:
			g.settings.Mute = !g.settings.Mute
			if g.sound != nil {
				g.sound.ToggleMute()
			}
		case 1:
			g.phase = g.prevPhase
			if g.prevPhase == PhasePaused {
				g.paused = true
			}
		}
	}
	return nil
}

// Draw renders the game to the screen. This method is called every frame.
func (g *Game) Draw(screen *ebiten.Image) {
	g.screen.Clear()
	if g.phase == PhaseMainMenu {
		g.mainMenu.Draw(g, g.screen)
		g.renderFrame(screen)
		return
	}
	if g.phase == PhasePreGame {
		g.preGame.Draw(g, g.screen)
		g.renderFrame(screen)
		return
	}
	drawBackgroundTilemap(g.screen)

	if g.phase == PhaseGameOver {
		opts := &text.DrawOptions{}
		opts.GeoM.Translate(900, 540)
		opts.ColorScale.ScaleWithColor(color.White)
		text.Draw(g.screen, "Game Over", BoldFont, opts)
		summary := []string{
			fmt.Sprintf("Score: %d", g.score),
			fmt.Sprintf("Accuracy: %.0f%%", g.typing.Accuracy()*100),
			fmt.Sprintf("WPM: %.1f", g.EffectiveWPM()),
		}
		if g.history != nil {
			summary = append(summary,
				fmt.Sprintf("Best Accuracy: %.0f%%", g.history.BestAccuracy*100),
				fmt.Sprintf("Best WPM: %.1f", g.history.BestWPM))
		}
		if len(g.achievements) > 0 {
			summary = append(summary, "Achievements:")
			for _, a := range g.achievements {
				summary = append(summary, " - "+a)
			}
		}
		drawMenu(g.screen, summary, 820, 580)
		g.renderFrame(screen)
		return
	}

	if g.shopOpen {
		// The main shop interface is now drawn by the HUD.
		// We can keep a minimal centered message or remove it entirely.
		// For now, let's remove the old one:
		// ebitenutil.DebugPrintAt(g.screen, "-- SHOP -- press Enter", 850, 520)
		// The HUD will display shop details.
	}

	g.base.Draw(g.screen)

	for i, t := range g.towers {
		t.Draw(g.screen)
		if i == g.selectedTower {
			bx, by, bw, bh := t.Bounds()
			vector.StrokeRect(g.screen, float32(bx-2), float32(by-2), float32(bw+4), float32(bh+4), 2, color.RGBA{255, 0, 0, 200}, false)
		}
	}

	if g.towerSelectMode {
		// Dim the background to highlight tower labels
		vector.DrawFilledRect(g.screen, 0, 0, 1920, 1080, color.RGBA{0, 0, 0, 120}, false)
		for i, t := range g.towers {
			label := ""
			for k, idx := range g.towerLabels {
				if idx == i {
					label = k
					break
				}
			}
			if label != "" {
				x, y := t.Position()
				opts := &text.DrawOptions{}
				opts.GeoM.Translate(x-6, y-30)
				opts.ColorScale.ScaleWithColor(color.White)
				text.Draw(g.screen, label, BoldFont, opts)
			}
		}
	}
	for _, p := range g.projectiles {
		p.Draw(g.screen)
	}
	for _, m := range g.mobs {
		m.Draw(g.screen)
	}
	if g.military != nil {
		for _, u := range g.military.Units() {
			u.Draw(g.screen)
		}
	}

	if !g.shopOpen {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(g.cursorX*TileSize), float64(TopMargin+g.cursorY*TileSize))
		if g.validTowerPosition(g.cursorX, g.cursorY) {
			g.screen.DrawImage(ImgHighlightTile, op)
		} else {
			// draw red rectangle for invalid position
			vector.DrawFilledRect(g.screen, float32(g.cursorX*TileSize), float32(TopMargin+g.cursorY*TileSize), float32(TileSize), float32(TileSize), color.RGBA{255, 0, 0, 100}, false)
		}
	}

	if g.phase == PhasePaused {
		lines := []string{"-- PAUSED --"}
		opts := []string{"Resume", "Restart", "Quit", "Settings"}
		for i, opt := range opts {
			prefix := "  "
			if i == g.pauseCursor {
				prefix = "> "
			}
			lines = append(lines, prefix+opt)
		}
		drawMenu(g.screen, lines, 860, 480)
		g.renderFrame(screen)
		return
	}
	if g.phase == PhaseSettings {
		lines := []string{"-- SETTINGS --"}
		mute := "Off"
		if g.settings.Mute {
			mute = "On"
		}
		opts := []string{"Toggle Mute: " + mute, "Back"}
		for i, opt := range opts {
			prefix := "  "
			if i == g.settingsCursor {
				prefix = "> "
			}
			lines = append(lines, prefix+opt)
		}
		drawMenu(g.screen, lines, 860, 480)
		g.renderFrame(screen)
		return
	}

	if g.hud != nil {
		g.hud.Draw(g.screen)
	}

	highlightHoverAndClickAndDrag(g.screen, "line")

	if g.flashTimer > 0 {
		alpha := uint8(255 * (g.flashTimer / jamFlashDuration))
		vector.DrawFilledRect(g.screen, 0, 0, 1920, 1080, color.RGBA{255, 0, 0, alpha}, false)
	}

	g.renderFrame(screen)
}

// renderFrame scales and draws the internal screen buffer to the actual screen
func (g *Game) renderFrame(screen *ebiten.Image) {
	screen.Clear()
	w, h := screen.Bounds().Dx(), screen.Bounds().Dy()
	scaleX := float64(w) / 1920.0
	scaleY := float64(h) / 1080.0
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(scaleX, scaleY)
	screen.DrawImage(g.screen, opts)
}

// Layout returns the size of the game screen in pixels.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 1920, 1080
}

// drawBackgroundTilemap draws the background tilemap on the screen.
func drawBackgroundTilemap(screen *ebiten.Image) {
	screen.DrawImage(ImgBackgroundBasicTiles, nil)
}

// spawnMob adds a new mob at the right side.
func (g *Game) spawnMob() {
	row := rand.Intn(32)
	x, y := tilePosition(59, row)
	hp := g.cfg.MobBaseHealth
	if hp == 0 {
		hp = 1
	}
	if g.cfg != nil {
		hp = int(float64(hp) + float64(g.currentWave-1)*g.cfg.N)
		if hp < 1 {
			hp = 1
		}
	}
	speed := g.cfg.MobSpeed * 0.3 // Much slower mobs
	if speed == 0 {
		speed = DefaultConfig.MobSpeed * 0.3
	}
	var m Enemy
	if g.currentWave%5 == 0 && g.mobsToSpawn == 1 {
		m = NewBossMob(float64(x+16), float64(y+16), g.base, hp*5, speed*0.5)
	} else {
		switch rand.Intn(3) {
		case 0:
			m = NewMob(float64(x+16), float64(y+16), g.base, hp, speed)
		case 1:
			m = NewArmoredMob(float64(x+16), float64(y+16), g.base, hp, 2, speed)
		default:
			m = NewFastMob(float64(x+16), float64(y+16), g.base, hp, speed, 2)
		}
	}
	g.mobs = append(g.mobs, m)
}

func (g *Game) validTowerPosition(tileX, tileY int) bool {
	if tileX < 0 || tileX > 59 || tileY < 0 || tileY > 33 {
		return false
	}
	tx, ty := tilePosition(tileX, tileY)
	px := float64(tx + TileSize/2)
	py := float64(ty + TileSize/2)
	bx, by, bw, bh := g.base.Bounds()
	if int(px) >= bx && int(px) <= bx+bw && int(py) >= by && int(py) <= by+bh {
		return false
	}
	for _, t := range g.towers {
		x, y, w, h := t.Bounds()
		if int(px) >= x && int(px) <= x+w && int(py) >= y && int(py) <= y+h {
			return false
		}
	}
	return true
}

func (g *Game) buildTowerAtCursorType(tt TowerType) {
	if g.cfg == nil {
		return
	}
	cost := g.cfg.TowerConstructionCost
	if cost == 0 {
		cost = DefaultConfig.TowerConstructionCost
	}
	if g.Gold() < cost {
		return
	}
	if !g.validTowerPosition(g.cursorX, g.cursorY) {
		return
	}
	tx, ty := tilePosition(g.cursorX, g.cursorY)
	t := NewTowerWithType(g, float64(tx+TileSize/2), float64(ty+TileSize/2), tt)
	t.ApplyModifiers(g.towerMods)
	g.towers = append(g.towers, t)
	g.SpendGold(cost)
}

// applyNextTech unlocks the next tech node and applies its effects.
func (g *Game) applyNextTech() {
	if g.techTree == nil || g.techTree.Completed() {
		return
	}
	letters, ach, mods := g.techTree.UnlockNext()
	if len(letters) > 0 {
		existing := make(map[rune]struct{})
		for _, r := range g.letterPool {
			existing[r] = struct{}{}
		}
		for _, r := range letters {
			if _, ok := existing[r]; !ok {
				g.letterPool = append(g.letterPool, r)
			}
		}
	}
	if mods != (TowerModifiers{}) {
		g.towerMods = g.towerMods.Merge(mods)
		for _, t := range g.towers {
			t.ApplyModifiers(mods)
		}
	}
	if ach != "" {
		g.achievements = append(g.achievements, ach)
	}
}

// applySkillEffects applies the effects of a newly unlocked skill node.
func (g *Game) applySkillEffects(n *SkillNode) {
	for k, v := range n.Effects {
		switch k {
		case "damage_mult":
			mod := TowerModifiers{DamageMult: v}
			g.towerMods = g.towerMods.Merge(mod)
			for _, t := range g.towers {
				t.ApplyModifiers(mod)
			}
		case "fire_rate_mult":
			mod := TowerModifiers{FireRateMult: v}
			g.towerMods = g.towerMods.Merge(mod)
			for _, t := range g.towers {
				t.ApplyModifiers(mod)
			}
		case "hp_add":
			if g.base != nil {
				g.base.health += int(v)
			}
		case "wpm_bonus":
			g.wpmBonus += int(v)
		case "auto_collect":
			g.autoCollect = true
		case "hotkeys":
			g.hotkeys = true
		}
	}
}

// filteredTechNodes returns remaining tech nodes matching the search buffer.
func (g *Game) filteredTechNodes() []TechNode {
	if g.techTree == nil {
		return nil
	}

	if g.input.TechMenu() {
		g.techMenuOpen = !g.techMenuOpen
		if g.techMenuOpen {
			g.searchBuffer = ""
			g.techCursor = 0
		}
		return nil
	}

	if g.techMenuOpen {
		for _, r := range g.input.TypedChars() {
			if unicode.IsPrint(r) {
				g.searchBuffer += string(r)
			}
		}
		if g.input.Backspace() && len(g.searchBuffer) > 0 {
			g.searchBuffer = g.searchBuffer[:len(g.searchBuffer)-1]
		}
		nodes := g.filteredTechNodes()
		if len(nodes) > 0 {
			if g.input.Down() {
				g.techCursor = (g.techCursor + 1) % len(nodes)
			}
			if g.input.Up() {
				g.techCursor = (g.techCursor - 1 + len(nodes)) % len(nodes)
			}
			if g.input.Enter() {
				node := nodes[g.techCursor]
				if g.techTree.stage < len(g.techTree.nodes) && node.Name == g.techTree.nodes[g.techTree.stage].Name {
					g.applyNextTech()
					g.techMenuOpen = false
				}
			}
		}
		return nil
	}
	var out []TechNode
	term := strings.ToLower(g.searchBuffer)
	for i := g.techTree.stage; i < len(g.techTree.nodes); i++ {
		n := g.techTree.nodes[i]
		if term == "" || strings.Contains(strings.ToLower(n.Name), term) {
			out = append(out, n)
		}
	}
	return out
}

// startWave initializes spawn counters for the next wave.
func (g *Game) startWave() {
	g.spawnTicker = 0
	base := g.cfg.MobsPerWave
	if base == 0 {
		base = DefaultConfig.MobsPerWave
	}
	inc := g.cfg.MobsPerWaveInc
	g.mobsToSpawn = base + inc*(g.currentWave-1)
	g.spawnInterval = g.cfg.SpawnInterval * 6.0 // Much slower spawning

	g.applyNextTech()
}

// randomReloadLetter returns a random letter from the current letter pool.
// If no letters have been unlocked, 'f' is returned as a safe default.
func (g *Game) randomReloadLetter() rune {
	if len(g.letterPool) == 0 {
		return 'f'
	}
	return g.letterPool[rand.Intn(len(g.letterPool))]
}

// skillNodesByCategory returns all skill nodes for the given category.
func (g *Game) skillNodesByCategory(cat SkillCategory) []*SkillNode {
	if g.skillTree == nil {
		return nil
	}
	var out []*SkillNode
	for _, n := range g.skillTree.Nodes {
		if n.Category == cat {
			out = append(out, n)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].ID < out[j].ID })
	return out
}

// handleStatsPanelInput toggles the stats panel visibility when Tab is pressed.
func (g *Game) handleStatsPanelInput() {
	if g.input.StatsPanel() {
		g.statsPanelOpen = !g.statsPanelOpen
	}
}

// handleSkillMenuInput processes keyboard input for the skill tree menu.
func (g *Game) handleSkillMenuInput() {
	if g.skillTree == nil {
		return
	}
	if g.input.SkillMenu() {
		g.skillMenuOpen = !g.skillMenuOpen
		if g.skillMenuOpen {
			g.skillCategory = 0
			g.skillCursor = 0
		}
		return
	}
	if !g.skillMenuOpen {
		return
	}
	categories := []SkillCategory{SkillOffense, SkillDefense, SkillTyping, SkillAutomation, SkillUtility}
	if g.input.Right() {
		g.skillCategory = (g.skillCategory + 1) % SkillCategory(len(categories))
		g.skillCursor = 0
		return
	}
	if g.input.Left() {
		g.skillCategory = (g.skillCategory - 1 + SkillCategory(len(categories))) % SkillCategory(len(categories))
		g.skillCursor = 0
		return
	}
	nodes := g.skillNodesByCategory(categories[g.skillCategory])
	if len(nodes) == 0 {
		return
	}
	if g.input.Down() {
		g.skillCursor = (g.skillCursor + 1) % len(nodes)
	}
	if g.input.Up() {
		g.skillCursor = (g.skillCursor - 1 + len(nodes)) % len(nodes)
	}
	if g.input.Enter() {
		node := nodes[g.skillCursor]
		if g.skillTree.Unlock(node.ID, &g.resources) {
			g.unlockedSkills[node.ID] = true
			g.applySkillEffects(node)
		}
	}
}

// handleSlotMenuInput manages the save/load slot selection overlay.
func (g *Game) handleSlotMenuInput() {
	if g.input.Save() {
		g.slotMenuOpen = true
		g.slotModeSave = true
		g.slotCursor = 0
	}
	if g.input.Load() {
		g.slotMenuOpen = true
		g.slotModeSave = false
		g.slotCursor = 0
	}
	if !g.slotMenuOpen {
		return
	}
	if g.input.Down() {
		g.slotCursor = (g.slotCursor + 1) % 3
	}
	if g.input.Up() {
		g.slotCursor = (g.slotCursor - 1 + 3) % 3
	}
	if g.input.Enter() {
		g.SetSaveSlot(g.slotCursor + 1)
		path := g.currentSavePath()
		if g.slotModeSave {
			g.saveGame(path)
			g.lastWaveSaved = g.currentWave
		} else {
			if err := g.loadGame(path); err != nil {
				fmt.Println("load game:", err)
			}
		}
		g.slotMenuOpen = false
	}
}

// enterTowerSelectMode assigns letter labels to towers and activates selection mode.
func (g *Game) enterTowerSelectMode() {
	g.towerLabels = make(map[string]int)
	letters := "abcdefghijklmnopqrstuvwxyz"
	for i := range g.towers {
		if i >= len(letters) {
			break
		}
		label := string(letters[i])
		g.towerLabels[label] = i
	}
	g.towerSelectMode = true
}

// MistypeFeedback triggers a red flash and "clank" sound for an incorrect key press.
// The flash duration is controlled by jamFlashDuration.
func (g *Game) MistypeFeedback() {
	if g == nil {
		return
	}
	g.flashTimer = jamFlashDuration
	if g.sound != nil {
		g.sound.PlayClank()
	}
}

// highlightHoverAndClickAndDrag highlights the tile under the mouse cursor.
func highlightHoverAndClickAndDrag(screen *ebiten.Image, shape string) {
	mouseX, mouseY := ebiten.CursorPosition()
	if mouseX < 0 || mouseY < TopMargin || mouseX >= 1920 || mouseY >= 1080-TopMargin {
		return
	}
	tileX, tileY := tileAtPosition(mouseX, mouseY)

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		if tileX >= 0 && tileX <= 59 && tileY >= 0 && tileY <= 33 {
			id := fmt.Sprintf("%d,%d", tileX, tileY)
			houses[id] = struct{}{}
		}
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		if !mousePressed {
			mousePressed = true
			clickedTileX = tileX
			clickedTileY = tileY
		}
	} else {
		mousePressed = false
	}

	if mousePressed {
		minX, maxX := clickedTileX, tileX
		if minX > maxX {
			minX, maxX = maxX, minX
		}
		minY, maxY := clickedTileY, tileY
		if minY > maxY {
			minY, maxY = maxY, minY
		}
		if minX < 0 {
			minX = 0
		}
		if maxX > 59 {
			maxX = 59
		}
		if minY < 0 {
			minY = 0
		}
		if maxY > 33 {
			maxY = 33
		}
		switch shape {
		case "rectangle":
			for x := minX; x <= maxX; x++ {
				for y := minY; y <= maxY; y++ {
					op := &ebiten.DrawImageOptions{}
					op.GeoM.Translate(float64(x*32), float64(TopMargin+y*32))
					screen.DrawImage(ImgHighlightTile, op)
				}
			}
		case "circle":
			centerTileX := (minX + maxX) / 2
			centerTileY := (minY + maxY) / 2
			radius := (maxX - minX) / 2
			for x := centerTileX - radius; x <= centerTileX+radius; x++ {
				for y := centerTileY - radius; y <= centerTileY+radius; y++ {
					if x < 0 || x > 59 || y < 0 || y > 31 {
						continue
					}
					dx := x - centerTileX
					dy := y - centerTileY
					if dx*dx+dy*dy <= radius*radius {
						op := &ebiten.DrawImageOptions{}
						op.GeoM.Translate(float64(x*32), float64(TopMargin+y*32))
						screen.DrawImage(ImgHighlightTile, op)
					}
				}
			}
		case "line":
			x0, y0 := clickedTileX, clickedTileY
			x1, y1 := tileX, tileY
			dx := math.Abs(float64(x1 - x0))
			dy := math.Abs(float64(y1 - y0))
			sx := -1
			if x0 < x1 {
				sx = 1
			}
			sy := -1
			if y0 < y1 {
				sy = 1
			}
			err := dx - dy
			for {
				if x0 >= 0 && x0 <= 59 && y0 >= 0 && y0 <= 33 {
					op := &ebiten.DrawImageOptions{}
					op.GeoM.Translate(float64(x0*32), float64(TopMargin+y0*32))
					screen.DrawImage(ImgHighlightTile, op)
				}
				if x0 == x1 && y0 == y1 {
					break
				}
				e2 := 2 * err
				if e2 > -dy {
					err -= dy
					x0 += sx
				}
				if e2 < dx {
					err += dx
					y0 += sy
				}
				if x0 < 0 || x0 > 59 || y0 < 0 || y0 > 33 {
					break
				}
			}
		default:
			for x := minX; x <= maxX; x++ {
				for y := minY; y <= maxY; y++ {
					op := &ebiten.DrawImageOptions{}
					op.GeoM.Translate(float64(x*32), float64(TopMargin+y*32))
					screen.DrawImage(ImgHighlightTile, op)
				}
			}
		}
	} else {
		if tileX >= 0 && tileX <= 59 && tileY >= 0 && tileY <= 33 {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(tileX*32), float64(TopMargin+tileY*32))
			screen.DrawImage(ImgHighlightTile, op)
		}
	}

	for id := range houses {
		var houseTileX, houseTileY int
		fmt.Sscanf(id, "%d,%d", &houseTileX, &houseTileY)
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(houseTileX*32), float64(TopMargin+houseTileY*32))
		screen.DrawImage(ImgHouseTile, op)
	}

	// Placeholder for additional debug UI if needed in the future.
}

// reloadConfig loads a Config from the given file and applies it to the game.
func (g *Game) reloadConfig(path string) error {
	cfg, err := LoadConfig(path)
	if err != nil {
		// still apply defaults when loading fails
		g.cfg = &cfg
		return err
	}
	g.cfg = &cfg

	// apply updated parameters
	hp := cfg.BaseHealth
	if hp == 0 {
		hp = int(cfg.J)
	}
	if hp <= 0 {
		hp = BaseStartingHealth
	}
	g.base.health = hp
	for _, t := range g.towers {
		t.ApplyConfig(cfg)
	}
	if cfg.SpawnInterval > 0 {
		g.spawnInterval = cfg.SpawnInterval
	}
	base := cfg.MobsPerWave
	if base == 0 {
		base = DefaultConfig.MobsPerWave
	}
	g.mobsToSpawn = base + cfg.MobsPerWaveInc*(g.currentWave-1)
	return nil
}

func (g *Game) saveGame(path string) {
	sg := savedGame{
		Version:  SaveVersion,
		Gold:     g.resources.GoldAmount(),
		Food:     g.resources.FoodAmount(),
		Wave:     g.currentWave,
		BaseHP:   g.base.Health(),
		Settings: g.settings,
		Skills:   make([]string, 0, len(g.unlockedSkills)),
	}
	for _, t := range g.towers {
		sg.Towers = append(sg.Towers, savedTower{
			X:            t.pos.X,
			Y:            t.pos.Y,
			Damage:       t.damage,
			Range:        t.rangeDst,
			Rate:         t.rate,
			AmmoCapacity: t.ammoCapacity,
		})
	}
	for id := range g.unlockedSkills {
		sg.Skills = append(sg.Skills, id)
	}
	b, err := json.MarshalIndent(sg, "", "  ")
	if err == nil {
		_ = os.WriteFile(path, b, 0644)
	}
}

func (g *Game) loadGame(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	var sg savedGame
	if err := json.Unmarshal(data, &sg); err != nil {
		return err
	}
	if sg.Version != SaveVersion {
		return ErrSaveVersion
	}
	*g = *NewGameWithConfig(*g.cfg)
	g.resources.Gold.Set(sg.Gold)
	g.resources.Food.Set(sg.Food)
	g.currentWave = sg.Wave
	g.base.health = sg.BaseHP
	g.settings = sg.Settings
	if g.sound != nil && g.settings.Mute {
		g.sound.mute = true
	}
	g.towers = nil
	for _, st := range sg.Towers {
		t := NewTower(g, st.X, st.Y)
		t.damage = st.Damage
		t.rangeDst = st.Range
		t.rangeImg = generateRangeImage(t.rangeDst)
		t.rate = st.Rate
		t.ammoCapacity = st.AmmoCapacity
		t.ammoQueue = make([]bool, t.ammoCapacity)
		for i := range t.ammoQueue {
			t.ammoQueue[i] = true
		}
		g.towers = append(g.towers, t)
	}
	for _, id := range sg.Skills {
		if node, ok := g.skillTree.Nodes[id]; ok {
			g.skillTree.unlocked[id] = true
			g.unlockedSkills[id] = true
			g.applySkillEffects(node)
		}
	}
	return nil
}

func (g *Game) Restart() {
	hist := g.history
	*g = *NewGameWithHistory(*g.cfg, hist)
}

// executeCommand runs a textual command entered via command mode.
func (g *Game) executeCommand(cmd string) {
	switch strings.ToLower(cmd) {
	case "quit":
		g.quit = true
	case "pause":
		g.paused = true
		g.phase = PhasePaused
	case "resume":
		g.paused = false
		g.phase = PhasePlaying
	}
}

// evaluatePerformanceAchievements awards achievements and gold based on typing performance.
func (g *Game) evaluatePerformanceAchievements() {
	wpm := g.EffectiveWPM()
	acc := g.typing.Accuracy()

	add := func(name string) {
		for _, a := range g.achievements {
			if a == name {
				return
			}
		}
		g.achievements = append(g.achievements, name)
	}

	if wpm >= 60 {
		add("Speed Demon")
		g.AddGold(5)
	}
	if acc >= 0.95 {
		add("Sharpshooter")
		g.AddGold(5)
	}
	if g.typing.MaxCombo() >= 10 {
		add("Combo Master")
		g.AddGold(5)
	}
}

// EffectiveWPM returns the player's WPM including any skill bonuses.
func (g *Game) EffectiveWPM() float64 {
	return g.typing.WPM() + float64(g.wpmBonus)
}
