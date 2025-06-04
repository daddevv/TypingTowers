package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type InputHandler interface {
	TypedChars() []rune // TypedChars returns any characters typed since the last Update call
	Update()            // Update processes input events and updates the Input state
	Reset()             // Reset resets the Input state to its default values
	Backspace() bool    // Backspace reports if backspace was pressed since the last Update call
	Space() bool        // Space reports if the space bar was pressed since the last Update call
	Quit() bool         // Quit returns whether the game should quit
	Reload() bool       // Reload returns whether config reload was requested
	Enter() bool        // Enter reports if the enter key was pressed
	Left() bool
	Right() bool
	Up() bool
	Down() bool
	Build() bool
	Save() bool
	Load() bool
}

type Input struct {
	quit      bool   // Whether the game should quit
	typed     []rune // Characters typed this frame
	backspace bool   // Whether backspace was pressed this frame
	space     bool   // Whether space was pressed this frame
	reload    bool   // Whether F5 was pressed this frame
	enter     bool   // Whether enter was pressed this frame
	left      bool
	right     bool
	up        bool
	down      bool
	build     bool
	save      bool
	load      bool
}

// NewInput creates a new Input instance with default values.
func NewInput() *Input {
	return &Input{
		quit:      false, // Default to not quitting
		typed:     nil,
		backspace: false,
		space:     false,
		reload:    false,
		enter:     false,
		left:      false,
		right:     false,
		up:        false,
		down:      false,
		build:     false,
		save:      false,
		load:      false,
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
	i.save = inpututil.IsKeyJustPressed(ebiten.KeyF2)
	i.load = inpututil.IsKeyJustPressed(ebiten.KeyF3)
	i.enter = inpututil.IsKeyJustPressed(ebiten.KeyEnter)

	i.left = inpututil.IsKeyJustPressed(ebiten.KeyH) || inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft)
	i.right = inpututil.IsKeyJustPressed(ebiten.KeyL) || inpututil.IsKeyJustPressed(ebiten.KeyArrowRight)
	i.up = inpututil.IsKeyJustPressed(ebiten.KeyK) || inpututil.IsKeyJustPressed(ebiten.KeyArrowUp)
	i.down = inpututil.IsKeyJustPressed(ebiten.KeyJ) || inpututil.IsKeyJustPressed(ebiten.KeyArrowDown)
	i.build = inpututil.IsKeyJustPressed(ebiten.KeyB)
}

// Reset resets the Input state to its default values.
func (i *Input) Reset() {
	i.quit = false // Reset quit state
	i.typed = i.typed[:0]
	i.backspace = false
	i.space = false
	i.reload = false
	i.enter = false
	i.left = false
	i.right = false
	i.up = false
	i.down = false
	i.build = false
	i.save = false
	i.load = false
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

// Enter reports if the enter key was pressed since the last Update call.
func (i *Input) Enter() bool {
	return i.enter
}

func (i *Input) Left() bool  { return i.left }
func (i *Input) Right() bool { return i.right }
func (i *Input) Up() bool    { return i.up }
func (i *Input) Down() bool  { return i.down }
func (i *Input) Build() bool { return i.build }
func (i *Input) Save() bool  { return i.save }
func (i *Input) Load() bool  { return i.load }
