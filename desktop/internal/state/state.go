package state

type EngineState int

const (
	MENU_MAIN EngineState = iota
	GAME_PLAYING
	GAME_PAUSED
	GAME_OVER
)
