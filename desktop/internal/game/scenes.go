package game

// enum CurrentScene
type CurrentScene int

const (
	MAINMENU CurrentScene = iota
	GAME
	OPTIONS
)

func NewCurrentScene() CurrentScene {
	return MAINMENU
}
