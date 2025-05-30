package ui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// Screen represents a buffer that can be drawn to the screen.
// It provides methods to update the screen state and draw it to the main canvas.
type Screen interface {
	// Update processes input and updates the screen state.
	Update() error
	// Draw renders the screen to the main canvas.
	Draw(canvas *ebiten.Image)
}

// BaseScreen provides a default implementation of the Screen interface.
type BaseScreen struct {
	// Frame is the internal buffer where the screen is drawn.
	Frame *ebiten.Image
}

// NewBaseScreen creates a new BaseScreen instance with a default frame size of 1920x1080.
func NewBaseScreen() *BaseScreen {
	return &BaseScreen{
		Frame: ebiten.NewImage(1920, 1080),
	}
}

// Update implements the Screen interface for BaseScreen.
func (s *BaseScreen) Update() error {
	// Default implementation does nothing, can be overridden by concrete screens.
	return nil
}

// Draw implements the Screen interface for BaseScreen.
func (s *BaseScreen) Draw(canvas *ebiten.Image) {
	// Draw the internal frame to the provided canvas.
	if s.Frame != nil {
		w, h := canvas.Bounds().Dx(), canvas.Bounds().Dy()
		scaleX := float64(w) / 1920.0
		scaleY := float64(h) / 1080.0
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Scale(scaleX, scaleY)
		canvas.DrawImage(s.Frame, opts)
	}
}

// Clear fills the screen with a transparent color, effectively clearing it.
func (s *BaseScreen) Clear() {
	s.Frame.Fill(color.Transparent) // Fill with transparent color
}
