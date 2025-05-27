package main

import (
	"log"

	"td/internal/game"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	// VERSION is the current version of the application.
	VERSION = "0.1.0"
	// UI_SCALE is the scale factor for the UI.
	UI_SCALE = 1
	// UI_WIDTH is the width of the UI.
	UI_WIDTH = 640
	// UI_HEIGHT is the height of the UI.
	UI_HEIGHT = 480
	// UI_TITLE is the title of the window.
	UI_TITLE = "Hello, World!"
)

func init() {
	ebiten.SetTPS(120) // Set the target frames per second
}

func main() {
	screenWidth, screenHeight := ebiten.Monitor().Size()
	engine := game.NewEngine(screenWidth, screenHeight, UI_TITLE, VERSION)

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle(UI_TITLE)

	if err := ebiten.RunGame(engine); err != nil {
		log.Fatal(err)
	}
}
