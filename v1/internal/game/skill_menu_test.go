//go:build test

package game

import (
	"testing"
	"time"
)

type skillInput struct {
	toggle bool
	up     bool
	down   bool
	left   bool
	right  bool
}

func (s *skillInput) TypedChars() []rune { return nil }
func (s *skillInput) Update()            {}
func (s *skillInput) Reset() {
	s.toggle, s.up, s.down, s.left, s.right = false, false, false, false, false
}
func (s *skillInput) Backspace() bool   { return false }
func (s *skillInput) Space() bool       { return false }
func (s *skillInput) Quit() bool        { return false }
func (s *skillInput) Reload() bool      { return false }
func (s *skillInput) Enter() bool       { return false }
func (s *skillInput) Left() bool        { v := s.left; s.left = false; return v }
func (s *skillInput) Right() bool       { v := s.right; s.right = false; return v }
func (s *skillInput) Up() bool          { v := s.up; s.up = false; return v }
func (s *skillInput) Down() bool        { v := s.down; s.down = false; return v }
func (s *skillInput) Build() bool       { return false }
func (s *skillInput) Save() bool        { return false }
func (s *skillInput) Load() bool        { return false }
func (s *skillInput) SelectTower() bool { return false }
func (s *skillInput) Command() bool     { return false }
func (s *skillInput) TechMenu() bool    { return false }
func (s *skillInput) SkillMenu() bool   { v := s.toggle; s.toggle = false; return v }

func TestSkillMenuToggle(t *testing.T) {
	g := NewGame()
	inp := &skillInput{toggle: true}
	g.input = inp
	g.lastUpdate = time.Now()
	if err := g.Update(); err != nil {
		t.Fatal(err)
	}
	if !g.skillMenuOpen {
		t.Fatalf("expected skill menu open")
	}
	inp.toggle = true
	if err := g.Update(); err != nil {
		t.Fatal(err)
	}
	if g.skillMenuOpen {
		t.Fatalf("expected skill menu closed")
	}
}
