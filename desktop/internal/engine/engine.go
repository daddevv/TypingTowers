package engine

import (
	"fmt"
	"image/color"

	"td/internal/building"
	"td/internal/enemy"
	"td/internal/goblin"
	"td/internal/player"
	"td/internal/ui"
	"td/internal/world"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// Engine represents the main game engine that manages the game state, menus, and rendering.
type Engine struct {
	DebugDisplay map[string]any // Debug display for various game states
	DebugEnabled bool // Flag to enable or disable debug display
	Screen     *ebiten.Image // Internal screen for rendering at 1920x1080
	Background *ebiten.Image // Background image for the game
	Goblins  []*enemy.Mob // List of goblins in the game
	TestAnimation *ui.Animation // Animation for testing purposes
	TestSpawner *goblin.GoblinSpawner
}

// NewEngine initializes a new game engine instance.
func NewEngine() *Engine {
	// Create a new image to hold the background
	bg := world.Example.Background

	// Draw 48x48 grid outline for debugging
	for x := range 1920 {
		for y := range 1080 {
			if x%48 == 0 || y%48 == 0 {
				bg.Set(x, y, color.RGBA{255, 0, 0, 255}) // Red lines for grid
			}
		}
	}

	debugDisplay := make(map[string]any) // Initialize the debug display map
	debugDisplay["FPS"] = 0 // Placeholder for FPS display
	mouseX, mouseY := ebiten.CursorPosition() // Get the current mouse position
	debugDisplay["MouseX"] = mouseX // Store mouse X position in debug display
	debugDisplay["MouseY"] = mouseY // Store mouse Y position in debug display
	debugDisplay["PlayerPosition"] = nil // Store player position in debug display
	debugDisplay["Goblins"] = 0 // Placeholder for goblin count
	

	return &Engine{
		DebugDisplay: debugDisplay, // Initialize the debug display
		DebugEnabled: true, // Debug display is initially enabled
		Goblins:  []*enemy.Mob{}, // Initialize the goblins list
		Screen:     ebiten.NewImage(1920, 1080),
		Background: bg,
		TestAnimation: player.WalkingAnimation,
		TestSpawner: goblin.NewGoblinSpawner(building.MobSpawnerLevel1, 400, 400, 60), // Example spawner
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
	// D for debug toggle
	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		e.DebugEnabled = !e.DebugEnabled // Toggle debug display
		if e.DebugEnabled {
			fmt.Println("Debug display enabled")
		} else {
			fmt.Println("Debug display disabled")
		}
	}
	// Esc to exit the game
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return ebiten.Termination // Exit the game
	}

	e.TestAnimation.Update()
	e.Goblins = e.TestSpawner.Update(e.Goblins) // Update the test spawner, passing the goblins list


	// Update the debug display with current FPS and mouse position
	e.DebugDisplay["FPS"] = ebiten.ActualFPS() // Update FPS in debug display
	mouseX, mouseY := ebiten.CursorPosition() // Get the current mouse position
	e.DebugDisplay["MouseX"] = mouseX // Store mouse X position in debug display
	e.DebugDisplay["MouseY"] = mouseY // Store mouse Y position in debug display
	// Update goblin count in debug display
	e.DebugDisplay["Goblins"] = len(e.Goblins) // Store the number of goblins in debug display
	


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
	e.Screen.DrawImage(e.TestAnimation.Frame(), opts)

	// Draw the test spawner
	opts = &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(e.TestSpawner.Position.X, e.TestSpawner.Position.Y) // Position the spawner
	e.Screen.DrawImage(e.TestSpawner.Sprite, opts) // Draw the spawner sprite

	// Draw the goblins
	for _, goblin := range e.Goblins {
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(goblin.Position.X, goblin.Position.Y) // Position the goblin
		e.Screen.DrawImage(goblin.Sprite, opts) // Draw the goblin sprite
	}

		//print the debug display if enabled
	if e.DebugEnabled {
		ebitenutil.DebugPrint(e.Screen, fmt.Sprintf("FPS: %.2f\nMouse: (%d, %d)\nGoblins: %d",
			e.DebugDisplay["FPS"].(float64),
			e.DebugDisplay["MouseX"].(int),
			e.DebugDisplay["MouseY"].(int),
			e.DebugDisplay["Goblins"].(int),
		))
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
