package game

import (
	"testing"
	"time"

	"github.com/daddevv/type-defense/internal/core"
)

type pauseInput struct {
	space bool
	enter bool
}

func (p *pauseInput) TypedChars() []rune { return nil }
func (p *pauseInput) Update()            {}
func (p *pauseInput) Reset()             { p.space, p.enter = false, false }
func (p *pauseInput) Backspace() bool    { return false }
func (p *pauseInput) Space() bool        { v := p.space; p.space = false; return v }
func (p *pauseInput) Quit() bool         { return false }
func (p *pauseInput) Reload() bool       { return false }
func (p *pauseInput) Enter() bool        { v := p.enter; p.enter = false; return v }
func (p *pauseInput) Left() bool         { return false }
func (p *pauseInput) Right() bool        { return false }
func (p *pauseInput) Up() bool           { return false }
func (p *pauseInput) Down() bool         { return false }
func (p *pauseInput) Build() bool        { return false }
func (p *pauseInput) Save() bool         { return false }
func (p *pauseInput) Load() bool         { return false }
func (p *pauseInput) SelectTower() bool  { return false }
func (p *pauseInput) Command() bool      { return false }
func (p *pauseInput) TechMenu() bool     { return false }
func (p *pauseInput) SkillMenu() bool    { return false }
func (p *pauseInput) StatsPanel() bool   { return false }

func TestPauseResumeTransition(t *testing.T) {
	g := NewGame()
	g.phase = core.PhasePlaying
	inp := &pauseInput{space: true}
	g.input = inp
	g.lastUpdate = time.Now()
	if err := g.Update(); err != nil {
		t.Fatal(err)
	}
	if g.phase != core.PhasePaused {
		t.Fatalf("expected PhasePaused got %v", g.phase)
	}
	g.pauseCursor = 0
	inp.enter = true
	if err := g.Update(); err != nil {
		t.Fatal(err)
	}
	if g.phase != core.PhasePlaying {
		t.Fatalf("expected PhasePlaying got %v", g.phase)
	}
}
