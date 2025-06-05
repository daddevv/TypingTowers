package game

import "testing"

func TestFarmerCooldownAndWordGeneration(t *testing.T) {
	f := NewFarmer()
	f.SetLetterPool([]rune{'f', 'j'})
	f.interval = 0.1
	f.cooldown = 0.1
	words := make(map[string]struct{})
	for i := 0; i < 10; i++ {
		word := f.Update(0.11)
		if word == "" {
			t.Errorf("Expected word on cooldown expiry, got empty string")
		}
		if len(word) < f.wordLenMin || len(word) > f.wordLenMax {
			t.Errorf("Word length out of bounds: %s", word)
		}
		for _, r := range word {
			if r != 'f' && r != 'j' {
				t.Errorf("Unexpected letter in word: %c", r)
			}
		}
		words[word] = struct{}{}
	}
	if len(words) < 2 {
		t.Errorf("Expected at least 2 unique words, got %d", len(words))
	}
}

func TestFarmerResourceOutput(t *testing.T) {
	f := NewFarmer()
	word := f.generateWord()
	f.pendingWord = word
	food := f.OnWordCompleted(word)
	if food != f.resourceOut {
		t.Errorf("Expected %d food, got %d", f.resourceOut, food)
	}
	food = f.OnWordCompleted("wrong")
	if food != 0 {
		t.Errorf("Expected 0 food for wrong word, got %d", food)
	}
}
