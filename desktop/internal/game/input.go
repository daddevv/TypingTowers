package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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
	input := &InputState{
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

	// Get mouse position
	input.MouseX, input.MouseY = ebiten.CursorPosition()

	// Get mouse buttons
	input.MouseButton1 = inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft)
	input.MouseButton2 = inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight)
	input.MouseButton3 = inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonMiddle)

	// Get keyboard keys
	for key := range input.KeyboardKeys {
		input.KeyboardKeys[key] = ebiten.IsKeyPressed(key)
	}

	// Check specific keys
	input.Up = ebiten.IsKeyPressed(ebiten.KeyArrowUp)
	input.Down = ebiten.IsKeyPressed(ebiten.KeyArrowDown)
	input.Left = ebiten.IsKeyPressed(ebiten.KeyArrowLeft)
	input.Right = ebiten.IsKeyPressed(ebiten.KeyArrowRight)
	input.Enter = ebiten.IsKeyPressed(ebiten.KeyEnter)
	input.Escape = ebiten.IsKeyPressed(ebiten.KeyEscape)
	input.Space = ebiten.IsKeyPressed(ebiten.KeySpace)

	return input
}

func (ui *InputState) Update() {
	// Get mouse position
	ui.MouseX, ui.MouseY = ebiten.CursorPosition()

	// Get mouse buttons
	ui.MouseButton1 = inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft)
	ui.MouseButton2 = inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight)
	ui.MouseButton3 = inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonMiddle)

	// Get keyboard keys
	for key := range ui.KeyboardKeys {
		ui.KeyboardKeys[key] = ebiten.IsKeyPressed(key)
	}

	// Check specific keys
	ui.Up = ebiten.IsKeyPressed(ebiten.KeyArrowUp)
	ui.Down = ebiten.IsKeyPressed(ebiten.KeyArrowDown)
	ui.Left = ebiten.IsKeyPressed(ebiten.KeyArrowLeft)
	ui.Right = ebiten.IsKeyPressed(ebiten.KeyArrowRight)
	ui.Enter = ebiten.IsKeyPressed(ebiten.KeyEnter)
	ui.Escape = ebiten.IsKeyPressed(ebiten.KeyEscape)
	ui.Space = ebiten.IsKeyPressed(ebiten.KeySpace)
}
