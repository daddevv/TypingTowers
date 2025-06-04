package game

import "testing"

func TestNewGame(t *testing.T) {
	g := NewGame()
	if g.screen.Bounds().Dx() != 1920 || g.screen.Bounds().Dy() != 1080 {
		t.Errorf("screen size expected 1920x1080 got %dx%d", g.screen.Bounds().Dx(), g.screen.Bounds().Dy())
	}
	if g.base == nil || g.base.Health() <= 0 {
		t.Errorf("base not initialized")
	}
}

func TestLetterUnlocking(t *testing.T) {
	g := NewGameWithConfig(DefaultConfig)
	if len(g.letterPool) != len(letterUnlockSequence[0]) {
		t.Fatalf("expected initial letter pool %d got %d", len(letterUnlockSequence[0]), len(g.letterPool))
	}

	g.currentWave = 2
	g.startWave()
	expected := len(letterUnlockSequence[0]) + len(letterUnlockSequence[1])
	if len(g.letterPool) != expected {
		t.Errorf("expected letter pool size %d after second wave got %d", expected, len(g.letterPool))
	}
}
