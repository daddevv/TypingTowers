//go:build test

package game

import (
	"testing"

	"github.com/daddevv/type-defense/internal/structure"
)

// TestCoreLoopSim runs the main game loop in headless mode and verifies core
// systems interact as expected.
func TestCoreLoopSim(t *testing.T) {
	g := NewGame()
	inp := &stubInput{}
	g.input = inp

	// Unlock the next letter stage for both buildings to widen pools.
	g.resources.AddKingsPoints(100)
	if !g.farmer.UnlockNext(&g.resources) {
		t.Fatalf("farmer unlock failed")
	}
	if !g.barracks.UnlockNext(&g.resources) {
		t.Fatalf("barracks unlock failed")
	}

	// Simulate ~5 seconds of game time with perfect typing.
	steps := 50
	dt := 0.1
	for i := 0; i < steps; i++ {
		if w, ok := g.Queue().Peek(); ok {
			inp.typed = []rune(w.Text)
		}
		if err := g.Step(dt); err != nil {
			t.Fatal(err)
		}
	}

	if g.Gold() == 0 {
		t.Errorf("expected gold to increase")
	}
	if g.Queue().Len() != 0 {
		t.Errorf("queue should be empty, got %d", g.Queue().Len())
	}
	if g.base.Health() != structure.BaseStartingHealth {
		t.Errorf("base should not take damage, hp=%d", g.base.Health())
	}
	if g.queueJam {
		t.Errorf("did not expect jam state")
	}
}
