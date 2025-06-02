package engine

import (
	"fmt"
	"image/color"
	"sort"
	"td/internal/game"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// Engine represents the main game engine that manages the game state, menus, and rendering.
type Engine struct {
	DebugDisplay map[string]any // Debug display for various game states
	DebugEnabled bool // Flag to enable or disable debug display
	Screen     *ebiten.Image // Internal screen for rendering at 1920x1080
	Game      *game.Game // The current game state, including player, enemies, etc.
}

// NewEngine initializes a new game engine instance.
func NewEngine() *Engine {
	return &Engine{
		DebugDisplay: make(map[string]any), // Initialize the debug display
		DebugEnabled: true, // Debug display is initially enabled
		Screen:     ebiten.NewImage(1920, 1080), // Create a new image for the internal screen
		Game:      game.NewGame(), // Initialize the game state
	}
}

// Update calls the appropriate update method based on the current state of the engine.
func (e *Engine) Update() error {
	err := e.Game.Update() // Update the game logic
	if err != nil {
		return err
	}

	// Handle input for quitting the game
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return ebiten.Termination
	}

	// Ctrl+F for fullscreen toggle
	if inpututil.IsKeyJustPressed(ebiten.KeyF) && ebiten.IsKeyPressed(ebiten.KeyControl) {
		if ebiten.IsFullscreen() {
			ebiten.SetFullscreen(false) // Exit fullscreen
		} else {
			ebiten.SetFullscreen(true) // Enter fullscreen
		}
	}

	// Ctrl+D for debug toggle
	if inpututil.IsKeyJustPressed(ebiten.KeyD) && ebiten.IsKeyPressed(ebiten.KeyControl) {
		e.DebugEnabled = !e.DebugEnabled // Toggle debug display
		if e.DebugEnabled {
			fmt.Println("Debug display enabled")
		} else {
			fmt.Println("Debug display disabled")
		}
	}

	// Update the debug display with current FPS and mouse position
	e.DebugDisplay["FPS"] = ebiten.ActualFPS() // Update FPS in debug display
	e.DebugDisplay["Mobs"] = len(e.Game.Mobs) // Update the number of mobs in debug display
	for i, mob := range e.Game.Mobs {
		// Add each mob and its target to the debug display
		if mob != nil {
			e.DebugDisplay[fmt.Sprintf("Mob %d", i)] = fmt.Sprintf("Position: (%.2f, %.2f), Target: (%.2f, %.2f)",
				mob.Position.X, mob.Position.Y, mob.Target.X, mob.Target.Y)
		}
	}

	return nil
}

// Draw constructs the frame on e.Screen and then renders it to the actual screen in one draw call
func (e *Engine) Draw(screen *ebiten.Image) {
	e.Screen.Clear() // Clear the internal screen

	e.Game.Draw(e.Screen) // Draw the game state to the internal screen

	//print the debug display if enabled
	var text string
	if e.DebugEnabled {
			// Draw 48x48 grid outline for debugging
		for x := range 1920 {
			for y := range 1080 {
				if x%48 == 0 || y%48 == 0 {
					e.Screen.Set(x, y, color.RGBA{255, 0, 0, 255}) // Red lines for grid
				}
			}
		}

		// sort the debug display keys for consistent output
		keys := make([]string, 0, len(e.DebugDisplay))
		for key := range e.DebugDisplay {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		for _, key := range keys {
			text += fmt.Sprintf("%s: %v\n", key, e.DebugDisplay[key]) // Format the debug display text
		}
		ebitenutil.DebugPrint(e.Screen, text) // Print the debug text on the screen
	}

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
