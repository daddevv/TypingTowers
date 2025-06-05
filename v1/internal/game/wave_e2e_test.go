//go:build test

package game

import (
	"testing"
	"time"
)

type waveInput struct {
	typed   []rune
	enter   bool
	command bool
}

func (w *waveInput) TypedChars() []rune { return w.typed }
func (w *waveInput) Update()            {}
func (w *waveInput) Reset()             { w.typed = nil; w.enter = false; w.command = false }
func (w *waveInput) Backspace() bool    { return false }
func (w *waveInput) Space() bool        { return false }
func (w *waveInput) Quit() bool         { return false }
func (w *waveInput) Reload() bool       { return false }
func (w *waveInput) Enter() bool        { return w.enter }
func (w *waveInput) Left() bool         { return false }
func (w *waveInput) Right() bool        { return false }
func (w *waveInput) Up() bool           { return false }
func (w *waveInput) Down() bool         { return false }
func (w *waveInput) Build() bool        { return false }
func (w *waveInput) Save() bool         { return false }
func (w *waveInput) Load() bool         { return false }
func (w *waveInput) SelectTower() bool  { return false }
func (w *waveInput) Command() bool      { c := w.command; w.command = false; return c }
func (w *waveInput) TechMenu() bool     { return false }
func (w *waveInput) SkillMenu() bool    { return false }

// TestSurviveFiveWaves simulates five waves with perfect typing input.
func TestSurviveFiveWaves(t *testing.T) {
	g := NewGame()
	inp := &waveInput{}
	g.input = inp

	// speed up wave spawning for deterministic test
	g.cfg.SpawnInterval = 0.1
	g.spawnInterval = g.cfg.SpawnInterval

	dt := 0.1
	completed := 0

	for steps := 0; steps < 2000 && completed < 5; steps++ {
		g.lastUpdate = time.Now().Add(-time.Duration(float64(time.Second) * dt))

		if g.shopOpen {
			completed++
			if completed >= 5 {
				break
			}
			inp.enter = true
			if err := g.Update(); err != nil {
				t.Fatal(err)
			}
			inp.enter = false
			continue
		}

		if w, ok := g.Queue().Peek(); ok && !g.queueJam {
			idx := g.Queue().Index()
			inp.typed = []rune{rune(w.Text[idx])}
		}

		if err := g.Update(); err != nil {
			t.Fatal(err)
		}
	}

	if completed < 5 {
		t.Fatalf("expected to complete 5 waves, got %d", completed)
	}
	if !g.base.Alive() {
		t.Errorf("base destroyed before wave completion")
	}
	if g.resources.GoldAmount() == 0 {
		t.Errorf("expected gold to accumulate")
	}
	if g.typing.Total() == 0 {
		t.Errorf("expected typing stats recorded")
	}
}
