//go:build test

package game

import (
	"testing"
	"time"
)

type techInput struct {
	toggle bool
	typed  []rune
	enter  bool
	up     bool
	down   bool
}

func (t *techInput) TypedChars() []rune { ch := t.typed; t.typed = nil; return ch }
func (t *techInput) Update()            {}
func (t *techInput) Reset() {
	t.toggle, t.enter, t.up, t.down = false, false, false, false
	t.typed = nil
}
func (t *techInput) Backspace() bool   { return false }
func (t *techInput) Space() bool       { return false }
func (t *techInput) Quit() bool        { return false }
func (t *techInput) Reload() bool      { return false }
func (t *techInput) Enter() bool       { v := t.enter; t.enter = false; return v }
func (t *techInput) Left() bool        { return false }
func (t *techInput) Right() bool       { return false }
func (t *techInput) Up() bool          { v := t.up; t.up = false; return v }
func (t *techInput) Down() bool        { v := t.down; t.down = false; return v }
func (t *techInput) Build() bool       { return false }
func (t *techInput) Save() bool        { return false }
func (t *techInput) Load() bool        { return false }
func (t *techInput) SelectTower() bool { return false }
func (t *techInput) Command() bool     { return false }
func (t *techInput) TechMenu() bool    { v := t.toggle; t.toggle = false; return v }
func (t *techInput) SkillMenu() bool   { return false }

func TestTechMenuToggle(t *testing.T) {
	g := NewGame()
	inp := &techInput{toggle: true}
	g.input = inp
	g.lastUpdate = time.Now()
	if err := g.Update(); err != nil {
		t.Fatal(err)
	}
	if !g.techMenuOpen {
		t.Fatalf("expected tech menu open")
	}
	inp.toggle = true
	if err := g.Update(); err != nil {
		t.Fatal(err)
	}
	if g.techMenuOpen {
		t.Fatalf("expected tech menu closed")
	}
}

func TestTechMenuPurchase(t *testing.T) {
	g := NewGame()
	inp := &techInput{toggle: true}
	g.input = inp
	g.lastUpdate = time.Now()
	g.Update() // open menu
	inp.enter = true
	if err := g.Update(); err != nil {
		t.Fatal(err)
	}
	if g.techTree.stage != 1 {
		t.Fatalf("expected tech stage 1 got %d", g.techTree.stage)
	}
}
