//go:build test

package game

import "testing"

// stubInput provides deterministic input for tests.
type stubInput struct {
	typed []rune
}

func (s *stubInput) TypedChars() []rune { return s.typed }
func (s *stubInput) Update()            {}
func (s *stubInput) Reset()             { s.typed = nil }
func (s *stubInput) Backspace() bool    { return false }
func (s *stubInput) Space() bool        { return false }
func (s *stubInput) Quit() bool         { return false }
func (s *stubInput) Reload() bool       { return false }
func (s *stubInput) Enter() bool        { return false }
func (s *stubInput) Left() bool         { return false }
func (s *stubInput) Right() bool        { return false }
func (s *stubInput) Up() bool           { return false }
func (s *stubInput) Down() bool         { return false }
func (s *stubInput) Build() bool        { return false }
func (s *stubInput) Save() bool         { return false }
func (s *stubInput) Load() bool         { return false }

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
	if g.base.Health() != BaseStartingHealth {
		t.Errorf("base should not take damage, hp=%d", g.base.Health())
	}
	if g.military.Count() == 0 {
		t.Errorf("expected units to spawn")
	}
	if g.queueJam {
		t.Errorf("did not expect jam state")
	}
}
