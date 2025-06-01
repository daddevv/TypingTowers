package engine

import (
	"fmt"
	"image/color"
	"sort"

	"td/internal/building"
	"td/internal/enemy"
	"td/internal/goblin"
	"td/internal/math"
	"td/internal/player"
	"td/internal/ui"
	"td/internal/world"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// Engine represents the main game engine that manages the game state, menus, and rendering.
type Engine struct {
	PlayerPos math.Vec2 // Player position in the game world
	DebugDisplay map[string]any // Debug display for various game states
	DebugEnabled bool // Flag to enable or disable debug display
	Screen     *ebiten.Image // Internal screen for rendering at 1920x1080
	Background *ebiten.Image // Background image for the game
	Goblins  []*enemy.Mob // List of goblins in the game
	TestAnimation *ui.Animation // Animation for testing purposes
	TestSpawner *goblin.GoblinSpawner
	TestMob *enemy.Mob // Example mob for testing
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
		PlayerPos: math.Vec2{X: 600, Y: 600}, // Initialize player position
		DebugDisplay: debugDisplay, // Initialize the debug display
		DebugEnabled: true, // Debug display is initially enabled
		Goblins:  []*enemy.Mob{}, // Initialize the goblins list
		Screen:     ebiten.NewImage(1920, 1080), // Create a new image for the internal screen
		Background: bg,
		TestAnimation: player.WalkingAnimation,
		TestSpawner: goblin.NewGoblinSpawner(building.MobSpawnerLevel1, 400, 400, 60), // Example spawner
		TestMob: enemy.NewMob("Goblin", 100, 200, 200, 0, 0, 100, 100), // Example mob for testing
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
	// Ctrl+D for debug toggle
	if inpututil.IsKeyJustPressed(ebiten.KeyD) && ebiten.IsKeyPressed(ebiten.KeyControl) {
		e.DebugEnabled = !e.DebugEnabled // Toggle debug display
		if e.DebugEnabled {
			fmt.Println("Debug display enabled")
		} else {
			fmt.Println("Debug display disabled")
		}
	}

	deltaVec := math.Vec2{X: 0, Y: 0} // Initialize delta vector for player movement
	// W for moving the player up
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		deltaVec.Y -= 10 // Move player up by 10 units
	}
	if ebiten.IsKeyPressed(ebiten.KeyE) {
		deltaVec.Y -= 10 // Move player up by 10 units (alternative key for testing)
		deltaVec.X += 10 // Move player right by 10 units (alternative key for testing)
	}
	// S for moving the player down
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		deltaVec.Y += 10 // Move player down by 10 units
	}
	// E for moving the player up and right (testing purposes)
	if ebiten.IsKeyPressed(ebiten.KeyC) {
		deltaVec.Y += 10 // Move player down by 10 units (alternative key for testing)
		deltaVec.X += 10 // Move player right by 10 units (alternative key for testing)
	}
	// A for moving the player left
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		deltaVec.X -= 10 // Move player left by 10 units
	}
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		deltaVec.X -= 10 // Move player left by 10 units (alternative key for testing)
		deltaVec.Y -= 10 // Move player up by 10 units (alternative key for testing)
	}
	// D for moving the player right
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		deltaVec.X += 10 // Move player right by 10 units
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		deltaVec.X += 10 // Move player right by 10 units (alternative key)
	}
	if ebiten.IsKeyPressed(ebiten.KeyZ) {
		deltaVec.X -= 10 // Move player right by 10 units (alternative key for testing)
		deltaVec.Y += 10 // Move player up by 10 units (alternative key for testing)
	}

	// Normalize the delta vector to ensure consistent movement speed
	deltaPos := deltaVec.Normalize()
	// Update player position based on input
	if deltaPos.X != 0 || deltaPos.Y != 0 {
		e.PlayerPos.X += deltaPos.X * 2 // Update player X position
		e.PlayerPos.Y += deltaPos.Y * 2 // Update player Y position
	}
	e.TestAnimation.Update()

	// Esc to exit the game
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return ebiten.Termination // Exit the game
	}
	

	e.Goblins = e.TestSpawner.Update(e.Goblins) // Update the test spawner, passing the goblins list

	// Update the test mob
	e.TestMob.Update(&math.Vec2{
		X: e.PlayerPos.X + 48, // Set the target position to the player's position plus an offset
		Y: e.PlayerPos.Y + 32, // Set the target position to the player's position plus an offset
	}) // Update the test mob's position towards the player

	// Update the debug display with current FPS and mouse position
	e.DebugDisplay["FPS"] = ebiten.ActualFPS() // Update FPS in debug display
	mouseX, mouseY := ebiten.CursorPosition() // Get the current mouse position
	e.DebugDisplay["MouseX"] = mouseX // Store mouse X position in debug display
	e.DebugDisplay["MouseY"] = mouseY // Store mouse Y position in debug display
	// Update goblin count in debug display
	e.DebugDisplay["PlayerPosition"] = e.PlayerPos // Update player position in debug display
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
	opts.GeoM.Translate(e.PlayerPos.X, e.PlayerPos.Y) // Position the animation at the player position
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
	var text string
	if e.DebugEnabled {
		// sort the debug display keys for consistent output
		keys := make([]string, 0, len(e.DebugDisplay))
		for key := range e.DebugDisplay {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		for _, key := range keys {
			text += fmt.Sprintf("%s: %v\n", key, e.DebugDisplay[key]) // Format the debug display text
		}
	}
	ebitenutil.DebugPrint(e.Screen, text) // Print the debug text on the screen

	// Draw the test mob
	opts = &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(e.TestMob.Position.X, e.TestMob.Position.Y) // Position the mob
	e.Screen.DrawImage(e.TestMob.Sprite, opts) // Draw the mob sprite

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
