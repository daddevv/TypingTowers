//go:build test

package game

import (
	"testing"
	"time"

	"github.com/daddevv/type-defense/internal/core"
	"github.com/daddevv/type-defense/internal/word"
	"github.com/hajimehoshi/ebiten/v2"
)

type panelInput struct{ toggle bool }

func (p *panelInput) TypedChars() []rune { return nil }
func (p *panelInput) Update()            {}
func (p *panelInput) Reset()             { p.toggle = false }
func (p *panelInput) Backspace() bool    { return false }
func (p *panelInput) Space() bool        { return false }
func (p *panelInput) Quit() bool         { return false }
func (p *panelInput) Reload() bool       { return false }
func (p *panelInput) Enter() bool        { return false }
func (p *panelInput) Left() bool         { return false }
func (p *panelInput) Right() bool        { return false }
func (p *panelInput) Up() bool           { return false }
func (p *panelInput) Down() bool         { return false }
func (p *panelInput) Build() bool        { return false }
func (p *panelInput) Save() bool         { return false }
func (p *panelInput) Load() bool         { return false }
func (p *panelInput) SelectTower() bool  { return false }
func (p *panelInput) TechMenu() bool     { return false }
func (p *panelInput) SkillMenu() bool    { return false }
func (p *panelInput) Command() bool      { return false }
func (p *panelInput) StatsPanel() bool   { v := p.toggle; p.toggle = false; return v }

func TestStatsPanelToggle(t *testing.T) {
	g := NewGame()
	inp := &panelInput{toggle: true}
	g.input = inp
	g.lastUpdate = time.Now()
	if err := g.Update(); err != nil {
		t.Fatal(err)
	}
	if !g.statsPanelOpen {
		t.Fatalf("expected panel open")
	}
	inp.toggle = true
	if err := g.Update(); err != nil {
		t.Fatal(err)
	}
	if g.statsPanelOpen {
		t.Fatalf("expected panel closed")
	}
}

func TestDrawStatsPanel(t *testing.T) {
	g := NewGame()
	g.statsPanelOpen = true
	g.wordHistory = []word.WordStat{{Text: "ab", Correct: 2, Incorrect: 0, Duration: time.Second}}
	hud := core.NewHUD()
	img := ebiten.NewImage(1920, 1080)
	hud.DrawStatsPanel(img)
}
