package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type InputHandler interface {
	TypedChars() []rune // TypedChars returns any characters typed since the last Update call
	Update()            // Update processes input events and updates the Input state
	Reset()             // Reset resets the Input state to its default values
	Quit() bool         // Quit returns whether the game should quit
	Reload() bool       // Reload returns whether config reload was requested
}

type Input struct {
	quit      bool   // Whether the game should quit
	typed     []rune // Characters typed this frame
	backspace bool   // Whether backspace was pressed this frame
	space     bool   // Whether space was pressed this frame
	reload    bool   // Whether F5 was pressed this frame
}

// NewInput creates a new Input instance with default values.
func NewInput() *Input {
	return &Input{
		quit:      false, // Default to not quitting
		typed:     nil,
		backspace: false,
		space:     false,
		reload:    false,
	}
}

// Update processes input events and updates the Input state.
func (i *Input) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		i.quit = true
	}
	i.typed = ebiten.AppendInputChars(i.typed[:0])
	i.backspace = inpututil.IsKeyJustPressed(ebiten.KeyBackspace)
	i.space = inpututil.IsKeyJustPressed(ebiten.KeySpace)
	i.reload = inpututil.IsKeyJustPressed(ebiten.KeyF5)
}

// Reset resets the Input state to its default values.
func (i *Input) Reset() {
	i.quit = false // Reset quit state
	i.typed = i.typed[:0]
	i.backspace = false
	i.space = false
	i.reload = false
}

// Quit returns whether the game should quit.
func (i *Input) Quit() bool {
	return i.quit
}

// TypedChars returns any characters typed since the last Update call.
func (i *Input) TypedChars() []rune {
	return i.typed
}

// Backspace reports if backspace was pressed since the last Update call.
func (i *Input) Backspace() bool {
	return i.backspace
}

// Space reports if the space bar was pressed since the last Update call.
func (i *Input) Space() bool {
	return i.space
}

// Reload reports if the F5 key was pressed since the last Update call.
func (i *Input) Reload() bool {
	return i.reload
}
