package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Input  *InputState
	State  *State
	Menu   *Menu
	Screen *ebiten.Image
}

func NewGame() *Game {
	return &Game{
		Input: NewUserInput(),
		State: NewState(),
		Menu:  NewMenu(),
	}
}

func (e *Game) Update() error {
	// Update the game state here
	// For example, handle input, update game objects, etc.
	e.Input.Update()
	switch e.State.CurrentScreen {
	case "main-menu":
		e.Menu.Update(e.Input)
	case "game":
		// Handle game logic
	case "game-over":
		// Handle game over logic
	default:
		// Handle default logic
	}
	return nil
}

func (e *Game) Draw(screen *ebiten.Image) {
	// Draw the game state here
	// For example, draw game objects, UI, etc.
	// e.Screen.Clear()
	// e.Screen.DrawImage(...)
}
