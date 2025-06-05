//go:build test

package game

import "testing"

// TestResourcesAccumulation simulates several minutes to ensure resources grow.
func TestResourcesAccumulation(t *testing.T) {
	g := NewGame()
	inp := &stubInput{}
	g.input = inp

	steps := 1800 // 180 seconds at 0.1 dt
	dt := 0.1
	for i := 0; i < steps; i++ {
		if w, ok := g.Queue().Peek(); ok {
			inp.typed = []rune(w.Text)
		}
		if err := g.Step(dt); err != nil {
			t.Fatal(err)
		}
	}

	if g.resources.GoldAmount() == 0 || g.resources.WoodAmount() == 0 || g.resources.StoneAmount() == 0 || g.resources.IronAmount() == 0 {
		t.Fatalf("expected resources to accumulate")
	}
}
