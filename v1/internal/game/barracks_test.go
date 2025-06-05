package game

import "testing"

func TestBarracksCooldownAndWordGeneration(t *testing.T) {
	b := NewBarracks()
	b.SetLetterPool([]rune{'f', 'j'})
	b.interval = 0.1
	b.cooldown = 0.1
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
	}
	if len(words) < 2 {
		t.Errorf("expected at least 2 unique words, got %d", len(words))
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
