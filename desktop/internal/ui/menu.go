package ui

import "github.com/hajimehoshi/ebiten/v2"

type Menu interface {
	// Update processes input and updates the menu state.
	Update() (*MenuSelection, error)
	// Draw renders the menu to the screen.
	Draw(screen *ebiten.Image)
}
