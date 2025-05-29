package entity

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Letter struct {
	Sprite *ebiten.Image // Image of the letter
	State  LetterState // State of the letter (active, target, inactive)
}

// NewLetter creates a new Letter with the specified image and state.
func NewLetter(image *ebiten.Image, state LetterState) Letter {
	return Letter{
		Sprite: image,
		State:  state,
	}
}
