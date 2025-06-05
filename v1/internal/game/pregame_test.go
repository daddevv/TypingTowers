//go:build test

package game

import (
	"testing"
	"time"
)

// pgInput implements InputHandler for pregame tests.
type pgInput struct {
	up, down, enter bool
	chars           []rune
}

func (p *pgInput) TypedChars() []rune { c := p.chars; p.chars = nil; return c }
func (p *pgInput) Update()            {}
func (p *pgInput) Reset()             { p.up, p.down, p.enter = false, false, false; p.chars = nil }
func (p *pgInput) Backspace() bool    { return false }
func (p *pgInput) Space() bool        { return false }
func (p *pgInput) Quit() bool         { return false }
func (p *pgInput) Reload() bool       { return false }
func (p *pgInput) Enter() bool        { return p.enter }
func (p *pgInput) Left() bool         { return false }
func (p *pgInput) Right() bool        { return false }
func (p *pgInput) Up() bool           { return p.up }
func (p *pgInput) Down() bool         { return p.down }
func (p *pgInput) Build() bool        { return false }
func (p *pgInput) Save() bool         { return false }
func (p *pgInput) Load() bool         { return false }

// TestPreGameFlow ensures the setup screens progress to playing state.
func TestPreGameFlow(t *testing.T) {
	g := NewGame()
	g.phase = PhasePreGame
	g.preGame = NewPreGame()
	inp := &pgInput{enter: true}
	g.input = inp
	g.lastUpdate = time.Now()

	// step 0 -> 1
	if err := g.Update(); err != nil {
		t.Fatal(err)
	}
	if g.preGame.step != 1 {
		t.Fatalf("expected step 1 got %d", g.preGame.step)
	}
	// step1 -> 2
	inp.enter = true
	if err := g.Update(); err != nil {
		t.Fatal(err)
	}
	if g.preGame.step != 2 {
		t.Fatalf("expected step 2 got %d", g.preGame.step)
	}
	// step2 -> 3
	inp.enter = true
	if err := g.Update(); err != nil {
		t.Fatal(err)
	}
	if g.preGame.step != 3 {
		t.Fatalf("expected step 3 got %d", g.preGame.step)
	}
	// typing test
	inp.chars = []rune{'r', 'e', 'a', 'd', 'y'}
	if err := g.Update(); err != nil {
		t.Fatal(err)
	}
	if g.preGame.step != 4 {
		t.Fatalf("expected step 4 got %d", g.preGame.step)
	}
	// final enter -> playing
	inp.enter = true
	if err := g.Update(); err != nil {
		t.Fatal(err)
	}
	if g.phase != PhasePlaying {
		t.Fatalf("expected PhasePlaying got %v", g.phase)
	}
}

// TestPreGameCursorWrap checks selection cursor wrapping.
func TestPreGameCursorWrap(t *testing.T) {
	g := NewGame()
	g.phase = PhasePreGame
	g.preGame = NewPreGame()
	inp := &pgInput{up: true}
	g.input = inp
	g.lastUpdate = time.Now()
	if err := g.Update(); err != nil {
		t.Fatal(err)
	}
	if g.preGame.charCursor != len(g.preGame.charOptions)-1 {
		t.Fatalf("expected wrap got %d", g.preGame.charCursor)
	}
}
