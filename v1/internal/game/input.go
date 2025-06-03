package game

import "github.com/hajimehoshi/ebiten/v2"

type InputHandler interface {
	Update()    // Update processes input events and updates the Input state
	Reset()     // Reset resets the Input state to its default values
	Quit() bool // Quit returns whether the game should quit
}

type Input struct {
	quit  bool   // Whether the game should quit
	typed []rune // Characters typed this frame
}

// NewInput creates a new Input instance with default values.
func NewInput() *Input {
	return &Input{
		quit:  false, // Default to not quitting
		typed: nil,
	}
}

// Update processes input events and updates the Input state.
func (i *Input) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		i.quit = true
	}
	i.typed = ebiten.AppendInputChars(i.typed[:0])
}

// Reset resets the Input state to its default values.
func (i *Input) Reset() {
	i.quit = false // Reset quit state
	i.typed = i.typed[:0]
}

// Quit returns whether the game should quit.
func (i *Input) Quit() bool {
	return i.quit
}

// TypedChars returns any characters typed since the last Update call.
func (i *Input) TypedChars() []rune {
	return i.typed
}
