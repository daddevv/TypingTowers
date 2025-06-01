package engine

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Engine represents the main game engine that manages the game state, menus, and rendering.
type Engine struct {
	Screen *ebiten.Image // Internal screen for rendering at 1920x1080
	Background *ebiten.Image // Background image for the game
}

// NewEngine initializes a new game engine instance.
func NewEngine() *Engine {
	

	return &Engine{
		Screen: ebiten.NewImage(1920, 1080),
		Background: ebiten.NewImage(1920, 1080), // Placeholder for background image

	}
}

// Update calls the appropriate update method based on the current state of the engine.
func (e *Engine) Update() error {

	return nil
}

// Draw constructs the frame on e.Screen and then renders it to the actual screen in one draw call
func (e *Engine) Draw(screen *ebiten.Image) {
	e.Screen.Clear() // Clear the internal screen

	// draw state to e.Screen
	e.Screen.DrawImage(e.Background, nil) // Draw the background image

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