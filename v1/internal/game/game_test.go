package game

import "testing"

func init() {
	TESTING = true
}

func TestNewGame(t *testing.T) {
	g := NewGame()
	if g.screen.Bounds().Dx() != 1920 || g.screen.Bounds().Dy() != 1080 {
		t.Errorf("screen size expected 1920x1080 got %dx%d", g.screen.Bounds().Dx(), g.screen.Bounds().Dy())
	}
	if g.base == nil || g.base.Health() <= 0 {
		t.Errorf("base not initialized")
	}
}
