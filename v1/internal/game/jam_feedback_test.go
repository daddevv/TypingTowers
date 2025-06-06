package game

import (
	"testing"
	"time"
)

func TestQueueJamMistypeFeedback(t *testing.T) {
	g := NewGame()
	g.phase = PhasePlaying // Ensure main update logic runs
	inp := &stubInput{}
	g.input = inp
	g.Queue().Enqueue(Word{Text: "f"})

	// Simulate a mistype: input 'g' when 'f' is expected
	inp.typed = []rune{'g'}
	g.lastUpdate = time.Now()
	_ = g.Update()

	if !g.queueJam {
		t.Fatalf("expected jam state after mistype")
	}
	if g.flashTimer <= 0 {
		t.Errorf("expected flash timer to be set")
	}

	// Clear jam with backspace
	inp.typed = nil
	inp.backspace = true
	g.lastUpdate = time.Now()
	_ = g.Update()
	if g.queueJam {
		t.Errorf("expected jam cleared after backspace")
	}
}
