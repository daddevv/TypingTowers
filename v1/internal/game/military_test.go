package game

import (
	"github.com/daddevv/type-defense/internal/entity"
	"testing"
)

func TestMilitaryAddAndCount(t *testing.T) {
	m := NewMilitary()
	if m.Count() != 0 {
		t.Fatalf("expected empty military")
	}
	m.AddUnit(entity.NewFootman(0, 0))
	if m.Count() != 1 {
		t.Fatalf("expected 1 unit got %d", m.Count())
	}
}

func TestBarracksAddsUnitToMilitary(t *testing.T) {
	b := NewBarracks()
	m := NewMilitary()
	b.SetMilitary(m)
	word := b.generateWord()
	b.pendingWord = word
	unit := b.OnWordCompleted(word)
	if unit == nil {
		t.Fatalf("expected footman spawn")
	}
	if m.Count() != 1 {
		t.Fatalf("military should track spawned unit")
	}
}
