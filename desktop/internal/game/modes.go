package game

// GameMode represents the different modes of play in the game engine.
type GameMode int

const (
	TRAINING GameMode = iota
	ENDLESS
	CHALLENGE
)
