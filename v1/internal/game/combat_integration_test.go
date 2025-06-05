//go:build test

package game

import "testing"

// TestCombatKillTimeUnderEight simulates perfect typing resulting in a Footman
// defeating an OrcGrunt. The battle should resolve in under eight seconds of
// simulated time.
func TestCombatKillTimeUnderEight(t *testing.T) {
    f := NewFootman(0, 0)
    f.speed = 0
    o := NewOrcGrunt(0, 0)
    o.speed = 0

    m := NewMilitary()
    m.AddUnit(f)

    dt := 0.1
    elapsed := 0.0
    for o.Alive() && elapsed < 10 {
        o.Update(dt)
        m.Update(dt, []*OrcGrunt{o})
        elapsed += dt
    }

    if o.Alive() {
        t.Fatalf("orc grunt should be defeated")
    }
    if elapsed >= 8.0 {
        t.Fatalf("expected kill in under 8s, got %.1fs", elapsed)
    }
}

