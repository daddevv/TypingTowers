//go:build test

package game

import (
	"testing"
	"time"
)

// menuInput implements InputHandler for menu navigation tests.
type menuInput struct {
	up, down, enter bool
}

func (m *menuInput) TypedChars() []rune { return nil }
func (m *menuInput) Update()            {}
func (m *menuInput) Reset()             { m.up, m.down, m.enter = false, false, false }
func (m *menuInput) Backspace() bool    { return false }
func (m *menuInput) Space() bool        { return false }
func (m *menuInput) Quit() bool         { return false }
func (m *menuInput) Reload() bool       { return false }
func (m *menuInput) Enter() bool        { return m.enter }
func (m *menuInput) Left() bool         { return false }
func (m *menuInput) Right() bool        { return false }
func (m *menuInput) Up() bool           { return m.up }
func (m *menuInput) Down() bool         { return m.down }
func (m *menuInput) Build() bool        { return false }
func (m *menuInput) Save() bool         { return false }
func (m *menuInput) Load() bool         { return false }
func (m *menuInput) SelectTower() bool  { return false }
func (m *menuInput) Command() bool      { return false }
func (m *menuInput) TechMenu() bool     { return false }
func (m *menuInput) SkillMenu() bool    { return false }

func TestMainMenuStartGame(t *testing.T) {
	g := NewGame()
	g.phase = PhaseMainMenu
	g.mainMenu = NewMainMenu()
	inp := &menuInput{enter: true}
	g.input = inp
	g.lastUpdate = time.Now()
	if err := g.Update(); err != nil {
		t.Fatal(err)
	}
	if g.phase != PhasePlaying {
		t.Fatalf("expected PhasePlaying, got %v", g.phase)
	}
}

func TestMainMenuCursorWrap(t *testing.T) {
	g := NewGame()
	g.phase = PhaseMainMenu
	g.mainMenu = NewMainMenu()
	inp := &menuInput{up: true}
	g.input = inp
	g.lastUpdate = time.Now()
	if err := g.Update(); err != nil {
		t.Fatal(err)
	}
	if g.mainMenu.cursor != len(g.mainMenu.options)-1 {
		t.Fatalf("expected cursor wrap, got %d", g.mainMenu.cursor)
	}
}

func TestMainMenuSettingsToggle(t *testing.T) {
	g := NewGame()
	g.phase = PhaseMainMenu
	g.mainMenu = NewMainMenu()
	g.mainMenu.cursor = 1
	inp := &menuInput{enter: true}
	g.input = inp
	g.lastUpdate = time.Now()
	if err := g.Update(); err != nil {
		t.Fatal(err)
	}
	if !g.mainMenu.inSettings {
		t.Fatalf("expected settings menu open")
	}
	inp.enter = true
	g.mainMenu.settingsCursor = 0
	if err := g.Update(); err != nil {
		t.Fatal(err)
	}
	if !g.settings.Mute {
		t.Fatalf("expected mute toggled")
	}
}
