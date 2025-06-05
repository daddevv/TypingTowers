package game

import (
	"image/color"
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
)

// stubInputOverlay implements InputHandler with no-op behavior for overlay tests.
type stubInputOverlay struct{}

func (s *stubInputOverlay) TypedChars() []rune { return nil }
func (s *stubInputOverlay) Update()            {}
func (s *stubInputOverlay) Reset()             {}
func (s *stubInputOverlay) Backspace() bool    { return false }
func (s *stubInputOverlay) Space() bool        { return false }
func (s *stubInputOverlay) Quit() bool         { return false }
func (s *stubInputOverlay) Reload() bool       { return false }
func (s *stubInputOverlay) Enter() bool        { return false }
func (s *stubInputOverlay) Left() bool         { return false }
func (s *stubInputOverlay) Right() bool        { return false }
func (s *stubInputOverlay) Up() bool           { return false }
func (s *stubInputOverlay) Down() bool         { return false }
func (s *stubInputOverlay) Build() bool        { return false }
func (s *stubInputOverlay) Save() bool         { return false }
func (s *stubInputOverlay) Load() bool         { return false }
func (s *stubInputOverlay) SelectTower() bool  { return false }
func (s *stubInputOverlay) Command() bool      { return false }
func (s *stubInputOverlay) TechMenu() bool     { return false }

// TestDrawTowerSelectionOverlay verifies the HUD draws overlay highlights without panic.
func TestDrawTowerSelectionOverlay(t *testing.T) {
	g := NewGame()
	g.input = &stubInputOverlay{}
	g.enterTowerSelectMode()
	hud := NewHUD(g)
	img := ebiten.NewImage(1920, 1080)
	hud.drawTowerSelectionOverlay(img)

	bx, by, _, _ := g.towers[0].Bounds()
	clr := color.RGBAModel.Convert(img.At(bx-4, by-4)).(color.RGBA)
	if clr.A == 0 {
		t.Fatalf("expected overlay pixel at tower bounds")
	}
}
