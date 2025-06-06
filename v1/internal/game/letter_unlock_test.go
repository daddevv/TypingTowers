package game

import (
	"testing"

	"github.com/daddevv/type-defense/internal/econ"
)

func TestLetterUnlockDeductsPoints(t *testing.T) {
	f := NewFarmer()
	pool := &econ.ResourcePool{}
	pool.AddKingsPoints(100)
	cost := f.NextUnlockCost()
	if !f.UnlockNext(pool) {
		t.Fatalf("unlock should succeed")
	}
	if pool.KingsAmount() != 100-cost {
		t.Fatalf("expected %d KP remaining got %d", 100-cost, pool.KingsAmount())
	}
	if len(f.letterPool) <= 2 {
		t.Fatalf("letters not added")
	}
}
