package game

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"

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
	screen       *ebiten.Image
	input        InputHandler
	towers       []*Tower
	mobs         []*Mob
	projectiles  []*Projectile
	spawnCounter int
}

// NewGame creates a new instance of the Game.
func NewGame() *Game {
	ebiten.SetWindowTitle("TypeDefense")
	ebiten.SetWindowSize(1920/8, 1080/8)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	// ebiten.SetFullscreen(true)

	rand.Seed(time.Now().UnixNano())

	g := &Game{
		screen: ebiten.NewImage(1920, 1080),
		input:  NewInput(),
	}
	tx, ty := tilePosition(2, 16)
	tower := NewTower(g, float64(tx+16), float64(ty+16))
	g.towers = []*Tower{tower}
	return g
}

// Update updates the game state. This method is called every frame.
func (g *Game) Update() error {
	g.input.Update()

	if g.input.Quit() {
		return ebiten.Termination
	}

	g.spawnCounter++
	if g.spawnCounter > 120 {
		g.spawnCounter = 0
		g.spawnMob()
	}

	for _, t := range g.towers {
		t.Update()
	}

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
		if !m.alive {
			g.mobs = append(g.mobs[:i], g.mobs[i+1:]...)
			continue
		}
		i++
	}

	return nil
}

// Draw renders the game to the screen. This method is called every frame.
func (g *Game) Draw(screen *ebiten.Image) {
	g.screen.Clear()
	drawBackgroundTilemap(g.screen)

	for _, t := range g.towers {
		t.Draw(g.screen)
	}
	for _, p := range g.projectiles {
		p.Draw(g.screen)
	}
	for _, m := range g.mobs {
		m.Draw(g.screen)
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
	m := NewMob(float64(x+16), float64(y+16))
	g.mobs = append(g.mobs, m)
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
