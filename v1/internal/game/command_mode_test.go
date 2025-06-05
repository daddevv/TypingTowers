//go:build test

package game

import (
	"testing"
	"time"
)

type cmdInput struct {
	typed   []rune
	enter   bool
	command bool
}

func (c *cmdInput) TypedChars() []rune { ch := c.typed; c.typed = nil; return ch }
func (c *cmdInput) Update()            {}
func (c *cmdInput) Reset()             { c.typed = nil; c.enter = false; c.command = false }
func (c *cmdInput) Backspace() bool    { return false }
func (c *cmdInput) Space() bool        { return false }
func (c *cmdInput) Quit() bool         { return false }
func (c *cmdInput) Reload() bool       { return false }
func (c *cmdInput) Enter() bool        { return c.enter }
func (c *cmdInput) Left() bool         { return false }
func (c *cmdInput) Right() bool        { return false }
func (c *cmdInput) Up() bool           { return false }
func (c *cmdInput) Down() bool         { return false }
func (c *cmdInput) Build() bool        { return false }
func (c *cmdInput) Save() bool         { return false }
func (c *cmdInput) Load() bool         { return false }
func (c *cmdInput) SelectTower() bool  { return false }
func (c *cmdInput) Command() bool      { v := c.command; c.command = false; return v }
func (c *cmdInput) TechMenu() bool     { return false }

func TestEnterCommandMode(t *testing.T) {
	g := NewGame()
	inp := &cmdInput{command: true}
	g.input = inp
	g.lastUpdate = time.Now()
	if err := g.Update(); err != nil {
		t.Fatal(err)
	}
	if !g.commandMode {
		t.Fatalf("expected command mode active")
	}
	inp.typed = []rune{'p', 'a', 'u', 's', 'e'}
	inp.enter = true
	if err := g.Update(); err != nil {
		t.Fatal(err)
	}
	if !g.paused {
		t.Fatalf("expected command executed")
	}
	if g.commandMode {
		t.Fatalf("expected command mode exit")
	}
}
