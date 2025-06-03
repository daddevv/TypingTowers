package game

import (
	"fmt"
	"math"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	mousePressed bool // Track if the mouse button is pressed
	clickedTileX int // X coordinate of the clicked tile
	clickedTileY int // Y coordinate of the clicked tile
	houses     = make(map[string]struct{}) // Store house tile positions as "x,y"
)

// Game represents the game state and implements ebiten.Game interface.
type Game struct {
	screen        *ebiten.Image // Internal screen buffer for rendering
	input 		InputHandler // Input handler for processing user input
}

// NewGame creates a new instance of the Game.
func NewGame() *Game {
	ebiten.SetWindowTitle("TypeDefense")
	ebiten.SetWindowSize(1920/8, 1080/8)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	// ebiten.SetFullscreen(true)

	return &Game{
		screen:        ebiten.NewImage(1920, 1080), // Create a new image for the game screen
		input: NewInput(),
	}
}

// Update updates the game state. This method is called every frame.
func (g *Game) Update() error {
	g.input.Update()

	if g.input.Quit() {
		return ebiten.Termination
	}

	// Update game logic here, such as player input, enemy movement, etc.
	// Return nil if the update is successful, or an error if something goes wrong.
	return nil
}

// Draw renders the game to the screen. This method is called every frame.
func (g *Game) Draw(screen *ebiten.Image) {
	
	drawBackgroundTilemap(screen) // Draw the background tilemap
	highlightHoverAndClickAndDrag(screen, "line") // Change shape as needed: "rectangle", "circle", "line", etc.

	g.renderFrame(screen) // Scale and draw the internal screen buffer to the actual screen
}

// renderFrame scales and draws the internal screen buffer to the actual screen
func (g *Game) renderFrame(screen *ebiten.Image) {
	screen.Clear()
	// Now scale the internal screen to the window size
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


// highlightHoverAndClickAndDrag highlights the tile under the mouse cursor.
func highlightHoverAndClickAndDrag(screen *ebiten.Image, shape string) {
	// Determine the tile at the current mouse position
	mouseX, mouseY := ebiten.CursorPosition()
	if mouseX < 0 || mouseY < TopMargin || mouseX >= 1920 || mouseY >= 1080- TopMargin {
		return // Ignore mouse position outside the screen
	}
	tileX, tileY := tileAtPosition(mouseX, mouseY)

	// Right Click to add a house tile (persistently)
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		if tileX >= 0 && tileX <= 59 && tileY >= 0 && tileY <= 33 {
			id := fmt.Sprintf("%d,%d", tileX, tileY)
			houses[id] = struct{}{}
		}
	}

	// Handle mouse press/release and track drag start
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		if !mousePressed {
			// Mouse just pressed
			mousePressed = true
			clickedTileX = tileX
			clickedTileY = tileY
		}
	} else {
		mousePressed = false
	}

	// Draw highlight(s)
	if mousePressed {
		// Dragging: highlight rectangle from clickedTile to current tile
		minX, maxX := clickedTileX, tileX
		if minX > maxX {
			minX, maxX = maxX, minX
		}
		minY, maxY := clickedTileY, tileY
		if minY > maxY {
			minY, maxY = maxY, minY
		}
		// Clamp rectangle bounds to grid
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
			// Draw a filled rectangle between the start and end tiles
			for x := minX; x <= maxX; x++ {
				for y := minY; y <= maxY; y++ {
					op := &ebiten.DrawImageOptions{}
					op.GeoM.Translate(float64(x*32), float64(TopMargin+y*32))
					screen.DrawImage(ImgHighlightTile, op)
				}
			}
		case "circle":
			// Draw a circle around the dragged area, but clamp to grid
			centerTileX := (minX + maxX) / 2
			centerTileY := (minY + maxY) / 2
			radius := (maxX - minX) / 2
			for x := centerTileX - radius; x <= centerTileX + radius; x++ {
				for y := centerTileY - radius; y <= centerTileY + radius; y++ {
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
			// Draw a pixelated line from clicked tile to current tile using Bresenham's algorithm, clamped to grid
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
				// Clamp to grid
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
				// Stop if out of grid
				if x0 < 0 || x0 > 59 || y0 < 0 || y0 > 33 {
					break
				}
			}
		// Add more shapes as needed
		default:
			// Default to rectangle if no shape specified
			for x := minX; x <= maxX; x++ {
				for y := minY; y <= maxY; y++ {
					op := &ebiten.DrawImageOptions{}
					op.GeoM.Translate(float64(x*32), float64(TopMargin+y*32))
					screen.DrawImage(ImgHighlightTile, op)
				}
			}
		}
	} else {
		// Not dragging: highlight only hovered tile, clamped to grid
		if tileX >= 0 && tileX <= 59 && tileY >= 0 && tileY <= 33 {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(tileX*32), float64(TopMargin+tileY*32))
			screen.DrawImage(ImgHighlightTile, op)
		}
	}

	// Draw all house tiles
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