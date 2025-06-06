package game

import (
	"testing"

	"github.com/daddevv/type-defense/internal/econ"
)

func TestMinerCooldownAndWordGeneration(t *testing.T) {
	m := NewMiner()
	m.SetLetterPool([]rune{'f', 'j'})
	m.SetInterval(0.1)
	m.SetCooldown(0.1)
	words := make(map[string]struct{})
	for i := 0; i < 10; i++ {
		w := m.Update(0.11)
		if w == "" {
			t.Fatalf("expected word on cooldown expiry")
		}
		if len(w) < m.wordLenMin || len(w) > m.wordLenMax {
			t.Errorf("word length out of bounds: %s", w)
		}
		words[w] = struct{}{}
		m.OnWordCompleted(w, nil)
	}
	if len(words) < 2 {
		t.Errorf("expected multiple unique words")
	}
}

func TestMinerResourceOutput(t *testing.T) {
	m := NewMiner()
	word := m.generateWord()
	m.pendingWord = word
	pool := &econ.ResourcePool{}
	stone, iron := m.OnWordCompleted(word, pool)
	if stone != m.stoneOut || iron != m.ironOut {
		t.Fatalf("unexpected resource amounts")
	}
	if pool.StoneAmount() != stone || pool.IronAmount() != iron {
		t.Fatalf("resources not added")
	}
}
