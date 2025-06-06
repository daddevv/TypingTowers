package gatherer

import (
	"testing"

	"github.com/daddevv/type-defense/internal/econ"
)

func TestLumberjackCooldownAndWordGeneration(t *testing.T) {
	l := NewLumberjack()
	l.SetLetterPool([]rune{'f', 'j'})
	l.SetInterval(0.1)
	l.SetCooldown(0.1)
	words := make(map[string]struct{})
	for i := 0; i < 10; i++ {
		w := l.Update(0.11)
		if w == "" {
			t.Fatalf("expected word on cooldown expiry")
		}
		if len(w) < l.WordLenMin || len(w) > l.WordLenMax {
			t.Errorf("word length out of bounds: %s", w)
		}
		words[w] = struct{}{}
		l.OnWordCompleted(w, nil)
	}
	if len(words) < 2 {
		t.Errorf("expected multiple unique words")
	}
}

func TestLumberjackResourceOutput(t *testing.T) {
	l := NewLumberjack()
	word := l.GenerateWord()
	l.PendingWord = word
	pool := &econ.ResourcePool{}
	wood := l.OnWordCompleted(word, pool)
	if wood != l.ResourceOut {
		t.Fatalf("expected %d wood got %d", l.ResourceOut, wood)
	}
	if pool.WoodAmount() != l.ResourceOut || pool.GoldAmount() != l.ResourceOut {
		t.Fatalf("resources not added")
	}
}
