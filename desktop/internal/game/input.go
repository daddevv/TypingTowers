package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type InputState struct {
	MouseX, MouseY int
	MouseButton1   bool
	MouseButton2   bool
	MouseButton3   bool
	KeyboardKeys   map[ebiten.Key]bool
	Up             bool
	Down           bool
	Left           bool
	Right          bool
	Enter          bool
	Escape         bool
	Space          bool
}

func NewUserInput() *InputState {
	return &InputState{
		MouseX:        0,
		MouseY:        0,
		MouseButton1:  false,
		MouseButton2:  false,
		MouseButton3:  false,
		KeyboardKeys: make(map[ebiten.Key]bool),
		Up:            false,
		Down:          false,
		Left:          false,
		Right:         false,
		Enter:         false,
		Escape:        false,
		Space:         false,
	}
}

func (ui *InputState) Update() {
	// Get mouse position
	ui.MouseX, ui.MouseY = ebiten.CursorPosition()

	// Get mouse buttons
	ui.MouseButton1 = ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
	ui.MouseButton2 = ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight)
	ui.MouseButton3 = ebiten.IsMouseButtonPressed(ebiten.MouseButtonMiddle)

	// Get keyboard keys
	for key := range ui.KeyboardKeys {
		ui.KeyboardKeys[key] = ebiten.IsKeyPressed(key)
	}

	// Update specific keys
	ui.Up = ebiten.IsKeyPressed(ebiten.KeyArrowUp)
	ui.Down = ebiten.IsKeyPressed(ebiten.KeyArrowDown)
	ui.Left = ebiten.IsKeyPressed(ebiten.KeyArrowLeft)
	ui.Right = ebiten.IsKeyPressed(ebiten.KeyArrowRight)
	ui.Enter = ebiten.IsKeyPressed(ebiten.KeyEnter)
	ui.Escape = ebiten.IsKeyPressed(ebiten.KeyEscape)
	ui.Space = ebiten.IsKeyPressed(ebiten.KeySpace)
	// Add more keys as needed
}
