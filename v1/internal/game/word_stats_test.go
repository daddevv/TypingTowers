package game

import (
	"testing"

	"github.com/daddevv/type-defense/internal/assets"
	"github.com/daddevv/type-defense/internal/core"
)

func TestWordStatsRecording(t *testing.T) {
	g := NewGame()
	g.phase = core.PhasePlaying
	s := &mockInput{}
	g.input = s

	g.queue.Enqueue(assets.Word{Text: "ab", Source: "Farmer"})

	s.typed = []rune{'a'}
	g.Update()
	s.typed = []rune{'x'}
	g.Update()
	if !g.queueJam {
		t.Fatalf("expected jam on wrong letter")
	}
	s.backspace = true
	g.Update()
	s.typed = []rune{'a'}
	g.Update()
	s.typed = []rune{'b'}
	g.Update()

	if len(g.wordHistory) != 1 {
		t.Fatalf("expected 1 word stat got %d", len(g.wordHistory))
	}
	ws := g.wordHistory[0]
	if ws.Text != "ab" || ws.Correct != 3 || ws.Incorrect != 1 {
		t.Fatalf("unexpected word stat %+v", ws)
	}
	if ws.Duration <= 0 {
		t.Errorf("expected duration recorded")
	}
}
