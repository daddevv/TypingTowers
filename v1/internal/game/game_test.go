package game

import (
	"testing"
	"time"
)

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
	tree := DefaultTechTree()
	firstLetters, _, _ := tree.UnlockNext()
	if len(g.letterPool) != len(firstLetters) {
		t.Fatalf("expected initial letter pool %d got %d", len(firstLetters), len(g.letterPool))
	}

	g.currentWave = 2
	g.startWave()
	tree = DefaultTechTree()
	tree.UnlockNext() // first stage
	secondLetters, _, _ := tree.UnlockNext()
	expected := len(firstLetters) + len(secondLetters)
	if len(g.letterPool) != expected {
		t.Errorf("expected letter pool size %d after second wave got %d", expected, len(g.letterPool))
	}
}

func TestGameBackPressureDamage(t *testing.T) {
	g := NewGame()
	for i := 0; i < 6; i++ {
		g.Queue().Enqueue(Word{Text: "w"})
	}
	g.lastUpdate = time.Now().Add(-1 * time.Second)
	g.Update()
	expected := BaseStartingHealth - 1
	if g.base.Health() != expected {
		t.Fatalf("expected base health %d got %d", expected, g.base.Health())
	}
}
