package game

import (
	"testing"

	"github.com/daddevv/type-defense/internal/assets"
	"github.com/daddevv/type-defense/internal/core"
)

// stubInput for deterministic input
type stubInputConveyor struct{ typed []rune }

func (s *stubInputConveyor) TypedChars() []rune { return s.typed }
func (s *stubInputConveyor) Update()            {}
func (s *stubInputConveyor) Reset()             { s.typed = nil }
func (s *stubInputConveyor) Backspace() bool    { return false }
func (s *stubInputConveyor) Space() bool        { return false }
func (s *stubInputConveyor) Quit() bool         { return false }
func (s *stubInputConveyor) Reload() bool       { return false }
func (s *stubInputConveyor) Enter() bool        { return false }
func (s *stubInputConveyor) Left() bool         { return false }
func (s *stubInputConveyor) Right() bool        { return false }
func (s *stubInputConveyor) Up() bool           { return false }
func (s *stubInputConveyor) Down() bool         { return false }
func (s *stubInputConveyor) Build() bool        { return false }
func (s *stubInputConveyor) Save() bool         { return false }
func (s *stubInputConveyor) Load() bool         { return false }
func (s *stubInputConveyor) SelectTower() bool  { return false }
func (s *stubInputConveyor) Command() bool      { return false }
func (s *stubInputConveyor) TechMenu() bool     { return false }
func (s *stubInputConveyor) SkillMenu() bool    { return false }
func (s *stubInputConveyor) StatsPanel() bool   { return false }

func TestConveyorOffsetMoves(t *testing.T) {
	g := NewGame()
	g.phase = core.PhasePlaying // Ensure main update logic runs
	inp := &stubInputConveyor{}
	g.input = inp
	g.Queue().Enqueue(assets.Word{Text: "ab"})

	// Ensure queue is not jammed and index is at 0
	g.queueJam = false
	g.Queue().ResetProgress()

	inp.typed = []rune{'a'}
	g.Step(0.1)
	if g.conveyorOffset <= 0 {
		t.Fatalf("expected offset to increase after typing")
	}
	off := g.conveyorOffset

	inp.typed = nil
	// simulate 1 second to let offset decay
	for i := 0; i < 60; i++ {
		g.Step(1.0 / 60.0)
	}
	if g.conveyorOffset >= off {
		t.Fatalf("expected offset to decrease over time")
	}
}
