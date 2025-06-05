package game

// GamePhase represents high level game states.
type GamePhase int

const (
	PhaseMenu GamePhase = iota
	PhasePlaying
)
