package core

// GamePhase represents high level game states.
type GamePhase int

const (
	PhaseMainMenu GamePhase = iota
	PhasePreGame
	PhasePlaying
	PhasePaused
	PhaseGameOver
	PhaseSettings
)

func (p GamePhase) String() string {
	switch p {
	case PhaseMainMenu:
		return "MainMenu"
	case PhasePreGame:
		return "PreGame"
	case PhasePlaying:
		return "Playing"
	case PhasePaused:
		return "Paused"
	case PhaseGameOver:
		return "GameOver"
	case PhaseSettings:
		return "Settings"
	}
	return "Unknown"
}
