package state

type State interface {
	EngineState
	Update() (EngineState, error)
	Draw() error
}

type EngineState int

const (
	MAIN_MENU EngineState = iota
	GAME_MENU
	PLAYER_SETTINGS_MENU
	GAME_PLAYING
	GAME_PAUSED
	GAME_OVER
	GAME_SETTINGS_MENU
)
