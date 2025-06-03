package game

import "github.com/hajimehoshi/ebiten/v2"

type InputHandler interface {
	Update()    // Update processes input events and updates the Input state
	Reset()     // Reset resets the Input state to its default values
	Quit() bool // Quit returns whether the game should quit
}

type Input struct {
	quit bool // Whether the game should quit
}

// NewInput creates a new Input instance with default values.
func NewInput() *Input {
	return &Input{
		quit: false, // Default to not quitting
	}
}

// Update processes input events and updates the Input state.
func (i *Input) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		i.quit = true
	}
}

// Reset resets the Input state to its default values.
func (i *Input) Reset() {
	i.quit = false // Reset quit state
}

// Quit returns whether the game should quit.
func (i *Input) Quit() bool {
	return i.quit
}
