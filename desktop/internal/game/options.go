package game

import "td/internal/world"

type GameOptions struct {
	Width           int
	Height          int
	Level           world.Level
	GameMode        GameMode
}
