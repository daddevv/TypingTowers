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

var (
	engine *game.Engine
)

func init() {
	// get screen size
	screenWidth, screenHeight := ebiten.Monitor().Size()
	engine = &game.Engine{
		Width:   screenWidth,
		Height:  screenHeight,
		Title:   UI_TITLE,
		Version: VERSION,
	}

	ebiten.SetTPS(100)
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle(UI_TITLE)
}

func main() {
	if err := ebiten.RunGame(engine); err != nil {
		log.Fatal(err)
	}
}
