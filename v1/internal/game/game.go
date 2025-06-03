package game

import (
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
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
	highlightHover(screen)
}

// Layout returns the size of the game screen in pixels.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 1920, 1080
}

// highlightHover highlights the tile under the mouse cursor.
func highlightHover(screen *ebiten.Image) {
	mouseX, mouseY := ebiten.CursorPosition()
	if mouseX < 0 || mouseY < 28 || mouseX >= 1920 || mouseY >= 1052 {
		return // Ignore mouse position outside the screen
	}
	tileX := mouseX / 32
	tileY := (mouseY - 28) / 32
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(tileX*32), float64(28+tileY*32))
	screen.DrawImage(ImgHighlightTile, op)
	ebitenutil.DebugPrintAt(screen, "Hovering over tile: "+strconv.Itoa(tileX)+", "+strconv.Itoa(tileY), 10, 2)
	ebitenutil.DebugPrintAt(screen, "Mouse Position: "+strconv.Itoa(mouseX)+", "+strconv.Itoa(mouseY), 10, 14)
}
