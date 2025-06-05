package game

import (
	"testing"
	"time"
)

// stubInput implements InputHandler for deterministic tests.
type stubInput struct {
	typed     []rune
	backspace bool
}

func (s *stubInput) TypedChars() []rune { return s.typed }
func (s *stubInput) Update()            {}
func (s *stubInput) Reset()             { s.typed = nil; s.backspace = false }
func (s *stubInput) Backspace() bool    { return s.backspace }
func (s *stubInput) Space() bool        { return false }
func (s *stubInput) Quit() bool         { return false }
func (s *stubInput) Reload() bool       { return false }
func (s *stubInput) Enter() bool        { return false }
func (s *stubInput) Left() bool         { return false }
func (s *stubInput) Right() bool        { return false }
func (s *stubInput) Up() bool           { return false }
func (s *stubInput) Down() bool         { return false }
func (s *stubInput) Build() bool        { return false }
func (s *stubInput) Save() bool         { return false }
func (s *stubInput) Load() bool         { return false }

func TestQueueJamMistypeFeedback(t *testing.T) {
	g := NewGame()
	inp := &stubInput{}
	g.input = inp
	g.Queue().Enqueue(Word{Text: "f"})

	inp.typed = []rune{'g'}
	g.lastUpdate = time.Now()
	if err := g.Update(); err != nil {
		t.Fatal(err)
	}

	if !g.queueJam {
		t.Fatalf("expected jam state after mistype")
	}
	if g.flashTimer <= 0 {
		t.Errorf("expected flash timer to be set")
	}

	// Clear jam with backspace
	inp.backspace = true
	g.lastUpdate = time.Now()
	if err := g.Update(); err != nil {
		t.Fatal(err)
	}
	if g.queueJam {
		t.Errorf("expected jam cleared after backspace")
	}
}
