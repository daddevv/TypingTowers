package game

import "testing"

func TestBarracksCooldownAndWordGeneration(t *testing.T) {
	b := NewBarracks()
	b.SetLetterPool([]rune{'f', 'j'})
	b.SetInterval(0.1)
	b.SetCooldown(0.1)
	words := make(map[string]struct{})
	for i := 0; i < 10; i++ {
		w := b.Update(0.11)
		if w == "" {
			t.Fatalf("expected word on cooldown expiry")
		}
		if len(w) < b.wordLenMin || len(w) > b.wordLenMax {
			t.Errorf("word length out of bounds: %s", w)
		}
		for _, r := range w {
			if r != 'f' && r != 'j' {
				t.Errorf("unexpected letter %c", r)
			}
		}
		words[w] = struct{}{}
		b.OnWordCompleted(w)
	}
	if len(words) < 2 {
		t.Errorf("expected at least 2 unique words, got %d", len(words))
	}
}

func TestBarracksWaitsForCompletion(t *testing.T) {
	b := NewBarracks()
	b.SetInterval(0.1)
	b.SetCooldown(0.1)
	first := b.Update(0.11)
	if first == "" {
		t.Fatalf("expected word on cooldown expiry")
	}
	if next := b.Update(0.11); next != "" {
		t.Fatalf("expected no new word until completion")
	}
	b.OnWordCompleted(first)
	if w := b.Update(0.11); w == "" {
		t.Fatalf("expected new word after completion")
	}
}

func TestBarracksUnitSpawn(t *testing.T) {
	b := NewBarracks()
	word := b.generateWord()
	b.pendingWord = word
	unit := b.OnWordCompleted(word)
	if unit == nil {
		t.Fatalf("expected Footman spawn")
	}
	if b.OnWordCompleted("bad") != nil {
		t.Errorf("unexpected spawn for wrong word")
	}
}
