package building

import (
	"testing"

	"github.com/daddevv/type-defense/internal/assets"
	"github.com/daddevv/type-defense/internal/word"
)

func TestBarracksCooldownAndWordGeneration(t *testing.T) {
	assets.InitImages()
	b := NewBarracks()
	b.SetLetterPool([]rune{'f', 'j'}) // Set letter pool for word generation
	b.SetInterval(0.1)
	b.SetCooldown(0.1)
	words := make(map[string]struct{})
	for i := 0; i < 10; i++ {
		w := b.Update(0.11)
		if w == "" {
			t.Fatalf("expected word on cooldown expiry")
		}
		if len(w) < b.WordLenMin || len(w) > b.WordLenMax {
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

func TestBarracksWaitsForCompletion(t *testing.T) {
	assets.InitImages()
	b := NewBarracks()
	b.SetLetterPool([]rune{'f', 'j'}) // Set letter pool for word generation
	b.SetInterval(0.1)
	b.SetCooldown(0.1)
	first := b.Update(0.11)
	if first == "" {
		t.Fatalf("expected word on cooldown expiry")
	}
	if next := b.Update(0.11); next != "" {
		t.Fatalf("expected no new word until completion")
	}
	if w := b.Update(0.11); w == "" {
		t.Fatalf("expected new word after completion")
	}
}

func TestBarracksLetterQueueIntegration(t *testing.T) {
	assets.InitImages()
	q := word.NewQueueManager()
	b := NewBarracks()
	b.SetQueue(q)
	b.SetLetterPool([]rune{'f', 'j'}) // Set letter pool for word generation
	b.SetInterval(0.1)
	b.SetCooldown(0.1)

	word := b.Update(0.11)
	if q.Len() != 1 {
		t.Fatalf("expected word enqueued")
	}

	for i, r := range word {
		match, done, _ := q.TryLetter(r)
		if !match {
			t.Fatalf("letter %d did not match", i)
		}
		if i < len(word)-1 && done {
			t.Fatalf("word completed too early")
		}
	}

	if q.Len() != 0 {
		t.Fatalf("queue should be empty after completion")
	}
}
