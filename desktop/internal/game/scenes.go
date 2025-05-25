package game

// enum CurrentScene
type CurrentScene int

const (
	MAINMENU CurrentScene = iota
	GAME
	OPTIONS
	GAMEOVER
)

func NewCurrentScene() CurrentScene {
	return MAINMENU
}
