package game

import (
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	mousePressed bool // Track if the mouse button is pressed
	clickedTileX int // X coordinate of the clicked tile
	clickedTileY int // Y coordinate of the clicked tile
)

// Game represents the game state and implements ebiten.Game interface.
type Game struct {
	// Add game state fields here, such as score, player state, etc.
}

// NewGame creates a new instance of the Game.
func NewGame() *Game {
	ebiten.SetWindowTitle("TypeDefense")
	ebiten.SetWindowSize(1920/4, 1080/4)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetFullscreen(true)
	// ebiten.SetWindowDecorated(false)
	return &Game{
		// Initialize game state fields here if needed.
	}
}

// Update updates the game state. This method is called every frame.
func (g *Game) Update() error {
	// Update game logic here, such as player input, enemy movement, etc.
	// Return nil if the update is successful, or an error if something goes wrong.
	return nil
}

// Draw renders the game to the screen. This method is called every frame.
func (g *Game) Draw(screen *ebiten.Image) {
	// Draw the game state to the screen.
	// Use ebiten's drawing functions to render images, text, etc.
	screen.Clear()
	screen.DrawImage(BACKGROUND_GRID, nil)
	highlightHoverAndClickAndDrag(screen, "line") // Change shape as needed: "rectangle", "circle", "line", etc.
}

// Layout returns the size of the game screen in pixels.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 1920, 1080
}

// highlightHoverAndClickAndDrag highlights the tile under the mouse cursor.
func highlightHoverAndClickAndDrag(screen *ebiten.Image, shape string) {
	mouseX, mouseY := ebiten.CursorPosition()
	if mouseX < 0 || mouseY < 28 || mouseX >= 1920 || mouseY >= 1052 {
		return // Ignore mouse position outside the screen
	}
	tileX, tileY := tileAtPosition(mouseX, mouseY)

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
		switch shape {
		case "rectangle":
			// Draw a filled rectangle between the start and end tiles
			for x := minX; x <= maxX; x++ {
				for y := minY; y <= maxY; y++ {
					op := &ebiten.DrawImageOptions{}
					op.GeoM.Translate(float64(x*32), float64(28+y*32))
					screen.DrawImage(ImgHighlightTile, op)
				}
			}
		case "circle":
			// Draw a circle around the dragged area
			centerX := (minX + maxX) * 16
			centerY := (minY + maxY) * 16 + 28 // Adjust for top margin
			radius := 16 * (maxX - minX + maxY - minY) / 2 // Average radius based on dragged area
			for x := -radius; x <= radius; x += 32 {
				for y := -radius; y <= radius; y += 32 {
					if x*x+y*y <= radius*radius {
						tX, tY := tileAtPosition(centerX+x, centerY+y)
						// Ensure we only draw within the bounds of the grid
						if tX >= 0 && tX < 60 && tY >= 0 && tY < 34 {
							op := &ebiten.DrawImageOptions{}
							op.GeoM.Translate(float64(centerX+x), float64(centerY+y))
							screen.DrawImage(ImgHighlightTile, op)
						}
					}
				}
			}
		case "line":
			// Draw a pixelated line from clicked tile to current tile using Bresenham's algorithm
			x0, y0 := clickedTileX, clickedTileY
			x1, y1 := tileX, tileY
			dx := abs(x1 - x0)
			dy := abs(y1 - y0)
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
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(x0*32), float64(28+y0*32))
				screen.DrawImage(ImgHighlightTile, op)
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
			}
		// Add more shapes as needed
		default:
			// Default to rectangle if no shape specified
			for x := minX; x <= maxX; x++ {
				for y := minY; y <= maxY; y++ {
					op := &ebiten.DrawImageOptions{}
					op.GeoM.Translate(float64(x*32), float64(28+y*32))
					screen.DrawImage(ImgHighlightTile, op)
				}
			}
		}
	} else {
		// Not dragging: highlight only hovered tile
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(tileX*32), float64(28+tileY*32))
		screen.DrawImage(ImgHighlightTile, op)
	}

	ebitenutil.DebugPrintAt(screen, "Hovering over tile: "+strconv.Itoa(tileX)+", "+strconv.Itoa(tileY), 10, 2)
	ebitenutil.DebugPrintAt(screen, "Mouse Position: "+strconv.Itoa(mouseX)+", "+strconv.Itoa(mouseY), 10, 14)
	if mousePressed {
		ebitenutil.DebugPrintAt(screen, "Dragging from: "+strconv.Itoa(clickedTileX)+", "+strconv.Itoa(clickedTileY), 190, 2)
	}
}

// Helper function for absolute value of int
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
