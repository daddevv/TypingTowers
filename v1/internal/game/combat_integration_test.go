//go:build test

package game

import (
	"testing"

	"github.com/daddevv/type-defense/internal/enemy"
)

// TestCombatKillTimeUnderEight simulates perfect typing resulting in a Footman
// defeating an OrcGrunt. The battle should resolve in under eight seconds of
// simulated time.
func TestCombatKillTimeUnderEight(t *testing.T) {
	f := enemy.NewFootman(0, 0)
	f.Speed = 0
	o := enemy.NewOrcGrunt(0, 0)
	o.Speed = 0

	dt := 0.1
	elapsed := 0.0
	for o.Alive && elapsed < 10 {
		o.Update(dt)
		f.Update(dt)
		elapsed += dt
	}

	if o.Alive {
		t.Fatalf("orc grunt should be defeated")
	}
	if elapsed >= 8.0 {
		t.Fatalf("expected kill in under 8s, got %.1fs", elapsed)
	}
}
