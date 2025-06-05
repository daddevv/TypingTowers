package game

import "testing"

func TestFarmerCooldownAndWordGeneration(t *testing.T) {
	f := NewFarmer()
	f.SetLetterPool([]rune{'f', 'j'})
	f.SetInterval(0.1)
	f.SetCooldown(0.1)
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
		// simulate typing completion to reset cooldown
		f.OnWordCompleted(word, nil)
	}
	if len(words) < 2 {
		t.Errorf("Expected at least 2 unique words, got %d", len(words))
	}
}

func TestFarmerWaitsForCompletion(t *testing.T) {
	f := NewFarmer()
	f.SetInterval(0.1)
	f.SetCooldown(0.1)
	first := f.Update(0.11)
	if first == "" {
		t.Fatalf("expected word on cooldown expiry")
	}
	if next := f.Update(0.11); next != "" {
		t.Fatalf("expected no new word until completion")
	}
	f.OnWordCompleted(first, nil)
	if w := f.Update(0.11); w == "" {
		t.Fatalf("expected new word after completion")
	}
}

func TestFarmerResourceOutput(t *testing.T) {
	f := NewFarmer()
	word := f.generateWord()
	f.pendingWord = word
	pool := &ResourcePool{}
	food := f.OnWordCompleted(word, pool)
	if food != f.resourceOut {
		t.Errorf("Expected %d food, got %d", f.resourceOut, food)
	}
	if pool.FoodAmount() != f.resourceOut || pool.GoldAmount() != f.resourceOut {
		t.Errorf("resources not added to pool")
	}
	food = f.OnWordCompleted("wrong", pool)
	if food != 0 {
		t.Errorf("Expected 0 food for wrong word, got %d", food)
	}
}

func TestFarmerAddsResourcesToPool(t *testing.T) {
	f := NewFarmer()
	pool := &ResourcePool{}
	word := f.Update(2.0)
	if word == "" {
		t.Fatalf("expected word generated")
	}
	f.OnWordCompleted(word, pool)
	if pool.GoldAmount() != f.resourceOut || pool.FoodAmount() != f.resourceOut {
		t.Fatalf("expected pool to have %d resources", f.resourceOut)
	}
}
