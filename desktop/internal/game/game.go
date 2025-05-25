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
	
	return nil
}

func (e *Game) Draw(screen *ebiten.Image) {
	// Draw the game state here
	// For example, draw game objects, UI, etc.
	// e.Screen.Clear()
	// e.Screen.DrawImage(...)
}
