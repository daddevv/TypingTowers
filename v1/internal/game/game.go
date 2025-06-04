package game

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	mousePressed bool
	clickedTileX int
	clickedTileY int
	houses       = make(map[string]struct{})
)

// Game represents the game state and implements ebiten.Game interface.
type Game struct {
	screen      *ebiten.Image
	input       InputHandler
	towers      []*Tower
	mobs        []*Mob
	projectiles []*Projectile
	base        *Base
	hud         *HUD
	gameOver    bool
	paused      bool
	gold        int

	shopOpen bool

	cfg *Config

	currentWave   int
	spawnInterval int
	spawnTicker   int
	mobsToSpawn   int
}

// NewGame creates a new instance of the Game.
func NewGame() *Game {
	return NewGameWithConfig(DefaultConfig)
}

// NewGameWithConfig creates a new instance of the Game using the provided configuration.
func NewGameWithConfig(cfg Config) *Game {
	ebiten.SetWindowTitle("TypingTowers")
	ebiten.SetWindowSize(1920/8, 1080/8)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	// ebiten.SetFullscreen(true)

	g := &Game{
		screen:        ebiten.NewImage(1920, 1080),
		input:         NewInput(),
		paused:        false,
		gold:          0,
		shopOpen:      false,
		currentWave:   1,
		spawnInterval: 60,
		spawnTicker:   0,
		mobsToSpawn:   3,
		cfg:           &cfg,

		mobs:        make([]*Mob, 0),
		projectiles: make([]*Projectile, 0),
	}

	tx, ty := tilePosition(1, 16)
	hp := int(cfg.J)
	if hp <= 0 {
		hp = BaseStartingHealth
	}
	g.base = NewBase(float64(tx+32), float64(ty+16), hp)

	tx, ty = tilePosition(2, 16)
	tower := NewTower(g, float64(tx+16), float64(ty+16))
	g.towers = []*Tower{tower}
	g.startWave()
	g.hud = NewHUD(g)
	return g
}

// Update updates the game state. This method is called every frame.
func (g *Game) Update() error {
	g.input.Update()

	if g.input.Reload() {
		if err := g.reloadConfig(ConfigFile); err != nil {
			fmt.Println("reload config:", err)
		}
	}

	if g.input.Space() {
		g.paused = !g.paused
	}

	if g.input.Quit() {
		return ebiten.Termination
	}

	if g.paused {
		return nil
	}

	if g.gameOver {
		return nil
	}

	if g.shopOpen {
		if g.input.Enter() {
			g.shopOpen = false
			g.currentWave++
			g.startWave()
		}
		return nil
	}

	if g.mobsToSpawn > 0 {
		g.spawnTicker++
		if g.spawnTicker >= g.spawnInterval {
			g.spawnTicker = 0
			g.spawnMob()
			g.mobsToSpawn--
		}
	} else if len(g.mobs) == 0 {
		g.shopOpen = true
	}

	for _, t := range g.towers {
		t.Update()
	}

	g.base.Update()

	for i := 0; i < len(g.projectiles); {
		p := g.projectiles[i]
		p.Update()
		if !p.alive {
			g.projectiles = append(g.projectiles[:i], g.projectiles[i+1:]...)
			continue
		}
		i++
	}

	for i := 0; i < len(g.mobs); {
		m := g.mobs[i]
		m.Update()
		bx, by, bw, bh := g.base.Bounds()
		dx := m.pos.X - float64(bx+bw/2)
		dy := m.pos.Y - float64(by+bh/2)
		if math.Hypot(dx, dy) < float64(m.width/2+bw/2) {
			g.base.Damage(1)
			m.alive = false
		}
		if !m.alive {
			g.mobs = append(g.mobs[:i], g.mobs[i+1:]...)
			g.gold++
			continue
		}
		i++
	}

	if !g.base.Alive() {
		g.gameOver = true
	}

	return nil
}

// Draw renders the game to the screen. This method is called every frame.
func (g *Game) Draw(screen *ebiten.Image) {
	g.screen.Clear()
	drawBackgroundTilemap(g.screen)

	if g.gameOver {
		ebitenutil.DebugPrintAt(g.screen, "Game Over", 900, 540)
		g.renderFrame(screen)
		return
	}

	if g.shopOpen {
		ebitenutil.DebugPrintAt(g.screen, "-- SHOP -- press Enter", 850, 520)
		g.renderFrame(screen)
		return
	}

	g.base.Draw(g.screen)

	for _, t := range g.towers {
		t.Draw(g.screen)
	}
	for _, p := range g.projectiles {
		p.Draw(g.screen)
	}
	for _, m := range g.mobs {
		m.Draw(g.screen)
	}

	if g.paused {
		ebitenutil.DebugPrintAt(g.screen, "-- PAUSED --", 900, 520)
		g.renderFrame(screen)
		return
	}

	if g.hud != nil {
		g.hud.Draw(g.screen)
	}

	highlightHoverAndClickAndDrag(g.screen, "line")

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
	hp := 1
	if g.cfg != nil {
		hp = int(1 + float64(g.currentWave-1)*g.cfg.N)
		if hp < 1 {
			hp = 1
		}
	}
	m := NewMob(float64(x+16), float64(y+16), g.base, hp)
	g.mobs = append(g.mobs, m)
}

// startWave initializes spawn counters for the next wave.
func (g *Game) startWave() {
	g.spawnTicker = 0
	g.mobsToSpawn = g.currentWave * 3
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

	ebitenutil.DebugPrintAt(screen, "Hovering over tile: "+strconv.Itoa(tileX)+", "+strconv.Itoa(tileY), 10, 2)
	ebitenutil.DebugPrintAt(screen, "Mouse Position: "+strconv.Itoa(mouseX)+", "+strconv.Itoa(mouseY), 10, 14)
	if mousePressed {
		ebitenutil.DebugPrintAt(screen, "Dragging from: "+strconv.Itoa(clickedTileX)+", "+strconv.Itoa(clickedTileY), 190, 2)
	}
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
	hp := int(cfg.J)
	if hp <= 0 {
		hp = BaseStartingHealth
	}
	g.base.health = hp
	for _, t := range g.towers {
		t.ApplyConfig(cfg)
	}
	return nil
}
