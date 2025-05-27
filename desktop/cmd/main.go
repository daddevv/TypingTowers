package main

import (
	"log"

	"td/internal/engine"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	// VERSION is the current version of the application.
	VERSION = "0.1.0"
	// UI_TITLE is the title of the window.
	UI_TITLE = "Hello, World!"
)

func init() {
	ebiten.SetTPS(120) // Set the target frames per second
}

func main() {
	screenWidth, screenHeight := ebiten.Monitor().Size()
	engine := engine.NewGame(screenWidth, screenHeight, UI_TITLE, VERSION)

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle(UI_TITLE)

	if err := ebiten.RunGame(engine); err != nil {
		log.Fatal(err)
	}
}
