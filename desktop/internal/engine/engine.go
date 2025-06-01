package engine

import (
	"image/color"

	"td/internal/player"
	"td/internal/tilemap"
	"td/internal/ui"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// Engine represents the main game engine that manages the game state, menus, and rendering.
type Engine struct {
	Screen     *ebiten.Image // Internal screen for rendering at 1920x1080
	Background *ebiten.Image // Background image for the game
	TestAnimation *ui.Animation // Animation for testing purposes
}

// NewEngine initializes a new game engine instance.
func NewEngine() *Engine {
	// Create a new image to hold the background
	islandWidth := 40     // Number of tiles horizontally (max 39)
	islandHeight := 18    // Number of tiles vertically (max 22)
	verticalOffset := 2   // Number of tiles to offset vertically
	horizontalOffset := 2 // Number of tiles to offset horizontally
	bg := ebiten.NewImage(1920, 1080)
	for y := verticalOffset * 48; y < 1080; y += 48 {
		for x := horizontalOffset * 48; x < 1920; x += 48 {
			tileX := x / 48
			tileY := y / 48
			var tileImg *ebiten.Image
			if tileX == horizontalOffset && tileY == verticalOffset {
				tileImg = tilemap.Cliff["top_left"]
			} else if tileX == horizontalOffset && tileY == verticalOffset+islandHeight-1 {
				tileImg = tilemap.Cliff["bottom_left"]
			} else if tileX == horizontalOffset+islandWidth-1 && tileY == verticalOffset {
				tileImg = tilemap.Cliff["top_right"]
			} else if tileX == horizontalOffset+islandWidth-1 && tileY == verticalOffset+islandHeight-1 {
				tileImg = tilemap.Cliff["bottom_right"]
			} else if tileY == verticalOffset && tileX > horizontalOffset && tileX < horizontalOffset+islandWidth-1 {
				tileImg = tilemap.Cliff["top"]
			} else if tileY == verticalOffset+islandHeight-1 && tileX > horizontalOffset && tileX < horizontalOffset+islandWidth-1 {
				tileImg = tilemap.Cliff["bottom"]
			} else if tileX == horizontalOffset && tileY > verticalOffset && tileY < verticalOffset+islandHeight-1 {
				tileImg = tilemap.Cliff["left"]
			} else if tileX == horizontalOffset+islandWidth-1 && tileY > verticalOffset && tileY < verticalOffset+islandHeight-1 {
				tileImg = tilemap.Cliff["right"]
			} else if tileX > horizontalOffset && tileX < horizontalOffset+islandWidth-1 && tileY > verticalOffset && tileY < verticalOffset+islandHeight-1 {
				tileImg = tilemap.Cliff["center"] // Default to center tile
			}

			if tileImg != nil {
				opts := &ebiten.DrawImageOptions{}
				opts.GeoM.Translate(float64(x), float64(y)) // Position the tile
				bg.DrawImage(tileImg, opts)
			}
		}
	}

	// Draw 48x48 grid outline for debugging
	for x := range 1920 {
		for y := range 1080 {
			if x%48 == 0 || y%48 == 0 {
				bg.Set(x, y, color.RGBA{255, 0, 0, 255}) // Red lines for grid
			}
		}
	}

	return &Engine{
		Screen:     ebiten.NewImage(1920, 1080),
		Background: bg,
		TestAnimation: player.IdleAnimation,
	}
}

// Update calls the appropriate update method based on the current state of the engine.
func (e *Engine) Update() error {
	// F for fullscreen toggle
	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		if ebiten.IsFullscreen() {
			ebiten.SetFullscreen(false) // Exit fullscreen
		} else {
			ebiten.SetFullscreen(true) // Enter fullscreen
		}
	}

	e.TestAnimation.Update()

	return nil
}

// Draw constructs the frame on e.Screen and then renders it to the actual screen in one draw call
func (e *Engine) Draw(screen *ebiten.Image) {

	e.Screen.Clear() // Clear the internal screen

	// draw state to e.Screen
	e.Screen.DrawImage(e.Background, nil) // Draw the background image

	// Draw the test animation at a fixed position
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(100, 100) // Position the animation at (100, 100)
	// Draw the current frame of the test animation
	e.Screen.DrawImage(e.TestAnimation.Frame(), opts)

	e.renderFrame(screen)
}

// renderFrame scales and draws the internal screen buffer to the actual screen
func (e *Engine) renderFrame(screen *ebiten.Image) {
	// Now scale the internal screen to the window size
	w, h := screen.Bounds().Dx(), screen.Bounds().Dy()
	scaleX := float64(w) / 1920.0
	scaleY := float64(h) / 1080.0
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(scaleX, scaleY)
	screen.DrawImage(e.Screen, opts)
}

// Layout defines the size of the internal canvas used by the engine.
// This is always 1920x1080, regardless of the actual window size.
func (g *Engine) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	// Always use 1920x1080 for the internal canvas
	return 1920, 1080
}
