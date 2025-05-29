package entity

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Letter struct {
	Sprite    *ebiten.Image // Image of the letter
	State     LetterState   // State of the letter (active, target, inactive)
	Character rune          // The actual character this letter represents
}

// NewLetter creates a new Letter with the specified image, state, and character.
func NewLetter(image *ebiten.Image, state LetterState, character rune) Letter {
	return Letter{
		Sprite:    image,
		State:     state,
		Character: character,
	}
}
