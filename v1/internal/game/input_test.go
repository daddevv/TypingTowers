package game

import "testing"

func TestInputReset(t *testing.T) {
	in := NewInput()
	in.quit = true
	in.Reset()
	if in.quit != false {
		t.Errorf("expected quit false got %v", in.quit)
	}
}

func TestInputQuit(t *testing.T) {
	in := NewInput()
	if in.Quit() {
		t.Errorf("expected default quit false")
	}
	in.quit = true
	if !in.Quit() {
		t.Errorf("expected quit true after set")
	}
}
