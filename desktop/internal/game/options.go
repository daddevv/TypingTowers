package game

import "td/internal/world"

type GameOptions struct {
	PossibleLetters []string
	Width           int
	Height          int
	Level           world.Level
	GameMode        GameMode
}
